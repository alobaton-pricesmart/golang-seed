package auth

import (
	"net/http"

	"github.com/go-oauth2/oauth2/v4/server"
)

type AuthService struct {
	srv *server.Server
}

func NewAuthService(srv *server.Server) *AuthService {
	return &AuthService{srv}
}

func (a AuthService) Authorize(w http.ResponseWriter, r *http.Request) {
	err := a.srv.HandleAuthorizeRequest(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func (a AuthService) Token(w http.ResponseWriter, r *http.Request) {
	a.srv.HandleTokenRequest(w, r)
}
