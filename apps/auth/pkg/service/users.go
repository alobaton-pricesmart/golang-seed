package service

import (
	"errors"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"golang-seed/apps/auth/pkg/authconst"
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
				authconst.GeneralErrorRegisterNotFoundParams,
				authconst.UsersUsers,
				fmt.Sprintf("id : %s", id))
		}

		return nil, httperror.ErrorCauseT(err, http.StatusInternalServerError, authconst.GeneralErrorAccessingDatabase)
	}

	return user, nil
}

func (s UsersService) Get(user *models.User) error {
	err := repo.Repo.Users().WhereModel(user).Get(user)
	if err != nil {
		if errors.Is(err, database.ErrRecordNotFound) {
			return httperror.ErrorCauseT(
				err,
				http.StatusNotFound,
				authconst.GeneralErrorRegisterNotFoundParams,
				authconst.UsersUsers,
				user.String())
		}

		httperror.ErrorCauseT(err, http.StatusInternalServerError, authconst.GeneralErrorAccessingDatabase)
	}

	return nil
}

func (s UsersService) GetAll(params map[string]interface{}, sort sorting.Sort) ([]models.User, error) {
	var users []models.User
	collection, err := repo.Repo.Users().WhereMap(params)
	if err != nil {
		httperror.ErrorCauseT(err, http.StatusBadRequest, err.Error())
	}

	collection, err = collection.Order(sort)
	if err != nil {
		httperror.ErrorCauseT(err, http.StatusBadRequest, err.Error())
	}

	err = collection.Find(&users)
	if err != nil {
		if errors.Is(err, database.ErrRecordNotFound) {
			return nil, httperror.ErrorCauseT(
				err,
				http.StatusNotFound,
				authconst.GeneralErrorRegisterNotFound,
				authconst.UsersUsers)
		}

		return nil, httperror.ErrorCauseT(err, http.StatusInternalServerError, authconst.GeneralErrorAccessingDatabase)
	}

	return users, nil
}

func (s UsersService) GetAllPaged(params map[string]interface{}, sort sorting.Sort, pageable pagination.Pageable) (*pagination.Page, error) {
	var users []models.User
	collection, err := repo.Repo.Users().WhereMap(params)
	if err != nil {
		httperror.ErrorCauseT(err, http.StatusBadRequest, err.Error())
	}

	collectiono, err := collection.Order(sort)
	if err != nil {
		httperror.ErrorCauseT(err, http.StatusBadRequest, err.Error())
	}

	err = collectiono.Pageable(pageable).Find(&users)
	if err != nil {
		return nil, httperror.ErrorCauseT(err, http.StatusInternalServerError, authconst.GeneralErrorAccessingDatabase)
	}

	var count int64
	err = collection.Count(&count)
	if err != nil {
		return nil, httperror.ErrorCauseT(err, http.StatusInternalServerError, authconst.GeneralErrorAccessingDatabase)
	}

	return pagination.NewPage(pageable, int(count), users), nil
}

func (s UsersService) Create(model *models.User) error {
	exists, err := repo.Repo.Users().Exists(model)
	if err != nil {
		httperror.ErrorCauseT(err, http.StatusInternalServerError, authconst.GeneralErrorAccessingDatabase)
	}

	if exists {
		return httperror.ErrorT(
			http.StatusConflict,
			authconst.GeneralErrorRegisterAlreadyExists,
			authconst.UsersUser,
			fmt.Sprintf("id : %s", model.ID))
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(model.Password), bcrypt.DefaultCost)
	if err != nil {
		return httperror.ErrorCauseT(err, http.StatusInternalServerError, authconst.OAuthInvalidUsernamePassword)
	}
	model.Password = string(hash)

	err = repo.Repo.Users().Create(model)
	if err != nil {
		httperror.ErrorCauseT(err, http.StatusInternalServerError, authconst.GeneralErrorAccessingDatabase)
	}

	return nil
}

func (s UsersService) Update(model *models.User) error {
	user := &models.User{ID: model.ID}
	exists, err := repo.Repo.Users().Exists(user)
	if err != nil {
		httperror.ErrorCauseT(err, http.StatusInternalServerError, authconst.GeneralErrorAccessingDatabase)
	}

	if !exists {
		return httperror.ErrorT(
			http.StatusNotFound,
			authconst.GeneralErrorRegisterNotFoundParams,
			authconst.ClientsClients,
			fmt.Sprintf("id : %s", model.ID))
	}

	model.CreatedAt = user.CreatedAt
	err = repo.Repo.Users().WhereModel(&models.User{ID: model.ID}).Update(model)
	if err != nil {
		httperror.ErrorCauseT(err, http.StatusInternalServerError, authconst.GeneralErrorAccessingDatabase)
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
				authconst.GeneralErrorRegisterNotFoundParams,
				authconst.UsersUsers,
				fmt.Sprintf("id : %s", id))
		}

		httperror.ErrorCauseT(err, http.StatusInternalServerError, authconst.GeneralErrorAccessingDatabase)
	}

	return nil
}
