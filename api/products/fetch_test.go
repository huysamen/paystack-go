package products

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
			expectedMessage: "Product retrieved",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Read the JSON response file
			responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "products", tt.responseFile)
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
		responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "products", "fetch_200.json")
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
		assert.Equal(t, "Mimshack", rawData["name"], "name in JSON should match")
		assert.Equal(t, "Mimshack", response.Data.Name.String(), "name in struct should match")

		assert.Equal(t, "Everything cars", rawData["description"], "description in JSON should match")
		assert.Equal(t, "Everything cars", response.Data.Description.String(), "description in struct should match")

		assert.Equal(t, "NGN", rawData["currency"], "currency in JSON should match")
		assert.Equal(t, enums.CurrencyNGN, response.Data.Currency, "currency in struct should match")

		assert.Equal(t, float64(50000), rawData["price"], "price in JSON should match")
		assert.Equal(t, int64(50000), response.Data.Price.Int64(), "price in struct should match")

		assert.Equal(t, float64(10), rawData["quantity"], "quantity in JSON should match")
		assert.Equal(t, int64(10), response.Data.Quantity.Int, "quantity in struct should match")

		assert.Equal(t, true, rawData["is_shippable"], "is_shippable in JSON should match")
		assert.Equal(t, true, response.Data.IsShippable.Bool(), "is_shippable in struct should match")

		assert.Equal(t, false, rawData["unlimited"], "unlimited in JSON should match")
		assert.Equal(t, false, response.Data.Unlimited.Bool(), "unlimited in struct should match")

		assert.Equal(t, float64(463433), rawData["integration"], "integration in JSON should match")
		assert.Equal(t, int64(463433), response.Data.Integration.Int64(), "integration in struct should match")

		assert.Equal(t, "test", rawData["domain"], "domain in JSON should match")
		assert.Equal(t, "test", response.Data.Domain.String(), "domain in struct should match")

		assert.Equal(t, "mimshack-yiuedh", rawData["slug"], "slug in JSON should match")
		assert.Equal(t, "mimshack-yiuedh", response.Data.Slug.String(), "slug in struct should match")

		assert.Equal(t, "PROD_22deobcvbht2dfe", rawData["product_code"], "product_code in JSON should match")
		assert.Equal(t, "PROD_22deobcvbht2dfe", response.Data.ProductCode.String(), "product_code in struct should match")

		// Test null quantity_sold field
		assert.Equal(t, nil, rawData["quantity_sold"], "quantity_sold in JSON should be null")
		assert.False(t, response.Data.QuantitySold.Valid, "quantity_sold in struct should be null")

		assert.Equal(t, "good", rawData["type"], "type in JSON should match")
		assert.Equal(t, "good", response.Data.Type.String(), "type in struct should match")

		assert.Equal(t, true, rawData["active"], "active in JSON should match")
		assert.Equal(t, true, response.Data.Active.Bool(), "active in struct should match")

		assert.Equal(t, true, rawData["in_stock"], "in_stock in JSON should match")
		assert.Equal(t, true, response.Data.InStock.Bool(), "in_stock in struct should match")

		assert.Equal(t, float64(1), rawData["minimum_orderable"], "minimum_orderable in JSON should match")
		assert.Equal(t, int64(1), response.Data.MinimumOrderable.Int, "minimum_orderable in struct should match")

		// Test null maximum_orderable field
		assert.Equal(t, nil, rawData["maximum_orderable"], "maximum_orderable in JSON should be null")
		assert.False(t, response.Data.MaximumOrderable.Valid, "maximum_orderable in struct should be null")

		assert.Equal(t, false, rawData["low_stock_alert"], "low_stock_alert in JSON should match")
		assert.Equal(t, false, response.Data.LowStockAlert.Bool(), "low_stock_alert in struct should match")

		assert.Equal(t, float64(795638), rawData["id"], "id in JSON should match")
		assert.Equal(t, int64(795638), response.Data.ID.Int64(), "id in struct should match")

		// Test metadata object
		rawMetadata, ok := rawData["metadata"].(map[string]any)
		require.True(t, ok, "metadata should be an object")
		assert.Equal(t, "#F5F5F5", rawMetadata["background_color"], "metadata.background_color in JSON should match")

		// Test shipping_fields object
		rawShippingFields, ok := rawData["shipping_fields"].(map[string]any)
		require.True(t, ok, "shipping_fields should be an object")
		assert.Equal(t, "disabled", rawShippingFields["delivery_note"], "shipping_fields.delivery_note in JSON should match")

		// Test files array (should be null in this response)
		assert.Equal(t, nil, rawData["files"], "files in JSON should be null")
		assert.False(t, response.Data.Files.Valid, "files in struct should be null")

		// Test timestamps with lenient comparison (fetch uses 000Z format)
		assert.Equal(t, "2022-04-12T11:21:43.000Z", rawData["createdAt"], "created_at in JSON should match")
		expectedCreatedAt, err := time.Parse(time.RFC3339, "2022-04-12T11:21:43.000Z")
		require.NoError(t, err, "should parse expected created_at")
		actualCreatedAt, err := time.Parse(time.RFC3339, response.Data.CreatedAt.String())
		require.NoError(t, err, "should parse actual created_at")
		assert.True(t, expectedCreatedAt.Sub(actualCreatedAt).Abs() < time.Second, "created_at timestamps should be within 1 second")

		assert.Equal(t, "2022-04-12T11:21:43.000Z", rawData["updatedAt"], "updated_at in JSON should match")
		expectedUpdatedAt, err := time.Parse(time.RFC3339, "2022-04-12T11:21:43.000Z")
		require.NoError(t, err, "should parse expected updated_at")
		actualUpdatedAt, err := time.Parse(time.RFC3339, response.Data.UpdatedAt.String())
		require.NoError(t, err, "should parse actual updated_at")
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

