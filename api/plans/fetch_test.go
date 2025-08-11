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
			expectedMessage: "Plan retrieved",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Read the JSON response file
			responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "plans", tt.responseFile)
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
		responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "plans", "fetch_200.json")
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
		assert.Equal(t, "Monthly retainer", rawData["name"], "name in JSON should match")
		assert.Equal(t, "Monthly retainer", response.Data.Name.String(), "name in struct should match")

		assert.Equal(t, float64(50000), rawData["amount"], "amount in JSON should match")
		assert.Equal(t, int64(50000), response.Data.Amount.Int64(), "amount in struct should match")

		assert.Equal(t, "monthly", rawData["interval"], "interval in JSON should match")
		assert.Equal(t, enums.IntervalMonthly, response.Data.Interval, "interval in struct should match")

		assert.Equal(t, float64(100032), rawData["integration"], "integration in JSON should match")
		assert.Equal(t, int64(100032), response.Data.Integration.Int64(), "integration in struct should match")

		assert.Equal(t, "test", rawData["domain"], "domain in JSON should match")
		assert.Equal(t, "test", response.Data.Domain.String(), "domain in struct should match")

		assert.Equal(t, "PLN_gx2wn530m0i3w3m", rawData["plan_code"], "plan_code in JSON should match")
		assert.Equal(t, "PLN_gx2wn530m0i3w3m", response.Data.PlanCode.String(), "plan_code in struct should match")

		// Null description field
		assert.Equal(t, nil, rawData["description"], "description in JSON should be null")
		assert.False(t, response.Data.Description.Valid, "description in struct should be null")

		assert.Equal(t, true, rawData["send_invoices"], "send_invoices in JSON should match")
		assert.Equal(t, true, response.Data.SendInvoices.Bool(), "send_invoices in struct should match")

		assert.Equal(t, true, rawData["send_sms"], "send_sms in JSON should match")
		assert.Equal(t, true, response.Data.SendSms.Bool(), "send_sms in struct should match")

		assert.Equal(t, false, rawData["hosted_page"], "hosted_page in JSON should match")
		assert.Equal(t, false, response.Data.HostedPage.Bool(), "hosted_page in struct should match")

		// Null hosted_page_url field
		assert.Equal(t, nil, rawData["hosted_page_url"], "hosted_page_url in JSON should be null")
		assert.False(t, response.Data.HostedPageURL.Valid, "hosted_page_url in struct should be null")

		// Null hosted_page_summary field
		assert.Equal(t, nil, rawData["hosted_page_summary"], "hosted_page_summary in JSON should be null")
		assert.False(t, response.Data.HostedPageSummary.Valid, "hosted_page_summary in struct should be null")

		assert.Equal(t, "NGN", rawData["currency"], "currency in JSON should match")
		assert.Equal(t, enums.CurrencyNGN, response.Data.Currency, "currency in struct should match")

		assert.Equal(t, float64(28), rawData["id"], "id in JSON should match")
		assert.Equal(t, uint64(28), response.Data.ID.Uint64(), "id in struct should match")

		assert.Equal(t, "2016-03-29T22:42:50.000Z", rawData["createdAt"], "created_at in JSON should match")
		expectedCreatedAt, err := time.Parse(time.RFC3339, "2016-03-29T22:42:50.000Z")
		require.NoError(t, err, "should parse expected created_at")
		actualCreatedAt, err := time.Parse(time.RFC3339, response.Data.CreatedAt.String())
		require.NoError(t, err, "should parse actual created_at")
		assert.True(t, expectedCreatedAt.Equal(actualCreatedAt), "created_at timestamps should be equal")

		assert.Equal(t, "2016-03-29T22:42:50.000Z", rawData["updatedAt"], "updated_at in JSON should match")
		expectedUpdatedAt, err := time.Parse(time.RFC3339, "2016-03-29T22:42:50.000Z")
		require.NoError(t, err, "should parse expected updated_at")
		actualUpdatedAt, err := time.Parse(time.RFC3339, response.Data.UpdatedAt.String())
		require.NoError(t, err, "should parse actual updated_at")
		assert.True(t, expectedUpdatedAt.Equal(actualUpdatedAt), "updated_at timestamps should be equal")

		// Verify subscriptions array is empty
		subscriptionsRaw, ok := rawData["subscriptions"].([]any)
		require.True(t, ok, "subscriptions should be an array")
		assert.Empty(t, subscriptionsRaw, "subscriptions should be empty array")
		assert.Empty(t, response.Data.Subscriptions, "struct subscriptions should be empty")

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
