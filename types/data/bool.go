package data

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// Bool represents a bool that can be unmarshaled from various JSON types
// Null values are converted to false instead of preserving null state
// Accepts: booleans, strings ("1", "true", "success" → true), numbers (non-zero → true), and null (→ false)
type Bool bool

// UnmarshalJSON implements json.Unmarshaler for Bool
func (b *Bool) UnmarshalJSON(data []byte) error {
	// Handle null - convert to false
	if data == nil || string(data) == "null" {
		*b = false
		return nil
	}

	// Try to unmarshal as boolean first
	var boolVal bool

	if err := json.Unmarshal(data, &boolVal); err == nil {
		*b = Bool(boolVal)
		return nil
	}

	// Try to unmarshal as string and parse as boolean
	var s string

	if err := json.Unmarshal(data, &s); err == nil {
		// Handle empty string and "null" string as false
		if s == "" || s == "null" {
			*b = false
			return nil
		}

		// Check for various string representations of true/false
		switch strings.ToLower(s) {
		case "true", "1", "success":
			*b = true
			return nil
		case "false", "0", "failure":
			*b = false
			return nil
		}

		// Try to parse as number - non-zero is true
		if parsed, err := strconv.ParseFloat(s, 64); err == nil {
			*b = Bool(parsed != 0)
			return nil
		}

		return fmt.Errorf("cannot parse string %q as boolean", s)
	}

	// Try to unmarshal as number - non-zero is true
	var floatVal float64

	if err := json.Unmarshal(data, &floatVal); err == nil {
		*b = Bool(floatVal != 0)
		return nil
	}

	// Try to unmarshal as integer - non-zero is true
	var intVal int64

	if err := json.Unmarshal(data, &intVal); err == nil {
		*b = Bool(intVal != 0)
		return nil
	}

	return fmt.Errorf("cannot unmarshal %s into Bool", string(data))
}

// MarshalJSON implements json.Marshaler for Bool
func (b Bool) MarshalJSON() ([]byte, error) {
	return json.Marshal(bool(b))
}

// Bool returns the bool value
func (b Bool) Bool() bool {
	return bool(b)
}

// String returns the string representation of the boolean
func (b Bool) String() string {
	if b {
		return "true"
	}
	return "false"
}

// NewBool creates a new Bool with the given value
func NewBool(value bool) Bool {
	return Bool(value)
}
