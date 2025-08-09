package data

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// Float represents a float64 that can be unmarshaled from various JSON types
// Null values are converted to zero instead of preserving null state
// Accepts: floats, integers, string representations of numbers, and null (→ 0.0)
type Float float64

// UnmarshalJSON implements json.Unmarshaler for Float
func (f *Float) UnmarshalJSON(data []byte) error {
	// Handle null - convert to zero
	if data == nil || string(data) == "null" {
		// TODO: Remove this debug print statement after testing
		fmt.Printf("[DEBUG] Float.UnmarshalJSON: received null value, converting to 0.0\n")
		*f = 0
		return nil
	}

	// Try to unmarshal as float64 first
	var floatVal float64

	if err := json.Unmarshal(data, &floatVal); err == nil {
		*f = Float(floatVal)
		return nil
	}

	// Try to unmarshal as integer (JSON numbers can be int)
	var intVal int64

	if err := json.Unmarshal(data, &intVal); err == nil {
		*f = Float(float64(intVal))
		return nil
	}

	// Try to unmarshal as string and parse as float
	var s string

	if err := json.Unmarshal(data, &s); err == nil {
		// Handle empty string and "null" string as zero
		if s == "" || s == "null" {
			// TODO: Remove this debug print statement after testing
			fmt.Printf("[DEBUG] Float.UnmarshalJSON: received null/empty string value, converting to 0.0\n")
			*f = 0
			return nil
		}

		// Try to parse string as float
		if parsed, err := strconv.ParseFloat(s, 64); err == nil {
			*f = Float(parsed)
			return nil
		}

		return fmt.Errorf("cannot parse string %q as float", s)
	}

	return fmt.Errorf("cannot unmarshal %s into Float", string(data))
}

// MarshalJSON implements json.Marshaler for Float
func (f Float) MarshalJSON() ([]byte, error) {
	return json.Marshal(float64(f))
}

// Float64 returns the float64 value
func (f Float) Float64() float64 {
	return float64(f)
}

// String returns the string representation of the float
func (f Float) String() string {
	return strconv.FormatFloat(float64(f), 'g', -1, 64)
}

// NewFloat creates a new Float with the given value
func NewFloat(value float64) Float {
	return Float(value)
}
