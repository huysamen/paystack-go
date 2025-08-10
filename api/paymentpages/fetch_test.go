package paymentpages

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFetchResponse_JSONDeserialization(t *testing.T) {
	tests := []struct {
		name            string
		responseFile    string
		expectedStatus  bool
		expectedMessage string
		expectedName    string
		expectedSlug    string
	}{
		{
			name:            "successful fetch response",
			responseFile:    "fetch_200.json",
			expectedStatus:  true,
			expectedMessage: "Page retrieved",
			expectedName:    "Offering collections",
			expectedSlug:    "5nApBwZkvY",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Read the response file
			responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "paymentpages", tt.responseFile)
			responseData, err := os.ReadFile(responseFilePath)
			require.NoError(t, err, "failed to read response file: %s", responseFilePath)

			// Deserialize the JSON response
			var response FetchResponse
			err = json.Unmarshal(responseData, &response)
			require.NoError(t, err, "failed to unmarshal JSON response")

			// Verify the response structure
			assert.Equal(t, tt.expectedStatus, response.Status.Bool(), "status should match")
			assert.Equal(t, tt.expectedMessage, response.Message, "message should match")

			// Verify the data structure
			assert.NotNil(t, response.Data, "data should not be nil")
			assert.Equal(t, tt.expectedName, response.Data.Name.String(), "name should match")
			assert.Equal(t, tt.expectedSlug, response.Data.Slug.String(), "slug should match")
			assert.Greater(t, response.Data.ID.Int64(), int64(0), "ID should be greater than 0")
		})
	}
}

