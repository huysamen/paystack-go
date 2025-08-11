package paymentrequests

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTotalsResponse_JSONDeserialization(t *testing.T) {
	tests := []struct {
		name            string
		responseFile    string
		expectedStatus  bool
		expectedMessage string
	}{
		{
			name:            "successful totals response",
			responseFile:    "totals_200.json",
			expectedStatus:  true,
			expectedMessage: "Payment request totals",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Read the response file
			responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "paymentrequests", tt.responseFile)
			responseData, err := os.ReadFile(responseFilePath)
			require.NoError(t, err, "failed to read response file: %s", responseFilePath)

			// Deserialize the JSON response
			var response TotalsResponse
			err = json.Unmarshal(responseData, &response)
			require.NoError(t, err, "failed to unmarshal JSON response")

			// Verify the response structure
			assert.Equal(t, tt.expectedStatus, response.Status.Bool(), "status should match")
			assert.Equal(t, tt.expectedMessage, response.Message, "message should match")

			// Verify the data structure
			assert.NotNil(t, response.Data, "data should not be nil")
			assert.Len(t, response.Data.Pending, 2, "should have 2 pending currency amounts")
			assert.Len(t, response.Data.Successful, 2, "should have 2 successful currency amounts")
			assert.Len(t, response.Data.Total, 2, "should have 2 total currency amounts")
		})
	}
}

func TestTotalsResponse_FieldByFieldValidation(t *testing.T) {
	t.Run("totals_200_json_comprehensive_field_validation", func(t *testing.T) {
		// Read the exact JSON response file
		responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "paymentrequests", "totals_200.json")
		responseData, err := os.ReadFile(responseFilePath)
		require.NoError(t, err, "failed to read totals_200.json")

		// Parse into our struct
		var response TotalsResponse
		err = json.Unmarshal(responseData, &response)
		require.NoError(t, err, "failed to unmarshal totals_200.json")

		// Parse the raw JSON to compare exact values
		var rawJSON map[string]any
		err = json.Unmarshal(responseData, &rawJSON)
		require.NoError(t, err, "failed to parse raw JSON for comparison")

		// Field-by-field validation against the exact JSON values
		assert.Equal(t, true, rawJSON["status"], "status in JSON should be true")
		assert.Equal(t, true, response.Status.Bool(), "status in struct should be true")

		assert.Equal(t, "Payment request totals", rawJSON["message"], "message in JSON should match")
		assert.Equal(t, "Payment request totals", response.Message, "message in struct should match")

		// Verify data field exists and has correct structure
		assert.Contains(t, rawJSON, "data", "JSON should contain data field")
		assert.NotNil(t, response.Data, "struct data field should not be nil")

		// Get the data portion from raw JSON
		rawData, ok := rawJSON["data"].(map[string]any)
		require.True(t, ok, "data field should be an object")

		// Validate the structure - each category should be an array
		pendingRaw, ok := rawData["pending"].([]any)
		require.True(t, ok, "pending should be an array")
		assert.Len(t, pendingRaw, 2, "pending should have 2 currency entries")
		assert.Len(t, response.Data.Pending, 2, "struct pending should have 2 entries")

		successfulRaw, ok := rawData["successful"].([]any)
		require.True(t, ok, "successful should be an array")
		assert.Len(t, successfulRaw, 2, "successful should have 2 currency entries")
		assert.Len(t, response.Data.Successful, 2, "struct successful should have 2 entries")

		totalRaw, ok := rawData["total"].([]any)
		require.True(t, ok, "total should be an array")
		assert.Len(t, totalRaw, 2, "total should have 2 currency entries")
		assert.Len(t, response.Data.Total, 2, "struct total should have 2 entries")

		// Verify specific values for the first currency (NGN)
		ngn := response.Data.Pending[0]
		assert.Equal(t, "NGN", ngn.Currency.String(), "first pending currency should be NGN")
		assert.Equal(t, int64(42000), ngn.Amount.Int64(), "first pending amount should be 42000")

		// Verify complete JSON structure matches our struct
		reconstituted, err := json.Marshal(response)
		require.NoError(t, err, "should be able to marshal struct back to JSON")

		var reconstitutedMap map[string]any
		err = json.Unmarshal(reconstituted, &reconstitutedMap)
		require.NoError(t, err, "should be able to parse reconstituted JSON")

		// Core fields should match
		assert.Equal(t, rawJSON["status"], reconstitutedMap["status"], "status should survive round-trip")
		assert.Equal(t, rawJSON["message"], reconstitutedMap["message"], "message should survive round-trip")

		// Verify data totals survive round-trip
		reconstitutedData := reconstitutedMap["data"].(map[string]any)
		assert.Equal(t, rawData["pending"], reconstitutedData["pending"], "pending should survive round-trip")
		assert.Equal(t, rawData["successful"], reconstitutedData["successful"], "successful should survive round-trip")
		assert.Equal(t, rawData["total"], reconstitutedData["total"], "total should survive round-trip")
	})
}
