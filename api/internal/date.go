package internal

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

func init() {
	var err error
	osloLoc, err = time.LoadLocation("Europe/Oslo")
	if err != nil {
		panic(fmt.Sprintf("failed to load Europe/Oslo location: %v", err))
	}
}

// Date wraps time.Time to handle JSON and SQL marshalling/unmarshalling
// @swaggertype string
// @example "2023-10-05"
type Date struct {
	time.Time
}

// --- JSON ---

func (t Date) MarshalJSON() ([]byte, error) {
	if t.Time.IsZero() {
		return []byte("null"), nil
	}
	// Format in Oslo time
	osloTime := t.In(osloLoc)
	return []byte(fmt.Sprintf(`"%s"`, osloTime.Format("2006-01-02"))), nil
}

func (t *Date) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	if s == "" {
		t.Time = time.Time{}
		return nil
	}
	parsed, err := time.ParseInLocation("2006-01-02", s, osloLoc)
	if err != nil {
		return err
	}
	t.Time = parsed
	return nil
}

// --- SQL ---

func (t Date) Value() (driver.Value, error) {
	if t.IsZero() {
		return nil, nil
	}
	return t.UTC(), nil
}

func (lt *Date) Scan(value any) error {
	if value == nil {
		lt.Time = time.Time{}
		return nil
	}

	var parsed time.Time
	switch v := value.(type) {
	case time.Time:
		parsed = v
	case []byte:
		t, err := time.ParseInLocation("2006-01-02", string(v), osloLoc)
		if err != nil {
			return err
		}
		parsed = t
	case string:
		t, err := time.ParseInLocation("2006-01-02", v, osloLoc)
		if err != nil {
			return err
		}
		parsed = t
	default:
		return fmt.Errorf("cannot scan type %T into LocalTime", value)
	}

	lt.Time = parsed.In(osloLoc)
	return nil
}
