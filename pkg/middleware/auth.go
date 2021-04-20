package middleware

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

type (
	// ValidateTokenFunc
	ValidateTokenFunc func(*http.Request) error

	ValidateTokenPermissionFunc func(*http.Request, string) error
)

func AuthorizeHandler(next http.Handler, permission string, vtpf ValidateTokenPermissionFunc) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		err := vtpf(r, permission)

		if err == nil {
			next.ServeHTTP(w, r)
			return
		}

		log.WithField("error", err).Error("an error accured")

		writeError(w, err)
	}

	return http.HandlerFunc(fn)
}

func AuthenticationHandler(next http.Handler, vtf ValidateTokenFunc) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		err := vtf(r)

		if err == nil {
			next.ServeHTTP(w, r)
			return
		}

		log.WithField("error", err).Error("an error accured")

		writeError(w, err)
	}

	return http.HandlerFunc(fn)
}
