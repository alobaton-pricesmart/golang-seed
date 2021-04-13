package store

import (
	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/store"
)

// NewClientStore create client store
func NewTokenStore() (oauth2.TokenStore, error) {
	return store.NewMemoryTokenStore()
}
