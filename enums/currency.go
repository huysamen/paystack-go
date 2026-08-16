package enums

import (
	"encoding/json"
)

// Currency represents currencies supported by Paystack
type Currency string

const (
	CurrencyZAR Currency = "ZAR"
	CurrencyNGN Currency = "NGN"
	CurrencyUSD Currency = "USD"
	CurrencyGHS Currency = "GHS"
	CurrencyKES Currency = "KES"
)

// String returns the string representation of the currency
func (c Currency) String() string {
	return string(c)
}

// MarshalJSON implements json.Marshaler
func (c Currency) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(c))
}

// UnmarshalJSON implements json.Unmarshaler.
//
// Paystack extends this field without notice, so an unrecognised value is
// accepted and preserved verbatim rather than rejected: failing here would
// fail the decode of the entire payload over a single field. Use IsValid to
// test whether the value is one this version of the library models.
func (c *Currency) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	*c = Currency(s)

	return nil
}

// IsValid returns true if the currency is a valid known value
func (c Currency) IsValid() bool {
	switch c {
	case CurrencyZAR, CurrencyNGN, CurrencyUSD, CurrencyGHS, CurrencyKES, "":
		return true
	default:
		return false
	}
}

// AllCurrencies returns all valid Currency values
func AllCurrencies() []Currency {
	return []Currency{
		CurrencyZAR,
		CurrencyNGN,
		CurrencyUSD,
		CurrencyGHS,
		CurrencyKES,
	}
}
