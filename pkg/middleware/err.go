package middleware

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

// httpError is an error whose details to be shared with client.
type httpError interface {
	Error() string
	// ResponseBody returns response body.
	ResponseBody() ([]byte, error)
	// ResponseHeaders returns http status code and headers.
	ResponseHeaders() (int, map[string]string)
}

// ErrorHandlerFunc type of the functions ErrorHandler wraps.
type ErrorHandlerFunc func(http.ResponseWriter, *http.Request) error

// ErrorHandler use as a wrapper around the ErrorHandlerFunc functions.
func ErrorHandler(next ErrorHandlerFunc) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		err := next(w, r)

		if err == nil {
			return
		}

		log.WithField("error", err).Error("an error accured")

		httpError, ok := err.(httpError)
		if !ok {
			w.WriteHeader(500)
			return
		}

		body, err := httpError.ResponseBody()
		if err != nil {
			log.WithField("error", err).Error("an error accured")
			w.WriteHeader(500)
			return
		}

		status, headers := httpError.ResponseHeaders()
		for k, v := range headers {
			w.Header().Set(k, v)
		}
		w.WriteHeader(status)
		w.Write(body)
	}

	return http.HandlerFunc(fn)
}
