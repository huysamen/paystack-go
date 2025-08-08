package charge

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSubmitPhoneResponse_JSONDeserialization(t *testing.T) {
	tests := []struct {
		name            string
		responseFile    string
		expectedStatus  bool
		expectedMessage string
	}{
		{
			name:            "successful submit phone response",
			responseFile:    "submit_phone_200.json",
			expectedStatus:  true,
			expectedMessage: "Charge attempted",
		},
		{
			name:            "failed submit phone response",
			responseFile:    "submit_phone_200_failed.json",
			expectedStatus:  true,
			expectedMessage: "Charge attempted",
		},
		{
			name:            "submit phone with pending status",
			responseFile:    "submit_phone_200_pending.json",
			expectedStatus:  true,
			expectedMessage: "Charge attempted",
		},
		{
			name:            "submit phone validation error",
			responseFile:    "submit_phone_400.json",
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
			var response SubmitPhoneResponse
			err = json.Unmarshal(responseData, &response)
			require.NoError(t, err, "failed to unmarshal JSON response")

			// Verify the response structure
			assert.Equal(t, tt.expectedStatus, response.Status.Bool(), "status should match")
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

func TestSubmitPhoneRequestBuilder(t *testing.T) {
	t.Run("builds request with phone and reference", func(t *testing.T) {
		builder := NewSubmitPhoneRequestBuilder("+2348123456789", "ref_123456789")
		request := builder.Build()

		assert.Equal(t, "+2348123456789", request.Phone, "phone should match")
		assert.Equal(t, "ref_123456789", request.Reference, "reference should match")
	})

	t.Run("builds request with empty values", func(t *testing.T) {
		builder := NewSubmitPhoneRequestBuilder("", "")
		request := builder.Build()

		assert.Equal(t, "", request.Phone, "phone should be empty")
		assert.Equal(t, "", request.Reference, "reference should be empty")
	})
}

func TestSubmitPhoneRequest_JSONSerialization(t *testing.T) {
	t.Run("serializes request correctly", func(t *testing.T) {
		builder := NewSubmitPhoneRequestBuilder("+2348123456789", "ref_123456789")
		request := builder.Build()

		jsonData, err := json.Marshal(request)
		require.NoError(t, err, "should marshal to JSON without error")

		var unmarshaled map[string]interface{}
		err = json.Unmarshal(jsonData, &unmarshaled)
		require.NoError(t, err, "should unmarshal JSON without error")

		assert.Equal(t, "+2348123456789", unmarshaled["phone"], "phone should match")
		assert.Equal(t, "ref_123456789", unmarshaled["reference"], "reference should match")
	})
}
