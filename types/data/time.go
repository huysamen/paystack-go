package data

import (
	"encoding/json"
	"fmt"
	"time"
)

// timeFormats are the time formats we'll try to parse during JSON unmarshaling
// This is shared between Time and NullTime to ensure consistent behavior
var timeFormats = []string{
	time.RFC3339,               // "2006-01-02T15:04:05Z07:00"
	time.RFC3339Nano,           // "2006-01-02T15:04:05.999999999Z07:00"
	"2006-01-02T15:04:05Z",     // UTC format (explicit Z)
	"2006-01-02T15:04:05.000Z", // Custom format with .000Z
	"2006-01-02 15:04:05",      // Simple datetime format
	"2006-01-02T15:04:05",      // ISO format without timezone
	"2006-01-02",               // Date only
}

// Time represents a time.Time that can be unmarshaled from various JSON types
// Null values are converted to zero time instead of preserving null state
// Accepts: RFC3339 strings, RFC3339Nano strings, and null (→ zero time)
type Time time.Time

// UnmarshalJSON implements json.Unmarshaler for Time
func (t *Time) UnmarshalJSON(data []byte) error {
	// Handle null - convert to zero time
	if data == nil || string(data) == "null" {
		*t = Time(time.Time{})
		return nil
	}

	// Try to unmarshal as string and parse as time
	var s string

	if err := json.Unmarshal(data, &s); err == nil {
		// Handle empty string and "null" string as zero time
		if s == "" || s == "null" {
			*t = Time(time.Time{})
			return nil
		}

		// Try various time formats
		for _, format := range timeFormats {
			if parsed, err := time.Parse(format, s); err == nil {
				*t = Time(parsed)
				return nil
			}
		}

		return fmt.Errorf("cannot parse string %q as time", s)
	}

	return fmt.Errorf("cannot unmarshal %s into Time", string(data))
}

// MarshalJSON implements json.Marshaler for Time
func (t Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(t).Format(time.RFC3339))
}

// Time returns the time.Time value
func (t Time) Time() time.Time {
	return time.Time(t)
}

// String returns the string representation of the time
func (t Time) String() string {
	return time.Time(t).Format(time.RFC3339)
}

// IsZero reports whether t represents the zero time instant
func (t Time) IsZero() bool {
	return time.Time(t).IsZero()
}

// NewTime creates a new Time with the given value
func NewTime(value time.Time) Time {
	return Time(value)
}

// NewTimeFromString creates a new Time by parsing a string
func NewTimeFromString(s string) (Time, error) {
	// Handle empty string and "null" string as zero time
	if s == "" || s == "null" {
		return Time(time.Time{}), nil
	}

	// Try various time formats
	for _, format := range timeFormats {
		if parsed, err := time.Parse(format, s); err == nil {
			return Time(parsed), nil
		}
	}

	return Time{}, fmt.Errorf("cannot parse string %q as time", s)
}
