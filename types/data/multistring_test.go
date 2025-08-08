package data

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMultiString_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		// String inputs
		{
			name:     "regular string",
			input:    `"hello"`,
			expected: "hello",
		},
		{
			name:     "empty string",
			input:    `""`,
			expected: "",
		},
		{
			name:     "numeric string",
			input:    `"123"`,
			expected: "123",
		},
		// Number inputs
		{
			name:     "integer number",
			input:    "42",
			expected: "42",
		},
		{
			name:     "float number",
			input:    "12.34",
			expected: "12",
		},
		{
			name:     "zero",
			input:    "0",
			expected: "0",
		},
		// Edge cases
		{
			name:     "null value",
			input:    "null",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var ms MultiString
			err := json.Unmarshal([]byte(tt.input), &ms)
			require.NoError(t, err, "should unmarshal without error")
			assert.Equal(t, tt.expected, ms.String(), "string value should match expected")
		})
	}
}

func TestMultiString_UnmarshalJSON_InvalidInput(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{
			name:  "boolean true",
			input: "true",
		},
		{
			name:  "boolean false",
			input: "false",
		},
		{
			name:  "array",
			input: `["test"]`,
		},
		{
			name:  "object",
			input: `{"key": "value"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var ms MultiString
			err := json.Unmarshal([]byte(tt.input), &ms)
			assert.Error(t, err, "should return error for invalid input")
		})
	}
}

func TestMultiString_MarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    MultiString
		expected string
	}{
		{
			name:     "regular string",
			input:    MultiString("hello"),
			expected: `"hello"`,
		},
		{
			name:     "empty string",
			input:    MultiString(""),
			expected: `""`,
		},
		{
			name:     "numeric string",
			input:    MultiString("123"),
			expected: `"123"`,
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

func TestMultiString_InStruct(t *testing.T) {
	// Test that MultiString works correctly in a struct
	type TestStruct struct {
		ExpMonth MultiString `json:"exp_month"`
		ExpYear  MultiString `json:"exp_year"`
		Name     string      `json:"name"`
	}

	tests := []struct {
		name             string
		jsonInput        string
		expectedExpMonth string
		expectedExpYear  string
		expectedName     string
	}{
		{
			name:             "string values",
			jsonInput:        `{"exp_month": "12", "exp_year": "2025", "name": "Test"}`,
			expectedExpMonth: "12",
			expectedExpYear:  "2025",
			expectedName:     "Test",
		},
		{
			name:             "number values",
			jsonInput:        `{"exp_month": 12, "exp_year": 2025, "name": "Test"}`,
			expectedExpMonth: "12",
			expectedExpYear:  "2025",
			expectedName:     "Test",
		},
		{
			name:             "mixed values",
			jsonInput:        `{"exp_month": "12", "exp_year": 2025, "name": "Test"}`,
			expectedExpMonth: "12",
			expectedExpYear:  "2025",
			expectedName:     "Test",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var testStruct TestStruct
			err := json.Unmarshal([]byte(tt.jsonInput), &testStruct)
			require.NoError(t, err, "should unmarshal struct without error")

			assert.Equal(t, tt.expectedExpMonth, testStruct.ExpMonth.String(), "exp_month should match expected")
			assert.Equal(t, tt.expectedExpYear, testStruct.ExpYear.String(), "exp_year should match expected")
			assert.Equal(t, tt.expectedName, testStruct.Name, "name should match expected")
		})
	}
}
