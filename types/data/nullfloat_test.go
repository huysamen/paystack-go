package data

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNullFloat_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expectedFloat float64
		expectedValid bool
		wantErr       bool
	}{
		// Float inputs
		{
			name:          "unmarshals positive float",
			input:         "42.5",
			expectedFloat: 42.5,
			expectedValid: true,
		},
		{
			name:          "unmarshals negative float",
			input:         "-10.25",
			expectedFloat: -10.25,
			expectedValid: true,
		},
		{
			name:          "unmarshals zero float",
			input:         "0.0",
			expectedFloat: 0.0,
			expectedValid: true,
		},
		{
			name:          "unmarshals scientific notation",
			input:         "1.23e4",
			expectedFloat: 12300.0,
			expectedValid: true,
		},
		{
			name:          "unmarshals negative scientific notation",
			input:         "-2.5e-3",
			expectedFloat: -0.0025,
			expectedValid: true,
		},
		// Integer inputs (should be treated as floats)
		{
			name:          "unmarshals integer as float",
			input:         "42",
			expectedFloat: 42.0,
			expectedValid: true,
		},
		{
			name:          "unmarshals negative integer as float",
			input:         "-10",
			expectedFloat: -10.0,
			expectedValid: true,
		},
		{
			name:          "unmarshals zero as float",
			input:         "0",
			expectedFloat: 0.0,
			expectedValid: true,
		},
		// String inputs
		{
			name:          "unmarshals string float",
			input:         `"123.45"`,
			expectedFloat: 123.45,
			expectedValid: true,
		},
		{
			name:          "unmarshals negative string float",
			input:         `"-456.78"`,
			expectedFloat: -456.78,
			expectedValid: true,
		},
		{
			name:          "unmarshals string integer as float",
			input:         `"789"`,
			expectedFloat: 789.0,
			expectedValid: true,
		},
		{
			name:          "unmarshals string scientific notation",
			input:         `"1.5e2"`,
			expectedFloat: 150.0,
			expectedValid: true,
		},
		{
			name:          "unmarshals empty string as null",
			input:         `""`,
			expectedFloat: 0.0,
			expectedValid: false,
		},
		// Null input
		{
			name:          "unmarshals null as invalid",
			input:         "null",
			expectedFloat: 0.0,
			expectedValid: false,
		},
		// Edge cases - special float values
		{
			name:          "unmarshals positive infinity",
			input:         `"Inf"`,
			expectedFloat: float64(0), // Will be +Inf, but we'll check differently
			expectedValid: true,
		},
		{
			name:          "unmarshals negative infinity",
			input:         `"-Inf"`,
			expectedFloat: float64(0), // Will be -Inf, but we'll check differently
			expectedValid: true,
		},
		{
			name:          "unmarshals NaN",
			input:         `"NaN"`,
			expectedFloat: float64(0), // Will be NaN, but we'll check differently
			expectedValid: true,
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
			var nf NullFloat
			err := json.Unmarshal([]byte(tt.input), &nf)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)

			// Special handling for infinity and NaN
			switch tt.name {
			case "unmarshals positive infinity":
				assert.True(t, nf.Valid)
				assert.True(t, nf.Float > 0 && nf.Float == nf.Float*2) // Positive infinity check
			case "unmarshals negative infinity":
				assert.True(t, nf.Valid)
				assert.True(t, nf.Float < 0 && nf.Float == nf.Float*2) // Negative infinity check
			case "unmarshals NaN":
				assert.True(t, nf.Valid)
				assert.True(t, nf.Float != nf.Float) // NaN check
			default:
				assert.Equal(t, tt.expectedFloat, nf.Float)
				assert.Equal(t, tt.expectedValid, nf.Valid)
			}
		})
	}
}

