package data

import (
	"encoding/json"
	"fmt"
	"time"
)

// MultiDateTime represents a time value that can be marshaled/unmarshaled from various formats
type MultiDateTime struct {
	time.Time
}

// NewMultiDateTime creates a new MultiDateTime from a time.Time
func NewMultiDateTime(t time.Time) MultiDateTime {
	return MultiDateTime{Time: t}
}

// MarshalJSON implements json.Marshaler for MultiDateTime
func (dt MultiDateTime) MarshalJSON() ([]byte, error) {
	if dt.IsZero() {
		return []byte("null"), nil
	}
	return json.Marshal(dt.Format(time.RFC3339))
}

// UnmarshalJSON implements json.Unmarshaler for MultiDateTime
func (dt *MultiDateTime) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	if s == "" || s == "null" {
		*dt = MultiDateTime{}
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
			*dt = MultiDateTime{Time: t}
			return nil
		}
	}

	return fmt.Errorf("unsupported time format: %s", s)
}

// String returns the string representation of MultiDateTime
func (dt MultiDateTime) String() string {
	if dt.IsZero() {
		return ""
	}
	return dt.Format(time.RFC3339)
}

// Unix returns the Unix timestamp
func (dt MultiDateTime) Unix() int64 {
	return dt.Time.Unix()
}
