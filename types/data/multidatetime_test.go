package data

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMultiDateTime_MarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		dateTime MultiDateTime
		expected string
	}{
		{
			name:     "zero time",
			dateTime: MultiDateTime{},
			expected: "null",
		},
		{
			name:     "valid time",
			dateTime: NewMultiDateTime(time.Date(2023, 1, 15, 14, 30, 45, 0, time.UTC)),
			expected: `"2023-01-15T14:30:45Z"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := json.Marshal(tt.dateTime)
			require.NoError(t, err)
			assert.Equal(t, tt.expected, string(result))
		})
	}
}

func TestMultiDateTime_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected time.Time
		hasError bool
	}{
		{
			name:     "RFC3339 format",
			input:    `"2023-01-15T14:30:45Z"`,
			expected: time.Date(2023, 1, 15, 14, 30, 45, 0, time.UTC),
			hasError: false,
		},
		{
			name:     "RFC3339 with milliseconds",
			input:    `"2023-01-15T14:30:45.123Z"`,
			expected: time.Date(2023, 1, 15, 14, 30, 45, 123000000, time.UTC),
			hasError: false,
		},
		{
			name:     "custom format with .000Z",
			input:    `"2023-01-15T14:30:45.000Z"`,
			expected: time.Date(2023, 1, 15, 14, 30, 45, 0, time.UTC),
			hasError: false,
		},
		{
			name:     "simple date time format",
			input:    `"2023-01-15 14:30:45"`,
			expected: time.Date(2023, 1, 15, 14, 30, 45, 0, time.UTC),
			hasError: false,
		},
		{
			name:     "null value",
			input:    `"null"`,
			expected: time.Time{},
			hasError: false,
		},
		{
			name:     "empty string",
			input:    `""`,
			expected: time.Time{},
			hasError: false,
		},
		{
			name:     "invalid format",
			input:    `"invalid-time"`,
			expected: time.Time{},
			hasError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var dt MultiDateTime
			err := json.Unmarshal([]byte(tt.input), &dt)

			if tt.hasError {
				assert.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.True(t, tt.expected.Equal(dt.Time),
					"expected %v, got %v", tt.expected, dt.Time)
			}
		})
	}
}

func TestMultiDateTime_String(t *testing.T) {
	tests := []struct {
		name     string
		dateTime MultiDateTime
		expected string
	}{
		{
			name:     "zero time",
			dateTime: MultiDateTime{},
			expected: "",
		},
		{
			name:     "valid time",
			dateTime: NewMultiDateTime(time.Date(2023, 1, 15, 14, 30, 45, 0, time.UTC)),
			expected: "2023-01-15T14:30:45Z",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.dateTime.String()
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestMultiDateTime_Unix(t *testing.T) {
	testTime := time.Date(2023, 1, 15, 14, 30, 45, 0, time.UTC)
	dt := NewMultiDateTime(testTime)

	expected := testTime.Unix()
	actual := dt.Unix()

	assert.Equal(t, expected, actual)
}

func TestMultiDateTime_RoundTrip(t *testing.T) {
	original := NewMultiDateTime(time.Date(2023, 1, 15, 14, 30, 45, 123456789, time.UTC))

	// Marshal to JSON
	jsonData, err := json.Marshal(original)
	require.NoError(t, err)

	// Unmarshal back
	var restored MultiDateTime
	err = json.Unmarshal(jsonData, &restored)
	require.NoError(t, err)

	// Should be equal (note: nanosecond precision might be lost in JSON)
	assert.True(t, original.Time.Truncate(time.Second).Equal(restored.Time.Truncate(time.Second)))
}

func TestNewMultiDateTime(t *testing.T) {
	testTime := time.Date(2023, 1, 15, 14, 30, 45, 0, time.UTC)
	dt := NewMultiDateTime(testTime)

	assert.True(t, testTime.Equal(dt.Time))
}
