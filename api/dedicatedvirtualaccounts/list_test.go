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

func TestListResponse_JSONDeserialization(t *testing.T) {
	tests := []struct {
		name            string
		responseFile    string
		expectedStatus  bool
		expectedMessage string
	}{
		{
			name:            "successful list dedicated virtual accounts response",
			responseFile:    "list_200.json",
			expectedStatus:  true,
			expectedMessage: "Managed accounts successfully retrieved",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Read the response file
			responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "dedicatedvirtualaccounts", tt.responseFile)
			responseData, err := os.ReadFile(responseFilePath)
			require.NoError(t, err, "failed to read response file: %s", responseFilePath)

			// Deserialize the JSON response
			var response ListResponse
			err = json.Unmarshal(responseData, &response)
			require.NoError(t, err, "failed to unmarshal JSON response")

			// Verify the response structure
			assert.Equal(t, tt.expectedStatus, response.Status.Bool(), "status should match")
			assert.Equal(t, tt.expectedMessage, response.Message, "message should match")

			// Only verify data structure for successful responses
			if tt.expectedStatus {
				assert.NotNil(t, response.Data, "data should not be nil")
				assert.Greater(t, len(response.Data), 0, "data should contain at least one account")

				// Verify first account structure
				account := response.Data[0]
				assert.NotEmpty(t, account.AccountNumber, "account number should not be empty")
				assert.NotEmpty(t, account.AccountName, "account name should not be empty")
				assert.NotEmpty(t, account.Currency, "currency should not be empty")
				assert.Greater(t, account.ID.Int64(), int64(0), "account ID should be greater than 0")

				// Verify bank data
				assert.NotEmpty(t, account.Bank.Name, "bank name should not be empty")
				assert.NotEmpty(t, account.Bank.Slug, "bank slug should not be empty")
				assert.Greater(t, account.Bank.ID.Int64(), int64(0), "bank ID should be greater than 0")

				// Verify customer data if present
				if account.Customer != nil {
					assert.NotEmpty(t, account.Customer.Email, "customer email should not be empty")
					assert.NotEmpty(t, account.Customer.CustomerCode, "customer code should not be empty")
					assert.Greater(t, account.Customer.ID.Uint64(), uint64(0), "customer ID should be greater than 0")
				}

				// Verify meta data if present
				if response.Meta != nil {
					if response.Meta.Total.Valid {
						assert.Greater(t, response.Meta.Total.Int, int64(0), "meta total should be greater than 0")
					}
					assert.Greater(t, response.Meta.PerPage, 0, "meta per page should be greater than 0")
					if response.Meta.Page.Valid {
						assert.GreaterOrEqual(t, response.Meta.Page.Int, int64(1), "meta page should be at least 1")
					}
				}
			}
		})
	}
}

func TestListRequestBuilder(t *testing.T) {
	t.Run("builds request with no filters", func(t *testing.T) {
		builder := NewListRequestBuilder()
		request := builder.Build()

		assert.Nil(t, request.Active, "active should be nil")
		assert.Empty(t, request.Currency, "currency should be empty")
		assert.Empty(t, request.ProviderSlug, "provider slug should be empty")
		assert.Empty(t, request.BankID, "bank ID should be empty")
		assert.Empty(t, request.Customer, "customer should be empty")
	})

	t.Run("builds request with active filter", func(t *testing.T) {
		builder := NewListRequestBuilder().
			Active(true)

		request := builder.Build()

		assert.NotNil(t, request.Active, "active should not be nil")
		assert.True(t, *request.Active, "active should be true")
	})

	t.Run("builds request with all filters", func(t *testing.T) {
		builder := NewListRequestBuilder().
			Active(false).
			Currency("NGN").
			ProviderSlug("wema-bank").
			BankID("20").
			Customer("CUS_customer_code")

		request := builder.Build()

		assert.NotNil(t, request.Active, "active should not be nil")
		assert.False(t, *request.Active, "active should be false")
		assert.Equal(t, "NGN", request.Currency, "currency should match")
		assert.Equal(t, "wema-bank", request.ProviderSlug, "provider slug should match")
		assert.Equal(t, "20", request.BankID, "bank ID should match")
		assert.Equal(t, "CUS_customer_code", request.Customer, "customer should match")
	})
}

