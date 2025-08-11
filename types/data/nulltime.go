package data

import (
	"encoding/json"
	"fmt"
	"time"
)

// NullTime represents a time.Time that may be null in JSON
// Unlike MultiDateTime, this preserves null values instead of converting them to zero time
// Accepts: RFC3339 strings, various datetime formats, and null
type NullTime struct {
	Time  time.Time
	Valid bool // true if Time is not null
}

// UnmarshalJSON implements json.Unmarshaler for NullTime
func (nt *NullTime) UnmarshalJSON(data []byte) error {
	// Handle null
	if data == nil || string(data) == "null" {
		nt.Time = time.Time{}
		nt.Valid = false

		return nil
	}

	// Try to unmarshal as string first
	var s string

	if err := json.Unmarshal(data, &s); err == nil {
		// Handle empty string and "null" string as null
		if s == "" || s == "null" {
			nt.Time = time.Time{}
			nt.Valid = false

			return nil
		}

		// Try to parse the string as time using various formats
		for _, format := range timeFormats {
			if parsedTime, err := time.Parse(format, s); err == nil {
				nt.Time = parsedTime
				nt.Valid = true

				return nil
			}
		}

		return fmt.Errorf("cannot parse string %q as time", s)
	}

	return fmt.Errorf("cannot unmarshal %s into NullTime", string(data))
}

// MarshalJSON implements json.Marshaler for NullTime
func (nt NullTime) MarshalJSON() ([]byte, error) {
	if !nt.Valid {
		return []byte("null"), nil
	}

	// Use RFC3339Nano if there are nanoseconds, otherwise RFC3339
	format := time.RFC3339
	if nt.Time.Nanosecond() != 0 {
		format = time.RFC3339Nano
	}

	return json.Marshal(nt.Time.Format(format))
}

// ValueOr returns the time value if valid, otherwise returns the fallback value
func (nt NullTime) ValueOr(fallback time.Time) time.Time {
	if nt.Valid {
		return nt.Time
	}

	return fallback
}

// String returns the string representation of the time, or "null" if not valid
func (nt NullTime) String() string {
	if !nt.Valid {
		return "null"
	}

	// Use RFC3339Nano if there are nanoseconds, otherwise RFC3339
	format := time.RFC3339
	if nt.Time.Nanosecond() != 0 {
		format = time.RFC3339Nano
	}

	return nt.Time.Format(format)
}

// IsZero returns true if the time is not valid or is the zero time
func (nt NullTime) IsZero() bool {
	return !nt.Valid || nt.Time.IsZero()
}

// Unix returns the Unix timestamp if valid, otherwise returns 0
func (nt NullTime) Unix() int64 {
	if !nt.Valid {
		return 0
	}

	return nt.Time.Unix()
}

// NewNullTime creates a new valid NullTime with the given value
func NewNullTime(value time.Time) NullTime {
	return NullTime{
		Time:  value,
		Valid: true,
	}
}

// NewNullTimeFromString creates a new NullTime by parsing the given string
func NewNullTimeFromString(value string) (NullTime, error) {
	// Handle empty string and "null" string as null
	if value == "" || value == "null" {
		return NullTime{
			Time:  time.Time{},
			Valid: false,
		}, nil
	}

	// Try to parse the string as time using various formats
	for _, format := range timeFormats {
		if parsedTime, err := time.Parse(format, value); err == nil {
			return NullTime{
				Time:  parsedTime,
				Valid: true,
			}, nil
		}
	}

	return NullTime{}, fmt.Errorf("cannot parse string %q as time", value)
}
