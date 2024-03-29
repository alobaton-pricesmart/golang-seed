package database

import (
	"fmt"
)

// Credentials configures the authentication and address of the remote MySQL database.
type Credentials struct {
	User      string
	Password  string
	Address   string
	Database  string
	Charset   string
	Collation string
	Protocol  string
}

// String returns the credentials with the exact format the Go MySQL driver needs
// to connect to it.
func (c Credentials) String() string {
	protocol := c.Protocol
	if protocol == "" {
		protocol = "tcp"
	}

	var charset string
	if c.Charset != "" {
		charset = fmt.Sprintf("&charset=%s", c.Charset)
	}
	var collation string
	if c.Collation != "" {
		collation = fmt.Sprintf("&collation=%s", c.Collation)
	}

	return fmt.Sprintf("%s:%s@%s(%s)/%s?parseTime=true%s%s", c.User, c.Password, protocol, c.Address, c.Database, charset, collation)
}
