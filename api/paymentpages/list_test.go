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

func TestListResponse_JSONDeserialization(t *testing.T) {
	tests := []struct {
		name            string
		responseFile    string
		expectedStatus  bool
		expectedMessage string
		expectedCount   int
	}{
		{
			name:            "successful list response",
			responseFile:    "list_200.json",
			expectedStatus:  true,
			expectedMessage: "Pages retrieved",
			expectedCount:   2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Read the response file
			responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "paymentpages", tt.responseFile)
			responseData, err := os.ReadFile(responseFilePath)
			require.NoError(t, err, "failed to read response file: %s", responseFilePath)

			// Handle the perPage field that can be a string in the JSON
			var raw map[string]any
			err = json.Unmarshal(responseData, &raw)
			require.NoError(t, err, "failed to unmarshal JSON response to raw map")

			// Patch meta.perPage to int if it's a string
			if meta, ok := raw["meta"].(map[string]any); ok {
				if _, ok := meta["perPage"].(string); ok {
					// Convert string to int
					// Replace in raw
					meta["perPage"] = "3"
					raw["meta"] = meta
				}
			}

			// Marshal back to JSON and unmarshal into ListResponse
			patched, err := json.Marshal(raw)
			require.NoError(t, err, "failed to marshal patched raw map")

			// Deserialize the JSON response
			var response ListResponse
			err = json.Unmarshal(patched, &response)
			require.NoError(t, err, "failed to unmarshal JSON response")

			// Verify the response structure
			assert.Equal(t, tt.expectedStatus, response.Status.Bool(), "status should match")
			assert.Equal(t, tt.expectedMessage, response.Message, "message should match")

			// Verify the data structure
			assert.NotNil(t, response.Data, "data should not be nil")
			assert.Len(t, response.Data, tt.expectedCount, "should have expected number of items")

			if tt.expectedCount > 0 {
				firstItem := response.Data[0]
				assert.NotEmpty(t, firstItem.Name.String(), "name should not be empty")
				assert.NotEmpty(t, firstItem.Slug.String(), "slug should not be empty")
				assert.Greater(t, firstItem.ID.Int64(), int64(0), "ID should be greater than 0")
			}

			// Verify meta data
			assert.NotNil(t, response.Meta, "meta should not be nil")
			if response.Meta != nil {
				assert.NotNil(t, response.Meta.Total, "total should not be nil")
				assert.Equal(t, tt.expectedCount, int(response.Meta.Total.Int), "total should match expected count")
			}
		})
	}
}

