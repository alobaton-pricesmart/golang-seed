package middleware

import (
	"net/http"

	"golang-seed/apps/auth/pkg/authconst"
	"golang-seed/pkg/httperror"

	log "github.com/sirupsen/logrus"
)

// RecoverHandler allow us to recover after a panic.
func RecoverHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %+v", err)

				httpError, ok := httperror.ErrorCauseT(err.(error), http.StatusInternalServerError, authconst.GeneralErrorInternalServerError).(httpError)
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
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
