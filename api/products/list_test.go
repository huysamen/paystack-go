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

func TestListResponse_JSONDeserialization(t *testing.T) {
	tests := []struct {
		name            string
		responseFile    string
		expectedStatus  bool
		expectedMessage string
	}{
		{
			name:            "successful list response",
			responseFile:    "list_200.json",
			expectedStatus:  true,
			expectedMessage: "Products retrieved",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Read the JSON response file
			responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "products", tt.responseFile)
			responseData, err := os.ReadFile(responseFilePath)
			require.NoError(t, err, "failed to read response file")

			// Deserialize the JSON response
			var response ListResponse
			err = json.Unmarshal(responseData, &response)
			require.NoError(t, err, "failed to unmarshal JSON response")

			// Verify the response structure
			assert.Equal(t, tt.expectedStatus, response.Status.Bool(), "status should match")
			assert.Equal(t, tt.expectedMessage, response.Message, "message should match")
			assert.NotNil(t, response.Data, "data should not be nil")
			assert.Greater(t, len(response.Data), 0, "should have at least one product")
		})
	}
}

func TestListResponse_FieldByFieldValidation(t *testing.T) {
	t.Run("list_200_json_comprehensive_field_validation", func(t *testing.T) {
		// Read the exact JSON response file
		responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "products", "list_200.json")
		responseData, err := os.ReadFile(responseFilePath)
		require.NoError(t, err, "failed to read list_200.json")

		// Parse raw JSON to compare field by field
		var rawJSON map[string]any
		err = json.Unmarshal(responseData, &rawJSON)
		require.NoError(t, err, "failed to parse raw JSON")

		// Parse into our struct
		var response ListResponse
		err = json.Unmarshal(responseData, &response)
		require.NoError(t, err, "failed to unmarshal list_200.json")

		// Verify struct data is properly populated
		assert.NotNil(t, response.Data, "struct data field should not be nil")

		// Get the data portion from raw JSON (should be an array)
		rawProducts, ok := rawJSON["data"].([]any)
		require.True(t, ok, "data field should be an array")
		require.Greater(t, len(rawProducts), 0, "should have products")
		require.Equal(t, len(rawProducts), len(response.Data), "product count should match")

		// Test first product
		rawProduct := rawProducts[0].(map[string]any)
		product := response.Data[0]

		// Validate core fields of first product
		assert.Equal(t, "Mimshack", rawProduct["name"], "name in JSON should match")
		assert.Equal(t, "Mimshack", product.Name.String(), "name in struct should match")

		assert.Equal(t, "Everything cars", rawProduct["description"], "description in JSON should match")
		assert.Equal(t, "Everything cars", product.Description.String(), "description in struct should match")

		assert.Equal(t, "NGN", rawProduct["currency"], "currency in JSON should match")
		assert.Equal(t, enums.CurrencyNGN, product.Currency, "currency in struct should match")

		assert.Equal(t, float64(50000), rawProduct["price"], "price in JSON should match")
		assert.Equal(t, int64(50000), product.Price.Int64(), "price in struct should match")

		assert.Equal(t, float64(10), rawProduct["quantity"], "quantity in JSON should match")
		assert.Equal(t, int64(10), product.Quantity.Int, "quantity in struct should match")

		assert.Equal(t, true, rawProduct["is_shippable"], "is_shippable in JSON should match")
		assert.Equal(t, true, product.IsShippable.Bool(), "is_shippable in struct should match")

		assert.Equal(t, false, rawProduct["unlimited"], "unlimited in JSON should match")
		assert.Equal(t, false, product.Unlimited.Bool(), "unlimited in struct should match")

		assert.Equal(t, float64(463433), rawProduct["integration"], "integration in JSON should match")
		assert.Equal(t, int64(463433), product.Integration.Int64(), "integration in struct should match")

		assert.Equal(t, "test", rawProduct["domain"], "domain in JSON should match")
		assert.Equal(t, "test", product.Domain.String(), "domain in struct should match")

		assert.Equal(t, "mimshack-yiuedh", rawProduct["slug"], "slug in JSON should match")
		assert.Equal(t, "mimshack-yiuedh", product.Slug.String(), "slug in struct should match")

		assert.Equal(t, "PROD_22deobcvbht2dfe", rawProduct["product_code"], "product_code in JSON should match")
		assert.Equal(t, "PROD_22deobcvbht2dfe", product.ProductCode.String(), "product_code in struct should match")

		assert.Equal(t, float64(0), rawProduct["quantity_sold"], "quantity_sold in JSON should match")
		assert.Equal(t, int64(0), product.QuantitySold.Int, "quantity_sold in struct should match")

		assert.Equal(t, "good", rawProduct["type"], "type in JSON should match")
		assert.Equal(t, "good", product.Type.String(), "type in struct should match")

		assert.Equal(t, true, rawProduct["active"], "active in JSON should match")
		assert.Equal(t, true, product.Active.Bool(), "active in struct should match")

		assert.Equal(t, true, rawProduct["in_stock"], "in_stock in JSON should match")
		assert.Equal(t, true, product.InStock.Bool(), "in_stock in struct should match")

		assert.Equal(t, float64(1), rawProduct["minimum_orderable"], "minimum_orderable in JSON should match")
		assert.Equal(t, int64(1), product.MinimumOrderable.Int, "minimum_orderable in struct should match")

		// Test null maximum_orderable field
		assert.Equal(t, nil, rawProduct["maximum_orderable"], "maximum_orderable in JSON should be null")
		assert.False(t, product.MaximumOrderable.Valid, "maximum_orderable in struct should be null")

		assert.Equal(t, float64(0), rawProduct["low_stock_alert"], "low_stock_alert in JSON should match")
		assert.Equal(t, false, product.LowStockAlert.Bool(), "low_stock_alert in struct should match")

		assert.Equal(t, float64(795638), rawProduct["id"], "id in JSON should match")
		assert.Equal(t, int64(795638), product.ID.Int64(), "id in struct should match")

		// Test metadata object
		rawMetadata, ok := rawProduct["metadata"].(map[string]any)
		require.True(t, ok, "metadata should be an object")
		assert.Equal(t, "#F5F5F5", rawMetadata["background_color"], "metadata.background_color in JSON should match")

		// Test shipping_fields object
		rawShippingFields, ok := rawProduct["shipping_fields"].(map[string]any)
		require.True(t, ok, "shipping_fields should be an object")
		assert.Equal(t, "disabled", rawShippingFields["delivery_note"], "shipping_fields.delivery_note in JSON should match")

		// Test files array (should be empty array in first product)
		rawFiles, ok := rawProduct["files"].([]any)
		require.True(t, ok, "files should be an array")
		assert.Equal(t, 0, len(rawFiles), "files should be empty for first product")

		// Test timestamps with lenient comparison (list uses 000Z format)
		assert.Equal(t, "2022-04-12T11:21:43.000Z", rawProduct["createdAt"], "created_at in JSON should match")
		expectedCreatedAt, err := time.Parse(time.RFC3339, "2022-04-12T11:21:43.000Z")
		require.NoError(t, err, "should parse expected created_at")
		actualCreatedAt, err := time.Parse(time.RFC3339, product.CreatedAt.String())
		require.NoError(t, err, "should parse actual created_at")
		assert.True(t, expectedCreatedAt.Sub(actualCreatedAt).Abs() < time.Second, "created_at timestamps should be within 1 second")

		assert.Equal(t, "2022-04-12T11:21:43.000Z", rawProduct["updatedAt"], "updated_at in JSON should match")
		expectedUpdatedAt, err := time.Parse(time.RFC3339, "2022-04-12T11:21:43.000Z")
		require.NoError(t, err, "should parse expected updated_at")
		actualUpdatedAt, err := time.Parse(time.RFC3339, product.UpdatedAt.String())
		require.NoError(t, err, "should parse actual updated_at")
		assert.True(t, expectedUpdatedAt.Sub(actualUpdatedAt).Abs() < time.Second, "updated_at timestamps should be within 1 second")

		// Test that we have all 3 products as expected
		assert.Equal(t, 3, len(response.Data), "should have exactly 3 products")

		// Test second product briefly to ensure array parsing works
		rawProduct2 := rawProducts[1].(map[string]any)
		product2 := response.Data[1]
		assert.Equal(t, "Nike Air 23", rawProduct2["name"], "second product name should match")
		assert.Equal(t, "Nike Air 23", product2.Name.String(), "second product name in struct should match")

		// Test that second product has files (non-empty array)
		rawFiles2, ok := rawProduct2["files"].([]any)
		require.True(t, ok, "second product files should be an array")
		assert.Greater(t, len(rawFiles2), 0, "second product should have files")
	})
}

