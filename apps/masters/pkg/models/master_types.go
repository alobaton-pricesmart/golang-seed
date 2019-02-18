package models

import (
	"github.com/altipla-consulting/database"
)

type MasterType struct {
	database.ModelTracking

	Code string `db:"code,pk"`
}

func (masterType *MasterType) TableName() string {
	return "master_types"
}
