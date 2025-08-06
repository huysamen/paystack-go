package types

import (
	"encoding/json"
	"strings"
	"time"
)

// DateTime wraps time.Time with custom JSON marshaling for Paystack API formats
type DateTime struct {
	time.Time
}

// Common time formats used by Paystack API
var paystackTimeFormats = []string{
	time.RFC3339,  // "2006-01-02T15:04:05Z07:00"
	time.DateTime, // "2006-01-02 15:04:05"
	"2006-01-02T15:04:05.000Z",
	"2006-01-02T15:04:05.000000Z",
	"2006-01-02", // Date only
}

// NewDateTime creates a new DateTime from time.Time
func NewDateTime(t time.Time) DateTime {
	return DateTime{Time: t}
}

// NewDateTimePtr creates a new DateTime pointer from time.Time
func NewDateTimePtr(t time.Time) *DateTime {
	return &DateTime{Time: t}
}

// MarshalJSON implements json.Marshaler interface
func (dt DateTime) MarshalJSON() ([]byte, error) {
	if dt.Time.IsZero() {
		return []byte("null"), nil
	}
	return json.Marshal(dt.Time.Format(time.RFC3339))
}

// UnmarshalJSON implements json.Unmarshaler interface with support for multiple Paystack date formats
func (dt *DateTime) UnmarshalJSON(b []byte) error {
	if string(b) == "null" || string(b) == `""` {
		dt.Time = time.Time{}
		return nil
	}

	str := strings.Trim(string(b), `"`)
	if str == "" {
		dt.Time = time.Time{}
		return nil
	}

	// Try each format until one succeeds
	for _, format := range paystackTimeFormats {
		if parsedTime, err := time.Parse(format, str); err == nil {
			dt.Time = parsedTime
			return nil
		}
	}

	// If none of the standard formats work, try parsing as Unix timestamp
	var timestamp int64
	if err := json.Unmarshal(b, &timestamp); err == nil {
		dt.Time = time.Unix(timestamp, 0)
		return nil
	}

	return &time.ParseError{
		Layout:     "Paystack datetime formats",
		Value:      str,
		LayoutElem: "any supported format",
		ValueElem:  str,
	}
}
