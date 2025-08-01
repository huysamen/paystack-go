package net

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/huysamen/paystack-go/types"
)

const apiURL = "https://api.paystack.co"

// Common Paystack error codes
const (
	ErrorCodeInvalidKey            = "invalid_Key"
	ErrorCodeValidationError       = "validation_error"
	ErrorCodeInsufficientFunds     = "insufficient_funds"
	ErrorCodeDuplicateReference    = "duplicate_reference"
	ErrorCodeTransactionNotFound   = "transaction_not_found"
	ErrorCodeCustomerNotFound      = "customer_not_found"
	ErrorCodePlanNotFound          = "plan_not_found"
	ErrorCodeAuthorizationNotFound = "authorization_not_found"
)

// Common Paystack error types
const (
	ErrorTypeValidation     = "validation_error"
	ErrorTypeAuthentication = "authentication_error"
	ErrorTypeAuthorization  = "authorization_error"
	ErrorTypeNotFound       = "not_found_error"
	ErrorTypeServerError    = "server_error"
	ErrorTypeRateLimit      = "rate_limit_error"
	ErrorTypeUnexpected     = "unexpected_error"
)

// PaystackError represents an error response from the Paystack API
type PaystackError struct {
	StatusCode int            `json:"-"`
	Message    string         `json:"message"`
	Status     bool           `json:"status"`
	Code       string         `json:"code,omitempty"`
	Type       string         `json:"type,omitempty"`
	Meta       map[string]any `json:"meta,omitempty"`
	Cause      error          `json:"-"` // Underlying cause of the error
}

func (e *PaystackError) Error() string {
	var errorMsg strings.Builder

	// Base error format
	errorMsg.WriteString(fmt.Sprintf("paystack api error (status %d): %s", e.StatusCode, e.Message))

	// Add error code if available
	if e.Code != "" {
		errorMsg.WriteString(fmt.Sprintf(" [code: %s]", e.Code))
	}

	// Add error type if available
	if e.Type != "" {
		errorMsg.WriteString(fmt.Sprintf(" [type: %s]", e.Type))
	}

	// Add meta information if available
	if len(e.Meta) > 0 {
		if nextStep, ok := e.Meta["nextStep"].(string); ok && nextStep != "" {
			errorMsg.WriteString(fmt.Sprintf(" [next step: %s]", nextStep))
		}
	}

	// Add underlying cause if available
	if e.Cause != nil {
		errorMsg.WriteString(fmt.Sprintf(", cause: %v", e.Cause))
	}

	return errorMsg.String()
}

// Unwrap returns the underlying cause error for error wrapping support
func (e *PaystackError) Unwrap() error {
	return e.Cause
}

// Is allows error comparison using errors.Is
func (e *PaystackError) Is(target error) bool {
	if t, ok := target.(*PaystackError); ok {
		return e.StatusCode == t.StatusCode && e.Code == t.Code
	}
	return false
}

// IsClientError returns true if the error is a client error (4xx status code)
func (e *PaystackError) IsClientError() bool {
	return e.StatusCode >= 400 && e.StatusCode < 500
}

// IsServerError returns true if the error is a server error (5xx status code)
func (e *PaystackError) IsServerError() bool {
	return e.StatusCode >= 500 && e.StatusCode < 600
}

// IsAuthenticationError returns true if the error is related to authentication
func (e *PaystackError) IsAuthenticationError() bool {
	return e.StatusCode == http.StatusUnauthorized ||
		e.Code == "invalid_Key" ||
		strings.Contains(strings.ToLower(e.Message), "invalid key")
}

// IsValidationError returns true if the error is a validation error
func (e *PaystackError) IsValidationError() bool {
	return e.StatusCode == http.StatusBadRequest ||
		e.StatusCode == http.StatusUnprocessableEntity ||
		e.Type == "validation_error" ||
		strings.Contains(strings.ToLower(e.Type), "validation")
}

// IsRateLimitError returns true if the error is due to rate limiting
func (e *PaystackError) IsRateLimitError() bool {
	return e.StatusCode == http.StatusTooManyRequests
}

// IsNotFoundError returns true if the resource was not found
func (e *PaystackError) IsNotFoundError() bool {
	return e.StatusCode == http.StatusNotFound
}

// GetNextStep returns the recommended next step if available in the error meta
func (e *PaystackError) GetNextStep() string {
	if len(e.Meta) > 0 {
		if nextStep, ok := e.Meta["nextStep"].(string); ok {
			return nextStep
		}
	}
	return ""
}

// NewPaystackError creates a new PaystackError with the given parameters
func NewPaystackError(statusCode int, message, code, errorType string) *PaystackError {
	return &PaystackError{
		StatusCode: statusCode,
		Message:    message,
		Status:     false,
		Code:       code,
		Type:       errorType,
	}
}

// NewValidationError creates a validation error
func NewValidationError(message string) *PaystackError {
	return NewPaystackError(http.StatusBadRequest, message, "", ErrorTypeValidation)
}

// NewAuthenticationError creates an authentication error
func NewAuthenticationError(message string) *PaystackError {
	return NewPaystackError(http.StatusUnauthorized, message, ErrorCodeInvalidKey, ErrorTypeAuthentication)
}

// NewNotFoundError creates a not found error
func NewNotFoundError(message string) *PaystackError {
	return NewPaystackError(http.StatusNotFound, message, "", ErrorTypeNotFound)
}

func getBaseURL(baseURL ...string) string {
	if len(baseURL) > 0 && baseURL[0] != "" {
		return baseURL[0]
	}
	return apiURL
}

