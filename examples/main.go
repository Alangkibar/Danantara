package main

import (
	"context"
	"fmt"
	"log"

	"github.com/alangkibar/dana"
	"github.com/google/uuid"
)

var (
	MerchantID   = "216620020002032273134"
	ClientID     = "2024122414041437752652"
	ClientSecret = "c54b954811eea0babc480e221120ebfe18e339ae31dde3aaed8a605df0470f8b"
	PublicKey    = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAj8cBR7T4nxIhhQqvvNYERU0Of83EIBzOVc2SWpwRCP6Ro4TrfoaOYSaTO4U1OoIYUfr0EQwBfRpaBapSk43vRLLjvjvBe8JOUHLw5uoujCIZnEQgDxqWMf/oYZ3+ruk1NNdxj6V8GviwDX782077upAw5X+EQOlAXgptBql049fSVn2xOaXSYbPIn4z4hrg/x5cinlopPj7KZGg5ixFGMcQvMR6iqHjFg1CaoDkDP3LKoQdSMhvk1MXeUFQKsLftJIDiiCTY0VqO3K3nuML2olGyBwtVr74QkaBo3bWJ7MidYNO8khXPLsJk+Elj87sNSDYzxrXHrCdgfsBO7aeAPQIDAQAB"
	PrivateKey   = "-----BEGIN RSA PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCPxwFHtPifEiGF\nCq+81gRFTQ5/zcQgHM5VzZJanBEI/pGjhOt+ho5hJpM7hTU6ghhR+vQRDAF9GloF\nqlKTje9EsuO+O8F7wk5QcvDm6i6MIhmcRCAPGpYx/+hhnf6u6TU013GPpXwa+LAN\nfvzbTvu6kDDlf4RA6UBeCm0GqXTj19JWfbE5pdJhs8ifjPiGuD/HlyKeWik+Pspk\naDmLEUYxxC8xHqKoeMWDUJqgOQM/csqhB1IyG+TUxd5QVAqwt+0kgOKIJNjRWo7c\nree4wvaiUbIHC1WvvhCRoGjdtYnsyJ1g07ySFc8uwmT4SWPzuw1INjPGtcesJ2B+\nwE7tp4A9AgMBAAECggEAGDBNExULeemZOZKLb9vaspUM8iv6ScPG4/EWpQdaJ1b6\nxhWj7/CvzgDpJsuZoasG7+3mN2L7ciJnkEnwJBJuNXLaNWREwC83sTmxNUwjQiCv\nidZ9RUU+DRHmdNvZn/AsqyPu31GiQEBkEBUz8aqHV8MP8uBh7fDAC0I6RvNl7rnd\nQKGlhlQvdkIgR/8Tu/2GchBEJyAEcRzUM6akDH56ZYjiJ58dvDkPqzixoL48sFUW\nLD6NUj4Yg2nNYPOpdUlEoHyhye494EV6osTN9sWM1dl3xMoljjlZXwa5cwK0XGG1\nS8wkvA1zup/OdFbhMWeFvQ6ptZhJzt/pLF7yBpqz2QKBgQC8NkhZmvoJfYjLqZxI\nbkzdH9j9Y2oAcCpNwUuAj5L0y59r4RteMm9n4iohepOFxm9ve8eAeF86wASx1W8K\nh05lZwaSHp1C7ifHYbIAyx502CZwJEmb+x0hbheOWP7wNXsManqPwh/t7IX1VlqY\npnHjHkDM5tiMzXB0t3EVB5TCyQKBgQDDj7YONa2wDIZ0vFO7ayhuTZG8Okc7iI1X\nwtOLlN+WjH/p9CAkjpVIPSJvzZmFy31wKXy7KvgXd+syslnGE474zUoPV13JCYqD\nEsULEly1HcP3ueH8McQr7ubw2VowQZQTitTM8GpZLgwALlOYkacyID52wXK6wkFd\ngdlU2qZ31QKBgFZJlLXwBT4gnXyx0AKs/vRJy5Ov5kM5KUvomJLiP/+W8wnkdEWl\nXcoEuEgNCGFGboRv/TD6/r3SltRpA7WR/vEiYLLQ8kxjHz2bnp4agWYUStsS5+Rp\nJhvTe105k/fQrQ8uKe6Kk1TjIS8vXe2qMYBsuEy4cQNslGgOSfvYM+1xAoGAZjXd\ndE3P6ryRJcAsjz8xkX0PV26qPk5feWgFguWbeqcST+fdSN4Q3gG91uuBevFhLVcx\nru2FX1a44fVfgGxGhHlfZOdHyEQpVdpgjCnbtHnLK9XGbDZijoe9wlIHlkMgLpqY\nIEjjfJrZgNq1rH8sAMHjg9QVr9DJsyKjwXrKGMUCgYEAjd6jYOCMPPyDiaYAjsN6\ns55u0UfJDleNrz1zZjzFiNSVderY5Wy+6l4LG/QuSx91D3TMO/cvtZiQNgR/eFAp\nynIEJ+KZYTPX9O54wLIaeAPPFFdtrmd9FSG975m+SMwgGafGFc9aZXwCiD6U0Nwl\nkTh/UrIAJ7XB5AGu9l0U2B8=\n-----END RSA PRIVATE KEY-----"
)

func main() {

	danaClient, err := dana.NewClient(dana.ClientConfig{
		MerchantID:   MerchantID,
		ClientID:     ClientID,
		ClientSecret: ClientSecret,
		PublicKey:    PublicKey,
		PrivateKey:   []byte(PrivateKey),
		UseSandbox:   true,
		Debug:        true,
	})

	if err != nil {
		log.Fatalf("Failed to initiate DANA client: %v", err)
	}

	ctx := context.Background()

	// UATScenario1(ctx, danaClient)
	// UATScenario2(ctx, danaClient)
	// UATScenario3a(ctx, danaClient)
	// UATScenario3b(ctx, danaClient)
	// UATScenario4(ctx, danaClient)
	// UATScenario5(ctx, danaClient)
	// UATScenario6a(ctx, danaClient)
	// UATScenario6b(ctx, danaClient)
	// UATScenario7(ctx, danaClient)
	// UATScenario8a(ctx, danaClient)
	// UATScenario8b(ctx, danaClient)
	// UATScenario9(ctx, danaClient)
	// UATScenario10(ctx, danaClient)
	// UATScenario11(ctx, danaClient)
	// UATScenario12(ctx, danaClient)
	// UATScenario13(ctx, danaClient)
	// UATScenario14(ctx, danaClient)
	// UATScenario15(ctx, danaClient)
	PaymentGateway(ctx, danaClient)
}

func PaymentGateway(ctx context.Context, danaClient *dana.Client) {
	partnerRefNo := uuid.NewString()
	externalUserID := uuid.NewString()
	amount := 2000

	req := dana.RequestPaymentGatewayDropInCreateOrder{
		PartnerReferenceNo: partnerRefNo,
		Amount: dana.Currency{
			Value:    fmt.Sprintf("%.2f", float64(amount)),
			Currency: "IDR",
		},
		// MerchantId is not needed here because the client sets it internally
		URLParams: []dana.PaymentGatewayNotifyURL{
			{
				URL:        "http://demo.localplace.id/api/v1/webhooks/dana",
				Type:       "NOTIFICATION",
				IsDeeplink: "N",
			},
			{
				URL:        "http://localhost",
				Type:       "PAY_RETURN",
				IsDeeplink: "N",
			},
		},
		AdditionalInfo: dana.PaymentGatewayAdditionalInfo{
			Order: &dana.PaymentGatewayOrderInfo{
				OrderTitle: "Payment Gateway Order",
				Buyer: dana.PaymentGatewayBuyerInfo{
					ExternalUserID: externalUserID,
				},
			},
			MCC: "5732",
			EnvInfo: dana.PaymentGatewayEnvInfo{
				SourcePlatform:    "IPG",
				TerminalType:      "WEB",
				OrderTerminalType: "WEB",
			},
		},
	}

	resp, err := danaClient.PaymentGatewayDropInCreateOrder(ctx, req)
	if err != nil {
		log.Fatalf("PaymentGatewayDropInCreateOrder failed: %v", err)
	}

	log.Printf("PaymentGatewayDropInCreateOrder Response: %+v\n", resp)
}

