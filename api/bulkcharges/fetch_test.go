package bulkcharges

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFetchResponse_JSONDeserialization(t *testing.T) {
	tests := []struct {
		name              string
		responseFile      string
		expectedStatus    bool
		expectedMessage   string
		expectedBatchCode string
	}{
		{
			name:              "successful fetch response",
			responseFile:      "fetch_200.json",
			expectedStatus:    true,
			expectedMessage:   "Bulk charge retrieved",
			expectedBatchCode: "BCH_180tl7oq7cayggh",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Read the response file
			responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "bulkcharges", tt.responseFile)
			responseData, err := os.ReadFile(responseFilePath)
			require.NoError(t, err, "failed to read response file: %s", responseFilePath)

			// Deserialize the JSON response
			var response FetchResponse
			err = json.Unmarshal(responseData, &response)
			require.NoError(t, err, "failed to unmarshal JSON response")

			// Verify the response structure
			assert.Equal(t, tt.expectedStatus, response.Status.Bool(), "status should match")
			assert.Equal(t, tt.expectedMessage, response.Message, "message should match")

			// Verify the data structure
			assert.NotNil(t, response.Data, "data should not be nil")
			assert.Equal(t, tt.expectedBatchCode, response.Data.BatchCode.String(), "batch code should match")
			assert.Greater(t, response.Data.ID.Int64(), int64(0), "ID should be greater than 0")
		})
	}
}

