package types

import "github.com/huysamen/paystack-go/types/data"

// Country represents a country supported by Paystack
type Country struct {
	ID                  data.Int             `json:"id"`
	Name                data.String          `json:"name"`
	ISOCode             data.String          `json:"iso_code"`
	DefaultCurrencyCode data.String          `json:"default_currency_code"`
	IntegrationDefaults Metadata             `json:"integration_defaults"`
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
	Type data.String `json:"type"`
	Data []any       `json:"data"`
}

// State represents a state for address verification
type State struct {
	Name         data.String `json:"name"`
	Slug         data.String `json:"slug"`
	Abbreviation data.String `json:"abbreviation"`
}
