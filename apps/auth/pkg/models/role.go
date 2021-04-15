package models

import "golang-seed/pkg/database"

type Role struct {
	database.AuditModel

	Code        string `gorm:"primaryKey" json:"code"`
	GroupID     string `json:"groupID"`
	Name        string `json:"name"`
	Description string `json:"description"`

	RolePermissions []RolePermission `gorm:"many2many:role_permission;" json:"rolePermissions"`
}

func (r Role) TableName() string {
	return "roles"
}
