package database

import (
	"time"
)

type (
	// Model provide us to determine wathever struct is a Model.
	Model interface {
		TableName() string
	}

	// AuditorModel provides the audit model columns.
	AuditModel struct {
		CreatedAt time.Time `json:"createdAt"`
		UpdatedAt time.Time `json:"updatedAt"`
	}
)
