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

		var unmarshaled map[string]any
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

		var unmarshaled map[string]any
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

func TestCreateResponse_FieldByFieldValidation(t *testing.T) {
	// Read the create_200.json file
	responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "dedicatedvirtualaccounts", "create_200.json")
	responseData, err := os.ReadFile(responseFilePath)
	require.NoError(t, err, "failed to read create_200.json")

	// Parse the raw JSON to get the original values
	var rawResponse map[string]any
	err = json.Unmarshal(responseData, &rawResponse)
	require.NoError(t, err, "failed to unmarshal raw JSON response")

	// Deserialize into the CreateResponse struct
	var response CreateResponse
	err = json.Unmarshal(responseData, &response)
	require.NoError(t, err, "failed to unmarshal into CreateResponse struct")

	// Validate top-level fields against the raw JSON
	assert.Equal(t, rawResponse["status"], response.Status.Bool(), "status field should match")
	assert.Equal(t, rawResponse["message"], response.Message, "message field should match")

	// Validate data object fields
	rawData := rawResponse["data"].(map[string]any)
	data := response.Data

	// Basic account fields
	assert.Equal(t, rawData["account_name"], data.AccountName, "account_name should match")
	assert.Equal(t, rawData["account_number"], data.AccountNumber, "account_number should match")
	assert.Equal(t, rawData["assigned"], data.Assigned, "assigned should match")
	assert.Equal(t, rawData["currency"], string(data.Currency), "currency should match")
	assert.Equal(t, rawData["active"], data.Active, "active should match")
	assert.Equal(t, rawData["id"], float64(data.ID), "id should match")

	// Metadata field (null in JSON)
	if rawData["metadata"] == nil {
		assert.Nil(t, data.Metadata, "metadata should be nil when null in JSON")
	} else {
		assert.Equal(t, rawData["metadata"], data.Metadata, "metadata should match")
	}

	// Bank object validation
	rawBank := rawData["bank"].(map[string]any)
	bank := data.Bank
	assert.Equal(t, rawBank["name"], bank.Name, "bank.name should match")
	assert.Equal(t, rawBank["id"], float64(bank.ID), "bank.id should match")
	assert.Equal(t, rawBank["slug"], bank.Slug, "bank.slug should match")

	// Customer object validation
	rawCustomer := rawData["customer"].(map[string]any)
	customer := data.Customer
	require.NotNil(t, customer, "customer should not be nil")
	assert.Equal(t, rawCustomer["id"], float64(customer.ID), "customer.id should match")
	assert.Equal(t, rawCustomer["first_name"], *customer.FirstName, "customer.first_name should match")
	assert.Equal(t, rawCustomer["last_name"], *customer.LastName, "customer.last_name should match")
	assert.Equal(t, rawCustomer["email"], customer.Email, "customer.email should match")
	assert.Equal(t, rawCustomer["customer_code"], customer.CustomerCode, "customer.customer_code should match")
	assert.Equal(t, rawCustomer["phone"], *customer.Phone, "customer.phone should match")
	assert.Equal(t, rawCustomer["risk_action"], customer.RiskAction, "customer.risk_action should match")

	// Assignment object validation
	rawAssignment := rawData["assignment"].(map[string]any)
	assignment := data.Assignment
	require.NotNil(t, assignment, "assignment should not be nil")
	assert.Equal(t, rawAssignment["integration"], float64(assignment.Integration), "assignment.integration should match")
	assert.Equal(t, rawAssignment["assignee_id"], float64(assignment.AssigneeID), "assignment.assignee_id should match")
	assert.Equal(t, rawAssignment["assignee_type"], assignment.AssigneeType, "assignment.assignee_type should match")
	assert.Equal(t, rawAssignment["expired"], assignment.Expired, "assignment.expired should match")
	assert.Equal(t, rawAssignment["account_type"], assignment.AccountType, "assignment.account_type should match")

	// Timestamp validation using MultiDateTime
	createdAtStr, ok := rawData["created_at"].(string)
	require.True(t, ok, "created_at should be a string")
	parsedCreatedAt, err := time.Parse("2006-01-02T15:04:05.000Z", createdAtStr)
	require.NoError(t, err, "should parse created_at timestamp")
	assert.Equal(t, 2019, parsedCreatedAt.Year(), "created_at year should be 2019")
	assert.Equal(t, 2019, data.CreatedAt.Time.Year(), "data CreatedAt year should match")

	updatedAtStr, ok := rawData["updated_at"].(string)
	require.True(t, ok, "updated_at should be a string")
	parsedUpdatedAt, err := time.Parse("2006-01-02T15:04:05.000Z", updatedAtStr)
	require.NoError(t, err, "should parse updated_at timestamp")
	assert.Equal(t, 2020, parsedUpdatedAt.Year(), "updated_at year should be 2020")
	assert.Equal(t, 2020, data.UpdatedAt.Time.Year(), "data UpdatedAt year should match")

	// Assignment timestamp validation
	assignedAtStr, ok := rawAssignment["assigned_at"].(string)
	require.True(t, ok, "assigned_at should be a string")
	parsedAssignedAt, err := time.Parse("2006-01-02T15:04:05.000Z", assignedAtStr)
	require.NoError(t, err, "should parse assigned_at timestamp")
	assert.Equal(t, 2020, parsedAssignedAt.Year(), "assigned_at year should be 2020")

	// Test round-trip serialization
	serialized, err := json.Marshal(response)
	require.NoError(t, err, "failed to marshal response back to JSON")

	var roundTripResponse CreateResponse
	err = json.Unmarshal(serialized, &roundTripResponse)
	require.NoError(t, err, "failed to unmarshal round-trip JSON")

	// Verify core fields survive round-trip
	assert.Equal(t, response.Status.Bool(), roundTripResponse.Status.Bool(), "status should survive round-trip")
	assert.Equal(t, response.Message, roundTripResponse.Message, "message should survive round-trip")
	assert.Equal(t, response.Data.AccountNumber, roundTripResponse.Data.AccountNumber, "account_number should survive round-trip")
	assert.Equal(t, response.Data.AccountName, roundTripResponse.Data.AccountName, "account_name should survive round-trip")
}
