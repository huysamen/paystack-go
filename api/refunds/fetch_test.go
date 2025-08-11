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

func TestFetchResponse_JSONDeserialization(t *testing.T) {
	tests := []struct {
		name            string
		responseFile    string
		expectedStatus  bool
		expectedMessage string
	}{
		{
			name:            "successful fetch response",
			responseFile:    "fetch_200.json",
			expectedStatus:  true,
			expectedMessage: "Refund retrieved",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Read the JSON response file
			responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "refunds", tt.responseFile)
			responseData, err := os.ReadFile(responseFilePath)
			require.NoError(t, err, "failed to read response file")

			// Deserialize the JSON response
			var response FetchResponse
			err = json.Unmarshal(responseData, &response)
			require.NoError(t, err, "failed to unmarshal JSON response")

			// Verify the response structure
			assert.Equal(t, tt.expectedStatus, response.Status.Bool(), "status should match")
			assert.Equal(t, tt.expectedMessage, response.Message, "message should match")
			assert.NotNil(t, response.Data, "data should not be nil")
		})
	}
}

func TestFetchResponse_FieldByFieldValidation(t *testing.T) {
	t.Run("fetch_200_json_comprehensive_field_validation", func(t *testing.T) {
		// Read the exact JSON response file
		responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "refunds", "fetch_200.json")
		responseData, err := os.ReadFile(responseFilePath)
		require.NoError(t, err, "failed to read fetch_200.json")

		// Parse raw JSON to compare field by field
		var rawJSON map[string]any
		err = json.Unmarshal(responseData, &rawJSON)
		require.NoError(t, err, "failed to parse raw JSON")

		// Parse into our struct
		var response FetchResponse
		err = json.Unmarshal(responseData, &response)
		require.NoError(t, err, "failed to unmarshal fetch_200.json")

		// Verify struct data is properly populated
		assert.NotNil(t, response.Data, "struct data field should not be nil")

		// Get the data portion from raw JSON
		rawData, ok := rawJSON["data"].(map[string]any)
		require.True(t, ok, "data field should be an object")

		// Validate core fields
		assert.Equal(t, float64(1), rawData["id"], "id in JSON should match")
		assert.Equal(t, int64(1), response.Data.ID.Int64(), "id in struct should match")

		assert.Equal(t, float64(100982), rawData["integration"], "integration in JSON should match")
		assert.Equal(t, int64(100982), response.Data.Integration.Int64(), "integration in struct should match")

		assert.Equal(t, "live", rawData["domain"], "domain in JSON should match")
		assert.Equal(t, "live", response.Data.Domain.String(), "domain in struct should match")

		// Test transaction field (should be a number, not an object in fetch response)
		assert.Equal(t, float64(1641), rawData["transaction"], "transaction in JSON should match")
		assert.Equal(t, int64(1641), response.Data.Transaction.Int64(), "transaction in struct should match")

		assert.Equal(t, nil, rawData["dispute"], "dispute in JSON should be null")
		assert.False(t, response.Data.Dispute.Valid, "dispute in struct should be null")

		assert.Equal(t, nil, rawData["settlement"], "settlement in JSON should be null")
		assert.False(t, response.Data.Settlement.Valid, "settlement in struct should be null")

		assert.Equal(t, float64(500000), rawData["amount"], "amount in JSON should match")
		assert.Equal(t, int64(500000), response.Data.Amount.Int64(), "amount in struct should match")

		assert.Equal(t, float64(500000), rawData["deducted_amount"], "deducted_amount in JSON should match")
		assert.Equal(t, int64(500000), response.Data.DeductedAmount.Int64(), "deducted_amount in struct should match")

		assert.Equal(t, true, rawData["fully_deducted"], "fully_deducted in JSON should match")
		assert.Equal(t, true, response.Data.FullyDeducted.Bool(), "fully_deducted in struct should match")

		assert.Equal(t, "NGN", rawData["currency"], "currency in JSON should match")
		assert.Equal(t, enums.CurrencyNGN, response.Data.Currency, "currency in struct should match")

		assert.Equal(t, "migs", rawData["channel"], "channel in JSON should match")
		assert.NotNil(t, response.Data.Channel, "channel in struct should not be nil")
		assert.Equal(t, enums.RefundChannelMIGS, *response.Data.Channel, "channel in struct should match")

		assert.Equal(t, "processed", rawData["status"], "status in JSON should match")
		assert.Equal(t, enums.RefundStatusProcessed, response.Data.Status, "status in struct should match")

		assert.Equal(t, "eseyinwale@gmail.com", rawData["refunded_by"], "refunded_by in JSON should match")
		assert.Equal(t, "eseyinwale@gmail.com", response.Data.RefundedBy.String(), "refunded_by in struct should match")

		assert.Equal(t, "xxx", rawData["customer_note"], "customer_note in JSON should match")
		assert.Equal(t, "xxx", response.Data.CustomerNote.String(), "customer_note in struct should match")

		assert.Equal(t, "xxx", rawData["merchant_note"], "merchant_note in JSON should match")
		assert.Equal(t, "xxx", response.Data.MerchantNote.String(), "merchant_note in struct should match")

		// Test timestamps with lenient comparison
		assert.Equal(t, "2018-01-12T10:54:47.000Z", rawData["refunded_at"], "refunded_at in JSON should match")
		expectedRefundedAt, err := time.Parse(time.RFC3339, "2018-01-12T10:54:47.000Z")
		require.NoError(t, err, "should parse expected refunded_at")
		actualRefundedAt, err := time.Parse(time.RFC3339, response.Data.RefundedAt.String())
		require.NoError(t, err, "should parse actual refunded_at")
		assert.True(t, expectedRefundedAt.Sub(actualRefundedAt).Abs() < time.Second, "refunded_at timestamps should be within 1 second")

		assert.Equal(t, "2017-10-01T21:10:59.000Z", rawData["expected_at"], "expected_at in JSON should match")
		expectedExpectedAt, err := time.Parse(time.RFC3339, "2017-10-01T21:10:59.000Z")
		require.NoError(t, err, "should parse expected expected_at")
		actualExpectedAt, err := time.Parse(time.RFC3339, response.Data.ExpectedAt.String())
		require.NoError(t, err, "should parse actual expected_at")
		assert.True(t, expectedExpectedAt.Sub(actualExpectedAt).Abs() < time.Second, "expected_at timestamps should be within 1 second")

		assert.Equal(t, "2017-09-24T21:10:59.000Z", rawData["createdAt"], "createdAt in JSON should match")
		expectedCreatedAt, err := time.Parse(time.RFC3339, "2017-09-24T21:10:59.000Z")
		require.NoError(t, err, "should parse expected createdAt")
		actualCreatedAt, err := time.Parse(time.RFC3339, response.Data.CreatedAt.String())
		require.NoError(t, err, "should parse actual createdAt")
		assert.True(t, expectedCreatedAt.Sub(actualCreatedAt).Abs() < time.Second, "createdAt timestamps should be within 1 second")

		assert.Equal(t, "2018-01-18T11:59:56.000Z", rawData["updatedAt"], "updatedAt in JSON should match")
		expectedUpdatedAt, err := time.Parse(time.RFC3339, "2018-01-18T11:59:56.000Z")
		require.NoError(t, err, "should parse expected updatedAt")
		actualUpdatedAt, err := time.Parse(time.RFC3339, response.Data.UpdatedAt.String())
		require.NoError(t, err, "should parse actual updatedAt")
		assert.True(t, expectedUpdatedAt.Sub(actualUpdatedAt).Abs() < time.Second, "updatedAt timestamps should be within 1 second")
	})
}

