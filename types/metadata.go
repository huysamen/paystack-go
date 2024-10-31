package types

type Metadata map[string]any

func (m Metadata) SetCustomFields(fields []CustomField) {
	m["custom_fields"] = fields
}

func (m Metadata) SetCancelAction(url string) {
	m["cancel_action"] = url
}

func (m Metadata) SetCustomFilters(filters CustomFilters) {
	m["custom_filters"] = filters
}

type CustomField struct {
	DisplayName  string `json:"display_name"`
	VariableName string `json:"variable_name"`
	Value        any    `json:"value"` // todo: restrict type with generics
}

type CustomFilters struct {
	Recurring                     bool        `json:"recurring,omitempty"`
	Banks                         []string    `json:"banks,omitempty"`
	CardBrands                    []CardBrand `json:"card_brands,omitempty"`
	SupportedMobileMoneyProviders []MoMo      `json:"supported_mobile_money_providers,omitempty"`
}
