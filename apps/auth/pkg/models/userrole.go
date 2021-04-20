package models

import "golang-seed/pkg/database"

type UserRole struct {
	database.AuditModel

	UserID   string `gorm:"primaryKey" json:"userID"`
	RoleCode string `gorm:"primaryKey" json:"roleCode"`
}

func (u UserRole) TableName() string {
	return "user_role"
}
