package enums

import (
	"encoding/json"
)

// TransactionSplitType represents the type of transaction split
type TransactionSplitType string

const (
	TransactionSplitTypePercentage TransactionSplitType = "percentage"
	TransactionSplitTypeFlat       TransactionSplitType = "flat"
)

// String returns the string representation of TransactionSplitType
func (tst TransactionSplitType) String() string {
	return string(tst)
}

// MarshalJSON implements json.Marshaler
func (tst TransactionSplitType) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(tst))
}

// UnmarshalJSON implements json.Unmarshaler.
//
// Paystack extends this field without notice, so an unrecognised value is
// accepted and preserved verbatim rather than rejected: failing here would
// fail the decode of the entire payload over a single field. Use IsValid to
// test whether the value is one this version of the library models.
func (tst *TransactionSplitType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	*tst = TransactionSplitType(s)

	return nil
}

// IsValid returns true if the transaction split type is a valid known value
func (tst TransactionSplitType) IsValid() bool {
	switch tst {
	case TransactionSplitTypePercentage, TransactionSplitTypeFlat:
		return true
	default:
		return false
	}
}

// AllTransactionSplitTypes returns all valid TransactionSplitType values
func AllTransactionSplitTypes() []TransactionSplitType {
	return []TransactionSplitType{
		TransactionSplitTypePercentage,
		TransactionSplitTypeFlat,
	}
}
