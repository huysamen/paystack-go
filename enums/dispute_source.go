package enums

import (
	"encoding/json"
	"fmt"
)

// DisputeSource represents the source of a dispute
type DisputeSource string

const (
	DisputeSourceBank DisputeSource = "bank"
	DisputeSourceCard DisputeSource = "card"
)

// String returns the string representation of DisputeSource
func (ds DisputeSource) String() string {
	return string(ds)
}

// MarshalJSON implements json.Marshaler
func (ds DisputeSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(ds))
}

// UnmarshalJSON implements json.Unmarshaler
func (ds *DisputeSource) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	source := DisputeSource(s)
	switch source {
	case DisputeSourceBank, DisputeSourceCard:
		*ds = source
		return nil
	default:
		return fmt.Errorf("invalid DisputeSource value: %s", s)
	}
}

// IsValid returns true if the dispute source is a valid known value
func (ds DisputeSource) IsValid() bool {
	switch ds {
	case DisputeSourceBank, DisputeSourceCard:
		return true
	default:
		return false
	}
}

// AllDisputeSources returns all valid DisputeSource values
func AllDisputeSources() []DisputeSource {
	return []DisputeSource{
		DisputeSourceBank,
		DisputeSourceCard,
	}
}
