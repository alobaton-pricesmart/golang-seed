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

type PermissionsService struct {
}

func NewPermissionsService() *PermissionsService {
	return &PermissionsService{}
}

func (s PermissionsService) GetByID(id string) (*models.Permission, error) {
	permission := &models.Permission{
		Code: id,
	}
	err := repo.Repo.Permissions().Get(permission)
	if err != nil {
		if errors.Is(err, database.ErrRecordNotFound) {
			return nil, httperror.ErrorCauseT(
				err,
				http.StatusNotFound,
				messagesconst.GeneralErrorRegisterNotFoundParams,
				messagesconst.PermissionsPermissions,
				fmt.Sprintf("code : %s", id))
		}

		return nil, httperror.ErrorCauseT(err, http.StatusInternalServerError, messagesconst.GeneralErrorAccessingDatabase)
	}

	return permission, nil
}

func (s PermissionsService) Get(permission *models.Permission) error {
	err := repo.Repo.Permissions().WhereModel(permission).Get(permission)
	if err != nil {
		if errors.Is(err, database.ErrRecordNotFound) {
			return httperror.ErrorCauseT(
				err,
				http.StatusNotFound,
				messagesconst.GeneralErrorRegisterNotFoundParams,
				messagesconst.PermissionsPermissions,
				permission.String())
		}

		httperror.ErrorCauseT(err, http.StatusInternalServerError, messagesconst.GeneralErrorAccessingDatabase)
	}

	return nil
}

func (s PermissionsService) GetAll(params map[string]interface{}, sort sorting.Sort) ([]models.Permission, error) {
	var permissions []models.Permission
	collection, err := repo.Repo.Permissions().WhereMap(params)
	if err != nil {
		httperror.ErrorCauseT(err, http.StatusBadRequest, err.Error())
	}

	collection, err = collection.Order(sort)
	if err != nil {
		httperror.ErrorCauseT(err, http.StatusBadRequest, err.Error())
	}

	err = collection.Find(&permissions)
	if err != nil {
		if errors.Is(err, database.ErrRecordNotFound) {
			return nil, httperror.ErrorCauseT(
				err,
				http.StatusNotFound,
				messagesconst.GeneralErrorRegisterNotFound,
				messagesconst.PermissionsPermissions)
		}

		return nil, httperror.ErrorCauseT(err, http.StatusInternalServerError, messagesconst.GeneralErrorAccessingDatabase)
	}

	return permissions, nil
}

func (s PermissionsService) GetAllPaged(params map[string]interface{}, sort sorting.Sort, pageable pagination.Pageable) (*pagination.Page, error) {
	var permissions []models.Permission
	collection, err := repo.Repo.Permissions().WhereMap(params)
	if err != nil {
		httperror.ErrorCauseT(err, http.StatusBadRequest, err.Error())
	}

	collectiono, err := collection.Order(sort)
	if err != nil {
		httperror.ErrorCauseT(err, http.StatusBadRequest, err.Error())
	}

	err = collectiono.Pageable(pageable).Find(&permissions)
	if err != nil {
		return nil, httperror.ErrorCauseT(err, http.StatusInternalServerError, messagesconst.GeneralErrorAccessingDatabase)
	}

	var count int64
	err = collection.Count(&count)
	if err != nil {
		return nil, httperror.ErrorCauseT(err, http.StatusInternalServerError, messagesconst.GeneralErrorAccessingDatabase)
	}

	return pagination.NewPage(pageable, int(count), permissions), nil
}

func (s PermissionsService) Create(model *models.Permission) error {
	exists, err := repo.Repo.Permissions().Exists(model)
	if err != nil {
		httperror.ErrorCauseT(err, http.StatusInternalServerError, messagesconst.GeneralErrorAccessingDatabase)
	}

	if exists {
		return httperror.ErrorT(
			http.StatusConflict,
			messagesconst.GeneralErrorRegisterAlreadyExists,
			messagesconst.PermissionsPermission,
			fmt.Sprintf("code : %s", model.Code))
	}

	err = repo.Repo.Permissions().Create(model)
	if err != nil {
		httperror.ErrorCauseT(err, http.StatusInternalServerError, messagesconst.GeneralErrorAccessingDatabase)
	}

	return nil
}

func (s PermissionsService) Update(model *models.Permission) error {
	permission := &models.Permission{Code: model.Code}
	exists, err := repo.Repo.Permissions().Exists(permission)
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

	model.CreatedAt = permission.CreatedAt
	err = repo.Repo.Permissions().WhereModel(&models.Permission{Code: model.Code}).Update(model)
	if err != nil {
		httperror.ErrorCauseT(err, http.StatusInternalServerError, messagesconst.GeneralErrorAccessingDatabase)
	}

	return nil
}

func (s PermissionsService) Delete(id string) error {
	permission := &models.Permission{
		Code: id,
	}
	err := repo.Repo.Permissions().Delete(permission)
	if err != nil {
		if errors.Is(err, database.ErrRecordNotFound) {
			return httperror.ErrorCauseT(
				err,
				http.StatusNotFound,
				messagesconst.GeneralErrorRegisterNotFoundParams,
				messagesconst.PermissionsPermissions,
				fmt.Sprintf("code : %s", id))
		}

		httperror.ErrorCauseT(err, http.StatusInternalServerError, messagesconst.GeneralErrorAccessingDatabase)
	}

	return nil
}
