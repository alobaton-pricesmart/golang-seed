package models

import "golang-seed/pkg/database"

type Permission struct {
	database.AuditModel

	Code        string `gorm:"primaryKey" json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (r Permission) TableName() string {
	return "permissions"
}
