package tserror

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type TSError struct {
	ErrType TSErrorTypes `json:"errType"`
	ErrMsg  string       `json:"errMsg"`
	Details string       `json:"details"`
}

func (e TSError) Error() string {
	errBytes, err := json.Marshal(e)
	if err != nil {
		return fmt.Sprintf("ErrorType:%v ErrorEnum:%v ErrorDetails:%v", e.ErrType, e.ErrMsg, e.Details)
	} else {
		return string(errBytes)
	}
}

func (e TSError) WithDetails(details string) *TSError {
	e.Details = details
	return &e
}

func (e TSError) GetHTTPErrorCode() int {
	if e.ErrType == ErrorTypeInvalidArgument || e.ErrType == ErrorTypeNotSupported {
		return http.StatusBadRequest
	}
	return http.StatusInternalServerError
}
