package applepay

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListDomainsResponse_JSONDeserialization(t *testing.T) {
	tests := []struct {
		name            string
		responseFile    string
		expectedStatus  bool
		expectedMessage string
		expectedDomains []string
	}{
		{
			name:            "successful list domains response",
			responseFile:    "list_domains_200.json",
			expectedStatus:  true,
			expectedMessage: "Apple Pay registered domains retrieved",
			expectedDomains: []string{"example.com"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Read the response file
			responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "applepay", tt.responseFile)
			responseData, err := os.ReadFile(responseFilePath)
			require.NoError(t, err, "failed to read response file: %s", responseFilePath)

			// Deserialize the JSON response
			var response ListDomainsResponse
			err = json.Unmarshal(responseData, &response)
			require.NoError(t, err, "failed to unmarshal JSON response")

			// Verify the response structure
			assert.Equal(t, tt.expectedStatus, response.Status.Bool(), "status should match")
			assert.Equal(t, tt.expectedMessage, response.Message, "message should match")

			// Verify the data structure
			assert.NotNil(t, response.Data, "data should not be nil")
			assert.Equal(t, tt.expectedDomains, response.Data.DomainNames, "domain names should match")
		})
	}
}

func TestListDomainsRequestBuilder(t *testing.T) {
	t.Run("builds request with all fields", func(t *testing.T) {
		builder := NewListDomainsRequest()
		request := builder.
			UseCursor(true).
			Next("next_cursor").
			Previous("prev_cursor").
			Build()

		assert.NotNil(t, request.UseCursor)
		assert.True(t, *request.UseCursor)
		assert.NotNil(t, request.Next)
		assert.Equal(t, "next_cursor", *request.Next)
		assert.NotNil(t, request.Previous)
		assert.Equal(t, "prev_cursor", *request.Previous)
	})

	t.Run("builds empty request", func(t *testing.T) {
		builder := NewListDomainsRequest()
		request := builder.Build()

		assert.Nil(t, request.UseCursor)
		assert.Nil(t, request.Next)
		assert.Nil(t, request.Previous)
	})

	t.Run("converts to query string correctly", func(t *testing.T) {
		builder := NewListDomainsRequest()
		request := builder.
			UseCursor(true).
			Next("next_cursor").
			Previous("prev_cursor").
			Build()

		query := request.toQuery()

		// The query parameters can be in any order
		assert.Contains(t, query, "use_cursor=true")
		assert.Contains(t, query, "next=next_cursor")
		assert.Contains(t, query, "previous=prev_cursor")
	})

	t.Run("converts empty request to empty query string", func(t *testing.T) {
		builder := NewListDomainsRequest()
		request := builder.Build()

		query := request.toQuery()
		assert.Empty(t, query)
	})
}

func TestListDomainsResponseData_Structure(t *testing.T) {
	t.Run("response data has correct field types", func(t *testing.T) {
		data := ListDomainsResponseData{
			DomainNames: []string{"test.com", "example.org"},
		}

		assert.IsType(t, []string{}, data.DomainNames)
		assert.Len(t, data.DomainNames, 2)
		assert.Equal(t, "test.com", data.DomainNames[0])
		assert.Equal(t, "example.org", data.DomainNames[1])
	})

	t.Run("response data handles empty domain list", func(t *testing.T) {
		data := ListDomainsResponseData{
			DomainNames: []string{},
		}

		assert.IsType(t, []string{}, data.DomainNames)
		assert.Empty(t, data.DomainNames)
	})
}
