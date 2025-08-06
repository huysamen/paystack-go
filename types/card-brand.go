package types

// CardBrand represents credit/debit card brands supported by Paystack
type CardBrand string

const (
	CardBrandUnknown    CardBrand = ""
	CardBrandVisa       CardBrand = "visa"
	CardBrandMasterCard CardBrand = "mastercard"
	CardBrandVerve      CardBrand = "verve"
)

// String returns the string representation of the card brand
func (c CardBrand) String() string {
	return string(c)
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