// func UATScenario1(ctx context.Context, danaClient *dana.Client) {
// 	// Scenario Name: Successfully doing customer top up

// 	uuid := uuid.NewString()
// 	phone_number := "62811742234"
// 	amount := 1

// 	// InquiryCustomerTopUpDisbursement
// 	req := dana.RequestAccountInquiryTopUpDisbursement{
// 		PartnerReferenceNo: uuid,
// 		CustomerNumber:     phone_number,
// 		Amount: dana.RequestCustomerTopUpDisbursementAmount{
// 			Value:    fmt.Sprintf("%.2f", float64(amount)),
// 			Currency: "IDR",
// 		},
// 		FeeAmount: dana.RequestCustomerTopUpDisbursementFeeAmount{
// 			Value:    "0.00",
// 			Currency: "IDR",
// 		},
// 		AdditionalInfo: dana.RequestCustomerTopUpDisbursementAdditionalInfo{
// 			FundType: "AGENT_TOPUP_FOR_USER_SETTLE",
// 		},
// 	}

// 	resp, err := danaClient.AccountInquiryTopUpDisbursement(ctx, req)
// 	if err != nil {
// 		log.Fatalln(resp)
// 	}

// 	log.Println("InquiryCustomerTopUpDisbursement Response:", resp, "\n")
// 	// End InquiryCustomerTopUpDisbursement

// 	// CustomerTopUpDisbursement
// 	reqTopUp := dana.RequestCustomerTopUpDisbursement{
// 		PartnerReferenceNo: uuid,
// 		CustomerNumber:     phone_number,
// 		Amount: dana.RequestCustomerTopUpDisbursementAmount{
// 			Value:    fmt.Sprintf("%.2f", float64(amount)),
// 			Currency: "IDR",
// 		},
// 		FeeAmount: dana.RequestCustomerTopUpDisbursementFeeAmount{
// 			Value:    "0.00",
// 			Currency: "IDR",
// 		},
// 		TransactionDate: dana.GenerateTimestamp(),
// 		AdditionalInfo: dana.RequestCustomerTopUpDisbursementAdditionalInfo{
// 			FundType: "AGENT_TOPUP_FOR_USER_SETTLE",
// 		},
// 	}

// 	respTopUp, err := danaClient.CustomerTopUpDisbursement(ctx, reqTopUp)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	log.Println("CustomerTopUpDisbursement Response:", respTopUp, "\n")
// 	// EndCustomerTopUpDisbursement
// }

// func UATScenario2(ctx context.Context, danaClient *dana.Client) {
// 	// Scenario Name: Failed to top up because merchant deposit is insufficient

// 	uuid := uuid.NewString()
// 	phone_number := "6281298055129"
// 	amount := 1

// 	// InquiryCustomerTopUpDisbursement
// 	req := dana.RequestAccountInquiryTopUpDisbursement{
// 		PartnerReferenceNo: uuid,
// 		CustomerNumber:     phone_number,
// 		Amount: dana.RequestCustomerTopUpDisbursementAmount{
// 			Value:    fmt.Sprintf("%.2f", float64(amount)),
// 			Currency: "IDR",
// 		},
// 		FeeAmount: dana.RequestCustomerTopUpDisbursementFeeAmount{
// 			Value:    "0.00",
// 			Currency: "IDR",
// 		},
// 		TransactionDate: dana.GenerateTimestamp(),
// 		AdditionalInfo: dana.RequestCustomerTopUpDisbursementAdditionalInfo{
// 			FundType: "AGENT_TOPUP_FOR_USER_SETTLE",
// 		},
// 	}

// 	resp, err := danaClient.AccountInquiryTopUpDisbursement(ctx, req)
// 	if err != nil {
// 		log.Fatalln(resp)
// 	}

// 	log.Println("InquiryCustomerTopUpDisbursement Response:", resp, "\n")
// 	// End InquiryCustomerTopUpDisbursement

// 	// CustomerTopUpDisbursement
// 	reqTopUp := dana.RequestCustomerTopUpDisbursement{
// 		PartnerReferenceNo: uuid,
// 		CustomerNumber:     phone_number,
// 		Amount: dana.RequestCustomerTopUpDisbursementAmount{
// 			Value:    fmt.Sprintf("%.2f", float64(amount)),
// 			Currency: "IDR",
// 		},
// 		FeeAmount: dana.RequestCustomerTopUpDisbursementFeeAmount{
// 			Value:    "0.00",
// 			Currency: "IDR",
// 		},
// 		TransactionDate: dana.GenerateTimestamp(),
// 		AdditionalInfo: dana.RequestCustomerTopUpDisbursementAdditionalInfo{
// 			FundType: "AGENT_TOPUP_FOR_USER_SETTLE",
// 		},
// 	}

// 	respTopUp, err := danaClient.CustomerTopUpDisbursement(ctx, reqTopUp)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	log.Println("CustomerTopUpDisbursement Response:", respTopUp, "\n")
// 	// EndCustomerTopUpDisbursement
// }

// func UATScenario3a(ctx context.Context, danaClient *dana.Client) {
// 	// Scenario Name: Gagal mendapatkan response Top Up (Timeout) yang kemudian status transaksinya didapatkan melalui H+1 settlement file setelah tetap tidak mendapatkan response pada retry mechanism dan inquiry status

// 	var (
// 		respTopUp *dana.ResponseCustomerTopUpDisbursement
// 		err       error
// 	)

// 	uuid := uuid.NewString()
// 	phone_number := "628551005454"
// 	amount := 1

// 	// InquiryCustomerTopUpDisbursement
// 	req := dana.RequestAccountInquiryTopUpDisbursement{
// 		PartnerReferenceNo: uuid,
// 		CustomerNumber:     phone_number,
// 		Amount: dana.RequestCustomerTopUpDisbursementAmount{
// 			Value:    fmt.Sprintf("%.2f", float64(amount)),
// 			Currency: "IDR",
// 		},
// 		FeeAmount: dana.RequestCustomerTopUpDisbursementFeeAmount{
// 			Value:    "0.00",
// 			Currency: "IDR",
// 		},
// 		TransactionDate: dana.GenerateTimestamp(),
// 		AdditionalInfo: dana.RequestCustomerTopUpDisbursementAdditionalInfo{
// 			FundType: "AGENT_TOPUP_FOR_USER_SETTLE",
// 		},
// 	}

// 	resp, err := danaClient.AccountInquiryTopUpDisbursement(ctx, req)
// 	if err != nil {
// 		log.Fatalln(resp)
// 	}

// 	log.Println("InquiryCustomerTopUpDisbursement Response:", resp, "\n")
// 	// End InquiryCustomerTopUpDisbursement

// 	// CustomerTopUpDisbursement
// 	// Retry intervals
// 	retryIntervals := []time.Duration{
// 		5 * time.Second,
// 		10 * time.Second,
// 		20 * time.Second,
// 		40 * time.Second,
// 		60 * time.Second,
// 	}

// 	reqTopUp := dana.RequestCustomerTopUpDisbursement{
// 		PartnerReferenceNo: uuid,
// 		CustomerNumber:     phone_number,
// 		Amount: dana.RequestCustomerTopUpDisbursementAmount{
// 			Value:    fmt.Sprintf("%.2f", float64(amount)),
// 			Currency: "IDR",
// 		},
// 		FeeAmount: dana.RequestCustomerTopUpDisbursementFeeAmount{
// 			Value:    "0.00",
// 			Currency: "IDR",
// 		},
// 		TransactionDate: dana.GenerateTimestamp(),
// 		AdditionalInfo: dana.RequestCustomerTopUpDisbursementAdditionalInfo{
// 			FundType: "AGENT_TOPUP_FOR_USER_SETTLE",
// 		},
// 	}

// 	for i, interval := range retryIntervals {
// 		respTopUp, err := danaClient.CustomerTopUpDisbursement(ctx, reqTopUp)
// 		if err == nil {
// 			log.Println("CustomerTopUpDisbursement Response:", respTopUp)
// 			break
// 		}

