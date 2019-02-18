package models

import (
	"github.com/altipla-consulting/database"
)

type Master struct {
	database.ModelTracking

	Code string `db:"code,pk"`
}

func (master *Master) TableName() string {
	return "masters"
}
