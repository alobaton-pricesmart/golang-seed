package store

import (
	"github.com/juju/errors"
	"gopkg.in/oauth2.v3"
	"gopkg.in/oauth2.v3/store"

	"github.com/alobaton/golang-seed/apps/auth/pkg/models"
)

// NewClientStore create client store
func NewClientStore() *ClientStore {
	return new(ClientStore)
}

// ClientStore client information store
type ClientStore store.ClientStore

// GetByID according to the ID for the client information
func (cs *ClientStore) GetByID(id string) (cli oauth2.ClientInfo, err error) {
	client := &models.Client{
		Code: id,
	}
	if err := models.Repo.Clients().Get(client); err != nil {
		return nil, errors.Trace(err)
	}

	return client, nil
}

// Set set client information
func (cs *ClientStore) Set(id string, cli oauth2.ClientInfo) (err error) {
	if err := models.Repo.Markets().Put(cli); err != nil {
		return errors.Trace(err)
	}

	return nil
}
