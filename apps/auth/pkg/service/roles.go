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

type RolesService struct {
}

func NewRolesService() *RolesService {
	return &RolesService{}
}

func (r RolesService) GetByID(id string) (*models.Role, error) {
	role := &models.Role{
		Code: id,
	}
	err := repo.Repo.Roles().Get(role)
	if err != nil {
		if errors.Is(err, database.ErrRecordNotFound) {
			return nil, httperror.ErrorCauseT(
				err,
				http.StatusNotFound,
				messagesconst.GeneralErrorRegisterNotFoundParams,
				messagesconst.RolesRoles,
				fmt.Sprintf("code : %s", id))
		}

		return nil, httperror.ErrorCauseT(err, http.StatusInternalServerError, messagesconst.GeneralErrorAccessingDatabase)
	}

	return role, nil
}

func (r RolesService) Get(role *models.Role) error {
	err := repo.Repo.Roles().Conditions(role).Get(role)
	if err != nil {
		if errors.Is(err, database.ErrRecordNotFound) {
			return httperror.ErrorCauseT(
				err,
				http.StatusNotFound,
				messagesconst.GeneralErrorRegisterNotFoundParams,
				messagesconst.RolesRoles,
				role.String())
		}

		return httperror.ErrorCauseT(err, http.StatusInternalServerError, messagesconst.GeneralErrorAccessingDatabase)
	}

	return nil
}

func (r RolesService) GetAll(params map[string]interface{}, sort sorting.Sort) ([]models.Role, error) {
	var roles []models.Role
	err := repo.Repo.Roles().Conditions(params).Order(sort).Find(&roles)
	if err != nil {
		if errors.Is(err, database.ErrRecordNotFound) {
			return nil, httperror.ErrorCauseT(
				err,
				http.StatusNotFound,
				messagesconst.GeneralErrorRegisterNotFound,
				messagesconst.RolesRoles)
		}

		return nil, httperror.ErrorCauseT(err, http.StatusInternalServerError, messagesconst.GeneralErrorAccessingDatabase)
	}

	return roles, nil
}

func (r RolesService) GetAllPaged(params map[string]interface{}, sort sorting.Sort, pageable pagination.Pageable) (*pagination.Page, error) {
	var roles []models.Role
	err := repo.Repo.Roles().Conditions(params).Order(sort).Pageable(pageable).Find(&roles)
	if err != nil {
		return nil, httperror.ErrorCauseT(err, http.StatusInternalServerError, messagesconst.GeneralErrorAccessingDatabase)
	}

	var count int64
	err = repo.Repo.Roles().Conditions(params).Count(&count)
	if err != nil {
		return nil, httperror.ErrorCauseT(err, http.StatusInternalServerError, messagesconst.GeneralErrorAccessingDatabase)
	}

	return pagination.NewPage(pageable, int(count), roles), nil
}

func (r RolesService) Create(model *models.Role) error {
	exists, err := repo.Repo.Roles().Exists(model)
	if err != nil {
		return httperror.ErrorCauseT(err, http.StatusInternalServerError, messagesconst.GeneralErrorAccessingDatabase)
	}

	if exists {
		return httperror.ErrorT(
			http.StatusConflict,
			messagesconst.GeneralErrorRegisterAlreadyExists,
			messagesconst.RolesRole,
			fmt.Sprintf("code : %s", model.Code))
	}

	err = repo.Repo.Roles().Create(model)
	if err != nil {
		return httperror.ErrorCauseT(err, http.StatusInternalServerError, messagesconst.GeneralErrorAccessingDatabase)
	}

	return nil
}

func (r RolesService) Update(model *models.Role) error {
	role := &models.Role{Code: model.Code}
	exists, err := repo.Repo.Roles().Exists(role)
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

	model.CreatedAt = role.CreatedAt
	err = repo.Repo.Roles().Conditions(&models.Role{Code: model.Code}).Update(model)
	if err != nil {
		return httperror.ErrorCauseT(err, http.StatusInternalServerError, messagesconst.GeneralErrorAccessingDatabase)
	}

	return nil
}

func (r RolesService) Delete(id string) error {
	role := &models.Role{
		Code: id,
	}
	err := repo.Repo.Roles().Delete(role)
	if err != nil {
		if errors.Is(err, database.ErrRecordNotFound) {
			return httperror.ErrorCauseT(
				err,
				http.StatusNotFound,
				messagesconst.GeneralErrorRegisterNotFoundParams,
				messagesconst.RolesRoles,
				fmt.Sprintf("code : %s", id))
		}

		return httperror.ErrorCauseT(err, http.StatusInternalServerError, messagesconst.GeneralErrorAccessingDatabase)
	}

	return nil
}