func TestFetchResponse_FieldByFieldValidation(t *testing.T) {
	t.Run("fetch_200_json_comprehensive_field_validation", func(t *testing.T) {
		// Read the exact JSON response file
		responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "paymentpages", "fetch_200.json")
		responseData, err := os.ReadFile(responseFilePath)
		require.NoError(t, err, "failed to read fetch_200.json")

		// Parse into our struct
		var response FetchResponse
		err = json.Unmarshal(responseData, &response)
		require.NoError(t, err, "failed to unmarshal fetch_200.json")

		// Parse the raw JSON to compare exact values
		var rawJSON map[string]any
		err = json.Unmarshal(responseData, &rawJSON)
		require.NoError(t, err, "failed to parse raw JSON for comparison")

		// Field-by-field validation against the exact JSON values
		assert.Equal(t, true, rawJSON["status"], "status in JSON should be true")
		assert.Equal(t, true, response.Status.Bool(), "status in struct should be true")

		assert.Equal(t, "Page retrieved", rawJSON["message"], "message in JSON should match")
		assert.Equal(t, "Page retrieved", response.Message, "message in struct should match")

		// Verify data field exists and has correct structure
		assert.Contains(t, rawJSON, "data", "JSON should contain data field")
		assert.NotNil(t, response.Data, "struct data field should not be nil")

		// Get the data portion from raw JSON
		rawData, ok := rawJSON["data"].(map[string]any)
		require.True(t, ok, "data field should be an object")

		// Validate each data field
		assert.Equal(t, float64(100032), rawData["integration"], "integration in JSON should match")
		assert.Equal(t, int64(100032), response.Data.Integration.Int64(), "integration in struct should match")

		assert.Equal(t, "test", rawData["domain"], "domain in JSON should match")
		assert.Equal(t, "test", response.Data.Domain.String(), "domain in struct should match")

		assert.Equal(t, "Offering collections", rawData["name"], "name in JSON should match")
		assert.Equal(t, "Offering collections", response.Data.Name.String(), "name in struct should match")

		assert.Equal(t, "Give unto the Lord, and it shall be multiplied ten-fold to you.", rawData["description"], "description in JSON should match")
		assert.Equal(t, "Give unto the Lord, and it shall be multiplied ten-fold to you.", response.Data.Description.String(), "description in struct should match")

		assert.Nil(t, rawData["amount"], "amount in JSON should be null")
		assert.False(t, response.Data.Amount.Valid, "amount in struct should be null")

		assert.Equal(t, "NGN", rawData["currency"], "currency in JSON should match")
		assert.Equal(t, "NGN", response.Data.Currency.String(), "currency in struct should match")

		assert.Equal(t, "5nApBwZkvY", rawData["slug"], "slug in JSON should match")
		assert.Equal(t, "5nApBwZkvY", response.Data.Slug.String(), "slug in struct should match")

		assert.Equal(t, true, rawData["active"], "active in JSON should match")
		assert.Equal(t, true, response.Data.Active.Bool(), "active in struct should match")

		assert.Equal(t, float64(18), rawData["id"], "id in JSON should match")
		assert.Equal(t, int64(18), response.Data.ID.Int64(), "id in struct should match")

		assert.Equal(t, "2016-03-30T00:49:57.000Z", rawData["createdAt"], "createdAt in JSON should match")
		// For timestamp comparison, check that we can parse both correctly rather than string format
		expectedCreatedAt, err := time.Parse(time.RFC3339, "2016-03-30T00:49:57.000Z")
		require.NoError(t, err, "should parse expected createdAt")
		actualCreatedAt, err := time.Parse(time.RFC3339, response.Data.CreatedAt.String())
		require.NoError(t, err, "should parse actual createdAt")
		assert.True(t, expectedCreatedAt.Equal(actualCreatedAt), "createdAt timestamps should be equal")

		assert.Equal(t, "2016-03-30T00:49:57.000Z", rawData["updatedAt"], "updatedAt in JSON should match")
		// For timestamp comparison, check that we can parse both correctly rather than string format
		expectedUpdatedAt, err := time.Parse(time.RFC3339, "2016-03-30T00:49:57.000Z")
		require.NoError(t, err, "should parse expected updatedAt")
		actualUpdatedAt, err := time.Parse(time.RFC3339, response.Data.UpdatedAt.String())
		require.NoError(t, err, "should parse actual updatedAt")
		assert.True(t, expectedUpdatedAt.Equal(actualUpdatedAt), "updatedAt timestamps should be equal")

		// Validate products array
		rawProducts, ok := rawData["products"].([]any)
		require.True(t, ok, "products field should be an array")
		assert.Len(t, response.Data.Products, 2, "should have 2 products")
		assert.Len(t, rawProducts, 2, "should have 2 products in JSON")

		// Validate first product
		rawProduct1, ok := rawProducts[0].(map[string]any)
		require.True(t, ok, "first product should be an object")
		product1 := response.Data.Products[0]

		assert.Equal(t, float64(523), rawProduct1["product_id"], "product_id in JSON should match")
		assert.Equal(t, int64(523), product1.ProductID.Int64(), "product_id in struct should match")

		assert.Equal(t, "Product Four", rawProduct1["name"], "product name in JSON should match")
		assert.Equal(t, "Product Four", product1.Name.String(), "product name in struct should match")

		assert.Equal(t, "Product Four Description", rawProduct1["description"], "product description in JSON should match")
		assert.Equal(t, "Product Four Description", product1.Description.String(), "product description in struct should match")

		assert.Equal(t, "PROD_l9p81u9pkjqjunb", rawProduct1["product_code"], "product_code in JSON should match")
		assert.Equal(t, "PROD_l9p81u9pkjqjunb", product1.ProductCode.String(), "product_code in struct should match")

		assert.Equal(t, float64(18), rawProduct1["page"], "product page in JSON should match")
		assert.Equal(t, int64(18), product1.Page.Int64(), "product page in struct should match")

		assert.Equal(t, float64(500000), rawProduct1["price"], "product price in JSON should match")
		assert.Equal(t, int64(500000), product1.Price.Int64(), "product price in struct should match")

		assert.Equal(t, "NGN", rawProduct1["currency"], "product currency in JSON should match")
		assert.Equal(t, "NGN", product1.Currency.String(), "product currency in struct should match")

		assert.Equal(t, float64(0), rawProduct1["quantity"], "product quantity in JSON should match")
		assert.Equal(t, int64(0), product1.Quantity.Int64(), "product quantity in struct should match")

		assert.Equal(t, "good", rawProduct1["type"], "product type in JSON should match")
		assert.Equal(t, "good", product1.Type.String(), "product type in struct should match")

		assert.Nil(t, rawProduct1["features"], "product features in JSON should be null")

		assert.Equal(t, float64(0), rawProduct1["is_shippable"], "product is_shippable in JSON should match")
		assert.Equal(t, int64(0), product1.IsShippable.Int64(), "product is_shippable in struct should match")

		assert.Equal(t, "test", rawProduct1["domain"], "product domain in JSON should match")
		assert.Equal(t, "test", product1.Domain.String(), "product domain in struct should match")

		assert.Equal(t, float64(343288), rawProduct1["integration"], "product integration in JSON should match")
		assert.Equal(t, int64(343288), product1.Integration.Int64(), "product integration in struct should match")

		assert.Equal(t, float64(1), rawProduct1["active"], "product active in JSON should match")
		assert.Equal(t, int64(1), product1.Active.Int64(), "product active in struct should match")

		assert.Equal(t, float64(1), rawProduct1["in_stock"], "product in_stock in JSON should match")
		assert.Equal(t, int64(1), product1.InStock.Int64(), "product in_stock in struct should match")

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