// 		log.Printf("Attempt %d failed: %v. Retrying in %v...\n", i+1, err, interval)
// 		select {
// 		case <-ctx.Done():
// 			log.Println("Context canceled or deadline exceeded, aborting retries.")
// 			return
// 		case <-time.After(interval):
// 			// continue to next retry
// 		}
// 	}

// 	if err != nil {
// 		log.Fatalf("Final attempt failed: %v\n", err)
// 	}

// 	log.Println("CustomerTopUpDisbursement Response:", respTopUp, "\n")
// 	// EndCustomerTopUpDisbursement

// 	// CustomerTopUpInquiryStatusDisbursement
// 	// reqTopUpStatus := dana.RequestCustomerTopUpInquiryStatusDisbursement{
// 	// 	OriginalPartnerReferenceNo: uuid,
// 	// 	ServiceCode: "38",
// 	// }

// 	// respTopUpStatus, err := danaClient.CustomerTopUpInquiryStatusDisbursement(ctx, reqTopUpStatus)
// 	// if err != nil {
// 	// 	log.Fatalln(err)
// 	// }
// 	// log.Println("CustomerTopUpInquiryStatusDisbursement Response:", respTopUpStatus, "\n")
// 	// EndCustomerTopUpInquiryStatusDisbursement
// }

// func UATScenario3b(ctx context.Context, danaClient *dana.Client) {
// 	// Scenario Name: Gagal mendapatkan response Top Up (Timeout) yang kemudian status transaksinya didapatkan melalui H+1 settlement file setelah tetap tidak mendapatkan response pada retry mechanism dan inquiry status

// 	var (
// 		respTopUp       *dana.ResponseCustomerTopUpDisbursement
// 		respTopUpStatus *dana.ResponseCustomerTopUpInquiryStatusDisbursement
// 		err             error
// 	)

// 	uuid := uuid.NewString()
// 	phone_number := "628551005454"
// 	amount := 1

// 	// InquiryCustomerTopUpDisbursement
// 	req := dana.RequestAccountInquiryTopUpDisbursement{
// 		PartnerReferenceNo: uuid,
// 		CustomerNumber:     phone_number,
// 		Amount: dana.RequestCustomerTopUpDisbursementAmount{
// 			Value:    fmt.Sprintf("%.2f", float64(amount)),
// 			Currency: "IDR",
// 		},
// 		FeeAmount: dana.RequestCustomerTopUpDisbursementFeeAmount{
// 			Value:    "0.00",
// 			Currency: "IDR",
// 		},
// 		TransactionDate: dana.GenerateTimestamp(),
// 		AdditionalInfo: dana.RequestCustomerTopUpDisbursementAdditionalInfo{
// 			FundType: "AGENT_TOPUP_FOR_USER_SETTLE",
// 		},
// 	}

// 	resp, err := danaClient.AccountInquiryTopUpDisbursement(ctx, req)
// 	if err != nil {
// 		log.Fatalln(resp)
// 	}

// 	log.Println("InquiryCustomerTopUpDisbursement Response:", resp, "\n")
// 	// End InquiryCustomerTopUpDisbursement

// 	// CustomerTopUpDisbursement
// 	// Retry intervals
// 	retryIntervals := []time.Duration{
// 		5 * time.Second,
// 		10 * time.Second,
// 		20 * time.Second,
// 		40 * time.Second,
// 		60 * time.Second,
// 	}

// 	reqTopUp := dana.RequestCustomerTopUpDisbursement{
// 		PartnerReferenceNo: uuid,
// 		CustomerNumber:     phone_number,
// 		Amount: dana.RequestCustomerTopUpDisbursementAmount{
// 			Value:    fmt.Sprintf("%.2f", float64(amount)),
// 			Currency: "IDR",
// 		},
// 		FeeAmount: dana.RequestCustomerTopUpDisbursementFeeAmount{
// 			Value:    "0.00",
// 			Currency: "IDR",
// 		},
// 		TransactionDate: dana.GenerateTimestamp(),
// 		AdditionalInfo: dana.RequestCustomerTopUpDisbursementAdditionalInfo{
// 			FundType: "AGENT_TOPUP_FOR_USER_SETTLE",
// 		},
// 	}

// 	for i, interval := range retryIntervals {
// 		respTopUp, err := danaClient.CustomerTopUpDisbursement(ctx, reqTopUp)
// 		if err == nil {
// 			log.Println("CustomerTopUpDisbursement Response:", respTopUp)
// 			break
// 		}

// 		log.Printf("Attempt %d failed: %v. Retrying in %v...\n", i+1, err, interval)
// 		select {
// 		case <-ctx.Done():
// 			log.Println("Context canceled or deadline exceeded, aborting retries.")
// 			return
// 		case <-time.After(interval):
// 			// continue to next retry
// 		}
// 	}

// 	log.Println("CustomerTopUpDisbursement Response:", respTopUp, "\n")
// 	// EndCustomerTopUpDisbursement

// 	// CustomerTopUpInquiryStatusDisbursement
// 	reqTopUpStatus := dana.RequestCustomerTopUpInquiryStatusDisbursement{
// 		OriginalPartnerReferenceNo: uuid,
// 		ServiceCode:                "XY",
// 	}

// 	for i, interval := range retryIntervals {
// 		respTopUpStatus, err := danaClient.CustomerTopUpInquiryStatusDisbursement(ctx, reqTopUpStatus)
// 		if err == nil {
// 			log.Println("CustomerTopUpInquiryStatusDisbursement Response:", respTopUpStatus)
// 			break
// 		}

// 		log.Printf("Attempt %d failed: %v. Retrying in %v...\n", i+1, err, interval)
// 		select {
// 		case <-ctx.Done():
// 			log.Println("Context canceled or deadline exceeded, aborting retries.")
// 			return
// 		case <-time.After(interval):
// 			// continue to next retry
// 		}
// 	}

// 	if err != nil {
// 		log.Fatalf("Final attempt failed: %v\n", err)
// 	}

// 	log.Println("CustomerTopUpInquiryStatusDisbursement Response:", respTopUpStatus, "\n")
// 	// EndCustomerTopUpInquiryStatusDisbursement
// }

// func UATScenario4(ctx context.Context, danaClient *dana.Client) {
// 	// Scenario Name: Failed to get a Top Up response (Timeout) and then the transaction status is obtained through a retry process (idempotent)
// 	var (
// 		respTopUp *dana.ResponseCustomerTopUpDisbursement
// 		err       error
// 	)

// 	uuid := "72a3375f-a56d-447d-8e31-78f6bde21831"
// 	phone_number := "628996647676"
// 	amount := 1

// 	// InquiryCustomerTopUpDisbursement
// 	// req := dana.RequestAccountInquiryTopUpDisbursement{
// 	// 	PartnerReferenceNo: uuid,
// 	// 	CustomerNumber:     phone_number,
// 	// 	Amount: dana.RequestCustomerTopUpDisbursementAmount{
// 	// 		Value:    fmt.Sprintf("%.2f", float64(amount)),
// 	// 		currency: "IDR",
// 	// 	},
// 	// 	FeeAmount: dana.RequestCustomerTopUpDisbursementFeeAmount{
// 	// 		Value:    "0.00",
// 	// 		currency: "IDR",
// 	// 	},
// 	// 	TransactionDate: dana.GenerateTimestamp(),
// 	// 	AdditionalInfo: dana.RequestCustomerTopUpDisbursementAdditionalInfo{
// 	// 		FundType: "AGENT_TOPUP_FOR_USER_SETTLE",
// 	// 	},
// 	// }

// 	// resp, err := danaClient.AccountInquiryTopUpDisbursement(ctx, req)
// 	// if err != nil {
// 	// 	log.Fatalln(resp)
// 	// }

// 	// log.Println("InquiryCustomerTopUpDisbursement Response:", resp, "\n")
// 	// End InquiryCustomerTopUpDisbursement

// 	// CustomerTopUpDisbursement
// 	// Retry intervals
// 	retryIntervals := []time.Duration{
// 		5 * time.Second,
// 		10 * time.Second,
// 		20 * time.Second,
// 		40 * time.Second,
// 		60 * time.Second,
// 	}

