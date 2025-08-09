package data

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// Uint represents a uint64 that converts null JSON values to zero
// This is the non-null variant that treats null as 0
// For preserving null values, use NullUint instead
type Uint uint64

// UnmarshalJSON implements json.Unmarshaler for Uint
func (u *Uint) UnmarshalJSON(data []byte) error {
	// Handle null
	if data == nil || string(data) == "null" {
		fmt.Printf("[DEBUG] Uint.UnmarshalJSON: received null value, converting to zero\n") // TODO: remove debug print
		*u = 0
		return nil
	}

	// Try to unmarshal as uint64 first
	var ui uint64
	if err := json.Unmarshal(data, &ui); err == nil {
		*u = Uint(ui)
		return nil
	}

	// Try to unmarshal as float64 (JSON numbers default to float64)
	var f float64
	if err := json.Unmarshal(data, &f); err == nil {
		if f < 0 {
			return fmt.Errorf("cannot convert negative number %v to uint", f)
		}
		*u = Uint(uint64(f)) // Truncate to uint64
		return nil
	}

	// Try to unmarshal as string and parse as uint
	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		// Handle empty string or "null" as zero
		if s == "" || s == "null" {
			fmt.Printf("[DEBUG] Uint.UnmarshalJSON: received null/empty string value, converting to zero\n") // TODO: remove debug print
			*u = 0
			return nil
		}

		// Try to parse string as uint64
		if parsed, err := strconv.ParseUint(s, 10, 64); err == nil {
			*u = Uint(parsed)
			return nil
		}

		// Try to parse string as float and truncate to uint64
		if parsed, err := strconv.ParseFloat(s, 64); err == nil {
			if parsed < 0 {
				return fmt.Errorf("cannot convert negative number %v to uint", parsed)
			}
			*u = Uint(uint64(parsed))
			return nil
		}

		return fmt.Errorf("cannot parse string %q as uint", s)
	}

	return fmt.Errorf("cannot unmarshal %s into Uint", string(data))
}

// MarshalJSON implements json.Marshaler for Uint
func (u Uint) MarshalJSON() ([]byte, error) {
	return json.Marshal(uint64(u))
}

// Uint64 returns the uint64 value
func (u Uint) Uint64() uint64 {
	return uint64(u)
}

// String returns the string representation of the uint
func (u Uint) String() string {
	return strconv.FormatUint(uint64(u), 10)
}

// NewUint creates a new Uint with the given value
func NewUint(value uint64) Uint {
	return Uint(value)
}
