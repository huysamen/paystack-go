package refunds

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
			name:            "successful create response",
			responseFile:    "create_200.json",
			expectedStatus:  true,
			expectedMessage: "Refund has been queued for processing",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Read the JSON response file
			responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "refunds", tt.responseFile)
			responseData, err := os.ReadFile(responseFilePath)
			require.NoError(t, err, "failed to read response file")

			// Deserialize the JSON response
			var response CreateResponse
			err = json.Unmarshal(responseData, &response)
			require.NoError(t, err, "failed to unmarshal JSON response")

			// Verify the response structure
			assert.Equal(t, tt.expectedStatus, response.Status.Bool(), "status should match")
			assert.Equal(t, tt.expectedMessage, response.Message, "message should match")
			assert.NotNil(t, response.Data, "data should not be nil")
		})
	}
}

func TestCreateResponse_FieldByFieldValidation(t *testing.T) {
	t.Run("create_200_json_comprehensive_field_validation", func(t *testing.T) {
		// Read the exact JSON response file
		responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "refunds", "create_200.json")
		responseData, err := os.ReadFile(responseFilePath)
		require.NoError(t, err, "failed to read create_200.json")

		// Parse raw JSON to compare field by field
		var rawJSON map[string]any
		err = json.Unmarshal(responseData, &rawJSON)
		require.NoError(t, err, "failed to parse raw JSON")

		// Parse into our struct
		var response CreateResponse
		err = json.Unmarshal(responseData, &response)
		require.NoError(t, err, "failed to unmarshal create_200.json")

		// Verify struct data is properly populated
		assert.NotNil(t, response.Data, "struct data field should not be nil")

		// Get the data portion from raw JSON
		rawData, ok := rawJSON["data"].(map[string]any)
		require.True(t, ok, "data field should be an object")

		// Validate core fields
		assert.Equal(t, float64(3018284), rawData["id"], "id in JSON should match")

		assert.Equal(t, float64(412829), rawData["integration"], "integration in JSON should match")

		assert.Equal(t, "live", rawData["domain"], "domain in JSON should match")

		assert.Equal(t, float64(10000), rawData["amount"], "amount in JSON should match")
		assert.Equal(t, int64(10000), response.Data.Amount.Int64(), "amount in struct should match")

		assert.Equal(t, float64(0), rawData["deducted_amount"], "deducted_amount in JSON should match")

		assert.Equal(t, "NGN", rawData["currency"], "currency in JSON should match")
		assert.Equal(t, "NGN", response.Data.Currency.String(), "currency in struct should match")

		assert.Equal(t, nil, rawData["channel"], "channel in JSON should be null")

		assert.Equal(t, false, rawData["fully_deducted"], "fully_deducted in JSON should match")

		assert.Equal(t, "pending", rawData["status"], "status in JSON should match")

		assert.Equal(t, "test@me.com", rawData["refunded_by"], "refunded_by in JSON should match")
		assert.Equal(t, "test@me.com", response.Data.RefundedBy.String(), "refunded_by in struct should match")

		assert.Equal(t, "Refund for transaction T685312322670591 by test@me.com", rawData["merchant_note"], "merchant_note in JSON should match")

		assert.Equal(t, "Refund for transaction T685312322670591", rawData["customer_note"], "customer_note in JSON should match")

		// Test nested transaction object
		rawTransaction, ok := rawData["transaction"].(map[string]any)
		require.True(t, ok, "transaction should be an object")
		assert.NotNil(t, response.Data.Transaction, "transaction in struct should not be nil")

		assert.Equal(t, float64(1004723697), rawTransaction["id"], "transaction.id in JSON should match")
		assert.Equal(t, "T685312322670591", rawTransaction["reference"], "transaction.reference in JSON should match")
		assert.Equal(t, float64(10000), rawTransaction["amount"], "transaction.amount in JSON should match")
		assert.Equal(t, "apple_pay", rawTransaction["channel"], "transaction.channel in JSON should match")
		assert.Equal(t, "NGN", rawTransaction["currency"], "transaction.currency in JSON should match")

		// Test timestamps with lenient comparison
		assert.Equal(t, "2021-12-16T09:21:17.016Z", rawData["expected_at"], "expected_at in JSON should match")

		assert.Equal(t, "2021-12-07T09:21:17.122Z", rawData["createdAt"], "createdAt in JSON should match")

		// Test that refunded_at is not in the response (it's a pending refund)
		_, hasRefundedAt := rawData["refunded_at"]
		assert.False(t, hasRefundedAt, "refunded_at should not be present in create response for pending refund")
	})
}

