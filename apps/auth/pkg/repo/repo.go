package repo

import (
	"github.com/juju/errors"

	"golang-seed/apps/auth/pkg/config"
	"golang-seed/apps/auth/pkg/models"
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

	Repo.database.Migrate(new(models.Client))
	Repo.database.Migrate(new(models.Permission))
	Repo.database.Migrate(new(models.Role))
	Repo.database.Migrate(new(models.RolePermission))
	Repo.database.SetupJoinTable(new(models.Role), "RolePermissions", new(models.RolePermission))
	Repo.database.Migrate(new(models.User))
	Repo.database.Migrate(new(models.UserRole))
	Repo.database.SetupJoinTable(new(models.User), "UserRoles", new(models.UserRole))

	return nil
}

type Repository struct {
	database *database.Database
}

func (r *Repository) Clients() *database.Collection {
	return r.database.Collection(new(models.Client))
}

func (r *Repository) Permissions() *database.Collection {
	return r.database.Collection(new(models.Permission))
}

func (r *Repository) Roles() *database.Collection {
	return r.database.Collection(new(models.Role))
}

func (r *Repository) RolePermissions() *database.Collection {
	return r.database.Collection(new(models.RolePermission))
}

func (r *Repository) Users() *database.Collection {
	return r.database.Collection(new(models.User))
}

func (r *Repository) UserRoles() *database.Collection {
	return r.database.Collection(new(models.UserRole))
}
