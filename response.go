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
