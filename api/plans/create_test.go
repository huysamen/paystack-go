package plans

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
			expectedMessage: "Plan created",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Read the JSON response file
			responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "plans", tt.responseFile)
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
		responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "plans", "create_200.json")
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

		// Validate each field
		assert.Equal(t, "Monthly retainer", rawData["name"], "name in JSON should match")
		assert.Equal(t, "Monthly retainer", response.Data.Name.String(), "name in struct should match")

		assert.Equal(t, float64(500000), rawData["amount"], "amount in JSON should match")
		assert.Equal(t, int64(500000), response.Data.Amount.Int64(), "amount in struct should match")

		assert.Equal(t, "monthly", rawData["interval"], "interval in JSON should match")
		assert.Equal(t, enums.IntervalMonthly, response.Data.Interval, "interval in struct should match")

		assert.Equal(t, float64(100032), rawData["integration"], "integration in JSON should match")
		assert.Equal(t, int64(100032), response.Data.Integration.Int64(), "integration in struct should match")

		assert.Equal(t, "test", rawData["domain"], "domain in JSON should match")
		assert.Equal(t, "test", response.Data.Domain.String(), "domain in struct should match")

		assert.Equal(t, "PLN_gx2wn530m0i3w3m", rawData["plan_code"], "plan_code in JSON should match")
		assert.Equal(t, "PLN_gx2wn530m0i3w3m", response.Data.PlanCode.String(), "plan_code in struct should match")

		assert.Equal(t, true, rawData["send_invoices"], "send_invoices in JSON should match")
		assert.Equal(t, true, response.Data.SendInvoices.Bool(), "send_invoices in struct should match")

		assert.Equal(t, true, rawData["send_sms"], "send_sms in JSON should match")
		assert.Equal(t, true, response.Data.SendSms.Bool(), "send_sms in struct should match")

		assert.Equal(t, false, rawData["hosted_page"], "hosted_page in JSON should match")
		assert.Equal(t, false, response.Data.HostedPage.Bool(), "hosted_page in struct should match")

		assert.Equal(t, "NGN", rawData["currency"], "currency in JSON should match")
		assert.Equal(t, enums.CurrencyNGN, response.Data.Currency, "currency in struct should match")

		assert.Equal(t, float64(28), rawData["id"], "id in JSON should match")
		assert.Equal(t, uint64(28), response.Data.ID.Uint64(), "id in struct should match")

		assert.Equal(t, "2016-03-29T22:42:50.811Z", rawData["createdAt"], "created_at in JSON should match")
		expectedCreatedAt, err := time.Parse(time.RFC3339, "2016-03-29T22:42:50.811Z")
		require.NoError(t, err, "should parse expected created_at")
		actualCreatedAt, err := time.Parse(time.RFC3339, response.Data.CreatedAt.String())
		require.NoError(t, err, "should parse actual created_at")
		// Allow for small timestamp differences due to precision
		assert.True(t, expectedCreatedAt.Sub(actualCreatedAt).Abs() < time.Second, "created_at timestamps should be within 1 second")

		assert.Equal(t, "2016-03-29T22:42:50.811Z", rawData["updatedAt"], "updated_at in JSON should match")
		expectedUpdatedAt, err := time.Parse(time.RFC3339, "2016-03-29T22:42:50.811Z")
		require.NoError(t, err, "should parse expected updated_at")
		actualUpdatedAt, err := time.Parse(time.RFC3339, response.Data.UpdatedAt.String())
		require.NoError(t, err, "should parse actual updated_at")
		// Allow for small timestamp differences due to precision
		assert.True(t, expectedUpdatedAt.Sub(actualUpdatedAt).Abs() < time.Second, "updated_at timestamps should be within 1 second")

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
	tests := []struct {
		name                 string
		setupBuilder         func() *CreateRequestBuilder
		expectedName         string
		expectedAmount       int
		expectedInterval     enums.Interval
		expectedDescription  *string
		expectedSendInvoices *bool
		expectedSendSMS      *bool
		expectedCurrency     enums.Currency
		expectedInvoiceLimit *int
	}{
		{
			name: "builds request with required fields only",
			setupBuilder: func() *CreateRequestBuilder {
				return NewCreateRequestBuilder("Basic Plan", 10000, enums.IntervalMonthly)
			},
			expectedName:     "Basic Plan",
			expectedAmount:   10000,
			expectedInterval: enums.IntervalMonthly,
		},
		{
			name: "builds request with all fields",
			setupBuilder: func() *CreateRequestBuilder {
				limit := 5
				return NewCreateRequestBuilder("Premium Plan", 50000, enums.IntervalWeekly).
					Description("A premium subscription plan").
					SendInvoices(true).
					SendSMS(false).
					Currency(enums.CurrencyUSD).
					InvoiceLimit(limit)
			},
			expectedName:         "Premium Plan",
			expectedAmount:       50000,
			expectedInterval:     enums.IntervalWeekly,
			expectedDescription:  stringPtr("A premium subscription plan"),
			expectedSendInvoices: boolPtr(true),
			expectedSendSMS:      boolPtr(false),
			expectedCurrency:     enums.CurrencyUSD,
			expectedInvoiceLimit: intPtr(5),
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
		})
	}
}

