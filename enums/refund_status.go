package enums

import (
	"encoding/json"
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

// UnmarshalJSON implements json.Unmarshaler.
//
// Paystack extends this field without notice, so an unrecognised value is
// accepted and preserved verbatim rather than rejected: failing here would
// fail the decode of the entire payload over a single field. Use IsValid to
// test whether the value is one this version of the library models.
func (rs *RefundStatus) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	*rs = RefundStatus(s)

	return nil
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
