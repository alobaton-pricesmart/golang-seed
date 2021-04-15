package clientsserv

import "golang-seed/apps/auth/pkg/models"

type ClientsService struct {
}

func NewClientsService() *ClientsService {
	return &ClientsService{}
}

func (s *ClientsService) Get(id string) (models.Client, error) {
	client := &models.Client{
		Code: id,
	}
	err := models.Repo.Clients().GetByID(client)

	return *client, err
}
