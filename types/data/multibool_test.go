package data

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMultiBool_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		// Boolean inputs
		{
			name:     "true boolean",
			input:    "true",
			expected: true,
		},
		{
			name:     "false boolean",
			input:    "false",
			expected: false,
		},
		// String inputs that should be true
		{
			name:     "true string",
			input:    `"true"`,
			expected: true,
		},
		{
			name:     "success string",
			input:    `"success"`,
			expected: true,
		},
		// String inputs that should be false
		{
			name:     "false string",
			input:    `"false"`,
			expected: false,
		},
		{
			name:     "failed string",
			input:    `"failed"`,
			expected: false,
		},
		{
			name:     "error string",
			input:    `"error"`,
			expected: false,
		},
		{
			name:     "empty string",
			input:    `""`,
			expected: false,
		},
		// Edge cases
		{
			name:     "null value",
			input:    "null",
			expected: false,
		},
		{
			name:     "number zero",
			input:    "0",
			expected: false,
		},
		{
			name:     "number one",
			input:    "1",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var mb MultiBool
			err := json.Unmarshal([]byte(tt.input), &mb)
			require.NoError(t, err, "should unmarshal without error")
			assert.Equal(t, tt.expected, mb.Bool(), "boolean value should match expected")
		})
	}
}

func TestMultiBool_MarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    MultiBool
		expected string
	}{
		{
			name:     "true value",
			input:    MultiBool(true),
			expected: "true",
		},
		{
			name:     "false value",
			input:    MultiBool(false),
			expected: "false",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := json.Marshal(tt.input)
			require.NoError(t, err, "should marshal without error")
			assert.Equal(t, tt.expected, string(data), "marshaled JSON should match expected")
		})
	}
}

func TestMultiBool_InResponse(t *testing.T) {
	// Test that MultiBool works correctly in a Response structure
	type TestResponse struct {
		Status  MultiBool `json:"status"`
		Message string    `json:"message"`
	}

	tests := []struct {
		name           string
		jsonInput      string
		expectedStatus bool
		expectedMsg    string
	}{
		{
			name:           "boolean true status",
			jsonInput:      `{"status": true, "message": "OK"}`,
			expectedStatus: true,
			expectedMsg:    "OK",
		},
		{
			name:           "string success status",
			jsonInput:      `{"status": "success", "message": "Operation completed"}`,
			expectedStatus: true,
			expectedMsg:    "Operation completed",
		},
		{
			name:           "string false status",
			jsonInput:      `{"status": "failed", "message": "Error occurred"}`,
			expectedStatus: false,
			expectedMsg:    "Error occurred",
		},
		{
			name:           "boolean false status",
			jsonInput:      `{"status": false, "message": "Failed"}`,
			expectedStatus: false,
			expectedMsg:    "Failed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var response TestResponse
			err := json.Unmarshal([]byte(tt.jsonInput), &response)
			require.NoError(t, err, "should unmarshal response without error")

			assert.Equal(t, tt.expectedStatus, response.Status.Bool(), "status should match expected")
			assert.Equal(t, tt.expectedMsg, response.Message, "message should match expected")
		})
	}
}
