package enums

import (
	"encoding/json"
)

// TransactionSplitBearerType represents who bears the transaction split charges
type TransactionSplitBearerType string

const (
	TransactionSplitBearerTypeAccount    TransactionSplitBearerType = "account"
	TransactionSplitBearerTypeSubaccount TransactionSplitBearerType = "subaccount"
	TransactionSplitBearerTypeAll        TransactionSplitBearerType = "all"
)

// String returns the string representation of TransactionSplitBearerType
func (tsbt TransactionSplitBearerType) String() string {
	return string(tsbt)
}

// MarshalJSON implements json.Marshaler
func (tsbt TransactionSplitBearerType) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(tsbt))
}

// UnmarshalJSON implements json.Unmarshaler.
//
// Paystack extends this field without notice, so an unrecognised value is
// accepted and preserved verbatim rather than rejected: failing here would
// fail the decode of the entire payload over a single field. Use IsValid to
// test whether the value is one this version of the library models.
func (tsbt *TransactionSplitBearerType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	*tsbt = TransactionSplitBearerType(s)

	return nil
}

// IsValid returns true if the transaction split bearer type is a valid known value
func (tsbt TransactionSplitBearerType) IsValid() bool {
	switch tsbt {
	case TransactionSplitBearerTypeAccount, TransactionSplitBearerTypeSubaccount, TransactionSplitBearerTypeAll:
		return true
	default:
		return false
	}
}

// AllTransactionSplitBearerTypes returns all valid TransactionSplitBearerType values
func AllTransactionSplitBearerTypes() []TransactionSplitBearerType {
	return []TransactionSplitBearerType{
		TransactionSplitBearerTypeAccount,
		TransactionSplitBearerTypeSubaccount,
		TransactionSplitBearerTypeAll,
	}
}
