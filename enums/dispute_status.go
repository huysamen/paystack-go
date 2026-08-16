package enums

import (
	"encoding/json"
)

// DisputeStatus represents the status of a dispute
type DisputeStatus string

const (
	DisputeStatusAwaitingMerchantFeedback DisputeStatus = "awaiting-merchant-feedback"
	DisputeStatusAwaitingBankFeedback     DisputeStatus = "awaiting-bank-feedback"
	DisputeStatusPending                  DisputeStatus = "pending"
	DisputeStatusResolved                 DisputeStatus = "resolved"
	DisputeStatusArchived                 DisputeStatus = "archived"
)

// String returns the string representation of DisputeStatus
func (ds DisputeStatus) String() string {
	return string(ds)
}

// MarshalJSON implements json.Marshaler
func (ds DisputeStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(ds))
}

// UnmarshalJSON implements json.Unmarshaler.
//
// Paystack extends this field without notice, so an unrecognised value is
// accepted and preserved verbatim rather than rejected: failing here would
// fail the decode of the entire payload over a single field. Use IsValid to
// test whether the value is one this version of the library models.
func (ds *DisputeStatus) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	*ds = DisputeStatus(s)

	return nil
}

// IsValid returns true if the dispute status is a valid known value
func (ds DisputeStatus) IsValid() bool {
	switch ds {
	case DisputeStatusAwaitingMerchantFeedback, DisputeStatusAwaitingBankFeedback,
		DisputeStatusPending, DisputeStatusResolved, DisputeStatusArchived:
		return true
	default:
		return false
	}
}

// AllDisputeStatuses returns all valid DisputeStatus values
func AllDisputeStatuses() []DisputeStatus {
	return []DisputeStatus{
		DisputeStatusAwaitingMerchantFeedback,
		DisputeStatusAwaitingBankFeedback,
		DisputeStatusPending,
		DisputeStatusResolved,
		DisputeStatusArchived,
	}
}
