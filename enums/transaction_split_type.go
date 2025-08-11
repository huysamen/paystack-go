package enums

import (
	"encoding/json"
	"fmt"
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

// UnmarshalJSON implements json.Unmarshaler
func (tst *TransactionSplitType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	splitType := TransactionSplitType(s)
	switch splitType {
	case TransactionSplitTypePercentage, TransactionSplitTypeFlat:
		*tst = splitType
		return nil
	default:
		return fmt.Errorf("invalid TransactionSplitType value: %s", s)
	}
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
