package models

import (
	"encoding/json"
	"fmt"
	"time"
)

// DateTime represents a time value that can be marshaled/unmarshaled from various formats
type DateTime struct {
	time.Time
}

// NewDateTime creates a new DateTime from a time.Time
func NewDateTime(t time.Time) DateTime {
	return DateTime{Time: t}
}

// MarshalJSON implements json.Marshaler for DateTime
func (dt DateTime) MarshalJSON() ([]byte, error) {
	if dt.IsZero() {
		return []byte("null"), nil
	}
	return json.Marshal(dt.Format(time.RFC3339))
}

// UnmarshalJSON implements json.Unmarshaler for DateTime
func (dt *DateTime) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	if s == "" || s == "null" {
		*dt = DateTime{}
		return nil
	}

	// Try different time formats
	formats := []string{
		time.RFC3339,
		time.RFC3339Nano,
		"2006-01-02T15:04:05.000Z",
		"2006-01-02T15:04:05Z",
		"2006-01-02 15:04:05",
	}

	for _, format := range formats {
		if t, err := time.Parse(format, s); err == nil {
			*dt = DateTime{Time: t}
			return nil
		}
	}

	return fmt.Errorf("unsupported time format: %s", s)
}

// String returns the string representation of DateTime
func (dt DateTime) String() string {
	if dt.IsZero() {
		return ""
	}
	return dt.Format(time.RFC3339)
}

// Unix returns the Unix timestamp
func (dt DateTime) Unix() int64 {
	return dt.Time.Unix()
}
