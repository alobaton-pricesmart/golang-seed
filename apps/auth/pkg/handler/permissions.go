package handler

import (
	"encoding/json"
	"golang-seed/apps/auth/pkg/authconst"
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

type PermissionsHandler struct {
	permissionsService *service.PermissionsService
}

func NewPermissionsHandler(permissionsService *service.PermissionsService) *PermissionsHandler {
	return &PermissionsHandler{permissionsService: permissionsService}
}

func (h PermissionsHandler) Get(w http.ResponseWriter, r *http.Request) error {
	pathVars := mux.Vars(r)
	id := pathVars["id"]

	_, err := uuid.Parse(id)
	if err != nil {
		return httperror.ErrorCauseT(err, http.StatusBadRequest, authconst.GeneralErrorInvalidField, "id")
	}

	permission, err := h.permissionsService.GetByID(id)
	if err != nil {
		return err
	}

	body, err := json.Marshal(permission)
	if err != nil {
		return httperror.ErrorCauseT(err, http.StatusInternalServerError, authconst.GeneralErrorMarshal)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
	return nil
}

func (h PermissionsHandler) GetAll(w http.ResponseWriter, r *http.Request) error {
	sort := sorting.Sortr(r)
	params := handler.Paramsr(r)

	permissions, err := h.permissionsService.GetAll(params, sort)
	if err != nil {
		return err
	}

	body, err := json.Marshal(permissions)
	if err != nil {
		return httperror.ErrorCauseT(err, http.StatusInternalServerError, authconst.GeneralErrorMarshal)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
	return nil
}

func (h PermissionsHandler) GetAllPaged(w http.ResponseWriter, r *http.Request) error {
	pageable, _ := pagination.Pageabler(r)
	sort := sorting.Sortr(r)
	params := handler.Paramsr(r)

	page, err := h.permissionsService.GetAllPaged(params, sort, pageable)
	if err != nil {
		return err
	}

	body, err := json.Marshal(page)
	if err != nil {
		return httperror.ErrorCauseT(err, http.StatusInternalServerError, authconst.GeneralErrorMarshal)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
	return nil
}

func (h PermissionsHandler) Create(w http.ResponseWriter, r *http.Request) error {
	permission := &models.Permission{}
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	err := decoder.Decode(permission)
	if err != nil {
		return httperror.ErrorCauseT(err, http.StatusBadRequest, authconst.GeneralErrorMarshal)
	}

	err = h.permissionsService.Create(permission)
	if err != nil {
		return err
	}

	body, err := json.Marshal(permission)
	if err != nil {
		return httperror.ErrorCauseT(err, http.StatusInternalServerError, authconst.GeneralErrorMarshal)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
	return nil
}

func (h PermissionsHandler) Update(w http.ResponseWriter, r *http.Request) error {
	permission := &models.Permission{}
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	err := decoder.Decode(permission)
	if err != nil {
		return httperror.ErrorCauseT(err, http.StatusBadRequest, authconst.GeneralErrorMarshal)
	}

	err = h.permissionsService.Update(permission)
	if err != nil {
		return err
	}

	body, err := json.Marshal(permission)
	if err != nil {
		return httperror.ErrorCauseT(err, http.StatusInternalServerError, authconst.GeneralErrorMarshal)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
	return nil
}

func (h PermissionsHandler) Delete(w http.ResponseWriter, r *http.Request) error {
	pathVars := mux.Vars(r)
	id := pathVars["id"]

	err := h.permissionsService.Delete(id)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	return nil
}
