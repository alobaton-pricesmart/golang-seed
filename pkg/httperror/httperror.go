package httperror

import (
	"encoding/json"
	"fmt"
	"golang-seed/pkg/messages"
)

// HTTPError implements ClientError interface.
type HTTPError struct {
	Cause            error  `json:"-"`
	CauseMessage     string `json:"error"`
	ErrorDescription string `json:"error_description"`
	Status           int    `json:"status"`
}

// Allows StatusError to satisfy the error interface.
func (e HTTPError) Error() string {
	if e.Cause == nil {
		return e.ErrorDescription
	}
	return e.ErrorDescription + " : " + e.Cause.Error()
}

// ResponseBody returns JSON response body.
func (e *HTTPError) ResponseBody() ([]byte, error) {
	body, err := json.Marshal(e)
	if err != nil {
		return nil, fmt.Errorf("Error while parsing response body: %v", err)
	}
	return body, nil
}

// ResponseHeaders returns http status code and headers.
func (e *HTTPError) ResponseHeaders() (int, map[string]string) {
	return e.Status, map[string]string{
		"Content-Type": "application/json; charset=utf-8",
	}
}

// Error Creates an HTTPError
func Error(status int, errorDescription string) error {
	return &HTTPError{
		ErrorDescription: errorDescription,
		Status:           status,
	}
}

// ErrorCase Creates an HTTPError with underliting error
func ErrorCause(err error, status int, errorDescription string) error {
	return &HTTPError{
		Cause:            err,
		CauseMessage:     err.Error(),
		ErrorDescription: errorDescription,
		Status:           status,
	}
}

// ErrorT Creates an HTTPError and traslate the key and args
func ErrorT(status int, key string, args ...string) error {
	return &HTTPError{
		ErrorDescription: messages.Get(key, translate(args)...),
		Status:           status,
	}
}

// ErrorCauseT Creates an HTTPError with underliting error and traslate the key and args
func ErrorCauseT(err error, status int, key string, args ...string) error {
	return &HTTPError{
		Cause:            err,
		CauseMessage:     err.Error(),
		ErrorDescription: messages.Get(key, translate(args)...),
		Status:           status,
	}
}

func translate(args []string) []interface{} {
	t := []interface{}{}
	for _, arg := range args {
		t = append(t, messages.Get(arg))
	}
	return t
}
