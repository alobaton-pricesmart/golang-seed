package models

import (
	"github.com/altipla-consulting/database"
)

type User struct {
	database.ModelTracking

	Code string `db:"code,pk"`
}

func (user *User) TableName() string {
	return "users"
}
