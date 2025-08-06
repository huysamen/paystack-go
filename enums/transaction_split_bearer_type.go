package enums

import (
	"encoding/json"
	"fmt"
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

// UnmarshalJSON implements json.Unmarshaler
func (tsbt *TransactionSplitBearerType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	bearerType := TransactionSplitBearerType(s)
	switch bearerType {
	case TransactionSplitBearerTypeAccount, TransactionSplitBearerTypeSubaccount, TransactionSplitBearerTypeAll:
		*tsbt = bearerType
		return nil
	default:
		return fmt.Errorf("invalid TransactionSplitBearerType value: %s", s)
	}
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
