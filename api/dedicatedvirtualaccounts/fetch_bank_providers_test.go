package dedicatedvirtualaccounts

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFetchBankProvidersResponse_JSONDeserialization(t *testing.T) {
	tests := []struct {
		name            string
		responseFile    string
		expectedStatus  bool
		expectedMessage string
	}{
		{
			name:            "successful fetch bank providers response",
			responseFile:    "fetch_bank_providers_200.json",
			expectedStatus:  true,
			expectedMessage: "Dedicated account providers retrieved",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Read the response file
			responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "dedicatedvirtualaccounts", tt.responseFile)
			responseData, err := os.ReadFile(responseFilePath)
			require.NoError(t, err, "failed to read response file: %s", responseFilePath)

			// Deserialize the JSON response
			var response FetchBankProvidersResponse
			err = json.Unmarshal(responseData, &response)
			require.NoError(t, err, "failed to unmarshal JSON response")

			// Verify the response structure
			assert.Equal(t, tt.expectedStatus, response.Status.Bool(), "status should match")
			assert.Equal(t, tt.expectedMessage, response.Message, "message should match")

			// Only verify data structure for successful responses
			if tt.expectedStatus {
				assert.NotNil(t, response.Data, "data should not be nil")
				assert.Greater(t, len(response.Data), 0, "data should contain at least one provider")

				// Verify first provider structure
				provider := response.Data[0]
				assert.Greater(t, provider.ID, 0, "provider ID should be greater than 0")
				assert.NotEmpty(t, provider.ProviderSlug, "provider slug should not be empty")
				assert.Greater(t, provider.BankID, 0, "bank ID should be greater than 0")
				assert.NotEmpty(t, provider.BankName, "bank name should not be empty")

				// Verify all providers have required fields
				for _, p := range response.Data {
					assert.Greater(t, p.ID, 0, "each provider ID should be greater than 0")
					assert.NotEmpty(t, p.ProviderSlug, "each provider slug should not be empty")
					assert.Greater(t, p.BankID, 0, "each bank ID should be greater than 0")
					assert.NotEmpty(t, p.BankName, "each bank name should not be empty")
				}
			}
		})
	}
}
