package enums

import (
	"encoding/json"
	"fmt"
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

// UnmarshalJSON implements json.Unmarshaler
func (mas *MandateAuthorizationStatus) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	status := MandateAuthorizationStatus(s)
	switch status {
	case MandateAuthorizationStatusPending, MandateAuthorizationStatusActive, MandateAuthorizationStatusInactive:
		*mas = status
		return nil
	default:
		return fmt.Errorf("invalid MandateAuthorizationStatus value: %s", s)
	}
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
