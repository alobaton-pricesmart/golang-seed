package middleware

import "net/http"

// HeaderHandler allow us to set up Headers (like Content-Type) around our responses or requests.
func HeaderHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
