package types

import "testing"

func TestResponse_IsSuccess(t *testing.T) {
	tests := []struct {
		name     string
		response Response[any]
		expected bool
	}{
		{
			name:     "successful response",
			response: Response[any]{Status: true},
			expected: true,
		},
		{
			name:     "failed response",
			response: Response[any]{Status: false},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.response.IsSuccess(); got != tt.expected {
				t.Errorf("IsSuccess() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestResponse_IsError(t *testing.T) {
	tests := []struct {
		name     string
		response Response[any]
		expected bool
	}{
		{
			name:     "successful response",
			response: Response[any]{Status: true},
			expected: false,
		},
		{
			name:     "failed response",
			response: Response[any]{Status: false},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.response.IsError(); got != tt.expected {
				t.Errorf("IsError() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestResponse_IsAuthenticationError(t *testing.T) {
	tests := []struct {
		name     string
		response Response[any]
		expected bool
	}{
		{
			name:     "successful response",
			response: Response[any]{Status: true},
			expected: false,
		},
		{
			name: "invalid_Key code",
			response: Response[any]{
				Status: false,
				Code:   "invalid_Key",
			},
			expected: true,
		},
		{
			name: "invalid key message",
			response: Response[any]{
				Status:  false,
				Message: "Invalid key provided",
			},
			expected: true,
		},
		{
			name: "unauthorized message",
			response: Response[any]{
				Status:  false,
				Message: "Unauthorized access",
			},
			expected: true,
		},
		{
			name: "validation error",
			response: Response[any]{
				Status:  false,
				Message: "Bad request",
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.response.IsAuthenticationError(); got != tt.expected {
				t.Errorf("IsAuthenticationError() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestResponse_IsValidationError(t *testing.T) {
	tests := []struct {
		name     string
		response Response[any]
		expected bool
	}{
		{
			name:     "successful response",
			response: Response[any]{Status: true},
			expected: false,
		},
		{
			name: "validation_error type",
			response: Response[any]{
				Status: false,
				Type:   "validation_error",
			},
			expected: true,
		},
		{
			name: "contains validation in type",
			response: Response[any]{
				Status: false,
				Type:   "form_validation_failed",
			},
			expected: true,
		},
		{
			name: "authentication error",
			response: Response[any]{
				Status: false,
				Code:   "invalid_Key",
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.response.IsValidationError(); got != tt.expected {
				t.Errorf("IsValidationError() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestResponse_IsNotFoundError(t *testing.T) {
	tests := []struct {
		name     string
		response Response[any]
		expected bool
	}{
		{
			name:     "successful response",
			response: Response[any]{Status: true},
			expected: false,
		},
		{
			name: "not_found_error type",
			response: Response[any]{
				Status: false,
				Type:   "not_found_error",
			},
			expected: true,
		},
		{
			name: "not found message",
			response: Response[any]{
				Status:  false,
				Message: "Resource not found",
			},
			expected: true,
		},
		{
			name: "validation error",
			response: Response[any]{
				Status: false,
				Type:   "validation_error",
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.response.IsNotFoundError(); got != tt.expected {
				t.Errorf("IsNotFoundError() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestResponse_IsRateLimitError(t *testing.T) {
	tests := []struct {
		name     string
		response Response[any]
		expected bool
	}{
		{
			name:     "successful response",
			response: Response[any]{Status: true},
			expected: false,
		},
		{
			name: "rate_limit_error type",
			response: Response[any]{
				Status: false,
				Type:   "rate_limit_error",
			},
			expected: true,
		},
		{
			name: "rate limit message",
			response: Response[any]{
				Status:  false,
				Message: "Too many requests, please try again later",
			},
			expected: true,
		},
		{
			name: "validation error",
			response: Response[any]{
				Status: false,
				Type:   "validation_error",
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.response.IsRateLimitError(); got != tt.expected {
				t.Errorf("IsRateLimitError() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestResponse_GetNextStep(t *testing.T) {
	tests := []struct {
		name     string
		response Response[any]
		expected string
	}{
		{
			name: "with next step",
			response: Response[any]{
				Meta: Meta{
					NextStep: "Check your API key",
				},
			},
			expected: "Check your API key",
		},
		{
			name:     "without next step",
			response: Response[any]{},
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.response.GetNextStep(); got != tt.expected {
				t.Errorf("GetNextStep() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestResponse_GetErrorMessage(t *testing.T) {
	tests := []struct {
		name     string
		response Response[any]
		expected string
	}{
		{
			name: "successful response",
			response: Response[any]{
				Status:  true,
				Message: "Success",
			},
			expected: "",
		},
		{
			name: "error response",
			response: Response[any]{
				Status:  false,
				Message: "Invalid API key",
			},
			expected: "Invalid API key",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.response.GetErrorMessage(); got != tt.expected {
				t.Errorf("GetErrorMessage() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestResponse_GetErrorCode(t *testing.T) {
	tests := []struct {
		name     string
		response Response[any]
		expected string
	}{
		{
			name: "successful response",
			response: Response[any]{
				Status: true,
				Code:   "success",
			},
			expected: "",
		},
		{
			name: "error response with code",
			response: Response[any]{
				Status: false,
				Code:   "invalid_Key",
			},
			expected: "invalid_Key",
		},
		{
			name: "error response without code",
			response: Response[any]{
				Status: false,
			},
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.response.GetErrorCode(); got != tt.expected {
				t.Errorf("GetErrorCode() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestResponse_GetErrorType(t *testing.T) {
	tests := []struct {
		name     string
		response Response[any]
		expected string
	}{
		{
			name: "successful response",
			response: Response[any]{
				Status: true,
				Type:   "success",
			},
			expected: "",
		},
		{
			name: "error response with type",
			response: Response[any]{
				Status: false,
				Type:   "validation_error",
			},
			expected: "validation_error",
		},
		{
			name: "error response without type",
			response: Response[any]{
				Status: false,
			},
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.response.GetErrorType(); got != tt.expected {
				t.Errorf("GetErrorType() = %v, want %v", got, tt.expected)
			}
		})
	}
}
