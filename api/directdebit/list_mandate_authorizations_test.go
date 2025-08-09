package directdebit

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/huysamen/paystack-go/enums"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListMandateAuthorizationsResponse_JSONDeserialization(t *testing.T) {
	tests := []struct {
		name            string
		responseFile    string
		expectedStatus  bool
		expectedMessage string
		expectedCount   int
	}{
		{
			name:            "successful list mandate authorizations response",
			responseFile:    "list_mandate_authorizations_200.json",
			expectedStatus:  true,
			expectedMessage: "Mandate authorizations retrieved successfully",
			expectedCount:   1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Read the response file
			responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "directdebit", tt.responseFile)
			responseData, err := os.ReadFile(responseFilePath)
			require.NoError(t, err, "failed to read response file: %s", responseFilePath)

			// Deserialize the JSON response
			var response ListMandateAuthorizationsResponse
			err = json.Unmarshal(responseData, &response)
			require.NoError(t, err, "failed to unmarshal JSON response")

			// Verify the response structure
			assert.Equal(t, tt.expectedStatus, response.Status.Bool(), "status should match")
			assert.Equal(t, tt.expectedMessage, response.Message, "message should match")

			// Only verify data structure for successful responses
			if tt.expectedStatus {
				require.NotNil(t, response.Data, "data should not be nil")
				assert.Len(t, response.Data, tt.expectedCount, "should have expected number of mandate authorizations")

				if len(response.Data) > 0 {
					mandateAuth := response.Data[0]
					assert.Greater(t, mandateAuth.ID.Int64(), int64(0), "mandate authorization ID should be greater than 0")
					assert.NotEmpty(t, mandateAuth.Status, "status should not be empty")
					assert.True(t, mandateAuth.Status.IsValid(), "status should be valid")
					assert.Greater(t, mandateAuth.MandateID.Int64(), int64(0), "mandate ID should be greater than 0")
					assert.Greater(t, mandateAuth.AuthorizationID.Int64(), int64(0), "authorization ID should be greater than 0")
					assert.NotEmpty(t, mandateAuth.AuthorizationCode.String(), "authorization code should not be empty")
					assert.Greater(t, mandateAuth.IntegrationID.Int64(), int64(0), "integration ID should be greater than 0")
					assert.NotEmpty(t, mandateAuth.AccountNumber.String(), "account number should not be empty")
					assert.NotEmpty(t, mandateAuth.BankCode.String(), "bank code should not be empty")
					assert.NotEmpty(t, mandateAuth.BankName.String(), "bank name should not be empty")
				}

				// Verify meta object if present
				if response.Meta != nil {
					assert.Greater(t, response.Meta.PerPage, 0, "per_page should be greater than 0")
					if response.Meta.Total.Valid {
						assert.Greater(t, response.Meta.Total.Int, int64(0), "total should be greater than 0")
					}
				}
			}
		})
	}
}

func TestListMandateAuthorizationsRequestBuilder(t *testing.T) {
	t.Run("builds request with no filters", func(t *testing.T) {
		builder := NewListMandateAuthorizationsRequestBuilder()
		request := builder.Build()

		assert.Empty(t, request.Cursor, "cursor should be empty")
		assert.Empty(t, request.Status, "status should be empty")
		assert.Zero(t, request.PerPage, "per_page should be zero")
	})

	t.Run("builds request with cursor", func(t *testing.T) {
		builder := NewListMandateAuthorizationsRequestBuilder().
			Cursor("next_page_cursor")

		request := builder.Build()

		assert.Equal(t, "next_page_cursor", request.Cursor, "cursor should match")
		assert.Empty(t, request.Status, "status should be empty")
		assert.Zero(t, request.PerPage, "per_page should be zero")
	})

	t.Run("builds request with status filter", func(t *testing.T) {
		builder := NewListMandateAuthorizationsRequestBuilder().
			Status(enums.MandateAuthorizationStatusActive)

		request := builder.Build()

		assert.Empty(t, request.Cursor, "cursor should be empty")
		assert.Equal(t, enums.MandateAuthorizationStatusActive, request.Status, "status should match")
		assert.Zero(t, request.PerPage, "per_page should be zero")
	})

	t.Run("builds request with all filters", func(t *testing.T) {
		builder := NewListMandateAuthorizationsRequestBuilder().
			Cursor("next_page_cursor").
			Status(enums.MandateAuthorizationStatusPending).
			PerPage(25)

		request := builder.Build()

		assert.Equal(t, "next_page_cursor", request.Cursor, "cursor should match")
		assert.Equal(t, enums.MandateAuthorizationStatusPending, request.Status, "status should match")
		assert.Equal(t, 25, request.PerPage, "per_page should match")
	})
}

