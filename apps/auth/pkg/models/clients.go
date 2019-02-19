package models

import (
	"gopkg.in/oauth2.v3"

	"github.com/altipla-consulting/database"
)

type Client struct {
	database.ModelTracking
	oauth2.ClientInfo

	Code     string `db:"code,pk"`
	Secret   string `db:"secret"`
	Domain   string `db:"domain"`
	UserCode string `db:"user_code"`
}

// GetID client id
func (c *Client) GetID() string {
	return c.Code
}

// GetSecret client domain
func (c *Client) GetSecret() string {
	return c.Secret
}

// GetDomain client domain
func (c *Client) GetDomain() string {
	return c.Domain
}

// GetUserID user id
func (c *Client) GetUserID() string {
	return c.UserCode
}

func (client *Client) TableName() string {
	return "clients"
}
