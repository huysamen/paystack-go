package enums

import (
	"encoding/json"
)

// SettlementStatus represents the status of a settlement
type SettlementStatus string

const (
	SettlementStatusSuccess SettlementStatus = "success"
	SettlementStatusPending SettlementStatus = "pending"
	SettlementStatusFailed  SettlementStatus = "failed"
)

// String returns the string representation of SettlementStatus
func (ss SettlementStatus) String() string {
	return string(ss)
}

// MarshalJSON implements json.Marshaler
func (ss SettlementStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(ss))
}

// UnmarshalJSON implements json.Unmarshaler.
//
// Paystack extends this field without notice, so an unrecognised value is
// accepted and preserved verbatim rather than rejected: failing here would
// fail the decode of the entire payload over a single field. Use IsValid to
// test whether the value is one this version of the library models.
func (ss *SettlementStatus) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	*ss = SettlementStatus(s)

	return nil
}

// IsValid returns true if the settlement status is a valid known value
func (ss SettlementStatus) IsValid() bool {
	switch ss {
	case SettlementStatusSuccess, SettlementStatusPending, SettlementStatusFailed:
		return true
	default:
		return false
	}
}

// AllSettlementStatuses returns all valid SettlementStatus values
func AllSettlementStatuses() []SettlementStatus {
	return []SettlementStatus{
		SettlementStatusSuccess,
		SettlementStatusPending,
		SettlementStatusFailed,
	}
}
