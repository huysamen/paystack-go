package types

import (
	"encoding/json"

	"github.com/huysamen/paystack-go/types/data"
)

// Response represents the standard Paystack API response wrapper
type Response[T any] struct {
	Status  data.Bool `json:"status"`
	Message string    `json:"message"`
	Data    T         `json:"data"`
	Meta    *Meta     `json:"meta,omitempty"`
}

// Meta represents pagination and other metadata
type Meta struct {
	Next      data.NullString `json:"next,omitempty"`
	Previous  data.NullString `json:"previous,omitempty"`
	PerPage   data.Int        `json:"perPage,omitempty"`
	Total     data.NullInt    `json:"total,omitempty"`
	Skipped   data.NullInt    `json:"skipped,omitempty"`
	Page      data.NullInt    `json:"page,omitempty"`
	PageCount data.NullInt    `json:"pageCount,omitempty"`
}

// UnmarshalJSON implements custom JSON unmarshaling to handle both perPage and per_page field names
func (m *Meta) UnmarshalJSON(jsonData []byte) error {
	// Create a temporary struct to unmarshal into
	type Alias Meta
	aux := &struct {
		PerPageAlt int `json:"per_page,omitempty"`
		*Alias
	}{
		Alias: (*Alias)(m),
	}

	if err := json.Unmarshal(jsonData, &aux); err != nil {
		return err
	}

	// If per_page was provided but perPage wasn't, use per_page value
	if aux.PerPageAlt > 0 && m.PerPage.Int64() == 0 {
		m.PerPage = data.Int(aux.PerPageAlt)
	}

	return nil
}

// ID represents various ID types used across the API
type ID interface {
	~int | ~int64 | ~uint | ~uint64 | ~string
}

// CustomField represents a custom field in metadata
type CustomField struct {
	DisplayName  string `json:"display_name"`
	VariableName string `json:"variable_name"`
	Value        any    `json:"value,omitempty"`
	Required     bool   `json:"required,omitempty"`
}

// CustomFields represents an array of custom fields
type CustomFields struct {
	CustomFields []CustomField `json:"custom_fields"`
}
