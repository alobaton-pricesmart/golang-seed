package middleware

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

type (
	// ValidateTokenFunc type to define a function that gets the token from http.Request and
	ValidateTokenFunc func(*http.Request) error

	// ValidatePermissionFunc
	ValidatePermissionFunc func(*http.Request, string) error
)

func AuthenticationHandler(fn ValidateTokenFunc) func(h http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			err := fn(r)

			if err == nil {
				next.ServeHTTP(w, r)
				return
			}

			log.WithField("error", err).Error("an error accured")

			writeError(w, err)
		})
	}
}

func AuthorizeHandler(permission string, fn ValidatePermissionFunc) func(h http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			err := fn(r, permission)

			if err == nil {
				next.ServeHTTP(w, r)
				return
			}

			log.WithField("error", err).Error("an error accured")

			writeError(w, err)
		})
	}
}
