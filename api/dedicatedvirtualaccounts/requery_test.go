package dedicatedvirtualaccounts

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRequeryResponse_JSONDeserialization(t *testing.T) {
	tests := []struct {
		name            string
		responseFile    string
		expectedStatus  bool
		expectedMessage string
	}{
		{
			name:            "successful requery dedicated virtual account response",
			responseFile:    "requery_200.json",
			expectedStatus:  true,
			expectedMessage: "We are checking the status of your transfer. We will send you a notification once it is confirmed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Read the response file
			responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "dedicatedvirtualaccounts", tt.responseFile)
			responseData, err := os.ReadFile(responseFilePath)
			require.NoError(t, err, "failed to read response file: %s", responseFilePath)

			// Deserialize the JSON response
			var response RequeryResponse
			err = json.Unmarshal(responseData, &response)
			require.NoError(t, err, "failed to unmarshal JSON response")

			// Verify the response structure
			assert.Equal(t, tt.expectedStatus, response.Status.Bool(), "status should match")
			assert.Equal(t, tt.expectedMessage, response.Message, "message should match")
		})
	}
}

func TestRequeryRequestBuilder(t *testing.T) {
	t.Run("builds request with required fields", func(t *testing.T) {
		builder := NewRequeryRequestBuilder().
			AccountNumber("1234567890").
			ProviderSlug("wema-bank")

		request := builder.Build()

		assert.Equal(t, "1234567890", request.AccountNumber, "account number should match")
		assert.Equal(t, "wema-bank", request.ProviderSlug, "provider slug should match")
		assert.Empty(t, request.Date, "date should be empty")
	})

	t.Run("builds request with all fields", func(t *testing.T) {
		builder := NewRequeryRequestBuilder().
			AccountNumber("1234567890").
			ProviderSlug("wema-bank").
			Date("2023-01-01")

		request := builder.Build()

		assert.Equal(t, "1234567890", request.AccountNumber, "account number should match")
		assert.Equal(t, "wema-bank", request.ProviderSlug, "provider slug should match")
		assert.Equal(t, "2023-01-01", request.Date, "date should match")
	})

	t.Run("builds request with empty values", func(t *testing.T) {
		builder := NewRequeryRequestBuilder()
		request := builder.Build()

		assert.Empty(t, request.AccountNumber, "account number should be empty")
		assert.Empty(t, request.ProviderSlug, "provider slug should be empty")
		assert.Empty(t, request.Date, "date should be empty")
	})
}

func TestRequeryRequest_QueryGeneration(t *testing.T) {
	t.Run("generates query with required parameters", func(t *testing.T) {
		builder := NewRequeryRequestBuilder().
			AccountNumber("1234567890").
			ProviderSlug("wema-bank")

		request := builder.Build()
		query := request.toQuery()

		assert.Contains(t, query, "account_number=1234567890", "query should contain account number parameter")
		assert.Contains(t, query, "provider_slug=wema-bank", "query should contain provider slug parameter")
	})

	t.Run("generates query with all parameters", func(t *testing.T) {
		builder := NewRequeryRequestBuilder().
			AccountNumber("1234567890").
			ProviderSlug("wema-bank").
			Date("2023-01-01")

		request := builder.Build()
		query := request.toQuery()

		assert.Contains(t, query, "account_number=1234567890", "query should contain account number parameter")
		assert.Contains(t, query, "provider_slug=wema-bank", "query should contain provider slug parameter")
		assert.Contains(t, query, "date=2023-01-01", "query should contain date parameter")
	})

	t.Run("generates query with empty date omitted", func(t *testing.T) {
		builder := NewRequeryRequestBuilder().
			AccountNumber("1234567890").
			ProviderSlug("wema-bank")

		request := builder.Build()
		query := request.toQuery()

		assert.Contains(t, query, "account_number=1234567890", "query should contain account number parameter")
		assert.Contains(t, query, "provider_slug=wema-bank", "query should contain provider slug parameter")
		assert.NotContains(t, query, "date=", "query should not contain empty date parameter")
	})
}
