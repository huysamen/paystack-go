package data

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestString_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		wantErr  bool
	}{
		// String inputs
		{
			name:     "unmarshals string",
			input:    `"hello"`,
			expected: "hello",
		},
		{
			name:     "unmarshals empty string",
			input:    `""`,
			expected: "",
		},
		{
			name:     "unmarshals string with spaces",
			input:    `"hello world"`,
			expected: "hello world",
		},
		{
			name:     "unmarshals string with special characters",
			input:    `"hello\nworld\t!"`,
			expected: "hello\nworld\t!",
		},
		{
			name:     "unmarshals string null as empty",
			input:    `"null"`,
			expected: "",
		},
		// Number inputs (should be converted to string)
		{
			name:     "unmarshals integer as string",
			input:    "42",
			expected: "42",
		},
		{
			name:     "unmarshals negative integer as string",
			input:    "-10",
			expected: "-10",
		},
		{
			name:     "unmarshals float as string",
			input:    "123.45",
			expected: "123.45",
		},
		{
			name:     "unmarshals negative float as string",
			input:    "-456.78",
			expected: "-456.78",
		},
		{
			name:     "unmarshals zero as string",
			input:    "0",
			expected: "0",
		},
		{
			name:     "unmarshals scientific notation as string",
			input:    "1.23e+4",
			expected: "12300",
		},
		// Null input
		{
			name:     "unmarshals null as empty string",
			input:    "null",
			expected: "",
		},
		// Note: boolean, object, array inputs will cause errors as expected
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s String
			err := json.Unmarshal([]byte(tt.input), &s)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.expected, s.String())
		})
	}
}

func TestString_UnmarshalJSON_NilData(t *testing.T) {
	var s String
	err := s.UnmarshalJSON(nil)
	require.NoError(t, err)
	assert.Equal(t, "", s.String())
}

func TestString_MarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    String
		expected string
	}{
		{
			name:     "marshals string",
			input:    String("hello"),
			expected: `"hello"`,
		},
		{
			name:     "marshals empty string",
			input:    String(""),
			expected: `""`,
		},
		{
			name:     "marshals string with spaces",
			input:    String("hello world"),
			expected: `"hello world"`,
		},
		{
			name:     "marshals string with special characters",
			input:    String("hello\nworld\t!"),
			expected: `"hello\nworld\t!"`,
		},
		{
			name:     "marshals numeric string",
			input:    String("123"),
			expected: `"123"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := json.Marshal(tt.input)
			require.NoError(t, err)
			assert.Equal(t, tt.expected, string(data))
		})
	}
}

func TestString_String(t *testing.T) {
	tests := []struct {
		name     string
		input    String
		expected string
	}{
		{
			name:     "returns correct string value",
			input:    String("hello"),
			expected: "hello",
		},
		{
			name:     "returns empty string",
			input:    String(""),
			expected: "",
		},
		{
			name:     "returns string with spaces",
			input:    String("hello world"),
			expected: "hello world",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.input.String())
		})
	}
}

func TestNewString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected String
	}{
		{
			name:     "creates new String with value",
			input:    "hello",
			expected: String("hello"),
		},
		{
			name:     "creates new String with empty value",
			input:    "",
			expected: String(""),
		},
		{
			name:     "creates new String with spaces",
			input:    "hello world",
			expected: String("hello world"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewString(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestString_RoundTrip(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{
			name:  "string",
			input: `"hello"`,
		},
		{
			name:  "empty string",
			input: `""`,
		},
		{
			name:  "integer",
			input: "42",
		},
		{
			name:  "float",
			input: "123.45",
		},
		{
			name:  "negative number",
			input: "-10",
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
			name:  "string null",
			input: `"null"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Unmarshal the input
			var s String
			err := json.Unmarshal([]byte(tt.input), &s)
			require.NoError(t, err)

			// Marshal it back
			data, err := json.Marshal(s)
			require.NoError(t, err)

			// Unmarshal again to verify consistency
			var s2 String
			err = json.Unmarshal(data, &s2)
			require.NoError(t, err)

			// Should be the same
			assert.Equal(t, s.String(), s2.String())
		})
	}
}

func TestString_InStruct(t *testing.T) {
	type TestStruct struct {
		Value String `json:"value"`
	}

	t.Run("unmarshals in struct with string value", func(t *testing.T) {
		input := `{"value": "hello"}`
		var ts TestStruct
		err := json.Unmarshal([]byte(input), &ts)
		require.NoError(t, err)
		assert.Equal(t, "hello", ts.Value.String())
	})

	t.Run("unmarshals in struct with number value", func(t *testing.T) {
		input := `{"value": 123}`
		var ts TestStruct
		err := json.Unmarshal([]byte(input), &ts)
		require.NoError(t, err)
		assert.Equal(t, "123", ts.Value.String())
	})

	t.Run("unmarshals in struct with null value", func(t *testing.T) {
		input := `{"value": null}`
		var ts TestStruct
		err := json.Unmarshal([]byte(input), &ts)
		require.NoError(t, err)
		assert.Equal(t, "", ts.Value.String())
	})

	t.Run("marshals from struct", func(t *testing.T) {
		ts := TestStruct{Value: String("world")}
		data, err := json.Marshal(ts)
		require.NoError(t, err)
		assert.Contains(t, string(data), `"value":"world"`)
	})
}

func TestString_SpecialNumberFormats(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "scientific notation",
			input:    "1.23e+4",
			expected: "12300",
		},
		{
			name:     "scientific notation negative exponent",
			input:    "1.23e-4",
			expected: "0.000123",
		},
		{
			name:     "large integer - precision may be lost in float conversion",
			input:    "9223372036854775807",
			expected: "9.223372036854776e+18", // Expected behavior due to float64 precision limits
		},
		{
			name:     "very small decimal",
			input:    "0.000001",
			expected: "1e-06",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s String
			err := json.Unmarshal([]byte(tt.input), &s)
			require.NoError(t, err)
			assert.Equal(t, tt.expected, s.String())
		})
	}
}
