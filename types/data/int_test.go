package data

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInt_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int64
		wantErr  bool
	}{
		// Integer inputs
		{
			name:     "unmarshals integer",
			input:    "42",
			expected: 42,
		},
		{
			name:     "unmarshals negative integer",
			input:    "-10",
			expected: -10,
		},
		{
			name:     "unmarshals zero",
			input:    "0",
			expected: 0,
		},
		{
			name:     "unmarshals large integer",
			input:    "9223372036854775807", // max int64
			expected: 9223372036854775807,
		},
		// Float inputs (should be truncated)
		{
			name:     "unmarshals float as integer (truncated)",
			input:    "42.9",
			expected: 42,
		},
		{
			name:     "unmarshals negative float as integer (truncated)",
			input:    "-42.9",
			expected: -42,
		},
		// String inputs
		{
			name:     "unmarshals string integer",
			input:    `"123"`,
			expected: 123,
		},
		{
			name:     "unmarshals negative string integer",
			input:    `"-456"`,
			expected: -456,
		},
		{
			name:     "unmarshals string float as integer",
			input:    `"78.9"`,
			expected: 78,
		},
		{
			name:     "unmarshals empty string as zero",
			input:    `""`,
			expected: 0,
		},
		{
			name:     "unmarshals string null as zero",
			input:    `"null"`,
			expected: 0,
		},
		// Null input
		{
			name:     "unmarshals null as zero",
			input:    "null",
			expected: 0,
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
			var i Int
			err := json.Unmarshal([]byte(tt.input), &i)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.expected, i.Int64())
		})
	}
}

func TestInt_UnmarshalJSON_NilData(t *testing.T) {
	var i Int
	err := i.UnmarshalJSON(nil)
	require.NoError(t, err)
	assert.Equal(t, int64(0), i.Int64())
}

func TestInt_MarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    Int
		expected string
	}{
		{
			name:     "marshals positive integer",
			input:    Int(42),
			expected: "42",
		},
		{
			name:     "marshals negative integer",
			input:    Int(-10),
			expected: "-10",
		},
		{
			name:     "marshals zero",
			input:    Int(0),
			expected: "0",
		},
		{
			name:     "marshals large integer",
			input:    Int(9223372036854775807),
			expected: "9223372036854775807",
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

func TestInt_Int64(t *testing.T) {
	tests := []struct {
		name     string
		input    Int
		expected int64
	}{
		{
			name:     "returns correct integer value",
			input:    Int(123),
			expected: 123,
		},
		{
			name:     "returns correct negative integer value",
			input:    Int(-456),
			expected: -456,
		},
		{
			name:     "returns zero",
			input:    Int(0),
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.input.Int64())
		})
	}
}

func TestInt_String(t *testing.T) {
	tests := []struct {
		name     string
		input    Int
		expected string
	}{
		{
			name:     "returns correct string representation",
			input:    Int(123),
			expected: "123",
		},
		{
			name:     "returns correct negative string representation",
			input:    Int(-456),
			expected: "-456",
		},
		{
			name:     "returns zero string",
			input:    Int(0),
			expected: "0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.input.String())
		})
	}
}

func TestNewInt(t *testing.T) {
	tests := []struct {
		name     string
		input    int64
		expected Int
	}{
		{
			name:     "creates new Int with positive value",
			input:    42,
			expected: Int(42),
		},
		{
			name:     "creates new Int with negative value",
			input:    -10,
			expected: Int(-10),
		},
		{
			name:     "creates new Int with zero",
			input:    0,
			expected: Int(0),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewInt(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestInt_RoundTrip(t *testing.T) {
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
		{
			name:  "string null",
			input: `"null"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Unmarshal the input
			var i Int
			err := json.Unmarshal([]byte(tt.input), &i)
			require.NoError(t, err)

			// Marshal it back
			data, err := json.Marshal(i)
			require.NoError(t, err)

			// Unmarshal again to verify consistency
			var i2 Int
			err = json.Unmarshal(data, &i2)
			require.NoError(t, err)

			// Should be the same
			assert.Equal(t, i.Int64(), i2.Int64())
		})
	}
}

func TestInt_InStruct(t *testing.T) {
	type TestStruct struct {
		Value Int `json:"value"`
	}

	t.Run("unmarshals in struct with string value", func(t *testing.T) {
		input := `{"value": "42"}`
		var ts TestStruct
		err := json.Unmarshal([]byte(input), &ts)
		require.NoError(t, err)
		assert.Equal(t, int64(42), ts.Value.Int64())
	})

	t.Run("unmarshals in struct with integer value", func(t *testing.T) {
		input := `{"value": 123}`
		var ts TestStruct
		err := json.Unmarshal([]byte(input), &ts)
		require.NoError(t, err)
		assert.Equal(t, int64(123), ts.Value.Int64())
	})

	t.Run("unmarshals in struct with null value", func(t *testing.T) {
		input := `{"value": null}`
		var ts TestStruct
		err := json.Unmarshal([]byte(input), &ts)
		require.NoError(t, err)
		assert.Equal(t, int64(0), ts.Value.Int64())
	})

	t.Run("marshals from struct", func(t *testing.T) {
		ts := TestStruct{Value: Int(456)}
		data, err := json.Marshal(ts)
		require.NoError(t, err)
		assert.Contains(t, string(data), `"value":456`)
	})
}
