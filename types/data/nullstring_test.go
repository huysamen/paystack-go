package data

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNullString_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expectedStr   string
		expectedValid bool
		wantErr       bool
	}{
		// String inputs
		{
			name:          "unmarshals regular string",
			input:         `"hello"`,
			expectedStr:   "hello",
			expectedValid: true,
		},
		{
			name:          "unmarshals string with spaces",
			input:         `"hello world"`,
			expectedStr:   "hello world",
			expectedValid: true,
		},
		{
			name:          "unmarshals numeric string",
			input:         `"123"`,
			expectedStr:   "123",
			expectedValid: true,
		},
		{
			name:          "unmarshals string with special characters",
			input:         `"hello@example.com"`,
			expectedStr:   "hello@example.com",
			expectedValid: true,
		},
		// Empty string handling
		{
			name:          "unmarshals empty string as null",
			input:         `""`,
			expectedStr:   "",
			expectedValid: false,
		},
		{
			name:          "unmarshals string null as null",
			input:         `"null"`,
			expectedStr:   "",
			expectedValid: false,
		},
		// Number inputs (converted to strings)
		{
			name:          "unmarshals positive integer as string",
			input:         "42",
			expectedStr:   "42",
			expectedValid: true,
		},
		{
			name:          "unmarshals negative integer as string",
			input:         "-10",
			expectedStr:   "-10",
			expectedValid: true,
		},
		{
			name:          "unmarshals zero as string",
			input:         "0",
			expectedStr:   "0",
			expectedValid: true,
		},
		{
			name:          "unmarshals float as string",
			input:         "3.14",
			expectedStr:   "3.14",
			expectedValid: true,
		},
		{
			name:          "unmarshals scientific notation as string",
			input:         "1.23e+10",
			expectedStr:   "1.23e+10",
			expectedValid: true,
		},
		// Null handling
		{
			name:          "unmarshals null as invalid",
			input:         "null",
			expectedStr:   "",
			expectedValid: false,
		},
		// Error cases
		{
			name:    "fails to unmarshal boolean",
			input:   "true",
			wantErr: true,
		},
		{
			name:    "fails to unmarshal object",
			input:   `{"key": "value"}`,
			wantErr: true,
		},
		{
			name:    "fails to unmarshal array",
			input:   `[1, 2, 3]`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var ns NullString
			err := json.Unmarshal([]byte(tt.input), &ns)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.expectedStr, ns.Str)
			assert.Equal(t, tt.expectedValid, ns.Valid)
		})
	}
}

func TestNullString_MarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    NullString
		expected string
	}{
		{
			name:     "marshals valid string",
			input:    NullString{Str: "hello", Valid: true},
			expected: `"hello"`,
		},
		{
			name:     "marshals valid empty string (but valid)",
			input:    NullString{Str: "", Valid: true},
			expected: `""`,
		},
		{
			name:     "marshals valid string with spaces",
			input:    NullString{Str: "hello world", Valid: true},
			expected: `"hello world"`,
		},
		{
			name:     "marshals valid numeric string",
			input:    NullString{Str: "123", Valid: true},
			expected: `"123"`,
		},
		{
			name:     "marshals invalid as null (with string value)",
			input:    NullString{Str: "hello", Valid: false},
			expected: "null",
		},
		{
			name:     "marshals invalid as null (with empty string)",
			input:    NullString{Str: "", Valid: false},
			expected: "null",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := json.Marshal(tt.input)
			require.NoError(t, err)
			assert.Equal(t, tt.expected, string(result))
		})
	}
}

