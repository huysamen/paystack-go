package enums

import (
	"encoding/json"
	"fmt"
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

// UnmarshalJSON implements json.Unmarshaler
func (ss *SettlementStatus) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	status := SettlementStatus(s)
	switch status {
	case SettlementStatusSuccess, SettlementStatusPending, SettlementStatusFailed:
		*ss = status
		return nil
	default:
		return fmt.Errorf("invalid SettlementStatus value: %s", s)
	}
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
