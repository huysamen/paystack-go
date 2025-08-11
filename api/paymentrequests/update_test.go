package paymentrequests

import (
	"encoding/json"
	"testing"

	"github.com/huysamen/paystack-go/types"
	"github.com/huysamen/paystack-go/types/data"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUpdateRequestBuilder(t *testing.T) {
	t.Run("builds request with minimal fields", func(t *testing.T) {
		builder := NewUpdateRequestBuilder()
		request := builder.
			Description("Updated description").
			Build()

		assert.Equal(t, "Updated description", request.Description)
		assert.Equal(t, "", request.Currency, "currency should be empty by default")
		assert.Nil(t, request.SendNotification, "send_notification should be nil by default")
	})

	t.Run("builds request with all fields", func(t *testing.T) {
		metadata := types.NewMetadata(map[string]any{
			"updated_field": "updated_value",
		})

		lineItem := types.LineItem{
			Name:   data.NewString("Updated Item"),
			Amount: data.NewInt(7500),
		}

		tax := types.Tax{
			Name:   data.NewString("Updated Tax"),
			Amount: data.NewInt(750),
		}

		builder := NewUpdateRequestBuilder()
		request := builder.
			DueDate("2024-12-25").
			Description("Updated Payment Request").
			Currency("USD").
			SendNotification(false).
			Draft(true).
			Metadata(metadata).
			AddLineItem(lineItem).
			AddTax(tax).
			Build()

		assert.Equal(t, "2024-12-25", request.DueDate)
		assert.Equal(t, "Updated Payment Request", request.Description)
		assert.Equal(t, "USD", request.Currency)
		assert.Equal(t, false, *request.SendNotification)
		assert.Equal(t, true, *request.Draft)
		assert.Equal(t, metadata, request.Metadata)
		assert.Len(t, request.LineItems, 1)
		assert.Equal(t, "Updated Item", request.LineItems[0].Name.String())
		assert.Equal(t, int64(7500), request.LineItems[0].Amount.Int64())
		assert.Len(t, request.Tax, 1)
		assert.Equal(t, "Updated Tax", request.Tax[0].Name.String())
		assert.Equal(t, int64(750), request.Tax[0].Amount.Int64())
	})

	t.Run("builds request with multiple line items and taxes", func(t *testing.T) {
		lineItems := []types.LineItem{
			{Name: data.NewString("Updated Item 1"), Amount: data.NewInt(3000)},
			{Name: data.NewString("Updated Item 2"), Amount: data.NewInt(4000)},
		}

		builder := NewUpdateRequestBuilder()
		request := builder.
			Description("Multi-item update").
			LineItems(lineItems).
			AddTax(types.Tax{Name: data.NewString("VAT"), Amount: data.NewInt(700)}).
			AddTax(types.Tax{Name: data.NewString("Service Tax"), Amount: data.NewInt(350)}).
			Build()

		assert.Equal(t, "Multi-item update", request.Description)
		assert.Len(t, request.LineItems, 2)
		assert.Equal(t, "Updated Item 1", request.LineItems[0].Name.String())
		assert.Equal(t, int64(3000), request.LineItems[0].Amount.Int64())
		assert.Equal(t, "Updated Item 2", request.LineItems[1].Name.String())
		assert.Equal(t, int64(4000), request.LineItems[1].Amount.Int64())

		assert.Len(t, request.Tax, 2)
		assert.Equal(t, "VAT", request.Tax[0].Name.String())
		assert.Equal(t, int64(700), request.Tax[0].Amount.Int64())
		assert.Equal(t, "Service Tax", request.Tax[1].Name.String())
		assert.Equal(t, int64(350), request.Tax[1].Amount.Int64())
	})
}

func TestUpdateRequest_JSONSerialization(t *testing.T) {
	t.Run("serializes minimal request correctly", func(t *testing.T) {
		builder := NewUpdateRequestBuilder()
		request := builder.
			Description("Updated description").
			Build()

		jsonData, err := json.Marshal(request)
		require.NoError(t, err, "should marshal to JSON without error")

		var unmarshaled map[string]any
		err = json.Unmarshal(jsonData, &unmarshaled)
		require.NoError(t, err, "should unmarshal JSON without error")

		assert.Equal(t, "Updated description", unmarshaled["description"])
	})

	t.Run("serializes complete request correctly", func(t *testing.T) {
		metadata := types.NewMetadata(map[string]any{
			"updated_order_id": "ORDER_456",
		})

		builder := NewUpdateRequestBuilder()
		request := builder.
			Description("Complete updated request").
			Currency("EUR").
			DueDate("2024-11-30").
			SendNotification(true).
			Draft(false).
			Metadata(metadata).
			AddLineItem(types.LineItem{Name: data.NewString("Updated Service"), Amount: data.NewInt(12000)}).
			AddTax(types.Tax{Name: data.NewString("Updated Tax"), Amount: data.NewInt(1200)}).
			Build()

		jsonData, err := json.Marshal(request)
		require.NoError(t, err, "should marshal to JSON without error")

		var unmarshaled map[string]any
		err = json.Unmarshal(jsonData, &unmarshaled)
		require.NoError(t, err, "should unmarshal JSON without error")

		assert.Equal(t, "Complete updated request", unmarshaled["description"])
		assert.Equal(t, "EUR", unmarshaled["currency"])
		assert.Equal(t, "2024-11-30", unmarshaled["due_date"])
		assert.Equal(t, true, unmarshaled["send_notification"])
		assert.Equal(t, false, unmarshaled["draft"])

		// Verify metadata
		metadataObj, ok := unmarshaled["metadata"].(map[string]any)
		require.True(t, ok, "metadata should be an object")
		assert.Equal(t, "ORDER_456", metadataObj["updated_order_id"])

		// Verify line items
		lineItems, ok := unmarshaled["line_items"].([]any)
		require.True(t, ok, "line_items should be an array")
		assert.Len(t, lineItems, 1)
		lineItem := lineItems[0].(map[string]any)
		assert.Equal(t, "Updated Service", lineItem["name"])
		assert.Equal(t, float64(12000), lineItem["amount"])

		// Verify tax items
		taxes, ok := unmarshaled["tax"].([]any)
		require.True(t, ok, "tax should be an array")
		assert.Len(t, taxes, 1)
		tax := taxes[0].(map[string]any)
		assert.Equal(t, "Updated Tax", tax["name"])
		assert.Equal(t, float64(1200), tax["amount"])
	})

	t.Run("omits empty optional fields", func(t *testing.T) {
		builder := NewUpdateRequestBuilder()
		request := builder.
			Description("Basic update").
			Build()

		jsonData, err := json.Marshal(request)
		require.NoError(t, err, "should marshal to JSON without error")

		var unmarshaled map[string]any
		err = json.Unmarshal(jsonData, &unmarshaled)
		require.NoError(t, err, "should unmarshal JSON without error")

		// Required field should be present
		assert.Contains(t, unmarshaled, "description")

		// Optional fields should be omitted when empty
		if currency, exists := unmarshaled["currency"]; exists {
			assert.Equal(t, "", currency, "currency should be empty string if present")
		}
		assert.NotContains(t, unmarshaled, "send_notification")
		assert.NotContains(t, unmarshaled, "draft")
	})
}
