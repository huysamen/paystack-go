package types

import (
	"encoding/json"
)

type CardBrand int

const (
	CardBrandUnknown CardBrand = iota
	CardBrandVisa
	CardBrandMasterCard
	CardBrandVerve
)

func (c CardBrand) String() string {
	switch c {
	case CardBrandVisa:
		return "visa"
	case CardBrandMasterCard:
		return "mastercard"
	case CardBrandVerve:
		return "verve"
	default:
		return ""
	}
}

func (c CardBrand) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.String())
}

func (c *CardBrand) UnmarshalJSON(data []byte) error {
	var str string

	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	switch str {
	case "visa":
		*c = CardBrandVisa
	case "mastercard":
		*c = CardBrandMasterCard
	case "verve":
		*c = CardBrandVerve
	default:
		*c = CardBrandUnknown
	}

	return nil
}
