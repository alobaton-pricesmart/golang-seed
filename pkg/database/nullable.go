package database

import (
	"database/sql"
	"encoding/json"
)

type (
	NullString struct {
		sql.NullString
	}
	NullInt64 struct {
		sql.NullInt64
	}
)

func (t NullString) MarshalJSON() ([]byte, error) {
	if t.Valid {
		return json.Marshal(t.String)
	} else {
		return json.Marshal(nil)
	}
}

func (t *NullString) UnmarshalJSON(data []byte) error {
	var v *string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	if v != nil {
		t.Valid = true
		t.String = *v
	} else {
		t.Valid = false
	}
	return nil
}

func (t NullInt64) MarshalJSON() ([]byte, error) {
	if t.Valid {
		return json.Marshal(t.Int64)
	} else {
		return json.Marshal(nil)
	}
}

func (t *NullInt64) UnmarshalJSON(data []byte) error {
	var v *int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	if v != nil {
		t.Valid = true
		t.Int64 = *v
	} else {
		t.Valid = false
	}
	return nil
}
