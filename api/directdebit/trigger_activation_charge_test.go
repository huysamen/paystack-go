package directdebit

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTriggerActivationChargeResponse_JSONDeserialization(t *testing.T) {
	tests := []struct {
		name            string
		responseFile    string
		expectedStatus  bool
		expectedMessage string
	}{
		{
			name:            "successful trigger activation charge response",
			responseFile:    "trigger_activation_charge_200.json",
			expectedStatus:  true,
			expectedMessage: "Mandate is queued for retry",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Read the response file
			responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "directdebit", tt.responseFile)
			responseData, err := os.ReadFile(responseFilePath)
			require.NoError(t, err, "failed to read response file: %s", responseFilePath)

			// Deserialize the JSON response
			var response TriggerActivationChargeResponse
			err = json.Unmarshal(responseData, &response)
			require.NoError(t, err, "failed to unmarshal JSON response")

			// Verify the response structure
			assert.Equal(t, tt.expectedStatus, response.Status.Bool(), "status should match")
			assert.Equal(t, tt.expectedMessage, response.Message, "message should match")
		})
	}
}

func TestTriggerActivationChargeRequestBuilder(t *testing.T) {
	t.Run("builds request with single customer ID", func(t *testing.T) {
		builder := NewTriggerActivationChargeRequestBuilder().
			CustomerID(12345)

		request := builder.Build()

		assert.Len(t, request.CustomerIDs, 1, "should have one customer ID")
		assert.Equal(t, uint64(12345), request.CustomerIDs[0], "customer ID should match")
	})

	t.Run("builds request with multiple customer IDs", func(t *testing.T) {
		builder := NewTriggerActivationChargeRequestBuilder().
			CustomerID(12345).
			CustomerID(67890)

		request := builder.Build()

		assert.Len(t, request.CustomerIDs, 2, "should have two customer IDs")
		assert.Equal(t, uint64(12345), request.CustomerIDs[0], "first customer ID should match")
		assert.Equal(t, uint64(67890), request.CustomerIDs[1], "second customer ID should match")
	})

	t.Run("builds request with customer IDs array", func(t *testing.T) {
		customerIDs := []uint64{11111, 22222, 33333}
		builder := NewTriggerActivationChargeRequestBuilder().
			CustomerIDs(customerIDs)

		request := builder.Build()

		assert.Len(t, request.CustomerIDs, 3, "should have three customer IDs")
		assert.Equal(t, customerIDs, request.CustomerIDs, "customer IDs should match")
	})

	t.Run("builds request with empty values", func(t *testing.T) {
		builder := NewTriggerActivationChargeRequestBuilder()
		request := builder.Build()

		assert.Empty(t, request.CustomerIDs, "customer IDs should be empty")
	})

	t.Run("builds request combining methods", func(t *testing.T) {
		builder := NewTriggerActivationChargeRequestBuilder().
			CustomerIDs([]uint64{11111, 22222}).
			CustomerID(33333)

		request := builder.Build()

		assert.Len(t, request.CustomerIDs, 3, "should have three customer IDs")
		assert.Equal(t, uint64(11111), request.CustomerIDs[0], "first customer ID should match")
		assert.Equal(t, uint64(22222), request.CustomerIDs[1], "second customer ID should match")
		assert.Equal(t, uint64(33333), request.CustomerIDs[2], "third customer ID should match")
	})
}

func TestTriggerActivationChargeRequest_JSONSerialization(t *testing.T) {
	t.Run("serializes request correctly", func(t *testing.T) {
		builder := NewTriggerActivationChargeRequestBuilder().
			CustomerIDs([]uint64{12345, 67890})

		request := builder.Build()

		jsonData, err := json.Marshal(request)
		require.NoError(t, err, "should marshal to JSON without error")

		var unmarshaled map[string]any
		err = json.Unmarshal(jsonData, &unmarshaled)
		require.NoError(t, err, "should unmarshal JSON without error")

		customerIDs := unmarshaled["customer_ids"].([]any)
		assert.Len(t, customerIDs, 2, "should have two customer IDs in JSON")
		assert.Equal(t, float64(12345), customerIDs[0], "first customer ID should match")
		assert.Equal(t, float64(67890), customerIDs[1], "second customer ID should match")
	})

	t.Run("serializes empty request", func(t *testing.T) {
		builder := NewTriggerActivationChargeRequestBuilder()
		request := builder.Build()

		jsonData, err := json.Marshal(request)
		require.NoError(t, err, "should marshal to JSON without error")

		var unmarshaled map[string]any
		err = json.Unmarshal(jsonData, &unmarshaled)
		require.NoError(t, err, "should unmarshal JSON without error")

		customerIDs := unmarshaled["customer_ids"]
		assert.Nil(t, customerIDs, "customer_ids should be nil in JSON for empty array")
	})
}

func TestTriggerActivationChargeResponse_FieldByFieldValidation(t *testing.T) {
	// Read the response file
	responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "directdebit", "trigger_activation_charge_200.json")
	responseData, err := os.ReadFile(responseFilePath)
	require.NoError(t, err, "failed to read response file")

	// Parse as raw JSON to get expected values
	var rawData map[string]any
	err = json.Unmarshal(responseData, &rawData)
	require.NoError(t, err, "failed to unmarshal raw JSON")

	// Deserialize into struct
	var response TriggerActivationChargeResponse
	err = json.Unmarshal(responseData, &response)
	require.NoError(t, err, "failed to unmarshal into struct")

	// Validate top-level fields
	expectedStatus := rawData["status"].(bool)
	assert.Equal(t, expectedStatus, response.Status.Bool(), "status should match")
	assert.Equal(t, rawData["message"], response.Message, "message should match")

	// Note: This endpoint returns a simple response with no data object,
	// so we just validate the basic structure

	// Test round-trip serialization
	serialized, err := json.Marshal(response)
	require.NoError(t, err, "should marshal back to JSON without error")

	var roundTripResponse TriggerActivationChargeResponse
	err = json.Unmarshal(serialized, &roundTripResponse)
	require.NoError(t, err, "should unmarshal round-trip JSON without error")

	// Verify round-trip integrity
	assert.Equal(t, response.Status.Bool(), roundTripResponse.Status.Bool(), "round-trip status should match")
	assert.Equal(t, response.Message, roundTripResponse.Message, "round-trip message should match")
}
