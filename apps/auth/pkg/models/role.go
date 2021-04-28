package models

import (
	"fmt"
	"golang-seed/pkg/database"
	"strings"
)

type Role struct {
	database.AuditModel

	Code        string `gorm:"primaryKey" json:"code"`
	GroupID     string `json:"groupID"`
	Name        string `json:"name"`
	Description string `json:"description"`

	RolePermissions []RolePermission `json:"rolePermissions"`
}

func (r Role) TableName() string {
	return "roles"
}

func (r Role) String() string {
	var b strings.Builder

	if len(r.Code) > 0 {
		fmt.Fprintf(&b, " code : %s ", r.Code)
	}

	if len(r.GroupID) > 0 {
		fmt.Fprintf(&b, " groupID : %s ", r.GroupID)
	}

	if len(r.Name) > 0 {
		fmt.Fprintf(&b, " name : %s ", r.Name)
	}

	if len(r.Description) > 0 {
		fmt.Fprintf(&b, " description : %s ", r.Description)
	}

	return b.String()
}
