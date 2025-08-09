package data

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// NullUint represents a uint64 that may be null in JSON
// Unlike Uint, this preserves null values instead of converting them to zero
// Accepts: unsigned integers, floats (truncated), string representations of numbers, and null
type NullUint struct {
	Uint  uint64
	Valid bool // true if Uint is not null
}

// UnmarshalJSON implements json.Unmarshaler for NullUint
func (nu *NullUint) UnmarshalJSON(data []byte) error {
	// Handle null
	if data == nil || string(data) == "null" {
		nu.Uint = 0
		nu.Valid = false
		return nil
	}

	// Try to unmarshal as uint64 first
	var ui uint64
	if err := json.Unmarshal(data, &ui); err == nil {
		nu.Uint = ui
		nu.Valid = true
		return nil
	}

	// Try to unmarshal as float64 (JSON numbers default to float64)
	var f float64
	if err := json.Unmarshal(data, &f); err == nil {
		if f < 0 {
			return fmt.Errorf("cannot convert negative number %v to uint", f)
		}
		nu.Uint = uint64(f) // Truncate to uint64
		nu.Valid = true
		return nil
	}

	// Try to unmarshal as string and parse as uint
	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		// Handle empty string as null
		if s == "" || s == "null" {
			nu.Uint = 0
			nu.Valid = false
			return nil
		}

		// Try to parse string as uint64
		if parsed, err := strconv.ParseUint(s, 10, 64); err == nil {
			nu.Uint = parsed
			nu.Valid = true
			return nil
		}

		// Try to parse string as float and truncate to uint64
		if parsed, err := strconv.ParseFloat(s, 64); err == nil {
			if parsed < 0 {
				return fmt.Errorf("cannot convert negative number %v to uint", parsed)
			}
			nu.Uint = uint64(parsed)
			nu.Valid = true
			return nil
		}

		return fmt.Errorf("cannot parse string %q as uint", s)
	}

	return fmt.Errorf("cannot unmarshal %s into NullUint", string(data))
}

// MarshalJSON implements json.Marshaler for NullUint
func (nu NullUint) MarshalJSON() ([]byte, error) {
	if !nu.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(nu.Uint)
}

// ValueOr returns the uint value if valid, otherwise returns the fallback value
func (nu NullUint) ValueOr(fallback uint64) uint64 {
	if nu.Valid {
		return nu.Uint
	}
	return fallback
}

// String returns the string representation of the uint, or "null" if not valid
func (nu NullUint) String() string {
	if !nu.Valid {
		return "null"
	}
	return strconv.FormatUint(nu.Uint, 10)
}

// NewNullUint creates a new valid NullUint with the given value
func NewNullUint(value uint64) NullUint {
	return NullUint{
		Uint:  value,
		Valid: true,
	}
}
