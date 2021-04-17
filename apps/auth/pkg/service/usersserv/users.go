package usersserv

import (
	"errors"
	"fmt"
	"golang-seed/apps/auth/pkg/messagesconst"
	"golang-seed/apps/auth/pkg/models"
	"golang-seed/pkg/database"
	"golang-seed/pkg/httperror"
	"net/http"

	"github.com/google/uuid"
)

type UsersService struct {
}

func NewUsersService() *UsersService {
	return &UsersService{}
}

func (s *UsersService) Get(id uuid.UUID) (*models.User, error) {
	user := &models.User{
		ID: id,
	}
	err := models.Repo.Users().Get(user)
	if err != nil {
		if errors.Is(err, database.ErrRecordNotFound) {
			return nil, httperror.NewHTTPErrorT(
				err,
				http.StatusNotFound,
				messagesconst.GeneralErrorRegisterNotFoundParams,
				messagesconst.UsersUsers,
				fmt.Sprintf("id %v", id))
		}

		return nil, httperror.NewHTTPErrorT(err, http.StatusInternalServerError, messagesconst.GeneralErrorAccessingDatabase)
	}

	return user, nil
}
