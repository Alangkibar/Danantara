package dana

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type EnvironmentType int8

const (
	_ EnvironmentType = iota

	// Sandbox : represent sandbox environment
	Sandbox

	// Production : represent production environment
	Production

	// libraryVersion : dana go library version
	libraryVersion = "v1.0.0"
)

var typeString = map[EnvironmentType]string{
	Sandbox:    "https://api.sandbox.dana.id",
	Production: "https://api.dana.id",
}

const (
	defaultRequestTimeout               = 80 * time.Second
	paymentGatewayDropInCreateOrderPath = "/v1.0/payment-gateway/payment.htm"
	accountInquiryTopUpDisbursementPath = "/v1.0/emoney/account-inquiry.htm"
)

var (
	Environment = Sandbox
)

var RsaPrivateKey []byte

// Refer to this documentation: https://dashboard.dana.id/api-docs/read/107
func (c *Client) AccountInquiryTopUpDisbursement(ctx context.Context, req RequestAccountInquiryTopUpDisbursement) (*ResponseAccountInquiryTopUpDisbursement, error) {
	httpRequest := c.prepareHttpRequest()
	body := requestCustomerTopUpDisbursement{
		PartnerReferenceNo: req.PartnerReferenceNo,
		CustomerNumber:     req.CustomerNumber,
		Amount:             req.Amount,
		FeeAmount:          req.FeeAmount,
		TransactionDate:    c.generateTimestamp(),
		AdditionalInfo:     req.AdditionalInfo,
	}
	headers, err := c.prepareHeaders(http.MethodPost, paymentGatewayDropInCreateOrderPath, body)
	if err != nil {
		return nil, err
	}
	for key, value := range headers {
		httpRequest.SetHeader(key, value)
	}

	var result ResponsePaymentGatewayDropInCreateOrder
	httpRequest.SetResult(result)
	httpRequest.SetContext(ctx)

	resp, err := httpRequest.Post(paymentGatewayDropInCreateOrderPath)
	if err != nil {
		return nil, &Error{
			Code:     "ERR000",
			Message:  "error",
			RawError: fmt.Errorf("error dana client: %w", err),
		}
	}

	if resp.IsError() {
		return nil, wrapError(resp)
	}

	return &result, nil
}

// Refer to this documentation: https://dashboard.dana.id/api-docs/read/119
func (c *Client) CustomerTopUpDisbursement(ctx context.Context, req RequestCustomerTopUpDisbursement) (*ResponseCustomerTopUpDisbursement, error) {
	var result ResponseCustomerTopUpDisbursement
	c.ResponseEntity = &result

	apiResp, err := c.doRequest(ctx, http.MethodPost, "/v1.0/emoney/topup.htm", req, nil)

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

	apiResp, err := c.doRequest(ctx, http.MethodPost, "/v1.0/emoney/topup-status.htm", req, nil)

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
	httpRequest := c.prepareHttpRequest()
	body := requestPaymentGatewayDropInCreateOrder{
		PartnerReferenceNo: req.PartnerReferenceNo,
		MerchantId:         c.merchantID,
		SubMerchantId:      req.SubMerchantId,
		Amount:             req.Amount,
		ExternalStoreId:    req.ExternalStoreId,
		ValidUpTo:          req.ValidUpTo,
		DisabledPayMethods: req.DisabledPayMethods,
		URLParams:          req.URLParams,
		PayOptionDetails:   req.PayOptionDetails,
		AdditionalInfo:     req.AdditionalInfo,
	}
	headers, err := c.prepareHeaders(http.MethodPost, paymentGatewayDropInCreateOrderPath, body)
	if err != nil {
		return nil, err
	}
	for key, value := range headers {
		httpRequest.SetHeader(key, value)
	}

	var result ResponsePaymentGatewayDropInCreateOrder
	httpRequest.SetResult(result)
	httpRequest.SetContext(ctx)

	resp, err := httpRequest.Post(paymentGatewayDropInCreateOrderPath)
	if err != nil {
		return nil, &Error{
			Code:     "ERR000",
			Message:  "error",
			RawError: fmt.Errorf("error dana client: %w", err),
		}
	}

	if resp.IsError() {
		return nil, wrapError(resp)
	}

	return &result, nil
}
