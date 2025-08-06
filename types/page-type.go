package types

// PageType represents the type of payment page
type PageType string

const (
	PageTypeUnknown      PageType = ""
	PageTypePayment      PageType = "payment"
	PageTypeSubscription PageType = "subscription"
	PageTypeProduct      PageType = "product"
	PageTypePlan         PageType = "plan"
)

// String returns the string representation of the page type
func (p PageType) String() string {
	return string(p)
}

// IsValid returns true if the page type is a valid known value
func (p PageType) IsValid() bool {
	switch p {
	case PageTypePayment, PageTypeSubscription, PageTypeProduct, PageTypePlan:
		return true
	default:
		return false
	}
}
