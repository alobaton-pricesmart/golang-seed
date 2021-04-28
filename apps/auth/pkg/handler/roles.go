package handler

import (
	"encoding/json"
	"golang-seed/apps/auth/pkg/messagesconst"
	"golang-seed/apps/auth/pkg/models"
	"golang-seed/apps/auth/pkg/service"
	"golang-seed/pkg/httperror"
	"golang-seed/pkg/pagination"
	"golang-seed/pkg/server/handler"
	"golang-seed/pkg/sorting"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type RolesHandler struct {
	rolesService *service.RolesService
}

func NewRolesHandler(rolesService *service.RolesService) *RolesHandler {
	return &RolesHandler{rolesService: rolesService}
}

func (h RolesHandler) Get(w http.ResponseWriter, r *http.Request) error {
	pathVars := mux.Vars(r)
	id := pathVars["id"]

	_, err := uuid.Parse(id)
	if err != nil {
		return httperror.ErrorCauseT(err, http.StatusBadRequest, messagesconst.GeneralErrorInvalidField, "id")
	}

	role, err := h.rolesService.GetByID(id)
	if err != nil {
		return err
	}

	body, err := json.Marshal(role)
	if err != nil {
		return httperror.ErrorCauseT(err, http.StatusInternalServerError, messagesconst.GeneralErrorMarshal)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
	return nil
}

func (h RolesHandler) GetAll(w http.ResponseWriter, r *http.Request) error {
	sort := sorting.Sortr(r)
	params := handler.Paramsr(r)

	roles, err := h.rolesService.GetAll(params, sort)
	if err != nil {
		return err
	}

	body, err := json.Marshal(roles)
	if err != nil {
		return httperror.ErrorCauseT(err, http.StatusInternalServerError, messagesconst.GeneralErrorMarshal)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
	return nil
}

func (h RolesHandler) GetAllPaged(w http.ResponseWriter, r *http.Request) error {
	pageable, _ := pagination.Pageabler(r)
	sort := sorting.Sortr(r)
	params := handler.Paramsr(r)

	page, err := h.rolesService.GetAllPaged(params, sort, pageable)
	if err != nil {
		return err
	}

	body, err := json.Marshal(page)
	if err != nil {
		return httperror.ErrorCauseT(err, http.StatusInternalServerError, messagesconst.GeneralErrorMarshal)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
	return nil
}

func (h RolesHandler) Create(w http.ResponseWriter, r *http.Request) error {
	role := &models.Role{}
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	err := decoder.Decode(role)
	if err != nil {
		return httperror.ErrorCauseT(err, http.StatusBadRequest, messagesconst.GeneralErrorMarshal)
	}

	err = h.rolesService.Create(role)
	if err != nil {
		return err
	}

	body, err := json.Marshal(role)
	if err != nil {
		return httperror.ErrorCauseT(err, http.StatusInternalServerError, messagesconst.GeneralErrorMarshal)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
	return nil
}

func (h RolesHandler) Update(w http.ResponseWriter, r *http.Request) error {
	role := &models.Role{}
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	err := decoder.Decode(role)
	if err != nil {
		return httperror.ErrorCauseT(err, http.StatusBadRequest, messagesconst.GeneralErrorMarshal)
	}

	err = h.rolesService.Update(role)
	if err != nil {
		return err
	}

	body, err := json.Marshal(role)
	if err != nil {
		return httperror.ErrorCauseT(err, http.StatusInternalServerError, messagesconst.GeneralErrorMarshal)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
	return nil
}

func (h RolesHandler) Delete(w http.ResponseWriter, r *http.Request) error {
	pathVars := mux.Vars(r)
	id := pathVars["id"]

	err := h.rolesService.Delete(id)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	return nil
}
