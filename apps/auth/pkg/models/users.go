package models

import (
	"github.com/altipla-consulting/database"
)

type User struct {
	database.ModelTracking

	Nickname string `db:"nickname,pk"`
}

func (user *User) TableName() string {
	return "users"
}