package paymentrequests

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/huysamen/paystack-go/types"
	"github.com/huysamen/paystack-go/types/data"
)

func TestCreateResponse_JSONDeserialization(t *testing.T) {
	tests := []struct {
		name                string
		responseFile        string
		expectedStatus      bool
		expectedMessage     string
		expectedAmount      int64
		expectedRequestCode string
	}{
		{
			name:                "successful create response",
			responseFile:        "create_200.json",
			expectedStatus:      true,
			expectedMessage:     "Payment request created",
			expectedAmount:      42000,
			expectedRequestCode: "PRQ_1weqqsn2wwzgft8",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Read the response file
			responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "paymentrequests", tt.responseFile)
			responseData, err := os.ReadFile(responseFilePath)
			require.NoError(t, err, "failed to read response file: %s", responseFilePath)

			// Deserialize the JSON response
			var response CreateResponse
			err = json.Unmarshal(responseData, &response)
			require.NoError(t, err, "failed to unmarshal JSON response")

			// Verify the response structure
			assert.Equal(t, tt.expectedStatus, response.Status.Bool(), "status should match")
			assert.Equal(t, tt.expectedMessage, response.Message, "message should match")

			// Verify the data structure
			assert.NotNil(t, response.Data, "data should not be nil")
			assert.Equal(t, tt.expectedAmount, response.Data.Amount.Int64(), "amount should match")
			assert.Equal(t, tt.expectedRequestCode, response.Data.RequestCode.String(), "request code should match")
			assert.Greater(t, response.Data.ID.Int64(), int64(0), "ID should be greater than 0")
		})
	}
}

