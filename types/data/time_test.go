package data

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTime_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected time.Time
		wantErr  bool
	}{
		// RFC3339 format
		{
			name:     "unmarshals RFC3339 time",
			input:    `"2023-12-25T15:30:45Z"`,
			expected: time.Date(2023, 12, 25, 15, 30, 45, 0, time.UTC),
		},
		{
			name:     "unmarshals RFC3339 time with timezone",
			input:    `"2023-12-25T15:30:45+02:00"`,
			expected: time.Date(2023, 12, 25, 13, 30, 45, 0, time.UTC), // converted to UTC
		},
		// RFC3339Nano format
		{
			name:     "unmarshals RFC3339Nano time",
			input:    `"2023-12-25T15:30:45.123456789Z"`,
			expected: time.Date(2023, 12, 25, 15, 30, 45, 123456789, time.UTC),
		},
		// Custom formats
		{
			name:     "unmarshals custom format with milliseconds",
			input:    `"2023-12-25T15:30:45.000Z"`,
			expected: time.Date(2023, 12, 25, 15, 30, 45, 0, time.UTC),
		},
		{
			name:     "unmarshals simple datetime format",
			input:    `"2023-12-25 15:30:45"`,
			expected: time.Date(2023, 12, 25, 15, 30, 45, 0, time.UTC),
		},
		{
			name:     "unmarshals ISO format without timezone",
			input:    `"2023-12-25T15:30:45"`,
			expected: time.Date(2023, 12, 25, 15, 30, 45, 0, time.UTC),
		},
		{
			name:     "unmarshals date only",
			input:    `"2023-12-25"`,
			expected: time.Date(2023, 12, 25, 0, 0, 0, 0, time.UTC),
		},
		// Empty and null inputs
		{
			name:     "unmarshals empty string as zero time",
			input:    `""`,
			expected: time.Time{},
		},
		{
			name:     "unmarshals string null as zero time",
			input:    `"null"`,
			expected: time.Time{},
		},
		{
			name:     "unmarshals null as zero time",
			input:    "null",
			expected: time.Time{},
		},
		// Error cases
		{
			name:    "fails to unmarshal invalid time string",
			input:   `"not-a-time"`,
			wantErr: true,
		},
		{
			name:    "fails to unmarshal invalid format",
			input:   `"25/12/2023"`,
			wantErr: true,
		},
		{
			name:    "fails to unmarshal boolean",
			input:   "true",
			wantErr: true,
		},
		{
			name:    "fails to unmarshal number",
			input:   "123456789",
			wantErr: true,
		},
		{
			name:    "fails to unmarshal object",
			input:   `{"year": 2023}`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var tm Time
			err := json.Unmarshal([]byte(tt.input), &tm)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.True(t, tt.expected.Equal(tm.Time()), "Expected %v, got %v", tt.expected, tm.Time())
		})
	}
}

func TestTime_UnmarshalJSON_NilData(t *testing.T) {
	var tm Time
	err := tm.UnmarshalJSON(nil)
	require.NoError(t, err)
	assert.True(t, tm.IsZero())
}

func TestTime_MarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    Time
		expected string
	}{
		{
			name:     "marshals time with RFC3339 format",
			input:    Time(time.Date(2023, 12, 25, 15, 30, 45, 0, time.UTC)),
			expected: `"2023-12-25T15:30:45Z"`,
		},
		{
			name:     "marshals zero time",
			input:    Time(time.Time{}),
			expected: `"0001-01-01T00:00:00Z"`,
		},
		{
			name:     "marshals time with timezone",
			input:    Time(time.Date(2023, 12, 25, 15, 30, 45, 0, time.FixedZone("CET", 3600))),
			expected: `"2023-12-25T15:30:45+01:00"`,
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

func TestTime_Time(t *testing.T) {
	testTime := time.Date(2023, 12, 25, 15, 30, 45, 123456789, time.UTC)
	tm := Time(testTime)
	assert.True(t, testTime.Equal(tm.Time()))
}

func TestTime_String(t *testing.T) {
	tests := []struct {
		name     string
		input    Time
		expected string
	}{
		{
			name:     "returns RFC3339 string",
			input:    Time(time.Date(2023, 12, 25, 15, 30, 45, 0, time.UTC)),
			expected: "2023-12-25T15:30:45Z",
		},
		{
			name:     "returns zero time string",
			input:    Time(time.Time{}),
			expected: "0001-01-01T00:00:00Z",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.input.String())
		})
	}
}

func TestTime_IsZero(t *testing.T) {
	tests := []struct {
		name     string
		input    Time
		expected bool
	}{
		{
			name:     "returns true for zero time",
			input:    Time(time.Time{}),
			expected: true,
		},
		{
			name:     "returns false for non-zero time",
			input:    Time(time.Date(2023, 12, 25, 15, 30, 45, 0, time.UTC)),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.input.IsZero())
		})
	}
}

