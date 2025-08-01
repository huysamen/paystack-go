package net

import (
	"encoding/json"
	"testing"
)

func TestPaystackError_Error(t *testing.T) {
	tests := []struct {
		name     string
		err      *PaystackError
		expected string
	}{
		{
			name: "basic error with just message",
			err: &PaystackError{
				StatusCode: 400,
				Message:    "Invalid request",
				Status:     false,
			},
			expected: "paystack api error (status 400): Invalid request",
		},
		{
			name: "error with code and type",
			err: &PaystackError{
				StatusCode: 401,
				Message:    "Invalid key",
				Status:     false,
				Code:       "invalid_Key",
				Type:       "validation_error",
			},
			expected: "paystack api error (status 401): Invalid key [code: invalid_Key] [type: validation_error]",
		},
		{
			name: "error with meta next step",
			err: &PaystackError{
				StatusCode: 401,
				Message:    "Invalid key",
				Status:     false,
				Code:       "invalid_Key",
				Type:       "validation_error",
				Meta: map[string]any{
					"nextStep": "Ensure that you provide the correct authorization key for the request",
				},
			},
			expected: "paystack api error (status 401): Invalid key [code: invalid_Key] [type: validation_error] [next step: Ensure that you provide the correct authorization key for the request]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.err.Error(); got != tt.expected {
				t.Errorf("PaystackError.Error() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestPaystackError_IsClientError(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
		expected   bool
	}{
		{"400 is client error", 400, true},
		{"404 is client error", 404, true},
		{"499 is client error", 499, true},
		{"200 is not client error", 200, false},
		{"500 is not client error", 500, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := &PaystackError{StatusCode: tt.statusCode}
			if got := err.IsClientError(); got != tt.expected {
				t.Errorf("PaystackError.IsClientError() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestPaystackError_IsServerError(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
		expected   bool
	}{
		{"500 is server error", 500, true},
		{"502 is server error", 502, true},
		{"599 is server error", 599, true},
		{"400 is not server error", 400, false},
		{"200 is not server error", 200, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := &PaystackError{StatusCode: tt.statusCode}
			if got := err.IsServerError(); got != tt.expected {
				t.Errorf("PaystackError.IsServerError() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestPaystackError_IsAuthenticationError(t *testing.T) {
	tests := []struct {
		name     string
		err      *PaystackError
		expected bool
	}{
		{
			name:     "401 status is authentication error",
			err:      &PaystackError{StatusCode: 401},
			expected: true,
		},
		{
			name:     "invalid_Key code is authentication error",
			err:      &PaystackError{StatusCode: 400, Code: "invalid_Key"},
			expected: true,
		},
		{
			name:     "invalid key message is authentication error",
			err:      &PaystackError{StatusCode: 400, Message: "Invalid key"},
			expected: true,
		},
		{
			name:     "normal 400 is not authentication error",
			err:      &PaystackError{StatusCode: 400, Message: "Bad request"},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.err.IsAuthenticationError(); got != tt.expected {
				t.Errorf("PaystackError.IsAuthenticationError() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestPaystackError_IsValidationError(t *testing.T) {
	tests := []struct {
		name     string
		err      *PaystackError
		expected bool
	}{
		{
			name:     "400 status is validation error",
			err:      &PaystackError{StatusCode: 400},
			expected: true,
		},
		{
			name:     "422 status is validation error",
			err:      &PaystackError{StatusCode: 422},
			expected: true,
		},
		{
			name:     "validation_error type is validation error",
			err:      &PaystackError{StatusCode: 500, Type: "validation_error"},
			expected: true,
		},
		{
			name:     "200 status is not validation error",
			err:      &PaystackError{StatusCode: 200},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.err.IsValidationError(); got != tt.expected {
				t.Errorf("PaystackError.IsValidationError() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestPaystackError_GetNextStep(t *testing.T) {
	tests := []struct {
		name     string
		err      *PaystackError
		expected string
	}{
		{
			name: "meta with nextStep",
			err: &PaystackError{
				Meta: map[string]any{
					"nextStep": "Check your API key",
				},
			},
			expected: "Check your API key",
		},
		{
			name: "meta without nextStep",
			err: &PaystackError{
				Meta: map[string]any{
					"otherField": "value",
				},
			},
			expected: "",
		},
		{
			name:     "no meta",
			err:      &PaystackError{},
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.err.GetNextStep(); got != tt.expected {
				t.Errorf("PaystackError.GetNextStep() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestErrorConstructors(t *testing.T) {
	t.Run("NewValidationError", func(t *testing.T) {
		err := NewValidationError("Invalid email format")
		if err.StatusCode != 400 {
			t.Errorf("Expected status code 400, got %d", err.StatusCode)
		}
		if err.Message != "Invalid email format" {
			t.Errorf("Expected message 'Invalid email format', got %s", err.Message)
		}
		if err.Type != ErrorTypeValidation {
			t.Errorf("Expected type %s, got %s", ErrorTypeValidation, err.Type)
		}
	})

	t.Run("NewAuthenticationError", func(t *testing.T) {
		err := NewAuthenticationError("Invalid API key")
		if err.StatusCode != 401 {
			t.Errorf("Expected status code 401, got %d", err.StatusCode)
		}
		if err.Code != ErrorCodeInvalidKey {
			t.Errorf("Expected code %s, got %s", ErrorCodeInvalidKey, err.Code)
		}
		if err.Type != ErrorTypeAuthentication {
			t.Errorf("Expected type %s, got %s", ErrorTypeAuthentication, err.Type)
		}
	})

	t.Run("NewNotFoundError", func(t *testing.T) {
		err := NewNotFoundError("Resource not found")
		if err.StatusCode != 404 {
			t.Errorf("Expected status code 404, got %d", err.StatusCode)
		}
		if err.Type != ErrorTypeNotFound {
			t.Errorf("Expected type %s, got %s", ErrorTypeNotFound, err.Type)
		}
	})
}

func TestGetHTTPErrorMessage(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
		body       []byte
		expected   string
	}{
		{
			name:       "JSON with message field",
			statusCode: 400,
			body:       []byte(`{"message": "Invalid request", "status": false}`),
			expected:   "Invalid request",
		},
		{
			name:       "JSON with error field",
			statusCode: 400,
			body:       []byte(`{"error": "Bad request", "status": false}`),
			expected:   "Bad request",
		},
		{
			name:       "401 fallback",
			statusCode: 401,
			body:       nil,
			expected:   "Unauthorized: Invalid or missing API key",
		},
		{
			name:       "404 fallback",
			statusCode: 404,
			body:       nil,
			expected:   "Not found: The requested resource does not exist",
		},
		{
			name:       "generic fallback",
			statusCode: 418,
			body:       nil,
			expected:   "HTTP 418: I'm a teapot",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getHTTPErrorMessage(tt.statusCode, tt.body); got != tt.expected {
				t.Errorf("getHTTPErrorMessage() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestPaystackErrorJSONUnmarshaling(t *testing.T) {
	// Test that we can properly unmarshal a Paystack error response
	jsonResponse := `{
		"status": false,
		"message": "Invalid key",
		"meta": {
			"nextStep": "Ensure that you provide the correct authorization key for the request"
		},
		"type": "validation_error",
		"code": "invalid_Key"
	}`

	var err PaystackError
	if unmarshalErr := json.Unmarshal([]byte(jsonResponse), &err); unmarshalErr != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", unmarshalErr)
	}

	if err.Message != "Invalid key" {
		t.Errorf("Expected message 'Invalid key', got %s", err.Message)
	}
	if err.Code != "invalid_Key" {
		t.Errorf("Expected code 'invalid_Key', got %s", err.Code)
	}
	if err.Type != "validation_error" {
		t.Errorf("Expected type 'validation_error', got %s", err.Type)
	}
	if err.Status != false {
		t.Errorf("Expected status false, got %v", err.Status)
	}

	nextStep := err.GetNextStep()
	expectedNextStep := "Ensure that you provide the correct authorization key for the request"
	if nextStep != expectedNextStep {
		t.Errorf("Expected nextStep '%s', got '%s'", expectedNextStep, nextStep)
	}
}
