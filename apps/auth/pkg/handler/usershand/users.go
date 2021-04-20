package usershand

import (
	"encoding/json"
	"net/http"

	"golang-seed/apps/auth/pkg/messagesconst"
	"golang-seed/apps/auth/pkg/service/usersserv"
	"golang-seed/pkg/httperror"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type UsersHandler struct {
	usersService *usersserv.UsersService
}

func NewUsersHandler(usersService *usersserv.UsersService) *UsersHandler {
	return &UsersHandler{usersService: usersService}
}

func (h UsersHandler) Get(w http.ResponseWriter, r *http.Request) error {
	pathVars := mux.Vars(r)
	id := pathVars["id"]

	uid, err := uuid.Parse(id)
	if err != nil {
		return httperror.ErrorCauseT(err, http.StatusBadRequest, messagesconst.GeneralErrorInvalidField, "id")
	}

	user, err := h.usersService.GetByID(uid)
	if err != nil {
		return err
	}

	user.Password = ""

	body, err := json.Marshal(user)
	if err != nil {
		return httperror.ErrorCauseT(err, http.StatusInternalServerError, messagesconst.GeneralErrorMarshal)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
	return nil
}

func (h UsersHandler) GetAll(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (h UsersHandler) GetAllPaged(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (h UsersHandler) Create(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (h UsersHandler) Update(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (h UsersHandler) Delete(w http.ResponseWriter, r *http.Request) error {
	return nil
}
