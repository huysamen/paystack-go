package enums

import (
	"encoding/json"
)

// MandateAuthorizationStatus represents the status of mandate authorization
type MandateAuthorizationStatus string

const (
	MandateAuthorizationStatusPending  MandateAuthorizationStatus = "pending"
	MandateAuthorizationStatusActive   MandateAuthorizationStatus = "active"
	MandateAuthorizationStatusInactive MandateAuthorizationStatus = "inactive"
)

// String returns the string representation of MandateAuthorizationStatus
func (mas MandateAuthorizationStatus) String() string {
	return string(mas)
}

// MarshalJSON implements json.Marshaler
func (mas MandateAuthorizationStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(mas))
}

// UnmarshalJSON implements json.Unmarshaler.
//
// Paystack extends this field without notice, so an unrecognised value is
// accepted and preserved verbatim rather than rejected: failing here would
// fail the decode of the entire payload over a single field. Use IsValid to
// test whether the value is one this version of the library models.
func (mas *MandateAuthorizationStatus) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	*mas = MandateAuthorizationStatus(s)

	return nil
}

// IsValid returns true if the mandate authorization status is a valid known value
func (mas MandateAuthorizationStatus) IsValid() bool {
	switch mas {
	case MandateAuthorizationStatusPending, MandateAuthorizationStatusActive, MandateAuthorizationStatusInactive:
		return true
	default:
		return false
	}
}

// AllMandateAuthorizationStatuses returns all valid MandateAuthorizationStatus values
func AllMandateAuthorizationStatuses() []MandateAuthorizationStatus {
	return []MandateAuthorizationStatus{
		MandateAuthorizationStatusPending,
		MandateAuthorizationStatusActive,
		MandateAuthorizationStatusInactive,
	}
}
