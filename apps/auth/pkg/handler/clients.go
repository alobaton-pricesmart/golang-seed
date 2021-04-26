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

	"github.com/gorilla/mux"
)

type ClientsHandler struct {
	clientsService *service.ClientsService
}

func NewClientsHandler(clientsService *service.ClientsService) *ClientsHandler {
	return &ClientsHandler{clientsService: clientsService}
}

func (h ClientsHandler) Get(w http.ResponseWriter, r *http.Request) error {
	pathVars := mux.Vars(r)
	id := pathVars["id"]

	client, err := h.clientsService.GetByID(id)
	if err != nil {
		return err
	}

	client.Secret = ""

	body, err := json.Marshal(client)
	if err != nil {
		return httperror.ErrorCauseT(err, http.StatusInternalServerError, messagesconst.GeneralErrorMarshal)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
	return nil
}

func (h ClientsHandler) GetAll(w http.ResponseWriter, r *http.Request) error {
	sort := sorting.Sortr(r)
	params := handler.Paramsr(r)

	clients, err := h.clientsService.GetAll(params, sort)
	if err != nil {
		return err
	}

	for _, client := range clients {
		client.Secret = ""
	}

	body, err := json.Marshal(clients)
	if err != nil {
		return httperror.ErrorCauseT(err, http.StatusInternalServerError, messagesconst.GeneralErrorMarshal)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
	return nil
}

func (h ClientsHandler) GetAllPaged(w http.ResponseWriter, r *http.Request) error {
	pageable, _ := pagination.Pageabler(r)
	sort := sorting.Sortr(r)
	params := handler.Paramsr(r)

	page, err := h.clientsService.GetAllPaged(params, sort, pageable)
	if err != nil {
		return err
	}

	clients, ok := page.Content.([]*models.Client)
	if ok {
		for _, client := range clients {
			client.Secret = ""
		}
		page.Content = clients
	}

	body, err := json.Marshal(page)
	if err != nil {
		return httperror.ErrorCauseT(err, http.StatusInternalServerError, messagesconst.GeneralErrorMarshal)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
	return nil
}

func (h ClientsHandler) Create(w http.ResponseWriter, r *http.Request) error {
	client := &models.Client{}
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	err := decoder.Decode(client)
	if err != nil {
		return httperror.ErrorCauseT(err, http.StatusBadRequest, messagesconst.GeneralErrorMarshal)
	}

	err = h.clientsService.Create(client)
	if err != nil {
		return err
	}

	client.Secret = ""

	body, err := json.Marshal(client)
	if err != nil {
		return httperror.ErrorCauseT(err, http.StatusInternalServerError, messagesconst.GeneralErrorMarshal)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
	return nil
}

func (h ClientsHandler) Update(w http.ResponseWriter, r *http.Request) error {
	client := &models.Client{}
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	err := decoder.Decode(client)
	if err != nil {
		return httperror.ErrorCauseT(err, http.StatusBadRequest, messagesconst.GeneralErrorMarshal)
	}

	err = h.clientsService.Update(client)
	if err != nil {
		return err
	}

	client.Secret = ""

	body, err := json.Marshal(client)
	if err != nil {
		return httperror.ErrorCauseT(err, http.StatusInternalServerError, messagesconst.GeneralErrorMarshal)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
	return nil
}

func (h ClientsHandler) Delete(w http.ResponseWriter, r *http.Request) error {
	pathVars := mux.Vars(r)
	id := pathVars["id"]

	err := h.clientsService.Delete(id)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	return nil
}
