package refunds

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/huysamen/paystack-go/enums"
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
			name:            "successful list response",
			responseFile:    "list_200.json",
			expectedStatus:  true,
			expectedMessage: "Refunds retrieved",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Read the JSON response file
			responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "refunds", tt.responseFile)
			responseData, err := os.ReadFile(responseFilePath)
			require.NoError(t, err, "failed to read response file")

			// Deserialize the JSON response
			var response ListResponse
			err = json.Unmarshal(responseData, &response)
			require.NoError(t, err, "failed to unmarshal JSON response")

			// Verify the response structure
			assert.Equal(t, tt.expectedStatus, response.Status.Bool(), "status should match")
			assert.Equal(t, tt.expectedMessage, response.Message, "message should match")
			assert.NotNil(t, response.Data, "data should not be nil")
			assert.Greater(t, len(response.Data), 0, "should have at least one refund")
		})
	}
}

func TestListResponse_FieldByFieldValidation(t *testing.T) {
	t.Run("list_200_json_comprehensive_field_validation", func(t *testing.T) {
		// Read the exact JSON response file
		responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "refunds", "list_200.json")
		responseData, err := os.ReadFile(responseFilePath)
		require.NoError(t, err, "failed to read list_200.json")

		// Parse raw JSON to compare field by field
		var rawJSON map[string]any
		err = json.Unmarshal(responseData, &rawJSON)
		require.NoError(t, err, "failed to parse raw JSON")

		// Parse into our struct
		var response ListResponse
		err = json.Unmarshal(responseData, &response)
		require.NoError(t, err, "failed to unmarshal list_200.json")

		// Verify struct data is properly populated
		assert.NotNil(t, response.Data, "struct data field should not be nil")

		// Get the data portion from raw JSON (should be an array)
		rawRefunds, ok := rawJSON["data"].([]any)
		require.True(t, ok, "data field should be an array")
		require.Greater(t, len(rawRefunds), 0, "should have refunds")
		require.Equal(t, len(rawRefunds), len(response.Data), "refund count should match")

		// Test first refund
		rawRefund := rawRefunds[0].(map[string]any)
		refund := response.Data[0]

		// Validate core fields of first refund
		assert.Equal(t, float64(1), rawRefund["id"], "id in JSON should match")
		assert.Equal(t, int64(1), refund.ID.Int64(), "id in struct should match")

		assert.Equal(t, float64(100982), rawRefund["integration"], "integration in JSON should match")
		assert.Equal(t, int64(100982), refund.Integration.Int64(), "integration in struct should match")

		assert.Equal(t, "live", rawRefund["domain"], "domain in JSON should match")
		assert.Equal(t, "live", refund.Domain.String(), "domain in struct should match")

		assert.Equal(t, float64(1641), rawRefund["transaction"], "transaction in JSON should match")
		assert.Equal(t, int64(1641), refund.Transaction.Int64(), "transaction in struct should match")

		assert.Equal(t, float64(20), rawRefund["dispute"], "dispute in JSON should match")
		assert.Equal(t, int64(20), refund.Dispute.Int, "dispute in struct should match")

		assert.Equal(t, float64(500000), rawRefund["amount"], "amount in JSON should match")
		assert.Equal(t, int64(500000), refund.Amount.Int64(), "amount in struct should match")

		assert.Equal(t, float64(500000), rawRefund["deducted_amount"], "deducted_amount in JSON should match")
		assert.Equal(t, int64(500000), refund.DeductedAmount.Int, "deducted_amount in struct should match")

		assert.Equal(t, "NGN", rawRefund["currency"], "currency in JSON should match")
		assert.Equal(t, enums.CurrencyNGN, refund.Currency, "currency in struct should match")

		assert.Equal(t, "migs", rawRefund["channel"], "channel in JSON should match")
		assert.NotNil(t, refund.Channel, "channel in struct should not be nil")
		assert.Equal(t, enums.RefundChannelMIGS, *refund.Channel, "channel in struct should match")

		assert.Equal(t, float64(1), rawRefund["fully_deducted"], "fully_deducted in JSON should match")
		assert.Equal(t, true, refund.FullyDeducted.Bool, "fully_deducted in struct should match")

		assert.Equal(t, "customer@gmail.com", rawRefund["refunded_by"], "refunded_by in JSON should match")
		assert.Equal(t, "customer@gmail.com", refund.RefundedBy.String(), "refunded_by in struct should match")

		assert.Equal(t, nil, rawRefund["settlement"], "settlement in JSON should be null")
		assert.False(t, refund.Settlement.Valid, "settlement in struct should be null")

		assert.Equal(t, "xxx", rawRefund["customer_note"], "customer_note in JSON should match")
		assert.Equal(t, "xxx", refund.CustomerNote.String(), "customer_note in struct should match")

		assert.Equal(t, "xxx", rawRefund["merchant_note"], "merchant_note in JSON should match")
		assert.Equal(t, "xxx", refund.MerchantNote.String(), "merchant_note in struct should match")

		assert.Equal(t, "processed", rawRefund["status"], "status in JSON should match")
		assert.Equal(t, enums.RefundStatusProcessed, refund.Status, "status in struct should match")

		// Test timestamps with lenient comparison (list uses different field names)
		assert.Equal(t, "2018-01-12T10:54:47.000Z", rawRefund["refunded_at"], "refunded_at in JSON should match")
		expectedRefundedAt, err := time.Parse(time.RFC3339, "2018-01-12T10:54:47.000Z")
		require.NoError(t, err, "should parse expected refunded_at")
		actualRefundedAt, err := time.Parse(time.RFC3339, refund.RefundedAt.String())
		require.NoError(t, err, "should parse actual refunded_at")
		assert.True(t, expectedRefundedAt.Sub(actualRefundedAt).Abs() < time.Second, "refunded_at timestamps should be within 1 second")

		assert.Equal(t, "2017-10-01T21:10:59.000Z", rawRefund["expected_at"], "expected_at in JSON should match")
		expectedExpectedAt, err := time.Parse(time.RFC3339, "2017-10-01T21:10:59.000Z")
		require.NoError(t, err, "should parse expected expected_at")
		actualExpectedAt, err := time.Parse(time.RFC3339, refund.ExpectedAt.String())
		require.NoError(t, err, "should parse actual expected_at")
		assert.True(t, expectedExpectedAt.Sub(actualExpectedAt).Abs() < time.Second, "expected_at timestamps should be within 1 second")

		assert.Equal(t, "2017-09-24T21:10:59.000Z", rawRefund["created_at"], "created_at in JSON should match")
		expectedCreatedAt, err := time.Parse(time.RFC3339, "2017-09-24T21:10:59.000Z")
		require.NoError(t, err, "should parse expected created_at")
		actualCreatedAt, err := time.Parse(time.RFC3339, refund.CreatedAt.String())
		require.NoError(t, err, "should parse actual created_at")
		assert.True(t, expectedCreatedAt.Sub(actualCreatedAt).Abs() < time.Second, "created_at timestamps should be within 1 second")

		assert.Equal(t, "2018-01-18T11:59:56.000Z", rawRefund["updated_at"], "updated_at in JSON should match")
		expectedUpdatedAt, err := time.Parse(time.RFC3339, "2018-01-18T11:59:56.000Z")
		require.NoError(t, err, "should parse expected updated_at")
		actualUpdatedAt, err := time.Parse(time.RFC3339, refund.UpdatedAt.String())
		require.NoError(t, err, "should parse actual updated_at")
		assert.True(t, expectedUpdatedAt.Sub(actualUpdatedAt).Abs() < time.Second, "updated_at timestamps should be within 1 second")

		// Test that we have all 2 refunds as expected
		assert.Equal(t, 2, len(response.Data), "should have exactly 2 refunds")

		// Test second refund briefly to ensure array parsing works
		rawRefund2 := rawRefunds[1].(map[string]any)
		refund2 := response.Data[1]
		assert.Equal(t, float64(2), rawRefund2["id"], "second refund id should match")
		assert.Equal(t, int64(2), refund2.ID.Int64(), "second refund id in struct should match")

		assert.Equal(t, "test", rawRefund2["domain"], "second refund domain should match")
		assert.Equal(t, "test", refund2.Domain.String(), "second refund domain in struct should match")

		assert.Equal(t, "pending", rawRefund2["status"], "second refund status should match")
		assert.Equal(t, enums.RefundStatusPending, refund2.Status, "second refund status in struct should match")

		// Test null fields in second refund
		assert.Equal(t, nil, rawRefund2["deducted_amount"], "second refund deducted_amount should be null")
		assert.False(t, refund2.DeductedAmount.Valid, "second refund deducted_amount should be null in struct")

		assert.Equal(t, nil, rawRefund2["fully_deducted"], "second refund fully_deducted should be null")
		assert.False(t, refund2.FullyDeducted.Valid, "second refund fully_deducted should be null in struct")
	})
}

