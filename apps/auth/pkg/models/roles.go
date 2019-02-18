package models

import (
	"github.com/altipla-consulting/database"
)

type Role struct {
	database.ModelTracking

	Code string `db:"code,pk"`
}

func (role *Role) TableName() string {
	return "roles"
}