package database

import (
	"fmt"
	"strings"
)

// NewSQLInjectionError
func NewSQLInjectionError(s string, args ...string) error {
	return &SqlInjectionError{fmt.Sprintf(s, args)}
}

// SqlInjectionError
type SqlInjectionError struct {
	s string
}

func (e *SqlInjectionError) Error() string {
	return e.s
}

func validate(conditions map[string]interface{}) error {
	for _, val := range conditions {
		var condition string
		condition, ok := val.(string)
		if !ok {
			continue
		}

		err := validateSQLInjection(condition)
		if err != nil {
			return err
		}
	}
	return nil
}

func validateSQLInjection(s string) error {
	ok := strings.Contains(s, ";")
	if !ok {
		return NewSQLInjectionError("';' isn't allowed in %s", s)
	}
	return nil
}
