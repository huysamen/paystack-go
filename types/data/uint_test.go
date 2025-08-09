package data

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUint_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected Uint
		hasError bool
	}{
		{
			name:     "unmarshals uint",
			input:    `123`,
			expected: 123,
		},
		{
			name:     "unmarshals zero",
			input:    `0`,
			expected: 0,
		},
		{
			name:     "unmarshals large uint",
			input:    `18446744073709551615`, // max uint64
			expected: 18446744073709551615,
		},
		{
			name:     "unmarshals float as uint (truncated)",
			input:    `123.78`,
			expected: 123,
		},
		{
			name:     "unmarshals string uint",
			input:    `"456"`,
			expected: 456,
		},
		{
			name:     "unmarshals string float as uint",
			input:    `"789.45"`,
			expected: 789,
		},
		{
			name:     "unmarshals empty string as zero",
			input:    `""`,
			expected: 0,
		},
		{
			name:     "unmarshals null as zero",
			input:    `null`,
			expected: 0,
		},
		{
			name:     "fails to unmarshal negative integer",
			input:    `-123`,
			hasError: true,
		},
		{
			name:     "fails to unmarshal negative float",
			input:    `-123.45`,
			hasError: true,
		},
		{
			name:     "fails to unmarshal negative string",
			input:    `"-456"`,
			hasError: true,
		},
		{
			name:     "fails to unmarshal invalid string",
			input:    `"invalid"`,
			hasError: true,
		},
		{
			name:     "fails to unmarshal boolean",
			input:    `true`,
			hasError: true,
		},
		{
			name:     "fails to unmarshal object",
			input:    `{"value": 123}`,
			hasError: true,
		},
		{
			name:     "fails to unmarshal array",
			input:    `[123]`,
			hasError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var u Uint
			err := json.Unmarshal([]byte(tt.input), &u)

			if tt.hasError {
				assert.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expected, u)
			}
		})
	}
}

func TestUint_MarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    Uint
		expected string
	}{
		{
			name:     "marshals positive uint",
			input:    Uint(123),
			expected: `123`,
		},
		{
			name:     "marshals zero",
			input:    Uint(0),
			expected: `0`,
		},
		{
			name:     "marshals large uint",
			input:    Uint(18446744073709551615),
			expected: `18446744073709551615`,
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

func TestUint_Uint64(t *testing.T) {
	tests := []struct {
		name     string
		input    Uint
		expected uint64
	}{
		{
			name:     "returns correct uint value",
			input:    Uint(123),
			expected: 123,
		},
		{
			name:     "returns zero",
			input:    Uint(0),
			expected: 0,
		},
		{
			name:     "returns large uint value",
			input:    Uint(18446744073709551615),
			expected: 18446744073709551615,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.input.Uint64()
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestUint_String(t *testing.T) {
	tests := []struct {
		name     string
		input    Uint
		expected string
	}{
		{
			name:     "returns correct string representation",
			input:    Uint(123),
			expected: "123",
		},
		{
			name:     "returns zero string",
			input:    Uint(0),
			expected: "0",
		},
		{
			name:     "returns large uint string",
			input:    Uint(18446744073709551615),
			expected: "18446744073709551615",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.input.String()
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestUint_RoundTrip(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{
			name:  "uint",
			input: `123`,
		},
		{
			name:  "string uint",
			input: `"456"`,
		},
		{
			name:  "string float",
			input: `"789.45"`,
		},
		{
			name:  "float",
			input: `123.78`,
		},
		{
			name:  "zero",
			input: `0`,
		},
		{
			name:  "null",
			input: `null`,
		},
		{
			name:  "empty string",
			input: `""`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var u Uint
			err := json.Unmarshal([]byte(tt.input), &u)
			require.NoError(t, err)

			marshaled, err := json.Marshal(u)
			require.NoError(t, err)

			var u2 Uint
			err = json.Unmarshal(marshaled, &u2)
			require.NoError(t, err)

			assert.Equal(t, u, u2, "round trip should preserve value")
		})
	}
}

func TestUint_InStruct(t *testing.T) {
	type TestStruct struct {
		Value Uint `json:"value"`
		Count Uint `json:"count"`
	}

	t.Run("unmarshals in struct with string value", func(t *testing.T) {
		jsonData := `{"value": "123", "count": "456"}`

		var ts TestStruct
		err := json.Unmarshal([]byte(jsonData), &ts)
		require.NoError(t, err)

		assert.Equal(t, Uint(123), ts.Value)
		assert.Equal(t, Uint(456), ts.Count)
	})

	t.Run("unmarshals in struct with uint value", func(t *testing.T) {
		jsonData := `{"value": 789, "count": 101112}`

		var ts TestStruct
		err := json.Unmarshal([]byte(jsonData), &ts)
		require.NoError(t, err)

		assert.Equal(t, Uint(789), ts.Value)
		assert.Equal(t, Uint(101112), ts.Count)
	})

	t.Run("unmarshals in struct with null value", func(t *testing.T) {
		jsonData := `{"value": null, "count": 0}`

		var ts TestStruct
		err := json.Unmarshal([]byte(jsonData), &ts)
		require.NoError(t, err)

		assert.Equal(t, Uint(0), ts.Value)
		assert.Equal(t, Uint(0), ts.Count)
	})

	t.Run("marshals from struct", func(t *testing.T) {
		ts := TestStruct{
			Value: Uint(123),
			Count: Uint(456),
		}

		jsonBytes, err := json.Marshal(ts)
		require.NoError(t, err)

		expected := `{"value":123,"count":456}`
		assert.Equal(t, expected, string(jsonBytes))
	})
}

func TestNewUint(t *testing.T) {
	tests := []struct {
		name     string
		input    uint64
		expected Uint
	}{
		{
			name:     "creates new Uint with value",
			input:    123,
			expected: Uint(123),
		},
		{
			name:     "creates new Uint with zero",
			input:    0,
			expected: Uint(0),
		},
		{
			name:     "creates new Uint with large value",
			input:    18446744073709551615,
			expected: Uint(18446744073709551615),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewUint(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
