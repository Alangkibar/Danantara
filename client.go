package dana

import (
	"context"
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
	HttpClient     *resty.Client
	BaseURL        *url.URL
	ResponseEntity interface{}
}

// NewClient creates a new DANA API client.
func NewClient(
	merchantID, clientID, clientSecret, publicKey, privateKey string, environment EnvironmentType,
) *Client {
	var baseURL string

	if merchantID == "" {
		log.Println("DANA: Merchant ID is required")
		return nil
	}
	MerchantID = merchantID

	if clientID == "" {
		log.Println("DANA: Client ID is required")
		return nil
	}
	ClientID = clientID

	if clientSecret == "" {
		log.Println("DANA: Client Secret is required")
		return nil
	}
	ClientSecret = clientSecret

	if publicKey == "" {
		log.Println("DANA: Public Key is required")
		return nil
	}
	PublicKey = publicKey

	if privateKey == "" {
		log.Println("DANA: Private Key is required")
		return nil
	}
	RsaPrivateKey = []byte(privateKey)

	client := resty.New()

	client.SetTimeout(DefaultRequestTimeout)
	client.SetHeader("Accept", "application/json")
	client.SetHeader("Content-Type", "application/json")

	if environment == Sandbox {
		baseURL = Sandbox.BaseUrl()
	} else if environment == Production {
		baseURL = Production.BaseUrl()
	} else {
		log.Println("DANA: Invalid environment")
		return nil
	}

	parsedBaseURL, err := url.Parse(baseURL)
	if err != nil {
		log.Println("DANA: Invalid URL:", err)
		return nil
	}

	return &Client{
		HttpClient: client,
		BaseURL:    parsedBaseURL,
	}
}

func (c *Client) DoRequest(ctx context.Context, method string, path string, body interface{}, headers map[string]string) (*ResponseAPI, error) {
	var resp *resty.Response
	var err error

	timestamp := GenerateTimestamp()
	// timestamp := "2025-04-21T13:08:00+07:00"

	bodyStringify, err := json.Marshal(body)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	signature, err := GetSignature(method, path, timestamp, string(bodyStringify))
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	// Begin to compose request client
	client := c.HttpClient

	client.SetDebug(false)

	req := client.R().
		SetHeader("X-TIMESTAMP", timestamp).
		SetHeader("X-SIGNATURE", signature).
		SetHeader("X-PARTNER-ID", ClientID).
		SetHeader("X-EXTERNAL-ID", time.Now().Format("02012006")).
		SetHeader("CHANNEL-ID", MerchantID)

	for key, value := range headers {
		req.SetHeader(key, value)
	}

	req.SetContext(ctx)

	req.SetResult(c.ResponseEntity)

	switch method {
	case http.MethodGet:
		resp, err = req.Get(c.BaseURL.ResolveReference(&url.URL{Path: path}).String())
	case http.MethodPost:
		if body != nil {
			req = req.SetBody(body)
		}
		resp, err = req.Post(c.BaseURL.ResolveReference(&url.URL{Path: path}).String())
	case http.MethodPut:
		if body != nil {
			req = req.SetBody(body)
		}
		resp, err = req.Put(c.BaseURL.ResolveReference(&url.URL{Path: path}).String())
	case http.MethodPatch:
		if body != nil {
			req = req.SetBody(body)
		}
		resp, err = req.Patch(c.BaseURL.ResolveReference(&url.URL{Path: path}).String())
	case http.MethodDelete:
		resp, err = req.Delete(c.BaseURL.ResolveReference(&url.URL{Path: path}).String())
	default:
		return nil, fmt.Errorf("unsupported method: %s", method)
	}

	// log.Println("==========")
	// log.Println("========== DANA API DEBUG (", path, ") ==========")
	// log.Println("===== REQUEST =====")
	// log.Printf("Request Details:\nURL Endpoint: %s\nHeader Request: %s\nRequest Body: %s",
	// 	c.BaseURL.ResolveReference(&url.URL{Path: path}).String(),
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

		code := danaErr.Code
		return &ResponseAPI{
			StatusCode:   resp.StatusCode(),
			RawBody:      resp.Body(),
			Error:        &danaErr,
			ResponseCode: code,
			CodeInfo:     GetDanaResponseInfo(code),
		}, &danaErr
	}

	return response, nil
}
