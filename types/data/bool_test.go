package data

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBool_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
		wantErr  bool
	}{
		// Boolean inputs
		{
			name:     "unmarshals true",
			input:    "true",
			expected: true,
		},
		{
			name:     "unmarshals false",
			input:    "false",
			expected: false,
		},
		// String inputs
		{
			name:     "unmarshals string true",
			input:    `"true"`,
			expected: true,
		},
		{
			name:     "unmarshals string false",
			input:    `"false"`,
			expected: false,
		},
		{
			name:     "unmarshals string 1 as true",
			input:    `"1"`,
			expected: true,
		},
		{
			name:     "unmarshals string 0 as false",
			input:    `"0"`,
			expected: false,
		},
		{
			name:     "unmarshals string success as true",
			input:    `"success"`,
			expected: true,
		},
		{
			name:     "unmarshals string failure as false",
			input:    `"failure"`,
			expected: false,
		},
		{
			name:     "unmarshals string SUCCESS as true (case insensitive)",
			input:    `"SUCCESS"`,
			expected: true,
		},
		{
			name:     "unmarshals string TRUE as true (case insensitive)",
			input:    `"TRUE"`,
			expected: true,
		},
		{
			name:     "unmarshals string FALSE as false (case insensitive)",
			input:    `"FALSE"`,
			expected: false,
		},
		{
			name:     "unmarshals empty string as false",
			input:    `""`,
			expected: false,
		},
		{
			name:     "unmarshals string null as false",
			input:    `"null"`,
			expected: false,
		},
		// Numeric string inputs
		{
			name:     "unmarshals string positive number as true",
			input:    `"42"`,
			expected: true,
		},
		{
			name:     "unmarshals string negative number as true",
			input:    `"-10"`,
			expected: true,
		},
		{
			name:     "unmarshals string zero as false",
			input:    `"0.0"`,
			expected: false,
		},
		{
			name:     "unmarshals string float as true",
			input:    `"3.14"`,
			expected: true,
		},
		// Numeric inputs
		{
			name:     "unmarshals positive integer as true",
			input:    "42",
			expected: true,
		},
		{
			name:     "unmarshals negative integer as true",
			input:    "-10",
			expected: true,
		},
		{
			name:     "unmarshals zero as false",
			input:    "0",
			expected: false,
		},
		{
			name:     "unmarshals positive float as true",
			input:    "3.14",
			expected: true,
		},
		{
			name:     "unmarshals negative float as true",
			input:    "-2.5",
			expected: true,
		},
		{
			name:     "unmarshals zero float as false",
			input:    "0.0",
			expected: false,
		},
		// Null input
		{
			name:     "unmarshals null as false",
			input:    "null",
			expected: false,
		},
		// Error cases
		{
			name:    "fails to unmarshal invalid string",
			input:   `"maybe"`,
			wantErr: true,
		},
		{
			name:    "fails to unmarshal object",
			input:   `{"key": "value"}`,
			wantErr: true,
		},
		{
			name:    "fails to unmarshal array",
			input:   `[true, false]`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var b Bool
			err := json.Unmarshal([]byte(tt.input), &b)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.expected, b.Bool())
		})
	}
}

func TestBool_UnmarshalJSON_NilData(t *testing.T) {
	var b Bool
	err := b.UnmarshalJSON(nil)
	require.NoError(t, err)
	assert.Equal(t, false, b.Bool())
}

func TestBool_MarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    Bool
		expected string
	}{
		{
			name:     "marshals true",
			input:    Bool(true),
			expected: "true",
		},
		{
			name:     "marshals false",
			input:    Bool(false),
			expected: "false",
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

func TestBool_Bool(t *testing.T) {
	tests := []struct {
		name     string
		input    Bool
		expected bool
	}{
		{
			name:     "returns true",
			input:    Bool(true),
			expected: true,
		},
		{
			name:     "returns false",
			input:    Bool(false),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.input.Bool())
		})
	}
}

func TestBool_String(t *testing.T) {
	tests := []struct {
		name     string
		input    Bool
		expected string
	}{
		{
			name:     "returns true string",
			input:    Bool(true),
			expected: "true",
		},
		{
			name:     "returns false string",
			input:    Bool(false),
			expected: "false",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.input.String())
		})
	}
}

func TestNewBool(t *testing.T) {
	tests := []struct {
		name     string
		input    bool
		expected Bool
	}{
		{
			name:     "creates new Bool with true",
			input:    true,
			expected: Bool(true),
		},
		{
			name:     "creates new Bool with false",
			input:    false,
			expected: Bool(false),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewBool(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestBool_RoundTrip(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{
			name:  "true boolean",
			input: "true",
		},
		{
			name:  "false boolean",
			input: "false",
		},
		{
			name:  "string true",
			input: `"true"`,
		},
		{
			name:  "string false",
			input: `"false"`,
		},
		{
			name:  "string 1",
			input: `"1"`,
		},
		{
			name:  "string 0",
			input: `"0"`,
		},
		{
			name:  "string success",
			input: `"success"`,
		},
		{
			name:  "string failure",
			input: `"failure"`,
		},
		{
			name:  "positive integer",
			input: "42",
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
			// Unmarshal the input
			var b Bool
			err := json.Unmarshal([]byte(tt.input), &b)
			require.NoError(t, err)

			// Marshal it back
			data, err := json.Marshal(b)
			require.NoError(t, err)

			// Unmarshal again to verify consistency
			var b2 Bool
			err = json.Unmarshal(data, &b2)
			require.NoError(t, err)

			// Should be the same
			assert.Equal(t, b.Bool(), b2.Bool())
		})
	}
}

func TestBool_InStruct(t *testing.T) {
	type TestStruct struct {
		Value Bool `json:"value"`
	}

	t.Run("unmarshals in struct with string value", func(t *testing.T) {
		input := `{"value": "true"}`
		var ts TestStruct
		err := json.Unmarshal([]byte(input), &ts)
		require.NoError(t, err)
		assert.Equal(t, true, ts.Value.Bool())
	})

	t.Run("unmarshals in struct with boolean value", func(t *testing.T) {
		input := `{"value": false}`
		var ts TestStruct
		err := json.Unmarshal([]byte(input), &ts)
		require.NoError(t, err)
		assert.Equal(t, false, ts.Value.Bool())
	})

	t.Run("unmarshals in struct with numeric value", func(t *testing.T) {
		input := `{"value": 1}`
		var ts TestStruct
		err := json.Unmarshal([]byte(input), &ts)
		require.NoError(t, err)
		assert.Equal(t, true, ts.Value.Bool())
	})

	t.Run("unmarshals in struct with null value", func(t *testing.T) {
		input := `{"value": null}`
		var ts TestStruct
		err := json.Unmarshal([]byte(input), &ts)
		require.NoError(t, err)
		assert.Equal(t, false, ts.Value.Bool())
	})

	t.Run("marshals from struct", func(t *testing.T) {
		ts := TestStruct{Value: Bool(true)}
		data, err := json.Marshal(ts)
		require.NoError(t, err)
		assert.Contains(t, string(data), `"value":true`)
	})
}
