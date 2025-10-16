package internal

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

// LocalTime wraps time.Time to customize JSON but remain SQL-compatible
type LocalTime struct {
	time.Time
}

// --- JSON ---

func (t LocalTime) MarshalJSON() ([]byte, error) {
	if t.Time.IsZero() {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf(`"%s"`, t.Format("2006-01-02T15:04:05"))), nil
}

func (t *LocalTime) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	if s == "" {
		t.Time = time.Time{}
		return nil
	}
	parsed, err := time.Parse("2006-01-02T15:04:05", s)
	if err != nil {
		return err
	}
	t.Time = parsed
	return nil
}

// --- SQL ---

// Value implements driver.Valuer so the DB can store it
func (t LocalTime) Value() (driver.Value, error) {
	if t.IsZero() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan implements sql.Scanner so sqlx can load it
func (lt *LocalTime) Scan(value any) error {
	if value == nil {
		lt.Time = time.Time{}
		return nil
	}

	switch v := value.(type) {
	case time.Time:
		lt.Time = v
	case []byte:
		t, err := time.Parse("2006-01-02 15:04:05", string(v))
		if err != nil {
			return err
		}
		lt.Time = t
	case string:
		t, err := time.Parse("2006-01-02 15:04:05", v)
		if err != nil {
			return err
		}
		lt.Time = t
	default:
		return fmt.Errorf("cannot scan type %T into LocalTime", value)
	}
	return nil
}
