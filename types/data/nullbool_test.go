package data

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNullBool_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expectedBool  bool
		expectedValid bool
		wantErr       bool
	}{
		// Boolean inputs
		{
			name:          "unmarshals true boolean",
			input:         "true",
			expectedBool:  true,
			expectedValid: true,
		},
		{
			name:          "unmarshals false boolean",
			input:         "false",
			expectedBool:  false,
			expectedValid: true,
		},
		// String inputs that should be true
		{
			name:          "unmarshals true string",
			input:         `"true"`,
			expectedBool:  true,
			expectedValid: true,
		},
		{
			name:          "unmarshals success string",
			input:         `"success"`,
			expectedBool:  true,
			expectedValid: true,
		},
		{
			name:          "unmarshals string 1 as true",
			input:         `"1"`,
			expectedBool:  true,
			expectedValid: true,
		},
		// String inputs that should be false
		{
			name:          "unmarshals false string",
			input:         `"false"`,
			expectedBool:  false,
			expectedValid: true,
		},
		{
			name:          "unmarshals failed string",
			input:         `"failed"`,
			expectedBool:  false,
			expectedValid: true,
		},
		{
			name:          "unmarshals error string",
			input:         `"error"`,
			expectedBool:  false,
			expectedValid: true,
		},
		{
			name:          "unmarshals random string as false",
			input:         `"random"`,
			expectedBool:  false,
			expectedValid: true,
		},
		{
			name:          "unmarshals string 0 as false",
			input:         `"0"`,
			expectedBool:  false,
			expectedValid: true,
		},
		{
			name:          "unmarshals empty string as null",
			input:         `""`,
			expectedBool:  false,
			expectedValid: false,
		},
		// Number inputs
		{
			name:          "unmarshals zero as false",
			input:         "0",
			expectedBool:  false,
			expectedValid: true,
		},
		{
			name:          "unmarshals one as true",
			input:         "1",
			expectedBool:  true,
			expectedValid: true,
		},
		{
			name:          "unmarshals positive number as true",
			input:         "42",
			expectedBool:  true,
			expectedValid: true,
		},
		{
			name:          "unmarshals negative number as true",
			input:         "-1",
			expectedBool:  true,
			expectedValid: true,
		},
		{
			name:          "unmarshals float zero as false",
			input:         "0.0",
			expectedBool:  false,
			expectedValid: true,
		},
		{
			name:          "unmarshals positive float as true",
			input:         "1.5",
			expectedBool:  true,
			expectedValid: true,
		},
		// Null input
		{
			name:          "unmarshals null as invalid",
			input:         "null",
			expectedBool:  false,
			expectedValid: false,
		},
		// Error cases
		{
			name:    "fails to unmarshal array",
			input:   `[1, 2, 3]`,
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
			var nb NullBool
			err := json.Unmarshal([]byte(tt.input), &nb)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.expectedBool, nb.Bool)
			assert.Equal(t, tt.expectedValid, nb.Valid)
		})
	}
}

func TestNullBool_MarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    NullBool
		expected string
	}{
		{
			name:     "marshals valid true",
			input:    NullBool{Bool: true, Valid: true},
			expected: "true",
		},
		{
			name:     "marshals valid false",
			input:    NullBool{Bool: false, Valid: true},
			expected: "false",
		},
		{
			name:     "marshals invalid as null (with true value)",
			input:    NullBool{Bool: true, Valid: false},
			expected: "null",
		},
		{
			name:     "marshals invalid as null (with false value)",
			input:    NullBool{Bool: false, Valid: false},
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

func TestNullBool_ValueOr(t *testing.T) {
	tests := []struct {
		name     string
		input    NullBool
		fallback bool
		expected bool
	}{
		{
			name:     "returns true when valid",
			input:    NullBool{Bool: true, Valid: true},
			fallback: false,
			expected: true,
		},
		{
			name:     "returns false when valid",
			input:    NullBool{Bool: false, Valid: true},
			fallback: true,
			expected: false,
		},
		{
			name:     "returns fallback when invalid (true fallback)",
			input:    NullBool{Bool: false, Valid: false},
			fallback: true,
			expected: true,
		},
		{
			name:     "returns fallback when invalid (false fallback)",
			input:    NullBool{Bool: true, Valid: false},
			fallback: false,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.input.ValueOr(tt.fallback))
		})
	}
}

func TestNullBool_String(t *testing.T) {
	tests := []struct {
		name     string
		input    NullBool
		expected string
	}{
		{
			name:     "returns true string when valid true",
			input:    NullBool{Bool: true, Valid: true},
			expected: "true",
		},
		{
			name:     "returns false string when valid false",
			input:    NullBool{Bool: false, Valid: true},
			expected: "false",
		},
		{
			name:     "returns null when invalid (with true value)",
			input:    NullBool{Bool: true, Valid: false},
			expected: "null",
		},
		{
			name:     "returns null when invalid (with false value)",
			input:    NullBool{Bool: false, Valid: false},
			expected: "null",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.input.String())
		})
	}
}

func TestNewNullBool(t *testing.T) {
	tests := []struct {
		name  string
		input bool
	}{
		{
			name:  "creates valid NullBool with true",
			input: true,
		},
		{
			name:  "creates valid NullBool with false",
			input: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nb := NewNullBool(tt.input)
			assert.Equal(t, tt.input, nb.Bool)
			assert.True(t, nb.Valid)
		})
	}
}

