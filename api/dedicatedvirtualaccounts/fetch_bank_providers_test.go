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

func TestFetchBankProvidersResponse_FieldByFieldValidation(t *testing.T) {
	// Read the fetch_bank_providers_200.json file
	responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "dedicatedvirtualaccounts", "fetch_bank_providers_200.json")
	responseData, err := os.ReadFile(responseFilePath)
	require.NoError(t, err, "failed to read fetch_bank_providers_200.json")

	// Parse the raw JSON to get the original values
	var rawResponse map[string]any
	err = json.Unmarshal(responseData, &rawResponse)
	require.NoError(t, err, "failed to unmarshal raw JSON response")

	// Deserialize into the FetchBankProvidersResponse struct
	var response FetchBankProvidersResponse
	err = json.Unmarshal(responseData, &response)
	require.NoError(t, err, "failed to unmarshal into FetchBankProvidersResponse struct")

	// Validate top-level fields against the raw JSON
	assert.Equal(t, rawResponse["status"], response.Status.Bool(), "status field should match")
	assert.Equal(t, rawResponse["message"], response.Message, "message field should match")

	// Validate data array
	rawData, hasData := rawResponse["data"]
	require.True(t, hasData, "data field should exist")
	rawDataArray, ok := rawData.([]any)
	require.True(t, ok, "data should be an array")
	require.Len(t, rawDataArray, 2, "should have 2 items")
	require.Len(t, response.Data, 2, "response data should have 2 items")

	// Validate first provider (Access Bank)
	rawProvider1 := rawDataArray[0].(map[string]any)
	provider1 := response.Data[0]
	assert.Equal(t, rawProvider1["provider_slug"], provider1.ProviderSlug, "provider1.provider_slug should match")
	assert.Equal(t, rawProvider1["bank_id"], float64(provider1.BankID), "provider1.bank_id should match")
	assert.Equal(t, rawProvider1["bank_name"], provider1.BankName, "provider1.bank_name should match")
	assert.Equal(t, rawProvider1["id"], float64(provider1.ID), "provider1.id should match")

	// Validate second provider (Wema Bank)
	rawProvider2 := rawDataArray[1].(map[string]any)
	provider2 := response.Data[1]
	assert.Equal(t, rawProvider2["provider_slug"], provider2.ProviderSlug, "provider2.provider_slug should match")
	assert.Equal(t, rawProvider2["bank_id"], float64(provider2.BankID), "provider2.bank_id should match")
	assert.Equal(t, rawProvider2["bank_name"], provider2.BankName, "provider2.bank_name should match")
	assert.Equal(t, rawProvider2["id"], float64(provider2.ID), "provider2.id should match")

	// Test round-trip serialization
	serialized, err := json.Marshal(response)
	require.NoError(t, err, "failed to marshal response back to JSON")

	var roundTripResponse FetchBankProvidersResponse
	err = json.Unmarshal(serialized, &roundTripResponse)
	require.NoError(t, err, "failed to unmarshal round-trip JSON")

	// Verify core fields survive round-trip
	assert.Equal(t, response.Status.Bool(), roundTripResponse.Status.Bool(), "status should survive round-trip")
	assert.Equal(t, response.Message, roundTripResponse.Message, "message should survive round-trip")
	assert.Equal(t, len(response.Data), len(roundTripResponse.Data), "data array length should survive round-trip")
	assert.Equal(t, response.Data[0].BankName, roundTripResponse.Data[0].BankName, "bank_name should survive round-trip")
}
