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

func (s *ClientsService) GetByID(id string) (*models.Client, error) {
	client := &models.Client{
		Code: id,
	}
	err := models.Repo.Clients().Get(client)
	if err != nil {
		if errors.Is(err, database.ErrRecordNotFound) {
			return nil, httperror.ErrorCauseT(
				err,
				http.StatusNotFound,
				messagesconst.GeneralErrorRegisterNotFoundParams,
				messagesconst.ClientsClients,
				fmt.Sprintf("code : %s", id))
		}

		return nil, httperror.ErrorCauseT(err, http.StatusInternalServerError, messagesconst.GeneralErrorAccessingDatabase)
	}

	return client, nil
}

func (s *ClientsService) Get(model *models.Client) error {
	err := models.Repo.Clients().Conditions(model).Get(model)
	if err != nil {
		if errors.Is(err, database.ErrRecordNotFound) {
			return httperror.ErrorCauseT(
				err,
				http.StatusNotFound,
				messagesconst.GeneralErrorRegisterNotFoundParams,
				messagesconst.ClientsClients,
				model.String())
		}

		return httperror.ErrorCauseT(err, http.StatusInternalServerError, messagesconst.GeneralErrorAccessingDatabase)
	}

	return nil
}

func (s *ClientsService) GetAll(params map[string]interface{}, sort database.Sort) ([]*models.Client, error) {
	var clients []*models.Client
	err := models.Repo.Clients().Conditions(params).Order(sort).Find(&clients)
	if err != nil {
		if errors.Is(err, database.ErrRecordNotFound) {
			return nil, httperror.ErrorCauseT(
				err,
				http.StatusNotFound,
				messagesconst.GeneralErrorRegisterNotFound,
				messagesconst.ClientsClients)
		}

		return nil, httperror.ErrorCauseT(err, http.StatusInternalServerError, messagesconst.GeneralErrorAccessingDatabase)
	}

	return clients, nil
}

func (s *ClientsService) GetAllPaged(params map[string]interface{}, sort database.Sort, pageable database.Pageable) (*database.Page, error) {
	var clients []*models.Client
	err := models.Repo.Clients().Conditions(params).Order(sort).Pageable(pageable).Find(&clients)
	if err != nil {
		return nil, httperror.ErrorCauseT(err, http.StatusInternalServerError, messagesconst.GeneralErrorAccessingDatabase)
	}

	var count int64
	err = models.Repo.Clients().Conditions(params).Count(&count)
	if err != nil {
		return nil, httperror.ErrorCauseT(err, http.StatusInternalServerError, messagesconst.GeneralErrorAccessingDatabase)
	}

	return database.NewPage(pageable, int(count), clients), nil
}

func (s *ClientsService) Create(model *models.Client) error {
	exists, err := models.Repo.Clients().Exists(model)
	if err != nil {
		return httperror.ErrorCauseT(err, http.StatusInternalServerError, messagesconst.GeneralErrorAccessingDatabase)
	}

	if exists {
		return httperror.ErrorT(
			http.StatusConflict,
			messagesconst.GeneralErrorRegisterAlreadyExists,
			messagesconst.ClientsClient,
			fmt.Sprintf("code : %s", model.Code))
	}

	err = models.Repo.Clients().Create(model)
	if err != nil {
		return httperror.ErrorCauseT(err, http.StatusInternalServerError, messagesconst.GeneralErrorAccessingDatabase)
	}

	return nil
}

func (s *ClientsService) Update(model *models.Client) error {
	client := &models.Client{Code: model.Code}
	exists, err := models.Repo.Clients().Exists(client)
	if err != nil {
		return httperror.ErrorCauseT(err, http.StatusInternalServerError, messagesconst.GeneralErrorAccessingDatabase)
	}

	if !exists {
		return httperror.ErrorT(
			http.StatusNotFound,
			messagesconst.GeneralErrorRegisterNotFoundParams,
			messagesconst.ClientsClients,
			fmt.Sprintf("code : %s", model.Code))
	}

	model.CreatedAt = client.CreatedAt
	err = models.Repo.Clients().Conditions(&models.Client{Code: model.Code}).Update(model)
	if err != nil {
		return httperror.ErrorCauseT(err, http.StatusInternalServerError, messagesconst.GeneralErrorAccessingDatabase)
	}

	return nil
}

func (s *ClientsService) Delete(id string) error {
	model := &models.Client{
		Code: id,
	}
	err := models.Repo.Clients().Delete(model)
	if err != nil {
		if errors.Is(err, database.ErrRecordNotFound) {
			return httperror.ErrorCauseT(
				err,
				http.StatusNotFound,
				messagesconst.GeneralErrorRegisterNotFoundParams,
				messagesconst.ClientsClients,
				fmt.Sprintf("code : %s", id))
		}

		return httperror.ErrorCauseT(err, http.StatusInternalServerError, messagesconst.GeneralErrorAccessingDatabase)
	}

	return nil
}
