package plans

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/huysamen/paystack-go/enums"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUpdateResponse_JSONDeserialization(t *testing.T) {
	tests := []struct {
		name            string
		responseFile    string
		expectedStatus  bool
		expectedMessage string
	}{
		{
			name:            "successful update response",
			responseFile:    "update_200.json",
			expectedStatus:  true,
			expectedMessage: "Plan updated. 1 subscription(s) affected",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Read the JSON response file
			responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "plans", tt.responseFile)
			responseData, err := os.ReadFile(responseFilePath)
			require.NoError(t, err, "failed to read response file")

			// Deserialize the JSON response
			var response UpdateResponse
			err = json.Unmarshal(responseData, &response)
			require.NoError(t, err, "failed to unmarshal JSON response")

			// Verify the response structure
			assert.Equal(t, tt.expectedStatus, response.Status.Bool(), "status should match")
			assert.Equal(t, tt.expectedMessage, response.Message, "message should match")

			// Update response doesn't return plan data, just status and message
		})
	}
}

func TestUpdateRequestBuilder(t *testing.T) {
	tests := []struct {
		name                                string
		setupBuilder                        func() *UpdateRequestBuilder
		expectedName                        string
		expectedAmount                      int
		expectedInterval                    enums.Interval
		expectedDescription                 *string
		expectedSendInvoices                *bool
		expectedSendSMS                     *bool
		expectedCurrency                    enums.Currency
		expectedInvoiceLimit                *int
		expectedUpdateExistingSubscriptions *bool
	}{
		{
			name: "builds request with required fields only",
			setupBuilder: func() *UpdateRequestBuilder {
				return NewUpdateRequestBuilder("Updated Plan", 15000, enums.IntervalMonthly)
			},
			expectedName:     "Updated Plan",
			expectedAmount:   15000,
			expectedInterval: enums.IntervalMonthly,
		},
		{
			name: "builds request with all fields",
			setupBuilder: func() *UpdateRequestBuilder {
				limit := 3
				return NewUpdateRequestBuilder("Premium Updated", 75000, enums.IntervalWeekly).
					Description("An updated premium subscription plan").
					SendInvoices(false).
					SendSMS(true).
					Currency(enums.CurrencyUSD).
					InvoiceLimit(limit).
					UpdateExistingSubscriptions(true)
			},
			expectedName:                        "Premium Updated",
			expectedAmount:                      75000,
			expectedInterval:                    enums.IntervalWeekly,
			expectedDescription:                 stringPtr("An updated premium subscription plan"),
			expectedSendInvoices:                boolPtr(false),
			expectedSendSMS:                     boolPtr(true),
			expectedCurrency:                    enums.CurrencyUSD,
			expectedInvoiceLimit:                intPtr(3),
			expectedUpdateExistingSubscriptions: boolPtr(true),
		},
		{
			name: "builds request with some optional fields",
			setupBuilder: func() *UpdateRequestBuilder {
				return NewUpdateRequestBuilder("Basic Updated", 8000, enums.IntervalAnnually).
					Description("Basic plan update").
					UpdateExistingSubscriptions(false)
			},
			expectedName:                        "Basic Updated",
			expectedAmount:                      8000,
			expectedInterval:                    enums.IntervalAnnually,
			expectedDescription:                 stringPtr("Basic plan update"),
			expectedUpdateExistingSubscriptions: boolPtr(false),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			builder := tt.setupBuilder()
			req := builder.Build()

			assert.Equal(t, tt.expectedName, req.Name)
			assert.Equal(t, tt.expectedAmount, req.Amount)
			assert.Equal(t, tt.expectedInterval, req.Interval)

			if tt.expectedDescription != nil {
				assert.Equal(t, *tt.expectedDescription, req.Description)
			} else {
				assert.Empty(t, req.Description)
			}

			if tt.expectedSendInvoices != nil {
				require.NotNil(t, req.SendInvoices)
				assert.Equal(t, *tt.expectedSendInvoices, *req.SendInvoices)
			} else {
				assert.Nil(t, req.SendInvoices)
			}

			if tt.expectedSendSMS != nil {
				require.NotNil(t, req.SendSMS)
				assert.Equal(t, *tt.expectedSendSMS, *req.SendSMS)
			} else {
				assert.Nil(t, req.SendSMS)
			}

			if tt.expectedCurrency != "" {
				assert.Equal(t, tt.expectedCurrency, req.Currency)
			} else {
				assert.Empty(t, req.Currency)
			}

			if tt.expectedInvoiceLimit != nil {
				require.NotNil(t, req.InvoiceLimit)
				assert.Equal(t, *tt.expectedInvoiceLimit, *req.InvoiceLimit)
			} else {
				assert.Nil(t, req.InvoiceLimit)
			}

			if tt.expectedUpdateExistingSubscriptions != nil {
				require.NotNil(t, req.UpdateExistingSubscriptions)
				assert.Equal(t, *tt.expectedUpdateExistingSubscriptions, *req.UpdateExistingSubscriptions)
			} else {
				assert.Nil(t, req.UpdateExistingSubscriptions)
			}
		})
	}
}

