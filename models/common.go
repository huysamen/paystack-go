package models

// Response represents the standard Paystack API response wrapper
type Response[T any] struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    T      `json:"data"`
	Meta    *Meta  `json:"meta,omitempty"`
}

// Meta represents pagination and other metadata
type Meta struct {
	Next      *string `json:"next,omitempty"`
	Previous  *string `json:"previous,omitempty"`
	PerPage   int     `json:"perPage,omitempty"`
	Total     *int    `json:"total,omitempty"`
	Skipped   *int    `json:"skipped,omitempty"`
	Page      *int    `json:"page,omitempty"`
	PageCount *int    `json:"pageCount,omitempty"`
}

// ID represents various ID types used across the API
type ID interface {
	~int | ~int64 | ~uint | ~uint64 | ~string
}

// Metadata represents arbitrary key-value pairs
type Metadata map[string]any

// IsEmpty returns true if metadata is nil or empty
func (m Metadata) IsEmpty() bool {
	return len(m) == 0
}

// CustomField represents a custom field in metadata
type CustomField struct {
	DisplayName  string `json:"display_name"`
	VariableName string `json:"variable_name"`
	Value        any    `json:"value"`
}

// CustomFields represents an array of custom fields
type CustomFields struct {
	CustomFields []CustomField `json:"custom_fields"`
}
