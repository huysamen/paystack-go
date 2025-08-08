package dedicatedvirtualaccounts

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

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
				assert.Greater(t, account.ID, 0, "account ID should be greater than 0")

				// Verify bank data
				assert.NotEmpty(t, account.Bank.Name, "bank name should not be empty")
				assert.NotEmpty(t, account.Bank.Slug, "bank slug should not be empty")
				assert.Greater(t, account.Bank.ID, 0, "bank ID should be greater than 0")

				// Verify customer data if present
				if account.Customer != nil {
					assert.NotEmpty(t, account.Customer.Email, "customer email should not be empty")
					assert.NotEmpty(t, account.Customer.CustomerCode, "customer code should not be empty")
					assert.Greater(t, account.Customer.ID, uint64(0), "customer ID should be greater than 0")
				}

				// Verify meta data if present
				if response.Meta != nil {
					if response.Meta.Total != nil {
						assert.Greater(t, *response.Meta.Total, 0, "meta total should be greater than 0")
					}
					assert.Greater(t, response.Meta.PerPage, 0, "meta per page should be greater than 0")
					if response.Meta.Page != nil {
						assert.GreaterOrEqual(t, *response.Meta.Page, 1, "meta page should be at least 1")
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