func TestCreateResponse_FieldByFieldValidation(t *testing.T) {
	t.Run("create_200_json_comprehensive_field_validation", func(t *testing.T) {
		// Read the exact JSON response file
		responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "paymentrequests", "create_200.json")
		responseData, err := os.ReadFile(responseFilePath)
		require.NoError(t, err, "failed to read create_200.json")

		// Parse into our struct
		var response CreateResponse
		err = json.Unmarshal(responseData, &response)
		require.NoError(t, err, "failed to unmarshal create_200.json")

		// Parse the raw JSON to compare exact values
		var rawJSON map[string]any
		err = json.Unmarshal(responseData, &rawJSON)
		require.NoError(t, err, "failed to parse raw JSON for comparison")

		// Field-by-field validation against the exact JSON values
		assert.Equal(t, true, rawJSON["status"], "status in JSON should be true")
		assert.Equal(t, true, response.Status.Bool(), "status in struct should be true")

		assert.Equal(t, "Payment request created", rawJSON["message"], "message in JSON should match")
		assert.Equal(t, "Payment request created", response.Message, "message in struct should match")

		// Verify data field exists and has correct structure
		assert.Contains(t, rawJSON, "data", "JSON should contain data field")
		assert.NotNil(t, response.Data, "struct data field should not be nil")

		// Get the data portion from raw JSON
		rawData, ok := rawJSON["data"].(map[string]any)
		require.True(t, ok, "data field should be an object")

		// Validate each data field
		assert.Equal(t, float64(3136406), rawData["id"], "id in JSON should match")
		assert.Equal(t, int64(3136406), response.Data.ID.Int64(), "id in struct should match")

		assert.Equal(t, "test", rawData["domain"], "domain in JSON should match")
		assert.Equal(t, "test", response.Data.Domain.String(), "domain in struct should match")

		assert.Equal(t, float64(42000), rawData["amount"], "amount in JSON should match")
		assert.Equal(t, int64(42000), response.Data.Amount.Int64(), "amount in struct should match")

		assert.Equal(t, "NGN", rawData["currency"], "currency in JSON should match")
		assert.Equal(t, "NGN", response.Data.Currency.String(), "currency in struct should match")

		assert.Equal(t, "2020-07-08T00:00:00.000Z", rawData["due_date"], "due_date in JSON should match")
		expectedDueDate, err := time.Parse(time.RFC3339, "2020-07-08T00:00:00.000Z")
		require.NoError(t, err, "should parse expected due date")
		actualDueDate, err := time.Parse(time.RFC3339, response.Data.DueDate.String())
		require.NoError(t, err, "should parse actual due date")
		assert.True(t, expectedDueDate.Equal(actualDueDate), "due date timestamps should be equal")

		assert.Equal(t, true, rawData["has_invoice"], "has_invoice in JSON should match")
		assert.Equal(t, true, response.Data.HasInvoice.Bool(), "has_invoice in struct should match")

		assert.Equal(t, float64(1), rawData["invoice_number"], "invoice_number in JSON should match")
		assert.Equal(t, int64(1), response.Data.InvoiceNumber.Int, "invoice_number in struct should match")

		assert.Equal(t, "a test invoice", rawData["description"], "description in JSON should match")
		assert.Equal(t, "a test invoice", response.Data.Description.String(), "description in struct should match")

		// Verify line items
		rawLineItems, ok := rawData["line_items"].([]any)
		require.True(t, ok, "line_items should be an array")
		assert.Len(t, rawLineItems, 2, "should have 2 line items in JSON")
		assert.Len(t, response.Data.LineItems, 2, "should have 2 line items in struct")

		// First line item
		rawLineItem1 := rawLineItems[0].(map[string]any)
		assert.Equal(t, "item 1", rawLineItem1["name"], "first line item name in JSON should match")
		assert.Equal(t, "item 1", response.Data.LineItems[0].Name.String(), "first line item name in struct should match")
		assert.Equal(t, float64(20000), rawLineItem1["amount"], "first line item amount in JSON should match")
		assert.Equal(t, int64(20000), response.Data.LineItems[0].Amount.Int64(), "first line item amount in struct should match")

		// Second line item
		rawLineItem2 := rawLineItems[1].(map[string]any)
		assert.Equal(t, "item 2", rawLineItem2["name"], "second line item name in JSON should match")
		assert.Equal(t, "item 2", response.Data.LineItems[1].Name.String(), "second line item name in struct should match")
		assert.Equal(t, float64(20000), rawLineItem2["amount"], "second line item amount in JSON should match")
		assert.Equal(t, int64(20000), response.Data.LineItems[1].Amount.Int64(), "second line item amount in struct should match")

		// Verify tax items
		rawTax, ok := rawData["tax"].([]any)
		require.True(t, ok, "tax should be an array")
		assert.Len(t, rawTax, 1, "should have 1 tax item in JSON")
		assert.Len(t, response.Data.Tax, 1, "should have 1 tax item in struct")

		rawTaxItem := rawTax[0].(map[string]any)
		assert.Equal(t, "VAT", rawTaxItem["name"], "tax item name in JSON should match")
		assert.Equal(t, "VAT", response.Data.Tax[0].Name.String(), "tax item name in struct should match")
		assert.Equal(t, float64(2000), rawTaxItem["amount"], "tax item amount in JSON should match")
		assert.Equal(t, int64(2000), response.Data.Tax[0].Amount.Int64(), "tax item amount in struct should match")

		assert.Equal(t, "PRQ_1weqqsn2wwzgft8", rawData["request_code"], "request_code in JSON should match")
		assert.Equal(t, "PRQ_1weqqsn2wwzgft8", response.Data.RequestCode.String(), "request_code in struct should match")

		assert.Equal(t, "pending", rawData["status"], "status in JSON should match")
		assert.Equal(t, "pending", response.Data.Status.String(), "status in struct should match")

		assert.Equal(t, false, rawData["paid"], "paid in JSON should match")
		assert.Equal(t, false, response.Data.Paid.Bool(), "paid in struct should match")

		assert.Equal(t, nil, rawData["metadata"], "metadata in JSON should be null")
		assert.False(t, response.Data.Metadata.Valid, "metadata in struct should be invalid")

		assert.Equal(t, "4286263136406", rawData["offline_reference"], "offline_reference in JSON should match")
		assert.Equal(t, "4286263136406", response.Data.OfflineReference.String(), "offline_reference in struct should match")

		assert.Equal(t, float64(25833615), rawData["customer"], "customer in JSON should match")
		// In create response, customer is just an integer ID
		customerInt := int64(rawData["customer"].(float64))
		assert.Equal(t, customerInt, response.Data.Customer.Int64(), "customer ID in struct should match")

		assert.Equal(t, "2020-06-29T16:07:33.073Z", rawData["created_at"], "created_at in JSON should match")
		expectedCreatedAt, err := time.Parse(time.RFC3339, "2020-06-29T16:07:33.073Z")
		require.NoError(t, err, "should parse expected created_at")
		actualCreatedAtStr := response.Data.CreatedAt.String()
		t.Logf("Expected timestamp: %s, Actual timestamp: %s", "2020-06-29T16:07:33.073Z", actualCreatedAtStr)
		actualCreatedAt, err := time.Parse(time.RFC3339, actualCreatedAtStr)
		require.NoError(t, err, "should parse actual created_at")
		// Allow for some millisecond precision differences
		timeDiff := expectedCreatedAt.Sub(actualCreatedAt)
		assert.True(t, timeDiff >= -time.Second && timeDiff <= time.Second, "created_at timestamps should be within 1 second of each other")

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

func TestCreateRequestBuilder(t *testing.T) {
	t.Run("builds request with required fields only", func(t *testing.T) {
		builder := NewCreateRequestBuilder()
		request := builder.
			Customer("CUS_xwaj0txjryg393b").
			Amount(10000).
			Build()

		assert.Equal(t, "CUS_xwaj0txjryg393b", request.Customer, "customer should match")
		assert.Equal(t, 10000, request.Amount, "amount should match")
		assert.Equal(t, "", request.Description, "description should be empty by default")
		assert.Equal(t, "", request.Currency, "currency should be empty by default")
		assert.Nil(t, request.SendNotification, "send_notification should be nil by default")
	})

	t.Run("builds request with all fields", func(t *testing.T) {
		metadata := types.NewMetadata(map[string]any{
			"custom_field": "custom_value",
		})

		lineItem := types.LineItem{
			Name:   data.NewString("Test Item"),
			Amount: data.NewInt(5000),
		}

		tax := types.Tax{
			Name:   data.NewString("VAT"),
			Amount: data.NewInt(500),
		}

		builder := NewCreateRequestBuilder()
		request := builder.
			Customer("CUS_xwaj0txjryg393b").
			Amount(10000).
			DueDate("2024-12-31").
			Description("Test Payment Request").
			Currency("NGN").
			SendNotification(true).
			Draft(false).
			HasInvoice(true).
			InvoiceNumber(123).
			SplitCode("SPL_test").
			Metadata(metadata).
			AddLineItem(lineItem).
			AddTax(tax).
			Build()

		assert.Equal(t, "CUS_xwaj0txjryg393b", request.Customer)
		assert.Equal(t, 10000, request.Amount)
		assert.Equal(t, "2024-12-31", request.DueDate)
		assert.Equal(t, "Test Payment Request", request.Description)
		assert.Equal(t, "NGN", request.Currency)
		assert.Equal(t, true, *request.SendNotification)
		assert.Equal(t, false, *request.Draft)
		assert.Equal(t, true, *request.HasInvoice)
		assert.Equal(t, 123, *request.InvoiceNumber)
		assert.Equal(t, "SPL_test", request.SplitCode)
		assert.Equal(t, metadata, request.Metadata)
		assert.Len(t, request.LineItems, 1)
		assert.Equal(t, "Test Item", request.LineItems[0].Name.String())
		assert.Equal(t, int64(5000), request.LineItems[0].Amount.Int64())
		assert.Len(t, request.Tax, 1)
		assert.Equal(t, "VAT", request.Tax[0].Name.String())
		assert.Equal(t, int64(500), request.Tax[0].Amount.Int64())
	})

	t.Run("builds request with multiple line items and taxes", func(t *testing.T) {
		lineItems := []types.LineItem{
			{Name: data.NewString("Item 1"), Amount: data.NewInt(2000)},
			{Name: data.NewString("Item 2"), Amount: data.NewInt(3000)},
		}

		builder := NewCreateRequestBuilder()
		request := builder.
			Customer("CUS_xwaj0txjryg393b").
			Amount(5500).
			LineItems(lineItems).
			AddTax(types.Tax{Name: data.NewString("VAT"), Amount: data.NewInt(500)}).
			Build()

		assert.Len(t, request.LineItems, 2)
		assert.Equal(t, "Item 1", request.LineItems[0].Name.String())
		assert.Equal(t, int64(2000), request.LineItems[0].Amount.Int64())
		assert.Equal(t, "Item 2", request.LineItems[1].Name.String())
		assert.Equal(t, int64(3000), request.LineItems[1].Amount.Int64())

		assert.Len(t, request.Tax, 1)
		assert.Equal(t, "VAT", request.Tax[0].Name.String())
		assert.Equal(t, int64(500), request.Tax[0].Amount.Int64())
	})
}

func TestCreateRequest_JSONSerialization(t *testing.T) {
	t.Run("serializes minimal request correctly", func(t *testing.T) {
		builder := NewCreateRequestBuilder()
		request := builder.
			Customer("CUS_xwaj0txjryg393b").
			Amount(10000).
			Build()

		jsonData, err := json.Marshal(request)
		require.NoError(t, err, "should marshal to JSON without error")

		var unmarshaled map[string]any
		err = json.Unmarshal(jsonData, &unmarshaled)
		require.NoError(t, err, "should unmarshal JSON without error")

		assert.Equal(t, "CUS_xwaj0txjryg393b", unmarshaled["customer"])
		assert.Equal(t, float64(10000), unmarshaled["amount"])
	})

	t.Run("serializes complete request correctly", func(t *testing.T) {
		metadata := types.NewMetadata(map[string]any{
			"order_id": "ORDER_123",
		})

		builder := NewCreateRequestBuilder()
		request := builder.
			Customer("CUS_xwaj0txjryg393b").
			Amount(15000).
			Description("Complete test request").
			Currency("USD").
			DueDate("2024-12-31").
			SendNotification(true).
			Draft(false).
			HasInvoice(true).
			InvoiceNumber(456).
			SplitCode("SPL_complete").
			Metadata(metadata).
			AddLineItem(types.LineItem{Name: data.NewString("Service"), Amount: data.NewInt(10000)}).
			AddTax(types.Tax{Name: data.NewString("Sales Tax"), Amount: data.NewInt(1000)}).
			Build()

		jsonData, err := json.Marshal(request)
		require.NoError(t, err, "should marshal to JSON without error")

		var unmarshaled map[string]any
		err = json.Unmarshal(jsonData, &unmarshaled)
		require.NoError(t, err, "should unmarshal JSON without error")

		assert.Equal(t, "CUS_xwaj0txjryg393b", unmarshaled["customer"])
		assert.Equal(t, float64(15000), unmarshaled["amount"])
		assert.Equal(t, "Complete test request", unmarshaled["description"])
		assert.Equal(t, "USD", unmarshaled["currency"])
		assert.Equal(t, "2024-12-31", unmarshaled["due_date"])
		assert.Equal(t, true, unmarshaled["send_notification"])
		assert.Equal(t, false, unmarshaled["draft"])
		assert.Equal(t, true, unmarshaled["has_invoice"])
		assert.Equal(t, float64(456), unmarshaled["invoice_number"])
		assert.Equal(t, "SPL_complete", unmarshaled["split_code"])

		// Verify metadata
		metadataObj, ok := unmarshaled["metadata"].(map[string]any)
		require.True(t, ok, "metadata should be an object")
		assert.Equal(t, "ORDER_123", metadataObj["order_id"])

		// Verify line items
		lineItems, ok := unmarshaled["line_items"].([]any)
		require.True(t, ok, "line_items should be an array")
		assert.Len(t, lineItems, 1)
		lineItem := lineItems[0].(map[string]any)
		assert.Equal(t, "Service", lineItem["name"])
		assert.Equal(t, float64(10000), lineItem["amount"])

		// Verify tax items
		taxes, ok := unmarshaled["tax"].([]any)
		require.True(t, ok, "tax should be an array")
		assert.Len(t, taxes, 1)
		tax := taxes[0].(map[string]any)
		assert.Equal(t, "Sales Tax", tax["name"])
		assert.Equal(t, float64(1000), tax["amount"])
	})

	t.Run("omits empty optional fields", func(t *testing.T) {
		builder := NewCreateRequestBuilder()
		request := builder.
			Customer("CUS_xwaj0txjryg393b").
			Amount(5000).
			Build()

		jsonData, err := json.Marshal(request)
		require.NoError(t, err, "should marshal to JSON without error")

		var unmarshaled map[string]any
		err = json.Unmarshal(jsonData, &unmarshaled)
		require.NoError(t, err, "should unmarshal JSON without error")

		// Required fields should be present
		assert.Contains(t, unmarshaled, "customer")
		assert.Contains(t, unmarshaled, "amount")

		// Optional fields should be omitted when empty
		if description, exists := unmarshaled["description"]; exists {
			assert.Equal(t, "", description)
		}
		if currency, exists := unmarshaled["currency"]; exists {
			assert.Equal(t, "", currency)
		}
		assert.NotContains(t, unmarshaled, "send_notification")
		assert.NotContains(t, unmarshaled, "draft")
		assert.NotContains(t, unmarshaled, "has_invoice")
		assert.NotContains(t, unmarshaled, "invoice_number")
	})
}
