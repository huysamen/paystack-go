package types

// Currency represents currencies supported by Paystack
type Currency string

const (
	CurrencyUnknown Currency = ""
	CurrencyZAR     Currency = "ZAR"
	CurrencyNGN     Currency = "NGN"
	CurrencyUSD     Currency = "USD"
	CurrencyGHS     Currency = "GHS"
	CurrencyKES     Currency = "KES"
)

// String returns the string representation of the currency
func (c Currency) String() string {
	return string(c)
}

// IsValid returns true if the currency is a valid known value
func (c Currency) IsValid() bool {
	switch c {
	case CurrencyZAR, CurrencyNGN, CurrencyUSD, CurrencyGHS, CurrencyKES:
		return true
	default:
		return false
	}
}
