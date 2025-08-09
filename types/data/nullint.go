package data

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// NullInt represents an int64 that may be null in JSON
// Unlike MultiInt, this preserves null values instead of converting them to zero
// Accepts: integers, floats (truncated), string representations of numbers, and null
type NullInt struct {
	Int   int64
	Valid bool // true if Int is not null
}

// UnmarshalJSON implements json.Unmarshaler for NullInt
func (ni *NullInt) UnmarshalJSON(data []byte) error {
	// Handle null
	if data == nil || string(data) == "null" {
		ni.Int = 0
		ni.Valid = false

		return nil
	}

	// Try to unmarshal as integer first
	var i int64

	if err := json.Unmarshal(data, &i); err == nil {
		ni.Int = i
		ni.Valid = true

		return nil
	}

	// Try to unmarshal as float64 (JSON numbers default to float64)
	var f float64

	if err := json.Unmarshal(data, &f); err == nil {
		ni.Int = int64(f) // Truncate to integer
		ni.Valid = true

		return nil
	}

	// Try to unmarshal as string and parse as integer
	var s string

	if err := json.Unmarshal(data, &s); err == nil {
		// Handle empty string as null
		if s == "" || s == "null" {
			ni.Int = 0
			ni.Valid = false

			return nil
		}

		// Try to parse string as integer
		if parsed, err := strconv.ParseInt(s, 10, 64); err == nil {
			ni.Int = parsed
			ni.Valid = true

			return nil
		}

		// Try to parse string as float and truncate to integer
		if parsed, err := strconv.ParseFloat(s, 64); err == nil {
			ni.Int = int64(parsed)
			ni.Valid = true

			return nil
		}

		return fmt.Errorf("cannot parse string %q as integer", s)
	}

	return fmt.Errorf("cannot unmarshal %s into NullInt", string(data))
}

// MarshalJSON implements json.Marshaler for NullInt
func (ni NullInt) MarshalJSON() ([]byte, error) {
	if !ni.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(ni.Int)
}

// ValueOr returns the integer value if valid, otherwise returns the fallback value
func (ni NullInt) ValueOr(fallback int64) int64 {
	if ni.Valid {
		return ni.Int
	}

	return fallback
}

// String returns the string representation of the integer, or "null" if not valid
func (ni NullInt) String() string {
	if !ni.Valid {
		return "null"
	}

	return strconv.FormatInt(ni.Int, 10)
}

// NewNullInt creates a new valid NullInt with the given value
func NewNullInt(value int64) NullInt {
	return NullInt{
		Int:   value,
		Valid: true,
	}
}
