package data

import (
	"encoding/json"
	"fmt"
)

// NullString represents a string that may be null in JSON
// Unlike MultiString, this preserves null values instead of converting them to empty strings
// Accepts: strings, numbers (converted to strings), and null
type NullString struct {
	Str   string
	Valid bool // true if Str is not null
}

// UnmarshalJSON implements json.Unmarshaler for NullString
func (ns *NullString) UnmarshalJSON(data []byte) error {
	// Handle null
	if data == nil || string(data) == "null" {
		ns.Str = ""
		ns.Valid = false

		return nil
	}

	// Try to unmarshal as string first
	var s string

	if err := json.Unmarshal(data, &s); err == nil {
		// Handle empty string and "null" string as null
		if s == "" || s == "null" {
			ns.Str = ""
			ns.Valid = false

			return nil
		}

		ns.Str = s
		ns.Valid = true

		return nil
	}

	// Try to unmarshal as number and convert to string
	var n float64

	if err := json.Unmarshal(data, &n); err == nil {
		ns.Str = fmt.Sprintf("%g", n) // Use %g for clean number representation
		ns.Valid = true

		return nil
	}

	return fmt.Errorf("cannot unmarshal %s into NullString", string(data))
}

// MarshalJSON implements json.Marshaler for NullString
func (ns NullString) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(ns.Str)
}

// ValueOr returns the string value if valid, otherwise returns the fallback value
func (ns NullString) ValueOr(fallback string) string {
	if ns.Valid {
		return ns.Str
	}

	return fallback
}

// String returns the string representation, or "null" if not valid
func (ns NullString) String() string {
	if !ns.Valid {
		return "null"
	}

	return ns.Str
}

// NewNullString creates a new valid NullString with the given value
func NewNullString(value string) NullString {
	return NullString{
		Str:   value,
		Valid: true,
	}
}
