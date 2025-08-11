package paymentrequests

import (
	"encoding/json"
	"net/url"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListResponse_JSONDeserialization(t *testing.T) {
	tests := []struct {
		name               string
		responseFile       string
		expectedStatus     bool
		expectedMessage    string
		expectedTotalCount int64
		minPaymentRequests int
	}{
		{
			name:               "successful list response",
			responseFile:       "list_200.json",
			expectedStatus:     true,
			expectedMessage:    "Payment requests retrieved",
			expectedTotalCount: 1,
			minPaymentRequests: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Read the response file
			responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "paymentrequests", tt.responseFile)
			responseData, err := os.ReadFile(responseFilePath)
			require.NoError(t, err, "failed to read response file: %s", responseFilePath)

			// Deserialize the JSON response
			var response ListResponse
			err = json.Unmarshal(responseData, &response)
			require.NoError(t, err, "failed to unmarshal JSON response")

			// Verify the response structure
			assert.Equal(t, tt.expectedStatus, response.Status.Bool(), "status should match")
			assert.Equal(t, tt.expectedMessage, response.Message, "message should match")

			// Verify the data structure
			assert.NotNil(t, response.Data, "data should not be nil")
			assert.GreaterOrEqual(t, len(response.Data), tt.minPaymentRequests, "should have at least expected payment requests")

			// Verify pagination data
			if len(response.Data) > 0 {
				paymentRequest := response.Data[0]
				assert.Greater(t, paymentRequest.ID.Int64(), int64(0), "payment request ID should be greater than 0")
				assert.NotEmpty(t, paymentRequest.RequestCode.String(), "request code should not be empty")
				assert.Greater(t, paymentRequest.Amount.Int64(), int64(0), "amount should be greater than 0")
			}

			// Verify meta information
			assert.NotNil(t, response.Meta, "meta should not be nil")
			assert.Equal(t, tt.expectedTotalCount, response.Meta.Total.Int, "total count should match")
		})
	}
}

