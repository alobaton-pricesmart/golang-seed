package models

import (
	"github.com/altipla-consulting/database"
	"github.com/juju/errors"

	"github.com/alobaton/golang-seed/apps/auth/pkg/config"
)

var Repo *MainDatabase

func ConnectRepo() error {
	credentials := database.Credentials{
		User:      config.Settings.Database.User,
		Password:  config.Settings.Database.Password,
		Address:   config.Settings.Database.Address,
		Database:  "app_auth",
		Charset:   "utf8mb4",
		Collation: "utf8mb4_bin",
	}

	sess, err := database.Open(credentials)
	if err != nil {
		return errors.Trace(err)
	}

	Repo = &MainDatabase{sess}

	return nil
}

type MainDatabase struct {
	sess *database.Database
}

func (repo *MainDatabase) Clients() *database.Collection {
	return repo.sess.Collection(new(Client))
}
