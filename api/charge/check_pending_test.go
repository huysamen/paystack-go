package charge

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCheckPendingResponse_JSONDeserialization(t *testing.T) {
	tests := []struct {
		name            string
		responseFile    string
		expectedStatus  bool
		expectedMessage string
	}{
		{
			name:            "successful check pending response",
			responseFile:    "check_pending_200.json",
			expectedStatus:  true,
			expectedMessage: "Charge attempted",
		},
		{
			name:            "check pending requiring OTP",
			responseFile:    "check_pending_200_otp.json",
			expectedStatus:  true,
			expectedMessage: "Charge attempted",
		},
		{
			name:            "check pending requiring PIN",
			responseFile:    "check_pensing_200_pin.json", // Note: typo in filename from API
			expectedStatus:  true,
			expectedMessage: "Charge attempted",
		},
		{
			name:            "check pending with birthday verification",
			responseFile:    "check_pending_200_birthday.json",
			expectedStatus:  true,
			expectedMessage: "Charge attempted",
		},
		{
			name:            "check pending with phone verification",
			responseFile:    "check_pending_200_phone.json",
			expectedStatus:  true,
			expectedMessage: "Charge attempted",
		},
		{
			name:            "check pending requiring bank authorization",
			responseFile:    "check_pending_200_bank_auth.json",
			expectedStatus:  true,
			expectedMessage: "Charge attempted",
		},
		{
			name:            "check pending with pending status",
			responseFile:    "check_pending_200_pending.json",
			expectedStatus:  true,
			expectedMessage: "Reference check successful",
		},
		{
			name:            "failed check pending response",
			responseFile:    "check_pending_200_failed.json",
			expectedStatus:  true,
			expectedMessage: "Reference check successful",
		},
		{
			name:            "check pending validation error",
			responseFile:    "check_pending_400.json",
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
			var response CheckPendingResponse
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

func TestCheckPendingRequestBuilder(t *testing.T) {
	t.Run("builds request with reference", func(t *testing.T) {
		builder := NewCheckPendingChargeRequestBuilder("ref_123456789")
		request := builder.Build()

		assert.Equal(t, "ref_123456789", request.Reference, "reference should match")
	})

	t.Run("builds request with empty reference", func(t *testing.T) {
		builder := NewCheckPendingChargeRequestBuilder("")
		request := builder.Build()

		assert.Equal(t, "", request.Reference, "reference should be empty")
	})
}

func TestCheckPendingRequest_JSONSerialization(t *testing.T) {
	t.Run("serializes request correctly", func(t *testing.T) {
		builder := NewCheckPendingChargeRequestBuilder("ref_123456789")
		request := builder.Build()

		jsonData, err := json.Marshal(request)
		require.NoError(t, err, "should marshal to JSON without error")

		var unmarshaled map[string]interface{}
		err = json.Unmarshal(jsonData, &unmarshaled)
		require.NoError(t, err, "should unmarshal JSON without error")

		assert.Equal(t, "ref_123456789", unmarshaled["reference"], "reference should match")
	})
}
