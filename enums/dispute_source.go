package enums

import (
	"encoding/json"
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

// UnmarshalJSON implements json.Unmarshaler.
//
// Paystack extends this field without notice, so an unrecognised value is
// accepted and preserved verbatim rather than rejected: failing here would
// fail the decode of the entire payload over a single field. Use IsValid to
// test whether the value is one this version of the library models.
func (ds *DisputeSource) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	*ds = DisputeSource(s)

	return nil
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
