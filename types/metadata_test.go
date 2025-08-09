package types

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMetadata_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expectedValid bool
		expectedData  map[string]any
		wantErr       bool
	}{
		{
			name:          "unmarshals valid object",
			input:         `{"key1": "value1", "key2": 42}`,
			expectedValid: true,
			expectedData:  map[string]any{"key1": "value1", "key2": float64(42)},
		},
		{
			name:          "unmarshals empty object",
			input:         `{}`,
			expectedValid: true,
			expectedData:  map[string]any{},
		},
		{
			name:          "unmarshals nested object",
			input:         `{"user": {"name": "John", "age": 30}, "active": true}`,
			expectedValid: true,
			expectedData:  map[string]any{"user": map[string]any{"name": "John", "age": float64(30)}, "active": true},
		},
		{
			name:          "handles null as invalid",
			input:         `null`,
			expectedValid: false,
			expectedData:  map[string]any{},
		},
		{
			name:    "fails on string input",
			input:   `"some string"`,
			wantErr: true,
		},
		{
			name:    "fails on number input",
			input:   `42`,
			wantErr: true,
		},
		{
			name:    "fails on boolean input",
			input:   `true`,
			wantErr: true,
		},
		{
			name:    "fails on array input",
			input:   `[1, 2, 3]`,
			wantErr: true,
		},
		{
			name:    "fails on invalid JSON",
			input:   `{invalid json}`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var m Metadata
			err := json.Unmarshal([]byte(tt.input), &m)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.expectedValid, m.Valid)
			assert.Equal(t, tt.expectedData, m.Metadata)
		})
	}
}

func TestMetadata_MarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    Metadata
		expected string
	}{
		{
			name: "marshals valid metadata",
			input: Metadata{
				Metadata: map[string]any{"key1": "value1", "key2": 42},
				Valid:    true,
			},
			expected: `{"key1":"value1","key2":42}`,
		},
		{
			name: "marshals empty valid metadata",
			input: Metadata{
				Metadata: map[string]any{},
				Valid:    true,
			},
			expected: `{}`,
		},
		{
			name: "marshals invalid metadata as null",
			input: Metadata{
				Metadata: map[string]any{"key": "value"},
				Valid:    false,
			},
			expected: `null`,
		},
		{
			name: "marshals invalid empty metadata as null",
			input: Metadata{
				Metadata: map[string]any{},
				Valid:    false,
			},
			expected: `null`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := json.Marshal(tt.input)
			require.NoError(t, err)
			assert.JSONEq(t, tt.expected, string(data))
		})
	}
}

func TestMetadata_IsEmpty(t *testing.T) {
	tests := []struct {
		name     string
		input    Metadata
		expected bool
	}{
		{
			name: "returns false for non-empty metadata",
			input: Metadata{
				Metadata: map[string]any{"key": "value"},
				Valid:    true,
			},
			expected: false,
		},
		{
			name: "returns true for empty metadata",
			input: Metadata{
				Metadata: map[string]any{},
				Valid:    true,
			},
			expected: true,
		},
		{
			name: "returns true for nil metadata map",
			input: Metadata{
				Metadata: nil,
				Valid:    false,
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.input.IsEmpty())
		})
	}
}

func TestNewMetadata(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]any
		expected Metadata
	}{
		{
			name:  "creates valid metadata from non-empty map",
			input: map[string]any{"key1": "value1", "key2": 42},
			expected: Metadata{
				Metadata: map[string]any{"key1": "value1", "key2": 42},
				Valid:    true,
			},
		},
		{
			name:  "creates valid metadata from empty map",
			input: map[string]any{},
			expected: Metadata{
				Metadata: map[string]any{},
				Valid:    true,
			},
		},
		{
			name:  "creates valid metadata from nil map",
			input: nil,
			expected: Metadata{
				Metadata: nil,
				Valid:    true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewMetadata(tt.input)
			assert.Equal(t, tt.expected.Metadata, result.Metadata)
			assert.Equal(t, tt.expected.Valid, result.Valid)
		})
	}
}

