package customers

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFetchCustomerResponse_JSONDeserialization(t *testing.T) {
	// Read the fetch_200.json file
	responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "customers", "fetch_200.json")
	responseData, err := os.ReadFile(responseFilePath)
	require.NoError(t, err, "failed to read fetch_200.json")

	var response FetchCustomerResponse
	err = json.Unmarshal(responseData, &response)
	require.NoError(t, err, "failed to unmarshal fetch customer response")

	// Basic validations
	assert.True(t, response.Status.Bool())
	assert.Equal(t, "Customer retrieved", response.Message)
	assert.NotNil(t, response.Data)

	// Validate customer data
	assert.Equal(t, "dom@gmail.com", response.Data.Email.String())
	assert.Equal(t, "CUS_c6wqvwmvwopw4ms", response.Data.CustomerCode.String())
	assert.Equal(t, uint64(90758908), response.Data.ID.Uint64())
	assert.False(t, response.Data.Identified.Bool())
	assert.Equal(t, "test", response.Data.Domain.String())

	// Validate related data
	assert.NotNil(t, response.Data.Transactions, "transactions should not be nil")
	assert.Empty(t, response.Data.Transactions, "transactions should be empty array")

	assert.NotNil(t, response.Data.Subscriptions, "subscriptions should not be nil")
	assert.Empty(t, response.Data.Subscriptions, "subscriptions should be empty array")

	assert.NotNil(t, response.Data.Authorizations, "authorizations should not be nil")
	assert.Len(t, response.Data.Authorizations, 1, "should have one authorization")

	// Validate authorization details
	auth := response.Data.Authorizations[0]
	assert.Equal(t, "AUTH_ekk8t49ogj", auth.AuthorizationCode.String())
	assert.Equal(t, "408408", auth.Bin.String())
	assert.Equal(t, "4081", auth.Last4.String())
	assert.Equal(t, "12", auth.ExpMonth.String())
	assert.Equal(t, "2030", auth.ExpYear.String())
}

