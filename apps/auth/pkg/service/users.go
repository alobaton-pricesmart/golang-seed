package service

import (
	"errors"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"golang-seed/apps/auth/pkg/messagesconst"
	"golang-seed/apps/auth/pkg/models"
	"golang-seed/apps/auth/pkg/repo"
	"golang-seed/pkg/database"
	"golang-seed/pkg/httperror"
	"golang-seed/pkg/pagination"
	"golang-seed/pkg/sorting"
)

type UsersService struct {
}

func NewUsersService() *UsersService {
	return &UsersService{}
}

func (s UsersService) GetByID(id string) (*models.User, error) {
	user := &models.User{
		ID: id,
	}
	err := repo.Repo.Users().Get(user)
	if err != nil {
		if errors.Is(err, database.ErrRecordNotFound) {
			return nil, httperror.ErrorCauseT(
				err,
				http.StatusNotFound,
				messagesconst.GeneralErrorRegisterNotFoundParams,
				messagesconst.UsersUsers,
				fmt.Sprintf("id : %s", id))
		}

		return nil, httperror.ErrorCauseT(err, http.StatusInternalServerError, messagesconst.GeneralErrorAccessingDatabase)
	}

	return user, nil
}

func (s UsersService) Get(user *models.User) error {
	err := repo.Repo.Users().Conditions(user).Get(user)
	if err != nil {
		if errors.Is(err, database.ErrRecordNotFound) {
			return httperror.ErrorCauseT(
				err,
				http.StatusNotFound,
				messagesconst.GeneralErrorRegisterNotFoundParams,
				messagesconst.UsersUsers,
				user.String())
		}

		return httperror.ErrorCauseT(err, http.StatusInternalServerError, messagesconst.GeneralErrorAccessingDatabase)
	}

	return nil
}

func (s UsersService) GetAll(params map[string]interface{}, sort sorting.Sort) ([]models.User, error) {
	var users []models.User
	err := repo.Repo.Users().Conditions(params).Order(sort).Find(&users)
	if err != nil {
		if errors.Is(err, database.ErrRecordNotFound) {
			return nil, httperror.ErrorCauseT(
				err,
				http.StatusNotFound,
				messagesconst.GeneralErrorRegisterNotFound,
				messagesconst.UsersUsers)
		}

		return nil, httperror.ErrorCauseT(err, http.StatusInternalServerError, messagesconst.GeneralErrorAccessingDatabase)
	}

	return users, nil
}

func (s UsersService) GetAllPaged(params map[string]interface{}, sort sorting.Sort, pageable pagination.Pageable) (*pagination.Page, error) {
	var users []models.User
	err := repo.Repo.Users().Conditions(params).Order(sort).Pageable(pageable).Find(&users)
	if err != nil {
		return nil, httperror.ErrorCauseT(err, http.StatusInternalServerError, messagesconst.GeneralErrorAccessingDatabase)
	}

	var count int64
	err = repo.Repo.Users().Conditions(params).Count(&count)
	if err != nil {
		return nil, httperror.ErrorCauseT(err, http.StatusInternalServerError, messagesconst.GeneralErrorAccessingDatabase)
	}

	return pagination.NewPage(pageable, int(count), users), nil
}

func (s UsersService) Create(model *models.User) error {
	exists, err := repo.Repo.Users().Exists(model)
	if err != nil {
		return httperror.ErrorCauseT(err, http.StatusInternalServerError, messagesconst.GeneralErrorAccessingDatabase)
	}

	if exists {
		return httperror.ErrorT(
			http.StatusConflict,
			messagesconst.GeneralErrorRegisterAlreadyExists,
			messagesconst.UsersUser,
			fmt.Sprintf("id : %s", model.ID))
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(model.Password), bcrypt.DefaultCost)
	if err != nil {
		return httperror.ErrorCauseT(err, http.StatusInternalServerError, messagesconst.OAuthInvalidUsernamePassword)
	}
	model.Password = string(hash)

	err = repo.Repo.Users().Create(model)
	if err != nil {
		return httperror.ErrorCauseT(err, http.StatusInternalServerError, messagesconst.GeneralErrorAccessingDatabase)
	}

	return nil
}

func (s UsersService) Update(model *models.User) error {
	user := &models.User{ID: model.ID}
	exists, err := repo.Repo.Users().Exists(user)
	if err != nil {
		return httperror.ErrorCauseT(err, http.StatusInternalServerError, messagesconst.GeneralErrorAccessingDatabase)
	}

	if !exists {
		return httperror.ErrorT(
			http.StatusNotFound,
			messagesconst.GeneralErrorRegisterNotFoundParams,
			messagesconst.ClientsClients,
			fmt.Sprintf("id : %s", model.ID))
	}

	model.CreatedAt = user.CreatedAt
	err = repo.Repo.Users().Conditions(&models.User{ID: model.ID}).Update(model)
	if err != nil {
		return httperror.ErrorCauseT(err, http.StatusInternalServerError, messagesconst.GeneralErrorAccessingDatabase)
	}

	return nil
}

func (s UsersService) Delete(id string) error {
	user := &models.User{
		ID: id,
	}
	err := repo.Repo.Users().Delete(user)
	if err != nil {
		if errors.Is(err, database.ErrRecordNotFound) {
			return httperror.ErrorCauseT(
				err,
				http.StatusNotFound,
				messagesconst.GeneralErrorRegisterNotFoundParams,
				messagesconst.UsersUsers,
				fmt.Sprintf("id : %s", id))
		}

		return httperror.ErrorCauseT(err, http.StatusInternalServerError, messagesconst.GeneralErrorAccessingDatabase)
	}

	return nil
}