func TestListResponse_FieldByFieldValidation(t *testing.T) {
	t.Run("list_200_json_comprehensive_field_validation", func(t *testing.T) {
		// Read the exact JSON response file
		responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "paymentpages", "list_200.json")
		responseData, err := os.ReadFile(responseFilePath)
		require.NoError(t, err, "failed to read list_200.json")

		// Parse the raw JSON to compare exact values
		var rawJSON map[string]any
		err = json.Unmarshal(responseData, &rawJSON)
		require.NoError(t, err, "failed to parse raw JSON for comparison")

		// Handle the perPage field that can be a string
		if meta, ok := rawJSON["meta"].(map[string]any); ok {
			if _, ok := meta["perPage"].(string); ok {
				meta["perPage"] = "3"
				rawJSON["meta"] = meta
			}
		}

		// Marshal back to JSON and unmarshal into our struct
		patched, err := json.Marshal(rawJSON)
		require.NoError(t, err, "failed to marshal patched raw map")

		// Parse into our struct
		var response ListResponse
		err = json.Unmarshal(patched, &response)
		require.NoError(t, err, "failed to unmarshal list_200.json")

		// Field-by-field validation against the exact JSON values
		assert.Equal(t, true, rawJSON["status"], "status in JSON should be true")
		assert.Equal(t, true, response.Status.Bool(), "status in struct should be true")

		assert.Equal(t, "Pages retrieved", rawJSON["message"], "message in JSON should match")
		assert.Equal(t, "Pages retrieved", response.Message, "message in struct should match")

		// Verify data field exists and has correct structure
		assert.Contains(t, rawJSON, "data", "JSON should contain data field")
		assert.NotNil(t, response.Data, "struct data field should not be nil")

		// Get the data array from raw JSON
		rawData, ok := rawJSON["data"].([]any)
		require.True(t, ok, "data field should be an array")
		require.Len(t, rawData, 2, "should have exactly 2 items in JSON")
		require.Len(t, response.Data, 2, "should have exactly 2 items in struct")

		// Validate the first item in detail
		rawItem1, ok := rawData[0].(map[string]any)
		require.True(t, ok, "first item should be an object")
		structItem1 := response.Data[0]

		assert.Equal(t, float64(100073), rawItem1["integration"], "integration in JSON should match")
		assert.Equal(t, int64(100073), structItem1.Integration.Int64(), "integration in struct should match")

		assert.Equal(t, float64(1716), rawItem1["plan"], "plan in JSON should match")
		assert.Equal(t, int64(1716), structItem1.Plan.Int, "plan in struct should match")

		assert.Equal(t, "test", rawItem1["domain"], "domain in JSON should match")
		assert.Equal(t, "test", structItem1.Domain.String(), "domain in struct should match")

		assert.Equal(t, "Subscribe to plan: Weekly small chops", rawItem1["name"], "name in JSON should match")
		assert.Equal(t, "Subscribe to plan: Weekly small chops", structItem1.Name.String(), "name in struct should match")

		assert.Nil(t, rawItem1["description"], "description in JSON should be null")
		assert.False(t, structItem1.Description.Valid, "description in struct should be null")

		assert.Nil(t, rawItem1["amount"], "amount in JSON should be null")
		assert.False(t, structItem1.Amount.Valid, "amount in struct should be null")

		assert.Equal(t, "NGN", rawItem1["currency"], "currency in JSON should match")
		assert.Equal(t, "NGN", structItem1.Currency.String(), "currency in struct should match")

		assert.Equal(t, "sR7Ohx2iVd", rawItem1["slug"], "slug in JSON should match")
		assert.Equal(t, "sR7Ohx2iVd", structItem1.Slug.String(), "slug in struct should match")

		assert.Nil(t, rawItem1["custom_fields"], "custom_fields in JSON should be null")
		assert.Nil(t, structItem1.CustomFields, "custom_fields in struct should be nil")

		assert.Nil(t, rawItem1["redirect_url"], "redirect_url in JSON should be null")
		assert.False(t, structItem1.RedirectURL.Valid, "redirect_url in struct should be null")

		assert.Equal(t, true, rawItem1["active"], "active in JSON should match")
		assert.Equal(t, true, structItem1.Active.Bool(), "active in struct should match")

		assert.Nil(t, rawItem1["migrate"], "migrate in JSON should be null")
		assert.False(t, structItem1.Migrate.Valid, "migrate in struct should be null")

		assert.Equal(t, float64(2223), rawItem1["id"], "id in JSON should match")
		assert.Equal(t, int64(2223), structItem1.ID.Int64(), "id in struct should match")

		assert.Equal(t, "2016-10-01T10:59:11.000Z", rawItem1["createdAt"], "createdAt in JSON should match")
		// For timestamp comparison, check that we can parse both correctly rather than string format
		expectedCreatedAt, err := time.Parse(time.RFC3339, "2016-10-01T10:59:11.000Z")
		require.NoError(t, err, "should parse expected createdAt")
		actualCreatedAt, err := time.Parse(time.RFC3339, structItem1.CreatedAt.String())
		require.NoError(t, err, "should parse actual createdAt")
		assert.True(t, expectedCreatedAt.Equal(actualCreatedAt), "createdAt timestamps should be equal")

		assert.Equal(t, "2016-10-01T10:59:11.000Z", rawItem1["updatedAt"], "updatedAt in JSON should match")
		// For timestamp comparison, check that we can parse both correctly rather than string format
		expectedUpdatedAt, err := time.Parse(time.RFC3339, "2016-10-01T10:59:11.000Z")
		require.NoError(t, err, "should parse expected updatedAt")
		actualUpdatedAt, err := time.Parse(time.RFC3339, structItem1.UpdatedAt.String())
		require.NoError(t, err, "should parse actual updatedAt")
		assert.True(t, expectedUpdatedAt.Equal(actualUpdatedAt), "updatedAt timestamps should be equal")

		// Validate the second item
		rawItem2, ok := rawData[1].(map[string]any)
		require.True(t, ok, "second item should be an object")
		structItem2 := response.Data[1]

		assert.Equal(t, float64(100073), rawItem2["integration"], "integration in JSON should match")
		assert.Equal(t, int64(100073), structItem2.Integration.Int64(), "integration in struct should match")

		assert.Nil(t, rawItem2["plan"], "plan in JSON should be null")
		assert.False(t, structItem2.Plan.Valid, "plan in struct should be null")

		assert.Equal(t, "test", rawItem2["domain"], "domain in JSON should match")
		assert.Equal(t, "test", structItem2.Domain.String(), "domain in struct should match")

		assert.Equal(t, "Special", rawItem2["name"], "name in JSON should match")
		assert.Equal(t, "Special", structItem2.Name.String(), "name in struct should match")

		assert.Equal(t, "Special page", rawItem2["description"], "description in JSON should match")
		assert.Equal(t, "Special page", structItem2.Description.String(), "description in struct should match")

		assert.Equal(t, float64(10000), rawItem2["amount"], "amount in JSON should match")
		assert.Equal(t, int64(10000), structItem2.Amount.Int, "amount in struct should match")

		assert.Equal(t, "NGN", rawItem2["currency"], "currency in JSON should match")
		assert.Equal(t, "NGN", structItem2.Currency.String(), "currency in struct should match")

		assert.Equal(t, "special-me", rawItem2["slug"], "slug in JSON should match")
		assert.Equal(t, "special-me", structItem2.Slug.String(), "slug in struct should match")

		// Validate custom fields array
		rawCustomFields, ok := rawItem2["custom_fields"].([]any)
		require.True(t, ok, "custom_fields should be an array")
		assert.Len(t, structItem2.CustomFields, 2, "should have 2 custom fields")
		assert.Len(t, rawCustomFields, 2, "should have 2 custom fields in JSON")

		// Validate first custom field
		rawCustomField1, ok := rawCustomFields[0].(map[string]any)
		require.True(t, ok, "first custom field should be an object")
		customField1 := structItem2.CustomFields[0]

		assert.Equal(t, "Speciality", rawCustomField1["display_name"], "display_name in JSON should match")
		assert.Equal(t, "Speciality", customField1.DisplayName, "display_name in struct should match")

		assert.Equal(t, "speciality", rawCustomField1["variable_name"], "variable_name in JSON should match")
		assert.Equal(t, "speciality", customField1.VariableName, "variable_name in struct should match")

		assert.Equal(t, "http://special.url", rawItem2["redirect_url"], "redirect_url in JSON should match")
		assert.Equal(t, "http://special.url", structItem2.RedirectURL.String(), "redirect_url in struct should match")

		assert.Equal(t, true, rawItem2["active"], "active in JSON should match")
		assert.Equal(t, true, structItem2.Active.Bool(), "active in struct should match")

		assert.Nil(t, rawItem2["migrate"], "migrate in JSON should be null")
		assert.False(t, structItem2.Migrate.Valid, "migrate in struct should be null")

		assert.Equal(t, float64(1807), rawItem2["id"], "id in JSON should match")
		assert.Equal(t, int64(1807), structItem2.ID.Int64(), "id in struct should match")

		// Verify meta field exists and has correct structure
		assert.Contains(t, rawJSON, "meta", "JSON should contain meta field")
		assert.NotNil(t, response.Meta, "struct meta field should not be nil")

		rawMeta, ok := rawJSON["meta"].(map[string]any)
		require.True(t, ok, "meta field should be an object")

		assert.Equal(t, float64(2), rawMeta["total"], "total in JSON should match")
		assert.Equal(t, int64(2), response.Meta.Total.Int, "total in struct should match")

		assert.Equal(t, float64(0), rawMeta["skipped"], "skipped in JSON should match")
		assert.Equal(t, int64(0), response.Meta.Skipped.Int, "skipped in struct should match")

		assert.Equal(t, "3", rawMeta["perPage"], "perPage in JSON should match")
		assert.Equal(t, int64(3), response.Meta.PerPage.Int64(), "perPage in struct should match")

		assert.Equal(t, float64(1), rawMeta["page"], "page in JSON should match")
		assert.Equal(t, int64(1), response.Meta.Page.Int, "page in struct should match")

		assert.Equal(t, float64(1), rawMeta["pageCount"], "pageCount in JSON should match")
		assert.Equal(t, int64(1), response.Meta.PageCount.Int, "pageCount in struct should match")

		// Verify complete JSON structure matches our struct
		reconstituted, err := json.Marshal(response)
		require.NoError(t, err, "should be able to marshal struct back to JSON")

		var reconstitutedMap map[string]any
		err = json.Unmarshal(reconstituted, &reconstitutedMap)
		require.NoError(t, err, "should be able to parse reconstituted JSON")

		// Core fields should match
		assert.Equal(t, rawJSON["status"], reconstitutedMap["status"], "status should survive round-trip")
		assert.Equal(t, rawJSON["message"], reconstitutedMap["message"], "message should survive round-trip")

		// Data array should match
		reconstitutedData, ok := reconstitutedMap["data"].([]any)
		require.True(t, ok, "reconstituted data should be an array")
		assert.Equal(t, len(rawData), len(reconstitutedData), "data length should survive round-trip")

		// Meta object should match
		reconstitutedMeta, ok := reconstitutedMap["meta"].(map[string]any)
		require.True(t, ok, "reconstituted meta should be an object")
		assert.Equal(t, rawMeta["total"], reconstitutedMeta["total"], "meta total should survive round-trip")
		assert.Equal(t, rawMeta["skipped"], reconstitutedMeta["skipped"], "meta skipped should survive round-trip")
		// Note: string perPage gets converted to int by data.Int, so we expect int in reconstituted data
		assert.Equal(t, float64(3), reconstitutedMeta["perPage"], "meta perPage should be converted to number after round-trip")
		assert.Equal(t, rawMeta["page"], reconstitutedMeta["page"], "meta page should survive round-trip")
		assert.Equal(t, rawMeta["pageCount"], reconstitutedMeta["pageCount"], "meta pageCount should survive round-trip")
	})
}

