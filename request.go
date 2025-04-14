package dana

// RequestCustomerTopUpDisbursement: Customer Top Up Disbursement
type RequestCustomerTopUpDisbursement struct {
	PartnerReferenceNo string                                         `json:"partnerReferenceNo"`
	CustomerNumber     string                                         `json:"customerNumber"`
	Amount             RequestCustomerTopUpDisbursementAmount         `json:"amount"`
	FeeAmount          RequestCustomerTopUpDisbursementFeeAmount      `json:"feeAmount"`
	TransactionDate    string                                         `json:"transactionDate"`
	AdditionalInfo     RequestCustomerTopUpDisbursementAdditionalInfo `json:"additionalInfo"`
}

type RequestCustomerTopUpDisbursementAmount struct {
	Value    string `json:"value"`
	Currency string `json:"currency"`
}

type RequestCustomerTopUpDisbursementFeeAmount struct {
	Value    string `json:"value"`
	Currency string `json:"currency"`
}

type RequestCustomerTopUpDisbursementAdditionalInfo struct {
	FundType string `json:"fundType"`
}

// AccountRequestAccountInquiryTopUpDisbursement : Inquiry Account Top Up Disbursement
type RequestAccountInquiryTopUpDisbursement RequestCustomerTopUpDisbursement
