package models

import (
	"fmt"
	"golang-seed/pkg/database"
	"strings"
)

type User struct {
	database.AuditModel

	ID       string              `gorm:"primaryKey" json:"id"`
	Nickname string              `gorm:"index" json:"nickname"`
	Name     string              `json:"name"`
	LastName database.NullString `json:"lastName"`
	Email    string              `gorm:"index" json:"email"`
	Password string              `json:"password,omitempty"`
	Locked   bool                `json:"locked"`
	Enabled  bool                `json:"enabled"`

	UserRoles []UserRole `gorm:"many2many:user_role;" json:"userRoles"`
}

func (u User) TableName() string {
	return "users"
}

func (u User) String() string {
	var b strings.Builder

	if len(u.ID) > 0 {
		fmt.Fprintf(&b, " id : %s ", u.ID)
	}

	if len(u.Nickname) > 0 {
		fmt.Fprintf(&b, " nickname : %s ", u.Nickname)
	}

	if len(u.Name) > 0 {
		fmt.Fprintf(&b, " name : %s ", u.Name)
	}

	if u.LastName.Valid {
		fmt.Fprintf(&b, " lastName : %s ", u.LastName.String)
	}

	if len(u.Email) > 0 {
		fmt.Fprintf(&b, " email : %s ", u.Email)
	}

	fmt.Fprintf(&b, " locked : %t ", u.Locked)

	fmt.Fprintf(&b, " enabled : %t ", u.Locked)

	return b.String()
}
