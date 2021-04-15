package database

import (
	"time"
)

type (
	Model interface {
		TableName() string
	}

	AuditModel struct {
		CreatedAt time.Time `json:"createdAt"`
		UpdatedAt time.Time `json:"updatedAt"`
	}
)
