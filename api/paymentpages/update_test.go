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

func TestUpdateResponse_JSONDeserialization(t *testing.T) {
	tests := []struct {
		name            string
		responseFile    string
		expectedStatus  bool
		expectedMessage string
		expectedName    string
	}{
		{
			name:            "successful update response",
			responseFile:    "update_200.json",
			expectedStatus:  true,
			expectedMessage: "Page updated",
			expectedName:    "Buttercup Brunch",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Read the response file
			responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "paymentpages", tt.responseFile)
			responseData, err := os.ReadFile(responseFilePath)
			require.NoError(t, err, "failed to read response file: %s", responseFilePath)

			// Deserialize the JSON response
			var response UpdateResponse
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

func TestUpdateResponse_FieldByFieldValidation(t *testing.T) {
	t.Run("update_200_json_comprehensive_field_validation", func(t *testing.T) {
		// Read the exact JSON response file
		responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "paymentpages", "update_200.json")
		responseData, err := os.ReadFile(responseFilePath)
		require.NoError(t, err, "failed to read update_200.json")

		// Parse into our struct
		var response UpdateResponse
		err = json.Unmarshal(responseData, &response)
		require.NoError(t, err, "failed to unmarshal update_200.json")

		// Parse the raw JSON to compare exact values
		var rawJSON map[string]any
		err = json.Unmarshal(responseData, &rawJSON)
		require.NoError(t, err, "failed to parse raw JSON for comparison")

		// Field-by-field validation against the exact JSON values
		assert.Equal(t, true, rawJSON["status"], "status in JSON should be true")
		assert.Equal(t, true, response.Status.Bool(), "status in struct should be true")

		assert.Equal(t, "Page updated", rawJSON["message"], "message in JSON should match")
		assert.Equal(t, "Page updated", response.Message, "message in struct should match")

		// Verify data field exists and has correct structure
		assert.Contains(t, rawJSON, "data", "JSON should contain data field")
		assert.NotNil(t, response.Data, "struct data field should not be nil")

		// Get the data portion from raw JSON
		rawData, ok := rawJSON["data"].(map[string]any)
		require.True(t, ok, "data field should be an object")

		// Validate key data fields
		assert.Equal(t, float64(100032), rawData["integration"], "integration in JSON should match")
		assert.Equal(t, int64(100032), response.Data.Integration.Int64(), "integration in struct should match")

		assert.Equal(t, "test", rawData["domain"], "domain in JSON should match")
		assert.Equal(t, "test", response.Data.Domain.String(), "domain in struct should match")

		assert.Equal(t, "Buttercup Brunch", rawData["name"], "name in JSON should match")
		assert.Equal(t, "Buttercup Brunch", response.Data.Name.String(), "name in struct should match")

		assert.Equal(t, "Gather your friends for the ritual that is brunch", rawData["description"], "description in JSON should match")
		assert.Equal(t, "Gather your friends for the ritual that is brunch", response.Data.Description.String(), "description in struct should match")

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

		assert.Equal(t, "2016-03-30T04:44:35.000Z", rawData["updatedAt"], "updatedAt in JSON should match")
		// For timestamp comparison, check that we can parse both correctly rather than string format
		expectedUpdatedAt, err := time.Parse(time.RFC3339, "2016-03-30T04:44:35.000Z")
		require.NoError(t, err, "should parse expected updatedAt")
		actualUpdatedAt, err := time.Parse(time.RFC3339, response.Data.UpdatedAt.String())
		require.NoError(t, err, "should parse actual updatedAt")
		assert.True(t, expectedUpdatedAt.Equal(actualUpdatedAt), "updatedAt timestamps should be equal")

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
	t.Run("builds empty request", func(t *testing.T) {
		builder := NewUpdateRequestBuilder()
		request := builder.Build()

		assert.Equal(t, "", request.Name)
		assert.Equal(t, "", request.Description)
		assert.Nil(t, request.Amount)
		assert.Nil(t, request.Active)
	})

	t.Run("builds request with all fields", func(t *testing.T) {
		builder := NewUpdateRequestBuilder()
		request := builder.
			Name("Updated Page").
			Description("Updated Description").
			Amount(15000).
			Active(true).
			Build()

		assert.Equal(t, "Updated Page", request.Name)
		assert.Equal(t, "Updated Description", request.Description)
		assert.Equal(t, 15000, *request.Amount)
		assert.Equal(t, true, *request.Active)
	})
}

func TestUpdateRequest_JSONSerialization(t *testing.T) {
	t.Run("serializes request correctly", func(t *testing.T) {
		builder := NewUpdateRequestBuilder().
			Name("Updated Page").
			Description("Updated Description").
			Amount(15000).
			Active(false)

		request := builder.Build()

		jsonData, err := json.Marshal(request)
		require.NoError(t, err, "should marshal to JSON without error")

		var unmarshaled map[string]any
		err = json.Unmarshal(jsonData, &unmarshaled)
		require.NoError(t, err, "should unmarshal JSON without error")

		assert.Equal(t, "Updated Page", unmarshaled["name"])
		assert.Equal(t, "Updated Description", unmarshaled["description"])
		assert.Equal(t, float64(15000), unmarshaled["amount"])
		assert.Equal(t, false, unmarshaled["active"])
	})
}