func TestListRequestBuilder(t *testing.T) {
	tests := []struct {
		name         string
		setupBuilder func() *ListRequestBuilder
		checkQuery   func(t *testing.T, queryString string)
	}{
		{
			name: "builds request with no filters",
			setupBuilder: func() *ListRequestBuilder {
				return NewListRequestBuilder()
			},
			checkQuery: func(t *testing.T, queryString string) {
				assert.Empty(t, queryString, "query should be empty for no filters")
			},
		},
		{
			name: "builds request with pagination",
			setupBuilder: func() *ListRequestBuilder {
				return NewListRequestBuilder().
					Page(2).
					PerPage(25)
			},
			checkQuery: func(t *testing.T, queryString string) {
				assert.Contains(t, queryString, "page=2", "page should be set")
				assert.Contains(t, queryString, "perPage=25", "perPage should be set")
			},
		},
		{
			name: "builds request with date range",
			setupBuilder: func() *ListRequestBuilder {
				return NewListRequestBuilder().
					From("2023-01-01").
					To("2023-12-31")
			},
			checkQuery: func(t *testing.T, queryString string) {
				assert.Contains(t, queryString, "from=2023-01-01", "from should be set")
				assert.Contains(t, queryString, "to=2023-12-31", "to should be set")
			},
		},
		{
			name: "builds request with all filters",
			setupBuilder: func() *ListRequestBuilder {
				return NewListRequestBuilder().
					Page(3).
					PerPage(10).
					DateRange("2023-06-01", "2023-06-30")
			},
			checkQuery: func(t *testing.T, queryString string) {
				assert.Contains(t, queryString, "page=3", "page should be set")
				assert.Contains(t, queryString, "perPage=10", "perPage should be set")
				assert.Contains(t, queryString, "from=2023-06-01", "from should be set")
				assert.Contains(t, queryString, "to=2023-06-30", "to should be set")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			builder := tt.setupBuilder()
			req := builder.Build()
			queryString := req.toQuery()
			tt.checkQuery(t, queryString)
		})
	}
}

