package models

import (
	"github.com/juju/errors"
	"gorm.io/gorm"

	"golang-seed/apps/auth/pkg/config"
	"golang-seed/pkg/database"
	"golang-seed/pkg/service"
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

	database, err := database.Open(credentials, service.IsLocal())
	if err != nil {
		return errors.Trace(err)
	}

	Repo = &Repository{database}

	Repo.database.Migrate(new(Client))
	Repo.database.Migrate(new(User))

	return nil
}

type Repository struct {
	database *database.Database
}

func (r *Repository) Clients() *gorm.DB {
	return r.database.Model(new(Client))
}

func (r *Repository) Users() *gorm.DB {
	return r.database.Model(new(User))
}
