package enums

import (
	"encoding/json"
	"fmt"
)

// CardBrand represents credit/debit card brands supported by Paystack
type CardBrand string

const (
	CardBrandVisa       CardBrand = "visa"
	CardBrandMasterCard CardBrand = "mastercard"
	CardBrandVerve      CardBrand = "verve"
)

// String returns the string representation of the card brand
func (c CardBrand) String() string {
	return string(c)
}

// MarshalJSON implements json.Marshaler
func (c CardBrand) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(c))
}

// UnmarshalJSON implements json.Unmarshaler
func (c *CardBrand) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	brand := CardBrand(s)
	switch brand {
	case CardBrandVisa, CardBrandMasterCard, CardBrandVerve:
		*c = brand
		return nil
	default:
		return fmt.Errorf("invalid CardBrand value: %s", s)
	}
}

// IsValid returns true if the card brand is a valid known value
func (c CardBrand) IsValid() bool {
	switch c {
	case CardBrandVisa, CardBrandMasterCard, CardBrandVerve:
		return true
	default:
		return false
	}
}

// AllCardBrands returns all valid CardBrand values
func AllCardBrands() []CardBrand {
	return []CardBrand{
		CardBrandVisa,
		CardBrandMasterCard,
		CardBrandVerve,
	}
}
