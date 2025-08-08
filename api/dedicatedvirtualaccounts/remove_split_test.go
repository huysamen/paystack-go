package dedicatedvirtualaccounts

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRemoveSplitResponse_JSONDeserialization(t *testing.T) {
	tests := []struct {
		name            string
		responseFile    string
		expectedStatus  bool
		expectedMessage string
	}{
		{
			name:            "successful remove split response",
			responseFile:    "remove_split_200.json",
			expectedStatus:  true,
			expectedMessage: "Subaccount unassigned",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Read the response file
			responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "dedicatedvirtualaccounts", tt.responseFile)
			responseData, err := os.ReadFile(responseFilePath)
			require.NoError(t, err, "failed to read response file: %s", responseFilePath)

			// Deserialize the JSON response directly - MultiBool will handle string status
			var response RemoveSplitResponse
			err = json.Unmarshal(responseData, &response)
			require.NoError(t, err, "failed to unmarshal JSON response")

			// Verify the response structure
			assert.Equal(t, tt.expectedStatus, response.Status.Bool(), "status should match")
			assert.Equal(t, tt.expectedMessage, response.Message, "message should match")

			// Only verify data structure for successful responses
			if tt.expectedStatus {
				assert.NotEmpty(t, response.Data.AccountNumber, "account number should not be empty")
				assert.NotEmpty(t, response.Data.AccountName, "account name should not be empty")
				assert.NotEmpty(t, response.Data.Currency, "currency should not be empty")
				assert.Greater(t, response.Data.ID, 0, "account ID should be greater than 0")
				assert.True(t, response.Data.Active, "account should be active")
				assert.True(t, response.Data.Assigned, "account should be assigned")

				// Verify split config is empty/removed
				assert.NotNil(t, response.Data.SplitConfig, "split config should not be nil but should be empty")
			}
		})
	}
}

func TestRemoveSplitRequestBuilder(t *testing.T) {
	t.Run("builds request with account number", func(t *testing.T) {
		builder := NewRemoveSplitRequestBuilder().
			AccountNumber("1234567890")

		request := builder.Build()

		assert.Equal(t, "1234567890", request.AccountNumber, "account number should match")
	})

	t.Run("builds request with empty values", func(t *testing.T) {
		builder := NewRemoveSplitRequestBuilder()
		request := builder.Build()

		assert.Empty(t, request.AccountNumber, "account number should be empty")
	})
}

func TestRemoveSplitRequest_JSONSerialization(t *testing.T) {
	t.Run("serializes request correctly", func(t *testing.T) {
		builder := NewRemoveSplitRequestBuilder().
			AccountNumber("1234567890")

		request := builder.Build()

		jsonData, err := json.Marshal(request)
		require.NoError(t, err, "should marshal to JSON without error")

		var unmarshaled map[string]any
		err = json.Unmarshal(jsonData, &unmarshaled)
		require.NoError(t, err, "should unmarshal JSON without error")

		assert.Equal(t, "1234567890", unmarshaled["account_number"], "account number should match")
	})
}

func TestRemoveSplitResponse_FieldByFieldValidation(t *testing.T) {
	// Read the response file
	responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "dedicatedvirtualaccounts", "remove_split_200.json")
	responseData, err := os.ReadFile(responseFilePath)
	require.NoError(t, err, "failed to read response file")

	// Parse as raw JSON to get expected values
	var rawData map[string]any
	err = json.Unmarshal(responseData, &rawData)
	require.NoError(t, err, "failed to unmarshal raw JSON")

	// Deserialize into struct
	var response RemoveSplitResponse
	err = json.Unmarshal(responseData, &response)
	require.NoError(t, err, "failed to unmarshal into struct")

	// Validate top-level fields - MultiBool doesn't have String() method
	expectedStatus := rawData["status"] == "success" || rawData["status"] == true
	assert.Equal(t, expectedStatus, response.Status.Bool(), "status should match")
	assert.Equal(t, rawData["message"], response.Message, "message should match")

	// Validate data object
	rawDataObj := rawData["data"].(map[string]any)
	assert.Equal(t, int(rawDataObj["id"].(float64)), response.Data.ID, "data.id should match")
	assert.Equal(t, rawDataObj["account_name"], response.Data.AccountName, "data.account_name should match")
	assert.Equal(t, rawDataObj["account_number"], response.Data.AccountNumber, "data.account_number should match")
	assert.Equal(t, rawDataObj["currency"], string(response.Data.Currency), "data.currency should match")
	assert.Equal(t, rawDataObj["assigned"], response.Data.Assigned, "data.assigned should match")
	assert.Equal(t, rawDataObj["active"], response.Data.Active, "data.active should match")

	// Validate timestamps
	expectedCreatedAt := rawDataObj["createdAt"].(string)
	expectedUpdatedAt := rawDataObj["updatedAt"].(string)
	if !response.Data.CreatedAt.Time.IsZero() {
		assert.Equal(t, expectedCreatedAt, response.Data.CreatedAt.Time.Format("2006-01-02T15:04:05.000Z"), "data.createdAt should match")
	}
	if !response.Data.UpdatedAt.Time.IsZero() {
		assert.Equal(t, expectedUpdatedAt, response.Data.UpdatedAt.Time.Format("2006-01-02T15:04:05.000Z"), "data.updatedAt should match")
	}

	// Validate split_config is empty after removal
	assert.NotNil(t, response.Data.SplitConfig, "data.split_config should not be nil")
	assert.Empty(t, *response.Data.SplitConfig, "data.split_config should be empty after removal")

	// Test round-trip serialization
	serialized, err := json.Marshal(response)
	require.NoError(t, err, "should marshal back to JSON without error")

	var roundTripResponse RemoveSplitResponse
	err = json.Unmarshal(serialized, &roundTripResponse)
	require.NoError(t, err, "should unmarshal round-trip JSON without error")

	// Verify round-trip integrity
	assert.Equal(t, response.Status.Bool(), roundTripResponse.Status.Bool(), "round-trip status should match")
	assert.Equal(t, response.Message, roundTripResponse.Message, "round-trip message should match")
	assert.Equal(t, response.Data.ID, roundTripResponse.Data.ID, "round-trip data.id should match")
	assert.Equal(t, response.Data.AccountName, roundTripResponse.Data.AccountName, "round-trip data.account_name should match")
	assert.Equal(t, response.Data.AccountNumber, roundTripResponse.Data.AccountNumber, "round-trip data.account_number should match")
	assert.Equal(t, response.Data.Currency, roundTripResponse.Data.Currency, "round-trip data.currency should match")
}
