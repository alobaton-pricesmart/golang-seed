package models

import (
	"github.com/juju/errors"

	"golang-seed/apps/auth/pkg/config"
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

	conf := database.Conf{
		MaxOpenConns: config.Settings.Database.MaxOpenConns,
		MaxIdleConns: config.Settings.Database.MaxIdleConns,
		Debug:        server.IsLocal(),
	}

	database, err := database.Open(credentials, conf)
	if err != nil {
		return errors.Trace(err)
	}

	Repo = &Repository{database}

	/**
	TODO(alobaton): Create a migrations cmd
	Repo.database.Migrate(new(Client))
	Repo.database.Migrate(new(Permission))
	Repo.database.Migrate(new(Role))
	Repo.database.Migrate(new(RolePermission))
	Repo.database.SetupJoinTable(new(Role), "RolePermissions", new(RolePermission))
	Repo.database.Migrate(new(User))
	Repo.database.Migrate(new(UserRole))
	Repo.database.SetupJoinTable(new(User), "UserRoles", new(UserRole))
	*/

	return nil
}

type Repository struct {
	database *database.Database
}

func (r *Repository) Clients() *database.Collection {
	return r.database.Collection(new(Client))
}

func (r *Repository) Permissions() *database.Collection {
	return r.database.Collection(new(Permission))
}

func (r *Repository) Roles() *database.Collection {
	return r.database.Collection(new(Role))
}

func (r *Repository) RolePermissions() *database.Collection {
	return r.database.Collection(new(RolePermission))
}

func (r *Repository) Users() *database.Collection {
	return r.database.Collection(new(User))
}

func (r *Repository) UserRoles() *database.Collection {
	return r.database.Collection(new(UserRole))
}