func TestFetchResponse_FieldByFieldValidation(t *testing.T) {
	t.Run("fetch_200_json_comprehensive_field_validation", func(t *testing.T) {
		// Read the exact JSON response file
		responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "bulkcharges", "fetch_200.json")
		responseData, err := os.ReadFile(responseFilePath)
		require.NoError(t, err, "failed to read fetch_200.json")

		// Parse into our struct
		var response FetchResponse
		err = json.Unmarshal(responseData, &response)
		require.NoError(t, err, "failed to unmarshal fetch_200.json")

		// Parse the raw JSON to compare exact values
		var rawJSON map[string]any
		err = json.Unmarshal(responseData, &rawJSON)
		require.NoError(t, err, "failed to parse raw JSON for comparison")

		// Field-by-field validation against the exact JSON values
		assert.Equal(t, true, rawJSON["status"], "status in JSON should be true")
		assert.Equal(t, true, response.Status.Bool(), "status in struct should be true")

		assert.Equal(t, "Bulk charge retrieved", rawJSON["message"], "message in JSON should match")
		assert.Equal(t, "Bulk charge retrieved", response.Message, "message in struct should match")

		// Verify data field exists and has correct structure
		assert.Contains(t, rawJSON, "data", "JSON should contain data field")
		assert.NotNil(t, response.Data, "struct data field should not be nil")

		// Get the data portion from raw JSON
		rawData, ok := rawJSON["data"].(map[string]any)
		require.True(t, ok, "data field should be an object")

		// Validate each data field
		assert.Equal(t, "test", rawData["domain"], "domain in JSON should match")
		assert.Equal(t, "test", response.Data.Domain.String(), "domain in struct should match")

		assert.Equal(t, "BCH_180tl7oq7cayggh", rawData["batch_code"], "batch_code in JSON should match")
		assert.Equal(t, "BCH_180tl7oq7cayggh", response.Data.BatchCode.String(), "batch_code in struct should match")

		assert.Equal(t, "complete", rawData["status"], "status in JSON should match")
		assert.Equal(t, "complete", response.Data.Status.String(), "status in struct should match")

		assert.Equal(t, float64(17), rawData["id"], "id in JSON should match")
		assert.Equal(t, int64(17), response.Data.ID.Int64(), "id in struct should match")

		assert.Equal(t, float64(0), rawData["total_charges"], "total_charges in JSON should match")
		assert.Equal(t, int64(0), response.Data.TotalCharges.Int64(), "total_charges in struct should match")

		assert.Equal(t, float64(0), rawData["pending_charges"], "pending_charges in JSON should match")
		assert.Equal(t, int64(0), response.Data.PendingCharges.Int64(), "pending_charges in struct should match")

		assert.Equal(t, "2017-02-04T05:44:19.000Z", rawData["createdAt"], "createdAt in JSON should match")
		// For timestamp comparison, check that we can parse both correctly rather than string format
		expectedCreatedAt, err := time.Parse(time.RFC3339, "2017-02-04T05:44:19.000Z")
		require.NoError(t, err, "should parse expected createdAt")
		actualCreatedAt, err := time.Parse(time.RFC3339, response.Data.CreatedAt.String())
		require.NoError(t, err, "should parse actual createdAt")
		assert.True(t, expectedCreatedAt.Equal(actualCreatedAt), "createdAt timestamps should be equal")

		assert.Equal(t, "2017-02-04T05:45:02.000Z", rawData["updatedAt"], "updatedAt in JSON should match")
		// For timestamp comparison, check that we can parse both correctly rather than string format
		expectedUpdatedAt, err := time.Parse(time.RFC3339, "2017-02-04T05:45:02.000Z")
		require.NoError(t, err, "should parse expected updatedAt")
		actualUpdatedAt, err := time.Parse(time.RFC3339, response.Data.UpdatedAt.String())
		require.NoError(t, err, "should parse actual updatedAt")
		assert.True(t, expectedUpdatedAt.Equal(actualUpdatedAt), "updatedAt timestamps should be equal")

		// Verify complete JSON structure matches our struct
		reconstituted, err := json.Marshal(response)
		require.NoError(t, err, "should be able to marshal struct back to JSON")

		var reconstitutedMap map[string]any
		err = json.Unmarshal(reconstituted, &reconstitutedMap)
		require.NoError(t, err, "should be able to parse reconstituted JSON")

		// Core fields should match
		assert.Equal(t, rawJSON["status"], reconstitutedMap["status"], "status should survive round-trip")
		assert.Equal(t, rawJSON["message"], reconstitutedMap["message"], "message should survive round-trip")

		// Data field should match
		reconstitutedData, ok := reconstitutedMap["data"].(map[string]any)
		require.True(t, ok, "reconstituted data should be an object")

		assert.Equal(t, rawData["domain"], reconstitutedData["domain"], "domain should survive round-trip")
		assert.Equal(t, rawData["batch_code"], reconstitutedData["batch_code"], "batch_code should survive round-trip")
		assert.Equal(t, rawData["status"], reconstitutedData["status"], "status should survive round-trip")
		assert.Equal(t, rawData["id"], reconstitutedData["id"], "id should survive round-trip")
		assert.Equal(t, rawData["total_charges"], reconstitutedData["total_charges"], "total_charges should survive round-trip")
		assert.Equal(t, rawData["pending_charges"], reconstitutedData["pending_charges"], "pending_charges should survive round-trip")

		// For timestamps, verify they represent the same moment in time
		originalCreatedAt, err := time.Parse(time.RFC3339, rawData["createdAt"].(string))
		require.NoError(t, err, "should parse original createdAt")
		roundTripCreatedAt, err := time.Parse(time.RFC3339, reconstitutedData["createdAt"].(string))
		require.NoError(t, err, "should parse round-trip createdAt")
		assert.True(t, originalCreatedAt.Equal(roundTripCreatedAt), "createdAt should survive round-trip")

		originalUpdatedAt, err := time.Parse(time.RFC3339, rawData["updatedAt"].(string))
		require.NoError(t, err, "should parse original updatedAt")
		roundTripUpdatedAt, err := time.Parse(time.RFC3339, reconstitutedData["updatedAt"].(string))
		require.NoError(t, err, "should parse round-trip updatedAt")
		assert.True(t, originalUpdatedAt.Equal(roundTripUpdatedAt), "updatedAt should survive round-trip")
	})
}
