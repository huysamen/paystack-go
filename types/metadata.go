package types

import (
	"encoding/json"
	"fmt"

	"github.com/huysamen/paystack-go/enums"
)

// Metadata represents arbitrary metadata that can be attached to Paystack objects
type Metadata map[string]any

// NewMetadata creates a new Metadata instance
func NewMetadata() Metadata {
	return make(Metadata)
}

// Set adds or updates a metadata field
func (m Metadata) Set(key string, value any) Metadata {
	if m == nil {
		m = make(Metadata)
	}
	m[key] = value
	return m
}

// Get retrieves a metadata field
func (m Metadata) Get(key string) (any, bool) {
	if m == nil {
		return nil, false
	}
	value, exists := m[key]
	return value, exists
}

// GetString retrieves a metadata field as a string
func (m Metadata) GetString(key string) (string, bool) {
	if value, exists := m.Get(key); exists {
		if str, ok := value.(string); ok {
			return str, true
		}
	}
	return "", false
}

// GetInt retrieves a metadata field as an int
func (m Metadata) GetInt(key string) (int, bool) {
	if value, exists := m.Get(key); exists {
		switch v := value.(type) {
		case int:
			return v, true
		case float64:
			return int(v), true
		case json.Number:
			if i, err := v.Int64(); err == nil {
				return int(i), true
			}
		}
	}
	return 0, false
}

// Delete removes a metadata field
func (m Metadata) Delete(key string) {
	if m != nil {
		delete(m, key)
	}
}

// SetCustomFields sets the custom_fields metadata for forms
func (m Metadata) SetCustomFields(fields []CustomField) Metadata {
	return m.Set("custom_fields", fields)
}

// SetCancelAction sets the cancel_action URL for payment pages
func (m Metadata) SetCancelAction(url string) Metadata {
	return m.Set("cancel_action", url)
}

// SetCustomFilters sets the custom_filters for payment channels
func (m Metadata) SetCustomFilters(filters CustomFilters) Metadata {
	return m.Set("custom_filters", filters)
}

// Merge combines this metadata with another, with the other taking precedence
func (m Metadata) Merge(other Metadata) Metadata {
	if m == nil {
		m = make(Metadata)
	}
	for key, value := range other {
		m[key] = value
	}
	return m
}

// Clone creates a deep copy of the metadata
func (m Metadata) Clone() Metadata {
	if m == nil {
		return nil
	}
	clone := make(Metadata, len(m))
	for key, value := range m {
		clone[key] = value
	}
	return clone
}

// String returns a string representation of the metadata
func (m Metadata) String() string {
	if len(m) == 0 {
		return "{}"
	}
	data, _ := json.Marshal(m)
	return string(data)
}

// Validate checks if the metadata contains valid values
func (m Metadata) Validate() error {
	if len(m) == 0 {
		return nil
	}

	// Check for JSON marshaling issues
	if _, err := json.Marshal(m); err != nil {
		return fmt.Errorf("metadata contains invalid values: %w", err)
	}

	return nil
}

// CustomField represents a custom field for forms and payment pages
type CustomField struct {
	DisplayName  string `json:"display_name"`
	VariableName string `json:"variable_name"`
	Value        any    `json:"value"`
}

// CustomFilters represents filters for customizing payment options
type CustomFilters struct {
	Recurring                     bool              `json:"recurring,omitempty"`
	Banks                         []string          `json:"banks,omitempty"`
	CardBrands                    []enums.CardBrand `json:"card_brands,omitempty"`
	SupportedMobileMoneyProviders []enums.MoMo      `json:"supported_mobile_money_providers,omitempty"`
}
