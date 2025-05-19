package dana

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

// Client holds the base URL and any necessary authentication tokens.
type Client struct {
	httpClient   *resty.Client
	merchantID   string
	clientID     string
	clientSecret string
	publicKey    string
	privateKey   []byte
	location     *time.Location
	debug        bool
}

type ClientConfig struct {
	MerchantID   string
	ClientID     string
	ClientSecret string
	PublicKey    string
	PrivateKey   []byte
	UseSandbox   bool
	Debug        bool
}

// NewClient creates a new DANA API client.
func NewClient(conf ClientConfig) (*Client, error) {
	if conf.MerchantID == "" {
		return nil, fmt.Errorf("merchant ID is required")
	}

	if conf.ClientID == "" {
		return nil, fmt.Errorf("clientID is required")
	}

	if conf.ClientSecret == "" {
		return nil, fmt.Errorf("client secret is required")
	}

	if conf.PublicKey == "" {
		return nil, fmt.Errorf("public key is required")
	}

	if conf.PrivateKey == nil {
		return nil, fmt.Errorf("private key is required")
	}

	client := resty.New()

	client.SetTimeout(defaultRequestTimeout)
	client.SetHeader("Accept", "application/json")
	client.SetHeader("Content-Type", "application/json")

	baseURL := Sandbox.BaseUrl()
	if !conf.UseSandbox {
		baseURL = Production.BaseUrl()
	}
	client.SetBaseURL(baseURL)

	// Get the current time in Jakarta time zone (UTC+7)
	jakartaLoc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		log.Panicf("Error loading Jakarta time zone: %v", err)
	}

	return &Client{
		httpClient:   client,
		merchantID:   conf.MerchantID,
		clientID:     conf.ClientID,
		clientSecret: conf.ClientSecret,
		publicKey:    conf.PublicKey,
		privateKey:   conf.PrivateKey,
		location:     jakartaLoc,
		debug:        conf.Debug,
	}, nil
}

func (c *Client) generateTimestamp() string {
	currentTime := time.Now().In(c.location)

	// Format the time according to the specified format
	timestamp := currentTime.Format("2006-01-02T15:04:05+07:00")

	return timestamp
}

func (c *Client) prepareHttpRequest() *resty.Request {
	client := c.httpClient.SetDebug(c.debug).R().
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/json").
		SetHeader("X-PARTNER-ID", c.clientID).
		SetHeader("X-EXTERNAL-ID", time.Now().Format("02012006")).
		SetHeader("CHANNEL-ID", c.merchantID)

	return client
}

func (c *Client) prepareHeaders(method, path string, body interface{}) (map[string]string, error) {
	timestamp := c.generateTimestamp()
	bodyStringify, err := json.Marshal(body)
	if err != nil {
		if c.debug {
			fmt.Println("Error bodyStringify:", err)
		}
		return nil, fmt.Errorf("error dana client: %w", err)
	}

	signature, err := c.getSignature(method, path, timestamp, string(bodyStringify))
	if err != nil {
		if c.debug {
			fmt.Println("Error getSignature: ", err)
		}
		return nil, fmt.Errorf("error dana client: %w", err)
	}

	headers := make(map[string]string)
	headers["X-TIMESTAMP"] = timestamp
	headers["X-SIGNATURE"] = signature
	headers["X-PARTNER-ID"] = c.clientID
	headers["X-EXTERNAL-ID"] = time.Now().Format("02012006")
	headers["CHANNEL-ID"] = c.merchantID

	return headers, nil
}

func (c *Client) getSignature(method string, path string, timestamp string, body string) (string, error) {
	// Hash the minified body
	hash := sha256.New()
	hash.Write([]byte(body))
	hashedPayload := fmt.Sprintf("%x", hash.Sum(nil)) // hex encode and lowercase

	// data = '<HTTP METHOD> + ”:” + <RELATIVE PATH URL> + “:“ + LowerCase(HexEncode(SHA-256(Minify(<HTTP BODY>)))) + “:“ + <X-TIMESTAMP>';
	data := method + ":" + path + ":" + hashedPayload + ":" + timestamp

	signature, err := generateSignature([]byte(data), c.privateKey)
	if err != nil {
		fmt.Println("signature generation error", err)
		return "", err
	}

	return signature, nil
}

func wrapError(resp *resty.Response) error {
	var danaErr Error
	if err := json.Unmarshal(resp.Body(), &danaErr); err != nil {
		if resp.StatusCode() == http.StatusGatewayTimeout {
			return &Error{
				Code:    danaErr.Code,
				Message: danaErr.Message,
			}
		}

		return err
	}

	return &Error{
		Code:    danaErr.Code,
		Message: danaErr.Message,
	}
}