func TestCreateRequestBuilder(t *testing.T) {
	tests := []struct {
		name         string
		setupBuilder func() *CreateRequestBuilder
		checkRequest func(t *testing.T, req *createRequest)
	}{
		{
			name: "builds request with only transaction",
			setupBuilder: func() *CreateRequestBuilder {
				return NewCreateRequestBuilder("T123456789")
			},
			checkRequest: func(t *testing.T, req *createRequest) {
				assert.Equal(t, "T123456789", req.Transaction, "transaction should be set")
				assert.Nil(t, req.Amount, "amount should be nil")
				assert.Nil(t, req.Currency, "currency should be nil")
				assert.Nil(t, req.CustomerNote, "customer_note should be nil")
				assert.Nil(t, req.MerchantNote, "merchant_note should be nil")
			},
		},
		{
			name: "builds request with amount",
			setupBuilder: func() *CreateRequestBuilder {
				return NewCreateRequestBuilder("T123456789").
					Amount(50000)
			},
			checkRequest: func(t *testing.T, req *createRequest) {
				assert.Equal(t, "T123456789", req.Transaction, "transaction should be set")
				require.NotNil(t, req.Amount, "amount should not be nil")
				assert.Equal(t, 50000, *req.Amount, "amount should match")
			},
		},
		{
			name: "builds request with currency",
			setupBuilder: func() *CreateRequestBuilder {
				return NewCreateRequestBuilder("T123456789").
					Currency("USD")
			},
			checkRequest: func(t *testing.T, req *createRequest) {
				assert.Equal(t, "T123456789", req.Transaction, "transaction should be set")
				require.NotNil(t, req.Currency, "currency should not be nil")
				assert.Equal(t, "USD", *req.Currency, "currency should match")
			},
		},
		{
			name: "builds request with notes",
			setupBuilder: func() *CreateRequestBuilder {
				return NewCreateRequestBuilder("T123456789").
					CustomerNote("Customer requested refund").
					MerchantNote("Approved by manager")
			},
			checkRequest: func(t *testing.T, req *createRequest) {
				assert.Equal(t, "T123456789", req.Transaction, "transaction should be set")
				require.NotNil(t, req.CustomerNote, "customer_note should not be nil")
				assert.Equal(t, "Customer requested refund", *req.CustomerNote, "customer_note should match")
				require.NotNil(t, req.MerchantNote, "merchant_note should not be nil")
				assert.Equal(t, "Approved by manager", *req.MerchantNote, "merchant_note should match")
			},
		},
		{
			name: "builds request with all fields",
			setupBuilder: func() *CreateRequestBuilder {
				return NewCreateRequestBuilder("T123456789").
					Amount(25000).
					Currency("NGN").
					CustomerNote("Refund requested").
					MerchantNote("Processing refund")
			},
			checkRequest: func(t *testing.T, req *createRequest) {
				assert.Equal(t, "T123456789", req.Transaction, "transaction should be set")
				require.NotNil(t, req.Amount, "amount should not be nil")
				assert.Equal(t, 25000, *req.Amount, "amount should match")
				require.NotNil(t, req.Currency, "currency should not be nil")
				assert.Equal(t, "NGN", *req.Currency, "currency should match")
				require.NotNil(t, req.CustomerNote, "customer_note should not be nil")
				assert.Equal(t, "Refund requested", *req.CustomerNote, "customer_note should match")
				require.NotNil(t, req.MerchantNote, "merchant_note should not be nil")
				assert.Equal(t, "Processing refund", *req.MerchantNote, "merchant_note should match")
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

func TestCreate(t *testing.T) {
	tests := []struct {
		name            string
		transaction     string
		expectedRequest createRequest
	}{
		{
			name:        "create with minimal request",
			transaction: "T123456789",
			expectedRequest: createRequest{
				Transaction: "T123456789",
			},
		},
		{
			name:        "create with full request",
			transaction: "T987654321",
			expectedRequest: createRequest{
				Transaction: "T987654321",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// This test verifies request structure and builder pattern
			// In a real scenario, we would mock the HTTP client
			builder := NewCreateRequestBuilder(tt.transaction)
			req := builder.Build()

			assert.Equal(t, tt.expectedRequest.Transaction, req.Transaction, "transaction should match")
		})
	}
}

func TestCreateResponse_JSONRoundTrip(t *testing.T) {
	t.Run("deserialize_and_serialize_maintains_structure", func(t *testing.T) {
		// Read the JSON response file
		responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "refunds", "create_200.json")
		originalData, err := os.ReadFile(responseFilePath)
		require.NoError(t, err, "failed to read create_200.json")

		// Parse into struct
		var response CreateResponse
		err = json.Unmarshal(originalData, &response)
		require.NoError(t, err, "failed to unmarshal create_200.json")

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

		// Test core refund fields survive round-trip (skip nested transaction which has many fields)
		assert.Equal(t, originalRefundData["amount"], serializedRefundData["amount"], "amount should survive round-trip")
		assert.Equal(t, originalRefundData["currency"], serializedRefundData["currency"], "currency should survive round-trip")
		assert.Equal(t, originalRefundData["status"], serializedRefundData["status"], "status should survive round-trip")
		assert.Equal(t, originalRefundData["refunded_by"], serializedRefundData["refunded_by"], "refunded_by should survive round-trip")
		assert.Equal(t, originalRefundData["domain"], serializedRefundData["domain"], "domain should survive round-trip")
	})
}