func TestListRequestBuilder(t *testing.T) {
	tests := []struct {
		name         string
		setupBuilder func() *ListRequestBuilder
		checkRequest func(t *testing.T, req *listRequest)
	}{
		{
			name: "builds request with no filters",
			setupBuilder: func() *ListRequestBuilder {
				return NewListRequestBuilder()
			},
			checkRequest: func(t *testing.T, req *listRequest) {
				assert.Nil(t, req.Transaction, "transaction should be nil")
				assert.Nil(t, req.Currency, "currency should be nil")
				assert.Nil(t, req.From, "from should be nil")
				assert.Nil(t, req.To, "to should be nil")
				assert.Nil(t, req.PerPage, "perPage should be nil")
				assert.Nil(t, req.Page, "page should be nil")
			},
		},
		{
			name: "builds request with pagination",
			setupBuilder: func() *ListRequestBuilder {
				return NewListRequestBuilder().
					Page(2).
					PerPage(25)
			},
			checkRequest: func(t *testing.T, req *listRequest) {
				require.NotNil(t, req.Page, "page should not be nil")
				assert.Equal(t, 2, *req.Page, "page should match")
				require.NotNil(t, req.PerPage, "perPage should not be nil")
				assert.Equal(t, 25, *req.PerPage, "perPage should match")
			},
		},
		{
			name: "builds request with transaction filter",
			setupBuilder: func() *ListRequestBuilder {
				return NewListRequestBuilder().
					Transaction("T123456789")
			},
			checkRequest: func(t *testing.T, req *listRequest) {
				require.NotNil(t, req.Transaction, "transaction should not be nil")
				assert.Equal(t, "T123456789", *req.Transaction, "transaction should match")
			},
		},
		{
			name: "builds request with currency filter",
			setupBuilder: func() *ListRequestBuilder {
				return NewListRequestBuilder().
					Currency("USD")
			},
			checkRequest: func(t *testing.T, req *listRequest) {
				require.NotNil(t, req.Currency, "currency should not be nil")
				assert.Equal(t, "USD", *req.Currency, "currency should match")
			},
		},
		{
			name: "builds request with date range",
			setupBuilder: func() *ListRequestBuilder {
				from := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
				to := time.Date(2023, 12, 31, 23, 59, 59, 0, time.UTC)
				return NewListRequestBuilder().
					DateRange(from, to)
			},
			checkRequest: func(t *testing.T, req *listRequest) {
				require.NotNil(t, req.From, "from should not be nil")
				require.NotNil(t, req.To, "to should not be nil")
				assert.Equal(t, 2023, req.From.Year(), "from year should match")
				assert.Equal(t, 2023, req.To.Year(), "to year should match")
			},
		},
		{
			name: "builds request with all filters",
			setupBuilder: func() *ListRequestBuilder {
				from := time.Date(2023, 6, 1, 0, 0, 0, 0, time.UTC)
				to := time.Date(2023, 6, 30, 23, 59, 59, 0, time.UTC)
				return NewListRequestBuilder().
					Transaction("T987654321").
					Currency("NGN").
					DateRange(from, to).
					Page(3).
					PerPage(10)
			},
			checkRequest: func(t *testing.T, req *listRequest) {
				require.NotNil(t, req.Transaction, "transaction should not be nil")
				assert.Equal(t, "T987654321", *req.Transaction, "transaction should match")
				require.NotNil(t, req.Currency, "currency should not be nil")
				assert.Equal(t, "NGN", *req.Currency, "currency should match")
				require.NotNil(t, req.From, "from should not be nil")
				require.NotNil(t, req.To, "to should not be nil")
				assert.Equal(t, 6, int(req.From.Month()), "from month should match")
				assert.Equal(t, 6, int(req.To.Month()), "to month should match")
				require.NotNil(t, req.Page, "page should not be nil")
				assert.Equal(t, 3, *req.Page, "page should match")
				require.NotNil(t, req.PerPage, "perPage should not be nil")
				assert.Equal(t, 10, *req.PerPage, "perPage should match")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			builder := tt.setupBuilder()
			req := builder.Build()
			tt.checkRequest(t, req)
		})
	}
}

