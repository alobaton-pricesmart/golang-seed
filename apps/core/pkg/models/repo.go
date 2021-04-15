package models

import (
	"github.com/juju/errors"

	"golang-seed/apps/core/pkg/config"
	"golang-seed/pkg/database"
	"golang-seed/pkg/server"
)

var Repo *Repository

func ConnectRepo() error {
	credentials := database.Credentials{
		User:      config.Settings.Database.User,
		Password:  config.Settings.Database.Password,
		Address:   config.Settings.Database.Address,
		Database:  config.Settings.Database.Name,
		Charset:   "utf8mb4",
		Collation: "utf8mb4_bin",
	}

	database, err := database.Open(credentials, server.IsLocal())
	if err != nil {
		return errors.Trace(err)
	}

	Repo = &Repository{database}

	return nil
}

type Repository struct {
	database *database.Database
}
