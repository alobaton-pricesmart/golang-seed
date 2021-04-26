package models

import (
	"fmt"
	"golang-seed/pkg/database"
	"strings"
)

type Permission struct {
	database.AuditModel

	Code        string `gorm:"primaryKey" json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (p Permission) TableName() string {
	return "permissions"
}

func (p Permission) String() string {
	var b strings.Builder

	if len(p.Code) > 0 {
		fmt.Fprintf(&b, " code : %s ", p.Code)
	}

	if len(p.Name) > 0 {
		fmt.Fprintf(&b, " name : %s ", p.Name)
	}

	if len(p.Description) > 0 {
		fmt.Fprintf(&b, " description : %s ", p.Description)
	}

	return b.String()
}