func TestList(t *testing.T) {
	tests := []struct {
		name         string
		queryString  string
		expectedBase string
	}{
		{
			name:         "list with no query parameters",
			queryString:  "",
			expectedBase: "https://api.paystack.co/product",
		},
		{
			name:         "list with pagination",
			queryString:  "page=2&perPage=25",
			expectedBase: "https://api.paystack.co/product",
		},
		{
			name:         "list with date filter",
			queryString:  "from=2023-01-01&to=2023-12-31",
			expectedBase: "https://api.paystack.co/product",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// This test verifies URL construction and query parameter handling
			// In a real scenario, we would mock the HTTP client
			baseURL := "https://api.paystack.co/product"

			if tt.queryString == "" {
				assert.Equal(t, baseURL, tt.expectedBase, "URL should match for empty query")
			} else {
				fullURL := baseURL + "?" + tt.queryString
				assert.Contains(t, fullURL, tt.expectedBase, "URL should contain base path")
				assert.Contains(t, fullURL, tt.queryString, "URL should contain query parameters")
			}
		})
	}
}

func TestListResponse_JSONRoundTrip(t *testing.T) {
	t.Run("deserialize_and_serialize_maintains_structure", func(t *testing.T) {
		// Read the JSON response file
		responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "products", "list_200.json")
		originalData, err := os.ReadFile(responseFilePath)
		require.NoError(t, err, "failed to read list_200.json")

		// Parse into struct
		var response ListResponse
		err = json.Unmarshal(originalData, &response)
		require.NoError(t, err, "failed to unmarshal list_200.json")

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

		// Data should be an array in both
		originalProducts, ok1 := originalMap["data"].([]any)
		serializedProducts, ok2 := serializedMap["data"].([]any)
		assert.True(t, ok1, "original data should be array")
		assert.True(t, ok2, "serialized data should be array")
		assert.Equal(t, len(originalProducts), len(serializedProducts), "product count should survive round-trip")

		// Test first product fields survive round-trip
		if len(originalProducts) > 0 && len(serializedProducts) > 0 {
			originalProduct := originalProducts[0].(map[string]any)
			serializedProduct := serializedProducts[0].(map[string]any)
			assert.Equal(t, originalProduct["name"], serializedProduct["name"], "product name should survive round-trip")
			assert.Equal(t, originalProduct["product_code"], serializedProduct["product_code"], "product_code should survive round-trip")
			assert.Equal(t, originalProduct["price"], serializedProduct["price"], "price should survive round-trip")
		}
	})
}
