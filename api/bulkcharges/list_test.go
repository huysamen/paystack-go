package bulkcharges

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

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
			assert.Equal(t, tt.expectedStatus, response.Status, "status should match")
			assert.Equal(t, tt.expectedMessage, response.Message, "message should match")

			// Verify the data structure
			assert.NotNil(t, response.Data, "data should not be nil")
			assert.Len(t, response.Data, tt.expectedCount, "should have expected number of items")

			if tt.expectedCount > 0 {
				firstItem := response.Data[0]
				assert.NotEmpty(t, firstItem.BatchCode, "batch code should not be empty")
				assert.Greater(t, firstItem.ID, 0, "ID should be greater than 0")
				assert.NotEmpty(t, firstItem.Status, "status should not be empty")
			}

			// Verify meta data
			assert.NotNil(t, response.Meta, "meta should not be nil")
			if response.Meta != nil {
				assert.NotNil(t, response.Meta.Total, "total should not be nil")
				assert.Equal(t, tt.expectedCount, *response.Meta.Total, "total should match expected count")
			}
		})
	}
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
