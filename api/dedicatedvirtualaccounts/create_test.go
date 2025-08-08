package dedicatedvirtualaccounts

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateResponse_JSONDeserialization(t *testing.T) {
	tests := []struct {
		name            string
		responseFile    string
		expectedStatus  bool
		expectedMessage string
	}{
		{
			name:            "successful create dedicated virtual account response",
			responseFile:    "create_200.json",
			expectedStatus:  true,
			expectedMessage: "NUBAN successfully created",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Read the response file
			responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "dedicatedvirtualaccounts", tt.responseFile)
			responseData, err := os.ReadFile(responseFilePath)
			require.NoError(t, err, "failed to read response file: %s", responseFilePath)

			// Deserialize the JSON response
			var response CreateResponse
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
			}
		})
	}
}

func TestCreateRequestBuilder(t *testing.T) {
	t.Run("builds request with customer only", func(t *testing.T) {
		builder := NewCreateRequestBuilder().
			Customer("CUS_customer_code")

		request := builder.Build()

		assert.Equal(t, "CUS_customer_code", request.Customer, "customer should match")
		assert.Empty(t, request.PreferredBank, "preferred bank should be empty")
		assert.Empty(t, request.Subaccount, "subaccount should be empty")
	})

	t.Run("builds request with all fields", func(t *testing.T) {
		builder := NewCreateRequestBuilder().
			Customer("CUS_customer_code").
			PreferredBank("wema-bank").
			Subaccount("ACCT_subaccount_code").
			SplitCode("SPL_split_code").
			FirstName("John").
			LastName("Doe").
			Phone("+2348012345678")

		request := builder.Build()

		assert.Equal(t, "CUS_customer_code", request.Customer, "customer should match")
		assert.Equal(t, "wema-bank", request.PreferredBank, "preferred bank should match")
		assert.Equal(t, "ACCT_subaccount_code", request.Subaccount, "subaccount should match")
		assert.Equal(t, "SPL_split_code", request.SplitCode, "split code should match")
		assert.Equal(t, "John", request.FirstName, "first name should match")
		assert.Equal(t, "Doe", request.LastName, "last name should match")
		assert.Equal(t, "+2348012345678", request.Phone, "phone should match")
	})

	t.Run("builds request with empty values", func(t *testing.T) {
		builder := NewCreateRequestBuilder()
		request := builder.Build()

		assert.Empty(t, request.Customer, "customer should be empty")
		assert.Empty(t, request.PreferredBank, "preferred bank should be empty")
		assert.Empty(t, request.Subaccount, "subaccount should be empty")
		assert.Empty(t, request.SplitCode, "split code should be empty")
		assert.Empty(t, request.FirstName, "first name should be empty")
		assert.Empty(t, request.LastName, "last name should be empty")
		assert.Empty(t, request.Phone, "phone should be empty")
	})
}

func TestCreateRequest_JSONSerialization(t *testing.T) {
	t.Run("serializes request correctly", func(t *testing.T) {
		builder := NewCreateRequestBuilder().
			Customer("CUS_customer_code").
			PreferredBank("wema-bank").
			FirstName("John").
			LastName("Doe")

		request := builder.Build()

		jsonData, err := json.Marshal(request)
		require.NoError(t, err, "should marshal to JSON without error")

		var unmarshaled map[string]interface{}
		err = json.Unmarshal(jsonData, &unmarshaled)
		require.NoError(t, err, "should unmarshal JSON without error")

		assert.Equal(t, "CUS_customer_code", unmarshaled["customer"], "customer should match")
		assert.Equal(t, "wema-bank", unmarshaled["preferred_bank"], "preferred bank should match")
		assert.Equal(t, "John", unmarshaled["first_name"], "first name should match")
		assert.Equal(t, "Doe", unmarshaled["last_name"], "last name should match")
	})

	t.Run("omits empty fields in JSON", func(t *testing.T) {
		builder := NewCreateRequestBuilder().
			Customer("CUS_customer_code")

		request := builder.Build()

		jsonData, err := json.Marshal(request)
		require.NoError(t, err, "should marshal to JSON without error")

		var unmarshaled map[string]interface{}
		err = json.Unmarshal(jsonData, &unmarshaled)
		require.NoError(t, err, "should unmarshal JSON without error")

		assert.Equal(t, "CUS_customer_code", unmarshaled["customer"], "customer should match")
		// Optional fields should not be present or should be empty
		_, hasPreferredBank := unmarshaled["preferred_bank"]
		if hasPreferredBank {
			assert.Empty(t, unmarshaled["preferred_bank"], "preferred bank should be empty if present")
		}
	})
}
