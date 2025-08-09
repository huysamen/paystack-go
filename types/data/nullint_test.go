package data

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNullInt_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expectedInt   int64
		expectedValid bool
		wantErr       bool
	}{
		// Integer inputs
		{
			name:          "unmarshals integer",
			input:         "42",
			expectedInt:   42,
			expectedValid: true,
		},
		{
			name:          "unmarshals negative integer",
			input:         "-10",
			expectedInt:   -10,
			expectedValid: true,
		},
		{
			name:          "unmarshals zero",
			input:         "0",
			expectedInt:   0,
			expectedValid: true,
		},
		// Float inputs (should be truncated)
		{
			name:          "unmarshals float as integer (truncated)",
			input:         "42.9",
			expectedInt:   42,
			expectedValid: true,
		},
		{
			name:          "unmarshals negative float as integer (truncated)",
			input:         "-78.3",
			expectedInt:   -78,
			expectedValid: true,
		},
		// String inputs
		{
			name:          "unmarshals string integer",
			input:         `"123"`,
			expectedInt:   123,
			expectedValid: true,
		},
		{
			name:          "unmarshals negative string integer",
			input:         `"-456"`,
			expectedInt:   -456,
			expectedValid: true,
		},
		{
			name:          "unmarshals string float as integer",
			input:         `"78.9"`,
			expectedInt:   78,
			expectedValid: true,
		},
		{
			name:          "unmarshals empty string as null",
			input:         `""`,
			expectedInt:   0,
			expectedValid: false,
		},
		// Null input
		{
			name:          "unmarshals null as invalid",
			input:         "null",
			expectedInt:   0,
			expectedValid: false,
		},
		// Error cases
		{
			name:    "fails to unmarshal invalid string",
			input:   `"not-a-number"`,
			wantErr: true,
		},
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
			var ni NullInt
			err := json.Unmarshal([]byte(tt.input), &ni)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.expectedInt, ni.Int)
			assert.Equal(t, tt.expectedValid, ni.Valid)
		})
	}
}

func TestNullInt_MarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    NullInt
		expected string
	}{
		{
			name:     "marshals valid positive integer",
			input:    NullInt{Int: 42, Valid: true},
			expected: "42",
		},
		{
			name:     "marshals valid negative integer",
			input:    NullInt{Int: -10, Valid: true},
			expected: "-10",
		},
		{
			name:     "marshals valid zero",
			input:    NullInt{Int: 0, Valid: true},
			expected: "0",
		},
		{
			name:     "marshals invalid as null",
			input:    NullInt{Int: 123, Valid: false},
			expected: "null",
		},
		{
			name:     "marshals zero invalid as null",
			input:    NullInt{Int: 0, Valid: false},
			expected: "null",
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

func TestNullInt_ValueOr(t *testing.T) {
	tests := []struct {
		name     string
		input    NullInt
		fallback int64
		expected int64
	}{
		{
			name:     "returns value when valid",
			input:    NullInt{Int: 123, Valid: true},
			fallback: 999,
			expected: 123,
		},
		{
			name:     "returns zero when valid",
			input:    NullInt{Int: 0, Valid: true},
			fallback: 999,
			expected: 0,
		},
		{
			name:     "returns fallback when invalid",
			input:    NullInt{Int: 123, Valid: false},
			fallback: 999,
			expected: 999,
		},
		{
			name:     "returns fallback when invalid with zero",
			input:    NullInt{Int: 0, Valid: false},
			fallback: 456,
			expected: 456,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.input.ValueOr(tt.fallback))
		})
	}
}

func TestNullInt_String(t *testing.T) {
	tests := []struct {
		name     string
		input    NullInt
		expected string
	}{
		{
			name:     "returns string representation when valid",
			input:    NullInt{Int: 123, Valid: true},
			expected: "123",
		},
		{
			name:     "returns negative string representation when valid",
			input:    NullInt{Int: -456, Valid: true},
			expected: "-456",
		},
		{
			name:     "returns zero string when valid",
			input:    NullInt{Int: 0, Valid: true},
			expected: "0",
		},
		{
			name:     "returns null when invalid",
			input:    NullInt{Int: 123, Valid: false},
			expected: "null",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.input.String())
		})
	}
}

func TestNewNullInt(t *testing.T) {
	tests := []struct {
		name  string
		input int64
	}{
		{
			name:  "creates valid NullInt with positive value",
			input: 123,
		},
		{
			name:  "creates valid NullInt with negative value",
			input: -456,
		},
		{
			name:  "creates valid NullInt with zero",
			input: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ni := NewNullInt(tt.input)
			assert.Equal(t, tt.input, ni.Int)
			assert.True(t, ni.Valid)
		})
	}
}

