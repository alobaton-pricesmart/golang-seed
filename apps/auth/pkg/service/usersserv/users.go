package usersserv

import (
	"golang-seed/apps/auth/pkg/models"

	"github.com/google/uuid"
)

type UsersService struct {
}

func NewUsersService() *UsersService {
	return &UsersService{}
}

func (s *UsersService) Get(id uuid.UUID) (models.User, error) {
	user := &models.User{
		ID: id,
	}
	err := models.Repo.Users().Get(user)

	return *user, err
}
