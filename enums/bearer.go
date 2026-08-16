package enums

import (
	"encoding/json"
)

// Bearer represents who bears the transaction charges
type Bearer string

const (
	BearerAccount    Bearer = "account"
	BearerSubaccount Bearer = "subaccount"
)

// String returns the string representation of the bearer
func (b Bearer) String() string {
	return string(b)
}

// MarshalJSON implements json.Marshaler
func (b Bearer) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(b))
}

// UnmarshalJSON implements json.Unmarshaler.
//
// Paystack extends this field without notice, so an unrecognised value is
// accepted and preserved verbatim rather than rejected: failing here would
// fail the decode of the entire payload over a single field. Use IsValid to
// test whether the value is one this version of the library models.
func (b *Bearer) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	*b = Bearer(s)

	return nil
}

// IsValid returns true if the bearer is a valid known value
func (b Bearer) IsValid() bool {
	switch b {
	case BearerAccount, BearerSubaccount:
		return true
	default:
		return false
	}
}

// AllBearers returns all valid Bearer values
func AllBearers() []Bearer {
	return []Bearer{
		BearerAccount,
		BearerSubaccount,
	}
}
