package usersserv

import (
	"errors"
	"fmt"
	"golang-seed/apps/auth/pkg/messagesconst"
	"golang-seed/apps/auth/pkg/models"
	"golang-seed/pkg/database"
	"golang-seed/pkg/httperror"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type UsersService struct {
}

func NewUsersService() *UsersService {
	return &UsersService{}
}

func (s *UsersService) GetByID(id string) (*models.User, error) {
	user := &models.User{
		ID: id,
	}
	err := models.Repo.Users().Get(user)
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

func (s *UsersService) Get(user *models.User) error {
	err := models.Repo.Users().Conditions(user).Get(user)
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

func (s *UsersService) GetAll(params map[string]interface{}, sort database.Sort) ([]models.User, error) {
	var users []models.User
	err := models.Repo.Users().Conditions(params).Order(sort).Find(&users)
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

func (s *UsersService) GetAllPaged(params map[string]interface{}, sort database.Sort, pageable database.Pageable) (*database.Page, error) {
	var users []models.User
	err := models.Repo.Users().Conditions(params).Order(sort).Pageable(pageable).Find(&users)
	if err != nil {
		return nil, httperror.ErrorCauseT(err, http.StatusInternalServerError, messagesconst.GeneralErrorAccessingDatabase)
	}

	var count int64
	err = models.Repo.Users().Conditions(params).Count(&count)
	if err != nil {
		return nil, httperror.ErrorCauseT(err, http.StatusInternalServerError, messagesconst.GeneralErrorAccessingDatabase)
	}

	return database.NewPage(pageable, int(count), users), nil
}

func (s *UsersService) Create(model *models.User) error {
	exists, err := models.Repo.Users().Exists(model)
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

	err = models.Repo.Users().Create(model)
	if err != nil {
		return httperror.ErrorCauseT(err, http.StatusInternalServerError, messagesconst.GeneralErrorAccessingDatabase)
	}

	return nil
}

func (s *UsersService) Delete(id string) error {
	user := &models.User{
		ID: id,
	}
	err := models.Repo.Users().Delete(user)
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