func TestNullFloat_MarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    NullFloat
		expected string
	}{
		{
			name:     "marshals valid positive float",
			input:    NullFloat{Float: 42.5, Valid: true},
			expected: "42.5",
		},
		{
			name:     "marshals valid negative float",
			input:    NullFloat{Float: -10.25, Valid: true},
			expected: "-10.25",
		},
		{
			name:     "marshals valid zero",
			input:    NullFloat{Float: 0.0, Valid: true},
			expected: "0",
		},
		{
			name:     "marshals valid integer-like float",
			input:    NullFloat{Float: 123.0, Valid: true},
			expected: "123",
		},
		{
			name:     "marshals invalid as null",
			input:    NullFloat{Float: 123.45, Valid: false},
			expected: "null",
		},
		{
			name:     "marshals zero invalid as null",
			input:    NullFloat{Float: 0.0, Valid: false},
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

func TestNullFloat_ValueOr(t *testing.T) {
	tests := []struct {
		name     string
		input    NullFloat
		fallback float64
		expected float64
	}{
		{
			name:     "returns value when valid",
			input:    NullFloat{Float: 123.45, Valid: true},
			fallback: 999.99,
			expected: 123.45,
		},
		{
			name:     "returns zero when valid",
			input:    NullFloat{Float: 0.0, Valid: true},
			fallback: 999.99,
			expected: 0.0,
		},
		{
			name:     "returns fallback when invalid",
			input:    NullFloat{Float: 123.45, Valid: false},
			fallback: 999.99,
			expected: 999.99,
		},
		{
			name:     "returns fallback when invalid with zero",
			input:    NullFloat{Float: 0.0, Valid: false},
			fallback: 456.78,
			expected: 456.78,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.input.ValueOr(tt.fallback))
		})
	}
}

func TestNullFloat_String(t *testing.T) {
	tests := []struct {
		name     string
		input    NullFloat
		expected string
	}{
		{
			name:     "returns string representation when valid",
			input:    NullFloat{Float: 123.45, Valid: true},
			expected: "123.45",
		},
		{
			name:     "returns negative string representation when valid",
			input:    NullFloat{Float: -456.78, Valid: true},
			expected: "-456.78",
		},
		{
			name:     "returns zero string when valid",
			input:    NullFloat{Float: 0.0, Valid: true},
			expected: "0",
		},
		{
			name:     "returns integer-like string when valid",
			input:    NullFloat{Float: 42.0, Valid: true},
			expected: "42",
		},
		{
			name:     "returns null when invalid",
			input:    NullFloat{Float: 123.45, Valid: false},
			expected: "null",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.input.String())
		})
	}
}

func TestNewNullFloat(t *testing.T) {
	tests := []struct {
		name  string
		input float64
	}{
		{
			name:  "creates valid NullFloat with positive value",
			input: 123.45,
		},
		{
			name:  "creates valid NullFloat with negative value",
			input: -456.78,
		},
		{
			name:  "creates valid NullFloat with zero",
			input: 0.0,
		},
		{
			name:  "creates valid NullFloat with integer-like value",
			input: 42.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nf := NewNullFloat(tt.input)
			assert.Equal(t, tt.input, nf.Float)
			assert.True(t, nf.Valid)
		})
	}
}

func TestNullFloat_ActualNull(t *testing.T) {
	// Test actual JSON null (not string "null")
	jsonData := `{"value": null}`

	var result struct {
		Value NullFloat `json:"value"`
	}

	err := json.Unmarshal([]byte(jsonData), &result)
	require.NoError(t, err)

	// Should be invalid (null)
	assert.False(t, result.Value.Valid, "Expected null value to be invalid")
	assert.Equal(t, 0.0, result.Value.Float, "Expected zero value for null")

	// Also test direct unmarshaling of null
	var nullFloat NullFloat
	err = json.Unmarshal([]byte("null"), &nullFloat)
	require.NoError(t, err)
	assert.False(t, nullFloat.Valid, "Direct null should be invalid")
	assert.Equal(t, 0.0, nullFloat.Float, "Direct null should have zero value")

	// Test string "null" (this should also work)
	var stringNullFloat NullFloat
	err = json.Unmarshal([]byte(`"null"`), &stringNullFloat)
	require.NoError(t, err)
	assert.False(t, stringNullFloat.Valid, "String null should be invalid")
	assert.Equal(t, 0.0, stringNullFloat.Float, "String null should have zero value")

	// Test nil data slice (edge case)
	var nilDataFloat NullFloat
	err = nilDataFloat.UnmarshalJSON(nil)
	require.NoError(t, err)
	assert.False(t, nilDataFloat.Valid, "Nil data should be invalid")
	assert.Equal(t, 0.0, nilDataFloat.Float, "Nil data should have zero value")
}

