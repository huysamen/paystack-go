package enums

import (
	"encoding/json"
)

// PageType represents the type of payment page
type PageType string

const (
	PageTypeInvoice          PageType = "invoice"
	PageTypePaymentRequest   PageType = "payment_request"
	PageTypePaymentPage      PageType = "payment_page"
	PageTypeProductPageSetup PageType = "product_page_setup"
)

// String returns the string representation of PageType
func (pt PageType) String() string {
	return string(pt)
}

// MarshalJSON implements json.Marshaler
func (pt PageType) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(pt))
}

// UnmarshalJSON implements json.Unmarshaler.
//
// Paystack extends this field without notice, so an unrecognised value is
// accepted and preserved verbatim rather than rejected: failing here would
// fail the decode of the entire payload over a single field. Use IsValid to
// test whether the value is one this version of the library models.
func (pt *PageType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	*pt = PageType(s)

	return nil
}

// IsValid returns true if the page type is a valid known value
func (pt PageType) IsValid() bool {
	switch pt {
	case PageTypeInvoice, PageTypePaymentRequest, PageTypePaymentPage, PageTypeProductPageSetup:
		return true
	default:
		return false
	}
}

// AllPageTypes returns all valid PageType values
func AllPageTypes() []PageType {
	return []PageType{
		PageTypeInvoice,
		PageTypePaymentRequest,
		PageTypePaymentPage,
		PageTypeProductPageSetup,
	}
}
