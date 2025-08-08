package data

import (
	"encoding/json"
)

// MultiBool handles boolean fields that may come as strings from the API
// Accepts: true, "true", "success" as true; everything else as false
type MultiBool bool

// UnmarshalJSON implements json.Unmarshaler for MultiBool
func (mb *MultiBool) UnmarshalJSON(data []byte) error {
	// Handle null
	if string(data) == "null" {
		*mb = false
		return nil
	}

	// Try to unmarshal as boolean first
	var b bool
	if err := json.Unmarshal(data, &b); err == nil {
		*mb = MultiBool(b)
		return nil
	}

	// Try to unmarshal as string
	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		// Consider "true" and "success" as true values
		*mb = MultiBool(s == "true" || s == "success")
		return nil
	}

	// Default to false for any other type or parsing error
	*mb = false
	return nil
}

// MarshalJSON implements json.Marshaler for MultiBool
func (mb MultiBool) MarshalJSON() ([]byte, error) {
	return json.Marshal(bool(mb))
}

// Bool returns the boolean value
func (mb MultiBool) Bool() bool {
	return bool(mb)
}
