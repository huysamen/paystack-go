package data

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNullUint_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expectedVal uint64
		expectedOK  bool
		hasError    bool
	}{
		{
			name:        "unmarshals uint",
			input:       `123`,
			expectedVal: 123,
			expectedOK:  true,
		},
		{
			name:        "unmarshals zero",
			input:       `0`,
			expectedVal: 0,
			expectedOK:  true,
		},
		{
			name:        "unmarshals large uint",
			input:       `18446744073709551615`, // max uint64
			expectedVal: 18446744073709551615,
			expectedOK:  true,
		},
		{
			name:        "unmarshals float as uint (truncated)",
			input:       `123.78`,
			expectedVal: 123,
			expectedOK:  true,
		},
		{
			name:        "unmarshals string uint",
			input:       `"456"`,
			expectedVal: 456,
			expectedOK:  true,
		},
		{
			name:        "unmarshals string float as uint",
			input:       `"789.45"`,
			expectedVal: 789,
			expectedOK:  true,
		},
		{
			name:        "unmarshals empty string as null",
			input:       `""`,
			expectedVal: 0,
			expectedOK:  false,
		},
		{
			name:        "unmarshals null as invalid",
			input:       `null`,
			expectedVal: 0,
			expectedOK:  false,
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
			var nu NullUint
			err := json.Unmarshal([]byte(tt.input), &nu)

			if tt.hasError {
				assert.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expectedVal, nu.Uint, "uint value should match")
				assert.Equal(t, tt.expectedOK, nu.Valid, "valid flag should match")
			}
		})
	}
}

