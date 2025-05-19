package dana

import (
	"context"
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
	defaultRequestTimeout = 80 * time.Second

	// API Paths
	paymentGatewayDropInCreateOrderPath = "/v1.0/payment-gateway/payment.htm"
	accountInquiryTopUpDisbursementPath = "/v1.0/emoney/account-inquiry.htm"
	customerTopUpDisbursementPath       = "/v1.0/emoney/topup.htm"
	customerTopUpInquiryStatusPath      = "/v1.0/emoney/topup-status.htm"
)

var (
	Environment = Sandbox
)

var RsaPrivateKey []byte

type Dana interface {
	// Account Inquiry (Disbursement)
	AccountInquiryTopUpDisbursement(ctx context.Context, req RequestAccountInquiryTopUpDisbursement) (*ResponseAccountInquiryTopUpDisbursement, error)

	// Customer Top Up (Disbursement)
	CustomerTopUpDisbursement(ctx context.Context, req RequestCustomerTopUpDisbursement) (*ResponseCustomerTopUpDisbursement, error)

	// Customer Top Up Inquiry Status (Disbursement)
	CustomerTopUpInquiryStatusDisbursement(ctx context.Context, req RequestCustomerTopUpInquiryStatusDisbursement) (*ResponseCustomerTopUpInquiryStatusDisbursement, error)

	// Payment Gateway Drop-In Create Order
	PaymentGatewayDropInCreateOrder(ctx context.Context, req RequestPaymentGatewayDropInCreateOrder) (*ResponsePaymentGatewayDropInCreateOrder, error)
}

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
	headers, err := c.prepareHeaders(http.MethodPost, accountInquiryTopUpDisbursementPath, body)
	if err != nil {
		return nil, err
	}
	for key, value := range headers {
		httpRequest.SetHeader(key, value)
	}

	httpRequest.SetBody(body)
	httpRequest.SetResult(&ResponseAccountInquiryTopUpDisbursement{})
	httpRequest.SetContext(ctx)

	resp, err := httpRequest.Post(accountInquiryTopUpDisbursementPath)
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

	return resp.Result().(*ResponseAccountInquiryTopUpDisbursement), nil
}

// Refer to this documentation: https://dashboard.dana.id/api-docs/read/119
func (c *Client) CustomerTopUpDisbursement(ctx context.Context, req RequestCustomerTopUpDisbursement) (*ResponseCustomerTopUpDisbursement, error) {
	httpRequest := c.prepareHttpRequest()
	body := requestCustomerTopUpDisbursement{
		PartnerReferenceNo: req.PartnerReferenceNo,
		CustomerNumber:     req.CustomerNumber,
		Amount:             req.Amount,
		FeeAmount:          req.FeeAmount,
		TransactionDate:    c.generateTimestamp(),
		AdditionalInfo:     req.AdditionalInfo,
	}

	headers, err := c.prepareHeaders(http.MethodPost, customerTopUpDisbursementPath, body)
	if err != nil {
		return nil, err
	}
	for key, value := range headers {
		httpRequest.SetHeader(key, value)
	}

	httpRequest.SetBody(body)
	httpRequest.SetResult(&ResponseCustomerTopUpDisbursement{})
	httpRequest.SetContext(ctx)

	resp, err := httpRequest.Post(customerTopUpDisbursementPath)
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

	return resp.Result().(*ResponseCustomerTopUpDisbursement), nil
}

// Refer to this documentation: https://dashboard.dana.id/api-docs/read/121
func (c *Client) CustomerTopUpInquiryStatusDisbursement(ctx context.Context, req RequestCustomerTopUpInquiryStatusDisbursement) (*ResponseCustomerTopUpInquiryStatusDisbursement, error) {
	httpRequest := c.prepareHttpRequest()

	headers, err := c.prepareHeaders(http.MethodPost, customerTopUpInquiryStatusPath, req)
	if err != nil {
		return nil, err
	}
	for key, value := range headers {
		httpRequest.SetHeader(key, value)
	}

	httpRequest.SetBody(req)
	httpRequest.SetResult(&ResponseCustomerTopUpInquiryStatusDisbursement{})
	httpRequest.SetContext(ctx)

	resp, err := httpRequest.Post(customerTopUpInquiryStatusPath)
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

	return resp.Result().(*ResponseCustomerTopUpInquiryStatusDisbursement), nil
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
		if c.debug {
			fmt.Println("Error prepare headers: ", err)
		}
		return nil, err
	}
	for key, value := range headers {
		httpRequest.SetHeader(key, value)
	}

	httpRequest.SetBody(body)

	httpRequest.SetResult(&ResponsePaymentGatewayDropInCreateOrder{})
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

	return resp.Result().(*ResponsePaymentGatewayDropInCreateOrder), nil
}
