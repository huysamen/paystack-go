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

func TestAddProductsResponse_JSONDeserialization(t *testing.T) {
	tests := []struct {
		name            string
		responseFile    string
		expectedStatus  bool
		expectedMessage string
		expectedName    string
	}{
		{
			name:            "successful add products response",
			responseFile:    "add_products_200.json",
			expectedStatus:  true,
			expectedMessage: "Products added to page",
			expectedName:    "Demo Products Page",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Read the response file
			responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "paymentpages", tt.responseFile)
			responseData, err := os.ReadFile(responseFilePath)
			require.NoError(t, err, "failed to read response file: %s", responseFilePath)

			// Deserialize the JSON response
			var response AddProductsResponse
			err = json.Unmarshal(responseData, &response)
			require.NoError(t, err, "failed to unmarshal JSON response")

			// Verify the response structure
			assert.Equal(t, tt.expectedStatus, response.Status.Bool(), "status should match")
			assert.Equal(t, tt.expectedMessage, response.Message, "message should match")

			// Verify the data structure
			assert.NotNil(t, response.Data, "data should not be nil")
			assert.Equal(t, tt.expectedName, response.Data.Name.String(), "name should match")
			assert.Greater(t, response.Data.ID.Int64(), int64(0), "ID should be greater than 0")
		})
	}
}

