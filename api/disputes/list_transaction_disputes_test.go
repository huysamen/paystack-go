package disputes

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListTransactionResponse_JSONDeserialization(t *testing.T) {
	tests := []struct {
		name            string
		responseFile    string
		expectedStatus  bool
		expectedMessage string
	}{
		{
			name:            "successful list transaction disputes response",
			responseFile:    "list_transaction_200.json",
			expectedStatus:  true,
			expectedMessage: "Dispute retrieved successfully",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Read the response file
			responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "disputes", tt.responseFile)
			responseData, err := os.ReadFile(responseFilePath)
			require.NoError(t, err, "failed to read response file: %s", responseFilePath)

			// Deserialize the JSON response
			var response ListTransactionResponse
			err = json.Unmarshal(responseData, &response)
			require.NoError(t, err, "failed to unmarshal JSON response")

			// Verify the response structure
			assert.Equal(t, tt.expectedStatus, response.Status.Bool(), "status should match")
			assert.Equal(t, "Dispute retrieved successfully", response.Message, "message should match")

			// Only verify data structure for successful responses
			if tt.expectedStatus {
				assert.NotNil(t, response.Data, "data should not be nil")

				// Verify history and messages arrays exist
				assert.NotNil(t, response.Data.History, "history should not be nil")
				assert.NotNil(t, response.Data.Messages, "messages should not be nil")

				// Verify dispute if present
				if response.Data.Dispute != nil {
					assert.Greater(t, response.Data.Dispute.ID, 0, "dispute ID should be positive")
				}
			}
		})
	}
}