func TestNewTime(t *testing.T) {
	testTime := time.Date(2023, 12, 25, 15, 30, 45, 0, time.UTC)
	tm := NewTime(testTime)
	assert.True(t, testTime.Equal(tm.Time()))
}

func TestNewTimeFromString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected time.Time
		wantErr  bool
	}{
		{
			name:     "parses RFC3339 time",
			input:    "2023-12-25T15:30:45Z",
			expected: time.Date(2023, 12, 25, 15, 30, 45, 0, time.UTC),
		},
		{
			name:     "parses RFC3339Nano time",
			input:    "2023-12-25T15:30:45.123456789Z",
			expected: time.Date(2023, 12, 25, 15, 30, 45, 123456789, time.UTC),
		},
		{
			name:     "parses date only",
			input:    "2023-12-25",
			expected: time.Date(2023, 12, 25, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "handles empty string as zero time",
			input:    "",
			expected: time.Time{},
		},
		{
			name:     "handles null string as zero time",
			input:    "null",
			expected: time.Time{},
		},
		{
			name:    "fails to parse invalid format",
			input:   "invalid-time",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := NewTimeFromString(tt.input)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.True(t, tt.expected.Equal(result.Time()))
		})
	}
}

func TestTime_RoundTrip(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{
			name:  "RFC3339",
			input: `"2023-12-25T15:30:45Z"`,
		},
		{
			name:  "RFC3339Nano - may lose nanosecond precision in round trip",
			input: `"2023-12-25T15:30:45.123456789Z"`,
		},
		{
			name:  "simple datetime",
			input: `"2023-12-25 15:30:45"`,
		},
		{
			name:  "date only",
			input: `"2023-12-25"`,
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
			var tm Time
			err := json.Unmarshal([]byte(tt.input), &tm)
			require.NoError(t, err)

			// Marshal it back
			data, err := json.Marshal(tm)
			require.NoError(t, err)

			// Unmarshal again to verify consistency
			var tm2 Time
			err = json.Unmarshal(data, &tm2)
			require.NoError(t, err)

			// Should be the same (allowing for nanosecond precision loss)
			diff := tm.Time().Sub(tm2.Time())
			if diff < 0 {
				diff = -diff
			}
			assert.True(t, diff < time.Second, "Time difference too large: %v", diff)
		})
	}
}

func TestTime_InStruct(t *testing.T) {
	type TestStruct struct {
		Value Time `json:"value"`
	}

	t.Run("unmarshals in struct with time string", func(t *testing.T) {
		input := `{"value": "2023-12-25T15:30:45Z"}`
		expected := time.Date(2023, 12, 25, 15, 30, 45, 0, time.UTC)

		var ts TestStruct
		err := json.Unmarshal([]byte(input), &ts)
		require.NoError(t, err)
		assert.True(t, expected.Equal(ts.Value.Time()))
	})

	t.Run("unmarshals in struct with null value", func(t *testing.T) {
		input := `{"value": null}`
		var ts TestStruct
		err := json.Unmarshal([]byte(input), &ts)
		require.NoError(t, err)
		assert.True(t, ts.Value.IsZero())
	})

	t.Run("marshals from struct", func(t *testing.T) {
		testTime := time.Date(2023, 12, 25, 15, 30, 45, 0, time.UTC)
		ts := TestStruct{Value: Time(testTime)}
		data, err := json.Marshal(ts)
		require.NoError(t, err)
		assert.Contains(t, string(data), `"value":"2023-12-25T15:30:45Z"`)
	})
}

func TestTime_EdgeCases(t *testing.T) {
	t.Run("handles various timezone formats", func(t *testing.T) {
		timezoneTests := []struct {
			input    string
			expected time.Time
		}{
			{
				input:    `"2023-12-25T15:30:45+02:00"`,
				expected: time.Date(2023, 12, 25, 13, 30, 45, 0, time.UTC), // converted to UTC
			},
			{
				input:    `"2023-12-25T15:30:45-05:00"`,
				expected: time.Date(2023, 12, 25, 20, 30, 45, 0, time.UTC), // converted to UTC
			},
		}

		for _, tt := range timezoneTests {
			var tm Time
			err := json.Unmarshal([]byte(tt.input), &tm)
			require.NoError(t, err)
			assert.True(t, tt.expected.Equal(tm.Time()))
		}
	})

	t.Run("handles leap year dates", func(t *testing.T) {
		input := `"2024-02-29T12:00:00Z"` // leap year
		expected := time.Date(2024, 2, 29, 12, 0, 0, 0, time.UTC)

		var tm Time
		err := json.Unmarshal([]byte(input), &tm)
		require.NoError(t, err)
		assert.True(t, expected.Equal(tm.Time()))
	})
}