// 	reqTopUp := dana.RequestCustomerTopUpDisbursement{
// 		PartnerReferenceNo: uuid,
// 		CustomerNumber:     phone_number,
// 		Amount:             dana.NewIDRCurrency(fmt.Sprintf("%.2f", float64(amount))),
// 		FeeAmount:          dana.NewIDRCurrency("0.00"),
// 		AdditionalInfo: dana.RequestCustomerTopUpDisbursementAdditionalInfo{
// 			FundType: "AGENT_TOPUP_FOR_USER_SETTLE",
// 		},
// 	}

// 	for i, interval := range retryIntervals {
// 		respTopUp, err := danaClient.CustomerTopUpDisbursement(ctx, reqTopUp)
// 		if err == nil {
// 			log.Println("CustomerTopUpDisbursement Response:", respTopUp)
// 			break
// 		}

// 		log.Printf("Attempt %d failed: %v. Retrying in %v...\n", i+1, err, interval)
// 		select {
// 		case <-ctx.Done():
// 			log.Println("Context canceled or deadline exceeded, aborting retries.")
// 			return
// 		case <-time.After(interval):
// 			// continue to next retry
// 		}
// 	}

// 	if err != nil {
// 		log.Fatalf("Final attempt failed: %v\n", err)
// 	}

// 	log.Println("CustomerTopUpDisbursement Response:", respTopUp, "\n")
// 	// EndCustomerTopUpDisbursement
// }

// func UATScenario5(ctx context.Context, danaClient *dana.Client) {
// 	// Scenario Name: Failed to get a Top Up response (Timeout) and then the transaction status was obtained via status inquiry (Successful transaction)

// 	var (
// 		respTopUp *dana.ResponseCustomerTopUpDisbursement
// 		err       error
// 	)

// 	uuid := "7f554eed-69fe-451d-9d79-5c40539e425d"
// 	phone_number := "6281322245545"
// 	amount := 1

// 	// InquiryCustomerTopUpDisbursement
// 	req := dana.RequestAccountInquiryTopUpDisbursement{
// 		PartnerReferenceNo: uuid,
// 		CustomerNumber:     phone_number,
// 		Amount: dana.RequestCustomerTopUpDisbursementAmount{
// 			Value:    fmt.Sprintf("%.2f", float64(amount)),
// 			Currency: "IDR",
// 		},
// 		FeeAmount: dana.RequestCustomerTopUpDisbursementFeeAmount{
// 			Value:    "0.00",
// 			Currency: "IDR",
// 		},
// 		TransactionDate: dana.GenerateTimestamp(),
// 		AdditionalInfo: dana.RequestCustomerTopUpDisbursementAdditionalInfo{
// 			FundType: "AGENT_TOPUP_FOR_USER_SETTLE",
// 		},
// 	}

// 	resp, err := danaClient.AccountInquiryTopUpDisbursement(ctx, req)
// 	if err != nil {
// 		log.Fatalln(resp)
// 	}

// 	log.Println("InquiryCustomerTopUpDisbursement Response:", resp, "\n")
// 	// End InquiryCustomerTopUpDisbursement

// 	// CustomerTopUpDisbursement
// 	// Retry intervals
// 	retryIntervals := []time.Duration{
// 		5 * time.Second,
// 		10 * time.Second,
// 		20 * time.Second,
// 		40 * time.Second,
// 		60 * time.Second,
// 	}

// 	reqTopUp := dana.RequestCustomerTopUpDisbursement{
// 		PartnerReferenceNo: uuid,
// 		CustomerNumber:     phone_number,
// 		Amount: dana.RequestCustomerTopUpDisbursementAmount{
// 			Value:    fmt.Sprintf("%.2f", float64(amount)),
// 			Currency: "IDR",
// 		},
// 		FeeAmount: dana.RequestCustomerTopUpDisbursementFeeAmount{
// 			Value:    "0.00",
// 			Currency: "IDR",
// 		},
// 		TransactionDate: dana.GenerateTimestamp(),
// 		AdditionalInfo: dana.RequestCustomerTopUpDisbursementAdditionalInfo{
// 			FundType: "AGENT_TOPUP_FOR_USER_SETTLE",
// 		},
// 	}

// 	for i, interval := range retryIntervals {
// 		respTopUp, err := danaClient.CustomerTopUpDisbursement(ctx, reqTopUp)
// 		if err == nil {
// 			log.Println("CustomerTopUpDisbursement Response:", respTopUp)
// 			break
// 		}

// 		log.Printf("Attempt %d failed: %v. Retrying in %v...\n", i+1, err, interval)
// 		select {
// 		case <-ctx.Done():
// 			log.Println("Context canceled or deadline exceeded, aborting retries.")
// 			return
// 		case <-time.After(interval):
// 			// continue to next retry
// 		}
// 	}

// 	if err != nil {
// 		log.Fatalf("Final attempt failed: %v\n", err)
// 	}

// 	log.Println("CustomerTopUpDisbursement Response:", respTopUp, "\n")
// 	// EndCustomerTopUpDisbursement

// 	// CustomerTopUpInquiryStatusDisbursement
// 	reqTopUpStatus := dana.RequestCustomerTopUpInquiryStatusDisbursement{
// 		OriginalPartnerReferenceNo: uuid,
// 		ServiceCode:                "38",
// 	}

// 	respTopUpStatus, err := danaClient.CustomerTopUpInquiryStatusDisbursement(ctx, reqTopUpStatus)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	log.Println("CustomerTopUpInquiryStatusDisbursement Response:", respTopUpStatus, "\n")
// 	// EndCustomerTopUpInquiryStatusDisbursement
// }

// func UATScenario6a(ctx context.Context, danaClient *dana.Client) {
// 	// Scenario Name: Failed to get a Top Up response (Timeout) and then the transaction status was obtained via status inquiry (Transaction failed)

// 	var (
// 		respTopUp *dana.ResponseCustomerTopUpDisbursement
// 		err       error
// 	)

// 	uuid := uuid.NewString()
// 	phone_number := "6281298055138"
// 	amount := 1

// 	// InquiryCustomerTopUpDisbursement
// 	req := dana.RequestAccountInquiryTopUpDisbursement{
// 		PartnerReferenceNo: uuid,
// 		CustomerNumber:     phone_number,
// 		Amount: dana.RequestCustomerTopUpDisbursementAmount{
// 			Value:    fmt.Sprintf("%.2f", float64(amount)),
// 			Currency: "IDR",
// 		},
// 		FeeAmount: dana.RequestCustomerTopUpDisbursementFeeAmount{
// 			Value:    "0.00",
// 			Currency: "IDR",
// 		},
// 		TransactionDate: dana.GenerateTimestamp(),
// 		AdditionalInfo: dana.RequestCustomerTopUpDisbursementAdditionalInfo{
// 			FundType: "AGENT_TOPUP_FOR_USER_SETTLE",
// 		},
// 	}

// 	resp, err := danaClient.AccountInquiryTopUpDisbursement(ctx, req)
// 	if err != nil {
// 		log.Fatalln(resp)
// 	}

// 	log.Println("InquiryCustomerTopUpDisbursement Response:", resp, "\n")
// 	// End InquiryCustomerTopUpDisbursement

// 	// CustomerTopUpDisbursement
// 	// Retry intervals
// 	retryIntervals := []time.Duration{
// 		5 * time.Second,
// 		10 * time.Second,
// 		20 * time.Second,
// 		40 * time.Second,
// 		60 * time.Second,
// 	}

// 	reqTopUp := dana.RequestCustomerTopUpDisbursement{
// 		PartnerReferenceNo: uuid,
// 		CustomerNumber:     phone_number,
// 		Amount: dana.RequestCustomerTopUpDisbursementAmount{
// 			Value:    fmt.Sprintf("%.2f", float64(amount)),
// 			Currency: "IDR",
// 		},
// 		FeeAmount: dana.RequestCustomerTopUpDisbursementFeeAmount{
// 			Value:    "0.00",
// 			Currency: "IDR",
// 		},
// 		TransactionDate: dana.GenerateTimestamp(),
// 		AdditionalInfo: dana.RequestCustomerTopUpDisbursementAdditionalInfo{
// 			FundType: "AGENT_TOPUP_FOR_USER_SETTLE",
// 		},
// 	}

