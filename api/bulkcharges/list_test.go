package bulkcharges

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
			expectedMessage: "Bulk charges retrieved",
			expectedCount:   1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Read the response file
			responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "bulkcharges", tt.responseFile)
			responseData, err := os.ReadFile(responseFilePath)
			require.NoError(t, err, "failed to read response file: %s", responseFilePath)

			// Deserialize the JSON response
			var response ListResponse
			err = json.Unmarshal(responseData, &response)
			require.NoError(t, err, "failed to unmarshal JSON response")

			// Verify the response structure
			assert.Equal(t, tt.expectedStatus, response.Status.Bool(), "status should match")
			assert.Equal(t, tt.expectedMessage, response.Message, "message should match")

			// Verify the data structure
			assert.NotNil(t, response.Data, "data should not be nil")
			assert.Len(t, response.Data, tt.expectedCount, "should have expected number of items")

			if tt.expectedCount > 0 {
				firstItem := response.Data[0]
				assert.NotEmpty(t, firstItem.BatchCode.String(), "batch code should not be empty")
				assert.Greater(t, firstItem.ID.Int64(), int64(0), "ID should be greater than 0")
				assert.NotEmpty(t, firstItem.Status.String(), "status should not be empty")
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
		responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "bulkcharges", "list_200.json")
		responseData, err := os.ReadFile(responseFilePath)
		require.NoError(t, err, "failed to read list_200.json")

		// Parse into our struct
		var response ListResponse
		err = json.Unmarshal(responseData, &response)
		require.NoError(t, err, "failed to unmarshal list_200.json")

		// Parse the raw JSON to compare exact values
		var rawJSON map[string]any
		err = json.Unmarshal(responseData, &rawJSON)
		require.NoError(t, err, "failed to parse raw JSON for comparison")

		// Field-by-field validation against the exact JSON values
		assert.Equal(t, true, rawJSON["status"], "status in JSON should be true")
		assert.Equal(t, true, response.Status.Bool(), "status in struct should be true")

		assert.Equal(t, "Bulk charges retrieved", rawJSON["message"], "message in JSON should match")
		assert.Equal(t, "Bulk charges retrieved", response.Message, "message in struct should match")

		// Verify data field exists and has correct structure
		assert.Contains(t, rawJSON, "data", "JSON should contain data field")
		assert.NotNil(t, response.Data, "struct data field should not be nil")

		// Get the data array from raw JSON
		rawData, ok := rawJSON["data"].([]any)
		require.True(t, ok, "data field should be an array")
		require.Len(t, rawData, 1, "should have exactly 1 item in JSON")
		require.Len(t, response.Data, 1, "should have exactly 1 item in struct")

		// Validate the first (and only) item in detail
		rawItem, ok := rawData[0].(map[string]any)
		require.True(t, ok, "first item should be an object")
		structItem := response.Data[0]

		assert.Equal(t, "test", rawItem["domain"], "domain in JSON should match")
		assert.Equal(t, "test", structItem.Domain.String(), "domain in struct should match")

		assert.Equal(t, "BCH_1nV4L1D7cayggh", rawItem["batch_code"], "batch_code in JSON should match")
		assert.Equal(t, "BCH_1nV4L1D7cayggh", structItem.BatchCode.String(), "batch_code in struct should match")

		assert.Equal(t, "complete", rawItem["status"], "status in JSON should match")
		assert.Equal(t, "complete", structItem.Status.String(), "status in struct should match")

		assert.Equal(t, float64(1733), rawItem["id"], "id in JSON should match")
		assert.Equal(t, int64(1733), structItem.ID.Int64(), "id in struct should match")

		assert.Equal(t, "2017-02-04T05:44:19.000Z", rawItem["createdAt"], "createdAt in JSON should match")
		// For timestamp comparison, check that we can parse both correctly rather than string format
		expectedCreatedAt, err := time.Parse(time.RFC3339, "2017-02-04T05:44:19.000Z")
		require.NoError(t, err, "should parse expected createdAt")
		actualCreatedAt, err := time.Parse(time.RFC3339, structItem.CreatedAt.String())
		require.NoError(t, err, "should parse actual createdAt")
		assert.True(t, expectedCreatedAt.Equal(actualCreatedAt), "createdAt timestamps should be equal")

		assert.Equal(t, "2017-02-04T05:45:02.000Z", rawItem["updatedAt"], "updatedAt in JSON should match")
		// For timestamp comparison, check that we can parse both correctly rather than string format
		expectedUpdatedAt, err := time.Parse(time.RFC3339, "2017-02-04T05:45:02.000Z")
		require.NoError(t, err, "should parse expected updatedAt")
		actualUpdatedAt, err := time.Parse(time.RFC3339, structItem.UpdatedAt.String())
		require.NoError(t, err, "should parse actual updatedAt")
		assert.True(t, expectedUpdatedAt.Equal(actualUpdatedAt), "updatedAt timestamps should be equal")

		// Verify meta field exists and has correct structure
		assert.Contains(t, rawJSON, "meta", "JSON should contain meta field")
		assert.NotNil(t, response.Meta, "struct meta field should not be nil")

		rawMeta, ok := rawJSON["meta"].(map[string]any)
		require.True(t, ok, "meta field should be an object")

		assert.Equal(t, float64(1), rawMeta["total"], "total in JSON should match")
		assert.Equal(t, int64(1), response.Meta.Total.Int, "total in struct should match")

		assert.Equal(t, float64(0), rawMeta["skipped"], "skipped in JSON should match")
		assert.Equal(t, int64(0), response.Meta.Skipped.Int, "skipped in struct should match")

		assert.Equal(t, float64(50), rawMeta["perPage"], "perPage in JSON should match")
		assert.Equal(t, int64(50), int64(response.Meta.PerPage), "perPage in struct should match")

		assert.Equal(t, float64(1), rawMeta["page"], "page in JSON should match")
		assert.Equal(t, int64(1), response.Meta.Page.Int, "page in struct should match")

		assert.Equal(t, float64(1), rawMeta["pageCount"], "pageCount in JSON should match")
		assert.Equal(t, int64(1), response.Meta.PageCount.Int, "pageCount in struct should match")

		assert.Equal(t, float64(0), rawMeta["skipped"], "skipped in JSON should match")
		assert.NotNil(t, response.Meta.Skipped, "skipped in struct should not be nil")
		assert.Equal(t, int64(0), response.Meta.Skipped.Int, "skipped in struct should match")

		assert.Equal(t, float64(50), rawMeta["perPage"], "perPage in JSON should match")
		assert.Equal(t, int64(50), response.Meta.PerPage.Int64(), "perPage in struct should match")

		assert.Equal(t, float64(1), rawMeta["page"], "page in JSON should match")
		assert.NotNil(t, response.Meta.Page, "page in struct should not be nil")
		assert.Equal(t, int64(1), response.Meta.Page.Int, "page in struct should match")

		assert.Equal(t, float64(1), rawMeta["pageCount"], "pageCount in JSON should match")
		assert.NotNil(t, response.Meta.PageCount, "pageCount in struct should not be nil")
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
		assert.Equal(t, rawMeta["perPage"], reconstitutedMeta["perPage"], "meta perPage should survive round-trip")
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

		assert.NotNil(t, request.PerPage)
		assert.Equal(t, 25, *request.PerPage)
		assert.NotNil(t, request.Page)
		assert.Equal(t, 2, *request.Page)
		assert.NotNil(t, request.From)
		assert.Equal(t, "2023-01-01", *request.From)
		assert.NotNil(t, request.To)
		assert.Equal(t, "2023-12-31", *request.To)
	})

	t.Run("builds request with date range helper", func(t *testing.T) {
		builder := NewListRequestBuilder()
		request := builder.
			DateRange("2023-06-01", "2023-06-30").
			Build()

		assert.NotNil(t, request.From)
		assert.Equal(t, "2023-06-01", *request.From)
		assert.NotNil(t, request.To)
		assert.Equal(t, "2023-06-30", *request.To)
	})

	t.Run("builds empty request", func(t *testing.T) {
		builder := NewListRequestBuilder()
		request := builder.Build()

		assert.Nil(t, request.PerPage)
		assert.Nil(t, request.Page)
		assert.Nil(t, request.From)
		assert.Nil(t, request.To)
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

func TestListRequestDateRange(t *testing.T) {
	t.Run("date range overwrites individual from/to", func(t *testing.T) {
		builder := NewListRequestBuilder()
		request := builder.
			From("2023-01-01").
			To("2023-01-31").
			DateRange("2023-06-01", "2023-06-30").
			Build()

		assert.Equal(t, "2023-06-01", *request.From)
		assert.Equal(t, "2023-06-30", *request.To)
	})

	t.Run("individual from/to overwrites date range", func(t *testing.T) {
		builder := NewListRequestBuilder()
		request := builder.
			DateRange("2023-06-01", "2023-06-30").
			From("2023-01-01").
			To("2023-01-31").
			Build()

		assert.Equal(t, "2023-01-01", *request.From)
		assert.Equal(t, "2023-01-31", *request.To)
	})
}
