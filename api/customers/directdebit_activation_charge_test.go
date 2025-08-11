package customers

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDirectDebitActivationChargeResponse_JSONDeserialization(t *testing.T) {
	// Read the directdebit_activation_charge_200.json file
	responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "customers", "directdebit_activation_charge_200.json")
	responseData, err := os.ReadFile(responseFilePath)
	require.NoError(t, err, "failed to read directdebit_activation_charge_200.json")

	var response DirectDebitActivationChargeResponse
	err = json.Unmarshal(responseData, &response)
	require.NoError(t, err, "failed to unmarshal direct debit activation charge response")

	// Basic validations
	assert.True(t, response.Status.Bool())
	assert.Equal(t, "Mandate is queued for retry", response.Message)
	// Data should be nil or empty for this endpoint
}

func TestDirectDebitActivationChargeRequestBuilder(t *testing.T) {
	t.Run("builds request with authorization ID", func(t *testing.T) {
		builder := NewDirectDebitActivationChargeRequestBuilder(12345)
		req := builder.Build()

		assert.Equal(t, 12345, req.AuthorizationID)
	})

	t.Run("builds request with zero authorization ID", func(t *testing.T) {
		builder := NewDirectDebitActivationChargeRequestBuilder(0)
		req := builder.Build()

		assert.Equal(t, 0, req.AuthorizationID)
	})

	t.Run("builds request with negative authorization ID", func(t *testing.T) {
		builder := NewDirectDebitActivationChargeRequestBuilder(-1)
		req := builder.Build()

		assert.Equal(t, -1, req.AuthorizationID)
	})
}

func TestDirectDebitActivationChargeRequest_JSONSerialization(t *testing.T) {
	t.Run("serializes request correctly", func(t *testing.T) {
		builder := NewDirectDebitActivationChargeRequestBuilder(54321)
		req := builder.Build()

		jsonData, err := json.Marshal(req)
		require.NoError(t, err, "should marshal without error")

		var unmarshaled map[string]any
		err = json.Unmarshal(jsonData, &unmarshaled)
		require.NoError(t, err, "should unmarshal JSON without error")

		// JSON numbers are unmarshaled as float64 in Go
		assert.Equal(t, float64(54321), unmarshaled["authorization_id"], "authorization_id should match")
	})

	t.Run("serializes zero authorization ID correctly", func(t *testing.T) {
		builder := NewDirectDebitActivationChargeRequestBuilder(0)
		req := builder.Build()

		jsonData, err := json.Marshal(req)
		require.NoError(t, err, "should marshal without error")

		var unmarshaled map[string]any
		err = json.Unmarshal(jsonData, &unmarshaled)
		require.NoError(t, err, "should unmarshal JSON without error")

		assert.Equal(t, float64(0), unmarshaled["authorization_id"], "authorization_id should be 0")
	})
}

func TestDirectDebitActivationChargeResponse_FieldByFieldValidation(t *testing.T) {
	// Read the directdebit_activation_charge_200.json file
	responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "customers", "directdebit_activation_charge_200.json")
	responseData, err := os.ReadFile(responseFilePath)
	require.NoError(t, err, "failed to read directdebit_activation_charge_200.json")

	// Parse the raw JSON to get the original values
	var rawResponse map[string]any
	err = json.Unmarshal(responseData, &rawResponse)
	require.NoError(t, err, "failed to unmarshal raw JSON response")

	// Deserialize into the DirectDebitActivationChargeResponse struct
	var response DirectDebitActivationChargeResponse
	err = json.Unmarshal(responseData, &response)
	require.NoError(t, err, "failed to unmarshal into DirectDebitActivationChargeResponse struct")

	// Validate top-level fields against the raw JSON
	assert.Equal(t, rawResponse["status"], response.Status.Bool(), "status field should match")
	assert.Equal(t, rawResponse["message"], response.Message, "message field should match")

	// Data field should not exist in this response
	_, hasData := rawResponse["data"]
	assert.False(t, hasData, "data field should not exist in direct debit activation charge response")

	// Test round-trip serialization
	serialized, err := json.Marshal(response)
	require.NoError(t, err, "failed to marshal response back to JSON")

	var roundTripResponse DirectDebitActivationChargeResponse
	err = json.Unmarshal(serialized, &roundTripResponse)
	require.NoError(t, err, "failed to unmarshal round-trip JSON")

	// Verify core fields survive round-trip
	assert.Equal(t, response.Status.Bool(), roundTripResponse.Status.Bool(), "status should survive round-trip")
	assert.Equal(t, response.Message, roundTripResponse.Message, "message should survive round-trip")
}
