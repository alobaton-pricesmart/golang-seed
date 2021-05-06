package service

import (
	"errors"
	"fmt"
	"net/http"

	"golang-seed/apps/auth/pkg/messagesconst"
	"golang-seed/apps/auth/pkg/models"
	"golang-seed/apps/auth/pkg/repo"
	"golang-seed/pkg/database"
	"golang-seed/pkg/httperror"
	"golang-seed/pkg/pagination"
	"golang-seed/pkg/sorting"
)

type ClientsService struct {
}

func NewClientsService() *ClientsService {
	return &ClientsService{}
}

func (s ClientsService) GetByID(id string) (*models.Client, error) {
	client := &models.Client{
		Code: id,
	}
	err := repo.Repo.Clients().Get(client)
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

func (s ClientsService) Get(model *models.Client) error {
	err := repo.Repo.Clients().WhereModel(model).Get(model)
	if err != nil {
		if errors.Is(err, database.ErrRecordNotFound) {
			return httperror.ErrorCauseT(
				err,
				http.StatusNotFound,
				messagesconst.GeneralErrorRegisterNotFoundParams,
				messagesconst.ClientsClients,
				model.String())
		}

		httperror.ErrorCauseT(err, http.StatusInternalServerError, messagesconst.GeneralErrorAccessingDatabase)
	}

	return nil
}

func (s ClientsService) GetAll(params map[string]interface{}, sort sorting.Sort) ([]*models.Client, error) {
	var clients []*models.Client
	collection, err := repo.Repo.Clients().WhereMap(params)
	if err != nil {
		httperror.ErrorCauseT(err, http.StatusBadRequest, err.Error())
	}

	collection, err = collection.Order(sort)
	if err != nil {
		httperror.ErrorCauseT(err, http.StatusBadRequest, err.Error())
	}

	collection.Find(&clients)
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

func (s ClientsService) GetAllPaged(params map[string]interface{}, sort sorting.Sort, pageable pagination.Pageable) (*pagination.Page, error) {
	var clients []*models.Client
	collection, err := repo.Repo.Clients().WhereMap(params)
	if err != nil {
		httperror.ErrorCauseT(err, http.StatusBadRequest, err.Error())
	}

	collectiono, err := collection.Order(sort)
	if err != nil {
		httperror.ErrorCauseT(err, http.StatusBadRequest, err.Error())
	}

	err = collectiono.Pageable(pageable).Find(&clients)
	if err != nil {
		return nil, httperror.ErrorCauseT(err, http.StatusInternalServerError, messagesconst.GeneralErrorAccessingDatabase)
	}

	var count int64
	err = collection.Count(&count)
	if err != nil {
		return nil, httperror.ErrorCauseT(err, http.StatusInternalServerError, messagesconst.GeneralErrorAccessingDatabase)
	}

	return pagination.NewPage(pageable, int(count), clients), nil
}

func (s ClientsService) Create(model *models.Client) error {
	exists, err := repo.Repo.Clients().Exists(model)
	if err != nil {
		httperror.ErrorCauseT(err, http.StatusInternalServerError, messagesconst.GeneralErrorAccessingDatabase)
	}

	if exists {
		return httperror.ErrorT(
			http.StatusConflict,
			messagesconst.GeneralErrorRegisterAlreadyExists,
			messagesconst.ClientsClient,
			fmt.Sprintf("code : %s", model.Code))
	}

	err = repo.Repo.Clients().Create(model)
	if err != nil {
		httperror.ErrorCauseT(err, http.StatusInternalServerError, messagesconst.GeneralErrorAccessingDatabase)
	}

	return nil
}

func (s ClientsService) Update(model *models.Client) error {
	client := &models.Client{Code: model.Code}
	exists, err := repo.Repo.Clients().Exists(client)
	if err != nil {
		httperror.ErrorCauseT(err, http.StatusInternalServerError, messagesconst.GeneralErrorAccessingDatabase)
	}

	if !exists {
		return httperror.ErrorT(
			http.StatusNotFound,
			messagesconst.GeneralErrorRegisterNotFoundParams,
			messagesconst.ClientsClients,
			fmt.Sprintf("code : %s", model.Code))
	}

	model.CreatedAt = client.CreatedAt
	err = repo.Repo.Clients().WhereModel(&models.Client{Code: model.Code}).Update(model)
	if err != nil {
		httperror.ErrorCauseT(err, http.StatusInternalServerError, messagesconst.GeneralErrorAccessingDatabase)
	}

	return nil
}

func (s ClientsService) Delete(id string) error {
	model := &models.Client{
		Code: id,
	}
	err := repo.Repo.Clients().Delete(model)
	if err != nil {
		if errors.Is(err, database.ErrRecordNotFound) {
			return httperror.ErrorCauseT(
				err,
				http.StatusNotFound,
				messagesconst.GeneralErrorRegisterNotFoundParams,
				messagesconst.ClientsClients,
				fmt.Sprintf("code : %s", id))
		}

		httperror.ErrorCauseT(err, http.StatusInternalServerError, messagesconst.GeneralErrorAccessingDatabase)
	}

	return nil
}