func TestNullFloat_RoundTrip(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{
			name:  "float",
			input: "42.5",
		},
		{
			name:  "integer",
			input: "123",
		},
		{
			name:  "string float",
			input: `"456.78"`,
		},
		{
			name:  "string integer",
			input: `"789"`,
		},
		{
			name:  "negative float",
			input: "-10.25",
		},
		{
			name:  "negative string",
			input: `"-20.5"`,
		},
		{
			name:  "zero",
			input: "0",
		},
		{
			name:  "scientific notation",
			input: "1.23e4",
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
			var nf NullFloat
			err := json.Unmarshal([]byte(tt.input), &nf)
			require.NoError(t, err)

			// Marshal it back
			data, err := json.Marshal(nf)
			require.NoError(t, err)

			// Unmarshal again to verify consistency
			var nf2 NullFloat
			err = json.Unmarshal(data, &nf2)
			require.NoError(t, err)

			// Should be the same
			assert.Equal(t, nf.Float, nf2.Float)
			assert.Equal(t, nf.Valid, nf2.Valid)
		})
	}
}

func TestNullFloat_InStruct(t *testing.T) {
	type TestStruct struct {
		Value NullFloat `json:"value"`
	}

	t.Run("unmarshals in struct with string value", func(t *testing.T) {
		input := `{"value": "42.5"}`
		var ts TestStruct
		err := json.Unmarshal([]byte(input), &ts)
		require.NoError(t, err)
		assert.Equal(t, 42.5, ts.Value.Float)
		assert.True(t, ts.Value.Valid)
	})

	t.Run("unmarshals in struct with float value", func(t *testing.T) {
		input := `{"value": 123.45}`
		var ts TestStruct
		err := json.Unmarshal([]byte(input), &ts)
		require.NoError(t, err)
		assert.Equal(t, 123.45, ts.Value.Float)
		assert.True(t, ts.Value.Valid)
	})

	t.Run("unmarshals in struct with integer value", func(t *testing.T) {
		input := `{"value": 789}`
		var ts TestStruct
		err := json.Unmarshal([]byte(input), &ts)
		require.NoError(t, err)
		assert.Equal(t, 789.0, ts.Value.Float)
		assert.True(t, ts.Value.Valid)
	})

	t.Run("unmarshals in struct with null value", func(t *testing.T) {
		input := `{"value": null}`
		var ts TestStruct
		err := json.Unmarshal([]byte(input), &ts)
		require.NoError(t, err)
		assert.Equal(t, 0.0, ts.Value.Float)
		assert.False(t, ts.Value.Valid)
	})

	t.Run("marshals valid value from struct", func(t *testing.T) {
		ts := TestStruct{Value: NullFloat{Float: 456.78, Valid: true}}
		data, err := json.Marshal(ts)
		require.NoError(t, err)
		assert.Contains(t, string(data), `"value":456.78`)
	})

	t.Run("marshals null value from struct", func(t *testing.T) {
		ts := TestStruct{Value: NullFloat{Float: 456.78, Valid: false}}
		data, err := json.Marshal(ts)
		require.NoError(t, err)
		assert.Contains(t, string(data), `"value":null`)
	})
}

func TestNullFloat_SpecialValues(t *testing.T) {
	t.Run("handles positive infinity", func(t *testing.T) {
		input := `"Inf"`
		var nf NullFloat
		err := json.Unmarshal([]byte(input), &nf)
		require.NoError(t, err)
		assert.True(t, nf.Valid)
		assert.True(t, nf.Float > 0 && nf.Float == nf.Float*2) // Positive infinity check
	})

	t.Run("handles negative infinity", func(t *testing.T) {
		input := `"-Inf"`
		var nf NullFloat
		err := json.Unmarshal([]byte(input), &nf)
		require.NoError(t, err)
		assert.True(t, nf.Valid)
		assert.True(t, nf.Float < 0 && nf.Float == nf.Float*2) // Negative infinity check
	})

	t.Run("handles NaN", func(t *testing.T) {
		input := `"NaN"`
		var nf NullFloat
		err := json.Unmarshal([]byte(input), &nf)
		require.NoError(t, err)
		assert.True(t, nf.Valid)
		assert.True(t, nf.Float != nf.Float) // NaN check
	})
}
