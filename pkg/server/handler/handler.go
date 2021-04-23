package handler

import "net/http"

func Paramsr(r *http.Request) map[string]interface{} {
	params := make(map[string]interface{})
	for k, v := range r.URL.Query() {
		params[k] = v
	}

	return params
}