// 	for i, interval := range retryIntervals {
// 		respTopUp, err := danaClient.CustomerTopUpDisbursement(ctx, reqTopUp)
// 		if err == nil {
// 			log.Println("CustomerTopUpDisbursement Response:", respTopUp)
// 			break
// 		}

// 		log.Printf("Attempt %d failed: %v. Retrying in %v...\n", i+1, err, interval)
// 		select {
// 		case <-ctx.Done():
// 			log.Println("Context canceled or deadline exceeded, aborting retries.")
// 			return
// 		case <-time.After(interval):
// 			// continue to next retry
// 		}
// 	}

// 	if err != nil {
// 		log.Fatalf("Final attempt failed: %v\n", err)
// 	}

// 	log.Println("CustomerTopUpDisbursement Response:", respTopUp, "\n")
// 	// EndCustomerTopUpDisbursement

// 	// CustomerTopUpInquiryStatusDisbursement
// 	reqTopUpStatus := dana.RequestCustomerTopUpInquiryStatusDisbursement{
// 		OriginalPartnerReferenceNo: uuid,
// 		ServiceCode:                "38",
// 	}

// 	respTopUpStatus, err := danaClient.CustomerTopUpInquiryStatusDisbursement(ctx, reqTopUpStatus)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	log.Println("CustomerTopUpInquiryStatusDisbursement Response:", respTopUpStatus, "\n")
// 	// EndCustomerTopUpInquiryStatusDisbursement
// }

// func UATScenario6b(ctx context.Context, danaClient *dana.Client) {
// 	// Scenario Name: Failed to get a Top Up response (Timeout) and then the transaction status was obtained via status inquiry (Transaction failed)

// 	var (
// 		respTopUp *dana.ResponseCustomerTopUpDisbursement
// 		err       error
// 	)

// 	uuid := uuid.NewString()
// 	phone_number := "628521470963"
// 	amount := 1

// 	// InquiryCustomerTopUpDisbursement
// 	req := dana.RequestAccountInquiryTopUpDisbursement{
// 		PartnerReferenceNo: uuid,
// 		CustomerNumber:     phone_number,
// 		Amount: dana.RequestCustomerTopUpDisbursementAmount{
// 			Value:    fmt.Sprintf("%.2f", float64(amount)),
// 			Currency: "IDR",
// 		},
// 		FeeAmount: dana.RequestCustomerTopUpDisbursementFeeAmount{
// 			Value:    "0.00",
// 			Currency: "IDR",
// 		},
// 		TransactionDate: dana.GenerateTimestamp(),
// 		AdditionalInfo: dana.RequestCustomerTopUpDisbursementAdditionalInfo{
// 			FundType: "AGENT_TOPUP_FOR_USER_SETTLE",
// 		},
// 	}

// 	resp, err := danaClient.AccountInquiryTopUpDisbursement(ctx, req)
// 	if err != nil {
// 		log.Fatalln(resp)
// 	}

// 	log.Println("InquiryCustomerTopUpDisbursement Response:", resp, "\n")
// 	// End InquiryCustomerTopUpDisbursement

// 	// CustomerTopUpDisbursement
// 	// Retry intervals
// 	retryIntervals := []time.Duration{
// 		5 * time.Second,
// 		10 * time.Second,
// 		20 * time.Second,
// 		40 * time.Second,
// 		60 * time.Second,
// 	}

// 	reqTopUp := dana.RequestCustomerTopUpDisbursement{
// 		PartnerReferenceNo: uuid,
// 		CustomerNumber:     phone_number,
// 		Amount: dana.RequestCustomerTopUpDisbursementAmount{
// 			Value:    fmt.Sprintf("%.2f", float64(amount)),
// 			Currency: "IDR",
// 		},
// 		FeeAmount: dana.RequestCustomerTopUpDisbursementFeeAmount{
// 			Value:    "0.00",
// 			Currency: "IDR",
// 		},
// 		TransactionDate: dana.GenerateTimestamp(),
// 		AdditionalInfo: dana.RequestCustomerTopUpDisbursementAdditionalInfo{
// 			FundType: "AGENT_TOPUP_FOR_USER_SETTLE",
// 		},
// 	}

// 	for i, interval := range retryIntervals {
// 		respTopUp, err := danaClient.CustomerTopUpDisbursement(ctx, reqTopUp)
// 		if err == nil {
// 			log.Println("CustomerTopUpDisbursement Response:", respTopUp)
// 			break
// 		}

// 		log.Printf("Attempt %d failed: %v. Retrying in %v...\n", i+1, err, interval)
// 		select {
// 		case <-ctx.Done():
// 			log.Println("Context canceled or deadline exceeded, aborting retries.")
// 			return
// 		case <-time.After(interval):
// 			// continue to next retry
// 		}
// 	}

// 	if err != nil {
// 		log.Fatalf("Final attempt failed: %v\n", err)
// 	}

// 	log.Println("CustomerTopUpDisbursement Response:", respTopUp, "\n")
// 	// EndCustomerTopUpDisbursement

// 	// CustomerTopUpInquiryStatusDisbursement
// 	reqTopUpStatus := dana.RequestCustomerTopUpInquiryStatusDisbursement{
// 		OriginalPartnerReferenceNo: uuid,
// 		ServiceCode:                "38",
// 	}

// 	respTopUpStatus, err := danaClient.CustomerTopUpInquiryStatusDisbursement(ctx, reqTopUpStatus)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	log.Println("CustomerTopUpInquiryStatusDisbursement Response:", respTopUpStatus, "\n")
// 	// EndCustomerTopUpInquiryStatusDisbursement
// }

// func UATScenario7(ctx context.Context, danaClient *dana.Client) {
// 	// Scenario Name: Failed to make account inquiry on blocked account (frozen account)
// 	uuid := uuid.NewString()
// 	phone_number := "628123456667"
// 	amount := 1

// 	// AccountInquiryTopUpDisbursement
// 	req := dana.RequestAccountInquiryTopUpDisbursement{
// 		PartnerReferenceNo: uuid,
// 		CustomerNumber:     phone_number,
// 		Amount: dana.RequestCustomerTopUpDisbursementAmount{
// 			Value:    fmt.Sprintf("%.2f", float64(amount)),
// 			Currency: "IDR",
// 		},
// 		FeeAmount: dana.RequestCustomerTopUpDisbursementFeeAmount{
// 			Value:    "0.00",
// 			Currency: "IDR",
// 		},
// 		TransactionDate: dana.GenerateTimestamp(),
// 		AdditionalInfo: dana.RequestCustomerTopUpDisbursementAdditionalInfo{
// 			FundType: "AGENT_TOPUP_FOR_USER_SETTLE",
// 		},
// 	}

// 	resp, err := danaClient.AccountInquiryTopUpDisbursement(ctx, req)
// 	if err != nil {
// 		log.Fatalln(resp)
// 	}

// 	log.Println("AccountInquiryTopUpDisbursement Response:", resp, "\n")
// 	// End AccountInquiryTopUpDisbursement
// }

// func UATScenario8a(ctx context.Context, danaClient *dana.Client) {
// 	// Scenario Name: Failed to Top Up on a blocked account (frozen account)

// 	var (
// 		respTopUp *dana.ResponseCustomerTopUpDisbursement
// 		err       error
// 	)

// 	uuid := uuid.NewString()
// 	phone_number := "628996647679"
// 	amount := 1

// 	// InquiryCustomerTopUpDisbursement
// 	req := dana.RequestAccountInquiryTopUpDisbursement{
// 		PartnerReferenceNo: uuid,
// 		CustomerNumber:     phone_number,
// 		Amount: dana.RequestCustomerTopUpDisbursementAmount{
// 			Value:    fmt.Sprintf("%.2f", float64(amount)),
// 			Currency: "IDR",
// 		},
// 		FeeAmount: dana.RequestCustomerTopUpDisbursementFeeAmount{
// 			Value:    "0.00",
// 			Currency: "IDR",
// 		},
// 		TransactionDate: dana.GenerateTimestamp(),
// 		AdditionalInfo: dana.RequestCustomerTopUpDisbursementAdditionalInfo{
// 			FundType: "AGENT_TOPUP_FOR_USER_SETTLE",
// 		},
// 	}

