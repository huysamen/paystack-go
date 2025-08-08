package integration

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFetchTimeoutResponse_JSONDeserialization(t *testing.T) {
	tests := []struct {
		name         string
		responseFile string
	}{
		{
			name:         "successful fetch timeout response",
			responseFile: "fetch_200.json",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Read the response file
			responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "integration", tt.responseFile)
			responseData, err := os.ReadFile(responseFilePath)
			require.NoError(t, err, "failed to read response file")

			// Deserialize the JSON response
			var response FetchTimeoutResponse
			err = json.Unmarshal(responseData, &response)
			require.NoError(t, err, "failed to unmarshal JSON response")

			// Basic response validation
			assert.True(t, response.Status.Bool(), "status should be true")
			assert.Equal(t, "Payment session timeout retrieved", response.Message, "message should match")
			require.NotNil(t, response.Data, "data should not be nil")
		})
	}
}

func TestFetchTimeoutResponse_FieldByFieldValidation(t *testing.T) {
	// Read the response file
	responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "integration", "fetch_200.json")
	responseData, err := os.ReadFile(responseFilePath)
	require.NoError(t, err, "failed to read response file")

	// Deserialize the JSON response
	var response FetchTimeoutResponse
	err = json.Unmarshal(responseData, &response)
	require.NoError(t, err, "failed to unmarshal JSON response")

	// Validate response structure
	assert.True(t, response.Status.Bool())
	assert.Equal(t, "Payment session timeout retrieved", response.Message)
	require.NotNil(t, response.Data)

	// Validate timeout data
	data := response.Data
	assert.Equal(t, 30, data.PaymentSessionTimeout, "payment session timeout should match")

	// Test JSON round-trip
	marshaled, err := json.Marshal(response)
	require.NoError(t, err)

	var roundTrip FetchTimeoutResponse
	err = json.Unmarshal(marshaled, &roundTrip)
	require.NoError(t, err)

	assert.Equal(t, response.Status, roundTrip.Status)
	assert.Equal(t, response.Message, roundTrip.Message)
	assert.Equal(t, response.Data.PaymentSessionTimeout, roundTrip.Data.PaymentSessionTimeout)
}
