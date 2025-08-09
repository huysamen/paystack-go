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
				assert.NotEmpty(t, response.Data.AccountNumber.String(), "account number should not be empty")
				assert.NotEmpty(t, response.Data.AccountName.String(), "account name should not be empty")
				assert.NotEmpty(t, response.Data.Currency.String(), "currency should not be empty")
				assert.Greater(t, response.Data.ID.Int64(), int64(0), "account ID should be greater than 0")
				assert.True(t, response.Data.Active.Bool(), "account should be active")
				assert.True(t, response.Data.Assigned.Bool(), "account should be assigned")

				// Verify bank data
				assert.NotEmpty(t, response.Data.Bank.Name.String(), "bank name should not be empty")
				assert.NotEmpty(t, response.Data.Bank.Slug.String(), "bank slug should not be empty")
				assert.Greater(t, response.Data.Bank.ID.Int64(), int64(0), "bank ID should be greater than 0")

				// Verify customer data
				if response.Data.Customer != nil {
					assert.NotEmpty(t, response.Data.Customer.Email.String(), "customer email should not be empty")
					assert.NotEmpty(t, response.Data.Customer.CustomerCode.String(), "customer code should not be empty")
					assert.Greater(t, response.Data.Customer.ID.Uint64(), uint64(0), "customer ID should be greater than 0")
				}

				// Verify assignment data
				if response.Data.Assignment != nil {
					assert.NotEmpty(t, response.Data.Assignment.AccountType.String(), "assignment account type should not be empty")
					assert.Greater(t, response.Data.Assignment.Integration.Int64(), int64(0), "assignment integration should be greater than 0")
					assert.Greater(t, response.Data.Assignment.AssigneeID.Int64(), int64(0), "assignment assignee ID should be greater than 0")
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

func TestSplitTransactionResponse_FieldByFieldValidation(t *testing.T) {
	// Read the response file
	responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "dedicatedvirtualaccounts", "split_transaction_200.json")
	responseData, err := os.ReadFile(responseFilePath)
	require.NoError(t, err, "failed to read response file")

	// Parse as raw JSON to get expected values
	var rawData map[string]any
	err = json.Unmarshal(responseData, &rawData)
	require.NoError(t, err, "failed to unmarshal raw JSON")

	// Deserialize into struct
	var response SplitTransactionResponse
	err = json.Unmarshal(responseData, &response)
	require.NoError(t, err, "failed to unmarshal into struct")

	// Validate top-level fields
	expectedStatus := rawData["status"].(bool)
	assert.Equal(t, expectedStatus, response.Status.Bool(), "status should match")
	assert.Equal(t, rawData["message"], response.Message, "message should match")

	// Validate data object
	rawDataObj := rawData["data"].(map[string]any)
	assert.Equal(t, int64(rawDataObj["id"].(float64)), response.Data.ID.Int64(), "data.id should match")
	assert.Equal(t, rawDataObj["account_name"], response.Data.AccountName.String(), "data.account_name should match")
	assert.Equal(t, rawDataObj["account_number"], response.Data.AccountNumber.String(), "data.account_number should match")
	assert.Equal(t, rawDataObj["currency"], response.Data.Currency.String(), "data.currency should match")
	assert.Equal(t, rawDataObj["assigned"], response.Data.Assigned.Bool(), "data.assigned should match")
	assert.Equal(t, rawDataObj["active"], response.Data.Active.Bool(), "data.active should match")

	// Validate timestamps
	expectedCreatedAt := rawDataObj["created_at"].(string)
	expectedUpdatedAt := rawDataObj["updated_at"].(string)
	assert.Equal(t, expectedCreatedAt, response.Data.CreatedAt.Time().Format("2006-01-02T15:04:05.000Z"), "data.created_at should match")
	assert.Equal(t, expectedUpdatedAt, response.Data.UpdatedAt.Time().Format("2006-01-02T15:04:05.000Z"), "data.updated_at should match")

	// Validate bank object
	rawBank := rawDataObj["bank"].(map[string]any)
	assert.Equal(t, int64(rawBank["id"].(float64)), response.Data.Bank.ID.Int64(), "data.bank.id should match")
	assert.Equal(t, rawBank["name"], response.Data.Bank.Name.String(), "data.bank.name should match")
	assert.Equal(t, rawBank["slug"], response.Data.Bank.Slug.String(), "data.bank.slug should match")

	// Validate customer object
	rawCustomer := rawDataObj["customer"].(map[string]any)
	assert.NotNil(t, response.Data.Customer, "data.customer should not be nil")
	assert.Equal(t, uint64(rawCustomer["id"].(float64)), response.Data.Customer.ID.Uint64(), "data.customer.id should match")
	if response.Data.Customer.FirstName.Valid && rawCustomer["first_name"] != nil {
		assert.Equal(t, rawCustomer["first_name"], response.Data.Customer.FirstName.String(), "data.customer.first_name should match")
	}
	if response.Data.Customer.LastName.Valid && rawCustomer["last_name"] != nil {
		assert.Equal(t, rawCustomer["last_name"], response.Data.Customer.LastName.String(), "data.customer.last_name should match")
	}
	assert.Equal(t, rawCustomer["email"], response.Data.Customer.Email.String(), "data.customer.email should match")
	assert.Equal(t, rawCustomer["customer_code"], response.Data.Customer.CustomerCode.String(), "data.customer.customer_code should match")

	// Validate assignment object
	rawAssignment := rawDataObj["assignment"].(map[string]any)
	assert.NotNil(t, response.Data.Assignment, "data.assignment should not be nil")
	assert.Equal(t, int64(rawAssignment["integration"].(float64)), response.Data.Assignment.Integration.Int64(), "data.assignment.integration should match")
	assert.Equal(t, int64(rawAssignment["assignee_id"].(float64)), response.Data.Assignment.AssigneeID.Int64(), "data.assignment.assignee_id should match")
	assert.Equal(t, rawAssignment["assignee_type"], response.Data.Assignment.AssigneeType.String(), "data.assignment.assignee_type should match")
	assert.Equal(t, rawAssignment["account_type"], response.Data.Assignment.AccountType.String(), "data.assignment.account_type should match")
	assert.Equal(t, rawAssignment["expired"], response.Data.Assignment.Expired.Bool(), "data.assignment.expired should match")

	// Validate assignment timestamps
	expectedAssignedAt := rawAssignment["assigned_at"].(string)
	assert.Equal(t, expectedAssignedAt, response.Data.Assignment.AssignedAt.Time().Format("2006-01-02T15:04:05.000Z"), "data.assignment.assigned_at should match")

	// Validate split_config object
	rawSplitConfig := rawDataObj["split_config"].(map[string]any)
	assert.NotNil(t, response.Data.SplitConfig, "data.split_config should not be nil")
	splitConfigMap := response.Data.SplitConfig.Metadata
	assert.Equal(t, rawSplitConfig["split_code"], splitConfigMap["split_code"], "data.split_config.split_code should match")

	// Test round-trip serialization
	serialized, err := json.Marshal(response)
	require.NoError(t, err, "should marshal back to JSON without error")

	var roundTripResponse SplitTransactionResponse
	err = json.Unmarshal(serialized, &roundTripResponse)
	require.NoError(t, err, "should unmarshal round-trip JSON without error")

	// Verify round-trip integrity
	assert.Equal(t, response.Status.Bool(), roundTripResponse.Status.Bool(), "round-trip status should match")
	assert.Equal(t, response.Message, roundTripResponse.Message, "round-trip message should match")
	assert.Equal(t, response.Data.ID.Int64(), roundTripResponse.Data.ID.Int64(), "round-trip data.id should match")
	assert.Equal(t, response.Data.AccountName.String(), roundTripResponse.Data.AccountName.String(), "round-trip data.account_name should match")
	assert.Equal(t, response.Data.AccountNumber.String(), roundTripResponse.Data.AccountNumber.String(), "round-trip data.account_number should match")
	assert.Equal(t, response.Data.Currency.String(), roundTripResponse.Data.Currency.String(), "round-trip data.currency should match")
	assert.Equal(t, response.Data.Bank.Name.String(), roundTripResponse.Data.Bank.Name.String(), "round-trip data.bank.name should match")
	assert.Equal(t, response.Data.Customer.Email.String(), roundTripResponse.Data.Customer.Email.String(), "round-trip data.customer.email should match")
}
