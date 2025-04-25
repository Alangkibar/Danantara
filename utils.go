package dana

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"time"
)

// BaseUrl To get Midtrans Base URL
func (e EnvironmentType) BaseUrl() string {
	for k, v := range typeString {
		if k == e {
			return v
		}
	}
	return "undefined"
}

func GenerateTimestamp() string {
	// Get the current time in Jakarta time zone (UTC+7)
	jakartaLoc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		fmt.Println("Error loading Jakarta time zone:", err)
		return ""
	}

	currentTime := time.Now().In(jakartaLoc)

	// Format the time according to the specified format
	timestamp := currentTime.Format("2006-01-02T15:04:05+07:00")

	return timestamp
}

func GetSignature(method string, path string, timestamp string, body string) (string, error) {
	// Hash the minified body
	hash := sha256.New()
	hash.Write([]byte(body))
	hashedPayload := fmt.Sprintf("%x", hash.Sum(nil)) // hex encode and lowercase

	// data = '<HTTP METHOD> + ”:” + <RELATIVE PATH URL> + “:“ + LowerCase(HexEncode(SHA-256(Minify(<HTTP BODY>)))) + “:“ + <X-TIMESTAMP>';
	data := method + ":" + path + ":" + hashedPayload + ":" + timestamp

	signature, err := GenerateSignature([]byte(data))
	if err != nil {
		fmt.Println("signature generation error", err)
		return "", err
	}

	return signature, nil
}

func GenerateSignature(data []byte) (string, error) {
	signer, err := parsePrivateKey(RsaPrivateKey)
	if err != nil {
		return "", err
	}

	signed, err := signer.Sign(data)
	if err != nil {
		return "", err
	}

	sig := base64.StdEncoding.EncodeToString(signed)

	return sig, nil
}

func parsePrivateKey(pemBytes []byte) (Signer, error) {
	block, _ := pem.Decode(pemBytes)
	if block == nil {
		return nil, errors.New("ssh: no key found")
	}

	var rawkey interface{}
	switch block.Type {
	case "RSA PRIVATE KEY":
		rsa, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		rawkey = rsa
	default:
		return nil, fmt.Errorf("ssh: unsupported key type %q", block.Type)
	}
	return newSignerFromKey(rawkey)
}

// A Signer is can create signatures that verify against a public key.
type Signer interface {
	// Sign returns raw signature for the given data. This method
	// will apply the hash specified for the keytype to the data.
	Sign(data []byte) ([]byte, error)
}

func newSignerFromKey(k interface{}) (Signer, error) {
	var sshKey Signer
	switch t := k.(type) {
	case *rsa.PrivateKey:
		sshKey = &rsaPrivateKey{t}
	default:
		return nil, fmt.Errorf("ssh: unsupported key type %T", k)
	}
	return sshKey, nil
}

type rsaPrivateKey struct {
	*rsa.PrivateKey
}

// Sign signs data with rsa-sha256
func (r *rsaPrivateKey) Sign(data []byte) ([]byte, error) {
	h := sha256.New()
	h.Write(data)
	d := h.Sum(nil)
	return rsa.SignPKCS1v15(rand.Reader, r.PrivateKey, crypto.SHA256, d)
}

// These response code mapping refer to DANA API Documentation
// https://dashboard.dana.id/api-docs/read/107
var DanaResponseCodeMap = map[string]ResponseCodeInfo{
	"2003700":    {"2003700", "Successful", "Success to be processed", "Mark process as Success"},
	"4003700":    {"4003700", "Bad Request", "General request failed error", "Retry request with proper parameter"},
	"4003701":    {"4003701", "Invalid Field Format", "Invalid format for certain field", "Retry request with proper parameter"},
	"4003702":    {"4003702", "Invalid Mandatory Field", "Missing or invalid format", "Retry request with proper parameter"},
	"4013700":    {"4013700", "Unauthorized", "General unauthorized error", "Retry request with proper parameter"},
	"4013701":    {"4013701", "Invalid Token (B2B)", "Token invalid/expired", "Retry request with proper parameter"},
	"4013702":    {"4013702", "Invalid Customer Token", "Customer token invalid/expired", "Retry request with proper parameter"},
	"4013704":    {"4013704", "Customer Token Not Found", "Token not found in system", "Retry request with proper parameter"},
	"4033702":    {"4033702", "Exceeds Transaction Amount Limit", "Transaction exceeds limit", "Retry with proper amount"},
	"4033705":    {"4033705", "Do Not Honor", "User/account status abnormal", "Contact DANA to check user/account"},
	"4033714":    {"4033714", "Insufficient Funds", "User has insufficient funds", "Contact DANA to top up"},
	"4033715":    {"4033715", "Transaction Not Permitted", "Transaction not allowed", "Contact DANA to check merchant config"},
	"4033718":    {"4033718", "Inactive Card/Account/Customer", "Account inactive", "Contact DANA"},
	"4033720":    {"4033720", "Merchant Limit Exceed", "Daily limit exceeded", "Contact DANA"},
	"4043708":    {"4043708", "Invalid Merchant", "Merchant not found or abnormal", "Contact DANA"},
	"4043711":    {"4043711", "Invalid Card/Account/Customer", "Card/account info invalid", "Contact DANA"},
	"4293700":    {"4293700", "Too Many Requests", "Max transaction rate exceeded", "Retry"},
	"5003700":    {"5003700", "General Error", "Non-retryable error", "Retry new Account Inquiry"},
	"5003701":    {"5003701", "Internal Server Error", "Retryable internal error", "Retry new Account Inquiry"},
	"TIMEOUT":    {"TIMEOUT", "Total Timeout", "Client received no response", "Retry"},
	"UNEXPECTED": {"UNEXPECTED", "Unexpected response", "Unexpected field/code", "Retry"},
}

func GetDanaResponseInfo(code string) ResponseCodeInfo {
	if val, ok := DanaResponseCodeMap[code]; ok {
		return val
	}
	return DanaResponseCodeMap["UNEXPECTED"]
}
