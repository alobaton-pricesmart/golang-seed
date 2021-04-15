package models

import "golang-seed/pkg/database"

type Permission struct {
	database.AuditModel

	Code        string `gorm:"primaryKey"`
	Name        string
	Description string
}

func (r Permission) TableName() string {
	return "permissions"
}
