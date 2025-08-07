package applepay

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRegisterDomainResponse_JSONDeserialization(t *testing.T) {
	tests := []struct {
		name            string
		responseFile    string
		expectedStatus  bool
		expectedMessage string
		expectedData    any
	}{
		{
			name:            "successful register domain response",
			responseFile:    "register_domain_200.json",
			expectedStatus:  true,
			expectedMessage: "Domain successfully registered on Apple Pay",
			expectedData:    nil, // The response doesn't include a data field
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Read the response file
			responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "applepay", tt.responseFile)
			responseData, err := os.ReadFile(responseFilePath)
			require.NoError(t, err, "failed to read response file: %s", responseFilePath) // Deserialize the JSON response
			var response RegisterDomainResponse
			err = json.Unmarshal(responseData, &response)
			require.NoError(t, err, "failed to unmarshal JSON response")

			// Verify the response structure
			assert.Equal(t, tt.expectedStatus, response.Status, "status should match")
			assert.Equal(t, tt.expectedMessage, response.Message, "message should match")

			// Verify the data field (should be nil/empty for register domain)
			assert.Equal(t, tt.expectedData, response.Data, "data should match expected value")
		})
	}
}

func TestRegisterDomainRequestBuilder(t *testing.T) {
	t.Run("builds request with domain name", func(t *testing.T) {
		domainName := "test.example.com"
		builder := NewRegisterDomainRequestBuilder(domainName)
		request := builder.Build()

		assert.Equal(t, domainName, request.DomainName, "domain name should match")
	})

	t.Run("builds request with empty domain name", func(t *testing.T) {
		domainName := ""
		builder := NewRegisterDomainRequestBuilder(domainName)
		request := builder.Build()

		assert.Equal(t, domainName, request.DomainName, "domain name should be empty")
	})

	t.Run("domain name is required in constructor", func(t *testing.T) {
		// This test verifies that the constructor requires a domain name parameter
		// The builder pattern enforces this at compile time
		domainName := "required.domain.com"
		builder := NewRegisterDomainRequestBuilder(domainName)
		request := builder.Build()

		assert.NotNil(t, request, "request should not be nil")
		assert.Equal(t, domainName, request.DomainName, "domain name should be set")
	})
}

func TestRegisterDomainRequest_JSONSerialization(t *testing.T) {
	t.Run("serializes request correctly", func(t *testing.T) {
		domainName := "test.example.com"
		builder := NewRegisterDomainRequestBuilder(domainName)
		request := builder.Build()

		jsonData, err := json.Marshal(request)
		require.NoError(t, err, "should marshal to JSON without error")

		var unmarshaled map[string]interface{}
		err = json.Unmarshal(jsonData, &unmarshaled)
		require.NoError(t, err, "should unmarshal JSON without error")

		assert.Equal(t, domainName, unmarshaled["domainName"], "domainName field should match")
	})
}
