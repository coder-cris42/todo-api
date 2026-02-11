package entities

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

const (
	DBDateFormat  = "2006-01-02 15:04:05"
	RFC3339Format = time.RFC3339
)

// DateTime is a custom time type that handles both MySQL DATETIME and RFC3339 formats
type DateTime struct {
	time.Time
}

// NewDateTime creates a new DateTime from a time.Time
func NewDateTime(t time.Time) DateTime {
	return DateTime{Time: t}
}

// Now returns the current time as DateTime
func Now() DateTime {
	return DateTime{Time: time.Now()}
}

// UnmarshalJSON handles unmarshaling datetime from JSON in multiple formats
func (dt *DateTime) UnmarshalJSON(b []byte) error {
	str := strings.Trim(string(b), `"`)

	if str == "" || str == "null" {
		dt.Time = time.Time{}
		return nil
	}

	// Try MySQL DATETIME format first
	t, err := time.Parse(DBDateFormat, str)
	if err == nil {
		dt.Time = t
		return nil
	}

	// Try RFC3339 format
	t, err = time.Parse(RFC3339Format, str)
	if err == nil {
		dt.Time = t
		return nil
	}

	// Try other common formats
	formats := []string{
		"2006-01-02T15:04:05Z07:00",
		"2006-01-02T15:04:05",
		"2006-01-02",
		time.RFC3339Nano,
	}

	for _, layout := range formats {
		t, err := time.Parse(layout, str)
		if err == nil {
			dt.Time = t
			return nil
		}
	}

	return fmt.Errorf("unable to parse datetime: %s", str)
}

// MarshalJSON handles marshaling datetime to JSON in RFC3339 format
func (dt DateTime) MarshalJSON() ([]byte, error) {
	if dt.Time.IsZero() {
		return []byte("null"), nil
	}
	return []byte(`"` + dt.Time.Format(RFC3339Format) + `"`), nil
}

// Value implements the driver.Valuer interface for database operations
func (dt DateTime) Value() (driver.Value, error) {
	// Always format the time, even if zero, to avoid NULL values in NOT NULL columns
	return dt.Time.Format(DBDateFormat), nil
}

// Scan implements the sql.Scanner interface for database reads
func (dt *DateTime) Scan(value interface{}) error {
	if value == nil {
		dt.Time = time.Time{}
		return nil
	}

	switch v := value.(type) {
	case time.Time:
		dt.Time = v
		return nil
	case string:
		// Try MySQL DATETIME format first
		t, err := time.Parse(DBDateFormat, v)
		if err == nil {
			dt.Time = t
			return nil
		}

		// Try RFC3339 format
		t, err = time.Parse(RFC3339Format, v)
		if err == nil {
			dt.Time = t
			return nil
		}

		return fmt.Errorf("unable to scan datetime: %s", v)
	case []byte:
		return dt.Scan(string(v))
	default:
		return fmt.Errorf("unable to scan type %T as DateTime", value)
	}
}
