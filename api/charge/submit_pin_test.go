package charge

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSubmitPINResponse_JSONDeserialization(t *testing.T) {
	tests := []struct {
		name            string
		responseFile    string
		expectedStatus  bool
		expectedMessage string
	}{
		{
			name:            "successful submit PIN response",
			responseFile:    "submit_pin_200.json",
			expectedStatus:  true,
			expectedMessage: "Charge attempted",
		},
		{
			name:            "failed submit PIN response",
			responseFile:    "submit_pin_200_failed.json",
			expectedStatus:  true,
			expectedMessage: "Charge attempted",
		},
		{
			name:            "submit PIN with pending status",
			responseFile:    "submit_pin_200_pending.json",
			expectedStatus:  true,
			expectedMessage: "Charge attempted",
		},
		{
			name:            "submit PIN validation error",
			responseFile:    "submit_pin_400.json",
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
			var response SubmitPINResponse
			err = json.Unmarshal(responseData, &response)
			require.NoError(t, err, "failed to unmarshal JSON response")

			// Verify the response structure
			assert.Equal(t, tt.expectedStatus, response.Status, "status should match")
			assert.Equal(t, tt.expectedMessage, response.Message, "message should match")

			// Only verify data structure for successful responses
			if tt.expectedStatus {
				assert.NotNil(t, response.Data, "data should not be nil")
				assert.NotEmpty(t, response.Data.Reference, "reference should not be empty")
				assert.NotEmpty(t, response.Data.Status, "status should not be empty")
			}
		})
	}
}

func TestSubmitPINRequestBuilder(t *testing.T) {
	t.Run("builds request with PIN and reference", func(t *testing.T) {
		builder := NewSubmitPINRequestBuilder("1234", "ref_123456789")
		request := builder.Build()

		assert.Equal(t, "1234", request.PIN, "PIN should match")
		assert.Equal(t, "ref_123456789", request.Reference, "reference should match")
	})

	t.Run("builds request with empty values", func(t *testing.T) {
		builder := NewSubmitPINRequestBuilder("", "")
		request := builder.Build()

		assert.Equal(t, "", request.PIN, "PIN should be empty")
		assert.Equal(t, "", request.Reference, "reference should be empty")
	})
}

func TestSubmitPINRequest_JSONSerialization(t *testing.T) {
	t.Run("serializes request correctly", func(t *testing.T) {
		builder := NewSubmitPINRequestBuilder("1234", "ref_123456789")
		request := builder.Build()

		jsonData, err := json.Marshal(request)
		require.NoError(t, err, "should marshal to JSON without error")

		var unmarshaled map[string]interface{}
		err = json.Unmarshal(jsonData, &unmarshaled)
		require.NoError(t, err, "should unmarshal JSON without error")

		assert.Equal(t, "1234", unmarshaled["pin"], "PIN should match")
		assert.Equal(t, "ref_123456789", unmarshaled["reference"], "reference should match")
	})
}
