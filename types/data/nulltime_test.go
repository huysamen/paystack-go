package data

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNullTime_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expectedTime  string // Use string for easier comparison
		expectedValid bool
		wantErr       bool
	}{
		// RFC3339 format
		{
			name:          "unmarshals RFC3339 format",
			input:         `"2023-12-25T10:30:45Z"`,
			expectedTime:  "2023-12-25T10:30:45Z",
			expectedValid: true,
		},
		{
			name:          "unmarshals RFC3339 with timezone",
			input:         `"2023-12-25T10:30:45+02:00"`,
			expectedTime:  "2023-12-25T10:30:45+02:00",
			expectedValid: true,
		},
		{
			name:          "unmarshals RFC3339 with nanoseconds",
			input:         `"2023-12-25T10:30:45.123456789Z"`,
			expectedTime:  "2023-12-25T10:30:45.123456789Z",
			expectedValid: true,
		},
		// Custom formats
		{
			name:          "unmarshals custom format with .000Z",
			input:         `"2023-12-25T10:30:45.000Z"`,
			expectedTime:  "2023-12-25T10:30:45Z",
			expectedValid: true,
		},
		{
			name:          "unmarshals simple datetime format",
			input:         `"2023-12-25 10:30:45"`,
			expectedTime:  "2023-12-25T10:30:45Z",
			expectedValid: true,
		},
		{
			name:          "unmarshals ISO format without timezone",
			input:         `"2023-12-25T10:30:45"`,
			expectedTime:  "2023-12-25T10:30:45Z",
			expectedValid: true,
		},
		{
			name:          "unmarshals date only",
			input:         `"2023-12-25"`,
			expectedTime:  "2023-12-25T00:00:00Z",
			expectedValid: true,
		},
		// Null handling
		{
			name:          "unmarshals empty string as null",
			input:         `""`,
			expectedTime:  "",
			expectedValid: false,
		},
		{
			name:          "unmarshals string null as null",
			input:         `"null"`,
			expectedTime:  "",
			expectedValid: false,
		},
		{
			name:          "unmarshals null as invalid",
			input:         "null",
			expectedTime:  "",
			expectedValid: false,
		},
		// Error cases
		{
			name:    "fails to unmarshal invalid date string",
			input:   `"not-a-date"`,
			wantErr: true,
		},
		{
			name:    "fails to unmarshal boolean",
			input:   "true",
			wantErr: true,
		},
		{
			name:    "fails to unmarshal number",
			input:   "12345",
			wantErr: true,
		},
		{
			name:    "fails to unmarshal object",
			input:   `{"year": 2023}`,
			wantErr: true,
		},
		{
			name:    "fails to unmarshal array",
			input:   `[2023, 12, 25]`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var nt NullTime
			err := json.Unmarshal([]byte(tt.input), &nt)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.expectedValid, nt.Valid)

			if tt.expectedValid && tt.expectedTime != "" {
				expectedTime, err := time.Parse(time.RFC3339, tt.expectedTime)
				require.NoError(t, err)
				assert.True(t, nt.Time.Equal(expectedTime), "Expected %v, got %v", expectedTime, nt.Time)
			}

			if !tt.expectedValid {
				assert.True(t, nt.Time.IsZero())
			}
		})
	}
}

