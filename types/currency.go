package types

import (
	"encoding/json"
)

type Currency int

const (
	CurrencyUnknown Currency = iota
	CurrencyZAR
	CurrencyNGN
	CurrencyUSD
	CurrencyGHS
	CurrencyKES
)

func (c Currency) String() string {
	switch c {
	case CurrencyZAR:
		return "ZAR"
	case CurrencyNGN:
		return "NGN"
	case CurrencyUSD:
		return "USD"
	case CurrencyGHS:
		return "GHS"
	case CurrencyKES:
		return "KES"
	default:
		return ""
	}
}

func (c Currency) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.String())
}

func (c *Currency) UnmarshalJSON(data []byte) error {
	var str string

	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	switch str {
	case "ZAR":
		*c = CurrencyZAR
	case "NGN":
		*c = CurrencyNGN
	case "USD":
		*c = CurrencyUSD
	case "GHS":
		*c = CurrencyGHS
	case "KES":
		*c = CurrencyKES
	default:
		*c = CurrencyUnknown
	}

	return nil
}
