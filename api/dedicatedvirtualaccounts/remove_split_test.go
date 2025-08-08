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
