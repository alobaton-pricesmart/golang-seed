package handler

import (
	"fmt"
	"net/http"

	"golang-seed/apps/auth/pkg/authconst"
	"golang-seed/apps/auth/pkg/config"
	"golang-seed/apps/auth/pkg/models"
	"golang-seed/apps/auth/pkg/service"
	"golang-seed/pkg/httperror"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/generates"
	"github.com/go-oauth2/oauth2/v4/server"
	"golang.org/x/crypto/bcrypt"
)

type (
	AuthHandler struct {
		srv          *server.Server
		usersService *service.UsersService
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

func NewAuthHandler(srv *server.Server, usersService *service.UsersService) *AuthHandler {
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
		return "", httperror.ErrorCauseT(err, http.StatusUnauthorized, authconst.OAuthInvalidUsernamePassword)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", httperror.ErrorT(http.StatusUnauthorized, authconst.OAuthInvalidUsernamePassword)
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
	tokenInfo, err := a.srv.ValidationBearerToken(r)
	if err != nil {
		return httperror.ErrorCauseT(err, http.StatusUnauthorized, authconst.OauthInvalidToken)
	}

	// Parse and verify jwt access token
	token, err := jwt.ParseWithClaims(tokenInfo.GetAccess(), &generates.JWTAccessClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("parse error")
		}
		return []byte(config.Settings.Security.Key), nil
	})
	if err != nil {
		httperror.ErrorCauseT(err, http.StatusUnauthorized, authconst.OauthInvalidToken)
	}

	_, ok := token.Claims.(*generates.JWTAccessClaims)
	if !ok || !token.Valid {
		httperror.ErrorCauseT(err, http.StatusUnauthorized, authconst.OauthInvalidToken)
	}

	return nil
}

func (a AuthHandler) ValidatePermission(r *http.Request, permission string) error {
	tokenInfo, err := a.srv.ValidationBearerToken(r)
	if err != nil {
		return httperror.ErrorCauseT(err, http.StatusUnauthorized, authconst.OauthInvalidToken)
	}

	// Consultar los permisos que tiene el usuario.
	tokenInfo.GetUserID()

	return nil
}
