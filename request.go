package dana

// RequestCustomerTopUpDisbursement: Customer Top Up Disbursement
type RequestCustomerTopUpDisbursement struct {
	PartnerReferenceNo string
	CustomerNumber     string
	Amount             Currency
	FeeAmount          Currency
	AdditionalInfo     RequestCustomerTopUpDisbursementAdditionalInfo
}

type requestCustomerTopUpDisbursement struct {
	PartnerReferenceNo string                                         `json:"partnerReferenceNo"`
	CustomerNumber     string                                         `json:"customerNumber"`
	Amount             Currency                                       `json:"amount"`
	FeeAmount          Currency                                       `json:"feeAmount"`
	TransactionDate    string                                         `json:"transactionDate"`
	AdditionalInfo     RequestCustomerTopUpDisbursementAdditionalInfo `json:"additionalInfo"`
}

type RequestCustomerTopUpDisbursementAdditionalInfo struct {
	FundType string `json:"fundType"`
}

// AccountRequestAccountInquiryTopUpDisbursement : Inquiry Account Top Up Disbursement
type RequestAccountInquiryTopUpDisbursement RequestCustomerTopUpDisbursement

type RequestCustomerTopUpInquiryStatusDisbursement struct {
	OriginalPartnerReferenceNo string                 `json:"originalPartnerReferenceNo"`    // Mandatory, max 64 characters
	OriginalReferenceNo        string                 `json:"originalReferenceNo,omitempty"` // Optional, max 64 characters
	OriginalExternalId         string                 `json:"originalExternalId,omitempty"`  // Optional, max 36 characters
	ServiceCode                string                 `json:"serviceCode"`                   // Mandatory, always "38"
	AdditionalInfo             map[string]interface{} `json:"additionalInfo,omitempty"`
}

type RequestPaymentGatewayDropInCreateOrder struct {
	PartnerReferenceNo string                       `json:"partnerReferenceNo"`           // Mandatory, max 64
	SubMerchantId      *string                      `json:"subMerchantId,omitempty"`      // Optional, max 32
	Amount             Currency                     `json:"amount"`                       // Mandatory
	ExternalStoreId    *string                      `json:"externalStoreId,omitempty"`    // Optional, max 64
	ValidUpTo          *string                      `json:"validUpTo,omitempty"`          // Optional, max 25, format: YYYY-MM-DDTHH:mm:ss+07:00
	DisabledPayMethods *string                      `json:"disabledPayMethods,omitempty"` // Optional, max 64
	URLParams          []PaymentGatewayNotifyURL    `json:"urlParams"`                    // Mandatory
	PayOptionDetails   *PaymentGatewayPayOption     `json:"payOptionDetails,omitempty"`   // Conditional
	AdditionalInfo     PaymentGatewayAdditionalInfo `json:"additionalInfo"`               // Optional
}

type requestPaymentGatewayDropInCreateOrder struct {
	PartnerReferenceNo string                       `json:"partnerReferenceNo"`           // Mandatory, max 64
	MerchantId         string                       `json:"merchantId"`                   // Mandatory, max 64
	SubMerchantId      *string                      `json:"subMerchantId,omitempty"`      // Optional, max 32
	Amount             Currency                     `json:"amount"`                       // Mandatory
	ExternalStoreId    *string                      `json:"externalStoreId,omitempty"`    // Optional, max 64
	ValidUpTo          *string                      `json:"validUpTo,omitempty"`          // Optional, max 25, format: YYYY-MM-DDTHH:mm:ss+07:00
	DisabledPayMethods *string                      `json:"disabledPayMethods,omitempty"` // Optional, max 64
	URLParams          []PaymentGatewayNotifyURL    `json:"urlParams"`                    // Mandatory
	PayOptionDetails   *PaymentGatewayPayOption     `json:"payOptionDetails,omitempty"`   // Conditional
	AdditionalInfo     PaymentGatewayAdditionalInfo `json:"additionalInfo"`               // Optional
}

type PaymentGatewayAmount struct {
	Value    string `json:"value"`    // Value of the transaction
	Currency string `json:"currency"` // ISO currency code
}

type Currency struct {
	Value    string `json:"value"`
	Currency string `json:"currency"`
}

func NewCurrency(value string, currency string) Currency {
	return Currency{
		Value:    value,
		Currency: currency,
	}
}

func NewIDRCurrency(value string) Currency {
	return NewCurrency(value, "IDR")
}

type PaymentGatewayNotifyURL struct {
	// Define actual notify URL fields based on expected schema
	URL        string `json:"url"`        // Mandatory
	Type       string `json:"type"`       // Mandatory
	IsDeeplink string `json:"isDeeplink"` // Mandatory
}

type PaymentGatewayPayOption struct {
	// Define actual pay option fields based on expected schema
	PayMethod string `json:"payMethod"`
}

type PaymentGatewayAdditionalInfo struct {
	Order      *PaymentGatewayOrderInfo `json:"order,omitempty"`
	MCC        string                   `json:"mcc"`                  // Mandatory
	ExtendInfo *string                  `json:"extendInfo,omitempty"` // Optional, max 4096
	EnvInfo    PaymentGatewayEnvInfo    `json:"envInfo"`              // Mandatory
}

type PaymentGatewayOrderInfo struct {
	OrderTitle        string                      `json:"orderTitle"`                  // Mandatory, max 64
	Scenario          string                      `json:"scenario"`                    // Mandatory, max 64
	MerchantTransType *string                     `json:"merchantTransType,omitempty"` // Optional, max 64
	Buyer             PaymentGatewayBuyerInfo     `json:"buyer"`                       // Mandatory
	Goods             []PaymentGatewayGoodsInfo   `json:"goods,omitempty"`             // Optional
	ShippingInfo      *PaymentGatewayShippingInfo `json:"shippingInfo,omitempty"`      // Optional
	ExtendInfo        *string                     `json:"extendInfo,omitempty"`        // Optional, max 4096
}

type PaymentGatewayBuyerInfo struct {
	// Define buyer info fields
	ExternalUserID string `json:"externalUserId"` // Mandatory
}

type PaymentGatewayGoodsInfo struct {
	// Define goods info fields
	Unit            string   `json:"unit"`
	Category        string   `json:"category"`
	Description     string   `json:"description"`
	Quantity        string   `json:"quantity"`
	Price           Currency `json:"price"`
	MerchantGoodsId string   `json:"merchantGoodsId"`
}

type PaymentGatewayShippingInfo struct {
	// Define shipping info fields
	Address string `json:"address"`
	City    string `json:"city"`
	ZipCode string `json:"zipCode"`
}

type PaymentGatewayEnvInfo struct {
	// Define environment info fields
	OsType            string `json:"osType"`
	DeviceId          string `json:"deviceId"`
	AppVersion        string `json:"appVersion"`
	SourcePlatform    string `json:"sourcePlatform"`
	OrderTerminalType string `json:"orderTerminalType"`
	TerminalType      string `json:"terminalType"`
}
