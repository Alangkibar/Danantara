package dana

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

type EnvironmentType int8

const (
	_ EnvironmentType = iota

	//Sandbox : represent sandbox environment
	Sandbox

	//Production : represent production environment
	Production

	//libraryVersion : dana go library version
	libraryVersion = "v1.0.0"
)

var typeString = map[EnvironmentType]string{
	Sandbox:    "https://api.sandbox.dana.id",
	Production: "https://api.dana.id",
}

var (
	Environment           = Sandbox
	DefaultRequestTimeout = 80 * time.Second
)

var RsaPrivateKey []byte

var MerchantID string

var PublicKey string

var ClientID string

var ClientSecret string

// Refer to this documentation: https://dashboard.dana.id/api-docs/read/107
func (c *Client) AccountInquiryTopUpDisbursement(ctx context.Context, req RequestAccountInquiryTopUpDisbursement) (*ResponseAccountInquiryTopUpDisbursement, error) {
	var result ResponseAccountInquiryTopUpDisbursement
	c.ResponseEntity = &result

	apiResp, err := c.DoRequest(ctx, http.MethodPost, "/v1.0/emoney/account-inquiry.htm", req, nil)

	if apiResp.RawBody != nil {
		if err := json.Unmarshal(apiResp.RawBody, &result); err != nil {
			return &result, err
		}
	}

	if err != nil {
		return &result, err
	}

	return &result, nil
}

// Refer to this documentation: https://dashboard.dana.id/api-docs/read/119
func (c *Client) CustomerTopUpDisbursement(ctx context.Context, req RequestCustomerTopUpDisbursement) (*ResponseCustomerTopUpDisbursement, error) {
	var result ResponseCustomerTopUpDisbursement
	c.ResponseEntity = &result

	apiResp, err := c.DoRequest(ctx, http.MethodPost, "/v1.0/emoney/topup.htm", req, nil)

	if apiResp.RawBody != nil {
		if err := json.Unmarshal(apiResp.RawBody, &result); err != nil {
			return &result, err
		}
	}

	if err != nil {
		return &result, err
	}

	return &result, nil
}

// Refer to this documentation: https://dashboard.dana.id/api-docs/read/121
func (c *Client) CustomerTopUpInquiryStatusDisbursement(ctx context.Context, req RequestCustomerTopUpInquiryStatusDisbursement) (*ResponseCustomerTopUpInquiryStatusDisbursement, error) {
	var result ResponseCustomerTopUpInquiryStatusDisbursement
	c.ResponseEntity = &result

	apiResp, err := c.DoRequest(ctx, http.MethodPost, "/v1.0/emoney/topup-status.htm", req, nil)

	if apiResp.RawBody != nil {
		if err := json.Unmarshal(apiResp.RawBody, &result); err != nil {
			return &result, err
		}
	}

	if err != nil {
		return &result, err
	}

	return &result, nil
}

// Refer to this documentation: https://dashboard.dana.id/api-docs/read/243
func (c *Client) PaymentGatewayDropInCreateOrder(ctx context.Context, req RequestPaymentGatewayDropInCreateOrder) (*ResponsePaymentGatewayDropInCreateOrder, error) {
	var result ResponsePaymentGatewayDropInCreateOrder
	c.ResponseEntity = &result

	apiResp, err := c.DoRequest(ctx, http.MethodPost, "/v1.0/payment-gateway/payment.htm", req, nil)

	if apiResp.RawBody != nil {
		if err := json.Unmarshal(apiResp.RawBody, &result); err != nil {
			return &result, err
		}
	}

	if err != nil {
		return &result, err
	}

	return &result, nil
}
