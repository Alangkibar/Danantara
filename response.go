package dana

// RequestCustomerTopUpDisbursement: Customer Top Up Disbursement
type ResponseCustomerTopUpDisbursementAmount struct {
	Value    string `json:"value"`
	Currency string `json:"currency"`
}

type ResponseCustomerTopUpDisbursement struct {
	Code               string                                  `json:"responseCode"`
	Message            string                                  `json:"responseMessage"`
	ReferenceNo        string                                  `json:"referenceNo"`
	PartnerReferenceNo string                                  `json:"partnerReferenceNo"`
	SessionId          string                                  `json:"sessionId"`
	CustomerNumber     string                                  `json:"customerNumber"`
	Amount             ResponseCustomerTopUpDisbursementAmount `json:"amount"`
	AdditionalInfo     map[string]interface{}                  `json:"additionalInfo"`
}

type ResponseAccountInquiryTopUpDisbursement struct {
	Code                   string                                  `json:"responseCode"`
	Message                string                                  `json:"responseMessage"`
	PartnerReferenceNo     string                                  `json:"partnerReferenceNo"`
	SessionId              string                                  `json:"sessionId"`
	CustomerNumber         string                                  `json:"customerNumber"`
	CustomerName           string                                  `json:"customerName"`
	FeeType                string                                  `json:"feeType"`
	Amount                 ResponseCustomerTopUpDisbursementAmount `json:"amount"`
	FeeAmount              ResponseCustomerTopUpDisbursementAmount `json:"feeAmount"`
	CustomerMonthlyInLimit string                                  `json:"customerMonthlyInLimit"`
	MinAmount              ResponseCustomerTopUpDisbursementAmount `json:"minAmount"`
	MaxAmount              ResponseCustomerTopUpDisbursementAmount `json:"maxAmount"`
}

type ResponseCustomerTopUpInquiryStatusDisbursement struct {
	ResponseCode            string                                  `json:"responseCode"`                  // Mandatory, max 7 characters
	ResponseMessage         string                                  `json:"responseMessage"`               // Mandatory, max 150 characters
	OriginalPartnerRefNo    string                                  `json:"originalPartnerReferenceNo"`    // Mandatory, max 64 characters
	OriginalReferenceNo     string                                  `json:"originalReferenceNo,omitempty"` // Optional, max 64 characters
	OriginalExternalId      string                                  `json:"originalExternalId,omitempty"`  // Optional, max 36 characters
	ServiceCode             string                                  `json:"serviceCode"`                   // Mandatory, always "38"
	Amount                  ResponseCustomerTopUpDisbursementAmount `json:"amount"`                        // Mandatory
	LatestTransactionStatus string                                  `json:"latestTransactionStatus"`       // Mandatory, fixed 2 chars
	TransactionStatusDesc   string                                  `json:"transactionStatusDesc"`         // Mandatory, max 50 characters
	AdditionalInfo          map[string]interface{}                  `json:"additionalInfo,omitempty"`      // Optional, JSON object
}

type ResponsePaymentGatewayDropInCreateOrder struct {
	ResponseCode       string                           `json:"responseCode"`             // Mandatory, max 7
	ResponseMessage    string                           `json:"responseMessage"`          // Mandatory, max 150
	ReferenceNo        string                           `json:"referenceNo,omitempty"`    // Conditional, max 64
	PartnerReferenceNo string                           `json:"partnerReferenceNo"`       // Mandatory, max 64
	WebRedirectURL     string                           `json:"webRedirectUrl,omitempty"` // Conditional, max 2048
	AdditionalInfo     PaymentGatewayAdditionalResponse `json:"additionalInfo,omitempty"` // Optional object, with conditional field
}

type PaymentGatewayAdditionalResponse struct {
	PaymentCode *string `json:"paymentCode,omitempty"` // Conditional, max 64
}

type ResponseAPI struct {
	StatusCode   int
	RawBody      []byte
	Result       interface{}
	Error        error
	ResponseCode string
	CodeInfo     ResponseCodeInfo
}

type ResponseCodeInfo struct {
	Code          string
	Message       string
	Description   string
	PartnerAction string
}
