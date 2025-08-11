package bulkcharges

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPauseResponse_JSONDeserialization(t *testing.T) {
	tests := []struct {
		name            string
		responseFile    string
		expectedStatus  bool
		expectedMessage string
	}{
		{
			name:            "successful pause response",
			responseFile:    "pause_200.json",
			expectedStatus:  true,
			expectedMessage: "Bulk charge batch has been paused",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Read the response file
			responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "bulkcharges", tt.responseFile)
			responseData, err := os.ReadFile(responseFilePath)
			require.NoError(t, err, "failed to read response file: %s", responseFilePath)

			// Deserialize the JSON response
			var response PauseResponse
			err = json.Unmarshal(responseData, &response)
			require.NoError(t, err, "failed to unmarshal JSON response")

			// Verify the response structure
			assert.Equal(t, tt.expectedStatus, response.Status.Bool(), "status should match")
			assert.Equal(t, tt.expectedMessage, response.Message, "message should match")

			// Data should be nil for pause response
			assert.Nil(t, response.Data, "data should be nil for pause response")
		})
	}
}

func TestPauseResponse_FieldByFieldValidation(t *testing.T) {
	// Read the pause_200.json file
	responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "bulkcharges", "pause_200.json")
	responseData, err := os.ReadFile(responseFilePath)
	require.NoError(t, err, "failed to read pause_200.json")

	// Parse the raw JSON to get the original values
	var rawResponse map[string]any
	err = json.Unmarshal(responseData, &rawResponse)
	require.NoError(t, err, "failed to unmarshal raw JSON response")

	// Deserialize into the PauseResponse struct
	var response PauseResponse
	err = json.Unmarshal(responseData, &response)
	require.NoError(t, err, "failed to unmarshal into PauseResponse struct")

	// Validate each field against the raw JSON
	assert.Equal(t, rawResponse["status"], response.Status.Bool(), "status field should match")
	assert.Equal(t, rawResponse["message"], response.Message, "message field should match")

	// Data should be nil (not present in JSON)
	assert.Nil(t, response.Data, "data field should be nil")

	// Test round-trip serialization
	serialized, err := json.Marshal(response)
	require.NoError(t, err, "failed to marshal response back to JSON")

	var roundTripResponse PauseResponse
	err = json.Unmarshal(serialized, &roundTripResponse)
	require.NoError(t, err, "failed to unmarshal round-trip JSON")

	// Verify round-trip consistency
	assert.Equal(t, response.Status.Bool(), roundTripResponse.Status.Bool(), "status should survive round-trip")
	assert.Equal(t, response.Message, roundTripResponse.Message, "message should survive round-trip")
	assert.Nil(t, roundTripResponse.Data, "data should remain nil after round-trip")
}
