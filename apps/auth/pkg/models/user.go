package models

import (
	"golang-seed/pkg/database"

	"github.com/google/uuid"
)

type User struct {
	database.AuditModel

	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	Nickname string    `gorm:"index" json:"nickname"`
	Name     string    `json:"name"`
	LastName string    `json:"lastName"`
	Email    string    `gorm:"index"`
	Password string    `json:"password"`
	Locked   bool      `json:"locked"`
	Enabled  bool      `json:"enabled"`

	UserRoles []UserRole `gorm:"many2many:user_role;" json:"userRoles"`
}

func (u User) TableName() string {
	return "users"
}