func TestFetchCustomerResponse_FieldByFieldValidation(t *testing.T) {
	// Read the fetch_200.json file
	responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "customers", "fetch_200.json")
	responseData, err := os.ReadFile(responseFilePath)
	require.NoError(t, err, "failed to read fetch_200.json")

	// Parse the raw JSON to get the original values
	var rawResponse map[string]any
	err = json.Unmarshal(responseData, &rawResponse)
	require.NoError(t, err, "failed to unmarshal raw JSON response")

	// Deserialize into the FetchCustomerResponse struct
	var response FetchCustomerResponse
	err = json.Unmarshal(responseData, &response)
	require.NoError(t, err, "failed to unmarshal into FetchCustomerResponse struct")

	// Validate top-level fields against the raw JSON
	assert.Equal(t, rawResponse["status"], response.Status.Bool(), "status field should match")
	assert.Equal(t, rawResponse["message"], response.Message, "message field should match")

	// Validate data object fields
	rawData := rawResponse["data"].(map[string]any)

	// Basic customer fields
	assert.Equal(t, rawData["email"], response.Data.Email.String(), "email should match")
	if response.Data.Integration.Valid {
		assert.Equal(t, int(rawData["integration"].(float64)), int(response.Data.Integration.Int), "integration should match")
	}
	assert.Equal(t, rawData["domain"], response.Data.Domain.String(), "domain should match")
	assert.Equal(t, rawData["customer_code"], response.Data.CustomerCode.String(), "customer_code should match")
	assert.Equal(t, uint64(rawData["id"].(float64)), response.Data.ID.Uint64(), "id should match")
	assert.Equal(t, rawData["identified"], response.Data.Identified.Bool(), "identified should match")
	assert.Equal(t, rawData["risk_action"], response.Data.RiskAction.String(), "risk_action should match")

	// Handle nullable fields
	if rawData["first_name"] == nil {
		assert.False(t, response.Data.FirstName.Valid, "first_name should be invalid")
	} else {
		assert.Equal(t, rawData["first_name"], response.Data.FirstName.String, "first_name should match")
	}

	if rawData["last_name"] == nil {
		assert.False(t, response.Data.LastName.Valid, "last_name should be invalid")
	} else {
		assert.Equal(t, rawData["last_name"], response.Data.LastName.String, "last_name should match")
	}

	if rawData["phone"] == nil {
		assert.False(t, response.Data.Phone.Valid, "phone should be invalid")
	} else {
		assert.Equal(t, rawData["phone"], response.Data.Phone.String, "phone should match")
	}

	// Handle metadata - null in JSON becomes invalid
	if rawData["metadata"] == nil {
		assert.False(t, response.Data.Metadata.Valid, "metadata should be invalid for null")
	} else {
		assert.Equal(t, rawData["metadata"], response.Data.Metadata.Metadata, "metadata should match")
	}

	if rawData["identifications"] == nil {
		assert.False(t, response.Data.Identifications.Valid, "identifications should be invalid")
	} else {
		assert.True(t, response.Data.Identifications.Valid, "identifications should be valid when present in JSON")
	}

	if rawData["dedicated_account"] == nil {
		assert.Nil(t, response.Data.DedicatedAccount, "dedicated_account should be nil")
	} else {
		assert.NotNil(t, response.Data.DedicatedAccount, "dedicated_account should not be nil when present in JSON")
	}

	// Additional fields specific to fetch response
	assert.Equal(t, int(rawData["total_transactions"].(float64)), int(response.Data.TotalTransactions), "total_transactions should match")
	assert.Equal(t, rawData["total_transaction_value"], response.Data.TotalTransactionValue, "total_transaction_value should match")

	// For timestamp comparisons, parse both and compare the actual time values
	expectedCreatedAt, err := time.Parse(time.RFC3339, rawData["createdAt"].(string))
	require.NoError(t, err, "should parse expected createdAt")
	assert.True(t, expectedCreatedAt.Equal(response.Data.CreatedAt.Time()), "createdAt should represent the same moment")

	expectedUpdatedAt, err := time.Parse(time.RFC3339, rawData["updatedAt"].(string))
	require.NoError(t, err, "should parse expected updatedAt")
	assert.True(t, expectedUpdatedAt.Equal(response.Data.UpdatedAt.Time()), "updatedAt should represent the same moment")

	// Validate the alternative timestamp fields (created_at, updated_at)
	expectedCreatedAtSnake, err := time.Parse(time.RFC3339, rawData["created_at"].(string))
	require.NoError(t, err, "should parse expected created_at")
	assert.True(t, expectedCreatedAtSnake.Equal(response.Data.CreatedAtSnake.Time), "created_at should represent the same moment")

	expectedUpdatedAtSnake, err := time.Parse(time.RFC3339, rawData["updated_at"].(string))
	require.NoError(t, err, "should parse expected updated_at")
	assert.True(t, expectedUpdatedAtSnake.Equal(response.Data.UpdatedAtSnake.Time), "updated_at should represent the same moment")

	// Validate related arrays
	rawTransactions := rawData["transactions"].([]any)
	assert.Len(t, response.Data.Transactions, len(rawTransactions), "transactions array length should match")

	rawSubscriptions := rawData["subscriptions"].([]any)
	assert.Len(t, response.Data.Subscriptions, len(rawSubscriptions), "subscriptions array length should match")

	rawAuthorizations := rawData["authorizations"].([]any)
	assert.Len(t, response.Data.Authorizations, len(rawAuthorizations), "authorizations array length should match")

	// Validate first authorization if present
	if len(rawAuthorizations) > 0 {
		rawAuth := rawAuthorizations[0].(map[string]any)
		auth := response.Data.Authorizations[0]

		assert.Equal(t, rawAuth["authorization_code"], auth.AuthorizationCode.String(), "authorization_code should match")
		assert.Equal(t, rawAuth["bin"], auth.Bin.String(), "bin should match")
		assert.Equal(t, rawAuth["last4"], auth.Last4.String(), "last4 should match")
		assert.Equal(t, rawAuth["exp_month"], auth.ExpMonth.String(), "exp_month should match")
		assert.Equal(t, rawAuth["exp_year"], auth.ExpYear.String(), "exp_year should match")
		assert.Equal(t, rawAuth["channel"], string(auth.Channel), "channel should match")
		assert.Equal(t, rawAuth["card_type"], auth.CardType.String(), "card_type should match")
		assert.Equal(t, rawAuth["bank"], auth.Bank.String(), "bank should match")
		assert.Equal(t, rawAuth["country_code"], auth.CountryCode.String(), "country_code should match")
		assert.Equal(t, rawAuth["brand"], auth.Brand.String(), "brand should match")
		assert.Equal(t, rawAuth["reusable"], auth.Reusable.Bool(), "reusable should match")
		assert.Equal(t, rawAuth["signature"], auth.Signature.String(), "signature should match")

		// Handle nullable account_name field
		if rawAuth["account_name"] == nil {
			assert.False(t, auth.AccountName.Valid, "account_name should be invalid")
		} else {
			assert.Equal(t, rawAuth["account_name"], auth.AccountName.String(), "account_name should match")
		}
	}

	// Test round-trip serialization
	serialized, err := json.Marshal(response)
	require.NoError(t, err, "failed to marshal response back to JSON")

	var roundTripResponse FetchCustomerResponse
	err = json.Unmarshal(serialized, &roundTripResponse)
	require.NoError(t, err, "failed to unmarshal round-trip JSON")

	// Verify core fields survive round-trip
	assert.Equal(t, response.Status.Bool(), roundTripResponse.Status.Bool(), "status should survive round-trip")
	assert.Equal(t, response.Message, roundTripResponse.Message, "message should survive round-trip")
	assert.Equal(t, response.Data.Email, roundTripResponse.Data.Email, "email should survive round-trip")
	assert.Equal(t, response.Data.CustomerCode, roundTripResponse.Data.CustomerCode, "customer_code should survive round-trip")
	assert.Equal(t, response.Data.ID, roundTripResponse.Data.ID, "id should survive round-trip")
	assert.Len(t, roundTripResponse.Data.Authorizations, len(response.Data.Authorizations), "authorizations length should survive round-trip")
}