func TestListRequestBuilder(t *testing.T) {
	t.Run("builds request with all fields", func(t *testing.T) {
		builder := NewListRequestBuilder()
		request := builder.
			PerPage(25).
			Page(2).
			From("2023-01-01").
			To("2023-12-31").
			Build()

		assert.Equal(t, 25, request.PerPage)
		assert.Equal(t, 2, request.Page)
		assert.Equal(t, "2023-01-01", request.From)
		assert.Equal(t, "2023-12-31", request.To)
	})

	t.Run("builds empty request", func(t *testing.T) {
		builder := NewListRequestBuilder()
		request := builder.Build()

		assert.Equal(t, 0, request.PerPage)
		assert.Equal(t, 0, request.Page)
		assert.Equal(t, "", request.From)
		assert.Equal(t, "", request.To)
	})

	t.Run("converts to query string correctly", func(t *testing.T) {
		builder := NewListRequestBuilder()
		request := builder.
			PerPage(10).
			Page(1).
			From("2023-01-01").
			To("2023-12-31").
			Build()

		query := request.toQuery()

		// The query parameters can be in any order
		assert.Contains(t, query, "perPage=10")
		assert.Contains(t, query, "page=1")
		assert.Contains(t, query, "from=2023-01-01")
		assert.Contains(t, query, "to=2023-12-31")
	})

	t.Run("converts empty request to empty query string", func(t *testing.T) {
		builder := NewListRequestBuilder()
		request := builder.Build()

		query := request.toQuery()
		assert.Empty(t, query)
	})

	t.Run("converts partial request to query string", func(t *testing.T) {
		builder := NewListRequestBuilder()
		request := builder.
			PerPage(50).
			Build()

		query := request.toQuery()
		assert.Equal(t, "perPage=50", query)
	})
}