func TestNullUint_MarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    NullUint
		expected string
	}{
		{
			name:     "marshals valid positive uint",
			input:    NullUint{Uint: 123, Valid: true},
			expected: `123`,
		},
		{
			name:     "marshals valid zero",
			input:    NullUint{Uint: 0, Valid: true},
			expected: `0`,
		},
		{
			name:     "marshals valid large uint",
			input:    NullUint{Uint: 18446744073709551615, Valid: true},
			expected: `18446744073709551615`,
		},
		{
			name:     "marshals invalid as null",
			input:    NullUint{Uint: 123, Valid: false},
			expected: `null`,
		},
		{
			name:     "marshals zero invalid as null",
			input:    NullUint{Uint: 0, Valid: false},
			expected: `null`,
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

func TestNullUint_ValueOr(t *testing.T) {
	tests := []struct {
		name     string
		input    NullUint
		fallback uint64
		expected uint64
	}{
		{
			name:     "returns value when valid",
			input:    NullUint{Uint: 123, Valid: true},
			fallback: 999,
			expected: 123,
		},
		{
			name:     "returns zero when valid",
			input:    NullUint{Uint: 0, Valid: true},
			fallback: 999,
			expected: 0,
		},
		{
			name:     "returns fallback when invalid",
			input:    NullUint{Uint: 123, Valid: false},
			fallback: 456,
			expected: 456,
		},
		{
			name:     "returns fallback when invalid with zero",
			input:    NullUint{Uint: 0, Valid: false},
			fallback: 789,
			expected: 789,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.input.ValueOr(tt.fallback)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestNullUint_String(t *testing.T) {
	tests := []struct {
		name     string
		input    NullUint
		expected string
	}{
		{
			name:     "returns string representation when valid",
			input:    NullUint{Uint: 123, Valid: true},
			expected: "123",
		},
		{
			name:     "returns zero string when valid",
			input:    NullUint{Uint: 0, Valid: true},
			expected: "0",
		},
		{
			name:     "returns large uint string when valid",
			input:    NullUint{Uint: 18446744073709551615, Valid: true},
			expected: "18446744073709551615",
		},
		{
			name:     "returns null when invalid",
			input:    NullUint{Uint: 123, Valid: false},
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

func TestNewNullUint(t *testing.T) {
	tests := []struct {
		name     string
		input    uint64
		expected NullUint
	}{
		{
			name:     "creates valid NullUint with positive value",
			input:    123,
			expected: NullUint{Uint: 123, Valid: true},
		},
		{
			name:     "creates valid NullUint with zero",
			input:    0,
			expected: NullUint{Uint: 0, Valid: true},
		},
		{
			name:     "creates valid NullUint with large value",
			input:    18446744073709551615,
			expected: NullUint{Uint: 18446744073709551615, Valid: true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewNullUint(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestNullUint_ActualNull(t *testing.T) {
	var nu NullUint
	jsonBytes, err := json.Marshal(nu)
	require.NoError(t, err)
	assert.Equal(t, `null`, string(jsonBytes), "uninitialized NullUint should marshal to null")

	var nu2 NullUint
	err = json.Unmarshal([]byte(`null`), &nu2)
	require.NoError(t, err)
	assert.False(t, nu2.Valid, "null should unmarshal to invalid NullUint")
	assert.Equal(t, uint64(0), nu2.Uint, "null should unmarshal to zero uint value")
}

func TestNullUint_RoundTrip(t *testing.T) {
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
			var nu NullUint
			err := json.Unmarshal([]byte(tt.input), &nu)
			require.NoError(t, err)

			marshaled, err := json.Marshal(nu)
			require.NoError(t, err)

			var nu2 NullUint
			err = json.Unmarshal(marshaled, &nu2)
			require.NoError(t, err)

			assert.Equal(t, nu, nu2, "round trip should preserve value and validity")
		})
	}
}

func TestNullUint_InStruct(t *testing.T) {
	type TestStruct struct {
		Value NullUint `json:"value"`
		Count NullUint `json:"count"`
	}

	t.Run("unmarshals in struct with string value", func(t *testing.T) {
		jsonData := `{"value": "123", "count": "456"}`

		var ts TestStruct
		err := json.Unmarshal([]byte(jsonData), &ts)
		require.NoError(t, err)

		assert.Equal(t, uint64(123), ts.Value.Uint)
		assert.True(t, ts.Value.Valid)
		assert.Equal(t, uint64(456), ts.Count.Uint)
		assert.True(t, ts.Count.Valid)
	})

	t.Run("unmarshals in struct with uint value", func(t *testing.T) {
		jsonData := `{"value": 789, "count": 101112}`

		var ts TestStruct
		err := json.Unmarshal([]byte(jsonData), &ts)
		require.NoError(t, err)

		assert.Equal(t, uint64(789), ts.Value.Uint)
		assert.True(t, ts.Value.Valid)
		assert.Equal(t, uint64(101112), ts.Count.Uint)
		assert.True(t, ts.Count.Valid)
	})

	t.Run("unmarshals in struct with null value", func(t *testing.T) {
		jsonData := `{"value": null, "count": null}`

		var ts TestStruct
		err := json.Unmarshal([]byte(jsonData), &ts)
		require.NoError(t, err)

		assert.False(t, ts.Value.Valid)
		assert.False(t, ts.Count.Valid)
	})

	t.Run("marshals valid value from struct", func(t *testing.T) {
		ts := TestStruct{
			Value: NewNullUint(123),
			Count: NewNullUint(456),
		}

		jsonBytes, err := json.Marshal(ts)
		require.NoError(t, err)

		expected := `{"value":123,"count":456}`
		assert.Equal(t, expected, string(jsonBytes))
	})

	t.Run("marshals null value from struct", func(t *testing.T) {
		ts := TestStruct{
			Value: NullUint{Valid: false},
			Count: NullUint{Valid: false},
		}

		jsonBytes, err := json.Marshal(ts)
		require.NoError(t, err)

		expected := `{"value":null,"count":null}`
		assert.Equal(t, expected, string(jsonBytes))
	})
}

func TestNullUint_CompareWithUint(t *testing.T) {
	t.Run("null handling difference", func(t *testing.T) {
		jsonData := `null`

		var u Uint
		var nu NullUint

		err := json.Unmarshal([]byte(jsonData), &u)
		require.NoError(t, err)
		assert.Equal(t, Uint(0), u, "Uint should convert null to 0")

		err = json.Unmarshal([]byte(jsonData), &nu)
		require.NoError(t, err)
		assert.False(t, nu.Valid, "NullUint should preserve null state")
		assert.Equal(t, uint64(0), nu.Uint, "NullUint should have zero value")

		// Verify marshaling behavior
		uMarshaled, err := json.Marshal(u)
		require.NoError(t, err)
		assert.Equal(t, `0`, string(uMarshaled), "Uint should marshal null as 0")

		nuMarshaled, err := json.Marshal(nu)
		require.NoError(t, err)
		assert.Equal(t, `null`, string(nuMarshaled), "NullUint should marshal null as null")
	})
}
