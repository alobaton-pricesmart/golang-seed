package clientshand

import (
	"encoding/json"
	"net/http"
	"strconv"

	"golang-seed/apps/auth/pkg/messagesconst"
	"golang-seed/apps/auth/pkg/service/clientsserv"
	"golang-seed/pkg/database"
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
		return err
	}

	body, err := json.Marshal(client)
	if err != nil {
		return httperror.ErrorCauseT(err, http.StatusInternalServerError, messagesconst.GeneralErrorMarshal)
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
	delete(params, "sort")

	clients, err := u.clientsService.GetAll(params, sort)
	if err != nil {
		return err
	}

	body, err := json.Marshal(clients)
	if err != nil {
		return httperror.ErrorCauseT(err, http.StatusInternalServerError, messagesconst.GeneralErrorMarshal)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
	return nil
}

func (u ClientsHandler) GetAllPaged(w http.ResponseWriter, r *http.Request) error {
	if len(r.URL.Query()["page"]) < 1 {
		return httperror.ErrorT(http.StatusBadRequest, messagesconst.GeneralErrorRequiredField, "page")
	}

	if len(r.URL.Query()["size"]) < 1 {
		return httperror.ErrorT(http.StatusBadRequest, messagesconst.GeneralErrorRequiredField, "size")
	}

	var err error
	var pagep int
	var sizep int

	pagep, err = strconv.Atoi(r.URL.Query()["page"][0])
	if err != nil {
		return httperror.ErrorCauseT(err, http.StatusInternalServerError, messagesconst.GeneralErrorMarshal)
	}

	sizep, err = strconv.Atoi(r.URL.Query()["size"][0])
	if err != nil {
		return httperror.ErrorCauseT(err, http.StatusInternalServerError, messagesconst.GeneralErrorMarshal)
	}

	pageable := database.NewPageable(pagep, sizep)

	sortp := r.URL.Query()["sort"]
	sort := database.NewSort(sortp)

	params := make(map[string]interface{})
	for k, v := range r.URL.Query() {
		params[k] = v
	}
	delete(params, "sort")
	delete(params, "page")
	delete(params, "size")

	clients, err := u.clientsService.GetAllPaged(params, sort, pageable)
	if err != nil {
		return err
	}

	body, err := json.Marshal(clients)
	if err != nil {
		return httperror.ErrorCauseT(err, http.StatusInternalServerError, messagesconst.GeneralErrorMarshal)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
	return nil
}

func (u ClientsHandler) Create(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (u ClientsHandler) Update(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (u ClientsHandler) Delete(w http.ResponseWriter, r *http.Request) error {
	pathVars := mux.Vars(r)
	id := pathVars["id"]

	err := u.clientsService.Delete(id)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	return nil
}