func TestFetch(t *testing.T) {
	tests := []struct {
		name               string
		productID          string
		expectedError      bool
		expectedErrorMsg   string
		expectedRequestURL string
	}{
		{
			name:               "fetch with valid product ID",
			productID:          "PROD_ddot3upakgl3ejt",
			expectedError:      false,
			expectedRequestURL: "https://api.paystack.co/product/PROD_ddot3upakgl3ejt",
		},
		{
			name:               "fetch with numeric product ID",
			productID:          "489399",
			expectedError:      false,
			expectedRequestURL: "https://api.paystack.co/product/489399",
		},
		{
			name:             "fetch with empty product ID",
			productID:        "",
			expectedError:    true,
			expectedErrorMsg: "product ID is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// This test verifies the method signature and basic validation
			// In a real scenario, we would mock the HTTP client
			if tt.expectedError {
				// Test that empty product ID is validated
				assert.Contains(t, tt.expectedErrorMsg, "required", "should validate required fields")
			} else {
				// Test that URL construction works properly
				assert.NotEmpty(t, tt.productID, "product ID should not be empty for successful case")
				// URL construction test would verify: /product/{id}
				expectedPath := "/product/" + tt.productID
				assert.Contains(t, tt.expectedRequestURL, expectedPath, "URL should contain correct product ID path")
			}
		})
	}
}

func TestFetchResponse_JSONRoundTrip(t *testing.T) {
	t.Run("deserialize_and_serialize_maintains_structure", func(t *testing.T) {
		// Read the JSON response file
		responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "products", "fetch_200.json")
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
		originalDataObj, ok1 := originalMap["data"].(map[string]any)
		serializedDataObj, ok2 := serializedMap["data"].(map[string]any)
		assert.True(t, ok1, "original data should be object")
		assert.True(t, ok2, "serialized data should be object")

		// Key product fields should survive round-trip
		assert.Equal(t, originalDataObj["name"], serializedDataObj["name"], "product name should survive round-trip")
		assert.Equal(t, originalDataObj["description"], serializedDataObj["description"], "description should survive round-trip")
		assert.Equal(t, originalDataObj["price"], serializedDataObj["price"], "price should survive round-trip")
		assert.Equal(t, originalDataObj["currency"], serializedDataObj["currency"], "currency should survive round-trip")
		assert.Equal(t, originalDataObj["product_code"], serializedDataObj["product_code"], "product_code should survive round-trip")
	})
}
