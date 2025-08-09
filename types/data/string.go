package data

import (
	"encoding/json"
	"fmt"
)

// String represents a string that can be unmarshaled from various JSON types
// Null values are converted to empty string instead of preserving null state
// Accepts: strings, numbers (converted to strings), and null (→ "")
type String string

// UnmarshalJSON implements json.Unmarshaler for String
func (s *String) UnmarshalJSON(data []byte) error {
	// Handle null - convert to empty string
	if data == nil || string(data) == "null" {
		// TODO: Remove this debug print statement after testing
		fmt.Printf("[DEBUG] String.UnmarshalJSON: received null value, converting to empty string\n")
		*s = ""
		return nil
	}

	// Try to unmarshal as string first
	var strVal string

	if err := json.Unmarshal(data, &strVal); err == nil {
		// Handle "null" string as empty string
		if strVal == "null" {
			// TODO: Remove this debug print statement after testing
			fmt.Printf("[DEBUG] String.UnmarshalJSON: received 'null' string value, converting to empty string\n")
			*s = ""
			return nil
		}

		*s = String(strVal)
		return nil
	}

	// Try to unmarshal as number and convert to string
	var n float64

	if err := json.Unmarshal(data, &n); err == nil {
		*s = String(fmt.Sprintf("%g", n)) // Use %g for clean number representation
		return nil
	}

	return fmt.Errorf("cannot unmarshal %s into String", string(data))
}

// MarshalJSON implements json.Marshaler for String
func (s String) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(s))
}

// String returns the string value (implements fmt.Stringer)
func (s String) String() string {
	return string(s)
}

// NewString creates a new String with the given value
func NewString(value string) String {
	return String(value)
}