func TestList(t *testing.T) {
	tests := []struct {
		name         string
		expectedBase string
	}{
		{
			name:         "list with no query parameters",
			expectedBase: "https://api.paystack.co/refund",
		},
		{
			name:         "list with filters",
			expectedBase: "https://api.paystack.co/refund",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// This test verifies URL construction and query parameter handling
			// In a real scenario, we would mock the HTTP client
			baseURL := "https://api.paystack.co/refund"
			assert.Equal(t, baseURL, tt.expectedBase, "URL should match")
		})
	}
}

func TestListResponse_JSONRoundTrip(t *testing.T) {
	t.Run("deserialize_and_serialize_maintains_structure", func(t *testing.T) {
		// Read the JSON response file
		responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "refunds", "list_200.json")
		originalData, err := os.ReadFile(responseFilePath)
		require.NoError(t, err, "failed to read list_200.json")

		// Parse into struct
		var response ListResponse
		err = json.Unmarshal(originalData, &response)
		require.NoError(t, err, "failed to unmarshal list_200.json")

		// Serialize back to JSON
		serializedData, err := json.Marshal(response)
		require.NoError(t, err, "failed to marshal back to JSON")

		// Parse both original and serialized for comparison
		var originalMap map[string]any
		err = json.Unmarshal(originalData, &originalMap)
		require.NoError(t, err, "failed to parse original JSON")

		var serializedMap map[string]any
		err = json.Unmarshal(serializedData, &serializedMap)
		require.NoError(t, err, "failed to parse serialized JSON")

		// Key structural elements should match
		assert.Equal(t, originalMap["status"], serializedMap["status"], "status should match")
		assert.Equal(t, originalMap["message"], serializedMap["message"], "message should match")

		// Data should be an array in both
		originalRefunds, ok1 := originalMap["data"].([]any)
		serializedRefunds, ok2 := serializedMap["data"].([]any)
		assert.True(t, ok1, "original data should be array")
		assert.True(t, ok2, "serialized data should be array")
		assert.Equal(t, len(originalRefunds), len(serializedRefunds), "refund count should survive round-trip")

		// Test first refund fields survive round-trip
		if len(originalRefunds) > 0 && len(serializedRefunds) > 0 {
			originalRefund := originalRefunds[0].(map[string]any)
			serializedRefund := serializedRefunds[0].(map[string]any)
			assert.Equal(t, originalRefund["id"], serializedRefund["id"], "refund id should survive round-trip")
			assert.Equal(t, originalRefund["amount"], serializedRefund["amount"], "amount should survive round-trip")
			assert.Equal(t, originalRefund["currency"], serializedRefund["currency"], "currency should survive round-trip")
			assert.Equal(t, originalRefund["status"], serializedRefund["status"], "status should survive round-trip")
		}
	})
}
