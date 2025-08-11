package types

import (
	"encoding/json"
	"fmt"
)

// NullMetadata represents arbitrary key-value pairs with nullability
type Metadata struct {
	Metadata map[string]any
	Valid    bool
}

// UnmarshalJSON implements json.Unmarshaler for Metadata
func (m *Metadata) UnmarshalJSON(data []byte) error {
	m.Metadata = make(map[string]any)

	if data == nil || string(data) == "null" {
		m.Valid = false
		return nil
	}

	// Handle empty string as invalid/null metadata
	if string(data) == `""` {
		m.Valid = false
		return nil
	}

	// Try to unmarshal as object first
	var obj map[string]any

	if err := json.Unmarshal(data, &obj); err == nil {
		m.Metadata = obj
		m.Valid = true
		return nil
	}

	// If that fails, try to unmarshal as string and then parse the string as JSON
	var str string
	if err := json.Unmarshal(data, &str); err == nil {
		// Handle empty string or literal "null" case
		if str == "" || str == "null" {
			m.Valid = false
			return nil
		}

		// Try to parse the string as JSON object
		var objFromStr map[string]any
		if err := json.Unmarshal([]byte(str), &objFromStr); err == nil {
			m.Metadata = objFromStr
			m.Valid = true
			return nil
		}

		// If the string contains valid JSON but not an object (e.g., number/bool/array), mark invalid
		var anyFromStr any
		if err := json.Unmarshal([]byte(str), &anyFromStr); err == nil {
			m.Valid = false
			return nil
		}

		// Not parseable JSON; treat as invalid but do not error
		m.Valid = false
		return nil
	}

	// Finally, if the value is a non-object JSON type (number, bool, array), tolerate by marking invalid
	var anyVal any
	if err := json.Unmarshal(data, &anyVal); err == nil {
		m.Valid = false
		return nil
	}

	// If still not parseable, return error to surface malformed JSON
	m.Valid = false
	return fmt.Errorf("cannot unmarshal %s into Metadata", string(data))
}

// MarshalJSON implements json.Marshaler for Metadata
func (m Metadata) MarshalJSON() ([]byte, error) {
	if !m.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(m.Metadata)
}

// IsEmpty returns true if metadata is nil or empty
func (m Metadata) IsEmpty() bool {
	return len(m.Metadata) == 0
}

// NewMetadata creates a valid Metadata from a map
func NewMetadata(md map[string]any) Metadata {
	return Metadata{Metadata: md, Valid: true}
}
