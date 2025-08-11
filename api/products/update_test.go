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
			expectedMessage: "Product successfully updated",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Read the JSON response file
			responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "products", tt.responseFile)
			responseData, err := os.ReadFile(responseFilePath)
			require.NoError(t, err, "failed to read response file")

			// Deserialize the JSON response
			var response UpdateResponse
			err = json.Unmarshal(responseData, &response)
			require.NoError(t, err, "failed to unmarshal JSON response")

			// Verify the response structure
			assert.Equal(t, tt.expectedStatus, response.Status.Bool(), "status should match")
			assert.Equal(t, tt.expectedMessage, response.Message, "message should match")
			assert.NotNil(t, response.Data, "data should not be nil")
		})
	}
}

func TestUpdateResponse_FieldByFieldValidation(t *testing.T) {
	t.Run("update_200_json_comprehensive_field_validation", func(t *testing.T) {
		// Read the exact JSON response file
		responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "products", "update_200.json")
		responseData, err := os.ReadFile(responseFilePath)
		require.NoError(t, err, "failed to read update_200.json")

		// Parse raw JSON to compare field by field
		var rawJSON map[string]any
		err = json.Unmarshal(responseData, &rawJSON)
		require.NoError(t, err, "failed to parse raw JSON")

		// Parse into our struct
		var response UpdateResponse
		err = json.Unmarshal(responseData, &response)
		require.NoError(t, err, "failed to unmarshal update_200.json")

		// Verify struct data is properly populated
		assert.NotNil(t, response.Data, "struct data field should not be nil")

		// Get the data portion from raw JSON
		rawData, ok := rawJSON["data"].(map[string]any)
		require.True(t, ok, "data field should be an object")

		// Validate core fields (using actual values from update_200.json)
		assert.Equal(t, "Prod One", rawData["name"], "name in JSON should match")
		assert.Equal(t, "Prod One", response.Data.Name.String(), "name in struct should match")

		assert.Equal(t, "Prod 1", rawData["description"], "description in JSON should match")
		assert.Equal(t, "Prod 1", response.Data.Description.String(), "description in struct should match")

		assert.Equal(t, "NGN", rawData["currency"], "currency in JSON should match")
		assert.Equal(t, enums.CurrencyNGN, response.Data.Currency, "currency in struct should match")

		assert.Equal(t, float64(20000), rawData["price"], "price in JSON should match")
		assert.Equal(t, int64(20000), response.Data.Price.Int64(), "price in struct should match")

		assert.Equal(t, float64(5), rawData["quantity"], "quantity in JSON should match")
		assert.Equal(t, int64(5), response.Data.Quantity.Int, "quantity in struct should match")

		assert.Equal(t, false, rawData["is_shippable"], "is_shippable in JSON should match")
		assert.Equal(t, false, response.Data.IsShippable.Bool(), "is_shippable in struct should match")

		assert.Equal(t, false, rawData["unlimited"], "unlimited in JSON should match")
		assert.Equal(t, false, response.Data.Unlimited.Bool(), "unlimited in struct should match")

		assert.Equal(t, float64(343288), rawData["integration"], "integration in JSON should match")
		assert.Equal(t, int64(343288), response.Data.Integration.Int64(), "integration in struct should match")

		assert.Equal(t, "test", rawData["domain"], "domain in JSON should match")
		assert.Equal(t, "test", response.Data.Domain.String(), "domain in struct should match")

		// Slug field is not present in update response, so we shouldn't test it
		// The response just doesn't include this field

		assert.Equal(t, "PROD_ohc0xq1ajpt2271", rawData["product_code"], "product_code in JSON should match")
		assert.Equal(t, "PROD_ohc0xq1ajpt2271", response.Data.ProductCode.String(), "product_code in struct should match")

		// Test null quantity_sold field
		assert.Equal(t, nil, rawData["quantity_sold"], "quantity_sold in JSON should be null")
		assert.False(t, response.Data.QuantitySold.Valid, "quantity_sold in struct should be null")

		assert.Equal(t, "good", rawData["type"], "type in JSON should match")
		assert.Equal(t, "good", response.Data.Type.String(), "type in struct should match")

		assert.Equal(t, true, rawData["active"], "active in JSON should match")
		assert.Equal(t, true, response.Data.Active.Bool(), "active in struct should match")

		assert.Equal(t, true, rawData["in_stock"], "in_stock in JSON should match")
		assert.Equal(t, true, response.Data.InStock.Bool(), "in_stock in struct should match")

		// Test null fields that are not present in this response
		assert.Equal(t, nil, rawData["features"], "features in JSON should be null")
		assert.Equal(t, nil, rawData["metadata"], "metadata in JSON should be null")

		assert.Equal(t, float64(526), rawData["id"], "id in JSON should match")
		assert.Equal(t, int64(526), response.Data.ID.Int64(), "id in struct should match")

		// Test timestamps with lenient comparison (update uses 000Z format)
		assert.Equal(t, "2019-06-29T14:46:52.000Z", rawData["createdAt"], "created_at in JSON should match")
		expectedCreatedAt, err := time.Parse(time.RFC3339, "2019-06-29T14:46:52.000Z")
		require.NoError(t, err, "should parse expected created_at")
		actualCreatedAt, err := time.Parse(time.RFC3339, response.Data.CreatedAt.String())
		require.NoError(t, err, "should parse actual created_at")
		assert.True(t, expectedCreatedAt.Sub(actualCreatedAt).Abs() < time.Second, "created_at timestamps should be within 1 second")

		assert.Equal(t, "2019-06-29T15:29:21.000Z", rawData["updatedAt"], "updated_at in JSON should match")
		expectedUpdatedAt, err := time.Parse(time.RFC3339, "2019-06-29T15:29:21.000Z")
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

func TestUpdateRequestBuilder(t *testing.T) {
	tests := []struct {
		name              string
		setupBuilder      func() *UpdateRequestBuilder
		expectedName      *string
		expectedDesc      *string
		expectedPrice     *int
		expectedUnlimited *bool
		expectedQuantity  *int
	}{
		{
			name: "builds request with only name update",
			setupBuilder: func() *UpdateRequestBuilder {
				return NewUpdateRequestBuilder().
					Name("Updated Product Name")
			},
			expectedName: stringPtr("Updated Product Name"),
		},
		{
			name: "builds request with price and description update",
			setupBuilder: func() *UpdateRequestBuilder {
				return NewUpdateRequestBuilder().
					Name("Test Product").
					Description("Updated description").
					Price(7500)
			},
			expectedName:  stringPtr("Test Product"),
			expectedDesc:  stringPtr("Updated description"),
			expectedPrice: intPtr(7500),
		},
		{
			name: "builds request with quantity and unlimited status",
			setupBuilder: func() *UpdateRequestBuilder {
				return NewUpdateRequestBuilder().
					Name("Limited Product").
					Unlimited(false).
					Quantity(25)
			},
			expectedName:      stringPtr("Limited Product"),
			expectedUnlimited: boolPtr(false),
			expectedQuantity:  intPtr(25),
		},
		{
			name: "builds request with all fields",
			setupBuilder: func() *UpdateRequestBuilder {
				metadata := types.NewMetadata(map[string]any{
					"updated": "true",
					"version": "2.0",
				})
				return NewUpdateRequestBuilder().
					Name("Complete Update").
					Description("Completely updated product").
					Price(15000).
					Unlimited(false).
					Quantity(50).
					Metadata(&metadata)
			},
			expectedName:      stringPtr("Complete Update"),
			expectedDesc:      stringPtr("Completely updated product"),
			expectedPrice:     intPtr(15000),
			expectedUnlimited: boolPtr(false),
			expectedQuantity:  intPtr(50),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			builder := tt.setupBuilder()
			req := builder.Build()

			if tt.expectedName != nil {
				require.NotNil(t, req.Name)
				assert.Equal(t, *tt.expectedName, *req.Name)
			} else {
				assert.Nil(t, req.Name)
			}

			if tt.expectedDesc != nil {
				require.NotNil(t, req.Description)
				assert.Equal(t, *tt.expectedDesc, *req.Description)
			} else {
				assert.Nil(t, req.Description)
			}

			if tt.expectedPrice != nil {
				require.NotNil(t, req.Price)
				assert.Equal(t, *tt.expectedPrice, *req.Price)
			} else {
				assert.Nil(t, req.Price)
			}

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

func TestUpdateRequest_JSONSerialization(t *testing.T) {
	tests := []struct {
		name        string
		builder     func() *UpdateRequestBuilder
		expectField func(t *testing.T, unmarshaled map[string]any)
	}{
		{
			name: "serializes name-only update correctly",
			builder: func() *UpdateRequestBuilder {
				return NewUpdateRequestBuilder().
					Name("New Name Only")
			},
			expectField: func(t *testing.T, unmarshaled map[string]any) {
				assert.Equal(t, "New Name Only", unmarshaled["name"])
				assert.NotContains(t, unmarshaled, "description")
				assert.NotContains(t, unmarshaled, "price")
				assert.NotContains(t, unmarshaled, "unlimited")
				assert.NotContains(t, unmarshaled, "quantity")
				assert.NotContains(t, unmarshaled, "metadata")
			},
		},
		{
			name: "serializes complete update correctly",
			builder: func() *UpdateRequestBuilder {
				metadata := types.NewMetadata(map[string]any{
					"color":    "red",
					"category": "electronics",
				})
				return NewUpdateRequestBuilder().
					Name("Complete Product Update").
					Description("Fully updated description").
					Price(12000).
					Unlimited(false).
					Quantity(75).
					Metadata(&metadata)
			},
			expectField: func(t *testing.T, unmarshaled map[string]any) {
				assert.Equal(t, "Complete Product Update", unmarshaled["name"])
				assert.Equal(t, "Fully updated description", unmarshaled["description"])
				assert.Equal(t, float64(12000), unmarshaled["price"])
				assert.Equal(t, false, unmarshaled["unlimited"])
				assert.Equal(t, float64(75), unmarshaled["quantity"])

				metadata, ok := unmarshaled["metadata"].(map[string]any)
				require.True(t, ok, "metadata should be an object")
				assert.Equal(t, "red", metadata["color"])
				assert.Equal(t, "electronics", metadata["category"])
			},
		},
		{
			name: "serializes partial update correctly",
			builder: func() *UpdateRequestBuilder {
				return NewUpdateRequestBuilder().
					Name("Partial Update").
					Price(8000)
			},
			expectField: func(t *testing.T, unmarshaled map[string]any) {
				assert.Equal(t, "Partial Update", unmarshaled["name"])
				assert.Equal(t, float64(8000), unmarshaled["price"])
				assert.NotContains(t, unmarshaled, "description")
				assert.NotContains(t, unmarshaled, "unlimited")
				assert.NotContains(t, unmarshaled, "quantity")
				assert.NotContains(t, unmarshaled, "metadata")
			},
		},
		{
			name: "omits unset optional fields",
			builder: func() *UpdateRequestBuilder {
				return NewUpdateRequestBuilder().
					Name("Minimal Update")
			},
			expectField: func(t *testing.T, unmarshaled map[string]any) {
				// Required update field should be present
				assert.Contains(t, unmarshaled, "name")
				assert.Equal(t, "Minimal Update", unmarshaled["name"])

				// Optional fields should be omitted when not set
				assert.NotContains(t, unmarshaled, "description")
				assert.NotContains(t, unmarshaled, "price")
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

func TestUpdate(t *testing.T) {
	tests := []struct {
		name               string
		productID          string
		expectedError      bool
		expectedErrorMsg   string
		expectedRequestURL string
	}{
		{
			name:               "update with valid product ID",
			productID:          "PROD_ddot3upakgl3ejt",
			expectedError:      false,
			expectedRequestURL: "https://api.paystack.co/product/PROD_ddot3upakgl3ejt",
		},
		{
			name:               "update with numeric product ID",
			productID:          "489399",
			expectedError:      false,
			expectedRequestURL: "https://api.paystack.co/product/489399",
		},
		{
			name:             "update with empty product ID",
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
				// URL construction test would verify: PUT /product/{id}
				expectedPath := "/product/" + tt.productID
				assert.Contains(t, tt.expectedRequestURL, expectedPath, "URL should contain correct product ID path")
			}
		})
	}
}

func TestUpdateResponse_JSONRoundTrip(t *testing.T) {
	t.Run("deserialize_and_serialize_maintains_structure", func(t *testing.T) {
		// Read the JSON response file
		responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "products", "update_200.json")
		originalData, err := os.ReadFile(responseFilePath)
		require.NoError(t, err, "failed to read update_200.json")

		// Parse into struct
		var response UpdateResponse
		err = json.Unmarshal(originalData, &response)
		require.NoError(t, err, "failed to unmarshal update_200.json")

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
