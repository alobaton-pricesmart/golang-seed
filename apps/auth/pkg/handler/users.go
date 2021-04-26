package handler

import (
	"encoding/json"
	"net/http"

	"golang-seed/apps/auth/pkg/messagesconst"
	"golang-seed/apps/auth/pkg/models"
	"golang-seed/apps/auth/pkg/service"
	"golang-seed/pkg/httperror"
	"golang-seed/pkg/pagination"
	"golang-seed/pkg/server/handler"
	"golang-seed/pkg/sorting"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type UsersHandler struct {
	usersService *service.UsersService
}

func NewUsersHandler(usersService *service.UsersService) *UsersHandler {
	return &UsersHandler{usersService: usersService}
}

func (h UsersHandler) Get(w http.ResponseWriter, r *http.Request) error {
	pathVars := mux.Vars(r)
	id := pathVars["id"]

	_, err := uuid.Parse(id)
	if err != nil {
		return httperror.ErrorCauseT(err, http.StatusBadRequest, messagesconst.GeneralErrorInvalidField, "id")
	}

	user, err := h.usersService.GetByID(id)
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
	sort := sorting.Sortr(r)
	params := handler.Paramsr(r)

	users, err := h.usersService.GetAll(params, sort)
	if err != nil {
		return err
	}

	for _, user := range users {
		user.Password = ""
	}

	body, err := json.Marshal(users)
	if err != nil {
		return httperror.ErrorCauseT(err, http.StatusInternalServerError, messagesconst.GeneralErrorMarshal)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
	return nil
}

func (h UsersHandler) GetAllPaged(w http.ResponseWriter, r *http.Request) error {
	pageable, _ := pagination.Pageabler(r)
	sort := sorting.Sortr(r)
	params := handler.Paramsr(r)

	page, err := h.usersService.GetAllPaged(params, sort, pageable)
	if err != nil {
		return err
	}

	users, ok := page.Content.([]*models.User)
	if ok {
		for _, user := range users {
			user.Password = ""
		}
		page.Content = users
	}

	body, err := json.Marshal(page)
	if err != nil {
		return httperror.ErrorCauseT(err, http.StatusInternalServerError, messagesconst.GeneralErrorMarshal)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
	return nil
}

func (h UsersHandler) Create(w http.ResponseWriter, r *http.Request) error {
	user := &models.User{}
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	err := decoder.Decode(user)
	if err != nil {
		return httperror.ErrorCauseT(err, http.StatusBadRequest, messagesconst.GeneralErrorMarshal)
	}

	user.ID = uuid.NewString()
	err = h.usersService.Create(user)
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

func (h UsersHandler) Update(w http.ResponseWriter, r *http.Request) error {
	user := &models.User{}
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	err := decoder.Decode(user)
	if err != nil {
		return httperror.ErrorCauseT(err, http.StatusBadRequest, messagesconst.GeneralErrorMarshal)
	}

	err = h.usersService.Update(user)
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

func (h UsersHandler) Delete(w http.ResponseWriter, r *http.Request) error {
	pathVars := mux.Vars(r)
	id := pathVars["id"]

	err := h.usersService.Delete(id)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	return nil
}