func TestCreateRequest_JSONSerialization(t *testing.T) {
	tests := []struct {
		name        string
		builder     func() *CreateRequestBuilder
		expectField func(t *testing.T, unmarshaled map[string]any)
	}{
		{
			name: "serializes minimal request correctly",
			builder: func() *CreateRequestBuilder {
				return NewCreateRequestBuilder("Basic Plan", 10000, enums.IntervalMonthly)
			},
			expectField: func(t *testing.T, unmarshaled map[string]any) {
				assert.Equal(t, "Basic Plan", unmarshaled["name"])
				assert.Equal(t, float64(10000), unmarshaled["amount"])
				assert.Equal(t, "monthly", unmarshaled["interval"])
				assert.NotContains(t, unmarshaled, "send_invoices")
				assert.NotContains(t, unmarshaled, "send_sms")
				assert.NotContains(t, unmarshaled, "invoice_limit")
			},
		},
		{
			name: "serializes complete request correctly",
			builder: func() *CreateRequestBuilder {
				return NewCreateRequestBuilder("Premium Plan", 50000, enums.IntervalWeekly).
					Description("Premium subscription").
					SendInvoices(true).
					SendSMS(false).
					Currency(enums.CurrencyUSD).
					InvoiceLimit(10)
			},
			expectField: func(t *testing.T, unmarshaled map[string]any) {
				assert.Equal(t, "Premium Plan", unmarshaled["name"])
				assert.Equal(t, float64(50000), unmarshaled["amount"])
				assert.Equal(t, "weekly", unmarshaled["interval"])
				assert.Equal(t, "Premium subscription", unmarshaled["description"])
				assert.Equal(t, true, unmarshaled["send_invoices"])
				assert.Equal(t, false, unmarshaled["send_sms"])
				assert.Equal(t, "USD", unmarshaled["currency"])
				assert.Equal(t, float64(10), unmarshaled["invoice_limit"])
			},
		},
		{
			name: "omits empty optional fields",
			builder: func() *CreateRequestBuilder {
				return NewCreateRequestBuilder("Basic Plan", 5000, enums.IntervalAnnually)
			},
			expectField: func(t *testing.T, unmarshaled map[string]any) {
				// Required field should be present
				assert.Contains(t, unmarshaled, "name")
				assert.Contains(t, unmarshaled, "amount")
				assert.Contains(t, unmarshaled, "interval")

				// Optional fields should be omitted when empty
				assert.NotContains(t, unmarshaled, "send_invoices")
				assert.NotContains(t, unmarshaled, "send_sms")
				assert.NotContains(t, unmarshaled, "invoice_limit")

				if description, exists := unmarshaled["description"]; exists {
					assert.Equal(t, "", description, "description should be empty string if present")
				}

				if currency, exists := unmarshaled["currency"]; exists {
					assert.Equal(t, "", currency, "currency should be empty string if present")
				}
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

// Helper functions
func stringPtr(s string) *string {
	return &s
}

func boolPtr(b bool) *bool {
	return &b
}

func intPtr(i int) *int {
	return &i
}