func TestNullString_ValueOr(t *testing.T) {
	tests := []struct {
		name     string
		input    NullString
		fallback string
		expected string
	}{
		{
			name:     "returns value when valid",
			input:    NullString{Str: "hello", Valid: true},
			fallback: "default",
			expected: "hello",
		},
		{
			name:     "returns empty string when valid",
			input:    NullString{Str: "", Valid: true},
			fallback: "default",
			expected: "",
		},
		{
			name:     "returns fallback when invalid",
			input:    NullString{Str: "hello", Valid: false},
			fallback: "default",
			expected: "default",
		},
		{
			name:     "returns empty fallback when invalid",
			input:    NullString{Str: "hello", Valid: false},
			fallback: "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.input.ValueOr(tt.fallback)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestNullString_String(t *testing.T) {
	tests := []struct {
		name     string
		input    NullString
		expected string
	}{
		{
			name:     "returns string when valid",
			input:    NullString{Str: "hello", Valid: true},
			expected: "hello",
		},
		{
			name:     "returns empty string when valid",
			input:    NullString{Str: "", Valid: true},
			expected: "",
		},
		{
			name:     "returns null when invalid (with string value)",
			input:    NullString{Str: "hello", Valid: false},
			expected: "null",
		},
		{
			name:     "returns null when invalid (with empty value)",
			input:    NullString{Str: "", Valid: false},
			expected: "null",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.input.String()
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestNewNullString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected NullString
	}{
		{
			name:     "creates valid NullString with string",
			input:    "hello",
			expected: NullString{Str: "hello", Valid: true},
		},
		{
			name:     "creates valid NullString with empty string",
			input:    "",
			expected: NullString{Str: "", Valid: true},
		},
		{
			name:     "creates valid NullString with numeric string",
			input:    "123",
			expected: NullString{Str: "123", Valid: true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewNullString(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestNullString_ActualNull(t *testing.T) {
	// Test actual JSON null (not string "null")
	jsonData := `{"value": null}`

	var result struct {
		Value NullString `json:"value"`
	}

	err := json.Unmarshal([]byte(jsonData), &result)
	require.NoError(t, err)

	// Should be invalid (null)
	assert.False(t, result.Value.Valid, "Expected null value to be invalid")
	assert.Equal(t, "", result.Value.Str, "Expected empty string for null")

	// Also test direct unmarshaling of null
	var nullString NullString
	err = json.Unmarshal([]byte("null"), &nullString)
	require.NoError(t, err)
	assert.False(t, nullString.Valid, "Direct null should be invalid")
	assert.Equal(t, "", nullString.Str, "Direct null should have empty string")

	// Test string "null" (this should also work)
	var stringNullString NullString
	err = json.Unmarshal([]byte(`"null"`), &stringNullString)
	require.NoError(t, err)
	assert.False(t, stringNullString.Valid, "String null should be invalid")
	assert.Equal(t, "", stringNullString.Str, "String null should have empty string")

	// Test nil data slice (edge case)
	var nilDataString NullString
	err = nilDataString.UnmarshalJSON(nil)
	require.NoError(t, err)
	assert.False(t, nilDataString.Valid, "Nil data should be invalid")
	assert.Equal(t, "", nilDataString.Str, "Nil data should have empty string")
}

func TestNullString_RoundTrip(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{
			name:  "string",
			input: `"hello"`,
		},
		{
			name:  "numeric string",
			input: `"123"`,
		},
		{
			name:  "integer",
			input: "42",
		},
		{
			name:  "float",
			input: "3.14",
		},
		{
			name:  "negative number",
			input: "-10",
		},
		{
			name:  "scientific notation",
			input: "1.23e+10",
		},
		{
			name:  "zero",
			input: "0",
		},
		{
			name:  "null",
			input: "null",
		},
		{
			name:  "empty string",
			input: `""`,
		},
		{
			name:  "string null",
			input: `"null"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var ns NullString
			err := json.Unmarshal([]byte(tt.input), &ns)
			require.NoError(t, err)

			// Marshal back to JSON
			marshaled, err := json.Marshal(ns)
			require.NoError(t, err)

			// For null cases, we expect "null" output
			if tt.input == "null" || tt.input == `""` || tt.input == `"null"` {
				assert.Equal(t, "null", string(marshaled))
			} else {
				// For valid cases, unmarshal again to verify round-trip
				var ns2 NullString
				err = json.Unmarshal(marshaled, &ns2)
				require.NoError(t, err)

				assert.Equal(t, ns.Str, ns2.Str)
				assert.Equal(t, ns.Valid, ns2.Valid)
			}
		})
	}
}

func TestNullString_InStruct(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected struct {
			Name     NullString `json:"name"`
			Email    NullString `json:"email"`
			Age      NullString `json:"age"`
			IsActive NullString `json:"is_active"`
		}
	}{
		{
			name:  "unmarshals in struct with string values",
			input: `{"name": "John", "email": "john@example.com", "age": "25", "is_active": "true"}`,
			expected: struct {
				Name     NullString `json:"name"`
				Email    NullString `json:"email"`
				Age      NullString `json:"age"`
				IsActive NullString `json:"is_active"`
			}{
				Name:     NullString{Str: "John", Valid: true},
				Email:    NullString{Str: "john@example.com", Valid: true},
				Age:      NullString{Str: "25", Valid: true},
				IsActive: NullString{Str: "true", Valid: true},
			},
		},
		{
			name:  "unmarshals in struct with number values",
			input: `{"name": "Jane", "email": "jane@example.com", "age": 30, "is_active": 1}`,
			expected: struct {
				Name     NullString `json:"name"`
				Email    NullString `json:"email"`
				Age      NullString `json:"age"`
				IsActive NullString `json:"is_active"`
			}{
				Name:     NullString{Str: "Jane", Valid: true},
				Email:    NullString{Str: "jane@example.com", Valid: true},
				Age:      NullString{Str: "30", Valid: true},
				IsActive: NullString{Str: "1", Valid: true},
			},
		},
		{
			name:  "unmarshals in struct with null values",
			input: `{"name": "Bob", "email": null, "age": "", "is_active": "null"}`,
			expected: struct {
				Name     NullString `json:"name"`
				Email    NullString `json:"email"`
				Age      NullString `json:"age"`
				IsActive NullString `json:"is_active"`
			}{
				Name:     NullString{Str: "Bob", Valid: true},
				Email:    NullString{Str: "", Valid: false},
				Age:      NullString{Str: "", Valid: false},
				IsActive: NullString{Str: "", Valid: false},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result struct {
				Name     NullString `json:"name"`
				Email    NullString `json:"email"`
				Age      NullString `json:"age"`
				IsActive NullString `json:"is_active"`
			}

			err := json.Unmarshal([]byte(tt.input), &result)
			require.NoError(t, err)
			assert.Equal(t, tt.expected, result)
		})
	}

	// Test marshaling from struct
	t.Run("marshals valid values from struct", func(t *testing.T) {
		input := struct {
			Name  NullString `json:"name"`
			Email NullString `json:"email"`
		}{
			Name:  NullString{Str: "Alice", Valid: true},
			Email: NullString{Str: "alice@example.com", Valid: true},
		}

		result, err := json.Marshal(input)
		require.NoError(t, err)
		assert.JSONEq(t, `{"name": "Alice", "email": "alice@example.com"}`, string(result))
	})

	t.Run("marshals null values from struct", func(t *testing.T) {
		input := struct {
			Name  NullString `json:"name"`
			Email NullString `json:"email"`
		}{
			Name:  NullString{Str: "Charlie", Valid: true},
			Email: NullString{Str: "", Valid: false},
		}

		result, err := json.Marshal(input)
		require.NoError(t, err)
		assert.JSONEq(t, `{"name": "Charlie", "email": null}`, string(result))
	})
}

func TestNullString_CompareWithMultiString(t *testing.T) {
	t.Run("null handling difference", func(t *testing.T) {
		jsonData := `{"multi": null, "null": null}`

		var result struct {
			// Multi MultiString `json:"multi"` (removed)
			Null NullString `json:"null"`
		}

		err := json.Unmarshal([]byte(jsonData), &result)
		require.NoError(t, err)

		// MultiString converts null to empty string
		// MultiString tests removed

		// NullString preserves null state
		assert.False(t, result.Null.Valid)
		assert.Equal(t, "", result.Null.Str)
		assert.Equal(t, "null", result.Null.String())
	})

	t.Run("empty string handling difference", func(t *testing.T) {
		jsonData := `{"multi": "", "null": ""}`

		var result struct {
			// Multi MultiString `json:"multi"` (removed)
			Null NullString `json:"null"`
		}

		err := json.Unmarshal([]byte(jsonData), &result)
		require.NoError(t, err)

		// MultiString keeps empty string as is
		// MultiString tests removed

		// NullString treats empty string as null
		assert.False(t, result.Null.Valid)
		assert.Equal(t, "", result.Null.Str)
		assert.Equal(t, "null", result.Null.String())
	})
}
