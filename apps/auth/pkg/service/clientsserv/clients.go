package clientsserv

import (
	"errors"
	"fmt"
	"golang-seed/apps/auth/pkg/messagesconst"
	"golang-seed/apps/auth/pkg/models"
	"golang-seed/pkg/database"
	"golang-seed/pkg/httperror"
	"net/http"
)

type ClientsService struct {
}

func NewClientsService() *ClientsService {
	return &ClientsService{}
}

func (s *ClientsService) Get(id string) (*models.Client, error) {
	client := &models.Client{
		Code: id,
	}
	err := models.Repo.Clients().Get(client)
	if err != nil {
		if errors.Is(err, database.ErrRecordNotFound) {
			return nil, httperror.NewHTTPErrorT(
				err,
				http.StatusNotFound,
				messagesconst.GeneralErrorRegisterNotFoundParams,
				messagesconst.ClientsClients,
				fmt.Sprintf("id : %s", id))
		}

		return nil, httperror.NewHTTPErrorT(err, http.StatusInternalServerError, messagesconst.GeneralErrorAccessingDatabase)
	}

	return client, nil
}

func (s *ClientsService) GetAll(params map[string]interface{}, sort database.Sort) ([]models.Client, error) {
	var clients []models.Client
	err := models.Repo.Clients().Conditions(params).Order(sort).Find(&clients)
	if err != nil {
		if errors.Is(err, database.ErrRecordNotFound) {
			return nil, httperror.NewHTTPErrorT(
				err,
				http.StatusNotFound,
				messagesconst.GeneralErrorRegisterNotFound,
				messagesconst.ClientsClients)
		}

		return nil, httperror.NewHTTPErrorT(err, http.StatusInternalServerError, messagesconst.GeneralErrorAccessingDatabase)
	}

	return clients, nil
}
