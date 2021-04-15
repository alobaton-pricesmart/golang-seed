package clientshand

import (
	"encoding/json"
	"net/http"

	"golang-seed/apps/auth/pkg/service/clientsserv"
	"golang-seed/pkg/httperror"

	"github.com/gorilla/mux"
)

type ClientsHandler struct {
	clientsService *clientsserv.ClientsService
}

func NewClientsHandler(clientsService *clientsserv.ClientsService) *ClientsHandler {
	return &ClientsHandler{clientsService: clientsService}
}

func (u ClientsHandler) Get(w http.ResponseWriter, r *http.Request) error {
	pathVars := mux.Vars(r)
	id := pathVars["id"]

	client, err := u.clientsService.Get(id)
	if err != nil {
		return httperror.NewHTTPError(err, http.StatusInternalServerError, "Error getting data from database")
	}

	body, err := json.Marshal(client)
	if err != nil {
		return httperror.NewHTTPError(err, http.StatusInternalServerError, "Error serializing data")
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
	return nil
}

func (u ClientsHandler) GetAll(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (u ClientsHandler) GetAllPaged(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (u ClientsHandler) Create(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (u ClientsHandler) Update(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (u ClientsHandler) Delete(w http.ResponseWriter, r *http.Request) error {
	return nil
}
