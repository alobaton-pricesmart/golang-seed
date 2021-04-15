package models

import "golang-seed/pkg/database"

type RolePermission struct {
	database.AuditModel

	RoleCode       string `gorm:"primaryKey" json:"roleCode"`
	PermissionCode string `gorm:"primaryKey" json:"permissionCode"`
}

func (u RolePermission) TableName() string {
	return "role_permission"
}
