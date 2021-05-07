package email

import (
	"bytes"
	"fmt"
	"net/smtp"
)

// Email holds the needed data to send an email
type Email struct {
	TemplatePath string
	Template     string
	Data         map[string]interface{}
	To           []string
	Subject      string

	Host     string
	Port     int
	From     string
	Password string
}

// Send send an email
func (e Email) Send() error {
	auth := smtp.PlainAuth("", e.From, e.Password, e.Host)

	t, err := tmpl(e.TemplatePath, e.Template)
	if err != nil {
		return err
	}

	var body bytes.Buffer
	body.Write([]byte(e.Subject))

	t.Execute(&body, e.Data)

	err = smtp.SendMail(fmt.Sprintf("%s:%d", e.Host, e.Port), auth, e.From, e.To, body.Bytes())
	if err != nil {
		return ErrSendingEmail
	}

	return nil
}
