package types

import (
	"encoding/json"

	"github.com/huysamen/paystack-go/types/data"
)

// Response represents the standard Paystack API response wrapper
type Response[T any] struct {
	Status  data.MultiBool `json:"status"`
	Message string         `json:"message"`
	Data    T              `json:"data"`
	Meta    *Meta          `json:"meta,omitempty"`
}

// Meta represents pagination and other metadata
type Meta struct {
	Next      data.NullString `json:"next,omitempty"`
	Previous  data.NullString `json:"previous,omitempty"`
	PerPage   int             `json:"perPage,omitempty"`
	Total     data.NullInt    `json:"total,omitempty"`
	Skipped   data.NullInt    `json:"skipped,omitempty"`
	Page      data.NullInt    `json:"page,omitempty"`
	PageCount data.NullInt    `json:"pageCount,omitempty"`
}

// UnmarshalJSON implements custom JSON unmarshaling to handle both perPage and per_page field names
func (m *Meta) UnmarshalJSON(data []byte) error {
	// Create a temporary struct to unmarshal into
	type Alias Meta
	aux := &struct {
		PerPageAlt int `json:"per_page,omitempty"`
		*Alias
	}{
		Alias: (*Alias)(m),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// If per_page was provided but perPage wasn't, use per_page value
	if aux.PerPageAlt > 0 && m.PerPage == 0 {
		m.PerPage = aux.PerPageAlt
	}

	return nil
}

// ID represents various ID types used across the API
type ID interface {
	~int | ~int64 | ~uint | ~uint64 | ~string
}

// Metadata represents arbitrary key-value pairs that can handle both object and string
type Metadata map[string]any

// UnmarshalJSON implements json.Unmarshaler for Metadata
func (m *Metadata) UnmarshalJSON(data []byte) error {
	// Try to unmarshal as string first (empty metadata case)
	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		if s == "" {
			*m = make(Metadata)
		}
		return nil
	}

	// Try to unmarshal as object
	var obj map[string]any
	if err := json.Unmarshal(data, &obj); err == nil {
		*m = Metadata(obj)
		return nil
	}

	return nil // Be lenient, just return empty metadata if we can't parse
}

// IsEmpty returns true if metadata is nil or empty
func (m Metadata) IsEmpty() bool {
	return len(m) == 0
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
