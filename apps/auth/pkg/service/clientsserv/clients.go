package clientsserv

import (
	"golang-seed/apps/auth/pkg/models"
	"golang-seed/pkg/database"
)

type ClientsService struct {
}

func NewClientsService() *ClientsService {
	return &ClientsService{}
}

func (s *ClientsService) Get(id string) (models.Client, error) {
	client := &models.Client{
		Code: id,
	}
	err := models.Repo.Clients().Get(client)

	return *client, err
}

func (s *ClientsService) GetAll(params map[string]interface{}, sort database.Sort) ([]models.Client, error) {
	var clients []models.Client
	err := models.Repo.Clients().Conditions(params).Find(&clients)

	return clients, err
}