func TestUpdateRequest_JSONSerialization(t *testing.T) {
	tests := []struct {
		name        string
		builder     func() *UpdateRequestBuilder
		expectField func(t *testing.T, unmarshaled map[string]any)
	}{
		{
			name: "serializes minimal request correctly",
			builder: func() *UpdateRequestBuilder {
				return NewUpdateRequestBuilder("Updated Plan", 12000, enums.IntervalMonthly)
			},
			expectField: func(t *testing.T, unmarshaled map[string]any) {
				assert.Equal(t, "Updated Plan", unmarshaled["name"])
				assert.Equal(t, float64(12000), unmarshaled["amount"])
				assert.Equal(t, "monthly", unmarshaled["interval"])
				assert.NotContains(t, unmarshaled, "send_invoices")
				assert.NotContains(t, unmarshaled, "send_sms")
				assert.NotContains(t, unmarshaled, "invoice_limit")
				assert.NotContains(t, unmarshaled, "update_existing_subscriptions")
			},
		},
		{
			name: "serializes complete request correctly",
			builder: func() *UpdateRequestBuilder {
				return NewUpdateRequestBuilder("Premium Updated", 80000, enums.IntervalWeekly).
					Description("Premium subscription updated").
					SendInvoices(false).
					SendSMS(true).
					Currency(enums.CurrencyGHS).
					InvoiceLimit(15).
					UpdateExistingSubscriptions(true)
			},
			expectField: func(t *testing.T, unmarshaled map[string]any) {
				assert.Equal(t, "Premium Updated", unmarshaled["name"])
				assert.Equal(t, float64(80000), unmarshaled["amount"])
				assert.Equal(t, "weekly", unmarshaled["interval"])
				assert.Equal(t, "Premium subscription updated", unmarshaled["description"])
				assert.Equal(t, false, unmarshaled["send_invoices"])
				assert.Equal(t, true, unmarshaled["send_sms"])
				assert.Equal(t, "GHS", unmarshaled["currency"])
				assert.Equal(t, float64(15), unmarshaled["invoice_limit"])
				assert.Equal(t, true, unmarshaled["update_existing_subscriptions"])
			},
		},
		{
			name: "omits empty optional fields",
			builder: func() *UpdateRequestBuilder {
				return NewUpdateRequestBuilder("Basic Updated", 6000, enums.IntervalAnnually)
			},
			expectField: func(t *testing.T, unmarshaled map[string]any) {
				// Required fields should be present
				assert.Contains(t, unmarshaled, "name")
				assert.Contains(t, unmarshaled, "amount")
				assert.Contains(t, unmarshaled, "interval")

				// Optional fields should be omitted when empty
				assert.NotContains(t, unmarshaled, "send_invoices")
				assert.NotContains(t, unmarshaled, "send_sms")
				assert.NotContains(t, unmarshaled, "invoice_limit")
				assert.NotContains(t, unmarshaled, "update_existing_subscriptions")

				if description, exists := unmarshaled["description"]; exists {
					assert.Equal(t, "", description, "description should be empty string if present")
				}

				if currency, exists := unmarshaled["currency"]; exists {
					assert.Equal(t, "", currency, "currency should be empty string if present")
				}
			},
		},
		{
			name: "serializes partial update correctly",
			builder: func() *UpdateRequestBuilder {
				return NewUpdateRequestBuilder("Partial Update", 9000, enums.IntervalBiannual).
					SendInvoices(true).
					UpdateExistingSubscriptions(false)
			},
			expectField: func(t *testing.T, unmarshaled map[string]any) {
				assert.Equal(t, "Partial Update", unmarshaled["name"])
				assert.Equal(t, float64(9000), unmarshaled["amount"])
				assert.Equal(t, "biannually", unmarshaled["interval"])
				assert.Equal(t, true, unmarshaled["send_invoices"])
				assert.Equal(t, false, unmarshaled["update_existing_subscriptions"])

				// These should not be present
				assert.NotContains(t, unmarshaled, "send_sms")
				assert.NotContains(t, unmarshaled, "invoice_limit")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			builder := tt.builder()
			req := builder.Build()

			// Serialize to JSON
			jsonData, err := json.Marshal(req)
			require.NoError(t, err, "should marshal request to JSON without error")

			// Parse back to verify structure
			var unmarshaled map[string]any
			err = json.Unmarshal(jsonData, &unmarshaled)
			require.NoError(t, err, "should unmarshal JSON without error")

			tt.expectField(t, unmarshaled)
		})
	}
}
