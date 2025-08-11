package paymentrequests

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
		name                string
		responseFile        string
		expectedStatus      bool
		expectedMessage     string
		expectedAmount      int64
		expectedRequestCode string
	}{
		{
			name:                "successful fetch response",
			responseFile:        "fetch_200.json",
			expectedStatus:      true,
			expectedMessage:     "Payment request retrieved",
			expectedAmount:      42000,
			expectedRequestCode: "PRQ_1weqqsn2wwzgft8",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Read the response file
			responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "paymentrequests", tt.responseFile)
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
			assert.Equal(t, tt.expectedAmount, response.Data.Amount.Int64(), "amount should match")
			assert.Equal(t, tt.expectedRequestCode, response.Data.RequestCode.String(), "request code should match")
			assert.Greater(t, response.Data.ID.Int64(), int64(0), "ID should be greater than 0")
		})
	}
}

func TestFetchResponse_FieldByFieldValidation(t *testing.T) {
	t.Run("fetch_200_json_comprehensive_field_validation", func(t *testing.T) {
		// Read the exact JSON response file
		responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "paymentrequests", "fetch_200.json")
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

		assert.Equal(t, "Payment request retrieved", rawJSON["message"], "message in JSON should match")
		assert.Equal(t, "Payment request retrieved", response.Message, "message in struct should match")

		// Verify data field exists and has correct structure
		assert.Contains(t, rawJSON, "data", "JSON should contain data field")
		assert.NotNil(t, response.Data, "struct data field should not be nil")

		// Get the data portion from raw JSON
		rawData, ok := rawJSON["data"].(map[string]any)
		require.True(t, ok, "data field should be an object")

		// Validate each data field
		assert.Equal(t, float64(3136406), rawData["id"], "id in JSON should match")
		assert.Equal(t, int64(3136406), response.Data.ID.Int64(), "id in struct should match")

		assert.Equal(t, "test", rawData["domain"], "domain in JSON should match")
		assert.Equal(t, "test", response.Data.Domain.String(), "domain in struct should match")

		assert.Equal(t, float64(42000), rawData["amount"], "amount in JSON should match")
		assert.Equal(t, int64(42000), response.Data.Amount.Int64(), "amount in struct should match")

		assert.Equal(t, "NGN", rawData["currency"], "currency in JSON should match")
		assert.Equal(t, "NGN", response.Data.Currency.String(), "currency in struct should match")

		assert.Equal(t, "PRQ_1weqqsn2wwzgft8", rawData["request_code"], "request_code in JSON should match")
		assert.Equal(t, "PRQ_1weqqsn2wwzgft8", response.Data.RequestCode.String(), "request_code in struct should match")

		assert.Equal(t, "pending", rawData["status"], "status in JSON should match")
		assert.Equal(t, "pending", response.Data.Status.String(), "status in struct should match")

		assert.Equal(t, false, rawData["paid"], "paid in JSON should match")
		assert.Equal(t, false, response.Data.Paid.Bool(), "paid in struct should match")

		assert.Equal(t, "2020-06-29T16:07:33.000Z", rawData["createdAt"], "created_at in JSON should match")
		expectedCreatedAt, err := time.Parse(time.RFC3339, "2020-06-29T16:07:33.000Z")
		require.NoError(t, err, "should parse expected created_at")
		actualCreatedAt, err := time.Parse(time.RFC3339, response.Data.CreatedAt.String())
		require.NoError(t, err, "should parse actual created_at")
		assert.True(t, expectedCreatedAt.Equal(actualCreatedAt), "created_at timestamps should be equal")

		// Verify complete JSON structure matches our struct
		reconstituted, err := json.Marshal(response)
		require.NoError(t, err, "should be able to marshal struct back to JSON")

		var reconstitutedMap map[string]any
		err = json.Unmarshal(reconstituted, &reconstitutedMap)
		require.NoError(t, err, "should be able to parse reconstituted JSON")

		// Core fields should match
		assert.Equal(t, rawJSON["status"], reconstitutedMap["status"], "status should survive round-trip")
		assert.Equal(t, rawJSON["message"], reconstitutedMap["message"], "message should survive round-trip")
	})
}
