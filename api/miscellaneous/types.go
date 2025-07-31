package miscellaneous

import (
	"time"

	"github.com/huysamen/paystack-go/types"
)

// Bank represents a bank in the system
type Bank struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Slug        string    `json:"slug"`
	Code        string    `json:"code"`
	LongCode    string    `json:"longcode"`
	Gateway     *string   `json:"gateway"`
	PayWithBank bool      `json:"pay_with_bank"`
	Active      bool      `json:"active"`
	IsDeleted   bool      `json:"is_deleted"`
	Country     string    `json:"country"`
	Currency    string    `json:"currency"`
	Type        string    `json:"type"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// Country represents a country supported by Paystack
type Country struct {
	ID                  int                  `json:"id"`
	Name                string               `json:"name"`
	ISOCode             string               `json:"iso_code"`
	DefaultCurrencyCode string               `json:"default_currency_code"`
	IntegrationDefaults map[string]any       `json:"integration_defaults"`
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

// Banks

// BankListRequest represents the request to list banks
type BankListRequest struct {
	Country                *string `json:"country,omitempty"`                  // Optional: country filter (ghana, kenya, nigeria, south africa)
	UseCursor              *bool   `json:"use_cursor,omitempty"`               // Optional: enable cursor pagination
	PerPage                *int    `json:"perPage,omitempty"`                  // Optional: records per page (default: 50, max: 100)
	PayWithBankTransfer    *bool   `json:"pay_with_bank_transfer,omitempty"`   // Optional: filter for transfer payment banks
	PayWithBank            *bool   `json:"pay_with_bank,omitempty"`            // Optional: filter for direct payment banks
	EnabledForVerification *bool   `json:"enabled_for_verification,omitempty"` // Optional: filter for verification-supported banks
	Next                   *string `json:"next,omitempty"`                     // Optional: cursor for next page
	Previous               *string `json:"previous,omitempty"`                 // Optional: cursor for previous page
	Gateway                *string `json:"gateway,omitempty"`                  // Optional: gateway type filter
	Type                   *string `json:"type,omitempty"`                     // Optional: financial channel type
	Currency               *string `json:"currency,omitempty"`                 // Optional: currency filter
	IncludeNIPSortCode     *bool   `json:"include_nip_sort_code,omitempty"`    // Optional: include NIP institution codes
}

// BankListResponse represents the response from listing banks
type BankListResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    []Bank      `json:"data"`
	Meta    *types.Meta `json:"meta,omitempty"`
}

// Countries

// CountryListResponse represents the response from listing countries
type CountryListResponse struct {
	Status  bool      `json:"status"`
	Message string    `json:"message"`
	Data    []Country `json:"data"`
}

// States

// StateListRequest represents the request to list states
type StateListRequest struct {
	Country string `json:"country"` // Required: country code
}

// StateListResponse represents the response from listing states
type StateListResponse struct {
	Status  bool    `json:"status"`
	Message string  `json:"message"`
	Data    []State `json:"data"`
}