func TestListMandateAuthorizationsRequest_QueryGeneration(t *testing.T) {
	t.Run("generates empty query for empty request", func(t *testing.T) {
		builder := NewListMandateAuthorizationsRequestBuilder()
		request := builder.Build()
		query := request.toQuery()

		assert.Empty(t, query, "query should be empty for empty request")
	})

	t.Run("generates query with cursor", func(t *testing.T) {
		builder := NewListMandateAuthorizationsRequestBuilder().
			Cursor("next_cursor")

		request := builder.Build()
		query := request.toQuery()

		assert.Contains(t, query, "cursor=next_cursor", "query should contain cursor parameter")
	})

	t.Run("generates query with status parameter", func(t *testing.T) {
		builder := NewListMandateAuthorizationsRequestBuilder().
			Status(enums.MandateAuthorizationStatusActive)

		request := builder.Build()
		query := request.toQuery()

		assert.Contains(t, query, "status=active", "query should contain status parameter")
	})

	t.Run("generates query with per_page parameter", func(t *testing.T) {
		builder := NewListMandateAuthorizationsRequestBuilder().
			PerPage(50)

		request := builder.Build()
		query := request.toQuery()

		assert.Contains(t, query, "per_page=50", "query should contain per_page parameter")
	})

	t.Run("generates query with multiple parameters", func(t *testing.T) {
		builder := NewListMandateAuthorizationsRequestBuilder().
			Cursor("next_cursor").
			Status(enums.MandateAuthorizationStatusInactive).
			PerPage(20)

		request := builder.Build()
		query := request.toQuery()

		assert.Contains(t, query, "cursor=next_cursor", "query should contain cursor parameter")
		assert.Contains(t, query, "status=inactive", "query should contain status parameter")
		assert.Contains(t, query, "per_page=20", "query should contain per_page parameter")
	})
}

func TestListMandateAuthorizationsResponse_FieldByFieldValidation(t *testing.T) {
	// Read the response file
	responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "directdebit", "list_mandate_authorizations_200.json")
	responseData, err := os.ReadFile(responseFilePath)
	require.NoError(t, err, "failed to read response file")

	// Parse as raw JSON to get expected values
	var rawData map[string]any
	err = json.Unmarshal(responseData, &rawData)
	require.NoError(t, err, "failed to unmarshal raw JSON")

	// Deserialize into struct
	var response ListMandateAuthorizationsResponse
	err = json.Unmarshal(responseData, &response)
	require.NoError(t, err, "failed to unmarshal into struct")

	// Validate top-level fields
	expectedStatus := rawData["status"].(bool)
	assert.Equal(t, expectedStatus, response.Status.Bool(), "status should match")
	assert.Equal(t, rawData["message"], response.Message, "message should match")

	// Validate data array
	rawDataArray := rawData["data"].([]any)
	require.NotNil(t, response.Data, "data should not be nil")
	assert.Len(t, response.Data, len(rawDataArray), "data array length should match")

	if len(rawDataArray) > 0 && len(response.Data) > 0 {
		rawMandateAuth := rawDataArray[0].(map[string]any)
		mandateAuth := response.Data[0]

		// Validate mandate authorization fields
		assert.Equal(t, int64(rawMandateAuth["id"].(float64)), mandateAuth.ID.Int64(), "id should match")
		assert.Equal(t, rawMandateAuth["status"], string(mandateAuth.Status), "status should match")
		assert.Equal(t, int64(rawMandateAuth["mandate_id"].(float64)), mandateAuth.MandateID.Int64(), "mandate_id should match")
		assert.Equal(t, int64(rawMandateAuth["authorization_id"].(float64)), mandateAuth.AuthorizationID.Int64(), "authorization_id should match")
		assert.Equal(t, rawMandateAuth["authorization_code"], mandateAuth.AuthorizationCode.String(), "authorization_code should match")
		assert.Equal(t, int64(rawMandateAuth["integration_id"].(float64)), mandateAuth.IntegrationID.Int64(), "integration_id should match")
		assert.Equal(t, rawMandateAuth["account_number"], mandateAuth.AccountNumber.String(), "account_number should match")
		assert.Equal(t, rawMandateAuth["bank_code"], mandateAuth.BankCode.String(), "bank_code should match")
		assert.Equal(t, rawMandateAuth["bank_name"], mandateAuth.BankName.String(), "bank_name should match")

		// Note: The JSON contains additional fields (customer, authorized_at) that aren't in the Go struct
		// This suggests the Go struct may need updating, but we'll test what's currently defined
	}

	// Validate meta object
	if rawData["meta"] != nil {
		rawMeta := rawData["meta"].(map[string]any)
		require.NotNil(t, response.Meta, "meta should not be nil")

		// Now that Meta struct supports both per_page and perPage, this should work correctly
		assert.Equal(t, int(rawMeta["per_page"].(float64)), response.Meta.PerPage, "meta.per_page should match")

		if rawMeta["total"] != nil && response.Meta.Total.Valid {
			assert.Equal(t, int64(rawMeta["total"].(float64)), response.Meta.Total.Int, "meta.total should match")
		}
		if rawMeta["next"] != nil && response.Meta.Next.Valid {
			assert.Equal(t, rawMeta["next"], response.Meta.Next.String(), "meta.next should match")
		}
	} // Test round-trip serialization
	serialized, err := json.Marshal(response)
	require.NoError(t, err, "should marshal back to JSON without error")

	var roundTripResponse ListMandateAuthorizationsResponse
	err = json.Unmarshal(serialized, &roundTripResponse)
	require.NoError(t, err, "should unmarshal round-trip JSON without error")

	// Verify round-trip integrity
	assert.Equal(t, response.Status.Bool(), roundTripResponse.Status.Bool(), "round-trip status should match")
	assert.Equal(t, response.Message, roundTripResponse.Message, "round-trip message should match")
	if len(response.Data) > 0 && len(roundTripResponse.Data) > 0 {
		assert.Equal(t, response.Data[0].ID.Int64(), roundTripResponse.Data[0].ID.Int64(), "round-trip data[0].id should match")
		assert.Equal(t, response.Data[0].Status, roundTripResponse.Data[0].Status, "round-trip data[0].status should match")
		assert.Equal(t, response.Data[0].AuthorizationCode.String(), roundTripResponse.Data[0].AuthorizationCode.String(), "round-trip data[0].authorization_code should match")
	}
}
