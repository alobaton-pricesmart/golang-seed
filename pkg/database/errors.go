package database

import "gorm.io/gorm"

var (
	// ErrRecordNotFound record not found error
	ErrRecordNotFound = gorm.ErrRecordNotFound
	// ErrRegistered registered
	ErrRegistered = gorm.ErrRegistered
)
