package types

import (
	"strings"
)

type Response[T any] struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    T      `json:"data,omitempty"`
	Meta    Meta   `json:"meta,omitempty"`
	Type    string `json:"type,omitempty"`
	Code    string `json:"code,omitempty"`
}

// IsSuccess returns true if the API call was successful
func (r *Response[T]) IsSuccess() bool {
	return r.Status
}

// IsError returns true if the API call resulted in an error
func (r *Response[T]) IsError() bool {
	return !r.Status
}

// IsAuthenticationError returns true if the error is related to authentication
func (r *Response[T]) IsAuthenticationError() bool {
	if r.Status {
		return false
	}
	return r.Code == "invalid_Key" ||
		strings.Contains(strings.ToLower(r.Message), "invalid key") ||
		strings.Contains(strings.ToLower(r.Message), "unauthorized")
}

// IsValidationError returns true if the error is a validation error
func (r *Response[T]) IsValidationError() bool {
	if r.Status {
		return false
	}
	return r.Type == "validation_error" ||
		strings.Contains(strings.ToLower(r.Type), "validation")
}

// IsNotFoundError returns true if the resource was not found
func (r *Response[T]) IsNotFoundError() bool {
	if r.Status {
		return false
	}
	return r.Type == "not_found_error" ||
		strings.Contains(strings.ToLower(r.Message), "not found")
}

// IsRateLimitError returns true if the error is due to rate limiting
func (r *Response[T]) IsRateLimitError() bool {
	if r.Status {
		return false
	}
	return r.Type == "rate_limit_error" ||
		strings.Contains(strings.ToLower(r.Message), "rate limit") ||
		strings.Contains(strings.ToLower(r.Message), "too many requests")
}

// GetNextStep returns the recommended next step if available in the error meta
func (r *Response[T]) GetNextStep() string {
	return r.Meta.NextStep
}

// GetErrorMessage returns the error message, or empty string if successful
func (r *Response[T]) GetErrorMessage() string {
	if r.Status {
		return ""
	}
	return r.Message
}

// GetErrorCode returns the error code, or empty string if successful or no code
func (r *Response[T]) GetErrorCode() string {
	if r.Status {
		return ""
	}
	return r.Code
}

// GetErrorType returns the error type, or empty string if successful or no type
func (r *Response[T]) GetErrorType() string {
	if r.Status {
		return ""
	}
	return r.Type
}

type Meta struct {
	Total     int    `json:"total,omitempty"`
	Skipped   int    `json:"skipped,omitempty"`
	PerPage   int    `json:"perPage,omitempty"`
	Page      int    `json:"page,omitempty"`
	PageCount int    `json:"pageCount,omitempty"`
	Next      string `json:"next,omitempty"`
	Previous  string `json:"previous,omitempty"`
	NextStep  string `json:"nextStep,omitempty"`
}
