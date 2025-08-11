package data

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// Int represents an int64 that can be unmarshaled from various JSON types
// Null values are converted to zero instead of preserving null state
// Accepts: integers, floats (truncated), string representations of numbers, and null (→ 0)
type Int int64

// UnmarshalJSON implements json.Unmarshaler for Int
func (i *Int) UnmarshalJSON(data []byte) error {
	// Handle null - convert to zero
	if data == nil || string(data) == "null" {
		*i = 0
		return nil
	}

	// Try to unmarshal as integer first
	var intVal int64

	if err := json.Unmarshal(data, &intVal); err == nil {
		*i = Int(intVal)
		return nil
	}

	// Try to unmarshal as float64 (JSON numbers default to float64)
	var floatVal float64

	if err := json.Unmarshal(data, &floatVal); err == nil {
		*i = Int(int64(floatVal)) // Truncate to integer
		return nil
	}

	// Try to unmarshal as string and parse as integer
	var s string

	if err := json.Unmarshal(data, &s); err == nil {
		// Handle empty string and "null" string as zero
		if s == "" || s == "null" {
			*i = 0
			return nil
		}

		// Try to parse string as integer
		if parsed, err := strconv.ParseInt(s, 10, 64); err == nil {
			*i = Int(parsed)
			return nil
		}

		// Try to parse string as float and truncate to integer
		if parsed, err := strconv.ParseFloat(s, 64); err == nil {
			*i = Int(int64(parsed))
			return nil
		}

		return fmt.Errorf("cannot parse string %q as integer", s)
	}

	return fmt.Errorf("cannot unmarshal %s into Int", string(data))
}

// MarshalJSON implements json.Marshaler for Int
func (i Int) MarshalJSON() ([]byte, error) {
	return json.Marshal(int64(i))
}

// Int64 returns the int64 value
func (i Int) Int64() int64 {
	return int64(i)
}

// String returns the string representation of the integer
func (i Int) String() string {
	return strconv.FormatInt(int64(i), 10)
}

// NewInt creates a new Int with the given value
func NewInt(value int64) Int {
	return Int(value)
}