func TestMetadata_RoundTrip(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{
			name:  "simple object",
			input: `{"name":"John","age":30}`,
		},
		{
			name:  "empty object",
			input: `{}`,
		},
		{
			name:  "nested object",
			input: `{"user":{"name":"John","details":{"age":30,"active":true}},"metadata":{"source":"api"}}`,
		},
		{
			name:  "null value",
			input: `null`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Unmarshal the input
			var m Metadata
			err := json.Unmarshal([]byte(tt.input), &m)
			require.NoError(t, err)

			// Marshal it back
			data, err := json.Marshal(m)
			require.NoError(t, err)

			// Unmarshal again to verify consistency
			var m2 Metadata
			err = json.Unmarshal(data, &m2)
			require.NoError(t, err)

			// Should be the same
			assert.Equal(t, m.Valid, m2.Valid)
			assert.Equal(t, m.Metadata, m2.Metadata)
		})
	}
}

func TestMetadata_InStruct(t *testing.T) {
	type TestStruct struct {
		ID       int      `json:"id"`
		Metadata Metadata `json:"metadata"`
	}

	t.Run("unmarshals in struct with valid object", func(t *testing.T) {
		input := `{"id": 123, "metadata": {"key": "value", "count": 42}}`
		var ts TestStruct
		err := json.Unmarshal([]byte(input), &ts)
		require.NoError(t, err)
		assert.Equal(t, 123, ts.ID)
		assert.True(t, ts.Metadata.Valid)
		assert.Equal(t, "value", ts.Metadata.Metadata["key"])
		assert.Equal(t, float64(42), ts.Metadata.Metadata["count"])
	})

	t.Run("unmarshals in struct with null metadata", func(t *testing.T) {
		input := `{"id": 123, "metadata": null}`
		var ts TestStruct
		err := json.Unmarshal([]byte(input), &ts)
		require.NoError(t, err)
		assert.Equal(t, 123, ts.ID)
		assert.False(t, ts.Metadata.Valid)
		assert.True(t, ts.Metadata.IsEmpty())
	})

	t.Run("fails to unmarshal in struct with string metadata", func(t *testing.T) {
		input := `{"id": 123, "metadata": "invalid"}`
		var ts TestStruct
		err := json.Unmarshal([]byte(input), &ts)
		assert.Error(t, err)
	})

	t.Run("marshals valid metadata from struct", func(t *testing.T) {
		ts := TestStruct{
			ID:       456,
			Metadata: NewMetadata(map[string]any{"source": "test", "version": 1}),
		}
		data, err := json.Marshal(ts)
		require.NoError(t, err)
		assert.Contains(t, string(data), `"id":456`)
		assert.Contains(t, string(data), `"metadata":{"source":"test","version":1}`)
	})

	t.Run("marshals null metadata from struct", func(t *testing.T) {
		ts := TestStruct{
			ID:       789,
			Metadata: Metadata{Valid: false},
		}
		data, err := json.Marshal(ts)
		require.NoError(t, err)
		assert.Contains(t, string(data), `"id":789`)
		assert.Contains(t, string(data), `"metadata":null`)
	})
}

func TestMetadata_EdgeCases(t *testing.T) {
	t.Run("handles nil data slice", func(t *testing.T) {
		var m Metadata
		err := m.UnmarshalJSON(nil)
		require.NoError(t, err)
		assert.False(t, m.Valid)
		assert.True(t, m.IsEmpty())
	})

	t.Run("handles empty data slice", func(t *testing.T) {
		var m Metadata
		err := m.UnmarshalJSON([]byte{})
		assert.Error(t, err)
	})

	t.Run("handles whitespace only", func(t *testing.T) {
		var m Metadata
		err := m.UnmarshalJSON([]byte("   "))
		assert.Error(t, err)
	})
}
