package data

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// NullFloat represents a float64 that may be null in JSON
// Accepts: floats, integers, string representations of numbers, and null
type NullFloat struct {
	Float float64
	Valid bool // true if Float is not null
}

// UnmarshalJSON implements json.Unmarshaler for NullFloat
func (nf *NullFloat) UnmarshalJSON(data []byte) error {
	// Handle null
	if data == nil || string(data) == "null" {
		nf.Float = 0.0
		nf.Valid = false

		return nil
	}

	// Try to unmarshal as float64 first
	var f float64

	if err := json.Unmarshal(data, &f); err == nil {
		nf.Float = f
		nf.Valid = true

		return nil
	}

	// Try to unmarshal as string and parse as float
	var s string

	if err := json.Unmarshal(data, &s); err == nil {
		// Handle empty string and "null" string as null
		if s == "" || s == "null" {
			nf.Float = 0.0
			nf.Valid = false

			return nil
		}

		// Try to parse string as float
		if parsed, err := strconv.ParseFloat(s, 64); err == nil {
			nf.Float = parsed
			nf.Valid = true

			return nil
		}

		return fmt.Errorf("cannot parse string %q as float", s)
	}

	return fmt.Errorf("cannot unmarshal %s into NullFloat", string(data))
}

// MarshalJSON implements json.Marshaler for NullFloat
func (nf NullFloat) MarshalJSON() ([]byte, error) {
	if !nf.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(nf.Float)
}

// ValueOr returns the float value if valid, otherwise returns the fallback value
func (nf NullFloat) ValueOr(fallback float64) float64 {
	if nf.Valid {
		return nf.Float
	}

	return fallback
}

// String returns the string representation of the float, or "null" if not valid
func (nf NullFloat) String() string {
	if !nf.Valid {
		return "null"
	}

	return strconv.FormatFloat(nf.Float, 'f', -1, 64)
}

// NewNullFloat creates a new valid NullFloat with the given value
func NewNullFloat(value float64) NullFloat {
	return NullFloat{
		Float: value,
		Valid: true,
	}
}