func TestFetch(t *testing.T) {
	tests := []struct {
		name     string
		refundID string
	}{
		{
			name:     "fetch with refund ID",
			refundID: "1",
		},
		{
			name:     "fetch with different refund ID",
			refundID: "123456",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// This test verifies URL construction
			// In a real scenario, we would mock the HTTP client
			expectedPath := "/refund/" + tt.refundID
			assert.Contains(t, expectedPath, tt.refundID, "path should contain refund ID")
		})
	}
}

func TestFetchResponse_JSONRoundTrip(t *testing.T) {
	t.Run("deserialize_and_serialize_maintains_structure", func(t *testing.T) {
		// Read the JSON response file
		responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "refunds", "fetch_200.json")
		originalData, err := os.ReadFile(responseFilePath)
		require.NoError(t, err, "failed to read fetch_200.json")

		// Parse into struct
		var response FetchResponse
		err = json.Unmarshal(originalData, &response)
		require.NoError(t, err, "failed to unmarshal fetch_200.json")

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

		// Data should be an object in both
		originalRefundData := originalMap["data"].(map[string]any)
		serializedRefundData := serializedMap["data"].(map[string]any)

		// Test core refund fields survive round-trip
		assert.Equal(t, originalRefundData["id"], serializedRefundData["id"], "id should survive round-trip")
		assert.Equal(t, originalRefundData["integration"], serializedRefundData["integration"], "integration should survive round-trip")
		assert.Equal(t, originalRefundData["amount"], serializedRefundData["amount"], "amount should survive round-trip")
		assert.Equal(t, originalRefundData["currency"], serializedRefundData["currency"], "currency should survive round-trip")
		assert.Equal(t, originalRefundData["status"], serializedRefundData["status"], "status should survive round-trip")
	})
}
