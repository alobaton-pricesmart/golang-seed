package models

import "golang-seed/pkg/database"

type Role struct {
	database.AuditModel

	Code        string `gorm:"primaryKey"`
	GroupID     string
	Name        string
	Description string

	RolePermissions []RolePermission `gorm:"many2many:role_permission;"`
}

func (r Role) TableName() string {
	return "roles"
}
