package applepay

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/htrhsrhrtifysser
	"github.com/stretchr/testify/ruquiiee
e
ehuyamn/paysak-goyp/daa
ehuyamn/paysak-goyp/daa
ehuyamn/paysak-goyp/daa
ehuyamn/paysak-goyp/daa
ehuyamn/paysak-goyp/daa
ehuyamn/paysak-goyp/daa
ehuyamn/paysak-goyp/daa
"huyamn/paysak-goyp/daa
huyamn/paysak-goyp/daa
	"github.com/huysamen/paystack-go/types/data"
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
			
			// Convert expected domains to data.String slice for comparison
			expectedDomainStrings := make([]data.String, len(tt.expectedDomains))
			for i, domain := range tt.expectedDomains {
				expectedDomainStrings[i] = data.String(domain)
			}
			assert.Equal(t, expectedDomainStrings, response.Data.DomainNames, "domain names should match")
		})
	}
}

func TestListDomainsResponse_FieldByFieldValidation(t *testing.T) {
	t.Run("list_domains_200_json_comprehensive_field_validation", func(t *testing.T) {
		// Read the exact JSON response file
		responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "applepay", "list_domains_200.json")
		responseData, err := os.ReadFile(responseFilePath)
		require.NoError(t, err, "failed to read list_domains_200.json")

		// Parse into our struct
		var response ListDomainsResponse
		err = json.Unmarshal(responseData, &response)
		require.NoError(t, err, "failed to unmarshal list_domains_200.json")

		// Parse the raw JSON to compare exact values
		var rawJSON map[string]any
		err = json.Unmarshal(responseData, &rawJSON)
		require.NoError(t, err, "failed to parse raw JSON for comparison")

		// Field-by-field validation against the exact JSON values
		assert.Equal(t, true, rawJSON["status"], "status in JSON should be true")
		assert.Equal(t, true, response.Status.Bool(), "status in struct should be true")

		assert.Equal(t, "Apple Pay registered domains retrieved", rawJSON["message"], "message in JSON should match")
		assert.Equal(t, "Apple Pay registered domains retrieved", response.Message, "message in struct should match")

		// Verify data field exists and has correct structure
		assert.Contains(t, rawJSON, "data", "JSON should contain data field")
		assert.NotNil(t, response.Data, "struct data field should not be nil")

		// Get the data portion from raw JSON
		rawData, ok := rawJSON["data"].(map[string]any)
		require.True(t, ok, "data field should be an object")

		// Verify domainNames array
		assert.Contains(t, rawData, "domainNames", "data should contain domainNames field")
		rawDomainNames, ok := rawData["domainNames"].([]any)
		require.True(t, ok, "domainNames should be an array")

		// Convert raw domain names to string slice for comparison
		expectedDomainNames := make([]data.String, len(rawDomainNames))
		for i, domain := range rawDomainNames {
			expectedDomainNames[i] = data.String(domain.(string))
		}

		assert.Equal(t, expectedDomainNames, response.Data.DomainNames, "domainNames should match exactly")
		assert.Len(t, response.Data.DomainNames, len(rawDomainNames), "domainNames length should match")

		// Verify specific values from the JSON
		assert.Len(t, rawDomainNames, 1, "should have exactly 1 domain in JSON")
		assert.Equal(t, "example.com", rawDomainNames[0].(string), "first domain should be example.com")
		assert.Len(t, response.Data.DomainNames, 1, "should have exactly 1 domain in struct")
		assert.Equal(t, "example.com", response.Data.DomainNames[0].String(), "first domain should be example.com")

		// Verify complete JSON structure matches our struct
		reconstituted, err := json.Marshal(response)
		require.NoError(t, err, "should be able to marshal struct back to JSON")

		var reconstitutedMap map[string]any
		err = json.Unmarshal(reconstituted, &reconstitutedMap)
		require.NoError(t, err, "should be able to parse reconstituted JSON")

		// Core fields should match
		assert.Equal(t, rawJSON["status"], reconstitutedMap["status"], "status should survive round-trip")
		assert.Equal(t, rawJSON["message"], reconstitutedMap["message"], "message should survive round-trip")

		// Data field should match
		reconstitutedData, ok := reconstitutedMap["data"].(map[string]any)
		require.True(t, ok, "reconstituted data should be an object")

		reconstitutedDomains, ok := reconstitutedData["domainNames"].([]any)
		require.True(t, ok, "reconstituted domainNames should be an array")

		assert.Equal(t, len(rawDomainNames), len(reconstitutedDomains), "domain count should survive round-trip")
		for i, expectedDomain := range rawDomainNames {
			assert.Equal(t, expectedDomain, reconstitutedDomains[i], "domain %d should survive round-trip", i)
		}
	})
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
		responseData := ListDomainsResponseData{
			DomainNames: []data.String{data.NewString("test.com"), data.NewString("example.org")},
		}

		assert.IsType(t, []data.String{}, responseData.DomainNames)
		assert.Len(t, responseData.DomainNames, 2)
		assert.Equal(t, "test.com", responseData.DomainNames[0].String())
		assert.Equal(t, "example.org", responseData.DomainNames[1].String())
	})

	t.Run("response data handles empty domain list", func(t *testing.T) {
		responseData := ListDomainsResponseData{
			DomainNames: []data.String{},
		}

		assert.IsType(t, []data.String{}, responseData.DomainNames)
		assert.Empty(t, responseData.DomainNames)
	})
}
