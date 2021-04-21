package authhand

import (
	"golang-seed/apps/auth/pkg/messagesconst"
	"golang-seed/apps/auth/pkg/models"
	"golang-seed/apps/auth/pkg/service/usersserv"
	"golang-seed/pkg/httperror"
	"net/http"

	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/server"
	"golang.org/x/crypto/bcrypt"
)

type (
	AuthHandler struct {
		srv          *server.Server
		usersService *usersserv.UsersService
	}
	// httpError is an error whose details to be shared with client.
	httpError interface {
		Error() string
		// ResponseBody returns response body.
		ResponseBody() ([]byte, error)
		// ResponseHeaders returns http status code and headers.
		ResponseHeaders() (int, map[string]string)
	}
)

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

func (a AuthHandler) InternalErrorHandler(err error) *errors.Response {
	httpError, ok := err.(httpError)
	if !ok {
		return nil
	}

	status, _ := httpError.ResponseHeaders()
	return &errors.Response{
		Error:       err,
		Description: httpError.Error(),
		StatusCode:  status,
	}
}

func (a AuthHandler) ValidateToken(r *http.Request) error {
	_, err := a.srv.ValidationBearerToken(r)
	if err != nil {
		return httperror.ErrorCauseT(err, http.StatusUnauthorized, messagesconst.OauthInvalidToken)
	}

	return nil
}

func (a AuthHandler) ValidatePermission(r *http.Request, permission string) error {
	tokenInfo, err := a.srv.ValidationBearerToken(r)
	if err != nil {
		return httperror.ErrorCauseT(err, http.StatusUnauthorized, messagesconst.OauthInvalidToken)
	}

	// Consultar los permisos que tiene el usuario.
	tokenInfo.GetUserID()

	return nil
}
