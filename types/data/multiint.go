package data

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// MultiInt handles integer fields that may come as strings or numbers from the API
// Accepts: integers, floats (truncated), and string representations of numbers
type MultiInt int

// UnmarshalJSON implements json.Unmarshaler for MultiInt
func (mi *MultiInt) UnmarshalJSON(data []byte) error {
	// Handle null
	if string(data) == "null" {
		*mi = 0
		return nil
	}

	// Try to unmarshal as integer first
	var i int
	if err := json.Unmarshal(data, &i); err == nil {
		*mi = MultiInt(i)
		return nil
	}

	// Try to unmarshal as float64 (JSON numbers default to float64)
	var f float64
	if err := json.Unmarshal(data, &f); err == nil {
		*mi = MultiInt(int(f)) // Truncate to integer
		return nil
	}

	// Try to unmarshal as string and parse as integer
	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		// Handle empty string as zero
		if s == "" {
			*mi = 0
			return nil
		}

		// Try to parse string as integer
		if parsed, err := strconv.Atoi(s); err == nil {
			*mi = MultiInt(parsed)
			return nil
		}

		// Try to parse string as float and truncate to integer
		if parsed, err := strconv.ParseFloat(s, 64); err == nil {
			*mi = MultiInt(int(parsed))
			return nil
		}

		return fmt.Errorf("cannot parse string %q as integer", s)
	}

	return fmt.Errorf("cannot unmarshal %s into MultiInt", string(data))
}

// MarshalJSON implements json.Marshaler for MultiInt
func (mi MultiInt) MarshalJSON() ([]byte, error) {
	return json.Marshal(int(mi))
}

// Int returns the integer value
func (mi MultiInt) Int() int {
	return int(mi)
}

// String returns the string representation
func (mi MultiInt) String() string {
	return strconv.Itoa(int(mi))
}
