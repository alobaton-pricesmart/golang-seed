package clientshand

import (
	"encoding/json"
	"net/http"

	"golang-seed/apps/auth/pkg/messagesconst"
	"golang-seed/apps/auth/pkg/service/clientsserv"
	"golang-seed/pkg/database"
	"golang-seed/pkg/httperror"
	"golang-seed/pkg/messages"

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
		return httperror.NewHTTPError(err, http.StatusInternalServerError, messages.Get(messagesconst.GeneralErrorGetting, messages.Get(messagesconst.ClientsClient)))
	}

	body, err := json.Marshal(client)
	if err != nil {
		return httperror.NewHTTPError(err, http.StatusInternalServerError, messages.Get(messagesconst.GeneralErrorMarshal))
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
	return nil
}

func (u ClientsHandler) GetAll(w http.ResponseWriter, r *http.Request) error {
	sparams := r.URL.Query()["sort"]
	sort := database.NewSort(sparams)

	params := make(map[string]interface{})
	for k, v := range r.URL.Query() {
		params[k] = v
	}

	clients, err := u.clientsService.GetAll(params, sort)
	if err != nil {
		return httperror.NewHTTPError(err, http.StatusInternalServerError, messages.Get(messagesconst.GeneralErrorGetting, messages.Get(messagesconst.ClientsClients)))
	}

	body, err := json.Marshal(clients)
	if err != nil {
		return httperror.NewHTTPError(err, http.StatusInternalServerError, messages.Get(messagesconst.GeneralErrorMarshal))
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
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
