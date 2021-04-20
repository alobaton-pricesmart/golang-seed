package models

import "golang-seed/pkg/database"

type UserRole struct {
	database.AuditModel

	UserID   uint64 `gorm:"primaryKey"`
	RoleCode string `gorm:"primaryKey"`
}

func (u UserRole) TableName() string {
	return "user_role"
}
