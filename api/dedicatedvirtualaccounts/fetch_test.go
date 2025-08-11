package dedicatedvirtualaccounts

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
				assert.Greater(t, response.Data.ID.Int64(), int64(0), "account ID should be greater than 0")
				assert.True(t, response.Data.Active.Bool(), "account should be active")
				assert.True(t, response.Data.Assigned.Bool(), "account should be assigned")

				// Verify bank data
				assert.NotEmpty(t, response.Data.Bank.Name, "bank name should not be empty")
				assert.NotEmpty(t, response.Data.Bank.Slug, "bank slug should not be empty")
				assert.Greater(t, response.Data.Bank.ID.Int64(), int64(0), "bank ID should be greater than 0")

				// Verify customer data
				if response.Data.Customer != nil {
					assert.NotEmpty(t, response.Data.Customer.Email, "customer email should not be empty")
					assert.NotEmpty(t, response.Data.Customer.CustomerCode, "customer code should not be empty")
					assert.Greater(t, response.Data.Customer.ID.Uint64(), uint64(0), "customer ID should be greater than 0")
				}
			}
		})
	}
}

func TestFetchResponse_FieldByFieldValidation(t *testing.T) {
	// Read the fetch_200.json file
	responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "dedicatedvirtualaccounts", "fetch_200.json")
	responseData, err := os.ReadFile(responseFilePath)
	require.NoError(t, err, "failed to read fetch_200.json")

	// Parse the raw JSON to get the original values
	var rawResponse map[string]any
	err = json.Unmarshal(responseData, &rawResponse)
	require.NoError(t, err, "failed to unmarshal raw JSON response")

	// Deserialize into the FetchResponse struct
	var response FetchResponse
	err = json.Unmarshal(responseData, &response)
	require.NoError(t, err, "failed to unmarshal into FetchResponse struct")

	// Validate top-level fields against the raw JSON
	assert.Equal(t, rawResponse["status"], response.Status.Bool(), "status field should match")
	assert.Equal(t, rawResponse["message"], response.Message, "message field should match")

	// Validate data object fields
	rawData := rawResponse["data"].(map[string]any)
	data := response.Data

	// Basic account fields
	assert.Equal(t, rawData["account_name"], data.AccountName.String(), "account_name should match")
	assert.Equal(t, rawData["account_number"], data.AccountNumber.String(), "account_number should match")
	assert.Equal(t, rawData["assigned"], data.Assigned.Bool(), "assigned should match")
	assert.Equal(t, rawData["currency"], data.Currency.String(), "currency should match")
	assert.Equal(t, rawData["active"], data.Active.Bool(), "active should match")
	assert.Equal(t, rawData["id"], float64(data.ID.Int64()), "id should match")

	// split_config field (JSON string in response)
	if rawData["split_config"] != nil {
		// In this response, split_config is a JSON string, but in the struct it might be parsed differently
		// We'll check if it exists and is not nil
		assert.NotNil(t, data.SplitConfig, "split_config should not be nil when present in JSON")
	}

	// Bank object validation
	rawBank := rawData["bank"].(map[string]any)
	bank := data.Bank
	assert.Equal(t, rawBank["name"], bank.Name.String(), "bank.name should match")
	assert.Equal(t, rawBank["id"], float64(bank.ID.Int64()), "bank.id should match")
	assert.Equal(t, rawBank["slug"], bank.Slug.String(), "bank.slug should match")

	// Customer object validation
	rawCustomer := rawData["customer"].(map[string]any)
	customer := data.Customer
	require.NotNil(t, customer, "customer should not be nil")
	assert.Equal(t, rawCustomer["id"], float64(customer.ID.Uint64()), "customer.id should match")
	assert.Equal(t, rawCustomer["first_name"], customer.FirstName.String(), "customer.first_name should match")
	assert.Equal(t, rawCustomer["last_name"], customer.LastName.String(), "customer.last_name should match")
	assert.Equal(t, rawCustomer["email"], customer.Email.String(), "customer.email should match")
	assert.Equal(t, rawCustomer["customer_code"], customer.CustomerCode.String(), "customer.customer_code should match")
	assert.Equal(t, rawCustomer["phone"], customer.Phone.String(), "customer.phone should match")
	assert.Equal(t, rawCustomer["risk_action"], customer.RiskAction.String(), "customer.risk_action should match")

	// Customer metadata (object in this response)
	rawCustomerMetadata := rawCustomer["metadata"].(map[string]any)
	assert.NotNil(t, customer.Metadata, "customer.metadata should not be nil")
	assert.NotEmpty(t, rawCustomerMetadata, "customer metadata should not be empty")

	// international_format_phone is null
	assert.Equal(t, rawCustomer["international_format_phone"], nil, "international_format_phone should be null")

	// Timestamp validation using MultiDateTime
	createdAtStr, ok := rawData["created_at"].(string)
	require.True(t, ok, "created_at should be a string")
	parsedCreatedAt, err := time.Parse("2006-01-02T15:04:05.000Z", createdAtStr)
	require.NoError(t, err, "should parse created_at timestamp")
	assert.Equal(t, 2021, parsedCreatedAt.Year(), "created_at year should be 2021")
	assert.Equal(t, 2021, data.CreatedAt.Time().Year(), "data CreatedAt year should match")

	updatedAtStr, ok := rawData["updated_at"].(string)
	require.True(t, ok, "updated_at should be a string")
	parsedUpdatedAt, err := time.Parse("2006-01-02T15:04:05.000Z", updatedAtStr)
	require.NoError(t, err, "should parse updated_at timestamp")
	assert.Equal(t, 2021, parsedUpdatedAt.Year(), "updated_at year should be 2021")
	assert.Equal(t, 2021, data.UpdatedAt.Time().Year(), "data UpdatedAt year should match")

	// Test round-trip serialization
	serialized, err := json.Marshal(response)
	require.NoError(t, err, "failed to marshal response back to JSON")

	var roundTripResponse FetchResponse
	err = json.Unmarshal(serialized, &roundTripResponse)
	require.NoError(t, err, "failed to unmarshal round-trip JSON")

	// Verify core fields survive round-trip
	assert.Equal(t, response.Status.Bool(), roundTripResponse.Status.Bool(), "status should survive round-trip")
	assert.Equal(t, response.Message, roundTripResponse.Message, "message should survive round-trip")
	assert.Equal(t, response.Data.AccountNumber, roundTripResponse.Data.AccountNumber, "account_number should survive round-trip")
	assert.Equal(t, response.Data.AccountName, roundTripResponse.Data.AccountName, "account_name should survive round-trip")
}
