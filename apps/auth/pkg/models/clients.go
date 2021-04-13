package models

import (
	"github.com/go-oauth2/oauth2/v4"
)

type Client struct {
	oauth2.ClientInfo `gorm:"-"`

	Code   string `gorm:"primaryKey"`
	Secret string
	Domain string
	UserID string
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

func (client *Client) TableName() string {
	return "clients"
}
