package enums

import (
	"encoding/json"
	"fmt"
)

// RefundStatus represents the status of a refund
type RefundStatus string

const (
	RefundStatusPending   RefundStatus = "pending"
	RefundStatusProcessed RefundStatus = "processed"
	RefundStatusFailed    RefundStatus = "failed"
)

// String returns the string representation of RefundStatus
func (rs RefundStatus) String() string {
	return string(rs)
}

// MarshalJSON implements json.Marshaler
func (rs RefundStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(rs))
}

// UnmarshalJSON implements json.Unmarshaler
func (rs *RefundStatus) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	status := RefundStatus(s)
	switch status {
	case RefundStatusPending, RefundStatusProcessed, RefundStatusFailed:
		*rs = status
		return nil
	default:
		return fmt.Errorf("invalid RefundStatus value: %s", s)
	}
}

// IsValid returns true if the refund status is a valid known value
func (rs RefundStatus) IsValid() bool {
	switch rs {
	case RefundStatusPending, RefundStatusProcessed, RefundStatusFailed:
		return true
	default:
		return false
	}
}

// AllRefundStatuses returns all valid RefundStatus values
func AllRefundStatuses() []RefundStatus {
	return []RefundStatus{
		RefundStatusPending,
		RefundStatusProcessed,
		RefundStatusFailed,
	}
}