func TestAddProductsResponse_FieldByFieldValidation(t *testing.T) {
	t.Run("add_products_200_json_comprehensive_field_validation", func(t *testing.T) {
		// Read the exact JSON response file
		responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "paymentpages", "add_products_200.json")
		responseData, err := os.ReadFile(responseFilePath)
		require.NoError(t, err, "failed to read add_products_200.json")

		// Parse into our struct
		var response AddProductsResponse
		err = json.Unmarshal(responseData, &response)
		require.NoError(t, err, "failed to unmarshal add_products_200.json")

		// Parse the raw JSON to compare exact values
		var rawJSON map[string]any
		err = json.Unmarshal(responseData, &rawJSON)
		require.NoError(t, err, "failed to parse raw JSON for comparison")

		// Field-by-field validation against the exact JSON values
		assert.Equal(t, true, rawJSON["status"], "status in JSON should be true")
		assert.Equal(t, true, response.Status.Bool(), "status in struct should be true")

		assert.Equal(t, "Products added to page", rawJSON["message"], "message in JSON should match")
		assert.Equal(t, "Products added to page", response.Message, "message in struct should match")

		// Verify data field exists and has correct structure
		assert.Contains(t, rawJSON, "data", "JSON should contain data field")
		assert.NotNil(t, response.Data, "struct data field should not be nil")

		// Get the data portion from raw JSON
		rawData, ok := rawJSON["data"].(map[string]any)
		require.True(t, ok, "data field should be an object")

		// Validate key data fields
		assert.Equal(t, float64(343288), rawData["integration"], "integration in JSON should match")
		assert.Equal(t, int64(343288), response.Data.Integration.Int64(), "integration in struct should match")

		assert.Equal(t, "test", rawData["domain"], "domain in JSON should match")
		assert.Equal(t, "test", response.Data.Domain.String(), "domain in struct should match")

		assert.Equal(t, "Demo Products Page", rawData["name"], "name in JSON should match")
		assert.Equal(t, "Demo Products Page", response.Data.Name.String(), "name in struct should match")

		assert.Equal(t, "Demo Products Page", rawData["description"], "description in JSON should match")
		assert.Equal(t, "Demo Products Page", response.Data.Description.String(), "description in struct should match")

		assert.Nil(t, rawData["amount"], "amount in JSON should be null")
		assert.False(t, response.Data.Amount.Valid, "amount in struct should be null")

		assert.Equal(t, "NGN", rawData["currency"], "currency in JSON should match")
		assert.Equal(t, "NGN", response.Data.Currency.String(), "currency in struct should match")

		assert.Equal(t, "demoproductspage", rawData["slug"], "slug in JSON should match")
		assert.Equal(t, "demoproductspage", response.Data.Slug.String(), "slug in struct should match")

		assert.Equal(t, true, rawData["active"], "active in JSON should match")
		assert.Equal(t, true, response.Data.Active.Bool(), "active in struct should match")

		assert.Equal(t, float64(102859), rawData["id"], "id in JSON should match")
		assert.Equal(t, int64(102859), response.Data.ID.Int64(), "id in struct should match")

		assert.Equal(t, "2019-06-29T16:21:24.000Z", rawData["createdAt"], "createdAt in JSON should match")
		// For timestamp comparison, check that we can parse both correctly rather than string format
		expectedCreatedAt, err := time.Parse(time.RFC3339, "2019-06-29T16:21:24.000Z")
		require.NoError(t, err, "should parse expected createdAt")
		actualCreatedAt, err := time.Parse(time.RFC3339, response.Data.CreatedAt.String())
		require.NoError(t, err, "should parse actual createdAt")
		assert.True(t, expectedCreatedAt.Equal(actualCreatedAt), "createdAt timestamps should be equal")

		assert.Equal(t, "2019-06-29T16:28:11.000Z", rawData["updatedAt"], "updatedAt in JSON should match")
		// For timestamp comparison, check that we can parse both correctly rather than string format
		expectedUpdatedAt, err := time.Parse(time.RFC3339, "2019-06-29T16:28:11.000Z")
		require.NoError(t, err, "should parse expected updatedAt")
		actualUpdatedAt, err := time.Parse(time.RFC3339, response.Data.UpdatedAt.String())
		require.NoError(t, err, "should parse actual updatedAt")
		assert.True(t, expectedUpdatedAt.Equal(actualUpdatedAt), "updatedAt timestamps should be equal")

		// Validate products array
		rawProducts, ok := rawData["products"].([]any)
		require.True(t, ok, "products field should be an array")
		assert.Greater(t, len(response.Data.Products), 0, "should have products")
		assert.Equal(t, len(rawProducts), len(response.Data.Products), "should have same number of products")

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

func TestAddProductsRequestBuilder(t *testing.T) {
	t.Run("builds request with single product", func(t *testing.T) {
		builder := NewAddProductsRequestBuilder()
		request := builder.
			AddProduct(123).
			Build()

		assert.Equal(t, []int{123}, request.Product)
	})

	t.Run("builds request with multiple products", func(t *testing.T) {
		builder := NewAddProductsRequestBuilder()
		request := builder.
			AddProduct(123).
			AddProduct(456).
			AddProduct(789).
			Build()

		assert.Equal(t, []int{123, 456, 789}, request.Product)
	})

	t.Run("builds request with products slice", func(t *testing.T) {
		builder := NewAddProductsRequestBuilder()
		products := []int{111, 222, 333}
		request := builder.
			Products(products).
			Build()

		assert.Equal(t, products, request.Product)
	})

	t.Run("builds empty request", func(t *testing.T) {
		builder := NewAddProductsRequestBuilder()
		request := builder.Build()

		assert.Equal(t, []int{}, request.Product)
	})
}

func TestAddProductsRequest_JSONSerialization(t *testing.T) {
	t.Run("serializes request correctly", func(t *testing.T) {
		builder := NewAddProductsRequestBuilder().
			AddProduct(123).
			AddProduct(456)

		request := builder.Build()

		jsonData, err := json.Marshal(request)
		require.NoError(t, err, "should marshal to JSON without error")

		var unmarshaled map[string]any
		err = json.Unmarshal(jsonData, &unmarshaled)
		require.NoError(t, err, "should unmarshal JSON without error")

		products, ok := unmarshaled["product"].([]any)
		require.True(t, ok, "product should be an array")
		assert.Equal(t, float64(123), products[0])
		assert.Equal(t, float64(456), products[1])
	})
}
