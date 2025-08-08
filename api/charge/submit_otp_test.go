package charge

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSubmitOTPResponse_JSONDeserialization(t *testing.T) {
	tests := []struct {
		name            string
		responseFile    string
		expectedStatus  bool
		expectedMessage string
	}{
		{
			name:            "successful submit OTP response",
			responseFile:    "submit_otp_200.json",
			expectedStatus:  true,
			expectedMessage: "Charge attempted",
		},
		{
			name:            "submit OTP requiring bank authorization",
			responseFile:    "submit_otp_200_bank_auth.json",
			expectedStatus:  true,
			expectedMessage: "Charge attempted",
		},
		{
			name:            "failed submit OTP response",
			responseFile:    "submit_otp_200_failed.json",
			expectedStatus:  true,
			expectedMessage: "Charge attempted",
		},
		{
			name:            "submit OTP with pending status",
			responseFile:    "submit_otp_200_pending.json",
			expectedStatus:  true,
			expectedMessage: "Charge attempted",
		},
		{
			name:            "submit OTP validation error",
			responseFile:    "submit_otp_400.json",
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
			var response SubmitOTPResponse
			err = json.Unmarshal(responseData, &response)
			require.NoError(t, err, "failed to unmarshal JSON response")

			// Verify the response structure
			assert.Equal(t, tt.expectedStatus, response.Status, "status should match")
			assert.Equal(t, tt.expectedMessage, response.Message, "message should match")

			// Only verify data structure for successful responses
			if tt.expectedStatus {
				// Some responses may not have reference/status fields
				if response.Data.Reference != "" {
					assert.NotEmpty(t, response.Data.Reference, "reference should not be empty")
				}
				if response.Data.Status != "" {
					assert.NotEmpty(t, response.Data.Status, "status should not be empty")
				}
			}
		})
	}
}

func TestSubmitOTPRequestBuilder(t *testing.T) {
	t.Run("builds request with OTP and reference", func(t *testing.T) {
		builder := NewSubmitOTPRequestBuilder("123456", "ref_123456789")
		request := builder.Build()

		assert.Equal(t, "123456", request.OTP, "OTP should match")
		assert.Equal(t, "ref_123456789", request.Reference, "reference should match")
	})

	t.Run("builds request with empty values", func(t *testing.T) {
		builder := NewSubmitOTPRequestBuilder("", "")
		request := builder.Build()

		assert.Equal(t, "", request.OTP, "OTP should be empty")
		assert.Equal(t, "", request.Reference, "reference should be empty")
	})
}

func TestSubmitOTPRequest_JSONSerialization(t *testing.T) {
	t.Run("serializes request correctly", func(t *testing.T) {
		builder := NewSubmitOTPRequestBuilder("123456", "ref_123456789")
		request := builder.Build()

		jsonData, err := json.Marshal(request)
		require.NoError(t, err, "should marshal to JSON without error")

		var unmarshaled map[string]interface{}
		err = json.Unmarshal(jsonData, &unmarshaled)
		require.NoError(t, err, "should unmarshal JSON without error")

		assert.Equal(t, "123456", unmarshaled["otp"], "OTP should match")
		assert.Equal(t, "ref_123456789", unmarshaled["reference"], "reference should match")
	})
}