func TestNullTime_MarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    NullTime
		expected string
	}{
		{
			name: "marshals valid time",
			input: NullTime{
				Time:  time.Date(2023, 12, 25, 10, 30, 45, 0, time.UTC),
				Valid: true,
			},
			expected: `"2023-12-25T10:30:45Z"`,
		},
		{
			name: "marshals valid time with timezone",
			input: NullTime{
				Time:  time.Date(2023, 12, 25, 10, 30, 45, 0, time.FixedZone("CET", 2*60*60)),
				Valid: true,
			},
			expected: `"2023-12-25T10:30:45+02:00"`,
		},
		{
			name: "marshals valid time with nanoseconds",
			input: NullTime{
				Time:  time.Date(2023, 12, 25, 10, 30, 45, 123456789, time.UTC),
				Valid: true,
			},
			expected: `"2023-12-25T10:30:45.123456789Z"`,
		},
		{
			name: "marshals invalid as null (with time value)",
			input: NullTime{
				Time:  time.Date(2023, 12, 25, 10, 30, 45, 0, time.UTC),
				Valid: false,
			},
			expected: "null",
		},
		{
			name: "marshals invalid as null (with zero time)",
			input: NullTime{
				Time:  time.Time{},
				Valid: false,
			},
			expected: "null",
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

func TestNullTime_ValueOr(t *testing.T) {
	fallbackTime := time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC)
	validTime := time.Date(2023, 12, 25, 10, 30, 45, 0, time.UTC)

	tests := []struct {
		name     string
		input    NullTime
		fallback time.Time
		expected time.Time
	}{
		{
			name: "returns value when valid",
			input: NullTime{
				Time:  validTime,
				Valid: true,
			},
			fallback: fallbackTime,
			expected: validTime,
		},
		{
			name: "returns fallback when invalid",
			input: NullTime{
				Time:  validTime,
				Valid: false,
			},
			fallback: fallbackTime,
			expected: fallbackTime,
		},
		{
			name: "returns zero fallback when invalid",
			input: NullTime{
				Time:  validTime,
				Valid: false,
			},
			fallback: time.Time{},
			expected: time.Time{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.input.ValueOr(tt.fallback)
			assert.True(t, result.Equal(tt.expected))
		})
	}
}

func TestNullTime_String(t *testing.T) {
	validTime := time.Date(2023, 12, 25, 10, 30, 45, 0, time.UTC)

	tests := []struct {
		name     string
		input    NullTime
		expected string
	}{
		{
			name: "returns RFC3339 string when valid",
			input: NullTime{
				Time:  validTime,
				Valid: true,
			},
			expected: "2023-12-25T10:30:45Z",
		},
		{
			name: "returns null when invalid (with time value)",
			input: NullTime{
				Time:  validTime,
				Valid: false,
			},
			expected: "null",
		},
		{
			name: "returns null when invalid (with zero time)",
			input: NullTime{
				Time:  time.Time{},
				Valid: false,
			},
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

func TestNullTime_IsZero(t *testing.T) {
	validTime := time.Date(2023, 12, 25, 10, 30, 45, 0, time.UTC)

	tests := []struct {
		name     string
		input    NullTime
		expected bool
	}{
		{
			name: "returns false when valid and non-zero",
			input: NullTime{
				Time:  validTime,
				Valid: true,
			},
			expected: false,
		},
		{
			name: "returns true when valid but zero time",
			input: NullTime{
				Time:  time.Time{},
				Valid: true,
			},
			expected: true,
		},
		{
			name: "returns true when invalid (with time value)",
			input: NullTime{
				Time:  validTime,
				Valid: false,
			},
			expected: true,
		},
		{
			name: "returns true when invalid (with zero time)",
			input: NullTime{
				Time:  time.Time{},
				Valid: false,
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.input.IsZero()
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestNullTime_Unix(t *testing.T) {
	validTime := time.Date(2023, 12, 25, 10, 30, 45, 0, time.UTC)
	expectedUnix := validTime.Unix()

	tests := []struct {
		name     string
		input    NullTime
		expected int64
	}{
		{
			name: "returns unix timestamp when valid",
			input: NullTime{
				Time:  validTime,
				Valid: true,
			},
			expected: expectedUnix,
		},
		{
			name: "returns zero when invalid",
			input: NullTime{
				Time:  validTime,
				Valid: false,
			},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.input.Unix()
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestNewNullTime(t *testing.T) {
	validTime := time.Date(2023, 12, 25, 10, 30, 45, 0, time.UTC)

	tests := []struct {
		name     string
		input    time.Time
		expected NullTime
	}{
		{
			name:  "creates valid NullTime with time",
			input: validTime,
			expected: NullTime{
				Time:  validTime,
				Valid: true,
			},
		},
		{
			name:  "creates valid NullTime with zero time",
			input: time.Time{},
			expected: NullTime{
				Time:  time.Time{},
				Valid: true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewNullTime(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestNewNullTimeFromString(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		expectErr bool
		expected  NullTime
	}{
		{
			name:  "creates valid NullTime from RFC3339 string",
			input: "2023-12-25T10:30:45Z",
			expected: NullTime{
				Time:  time.Date(2023, 12, 25, 10, 30, 45, 0, time.UTC),
				Valid: true,
			},
		},
		{
			name:  "creates invalid NullTime from empty string",
			input: "",
			expected: NullTime{
				Time:  time.Time{},
				Valid: false,
			},
		},
		{
			name:      "returns error for invalid time string",
			input:     "not-a-date",
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := NewNullTimeFromString(tt.input)

			if tt.expectErr {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.expected.Valid, result.Valid)
			if tt.expected.Valid {
				assert.True(t, result.Time.Equal(tt.expected.Time))
			}
		})
	}
}

func TestNullTime_ActualNull(t *testing.T) {
	// Test actual JSON null (not string "null")
	jsonData := `{"value": null}`

	var result struct {
		Value NullTime `json:"value"`
	}

	err := json.Unmarshal([]byte(jsonData), &result)
	require.NoError(t, err)

	// Should be invalid (null)
	assert.False(t, result.Value.Valid, "Expected null value to be invalid")
	assert.True(t, result.Value.Time.IsZero(), "Expected zero time for null")

	// Also test direct unmarshaling of null
	var nullTime NullTime
	err = json.Unmarshal([]byte("null"), &nullTime)
	require.NoError(t, err)
	assert.False(t, nullTime.Valid, "Direct null should be invalid")
	assert.True(t, nullTime.Time.IsZero(), "Direct null should have zero time")

	// Test string "null" (this should also work)
	var stringNullTime NullTime
	err = json.Unmarshal([]byte(`"null"`), &stringNullTime)
	require.NoError(t, err)
	assert.False(t, stringNullTime.Valid, "String null should be invalid")
	assert.True(t, stringNullTime.Time.IsZero(), "String null should have zero time")

	// Test nil data slice (edge case)
	var nilDataTime NullTime
	err = nilDataTime.UnmarshalJSON(nil)
	require.NoError(t, err)
	assert.False(t, nilDataTime.Valid, "Nil data should be invalid")
	assert.True(t, nilDataTime.Time.IsZero(), "Nil data should have zero time")
}

func TestNullTime_RoundTrip(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{
			name:  "RFC3339",
			input: `"2023-12-25T10:30:45Z"`,
		},
		{
			name:  "RFC3339 with timezone",
			input: `"2023-12-25T10:30:45+02:00"`,
		},
		{
			name:  "RFC3339 with nanoseconds",
			input: `"2023-12-25T10:30:45.123456789Z"`,
		},
		{
			name:  "simple datetime",
			input: `"2023-12-25 10:30:45"`,
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
			var nt NullTime
			err := json.Unmarshal([]byte(tt.input), &nt)
			require.NoError(t, err)

			// Marshal back to JSON
			marshaled, err := json.Marshal(nt)
			require.NoError(t, err)

			// For null cases, we expect "null" output
			if tt.input == "null" || tt.input == `""` || tt.input == `"null"` {
				assert.Equal(t, "null", string(marshaled))
			} else {
				// For valid cases, unmarshal again to verify round-trip
				var nt2 NullTime
				err = json.Unmarshal(marshaled, &nt2)
				require.NoError(t, err)

				assert.True(t, nt.Time.Equal(nt2.Time))
				assert.Equal(t, nt.Valid, nt2.Valid)
			}
		})
	}
}

func TestNullTime_InStruct(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected struct {
			CreatedAt NullTime `json:"created_at"`
			UpdatedAt NullTime `json:"updated_at"`
			DeletedAt NullTime `json:"deleted_at"`
		}
	}{
		{
			name:  "unmarshals in struct with valid times",
			input: `{"created_at": "2023-12-25T10:30:45Z", "updated_at": "2023-12-25T11:30:45Z", "deleted_at": null}`,
			expected: struct {
				CreatedAt NullTime `json:"created_at"`
				UpdatedAt NullTime `json:"updated_at"`
				DeletedAt NullTime `json:"deleted_at"`
			}{
				CreatedAt: NullTime{Time: time.Date(2023, 12, 25, 10, 30, 45, 0, time.UTC), Valid: true},
				UpdatedAt: NullTime{Time: time.Date(2023, 12, 25, 11, 30, 45, 0, time.UTC), Valid: true},
				DeletedAt: NullTime{Time: time.Time{}, Valid: false},
			},
		},
		{
			name:  "unmarshals in struct with mixed formats",
			input: `{"created_at": "2023-12-25 10:30:45", "updated_at": "2023-12-25", "deleted_at": ""}`,
			expected: struct {
				CreatedAt NullTime `json:"created_at"`
				UpdatedAt NullTime `json:"updated_at"`
				DeletedAt NullTime `json:"deleted_at"`
			}{
				CreatedAt: NullTime{Time: time.Date(2023, 12, 25, 10, 30, 45, 0, time.UTC), Valid: true},
				UpdatedAt: NullTime{Time: time.Date(2023, 12, 25, 0, 0, 0, 0, time.UTC), Valid: true},
				DeletedAt: NullTime{Time: time.Time{}, Valid: false},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result struct {
				CreatedAt NullTime `json:"created_at"`
				UpdatedAt NullTime `json:"updated_at"`
				DeletedAt NullTime `json:"deleted_at"`
			}

			err := json.Unmarshal([]byte(tt.input), &result)
			require.NoError(t, err)

			assert.Equal(t, tt.expected.CreatedAt.Valid, result.CreatedAt.Valid)
			assert.Equal(t, tt.expected.UpdatedAt.Valid, result.UpdatedAt.Valid)
			assert.Equal(t, tt.expected.DeletedAt.Valid, result.DeletedAt.Valid)

			if tt.expected.CreatedAt.Valid {
				assert.True(t, tt.expected.CreatedAt.Time.Equal(result.CreatedAt.Time))
			}
			if tt.expected.UpdatedAt.Valid {
				assert.True(t, tt.expected.UpdatedAt.Time.Equal(result.UpdatedAt.Time))
			}
			if tt.expected.DeletedAt.Valid {
				assert.True(t, tt.expected.DeletedAt.Time.Equal(result.DeletedAt.Time))
			}
		})
	}

	// Test marshaling from struct
	t.Run("marshals valid times from struct", func(t *testing.T) {
		input := struct {
			CreatedAt NullTime `json:"created_at"`
			UpdatedAt NullTime `json:"updated_at"`
		}{
			CreatedAt: NullTime{Time: time.Date(2023, 12, 25, 10, 30, 45, 0, time.UTC), Valid: true},
			UpdatedAt: NullTime{Time: time.Date(2023, 12, 25, 11, 30, 45, 0, time.UTC), Valid: true},
		}

		result, err := json.Marshal(input)
		require.NoError(t, err)
		assert.JSONEq(t, `{"created_at": "2023-12-25T10:30:45Z", "updated_at": "2023-12-25T11:30:45Z"}`, string(result))
	})

	t.Run("marshals null times from struct", func(t *testing.T) {
		input := struct {
			CreatedAt NullTime `json:"created_at"`
			DeletedAt NullTime `json:"deleted_at"`
		}{
			CreatedAt: NullTime{Time: time.Date(2023, 12, 25, 10, 30, 45, 0, time.UTC), Valid: true},
			DeletedAt: NullTime{Time: time.Time{}, Valid: false},
		}

		result, err := json.Marshal(input)
		require.NoError(t, err)
		assert.JSONEq(t, `{"created_at": "2023-12-25T10:30:45Z", "deleted_at": null}`, string(result))
	})
}

func TestNullTime_CompareWithMultiDateTime(t *testing.T) {
	t.Run("null handling difference", func(t *testing.T) {
		jsonData := `{"multi": null, "null": null}`

		var result struct {
			Multi MultiDateTime `json:"multi"`
			Null  NullTime      `json:"null"`
		}

		err := json.Unmarshal([]byte(jsonData), &result)
		require.NoError(t, err)

		// MultiDateTime converts null to zero time
		assert.True(t, result.Multi.Time.IsZero())

		// NullTime preserves null state
		assert.False(t, result.Null.Valid)
		assert.True(t, result.Null.Time.IsZero())
		assert.Equal(t, "null", result.Null.String())
	})

	t.Run("empty string handling difference", func(t *testing.T) {
		jsonData := `{"multi": "", "null": ""}`

		var result struct {
			Multi MultiDateTime `json:"multi"`
			Null  NullTime      `json:"null"`
		}

		err := json.Unmarshal([]byte(jsonData), &result)
		require.NoError(t, err)

		// MultiDateTime converts empty string to zero time
		assert.True(t, result.Multi.Time.IsZero())

		// NullTime treats empty string as null
		assert.False(t, result.Null.Valid)
		assert.True(t, result.Null.Time.IsZero())
		assert.Equal(t, "null", result.Null.String())
	})
}
