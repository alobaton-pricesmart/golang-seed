package models

import (
	"fmt"
	"golang-seed/pkg/database"
	"strings"

	"github.com/go-oauth2/oauth2/v4"
)

type Client struct {
	database.AuditModel
	oauth2.ClientInfo `gorm:"-" json:"-"`

	Code   string              `gorm:"primaryKey" json:"code"`
	Secret string              `json:"secret,omitempty"`
	Domain database.NullString `json:"domain"`
	UserID database.NullString `json:"-"`
}

// Implements oauth2.ClientInfo interface
func (c Client) GetID() string {
	return c.Code
}

func (c Client) GetSecret() string {
	return c.Secret
}

func (c Client) GetDomain() string {
	if c.Domain.Valid {
		return c.Domain.String
	}
	return ""
}

func (c Client) GetUserID() string {
	if c.UserID.Valid {
		return c.UserID.String
	}
	return ""
}

func (c Client) TableName() string {
	return "clients"
}

func (c Client) String() string {
	var b strings.Builder

	if len(c.Code) > 0 {
		fmt.Fprintf(&b, " code : %s ", c.Code)
	}

	if c.Domain.Valid {
		fmt.Fprintf(&b, " domain : %s ", c.Domain.String)
	}

	if c.UserID.Valid {
		fmt.Fprintf(&b, " userID : %s ", c.UserID.String)
	}

	return b.String()
}
