package service

import (
	"errors"
	"fmt"

	"golang-seed/apps/auth/pkg/messagesconst"
	"golang-seed/apps/auth/pkg/models"
	"golang-seed/pkg/database"
	"golang-seed/pkg/httperror"
	"golang-seed/pkg/pagination"
	"golang-seed/pkg/sorting"
	"net/http"
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
	err := models.Repo.Permissions().Get(permission)
	if err != nil {
		if errors.Is(err, database.ErrRecordNotFound) {
			return nil, httperror.ErrorCauseT(
				err,
				http.StatusNotFound,
				messagesconst.GeneralErrorRegisterNotFoundParams,
				messagesconst.PermissionsPermissions,
				fmt.Sprintf("id : %s", id))
		}

		return nil, httperror.ErrorCauseT(err, http.StatusInternalServerError, messagesconst.GeneralErrorAccessingDatabase)
	}

	return permission, nil
}

func (s PermissionsService) Get(permission *models.Permission) error {
	err := models.Repo.Permissions().Conditions(permission).Get(permission)
	if err != nil {
		if errors.Is(err, database.ErrRecordNotFound) {
			return httperror.ErrorCauseT(
				err,
				http.StatusNotFound,
				messagesconst.GeneralErrorRegisterNotFoundParams,
				messagesconst.PermissionsPermissions,
				permission.String())
		}

		return httperror.ErrorCauseT(err, http.StatusInternalServerError, messagesconst.GeneralErrorAccessingDatabase)
	}

	return nil
}

func (s PermissionsService) GetAll(params map[string]interface{}, sort sorting.Sort) ([]models.Permission, error) {
	var permissions []models.Permission
	err := models.Repo.Permissions().Conditions(params).Order(sort).Find(&permissions)
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
	err := models.Repo.Permissions().Conditions(params).Order(sort).Pageable(pageable).Find(&permissions)
	if err != nil {
		return nil, httperror.ErrorCauseT(err, http.StatusInternalServerError, messagesconst.GeneralErrorAccessingDatabase)
	}

	var count int64
	err = models.Repo.Permissions().Conditions(params).Count(&count)
	if err != nil {
		return nil, httperror.ErrorCauseT(err, http.StatusInternalServerError, messagesconst.GeneralErrorAccessingDatabase)
	}

	return pagination.NewPage(pageable, int(count), permissions), nil
}

func (s PermissionsService) Create(model *models.Permission) error {
	exists, err := models.Repo.Permissions().Exists(model)
	if err != nil {
		return httperror.ErrorCauseT(err, http.StatusInternalServerError, messagesconst.GeneralErrorAccessingDatabase)
	}

	if exists {
		return httperror.ErrorT(
			http.StatusConflict,
			messagesconst.GeneralErrorRegisterAlreadyExists,
			messagesconst.PermissionsPermission,
			fmt.Sprintf("code : %s", model.Code))
	}

	err = models.Repo.Permissions().Create(model)
	if err != nil {
		return httperror.ErrorCauseT(err, http.StatusInternalServerError, messagesconst.GeneralErrorAccessingDatabase)
	}

	return nil
}

func (s PermissionsService) Update(model *models.Permission) error {
	permission := &models.Permission{Code: model.Code}
	exists, err := models.Repo.Permissions().Exists(permission)
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

	model.CreatedAt = permission.CreatedAt
	err = models.Repo.Permissions().Conditions(&models.Permission{Code: model.Code}).Update(model)
	if err != nil {
		return httperror.ErrorCauseT(err, http.StatusInternalServerError, messagesconst.GeneralErrorAccessingDatabase)
	}

	return nil
}

func (s PermissionsService) Delete(id string) error {
	permission := &models.Permission{
		Code: id,
	}
	err := models.Repo.Permissions().Delete(permission)
	if err != nil {
		if errors.Is(err, database.ErrRecordNotFound) {
			return httperror.ErrorCauseT(
				err,
				http.StatusNotFound,
				messagesconst.GeneralErrorRegisterNotFoundParams,
				messagesconst.PermissionsPermissions,
				fmt.Sprintf("id : %s", id))
		}

		return httperror.ErrorCauseT(err, http.StatusInternalServerError, messagesconst.GeneralErrorAccessingDatabase)
	}

	return nil
}
