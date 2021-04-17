package models

import (
	"golang-seed/pkg/database"

	"github.com/go-oauth2/oauth2/v4"
)

type Client struct {
	database.AuditModel
	oauth2.ClientInfo `gorm:"-" json:"-"`

	Code   string `gorm:"primaryKey" json:"code"`
	Secret string `json:"-"`
	Domain string `json:"domain"`
	UserID string `json:"-"`
}

// Implements oauth2.ClientInfo interface
func (c Client) GetID() string {
	return c.Code
}

func (c Client) GetSecret() string {
	return c.Secret
}

func (c Client) GetDomain() string {
	return c.Domain
}

func (c Client) GetUserID() string {
	return c.UserID
}

func (c Client) TableName() string {
	return "clients"
}
