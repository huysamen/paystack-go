package data

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// MultiString handles both string and number JSON values
type MultiString string

// UnmarshalJSON implements json.Unmarshaler for MultiString
func (ms *MultiString) UnmarshalJSON(data []byte) error {
	// Handle null
	if string(data) == "null" {
		*ms = ""
		return nil
	}

	// Try to unmarshal as string first
	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		*ms = MultiString(s)
		return nil
	}

	// Try to unmarshal as number
	var n float64
	if err := json.Unmarshal(data, &n); err == nil {
		*ms = MultiString(strconv.FormatFloat(n, 'f', 0, 64))
		return nil
	}

	return fmt.Errorf("cannot unmarshal %s into MultiString", string(data))
}

// String returns the string representation
func (ms MultiString) String() string {
	return string(ms)
}

// MarshalJSON implements json.Marshaler
func (ms MultiString) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(ms))
}