func TestListResponse_FieldByFieldValidation(t *testing.T) {
	t.Run("list_200_json_comprehensive_field_validation", func(t *testing.T) {
		// Read the exact JSON response file
		responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "paymentrequests", "list_200.json")
		responseData, err := os.ReadFile(responseFilePath)
		require.NoError(t, err, "failed to read list_200.json")

		// Parse into our struct
		var response ListResponse
		err = json.Unmarshal(responseData, &response)
		require.NoError(t, err, "failed to unmarshal list_200.json")

		// Parse the raw JSON to compare exact values
		var rawJSON map[string]any
		err = json.Unmarshal(responseData, &rawJSON)
		require.NoError(t, err, "failed to parse raw JSON for comparison")

		// Field-by-field validation against the exact JSON values
		assert.Equal(t, true, rawJSON["status"], "status in JSON should be true")
		assert.Equal(t, true, response.Status.Bool(), "status in struct should be true")

		assert.Equal(t, "Payment requests retrieved", rawJSON["message"], "message in JSON should match")
		assert.Equal(t, "Payment requests retrieved", response.Message, "message in struct should match")

		// Verify data field exists and has correct structure
		assert.Contains(t, rawJSON, "data", "JSON should contain data field")
		assert.NotNil(t, response.Data, "struct data field should not be nil")

		// Get the data portion from raw JSON
		rawData, ok := rawJSON["data"].([]any)
		require.True(t, ok, "data field should be an array")
		require.Len(t, rawData, 1, "should have 1 payment request in response")
		assert.Len(t, response.Data, 1, "should have 1 payment request in struct")

		// Validate the first payment request
		rawPaymentRequest := rawData[0].(map[string]any)
		paymentRequest := response.Data[0]

		assert.Equal(t, float64(3136406), rawPaymentRequest["id"], "id in JSON should match")
		assert.Equal(t, int64(3136406), paymentRequest.ID.Int64(), "id in struct should match")

		assert.Equal(t, "test", rawPaymentRequest["domain"], "domain in JSON should match")
		assert.Equal(t, "test", paymentRequest.Domain.String(), "domain in struct should match")

		assert.Equal(t, float64(42000), rawPaymentRequest["amount"], "amount in JSON should match")
		assert.Equal(t, int64(42000), paymentRequest.Amount.Int64(), "amount in struct should match")

		assert.Equal(t, "NGN", rawPaymentRequest["currency"], "currency in JSON should match")
		assert.Equal(t, "NGN", paymentRequest.Currency.String(), "currency in struct should match")

		assert.Equal(t, "2020-07-08T00:00:00.000Z", rawPaymentRequest["due_date"], "due_date in JSON should match")
		expectedDueDate, err := time.Parse(time.RFC3339, "2020-07-08T00:00:00.000Z")
		require.NoError(t, err, "should parse expected due date")
		actualDueDate, err := time.Parse(time.RFC3339, paymentRequest.DueDate.String())
		require.NoError(t, err, "should parse actual due date")
		assert.True(t, expectedDueDate.Equal(actualDueDate), "due date timestamps should be equal")

		assert.Equal(t, true, rawPaymentRequest["has_invoice"], "has_invoice in JSON should match")
		assert.Equal(t, true, paymentRequest.HasInvoice.Bool(), "has_invoice in struct should match")

		assert.Equal(t, float64(1), rawPaymentRequest["invoice_number"], "invoice_number in JSON should match")
		assert.Equal(t, int64(1), paymentRequest.InvoiceNumber.Int, "invoice_number in struct should match")

		assert.Equal(t, "a test invoice", rawPaymentRequest["description"], "description in JSON should match")
		assert.Equal(t, "a test invoice", paymentRequest.Description.String(), "description in struct should match")

		assert.Equal(t, "PRQ_1weqqsn2wwzgft8", rawPaymentRequest["request_code"], "request_code in JSON should match")
		assert.Equal(t, "PRQ_1weqqsn2wwzgft8", paymentRequest.RequestCode.String(), "request_code in struct should match")

		assert.Equal(t, "pending", rawPaymentRequest["status"], "status in JSON should match")
		assert.Equal(t, "pending", paymentRequest.Status.String(), "status in struct should match")

		assert.Equal(t, false, rawPaymentRequest["paid"], "paid in JSON should match")
		assert.Equal(t, false, paymentRequest.Paid.Bool(), "paid in struct should match")

		assert.Equal(t, "2020-06-29T16:07:33.000Z", rawPaymentRequest["created_at"], "created_at in JSON should match")
		expectedCreatedAt, err := time.Parse(time.RFC3339, "2020-06-29T16:07:33.000Z")
		require.NoError(t, err, "should parse expected created_at")
		actualCreatedAt, err := time.Parse(time.RFC3339, paymentRequest.CreatedAt.String())
		require.NoError(t, err, "should parse actual created_at")
		assert.True(t, expectedCreatedAt.Equal(actualCreatedAt), "created_at timestamps should be equal")

		// Verify meta field
		assert.Contains(t, rawJSON, "meta", "JSON should contain meta field")
		assert.NotNil(t, response.Meta, "struct meta field should not be nil")

		rawMeta := rawJSON["meta"].(map[string]any)
		assert.Equal(t, float64(1), rawMeta["total"], "total in JSON should match")
		assert.Equal(t, int64(1), response.Meta.Total.Int, "total in struct should match")

		assert.Equal(t, float64(50), rawMeta["perPage"], "perPage in JSON should match")
		assert.Equal(t, int64(50), response.Meta.PerPage.Int64(), "perPage in struct should match")

		assert.Equal(t, float64(1), rawMeta["page"], "page in JSON should match")
		assert.Equal(t, int64(1), response.Meta.Page.Int, "page in struct should match")

		assert.Equal(t, float64(1), rawMeta["pageCount"], "pageCount in JSON should match")
		assert.Equal(t, int64(1), response.Meta.PageCount.Int, "pageCount in struct should match")

		// Verify complete JSON structure matches our struct
		reconstituted, err := json.Marshal(response)
		require.NoError(t, err, "should be able to marshal struct back to JSON")

		var reconstitutedMap map[string]any
		err = json.Unmarshal(reconstituted, &reconstitutedMap)
		require.NoError(t, err, "should be able to parse reconstituted JSON")

		// Core fields should match
		assert.Equal(t, rawJSON["status"], reconstitutedMap["status"], "status should survive round-trip")
		assert.Equal(t, rawJSON["message"], reconstitutedMap["message"], "message should survive round-trip")
	})
}

