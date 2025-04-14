package main

import (
	"context"
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

	danaClient := dana.NewClient(MerchantID, ClientID, ClientSecret, PublicKey, PrivateKey, dana.Sandbox)

	ctx := context.Background()

	// InquiryCustomerTopUpDisbursement
	// req := dana.RequestInquiryCustomerTopUpDisbursement{
	// 	PartnerReferenceNo: uuid.NewString(),
	// 	CustomerNumber:     "6281328076003",
	// 	Amount: dana.RequestCustomerTopUpDisbursementAmount{
	// 		Value:    "1.00",
	// 		Currency: "IDR",
	// 	},
	// 	FeeAmount: dana.RequestCustomerTopUpDisbursementFeeAmount{
	// 		Value:    "0.00",
	// 		Currency: "IDR",
	// 	},
	// 	TransactionDate: dana.GenerateTimestamp(),
	// 	AdditionalInfo: dana.RequestCustomerTopUpDisbursementAdditionalInfo{
	// 		FundType: "AGENT_TOPUP_FOR_USER_SETTLE",
	// 	},
	// }

	// resp, err := danaClient.InquiryCustomerTopUpDisbursement(ctx, req)
	// if err != nil {
	// 	log.Fatalln(resp)
	// }
	// End InquiryCustomerTopUpDisbursement

	// CustomerTopUpDisbursement
	req := dana.RequestCustomerTopUpDisbursement{
		PartnerReferenceNo: uuid.NewString(),
		CustomerNumber:     "6281328076003",
		Amount: dana.RequestCustomerTopUpDisbursementAmount{
			Value:    "1.00",
			Currency: "IDR",
		},
		FeeAmount: dana.RequestCustomerTopUpDisbursementFeeAmount{
			Value:    "0.00",
			Currency: "IDR",
		},
		TransactionDate: dana.GenerateTimestamp(),
		AdditionalInfo: dana.RequestCustomerTopUpDisbursementAdditionalInfo{
			FundType: "AGENT_TOPUP_FOR_USER_SETTLE",
		},
	}

	resp, err := danaClient.CustomerTopUpDisbursement(ctx, req)
	if err != nil {
		log.Fatalln(err)
	}
	// EndCustomerTopUpDisbursement

	log.Printf("Disbursement: %+v", resp)
}
