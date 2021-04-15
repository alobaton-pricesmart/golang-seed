package authhand

import (
	"net/http"

	"github.com/go-oauth2/oauth2/v4/server"
)

type AuthHandler struct {
	srv *server.Server
}

func NewAuthHandler(srv *server.Server) *AuthHandler {
	return &AuthHandler{srv}
}

func (a AuthHandler) Authorize(w http.ResponseWriter, r *http.Request) {
	err := a.srv.HandleAuthorizeRequest(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func (a AuthHandler) Token(w http.ResponseWriter, r *http.Request) {
	a.srv.HandleTokenRequest(w, r)
}