// 	resp, err := danaClient.AccountInquiryTopUpDisbursement(ctx, req)
// 	if err != nil {
// 		log.Fatalln(resp)
// 	}

// 	log.Println("InquiryCustomerTopUpDisbursement Response:", resp, "\n")
// 	// End InquiryCustomerTopUpDisbursement

// 	// CustomerTopUpDisbursement
// 	// Retry intervals
// 	retryIntervals := []time.Duration{
// 		5 * time.Second,
// 		10 * time.Second,
// 		20 * time.Second,
// 		40 * time.Second,
// 		60 * time.Second,
// 	}

// 	reqTopUp := dana.RequestCustomerTopUpDisbursement{
// 		PartnerReferenceNo: uuid,
// 		CustomerNumber:     phone_number,
// 		Amount: dana.RequestCustomerTopUpDisbursementAmount{
// 			Value:    fmt.Sprintf("%.2f", float64(amount)),
// 			Currency: "IDR",
// 		},
// 		FeeAmount: dana.RequestCustomerTopUpDisbursementFeeAmount{
// 			Value:    "0.00",
// 			Currency: "IDR",
// 		},
// 		TransactionDate: dana.GenerateTimestamp(),
// 		AdditionalInfo: dana.RequestCustomerTopUpDisbursementAdditionalInfo{
// 			FundType: "AGENT_TOPUP_FOR_USER_SETTLE",
// 		},
// 	}

// 	for i, interval := range retryIntervals {
// 		respTopUp, err := danaClient.CustomerTopUpDisbursement(ctx, reqTopUp)
// 		if err == nil {
// 			log.Println("CustomerTopUpDisbursement Response:", respTopUp)
// 			break
// 		}

// 		log.Printf("Attempt %d failed: %v. Retrying in %v...\n", i+1, err, interval)
// 		select {
// 		case <-ctx.Done():
// 			log.Println("Context canceled or deadline exceeded, aborting retries.")
// 			return
// 		case <-time.After(interval):
// 			// continue to next retry
// 		}
// 	}

// 	if err != nil {
// 		log.Fatalf("Final attempt failed: %v\n", err)
// 	}

// 	log.Println("CustomerTopUpDisbursement Response:", respTopUp, "\n")
// 	// EndCustomerTopUpDisbursement
// }

// func UATScenario8b(ctx context.Context, danaClient *dana.Client) {
// 	// Scenario Name: Failed to make account inquiry on unregistered account
// 	uuid := uuid.NewString()
// 	phone_number := "628152768647"
// 	amount := 1

// 	// AccountInquiryTopUpDisbursement
// 	req := dana.RequestAccountInquiryTopUpDisbursement{
// 		PartnerReferenceNo: uuid,
// 		CustomerNumber:     phone_number,
// 		Amount: dana.RequestCustomerTopUpDisbursementAmount{
// 			Value:    fmt.Sprintf("%.2f", float64(amount)),
// 			Currency: "IDR",
// 		},
// 		FeeAmount: dana.RequestCustomerTopUpDisbursementFeeAmount{
// 			Value:    "0.00",
// 			Currency: "IDR",
// 		},
// 		TransactionDate: dana.GenerateTimestamp(),
// 		AdditionalInfo: dana.RequestCustomerTopUpDisbursementAdditionalInfo{
// 			FundType: "AGENT_TOPUP_FOR_USER_SETTLE",
// 		},
// 	}

// 	resp, err := danaClient.AccountInquiryTopUpDisbursement(ctx, req)
// 	if err != nil {
// 		log.Fatalln(resp)
// 	}

// 	log.Println("AccountInquiryTopUpDisbursement Response:", resp, "\n")
// 	// End AccountInquiryTopUpDisbursement
// }

// func UATScenario9(ctx context.Context, danaClient *dana.Client) {
// 	// Scenario Name: Failed to make account inquiry on unregistered account
// 	uuid := uuid.NewString()
// 	phone_number := "62811742234"
// 	amount := 21000000

// 	// AccountInquiryTopUpDisbursement
// 	req := dana.RequestAccountInquiryTopUpDisbursement{
// 		PartnerReferenceNo: uuid,
// 		CustomerNumber:     phone_number,
// 		Amount: dana.RequestCustomerTopUpDisbursementAmount{
// 			Value:    fmt.Sprintf("%.2f", float64(amount)),
// 			Currency: "IDR",
// 		},
// 		FeeAmount: dana.RequestCustomerTopUpDisbursementFeeAmount{
// 			Value:    "0.00",
// 			Currency: "IDR",
// 		},
// 		TransactionDate: dana.GenerateTimestamp(),
// 		AdditionalInfo: dana.RequestCustomerTopUpDisbursementAdditionalInfo{
// 			FundType: "AGENT_TOPUP_FOR_USER_SETTLE",
// 		},
// 	}

// 	resp, err := danaClient.AccountInquiryTopUpDisbursement(ctx, req)
// 	if err != nil {
// 		log.Fatalln(resp)
// 	}

// 	log.Println("AccountInquiryTopUpDisbursement Response:", resp, "\n")
// 	// End AccountInquiryTopUpDisbursement
// }

// func UATScenario10(ctx context.Context, danaClient *dana.Client) {
// 	// Scenario Name: Failed to Top Up on a blocked account (frozen account)

// 	var (
// 		respTopUp *dana.ResponseCustomerTopUpDisbursement
// 		err       error
// 	)

// 	uuid := uuid.NewString()
// 	phone_number := "6287825574103"
// 	amount := 1

// 	// InquiryCustomerTopUpDisbursement
// 	req := dana.RequestAccountInquiryTopUpDisbursement{
// 		PartnerReferenceNo: uuid,
// 		CustomerNumber:     phone_number,
// 		Amount: dana.RequestCustomerTopUpDisbursementAmount{
// 			Value:    fmt.Sprintf("%.2f", float64(amount)),
// 			Currency: "IDR",
// 		},
// 		FeeAmount: dana.RequestCustomerTopUpDisbursementFeeAmount{
// 			Value:    "0.00",
// 			Currency: "IDR",
// 		},
// 		TransactionDate: dana.GenerateTimestamp(),
// 		AdditionalInfo: dana.RequestCustomerTopUpDisbursementAdditionalInfo{
// 			FundType: "AGENT_TOPUP_FOR_USER_SETTLE",
// 		},
// 	}

// 	resp, err := danaClient.AccountInquiryTopUpDisbursement(ctx, req)
// 	if err != nil {
// 		log.Fatalln(resp)
// 	}

// 	log.Println("InquiryCustomerTopUpDisbursement Response:", resp, "\n")
// 	// End InquiryCustomerTopUpDisbursement

// 	// CustomerTopUpDisbursement
// 	// Retry intervals
// 	retryIntervals := []time.Duration{
// 		5 * time.Second,
// 		10 * time.Second,
// 		20 * time.Second,
// 		40 * time.Second,
// 		60 * time.Second,
// 	}

// 	reqTopUp := dana.RequestCustomerTopUpDisbursement{
// 		PartnerReferenceNo: uuid,
// 		CustomerNumber:     phone_number,
// 		Amount: dana.RequestCustomerTopUpDisbursementAmount{
// 			Value:    fmt.Sprintf("%.2f", float64(amount)),
// 			Currency: "IDR",
// 		},
// 		FeeAmount: dana.RequestCustomerTopUpDisbursementFeeAmount{
// 			Value:    "0.00",
// 			Currency: "IDR",
// 		},
// 		TransactionDate: dana.GenerateTimestamp(),
// 		AdditionalInfo: dana.RequestCustomerTopUpDisbursementAdditionalInfo{
// 			FundType: "AGENT_TOPUP_FOR_USER_SETTLE",
// 		},
// 	}

