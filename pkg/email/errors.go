package email

import "errors"

var (
	ErrInvalidTemplate = errors.New("no valid template")
	ErrSendingEmail    = errors.New("unable to sent email")
)
