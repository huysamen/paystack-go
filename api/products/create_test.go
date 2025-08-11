package products

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/huysamen/paystack-go/enums"
	"github.com/huysamen/paystack-go/types"
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
			expectedMessage: "Product successfully created",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Read the JSON response file
			responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "products", tt.responseFile)
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
		responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "products", "create_200.json")
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

		// Validate core fields
		assert.Equal(t, "Puff Puff", rawData["name"], "name in JSON should match")
		assert.Equal(t, "Puff Puff", response.Data.Name.String(), "name in struct should match")

		assert.Equal(t, "Crispy flour ball with fluffy interior", rawData["description"], "description in JSON should match")
		assert.Equal(t, "Crispy flour ball with fluffy interior", response.Data.Description.String(), "description in struct should match")

		assert.Equal(t, "NGN", rawData["currency"], "currency in JSON should match")
		assert.Equal(t, enums.CurrencyNGN, response.Data.Currency, "currency in struct should match")

		assert.Equal(t, float64(5000), rawData["price"], "price in JSON should match")
		assert.Equal(t, int64(5000), response.Data.Price.Int64(), "price in struct should match")

		assert.Equal(t, float64(100), rawData["quantity"], "quantity in JSON should match")
		assert.Equal(t, int64(100), response.Data.Quantity.Int, "quantity in struct should match")

		assert.Equal(t, true, rawData["is_shippable"], "is_shippable in JSON should match")
		assert.Equal(t, true, response.Data.IsShippable.Bool(), "is_shippable in struct should match")

		assert.Equal(t, false, rawData["unlimited"], "unlimited in JSON should match")
		assert.Equal(t, false, response.Data.Unlimited.Bool(), "unlimited in struct should match")

		assert.Equal(t, float64(463433), rawData["integration"], "integration in JSON should match")
		assert.Equal(t, int64(463433), response.Data.Integration.Int64(), "integration in struct should match")

		assert.Equal(t, "test", rawData["domain"], "domain in JSON should match")
		assert.Equal(t, "test", response.Data.Domain.String(), "domain in struct should match")

		assert.Equal(t, "puff-puff-prqnxc", rawData["slug"], "slug in JSON should match")
		assert.Equal(t, "puff-puff-prqnxc", response.Data.Slug.String(), "slug in struct should match")

		assert.Equal(t, "PROD_ddot3upakgl3ejt", rawData["product_code"], "product_code in JSON should match")
		assert.Equal(t, "PROD_ddot3upakgl3ejt", response.Data.ProductCode.String(), "product_code in struct should match")

		assert.Equal(t, float64(0), rawData["quantity_sold"], "quantity_sold in JSON should match")
		assert.Equal(t, int64(0), response.Data.QuantitySold.Int, "quantity_sold in struct should match")

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

		assert.Equal(t, float64(489399), rawData["id"], "id in JSON should match")
		assert.Equal(t, int64(489399), response.Data.ID.Int64(), "id in struct should match")

		// Test metadata object
		rawMetadata, ok := rawData["metadata"].(map[string]any)
		require.True(t, ok, "metadata should be an object")
		assert.Equal(t, "#F5F5F5", rawMetadata["background_color"], "metadata.background_color in JSON should match")

		// Test shipping_fields object
		rawShippingFields, ok := rawData["shipping_fields"].(map[string]any)
		require.True(t, ok, "shipping_fields should be an object")
		assert.Equal(t, "disabled", rawShippingFields["delivery_note"], "shipping_fields.delivery_note in JSON should match")

		// Test timestamps with lenient comparison
		assert.Equal(t, "2021-11-08T14:39:37.303Z", rawData["createdAt"], "created_at in JSON should match")
		expectedCreatedAt, err := time.Parse(time.RFC3339, "2021-11-08T14:39:37.303Z")
		require.NoError(t, err, "should parse expected created_at")
		actualCreatedAt, err := time.Parse(time.RFC3339, response.Data.CreatedAt.String())
		require.NoError(t, err, "should parse actual created_at")
		assert.True(t, expectedCreatedAt.Sub(actualCreatedAt).Abs() < time.Second, "created_at timestamps should be within 1 second")

		assert.Equal(t, "2021-11-08T14:39:37.303Z", rawData["updatedAt"], "updated_at in JSON should match")
		expectedUpdatedAt, err := time.Parse(time.RFC3339, "2021-11-08T14:39:37.303Z")
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

