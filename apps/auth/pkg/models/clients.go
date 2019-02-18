package models

import (
	"gopkg.in/oauth2.v3"

	"github.com/altipla-consulting/database"
)

type Client struct {
	database.ModelTracking

	Code   string `db:"code,pk"`
	Secret string `db:"secret,pk"`
	Domain string `db:"domain,pk"`

	oauth2.ClientInfo
}

func (client *Client) TableName() string {
	return "clients"
}
