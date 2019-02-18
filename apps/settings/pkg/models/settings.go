package models

import (
	"github.com/altipla-consulting/database"
)

type Setting struct {
	database.ModelTracking

	Code string `db:"code,pk"`
}

func (setting *Setting) TableName() string {
	return "settings"
}