func TestNullBool_ActualNull(t *testing.T) {
	// Test actual JSON null (not string "null")
	jsonData := `{"value": null}`

	var result struct {
		Value NullBool `json:"value"`
	}

	err := json.Unmarshal([]byte(jsonData), &result)
	require.NoError(t, err)

	// Should be invalid (null)
	assert.False(t, result.Value.Valid, "Expected null value to be invalid")
	assert.False(t, result.Value.Bool, "Expected zero value for null")

	// Also test direct unmarshaling of null
	var nullBool NullBool
	err = json.Unmarshal([]byte("null"), &nullBool)
	require.NoError(t, err)
	assert.False(t, nullBool.Valid, "Direct null should be invalid")
	assert.False(t, nullBool.Bool, "Direct null should have zero value")

	// Test string "null" (this should also work)
	var stringNullBool NullBool
	err = json.Unmarshal([]byte(`"null"`), &stringNullBool)
	require.NoError(t, err)
	assert.False(t, stringNullBool.Valid, "String null should be invalid")
	assert.False(t, stringNullBool.Bool, "String null should have zero value")

	// Test nil data slice (edge case)
	var nilDataBool NullBool
	err = nilDataBool.UnmarshalJSON(nil)
	require.NoError(t, err)
	assert.False(t, nilDataBool.Valid, "Nil data should be invalid")
	assert.False(t, nilDataBool.Bool, "Nil data should have zero value")
}

func TestNullBool_RoundTrip(t *testing.T) {
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
			name:  "true string",
			input: `"true"`,
		},
		{
			name:  "success string",
			input: `"success"`,
		},
		{
			name:  "string 1",
			input: `"1"`,
		},
		{
			name:  "false string",
			input: `"false"`,
		},
		{
			name:  "failed string",
			input: `"failed"`,
		},
		{
			name:  "number zero",
			input: "0",
		},
		{
			name:  "number one",
			input: "1",
		},
		{
			name:  "positive number",
			input: "42",
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
			var nb NullBool
			err := json.Unmarshal([]byte(tt.input), &nb)
			require.NoError(t, err)

			// Marshal it back
			data, err := json.Marshal(nb)
			require.NoError(t, err)

			// Unmarshal again to verify consistency
			var nb2 NullBool
			err = json.Unmarshal(data, &nb2)
			require.NoError(t, err)

			// Should be the same
			assert.Equal(t, nb.Bool, nb2.Bool)
			assert.Equal(t, nb.Valid, nb2.Valid)
		})
	}
}

func TestNullBool_InStruct(t *testing.T) {
	type TestStruct struct {
		Value NullBool `json:"value"`
	}

	t.Run("unmarshals in struct with string true value", func(t *testing.T) {
		input := `{"value": "true"}`
		var ts TestStruct
		err := json.Unmarshal([]byte(input), &ts)
		require.NoError(t, err)
		assert.True(t, ts.Value.Bool)
		assert.True(t, ts.Value.Valid)
	})

	t.Run("unmarshals in struct with boolean false value", func(t *testing.T) {
		input := `{"value": false}`
		var ts TestStruct
		err := json.Unmarshal([]byte(input), &ts)
		require.NoError(t, err)
		assert.False(t, ts.Value.Bool)
		assert.True(t, ts.Value.Valid)
	})

	t.Run("unmarshals in struct with success string", func(t *testing.T) {
		input := `{"value": "success"}`
		var ts TestStruct
		err := json.Unmarshal([]byte(input), &ts)
		require.NoError(t, err)
		assert.True(t, ts.Value.Bool)
		assert.True(t, ts.Value.Valid)
	})

	t.Run("unmarshals in struct with null value", func(t *testing.T) {
		input := `{"value": null}`
		var ts TestStruct
		err := json.Unmarshal([]byte(input), &ts)
		require.NoError(t, err)
		assert.False(t, ts.Value.Bool)
		assert.False(t, ts.Value.Valid)
	})

	t.Run("marshals valid true value from struct", func(t *testing.T) {
		ts := TestStruct{Value: NullBool{Bool: true, Valid: true}}
		data, err := json.Marshal(ts)
		require.NoError(t, err)
		assert.Contains(t, string(data), `"value":true`)
	})

	t.Run("marshals valid false value from struct", func(t *testing.T) {
		ts := TestStruct{Value: NullBool{Bool: false, Valid: true}}
		data, err := json.Marshal(ts)
		require.NoError(t, err)
		assert.Contains(t, string(data), `"value":false`)
	})

	t.Run("marshals null value from struct", func(t *testing.T) {
		ts := TestStruct{Value: NullBool{Bool: true, Valid: false}}
		data, err := json.Marshal(ts)
		require.NoError(t, err)
		assert.Contains(t, string(data), `"value":null`)
	})
}

func TestNullBool_CompareWithMultiBool(t *testing.T) {
	t.Run("null handling difference", func(t *testing.T) {
		input := "null"

		var nb NullBool
		err := json.Unmarshal([]byte(input), &nb)
		require.NoError(t, err)
		// NullBool preserves null state
		assert.False(t, nb.Valid)
		assert.False(t, nb.Bool)
		// MultiBool tests removed
		nbData, err := json.Marshal(nb)
		require.NoError(t, err)
		assert.Equal(t, "null", string(nbData))
	})

	t.Run("empty string handling difference", func(t *testing.T) {
		input := `""`

		var nb NullBool
		err := json.Unmarshal([]byte(input), &nb)
		require.NoError(t, err)
		// NullBool treats empty string as null
		assert.False(t, nb.Valid)
		assert.False(t, nb.Bool)
		// MultiBool tests removed
		nbData, err := json.Marshal(nb)
		require.NoError(t, err)
		assert.Equal(t, "null", string(nbData))
	})
}
