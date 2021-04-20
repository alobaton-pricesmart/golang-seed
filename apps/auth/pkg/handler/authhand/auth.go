package authhand

import (
	"golang-seed/apps/auth/pkg/messagesconst"
	"golang-seed/apps/auth/pkg/models"
	"golang-seed/apps/auth/pkg/service/usersserv"
	"golang-seed/pkg/httperror"
	"net/http"

	"github.com/go-oauth2/oauth2/v4/server"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	srv          *server.Server
	usersService *usersserv.UsersService
}

func NewAuthHandler(srv *server.Server, usersService *usersserv.UsersService) *AuthHandler {
	return &AuthHandler{srv, usersService}
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

func (a AuthHandler) PasswordAuthorizationHandler(username, password string) (string, error) {
	var err error
	user := &models.User{
		Nickname: username,
	}
	err = a.usersService.Get(user)
	if err != nil {
		return "", httperror.ErrorCauseT(err, http.StatusUnauthorized, messagesconst.OAuthInvalidUsernamePassword)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", httperror.ErrorT(http.StatusUnauthorized, messagesconst.OAuthInvalidUsernamePassword)
	}

	return user.ID, nil
}