func TestListRequest_QueryGeneration(t *testing.T) {
	t.Run("generates empty query for empty request", func(t *testing.T) {
		builder := NewListRequestBuilder()
		request := builder.Build()

		query := request.toQuery()
		assert.Empty(t, query, "query should be empty")
	})

	t.Run("generates query with single parameter", func(t *testing.T) {
		builder := NewListRequestBuilder().
			Currency("NGN")

		request := builder.Build()
		query := request.toQuery()

		assert.Equal(t, "currency=NGN", query, "query should match")
	})

	t.Run("generates query with multiple parameters", func(t *testing.T) {
		builder := NewListRequestBuilder().
			Active(true).
			Currency("NGN").
			ProviderSlug("wema-bank")

		request := builder.Build()
		query := request.toQuery()

		assert.Contains(t, query, "active=true", "query should contain active parameter")
		assert.Contains(t, query, "currency=NGN", "query should contain currency parameter")
		assert.Contains(t, query, "provider_slug=wema-bank", "query should contain provider_slug parameter")
	})

	t.Run("generates query with boolean false", func(t *testing.T) {
		builder := NewListRequestBuilder().
			Active(false)

		request := builder.Build()
		query := request.toQuery()

		assert.Equal(t, "active=false", query, "query should match")
	})
}

