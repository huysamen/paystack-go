package data

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFloat_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected float64
		wantErr  bool
	}{
		// Float inputs
		{
			name:     "unmarshals float",
			input:    "42.5",
			expected: 42.5,
		},
		{
			name:     "unmarshals negative float",
			input:    "-10.3",
			expected: -10.3,
		},
		{
			name:     "unmarshals zero float",
			input:    "0.0",
			expected: 0.0,
		},
		{
			name:     "unmarshals large float",
			input:    "1.7976931348623157e+308", // close to max float64
			expected: 1.7976931348623157e+308,
		},
		// Integer inputs (should be converted to float)
		{
			name:     "unmarshals integer as float",
			input:    "42",
			expected: 42.0,
		},
		{
			name:     "unmarshals negative integer as float",
			input:    "-10",
			expected: -10.0,
		},
		// String inputs
		{
			name:     "unmarshals string float",
			input:    `"123.45"`,
			expected: 123.45,
		},
		{
			name:     "unmarshals negative string float",
			input:    `"-456.78"`,
			expected: -456.78,
		},
		{
			name:     "unmarshals string integer as float",
			input:    `"789"`,
			expected: 789.0,
		},
		{
			name:     "unmarshals empty string as zero",
			input:    `""`,
			expected: 0.0,
		},
		{
			name:     "unmarshals string null as zero",
			input:    `"null"`,
			expected: 0.0,
		},
		// Scientific notation
		{
			name:     "unmarshals scientific notation",
			input:    "1.23e+4",
			expected: 12300.0,
		},
		{
			name:     "unmarshals string scientific notation",
			input:    `"1.23e-4"`,
			expected: 0.000123,
		},
		// Null input
		{
			name:     "unmarshals null as zero",
			input:    "null",
			expected: 0.0,
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
			input:   `[1.1, 2.2, 3.3]`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var f Float
			err := json.Unmarshal([]byte(tt.input), &f)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.expected, f.Float64())
		})
	}
}

func TestFloat_UnmarshalJSON_NilData(t *testing.T) {
	var f Float
	err := f.UnmarshalJSON(nil)
	require.NoError(t, err)
	assert.Equal(t, 0.0, f.Float64())
}

func TestFloat_MarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    Float
		expected string
	}{
		{
			name:     "marshals positive float",
			input:    Float(42.5),
			expected: "42.5",
		},
		{
			name:     "marshals negative float",
			input:    Float(-10.3),
			expected: "-10.3",
		},
		{
			name:     "marshals zero",
			input:    Float(0.0),
			expected: "0",
		},
		{
			name:     "marshals integer as float",
			input:    Float(123),
			expected: "123",
		},
		{
			name:     "marshals scientific notation as decimal",
			input:    Float(1.23e+10),
			expected: "12300000000", // Go's json.Marshal converts scientific to decimal when possible
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

func TestFloat_Float64(t *testing.T) {
	tests := []struct {
		name     string
		input    Float
		expected float64
	}{
		{
			name:     "returns correct float value",
			input:    Float(123.45),
			expected: 123.45,
		},
		{
			name:     "returns correct negative float value",
			input:    Float(-456.78),
			expected: -456.78,
		},
		{
			name:     "returns zero",
			input:    Float(0.0),
			expected: 0.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.input.Float64())
		})
	}
}

func TestFloat_String(t *testing.T) {
	tests := []struct {
		name     string
		input    Float
		expected string
	}{
		{
			name:     "returns correct string representation",
			input:    Float(123.45),
			expected: "123.45",
		},
		{
			name:     "returns correct negative string representation",
			input:    Float(-456.78),
			expected: "-456.78",
		},
		{
			name:     "returns zero string",
			input:    Float(0.0),
			expected: "0",
		},
		{
			name:     "returns scientific notation for large numbers",
			input:    Float(1.23e+10),
			expected: "1.23e+10",
		},
		{
			name:     "returns scientific notation for small numbers",
			input:    Float(1.23e-10),
			expected: "1.23e-10",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.input.String())
		})
	}
}

func TestNewFloat(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected Float
	}{
		{
			name:     "creates new Float with positive value",
			input:    42.5,
			expected: Float(42.5),
		},
		{
			name:     "creates new Float with negative value",
			input:    -10.3,
			expected: Float(-10.3),
		},
		{
			name:     "creates new Float with zero",
			input:    0.0,
			expected: Float(0.0),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewFloat(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestFloat_RoundTrip(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{
			name:  "float",
			input: "42.5",
		},
		{
			name:  "string float",
			input: `"123.45"`,
		},
		{
			name:  "integer",
			input: "789",
		},
		{
			name:  "string integer",
			input: `"456"`,
		},
		{
			name:  "negative float",
			input: "-10.3",
		},
		{
			name:  "negative string float",
			input: `"-20.7"`,
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
		{
			name:  "scientific notation",
			input: "1.23e+4",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Unmarshal the input
			var f Float
			err := json.Unmarshal([]byte(tt.input), &f)
			require.NoError(t, err)

			// Marshal it back
			data, err := json.Marshal(f)
			require.NoError(t, err)

			// Unmarshal again to verify consistency
			var f2 Float
			err = json.Unmarshal(data, &f2)
			require.NoError(t, err)

			// Should be the same (within floating point precision)
			assert.True(t, math.Abs(f.Float64()-f2.Float64()) < 1e-10)
		})
	}
}

func TestFloat_InStruct(t *testing.T) {
	type TestStruct struct {
		Value Float `json:"value"`
	}

	t.Run("unmarshals in struct with string value", func(t *testing.T) {
		input := `{"value": "42.5"}`
		var ts TestStruct
		err := json.Unmarshal([]byte(input), &ts)
		require.NoError(t, err)
		assert.Equal(t, 42.5, ts.Value.Float64())
	})

	t.Run("unmarshals in struct with float value", func(t *testing.T) {
		input := `{"value": 123.45}`
		var ts TestStruct
		err := json.Unmarshal([]byte(input), &ts)
		require.NoError(t, err)
		assert.Equal(t, 123.45, ts.Value.Float64())
	})

	t.Run("unmarshals in struct with integer value", func(t *testing.T) {
		input := `{"value": 789}`
		var ts TestStruct
		err := json.Unmarshal([]byte(input), &ts)
		require.NoError(t, err)
		assert.Equal(t, 789.0, ts.Value.Float64())
	})

	t.Run("unmarshals in struct with null value", func(t *testing.T) {
		input := `{"value": null}`
		var ts TestStruct
		err := json.Unmarshal([]byte(input), &ts)
		require.NoError(t, err)
		assert.Equal(t, 0.0, ts.Value.Float64())
	})

	t.Run("marshals from struct", func(t *testing.T) {
		ts := TestStruct{Value: Float(456.78)}
		data, err := json.Marshal(ts)
		require.NoError(t, err)
		assert.Contains(t, string(data), `"value":456.78`)
	})
}
