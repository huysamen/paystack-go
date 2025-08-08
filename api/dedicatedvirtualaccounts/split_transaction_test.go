package dedicatedvirtualaccounts

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSplitTransactionResponse_JSONDeserialization(t *testing.T) {
	tests := []struct {
		name            string
		responseFile    string
		expectedStatus  bool
		expectedMessage string
	}{
		{
			name:            "successful split transaction response",
			responseFile:    "split_transaction_200.json",
			expectedStatus:  true,
			expectedMessage: "Assigned Managed Account Successfully Created",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Read the response file
			responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "dedicatedvirtualaccounts", tt.responseFile)
			responseData, err := os.ReadFile(responseFilePath)
			require.NoError(t, err, "failed to read response file: %s", responseFilePath)

			// Deserialize the JSON response
			var response SplitTransactionResponse
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

				// Verify assignment data
				if response.Data.Assignment != nil {
					assert.NotEmpty(t, response.Data.Assignment.AccountType, "assignment account type should not be empty")
					assert.Greater(t, response.Data.Assignment.Integration, 0, "assignment integration should be greater than 0")
					assert.Greater(t, response.Data.Assignment.AssigneeID, 0, "assignment assignee ID should be greater than 0")
				}

				// Verify split config (should be present for split transaction)
				assert.NotNil(t, response.Data.SplitConfig, "split config should not be nil for split transaction")
			}
		})
	}
}

func TestSplitTransactionRequestBuilder(t *testing.T) {
	t.Run("builds request with customer only", func(t *testing.T) {
		builder := NewSplitTransactionRequestBuilder().
			Customer("CUS_customer_code")

		request := builder.Build()

		assert.Equal(t, "CUS_customer_code", request.Customer, "customer should match")
		assert.Empty(t, request.Subaccount, "subaccount should be empty")
		assert.Empty(t, request.SplitCode, "split code should be empty")
		assert.Empty(t, request.PreferredBank, "preferred bank should be empty")
	})

	t.Run("builds request with all fields", func(t *testing.T) {
		builder := NewSplitTransactionRequestBuilder().
			Customer("CUS_customer_code").
			Subaccount("ACCT_subaccount_code").
			SplitCode("SPL_split_code").
			PreferredBank("wema-bank")

		request := builder.Build()

		assert.Equal(t, "CUS_customer_code", request.Customer, "customer should match")
		assert.Equal(t, "ACCT_subaccount_code", request.Subaccount, "subaccount should match")
		assert.Equal(t, "SPL_split_code", request.SplitCode, "split code should match")
		assert.Equal(t, "wema-bank", request.PreferredBank, "preferred bank should match")
	})

	t.Run("builds request with empty values", func(t *testing.T) {
		builder := NewSplitTransactionRequestBuilder()
		request := builder.Build()

		assert.Empty(t, request.Customer, "customer should be empty")
		assert.Empty(t, request.Subaccount, "subaccount should be empty")
		assert.Empty(t, request.SplitCode, "split code should be empty")
		assert.Empty(t, request.PreferredBank, "preferred bank should be empty")
	})
}

func TestSplitTransactionRequest_JSONSerialization(t *testing.T) {
	t.Run("serializes request correctly", func(t *testing.T) {
		builder := NewSplitTransactionRequestBuilder().
			Customer("CUS_customer_code").
			Subaccount("ACCT_subaccount_code").
			SplitCode("SPL_split_code")

		request := builder.Build()

		jsonData, err := json.Marshal(request)
		require.NoError(t, err, "should marshal to JSON without error")

		var unmarshaled map[string]any
		err = json.Unmarshal(jsonData, &unmarshaled)
		require.NoError(t, err, "should unmarshal JSON without error")

		assert.Equal(t, "CUS_customer_code", unmarshaled["customer"], "customer should match")
		assert.Equal(t, "ACCT_subaccount_code", unmarshaled["subaccount"], "subaccount should match")
		assert.Equal(t, "SPL_split_code", unmarshaled["split_code"], "split code should match")
	})

	t.Run("omits empty fields in JSON", func(t *testing.T) {
		builder := NewSplitTransactionRequestBuilder().
			Customer("CUS_customer_code")

		request := builder.Build()

		jsonData, err := json.Marshal(request)
		require.NoError(t, err, "should marshal to JSON without error")

		var unmarshaled map[string]any
		err = json.Unmarshal(jsonData, &unmarshaled)
		require.NoError(t, err, "should unmarshal JSON without error")

		assert.Equal(t, "CUS_customer_code", unmarshaled["customer"], "customer should match")
		// Optional fields should not be present or should be empty
		_, hasSubaccount := unmarshaled["subaccount"]
		if hasSubaccount {
			assert.Empty(t, unmarshaled["subaccount"], "subaccount should be empty if present")
		}
	})
}
