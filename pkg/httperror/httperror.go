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

func NewHTTPError(err error, status int, errorDescription string) error {
	return &HTTPError{
		Cause:            err,
		CauseMessage:     err.Error(),
		ErrorDescription: errorDescription,
		Status:           status,
	}
}

func NewHTTPErrorT(err error, status int, key string, args ...string) error {
	t := translate(args)
	a := make([]interface{}, len(t))
	for i, v := range t {
		a[i] = v
	}
	return &HTTPError{
		Cause:            err,
		CauseMessage:     err.Error(),
		ErrorDescription: messages.Get(key, a...),
		Status:           status,
	}
}

func translate(args []string) []string {
	t := []string{}
	for _, arg := range args {
		t = append(t, messages.Get(arg))
	}
	return t
}
