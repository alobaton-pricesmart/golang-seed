package models

import (
	"golang-seed/pkg/database"

	"github.com/google/uuid"
)

type User struct {
	database.AuditModel

	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Nickname string    `gorm:"index"`
	Name     string
	LastName string
	Email    string `gorm:"index"`
	Password string
	Locked   bool
	Enabled  bool

	UserRoles []UserRole `gorm:"many2many:user_role;"`
}

func (u User) TableName() string {
	return "users"
}
