package types

import (
	"encoding/json"
)

// NullMetadata represents arbitrary key-value pairs with nullability
type Metadata struct {
	Metadata map[string]any
	Valid    bool
}

// UnmarshalJSON implements json.Unmarshaler for Metadata
func (m Metadata) UnmarshalJSON(data []byte) error {
	m.Metadata = make(map[string]any)

	if data == nil || string(data) == "null" {
		m.Valid = false

		return nil
	}

	// Try to unmarshal as object
	var obj map[string]any

	if err := json.Unmarshal(data, &obj); err != nil {
		// If not a valid object, fail
		m.Valid = false
		return err
	}

	m.Metadata = obj
	m.Valid = true
	return nil
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