func TestListRequestBuilder(t *testing.T) {
	t.Run("builds request with no filters", func(t *testing.T) {
		builder := NewListRequestBuilder()
		request := builder.Build()

		// Default request should be empty
		assert.Equal(t, 0, request.PerPage, "per_page should be default")
		assert.Equal(t, 0, request.Page, "page should be default")
		assert.Equal(t, "", request.Customer, "customer should be empty by default")
		assert.Equal(t, "", request.Status, "status should be empty by default")
		assert.Equal(t, "", request.Currency, "currency should be empty by default")
	})

	t.Run("builds request with pagination", func(t *testing.T) {
		builder := NewListRequestBuilder()
		request := builder.
			Page(2).
			PerPage(25).
			Build()

		assert.Equal(t, 2, request.Page)
		assert.Equal(t, 25, request.PerPage)
	})

	t.Run("builds request with filters", func(t *testing.T) {
		builder := NewListRequestBuilder()
		request := builder.
			Customer("CUS_xwaj0txjryg393b").
			Status("pending").
			Currency("NGN").
			From("2024-01-01").
			To("2024-12-31").
			Build()

		assert.Equal(t, "CUS_xwaj0txjryg393b", request.Customer)
		assert.Equal(t, "pending", request.Status)
		assert.Equal(t, "NGN", request.Currency)
		assert.Equal(t, "2024-01-01", request.From)
		assert.Equal(t, "2024-12-31", request.To)
	})

	t.Run("builds request with complete filters", func(t *testing.T) {
		builder := NewListRequestBuilder()
		request := builder.
			Page(3).
			PerPage(100).
			Customer("CUS_test123").
			Status("successful").
			Currency("USD").
			From("2023-01-01").
			To("2023-12-31").
			Build()

		assert.Equal(t, 3, request.Page)
		assert.Equal(t, 100, request.PerPage)
		assert.Equal(t, "CUS_test123", request.Customer)
		assert.Equal(t, "successful", request.Status)
		assert.Equal(t, "USD", request.Currency)
		assert.Equal(t, "2023-01-01", request.From)
		assert.Equal(t, "2023-12-31", request.To)
	})
}

func TestListRequest_URLQuery(t *testing.T) {
	t.Run("converts to URL query correctly", func(t *testing.T) {
		builder := NewListRequestBuilder()
		request := builder.
			Page(2).
			PerPage(50).
			Customer("CUS_test").
			Status("pending").
			Currency("NGN").
			From("2024-01-01").
			To("2024-12-31").
			Build()

		query := request.toQuery()

		// Parse the query string to verify parameters
		values, err := url.ParseQuery(query)
		require.NoError(t, err, "should parse query string without error")

		assert.Equal(t, "2", values.Get("page"))
		assert.Equal(t, "50", values.Get("perPage"))
		assert.Equal(t, "CUS_test", values.Get("customer"))
		assert.Equal(t, "pending", values.Get("status"))
		assert.Equal(t, "NGN", values.Get("currency"))
		assert.Equal(t, "2024-01-01", values.Get("from"))
		assert.Equal(t, "2024-12-31", values.Get("to"))
	})

	t.Run("omits empty values", func(t *testing.T) {
		builder := NewListRequestBuilder()
		request := builder.
			Page(1).
			Status("pending").
			Build()

		query := request.toQuery()
		values, err := url.ParseQuery(query)
		require.NoError(t, err, "should parse query string without error")

		assert.Equal(t, "1", values.Get("page"))
		assert.Equal(t, "pending", values.Get("status"))
		assert.Equal(t, "", values.Get("customer")) // Empty values should not be set
		assert.Equal(t, "", values.Get("currency"))
		assert.Equal(t, "", values.Get("from"))
		assert.Equal(t, "", values.Get("to"))
	})
}