func TestCreateRequestBuilder(t *testing.T) {
	tests := []struct {
		name              string
		setupBuilder      func() *CreateRequestBuilder
		expectedName      string
		expectedDesc      string
		expectedPrice     int
		expectedCurrency  string
		expectedUnlimited *bool
		expectedQuantity  *int
	}{
		{
			name: "builds request with required fields only",
			setupBuilder: func() *CreateRequestBuilder {
				return NewCreateRequestBuilder("Test Product", "A test product", 1000, "NGN")
			},
			expectedName:     "Test Product",
			expectedDesc:     "A test product",
			expectedPrice:    1000,
			expectedCurrency: "NGN",
		},
		{
			name: "builds request with all fields",
			setupBuilder: func() *CreateRequestBuilder {
				metadata := types.NewMetadata(map[string]any{
					"color": "blue",
					"size":  "large",
				})
				return NewCreateRequestBuilder("Premium Product", "Premium description", 5000, "USD").
					Unlimited(true).
					Quantity(50).
					Metadata(&metadata)
			},
			expectedName:      "Premium Product",
			expectedDesc:      "Premium description",
			expectedPrice:     5000,
			expectedCurrency:  "USD",
			expectedUnlimited: boolPtr(true),
			expectedQuantity:  intPtr(50),
		},
		{
			name: "builds request with limited quantity",
			setupBuilder: func() *CreateRequestBuilder {
				return NewCreateRequestBuilder("Limited Product", "Limited quantity product", 2500, "GHS").
					Unlimited(false).
					Quantity(10)
			},
			expectedName:      "Limited Product",
			expectedDesc:      "Limited quantity product",
			expectedPrice:     2500,
			expectedCurrency:  "GHS",
			expectedUnlimited: boolPtr(false),
			expectedQuantity:  intPtr(10),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			builder := tt.setupBuilder()
			req := builder.Build()

			assert.Equal(t, tt.expectedName, req.Name)
			assert.Equal(t, tt.expectedDesc, req.Description)
			assert.Equal(t, tt.expectedPrice, req.Price)
			assert.Equal(t, tt.expectedCurrency, req.Currency)

			if tt.expectedUnlimited != nil {
				require.NotNil(t, req.Unlimited)
				assert.Equal(t, *tt.expectedUnlimited, *req.Unlimited)
			} else {
				assert.Nil(t, req.Unlimited)
			}

			if tt.expectedQuantity != nil {
				require.NotNil(t, req.Quantity)
				assert.Equal(t, *tt.expectedQuantity, *req.Quantity)
			} else {
				assert.Nil(t, req.Quantity)
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
				return NewCreateRequestBuilder("Simple Product", "Simple description", 1500, "NGN")
			},
			expectField: func(t *testing.T, unmarshaled map[string]any) {
				assert.Equal(t, "Simple Product", unmarshaled["name"])
				assert.Equal(t, "Simple description", unmarshaled["description"])
				assert.Equal(t, float64(1500), unmarshaled["price"])
				assert.Equal(t, "NGN", unmarshaled["currency"])
				assert.NotContains(t, unmarshaled, "unlimited")
				assert.NotContains(t, unmarshaled, "quantity")
				assert.NotContains(t, unmarshaled, "metadata")
			},
		},
		{
			name: "serializes complete request correctly",
			builder: func() *CreateRequestBuilder {
				metadata := types.NewMetadata(map[string]any{
					"category": "electronics",
					"brand":    "TechCorp",
				})
				return NewCreateRequestBuilder("Full Product", "Complete product description", 25000, "USD").
					Unlimited(false).
					Quantity(100).
					Metadata(&metadata)
			},
			expectField: func(t *testing.T, unmarshaled map[string]any) {
				assert.Equal(t, "Full Product", unmarshaled["name"])
				assert.Equal(t, "Complete product description", unmarshaled["description"])
				assert.Equal(t, float64(25000), unmarshaled["price"])
				assert.Equal(t, "USD", unmarshaled["currency"])
				assert.Equal(t, false, unmarshaled["unlimited"])
				assert.Equal(t, float64(100), unmarshaled["quantity"])

				metadata, ok := unmarshaled["metadata"].(map[string]any)
				require.True(t, ok, "metadata should be an object")
				assert.Equal(t, "electronics", metadata["category"])
				assert.Equal(t, "TechCorp", metadata["brand"])
			},
		},
		{
			name: "serializes unlimited product correctly",
			builder: func() *CreateRequestBuilder {
				return NewCreateRequestBuilder("Unlimited Product", "Digital service", 10000, "KES").
					Unlimited(true)
			},
			expectField: func(t *testing.T, unmarshaled map[string]any) {
				assert.Equal(t, "Unlimited Product", unmarshaled["name"])
				assert.Equal(t, "Digital service", unmarshaled["description"])
				assert.Equal(t, float64(10000), unmarshaled["price"])
				assert.Equal(t, "KES", unmarshaled["currency"])
				assert.Equal(t, true, unmarshaled["unlimited"])
				assert.NotContains(t, unmarshaled, "quantity") // quantity shouldn't be set for unlimited
			},
		},
		{
			name: "omits empty optional fields",
			builder: func() *CreateRequestBuilder {
				return NewCreateRequestBuilder("Basic Product", "Basic product", 800, "ZAR")
			},
			expectField: func(t *testing.T, unmarshaled map[string]any) {
				// Required fields should be present
				assert.Contains(t, unmarshaled, "name")
				assert.Contains(t, unmarshaled, "description")
				assert.Contains(t, unmarshaled, "price")
				assert.Contains(t, unmarshaled, "currency")

				// Optional fields should be omitted when empty
				assert.NotContains(t, unmarshaled, "unlimited")
				assert.NotContains(t, unmarshaled, "quantity")
				assert.NotContains(t, unmarshaled, "metadata")
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
func boolPtr(b bool) *bool {
	return &b
}

func intPtr(i int) *int {
	return &i
}

func stringPtr(s string) *string {
	return &s
}
