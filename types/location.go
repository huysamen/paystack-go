package types

// Country represents a country supported by Paystack
type Country struct {
	ID                  int                  `json:"id"`
	Name                string               `json:"name"`
	ISOCode             string               `json:"iso_code"`
	DefaultCurrencyCode string               `json:"default_currency_code"`
	IntegrationDefaults *Metadata            `json:"integration_defaults"`
	Relationships       CountryRelationships `json:"relationships"`
}

// CountryRelationships represents the relationships for a country
type CountryRelationships struct {
	Currency           CountryRelationshipData `json:"currency"`
	IntegrationFeature CountryRelationshipData `json:"integration_feature"`
	IntegrationType    CountryRelationshipData `json:"integration_type"`
}

// CountryRelationshipData represents relationship data
type CountryRelationshipData struct {
	Type string `json:"type"`
	Data []any  `json:"data"`
}

// State represents a state for address verification
type State struct {
	Name         string `json:"name"`
	Slug         string `json:"slug"`
	Abbreviation string `json:"abbreviation"`
}
