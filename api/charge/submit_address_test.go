package charge

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSubmitAddressResponse_JSONDeserialization(t *testing.T) {
	tests := []struct {
		name            string
		responseFile    string
		expectedStatus  bool
		expectedMessage string
	}{
		{
			name:            "successful submit address response",
			responseFile:    "submit_address_200.json",
			expectedStatus:  true,
			expectedMessage: "Charge attempted",
		},
		{
			name:            "failed submit address response",
			responseFile:    "submit_address_200_failed.json",
			expectedStatus:  false,
			expectedMessage: "Charge attempted",
		},
		{
			name:            "submit address validation error",
			responseFile:    "submit_address_400.json",
			expectedStatus:  false,
			expectedMessage: "Transaction reference is invalid",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Read the response file
			responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "charge", tt.responseFile)
			responseData, err := os.ReadFile(responseFilePath)
			require.NoError(t, err, "failed to read response file: %s", responseFilePath)

			// Deserialize the JSON response
			var response SubmitAddressResponse
			err = json.Unmarshal(responseData, &response)
			require.NoError(t, err, "failed to unmarshal JSON response")

			// Verify the response structure
			assert.Equal(t, tt.expectedStatus, response.Status.Bool(), "status should match")
			assert.Equal(t, tt.expectedMessage, response.Message, "message should match")

			// Only verify data structure for successful responses
			if tt.expectedStatus {
				assert.NotEmpty(t, response.Data.Reference, "reference should not be empty")
				assert.NotEmpty(t, response.Data.Status, "status should not be empty")
			}
		})
	}
}

func TestSubmitAddressRequestBuilder(t *testing.T) {
	t.Run("builds request with all address components", func(t *testing.T) {
		builder := NewSubmitAddressRequestBuilder(
			"123 Main Street",
			"Lagos",
			"Lagos",
			"101001",
			"ref_123456789",
		)
		request := builder.Build()

		assert.Equal(t, "123 Main Street", request.Address, "address should match")
		assert.Equal(t, "Lagos", request.City, "city should match")
		assert.Equal(t, "Lagos", request.State, "state should match")
		assert.Equal(t, "101001", request.ZipCode, "zipcode should match")
		assert.Equal(t, "ref_123456789", request.Reference, "reference should match")
	})

	t.Run("builds request with empty values", func(t *testing.T) {
		builder := NewSubmitAddressRequestBuilder("", "", "", "", "")
		request := builder.Build()

		assert.Equal(t, "", request.Address, "address should be empty")
		assert.Equal(t, "", request.City, "city should be empty")
		assert.Equal(t, "", request.State, "state should be empty")
		assert.Equal(t, "", request.ZipCode, "zipcode should be empty")
		assert.Equal(t, "", request.Reference, "reference should be empty")
	})
}

func TestSubmitAddressRequest_JSONSerialization(t *testing.T) {
	t.Run("serializes request correctly", func(t *testing.T) {
		builder := NewSubmitAddressRequestBuilder(
			"123 Main Street",
			"Lagos",
			"Lagos",
			"101001",
			"ref_123456789",
		)
		request := builder.Build()

		jsonData, err := json.Marshal(request)
		require.NoError(t, err, "should marshal to JSON without error")

		var unmarshaled map[string]interface{}
		err = json.Unmarshal(jsonData, &unmarshaled)
		require.NoError(t, err, "should unmarshal JSON without error")

		assert.Equal(t, "123 Main Street", unmarshaled["address"], "address should match")
		assert.Equal(t, "Lagos", unmarshaled["city"], "city should match")
		assert.Equal(t, "Lagos", unmarshaled["state"], "state should match")
		assert.Equal(t, "101001", unmarshaled["zipcode"], "zipcode should match")
		assert.Equal(t, "ref_123456789", unmarshaled["reference"], "reference should match")
	})
}