// 	for i, interval := range retryIntervals {
// 		respTopUp, err := danaClient.CustomerTopUpDisbursement(ctx, reqTopUp)
// 		if err == nil {
// 			log.Println("CustomerTopUpDisbursement Response:", respTopUp)
// 			break
// 		}

// 		log.Printf("Attempt %d failed: %v. Retrying in %v...\n", i+1, err, interval)
// 		select {
// 		case <-ctx.Done():
// 			log.Println("Context canceled or deadline exceeded, aborting retries.")
// 			return
// 		case <-time.After(interval):
// 			// continue to next retry
// 		}
// 	}

// 	if err != nil {
// 		log.Fatalf("Final attempt failed: %v\n", err)
// 	}

// 	log.Println("CustomerTopUpDisbursement Response:", respTopUp, "\n")
// 	// EndCustomerTopUpDisbursement
// }

// func UATScenario11a(ctx context.Context, danaClient *dana.Client) {
// 	// Scenario Name: Any Services

// 	uuid := uuid.NewString()
// 	phone_number := "6287720766990"
// 	amount := 1

// 	// InquiryCustomerTopUpDisbursement
// 	req := dana.RequestAccountInquiryTopUpDisbursement{
// 		PartnerReferenceNo: uuid,
// 		CustomerNumber:     phone_number,
// 		Amount: dana.RequestCustomerTopUpDisbursementAmount{
// 			Value:    fmt.Sprintf("%.2f", float64(amount)),
// 			Currency: "IDR",
// 		},
// 		FeeAmount: dana.RequestCustomerTopUpDisbursementFeeAmount{
// 			Value:    "0.00",
// 			Currency: "IDR",
// 		},
// 		TransactionDate: dana.GenerateTimestamp(),
// 		AdditionalInfo: dana.RequestCustomerTopUpDisbursementAdditionalInfo{
// 			FundType: "AGENT_TOPUP_FOR_USER_SETTLE",
// 		},
// 	}

// 	resp, err := danaClient.AccountInquiryTopUpDisbursement(ctx, req)
// 	if err != nil {
// 		log.Fatalln(resp)
// 	}

// 	log.Println("InquiryCustomerTopUpDisbursement Response:", resp, "\n")
// 	// End InquiryCustomerTopUpDisbursement
// }

// func UATScenario12(ctx context.Context, danaClient *dana.Client) {
// 	// Scenario Name: Successfully Top Up to make inconsistent requests

// 	uuid := uuid.NewString()
// 	phone_number := "6287720766990"
// 	amount := 1

// 	// InquiryCustomerTopUpDisbursement
// 	req := dana.RequestAccountInquiryTopUpDisbursement{
// 		PartnerReferenceNo: uuid,
// 		CustomerNumber:     phone_number,
// 		Amount: dana.RequestCustomerTopUpDisbursementAmount{
// 			Value:    fmt.Sprintf("%.2f", float64(amount)),
// 			Currency: "IDR",
// 		},
// 		FeeAmount: dana.RequestCustomerTopUpDisbursementFeeAmount{
// 			Value:    "0.00",
// 			Currency: "IDR",
// 		},
// 		TransactionDate: dana.GenerateTimestamp(),
// 		AdditionalInfo: dana.RequestCustomerTopUpDisbursementAdditionalInfo{
// 			FundType: "AGENT_TOPUP_FOR_USER_SETTLE",
// 		},
// 	}

// 	resp, err := danaClient.AccountInquiryTopUpDisbursement(ctx, req)
// 	if err != nil {
// 		log.Fatalln(resp)
// 	}

// 	log.Println("InquiryCustomerTopUpDisbursement Response:", resp, "\n")
// 	// End InquiryCustomerTopUpDisbursement

// 	// CustomerTopUpDisbursement
// 	reqTopUp := dana.RequestCustomerTopUpDisbursement{
// 		PartnerReferenceNo: uuid,
// 		CustomerNumber:     phone_number,
// 		Amount: dana.RequestCustomerTopUpDisbursementAmount{
// 			Value:    fmt.Sprintf("%.2f", float64(amount)),
// 			Currency: "IDR",
// 		},
// 		FeeAmount: dana.RequestCustomerTopUpDisbursementFeeAmount{
// 			Value:    "0.00",
// 			Currency: "IDR",
// 		},
// 		TransactionDate: dana.GenerateTimestamp(),
// 		AdditionalInfo: dana.RequestCustomerTopUpDisbursementAdditionalInfo{
// 			FundType: "AGENT_TOPUP_FOR_USER_SETTLE",
// 		},
// 	}

// 	respTopUp, err := danaClient.CustomerTopUpDisbursement(ctx, reqTopUp)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	log.Println("CustomerTopUpDisbursement Response:", respTopUp, "\n")
// 	// EndCustomerTopUpDisbursement

// 	// CustomerTopUpDisbursement
// 	reqTopUp2 := dana.RequestCustomerTopUpDisbursement{
// 		PartnerReferenceNo: uuid,
// 		CustomerNumber:     phone_number,
// 		Amount: dana.RequestCustomerTopUpDisbursementAmount{
// 			Value:    fmt.Sprintf("%.2f", float64(amount)+1),
// 			Currency: "IDR",
// 		},
// 		FeeAmount: dana.RequestCustomerTopUpDisbursementFeeAmount{
// 			Value:    "0.00",
// 			Currency: "IDR",
// 		},
// 		TransactionDate: dana.GenerateTimestamp(),
// 		AdditionalInfo: dana.RequestCustomerTopUpDisbursementAdditionalInfo{
// 			FundType: "AGENT_TOPUP_FOR_USER_SETTLE",
// 		},
// 	}

// 	respTopUp2, err := danaClient.CustomerTopUpDisbursement(ctx, reqTopUp2)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	log.Println("CustomerTopUpDisbursement Response:", respTopUp2, "\n")
// 	// EndCustomerTopUpDisbursement
// }

// func UATScenario13(ctx context.Context, danaClient *dana.Client) {
// 	// Scenario Name: Failed to get top up status due to internal server error
// 	var (
// 		respTopUp *dana.ResponseCustomerTopUpDisbursement
// 		err       error
// 	)

// 	uuid := uuid.NewString()
// 	phone_number := "628551008794"
// 	amount := 1

// 	// InquiryCustomerTopUpDisbursement
// 	req := dana.RequestAccountInquiryTopUpDisbursement{
// 		PartnerReferenceNo: uuid,
// 		CustomerNumber:     phone_number,
// 		Amount: dana.RequestCustomerTopUpDisbursementAmount{
// 			Value:    fmt.Sprintf("%.2f", float64(amount)),
// 			Currency: "IDR",
// 		},
// 		FeeAmount: dana.RequestCustomerTopUpDisbursementFeeAmount{
// 			Value:    "0.00",
// 			Currency: "IDR",
// 		},
// 		TransactionDate: dana.GenerateTimestamp(),
// 		AdditionalInfo: dana.RequestCustomerTopUpDisbursementAdditionalInfo{
// 			FundType: "AGENT_TOPUP_FOR_USER_SETTLE",
// 		},
// 	}

// 	resp, err := danaClient.AccountInquiryTopUpDisbursement(ctx, req)
// 	if err != nil {
// 		log.Fatalln(resp)
// 	}

// 	log.Println("InquiryCustomerTopUpDisbursement Response:", resp, "\n")
// 	// End InquiryCustomerTopUpDisbursement

// 	// CustomerTopUpDisbursement
// 	// Retry intervals
// 	retryIntervals := []time.Duration{
// 		5 * time.Second,
// 		10 * time.Second,
// 		20 * time.Second,
// 		40 * time.Second,
// 		60 * time.Second,
// 	}

// 	reqTopUp := dana.RequestCustomerTopUpDisbursement{
// 		PartnerReferenceNo: uuid,
// 		CustomerNumber:     phone_number,
// 		Amount: dana.RequestCustomerTopUpDisbursementAmount{
// 			Value:    fmt.Sprintf("%.2f", float64(amount)),
// 			Currency: "IDR",
// 		},
// 		FeeAmount: dana.RequestCustomerTopUpDisbursementFeeAmount{
// 			Value:    "0.00",
// 			Currency: "IDR",
// 		},
// 		TransactionDate: dana.GenerateTimestamp(),
// 		AdditionalInfo: dana.RequestCustomerTopUpDisbursementAdditionalInfo{
// 			FundType: "AGENT_TOPUP_FOR_USER_SETTLE",
// 		},
// 	}