func TestListResponse_FieldByFieldValidation(t *testing.T) {
	// Read the list_200.json file
	responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "dedicatedvirtualaccounts", "list_200.json")
	responseData, err := os.ReadFile(responseFilePath)
	require.NoError(t, err, "failed to read list_200.json")

	// Parse the raw JSON to get the original values
	var rawResponse map[string]any
	err = json.Unmarshal(responseData, &rawResponse)
	require.NoError(t, err, "failed to unmarshal raw JSON response")

	// Deserialize into the ListResponse struct
	var response ListResponse
	err = json.Unmarshal(responseData, &response)
	require.NoError(t, err, "failed to unmarshal into ListResponse struct")

	// Validate top-level fields against the raw JSON
	assert.Equal(t, rawResponse["status"], response.Status.Bool(), "status field should match")
	assert.Equal(t, rawResponse["message"], response.Message, "message field should match")

	// Validate data array
	rawData, hasData := rawResponse["data"]
	require.True(t, hasData, "data field should exist")
	rawDataArray, ok := rawData.([]any)
	require.True(t, ok, "data should be an array")
	require.Len(t, rawDataArray, 1, "should have 1 item")
	require.Len(t, response.Data, 1, "response data should have 1 item")

	// Validate first account
	rawAccount := rawDataArray[0].(map[string]any)
	account := response.Data[0]

	// Basic account fields
	assert.Equal(t, rawAccount["account_name"], account.AccountName.String(), "account_name should match")
	assert.Equal(t, rawAccount["account_number"], account.AccountNumber.String(), "account_number should match")
	assert.Equal(t, rawAccount["assigned"], account.Assigned.Bool(), "assigned should match")
	assert.Equal(t, rawAccount["currency"], account.Currency.String(), "currency should match")
	assert.Equal(t, rawAccount["active"], account.Active.Bool(), "active should match")
	assert.Equal(t, rawAccount["id"], float64(account.ID.Int64()), "id should match")

	// Bank object validation
	rawBank := rawAccount["bank"].(map[string]any)
	bank := account.Bank
	assert.Equal(t, rawBank["name"], bank.Name.String(), "bank.name should match")
	assert.Equal(t, rawBank["id"], float64(bank.ID.Int64()), "bank.id should match")
	assert.Equal(t, rawBank["slug"], bank.Slug.String(), "bank.slug should match")

	// Customer object validation
	rawCustomer := rawAccount["customer"].(map[string]any)
	customer := account.Customer
	require.NotNil(t, customer, "customer should not be nil")
	assert.Equal(t, rawCustomer["id"], float64(customer.ID.Uint64()), "customer.id should match")
	assert.Equal(t, rawCustomer["first_name"], customer.FirstName.String(), "customer.first_name should match")
	assert.Equal(t, rawCustomer["last_name"], customer.LastName.String(), "customer.last_name should match")
	assert.Equal(t, rawCustomer["email"], customer.Email.String(), "customer.email should match")
	assert.Equal(t, rawCustomer["customer_code"], customer.CustomerCode.String(), "customer.customer_code should match")
	assert.Equal(t, rawCustomer["phone"], customer.Phone.String(), "customer.phone should match")
	assert.Equal(t, rawCustomer["risk_action"], customer.RiskAction.String(), "customer.risk_action should match")

	// international_format_phone is null
	assert.Equal(t, rawCustomer["international_format_phone"], nil, "international_format_phone should be null")

	// split_config object validation (is a map in this response)
	rawSplitConfig := rawAccount["split_config"].(map[string]any)
	assert.NotNil(t, account.SplitConfig, "split_config should not be nil")
	assert.NotEmpty(t, rawSplitConfig, "split_config should not be empty")

	// Timestamp validation using MultiDateTime
	createdAtStr, ok := rawAccount["created_at"].(string)
	require.True(t, ok, "created_at should be a string")
	parsedCreatedAt, err := time.Parse("2006-01-02T15:04:05.000Z", createdAtStr)
	require.NoError(t, err, "should parse created_at timestamp")
	assert.Equal(t, 2019, parsedCreatedAt.Year(), "created_at year should be 2019")
	assert.Equal(t, 2019, account.CreatedAt.Time().Year(), "account CreatedAt year should match")

	updatedAtStr, ok := rawAccount["updated_at"].(string)
	require.True(t, ok, "updated_at should be a string")
	parsedUpdatedAt, err := time.Parse("2006-01-02T15:04:05.000Z", updatedAtStr)
	require.NoError(t, err, "should parse updated_at timestamp")
	assert.Equal(t, 2020, parsedUpdatedAt.Year(), "updated_at year should be 2020")
	assert.Equal(t, 2020, account.UpdatedAt.Time().Year(), "account UpdatedAt year should match")

	// Validate meta fields
	rawMeta := rawResponse["meta"].(map[string]any)
	require.NotNil(t, response.Meta, "meta should not be nil")
	meta := response.Meta
	assert.Equal(t, rawMeta["total"], float64(meta.Total.Int), "meta.total should match")
	assert.Equal(t, rawMeta["skipped"], float64(meta.Skipped.Int), "meta.skipped should match")
	assert.Equal(t, rawMeta["perPage"], float64(meta.PerPage), "meta.perPage should match")
	assert.Equal(t, rawMeta["page"], float64(meta.Page.Int), "meta.page should match")
	assert.Equal(t, rawMeta["pageCount"], float64(meta.PageCount.Int), "meta.pageCount should match")

	// Test round-trip serialization
	serialized, err := json.Marshal(response)
	require.NoError(t, err, "failed to marshal response back to JSON")

	var roundTripResponse ListResponse
	err = json.Unmarshal(serialized, &roundTripResponse)
	require.NoError(t, err, "failed to unmarshal round-trip JSON")

	// Verify core fields survive round-trip
	assert.Equal(t, response.Status.Bool(), roundTripResponse.Status.Bool(), "status should survive round-trip")
	assert.Equal(t, response.Message, roundTripResponse.Message, "message should survive round-trip")
	assert.Equal(t, len(response.Data), len(roundTripResponse.Data), "data array length should survive round-trip")
	assert.Equal(t, response.Data[0].AccountNumber, roundTripResponse.Data[0].AccountNumber, "account_number should survive round-trip")
}
