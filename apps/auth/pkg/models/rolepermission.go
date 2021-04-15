package models

import "golang-seed/pkg/database"

type RolePermission struct {
	database.AuditModel

	RoleCode       string `gorm:"primaryKey"`
	PermissionCode string `gorm:"primaryKey"`
}

func (u RolePermission) TableName() string {
	return "role_permission"
}