// 	for i, interval := range retryIntervals {
// 		respTopUp, err := danaClient.CustomerTopUpDisbursement(ctx, reqTopUp)
// 		if err == nil {
// 			log.Println("CustomerTopUpDisbursement Response:", respTopUp)
// 			break
// 		}

// 		log.Printf("Attempt %d failed: %v. Retrying in %v...\n", i+1, err, interval)
// 		select {
// 		case <-ctx.Done():
// 			log.Println("Context canceled or deadline exceeded, aborting retries.")
// 			return
// 		case <-time.After(interval):
// 			// continue to next retry
// 		}
// 	}

// 	if err != nil {
// 		log.Fatalf("Final attempt failed: %v\n", err)
// 	}

// 	log.Println("CustomerTopUpDisbursement Response:", respTopUp, "\n")
// 	// EndCustomerTopUpDisbursement
// }

// func UATScenario14(ctx context.Context, danaClient *dana.Client) {
// 	// Scenario Name: Getting an abnormal response
// 	var (
// 		respTopUp *dana.ResponseCustomerTopUpDisbursement
// 		err       error
// 	)

// 	uuid := uuid.NewString()
// 	phone_number := "628551001237"
// 	amount := 1

// 	// InquiryCustomerTopUpDisbursement
// 	req := dana.RequestAccountInquiryTopUpDisbursement{
// 		PartnerReferenceNo: uuid,
// 		CustomerNumber:     phone_number,
// 		Amount: dana.RequestCustomerTopUpDisbursementAmount{
// 			Value:    fmt.Sprintf("%.2f", float64(amount)),
// 			Currency: "IDR",
// 		},
// 		FeeAmount: dana.RequestCustomerTopUpDisbursementFeeAmount{
// 			Value:    "0.00",
// 			Currency: "IDR",
// 		},
// 		TransactionDate: dana.GenerateTimestamp(),
// 		AdditionalInfo: dana.RequestCustomerTopUpDisbursementAdditionalInfo{
// 			FundType: "AGENT_TOPUP_FOR_USER_SETTLE",
// 		},
// 	}

// 	resp, err := danaClient.AccountInquiryTopUpDisbursement(ctx, req)
// 	if err != nil {
// 		log.Fatalln(resp)
// 	}

// 	log.Println("InquiryCustomerTopUpDisbursement Response:", resp, "\n")
// 	// End InquiryCustomerTopUpDisbursement

// 	// CustomerTopUpDisbursement
// 	// Retry intervals
// 	retryIntervals := []time.Duration{
// 		5 * time.Second,
// 		10 * time.Second,
// 		20 * time.Second,
// 		40 * time.Second,
// 		60 * time.Second,
// 	}

// 	reqTopUp := dana.RequestCustomerTopUpDisbursement{
// 		PartnerReferenceNo: uuid,
// 		CustomerNumber:     phone_number,
// 		Amount: dana.RequestCustomerTopUpDisbursementAmount{
// 			Value:    fmt.Sprintf("%.2f", float64(amount)),
// 			Currency: "IDR",
// 		},
// 		FeeAmount: dana.RequestCustomerTopUpDisbursementFeeAmount{
// 			Value:    "0.00",
// 			Currency: "IDR",
// 		},
// 		TransactionDate: dana.GenerateTimestamp(),
// 		AdditionalInfo: dana.RequestCustomerTopUpDisbursementAdditionalInfo{
// 			FundType: "AGENT_TOPUP_FOR_USER_SETTLE",
// 		},
// 	}

// 	for i, interval := range retryIntervals {
// 		respTopUp, err := danaClient.CustomerTopUpDisbursement(ctx, reqTopUp)
// 		if err == nil {
// 			log.Println("CustomerTopUpDisbursement Response:", respTopUp)
// 			break
// 		}

// 		log.Printf("Attempt %d failed: %v. Retrying in %v...\n", i+1, err, interval)
// 		select {
// 		case <-ctx.Done():
// 			log.Println("Context canceled or deadline exceeded, aborting retries.")
// 			return
// 		case <-time.After(interval):
// 			// continue to next retry
// 		}
// 	}

// 	if err != nil {
// 		log.Fatalf("Final attempt failed: %v\n", err)
// 	}

// 	log.Println("CustomerTopUpDisbursement Response:", respTopUp, "\n")
// 	// EndCustomerTopUpDisbursement
// }

// func UATScenario15(ctx context.Context, danaClient *dana.Client) {
// 	// Scenario Name: Getting an abnormal response
// 	var (
// 		respTopUp *dana.ResponseCustomerTopUpDisbursement
// 		err       error
// 	)

// 	uuid := uuid.NewString()
// 	phone_number := "628121111111"
// 	amount := 1

// 	// InquiryCustomerTopUpDisbursement
// 	req := dana.RequestAccountInquiryTopUpDisbursement{
// 		PartnerReferenceNo: uuid,
// 		CustomerNumber:     phone_number,
// 		Amount: dana.RequestCustomerTopUpDisbursementAmount{
// 			Value:    fmt.Sprintf("%.2f", float64(amount)),
// 			Currency: "IDR",
// 		},
// 		FeeAmount: dana.RequestCustomerTopUpDisbursementFeeAmount{
// 			Value:    "0.00",
// 			Currency: "IDR",
// 		},
// 		TransactionDate: dana.GenerateTimestamp(),
// 		AdditionalInfo: dana.RequestCustomerTopUpDisbursementAdditionalInfo{
// 			FundType: "AGENT_TOPUP_FOR_USER_SETTLE",
// 		},
// 	}

// 	resp, err := danaClient.AccountInquiryTopUpDisbursement(ctx, req)
// 	if err != nil {
// 		log.Fatalln(resp)
// 	}

// 	log.Println("InquiryCustomerTopUpDisbursement Response:", resp, "\n")
// 	// End InquiryCustomerTopUpDisbursement

// 	// CustomerTopUpDisbursement
// 	// Retry intervals
// 	retryIntervals := []time.Duration{
// 		5 * time.Second,
// 		10 * time.Second,
// 		20 * time.Second,
// 		40 * time.Second,
// 		60 * time.Second,
// 	}

// 	reqTopUp := dana.RequestCustomerTopUpDisbursement{
// 		PartnerReferenceNo: uuid,
// 		CustomerNumber:     phone_number,
// 		Amount: dana.RequestCustomerTopUpDisbursementAmount{
// 			Value:    fmt.Sprintf("%.2f", float64(amount)),
// 			Currency: "IDR",
// 		},
// 		FeeAmount: dana.RequestCustomerTopUpDisbursementFeeAmount{
// 			Value:    "0.00",
// 			Currency: "IDR",
// 		},
// 		TransactionDate: dana.GenerateTimestamp(),
// 		AdditionalInfo: dana.RequestCustomerTopUpDisbursementAdditionalInfo{
// 			FundType: "AGENT_TOPUP_FOR_USER_SETTLE",
// 		},
// 	}

// 	for i, interval := range retryIntervals {
// 		respTopUp, err := danaClient.CustomerTopUpDisbursement(ctx, reqTopUp)
// 		if err == nil {
// 			log.Println("CustomerTopUpDisbursement Response:", respTopUp)
// 			break
// 		}

// 		log.Printf("Attempt %d failed: %v. Retrying in %v...\n", i+1, err, interval)
// 		select {
// 		case <-ctx.Done():
// 			log.Println("Context canceled or deadline exceeded, aborting retries.")
// 			return
// 		case <-time.After(interval):
// 			// continue to next retry
// 		}
// 	}

// 	if err != nil {
// 		log.Fatalf("Final attempt failed: %v\n", err)
// 	}

// 	log.Println("CustomerTopUpDisbursement Response:", respTopUp, "\n")
// 	// EndCustomerTopUpDisbursement
// }
