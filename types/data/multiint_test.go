package data

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMultiInt_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
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
		// Float inputs (should be truncated)
		{
			name:     "unmarshals float as integer (truncated)",
			input:    "42.9",
			expected: 42,
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var mi MultiInt
			err := json.Unmarshal([]byte(tt.input), &mi)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.expected, mi.Int())
		})
	}
}

func TestMultiInt_MarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    MultiInt
		expected string
	}{
		{
			name:     "marshals positive integer",
			input:    MultiInt(42),
			expected: "42",
		},
		{
			name:     "marshals negative integer",
			input:    MultiInt(-10),
			expected: "-10",
		},
		{
			name:     "marshals zero",
			input:    MultiInt(0),
			expected: "0",
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

func TestMultiInt_Int(t *testing.T) {
	tests := []struct {
		name     string
		input    MultiInt
		expected int
	}{
		{
			name:     "returns correct integer value",
			input:    MultiInt(123),
			expected: 123,
		},
		{
			name:     "returns correct negative integer value",
			input:    MultiInt(-456),
			expected: -456,
		},
		{
			name:     "returns zero",
			input:    MultiInt(0),
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.input.Int())
		})
	}
}

func TestMultiInt_String(t *testing.T) {
	tests := []struct {
		name     string
		input    MultiInt
		expected string
	}{
		{
			name:     "returns correct string representation",
			input:    MultiInt(123),
			expected: "123",
		},
		{
			name:     "returns correct negative string representation",
			input:    MultiInt(-456),
			expected: "-456",
		},
		{
			name:     "returns zero string",
			input:    MultiInt(0),
			expected: "0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.input.String())
		})
	}
}

func TestMultiInt_RoundTrip(t *testing.T) {
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
			var mi MultiInt
			err := json.Unmarshal([]byte(tt.input), &mi)
			require.NoError(t, err)

			// Marshal it back
			data, err := json.Marshal(mi)
			require.NoError(t, err)

			// Unmarshal again to verify consistency
			var mi2 MultiInt
			err = json.Unmarshal(data, &mi2)
			require.NoError(t, err)

			// Should be the same
			assert.Equal(t, mi.Int(), mi2.Int())
		})
	}
}

func TestMultiInt_InStruct(t *testing.T) {
	type TestStruct struct {
		Value MultiInt `json:"value"`
	}

	t.Run("unmarshals in struct with string value", func(t *testing.T) {
		input := `{"value": "42"}`
		var ts TestStruct
		err := json.Unmarshal([]byte(input), &ts)
		require.NoError(t, err)
		assert.Equal(t, 42, ts.Value.Int())
	})

	t.Run("unmarshals in struct with integer value", func(t *testing.T) {
		input := `{"value": 123}`
		var ts TestStruct
		err := json.Unmarshal([]byte(input), &ts)
		require.NoError(t, err)
		assert.Equal(t, 123, ts.Value.Int())
	})

	t.Run("marshals from struct", func(t *testing.T) {
		ts := TestStruct{Value: MultiInt(456)}
		data, err := json.Marshal(ts)
		require.NoError(t, err)
		assert.Contains(t, string(data), `"value":456`)
	})
}
