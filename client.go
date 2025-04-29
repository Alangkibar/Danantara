package dana

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
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
	}, nil
}

// Deprecated
func (c *Client) doRequest(ctx context.Context, method string, path string, body interface{}, headers map[string]string) (*ResponseAPI, error) {
	var resp *resty.Response
	var err error

	timestamp := c.generateTimestamp()
	bodyStringify, err := json.Marshal(body)
	if err != nil {
		if c.debug {
			fmt.Println("Error:", err)
		}
		return nil, fmt.Errorf("error dana client: %w", err)
	}

	signature, err := getSignature(method, path, timestamp, string(bodyStringify))
	if err != nil {
		fmt.Println("Error:", err)
		return nil, fmt.Errorf("error dana client: %w", err)
	}

	// Begin to compose request client
	client := c.httpClient

	client.SetDebug(false)

	req := client.R().
		SetHeader("X-TIMESTAMP", timestamp).
		SetHeader("X-SIGNATURE", signature).
		SetHeader("X-PARTNER-ID", c.clientID).
		SetHeader("X-EXTERNAL-ID", time.Now().Format("02012006")).
		SetHeader("CHANNEL-ID", c.merchantID)

	for key, value := range headers {
		req.SetHeader(key, value)
	}

	req.SetContext(ctx)

	req.SetResult(c.ResponseEntity)

	switch method {
	case http.MethodGet:
		resp, err = req.Get(c.baseURL.ResolveReference(&url.URL{Path: path}).String())
	case http.MethodPost:
		if body != nil {
			req = req.SetBody(body)
		}
		resp, err = req.Post(c.baseURL.ResolveReference(&url.URL{Path: path}).String())
	case http.MethodPut:
		if body != nil {
			req = req.SetBody(body)
		}
		resp, err = req.Put(c.baseURL.ResolveReference(&url.URL{Path: path}).String())
	case http.MethodPatch:
		if body != nil {
			req = req.SetBody(body)
		}
		resp, err = req.Patch(c.baseURL.ResolveReference(&url.URL{Path: path}).String())
	case http.MethodDelete:
		resp, err = req.Delete(c.baseURL.ResolveReference(&url.URL{Path: path}).String())
	default:
		return nil, fmt.Errorf("unsupported method: %s", method)
	}

	// log.Println("==========")
	// log.Println("========== DANA API DEBUG (", path, ") ==========")
	// log.Println("===== REQUEST =====")
	// log.Printf("Request Details:\nURL Endpoint: %s\nHeader Request: %s\nRequest Body: %s",
	// 	c.baseURL.ResolveReference(&url.URL{Path: path}).String(),
	// 	func() string {
	// 		flatHeaders := make(map[string]string)
	// 		for key, values := range req.Header {
	// 			if len(values) > 0 {
	// 				flatHeaders[key] = values[0]
	// 			}
	// 		}
	// 		headerJSON, err := json.Marshal(flatHeaders)
	// 		if err != nil {
	// 			return fmt.Sprintf("Error marshaling headers to JSON: %v", err)
	// 		}
	// 		return string(headerJSON)
	// 	}(),
	// 	string(bodyStringify),
	// )
	// log.Println("===== RESPONSE =====")
	// log.Println("Response Body:", string(resp.Body()))
	// log.Println("==========")

	// Base response wrapper
	response := &ResponseAPI{
		StatusCode: resp.StatusCode(),
		RawBody:    resp.Body(),
		Result:     resp.Result(),
	}

	// Handle request error
	if err != nil {
		response.Error = err
		return response, err
	}

	if resp.IsError() {
		var danaErr Error
		if err := json.Unmarshal(resp.Body(), &danaErr); err != nil {

			if resp.StatusCode() == http.StatusGatewayTimeout {
				return &ResponseAPI{
					StatusCode:   resp.StatusCode(),
					RawBody:      resp.Body(),
					Error:        fmt.Errorf("failed to decode error body: %w", err),
					ResponseCode: "UNEXPECTED",
					CodeInfo:     GetDanaResponseInfo("UNEXPECTED"),
				}, context.DeadlineExceeded
			}

			return &ResponseAPI{
				StatusCode:   resp.StatusCode(),
				RawBody:      resp.Body(),
				Error:        fmt.Errorf("failed to decode error body: %w", err),
				ResponseCode: "UNEXPECTED",
				CodeInfo:     GetDanaResponseInfo("UNEXPECTED"),
			}, err
		}

		// code := danaErr.Code
		// return &ResponseAPI{
		// 	StatusCode:   resp.StatusCode(),
		// 	RawBody:      resp.Body(),
		// 	Error:        &danaErr,
		// 	ResponseCode: code,
		// 	CodeInfo:     GetDanaResponseInfo(code),
		// }, &danaErr
		return nil, &Error{
			Code:    danaErr.Code,
			Message: danaErr.Message,
		}
	}

	return response, nil
}

func (c *Client) generateTimestamp() string {
	currentTime := time.Now().In(c.location)

	// Format the time according to the specified format
	timestamp := currentTime.Format("2006-01-02T15:04:05+07:00")

	return timestamp
}

func (c *Client) prepareHttpRequest() *resty.Request {
	client := c.httpClient.R().
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
			fmt.Println("Error:", err)
		}
		return nil, fmt.Errorf("error dana client: %w", err)
	}

	signature, err := c.getSignature(method, path, timestamp, string(bodyStringify))
	if err != nil {
		fmt.Println("Error:", err)
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
