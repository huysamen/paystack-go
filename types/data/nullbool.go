package data

import (
	"encoding/json"
	"fmt"
)

// NullBool represents a bool that may be null in JSON
// Unlike MultiBool, this preserves null values instead of converting them to false
// Accepts: booleans, string representations ("true", "success", "1" as true), and null
type NullBool struct {
	Bool  bool
	Valid bool // true if Bool is not null
}

// UnmarshalJSON implements json.Unmarshaler for NullBool
func (nb *NullBool) UnmarshalJSON(data []byte) error {
	// Handle null
	if data == nil || string(data) == "null" {
		nb.Bool = false
		nb.Valid = false

		return nil
	}

	// Try to unmarshal as boolean first
	var b bool

	if err := json.Unmarshal(data, &b); err == nil {
		nb.Bool = b
		nb.Valid = true

		return nil
	}

	// Try to unmarshal as string
	var s string

	if err := json.Unmarshal(data, &s); err == nil {
		// Handle empty string and "null" string as null
		if s == "" || s == "null" {
			nb.Bool = false
			nb.Valid = false

			return nil
		}

		// Consider "true", "success", and "1" as true values, everything else as false
		nb.Bool = (s == "true" || s == "success" || s == "1")
		nb.Valid = true

		return nil
	}

	// Try to unmarshal as number (following MultiBool pattern of handling numbers)
	var n float64

	if err := json.Unmarshal(data, &n); err == nil {
		// Any non-zero number is considered true
		nb.Bool = n != 0
		nb.Valid = true

		return nil
	}

	return fmt.Errorf("cannot unmarshal %s into NullBool", string(data))
}

// MarshalJSON implements json.Marshaler for NullBool
func (nb NullBool) MarshalJSON() ([]byte, error) {
	if !nb.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(nb.Bool)
}

// ValueOr returns the boolean value if valid, otherwise returns the fallback value
func (nb NullBool) ValueOr(fallback bool) bool {
	if nb.Valid {
		return nb.Bool
	}

	return fallback
}

// String returns the string representation of the boolean, or "null" if not valid
func (nb NullBool) String() string {
	if !nb.Valid {
		return "null"
	}

	if nb.Bool {
		return "true"
	}

	return "false"
}

// NewNullBool creates a new valid NullBool with the given value
func NewNullBool(value bool) NullBool {
	return NullBool{
		Bool:  value,
		Valid: true,
	}
}
