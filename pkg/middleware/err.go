package middleware

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

// ClientError is an error whose details to be shared with client.
type ClientError interface {
	Error() string
	// ResponseBody returns response body.
	ResponseBody() ([]byte, error)
	// ResponseHeaders returns http status code and headers.
	ResponseHeaders() (int, map[string]string)
}

// Use as a wrapper around the handler functions.
type ErrorHandlerFunc func(http.ResponseWriter, *http.Request) error

func ErrorHandler(next ErrorHandlerFunc) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		err := next(w, r)

		if err == nil {
			return
		}

		// This is where our error handling logic starts.
		log.WithField("error", err).Error("an error accured")

		// Check if it is a ClientError.
		clientError, ok := err.(ClientError)
		if !ok {
			// If the error is not ClientError, assume that it is ServerError.
			// return 500 Internal Server Error.
			w.WriteHeader(500)
			return
		}

		// Try to get response body of ClientError.
		body, err := clientError.ResponseBody()
		if err != nil {
			log.WithField("error", err).Error("an error accured")
			w.WriteHeader(500)
			return
		}

		// Get http status code and headers.
		status, headers := clientError.ResponseHeaders()
		for k, v := range headers {
			w.Header().Set(k, v)
		}
		w.WriteHeader(status)
		w.Write(body)
	}

	return http.HandlerFunc(fn)
}
