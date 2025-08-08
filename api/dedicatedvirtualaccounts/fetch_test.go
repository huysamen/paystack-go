package dedicatedvirtualaccounts

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFetchResponse_JSONDeserialization(t *testing.T) {
	tests := []struct {
		name            string
		responseFile    string
		expectedStatus  bool
		expectedMessage string
	}{
		{
			name:            "successful fetch dedicated virtual account response",
			responseFile:    "fetch_200.json",
			expectedStatus:  true,
			expectedMessage: "Managed account successfully retrieved",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Read the response file
			responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "dedicatedvirtualaccounts", tt.responseFile)
			responseData, err := os.ReadFile(responseFilePath)
			require.NoError(t, err, "failed to read response file: %s", responseFilePath)

			// Deserialize the JSON response
			var response FetchResponse
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

				// Verify bank data
				assert.NotEmpty(t, response.Data.Bank.Name, "bank name should not be empty")
				assert.NotEmpty(t, response.Data.Bank.Slug, "bank slug should not be empty")
				assert.Greater(t, response.Data.Bank.ID, 0, "bank ID should be greater than 0")

				// Verify customer data
				if response.Data.Customer != nil {
					assert.NotEmpty(t, response.Data.Customer.Email, "customer email should not be empty")
					assert.NotEmpty(t, response.Data.Customer.CustomerCode, "customer code should not be empty")
					assert.Greater(t, response.Data.Customer.ID, uint64(0), "customer ID should be greater than 0")
				}
			}
		})
	}
}