func TestNullInt_ActualNull(t *testing.T) {
	// Test actual JSON null (not string "null")
	jsonData := `{"value": null}`

	var result struct {
		Value NullInt `json:"value"`
	}

	err := json.Unmarshal([]byte(jsonData), &result)
	require.NoError(t, err)

	// Should be invalid (null)
	assert.False(t, result.Value.Valid, "Expected null value to be invalid")
	assert.Equal(t, int64(0), result.Value.Int, "Expected zero value for null")

	// Also test direct unmarshaling of null
	var nullInt NullInt
	err = json.Unmarshal([]byte("null"), &nullInt)
	require.NoError(t, err)
	assert.False(t, nullInt.Valid, "Direct null should be invalid")
	assert.Equal(t, int64(0), nullInt.Int, "Direct null should have zero value")

	// Test string "null" (this should also work)
	var stringNullInt NullInt
	err = json.Unmarshal([]byte(`"null"`), &stringNullInt)
	require.NoError(t, err)
	assert.False(t, stringNullInt.Valid, "String null should be invalid")
	assert.Equal(t, int64(0), stringNullInt.Int, "String null should have zero value")

	// Test nil data slice (edge case)
	var nilDataInt NullInt
	err = nilDataInt.UnmarshalJSON(nil)
	require.NoError(t, err)
	assert.False(t, nilDataInt.Valid, "Nil data should be invalid")
	assert.Equal(t, int64(0), nilDataInt.Int, "Nil data should have zero value")
}

func TestNullInt_RoundTrip(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{
			name:  "integer",
			input: "42",
		},
		{
			name:  "string integer",
			input: `"123"`,
		},
		{
			name:  "string float",
			input: `"456.7"`,
		},
		{
			name:  "float",
			input: "789.1",
		},
		{
			name:  "negative integer",
			input: "-10",
		},
		{
			name:  "negative string",
			input: `"-20"`,
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Unmarshal the input
			var ni NullInt
			err := json.Unmarshal([]byte(tt.input), &ni)
			require.NoError(t, err)

			// Marshal it back
			data, err := json.Marshal(ni)
			require.NoError(t, err)

			// Unmarshal again to verify consistency
			var ni2 NullInt
			err = json.Unmarshal(data, &ni2)
			require.NoError(t, err)

			// Should be the same
			assert.Equal(t, ni.Int, ni2.Int)
			assert.Equal(t, ni.Valid, ni2.Valid)
		})
	}
}

func TestNullInt_InStruct(t *testing.T) {
	type TestStruct struct {
		Value NullInt `json:"value"`
	}

	t.Run("unmarshals in struct with string value", func(t *testing.T) {
		input := `{"value": "42"}`
		var ts TestStruct
		err := json.Unmarshal([]byte(input), &ts)
		require.NoError(t, err)
		assert.Equal(t, int64(42), ts.Value.Int)
		assert.True(t, ts.Value.Valid)
	})

	t.Run("unmarshals in struct with integer value", func(t *testing.T) {
		input := `{"value": 123}`
		var ts TestStruct
		err := json.Unmarshal([]byte(input), &ts)
		require.NoError(t, err)
		assert.Equal(t, int64(123), ts.Value.Int)
		assert.True(t, ts.Value.Valid)
	})

	t.Run("unmarshals in struct with null value", func(t *testing.T) {
		input := `{"value": null}`
		var ts TestStruct
		err := json.Unmarshal([]byte(input), &ts)
		require.NoError(t, err)
		assert.Equal(t, int64(0), ts.Value.Int)
		assert.False(t, ts.Value.Valid)
	})

	t.Run("marshals valid value from struct", func(t *testing.T) {
		ts := TestStruct{Value: NullInt{Int: 456, Valid: true}}
		data, err := json.Marshal(ts)
		require.NoError(t, err)
		assert.Contains(t, string(data), `"value":456`)
	})

	t.Run("marshals null value from struct", func(t *testing.T) {
		ts := TestStruct{Value: NullInt{Int: 456, Valid: false}}
		data, err := json.Marshal(ts)
		require.NoError(t, err)
		assert.Contains(t, string(data), `"value":null`)
	})
}

func TestNullInt_CompareWithMultiInt(t *testing.T) {
	t.Run("null handling difference", func(t *testing.T) {
		input := "null"

		var ni NullInt
		var mi MultiInt

		err := json.Unmarshal([]byte(input), &ni)
		require.NoError(t, err)
		err = json.Unmarshal([]byte(input), &mi)
		require.NoError(t, err)

		// NullInt preserves null state
		assert.False(t, ni.Valid)
		assert.Equal(t, int64(0), ni.Int)

		// MultiInt converts null to zero
		assert.Equal(t, 0, mi.Int())

		// But when marshaling back, they differ
		niData, err := json.Marshal(ni)
		require.NoError(t, err)
		assert.Equal(t, "null", string(niData))

		miData, err := json.Marshal(mi)
		require.NoError(t, err)
		assert.Equal(t, "0", string(miData))
	})
}