// getHTTPErrorMessage generates a meaningful error message from HTTP status and response body
func getHTTPErrorMessage(statusCode int, body []byte) string {
	// Try to extract a meaningful message from the response body
	if len(body) > 0 {
		// Try to parse as JSON and extract message
		var response map[string]any
		if err := json.Unmarshal(body, &response); err == nil {
			if msg, ok := response["message"].(string); ok && msg != "" {
				return msg
			}
			if msg, ok := response["error"].(string); ok && msg != "" {
				return msg
			}
		}
		// If not JSON or no message field, use raw body if it's short and looks like text
		if len(body) < 200 && !json.Valid(body) {
			return strings.TrimSpace(string(body))
		}
	}

	// Fallback to standard HTTP status text
	switch statusCode {
	case http.StatusBadRequest:
		return "Bad request: The request was invalid or malformed"
	case http.StatusUnauthorized:
		return "Unauthorized: Invalid or missing API key"
	case http.StatusForbidden:
		return "Forbidden: Access denied or insufficient permissions"
	case http.StatusNotFound:
		return "Not found: The requested resource does not exist"
	case http.StatusUnprocessableEntity:
		return "Unprocessable entity: Validation failed"
	case http.StatusTooManyRequests:
		return "Rate limited: Too many requests, please try again later"
	case http.StatusConflict:
		return "Conflict: The request conflicts with the current state"
	case http.StatusPreconditionFailed:
		return "Precondition failed: Required conditions were not met"
	default:
		return fmt.Sprintf("HTTP %d: %s", statusCode, http.StatusText(statusCode))
	}
}

// Get makes a GET request with context support
func Get[O any](ctx context.Context, client *http.Client, secret, path string, baseURL ...string) (*types.Response[O], error) {
	url := getBaseURL(baseURL...)
	body, err := doReq(ctx, client, http.MethodGet, secret, url+path, nil)
	if err != nil {
		return nil, err
	}

	rsp := new(types.Response[O])

	if len(body) > 0 {
		err = json.Unmarshal(body, rsp)
		if err != nil {
			return nil, err
		}
	}

	return rsp, nil
}

// Post makes a POST request with context support
func Post[I any, O any](ctx context.Context, client *http.Client, secret, path string, payload *I, baseURL ...string) (*types.Response[O], error) {
	url := getBaseURL(baseURL...)
	return putOrPost[I, O](ctx, client, http.MethodPost, secret, url+path, payload)
}

// Put makes a PUT request with context support
func Put[I any, O any](ctx context.Context, client *http.Client, secret, path string, payload *I, baseURL ...string) (*types.Response[O], error) {
	url := getBaseURL(baseURL...)
	return putOrPost[I, O](ctx, client, http.MethodPut, secret, url+path, payload)
}

// Delete makes a DELETE request with context support
func Delete[O any](ctx context.Context, client *http.Client, secret, path string, baseURL ...string) (*types.Response[O], error) {
	url := getBaseURL(baseURL...)
	body, err := doReq(ctx, client, http.MethodDelete, secret, url+path, nil)
	if err != nil {
		return nil, err
	}

	rsp := new(types.Response[O])

	if len(body) > 0 {
		err = json.Unmarshal(body, rsp)
		if err != nil {
			return nil, err
		}
	}

	return rsp, nil
}

func putOrPost[I any, O any](ctx context.Context, client *http.Client, method, secret, fullURL string, payload *I) (*types.Response[O], error) {
	body, err := doReq(ctx, client, method, secret, fullURL, payload)
	if err != nil {
		return nil, err
	}

	rsp := new(types.Response[O])

	if len(body) > 0 {
		err = json.Unmarshal(body, rsp)
		if err != nil {
			return nil, err
		}
	}

	return rsp, nil
}

func doReq(ctx context.Context, client *http.Client, method, secret, fullURL string, data any) ([]byte, error) {
	var req *http.Request
	var err error

	if data != nil {
		d, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}

		req, err = http.NewRequestWithContext(ctx, method, fullURL, bytes.NewBuffer(d))
		if err != nil {
			return nil, err
		}
	} else {
		req, err = http.NewRequestWithContext(ctx, method, fullURL, nil)
		if err != nil {
			return nil, err
		}
	}

	req.Header.Add("Authorization", "Bearer "+secret)

	if data != nil {
		req.Header.Add("Content-Type", "application/json")
	}

	rsp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() { _ = rsp.Body.Close() }()

	body, err := io.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}

	switch rsp.StatusCode {
	case http.StatusOK, http.StatusCreated:
		return body, nil
	case http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden,
		http.StatusNotFound, http.StatusUnprocessableEntity, http.StatusTooManyRequests,
		http.StatusConflict, http.StatusPreconditionFailed:
		// Try to parse Paystack error response
		var paystackErr PaystackError
		if len(body) > 0 {
			if err := json.Unmarshal(body, &paystackErr); err == nil {
				paystackErr.StatusCode = rsp.StatusCode
				return nil, &paystackErr
			}
		}
		// Fallback to generic error if parsing fails
		return nil, &PaystackError{
			StatusCode: rsp.StatusCode,
			Message:    getHTTPErrorMessage(rsp.StatusCode, body),
			Status:     false,
		}
	case http.StatusInternalServerError, http.StatusBadGateway, http.StatusServiceUnavailable,
		http.StatusGatewayTimeout:
		// Server errors - these should be reported to Paystack
		return nil, &PaystackError{
			StatusCode: rsp.StatusCode,
			Message:    fmt.Sprintf("Paystack server error (HTTP %d). Please report this to Paystack support.", rsp.StatusCode),
			Status:     false,
			Type:       "server_error",
		}
	default:
		// Unexpected status codes
		return nil, &PaystackError{
			StatusCode: rsp.StatusCode,
			Message:    fmt.Sprintf("unexpected HTTP status code: %d", rsp.StatusCode),
			Status:     false,
			Type:       "unexpected_error",
		}
	}
}
