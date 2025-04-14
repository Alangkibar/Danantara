package dana

import (
	"errors"
	"fmt"

	"github.com/go-resty/resty/v2"
)

type ErrorAdditionalInfo struct {
	ResultMsg string `json:"resultMsg"`
}

type Error struct {
	Code           string              `json:"responseCode"`
	Message        string              `json:"responseMessage"`
	Reason         string              `json:"-"`
	AdditionalInfo ErrorAdditionalInfo `json:"additionalInfo"`
	RawError       error               `json:"-"`
	RawResponse    *resty.Response     `json:"-"`
}

func (e *Error) Error() string {
	if e.RawError != nil {
		return fmt.Sprintf("%s: %s", e.Message, e.RawError.Error())
	}
	return e.Message
}

func (e *Error) Unwrap() error {
	return e.RawError
}

func IsError(err error) bool {
	var danaError *Error
	return errors.As(err, &danaError)
}
