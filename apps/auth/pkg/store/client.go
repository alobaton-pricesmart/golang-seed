package store

import (
	"context"

	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/store"

	"golang-seed/apps/auth/pkg/models"
)

// NewClientStore creates client store
func NewClientStore() *ClientStore {
	return new(ClientStore)
}

// ClientStore client information store
type ClientStore store.ClientStore

func serializeClientInfo(id string, cli oauth2.ClientInfo) *models.Client {
	client := &models.Client{
		Code:   id,
		Secret: cli.GetSecret(),
		Domain: cli.GetDomain(),
	}

	return client
}

// GetByID according to the ID for the client information
func (cs *ClientStore) GetByID(ctx context.Context, id string) (oauth2.ClientInfo, error) {
	cs.RLock()
	defer cs.RUnlock()

	client := &models.Client{
		Code: id,
	}
	err := models.Repo.Clients().Get(client)

	return client, err
}

// Set set client information
func (cs *ClientStore) Set(id string, cli oauth2.ClientInfo) error {
	cs.RLock()
	defer cs.RUnlock()

	client := serializeClientInfo(id, cli)

	return models.Repo.Clients().Update(client)
}
