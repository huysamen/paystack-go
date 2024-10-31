package types

import (
	"encoding/json"
)

type PageType int

const (
	PageTypeUnknown PageType = iota
	PageTypePayment
	PageTypeSubscription
	PageTypeProduct
	PageTypePlan
)

func (p PageType) String() string {
	switch p {
	case PageTypePayment:
		return "payment"
	case PageTypeSubscription:
		return "subscription"
	case PageTypeProduct:
		return "product"
	case PageTypePlan:
		return "plan"
	default:
		return ""
	}
}

func (p PageType) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.String())
}

func (p *PageType) UnmarshalJSON(data []byte) error {
	var str string

	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	switch str {
	case "payment":
		*p = PageTypePayment
	case "subscription":
		*p = PageTypeSubscription
	case "product":
		*p = PageTypeProduct
	case "plan":
		*p = PageTypePlan
	default:
		*p = PageTypeUnknown
	}

	return nil
}
