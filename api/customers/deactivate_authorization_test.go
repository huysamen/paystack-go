package customers

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDeactivateAuthorizationResponse_JSONDeserialization(t *testing.T) {
	// Read the deactivate_authorization_200.json file
	responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "customers", "deactivate_authorization_200.json")
	responseData, err := os.ReadFile(responseFilePath)
	require.NoError(t, err, "failed to read deactivate_authorization_200.json")

	var response DeactivateAuthorizationResponse
	err = json.Unmarshal(responseData, &response)
	require.NoError(t, err, "failed to unmarshal deactivate authorization response")

	// Basic validations
	assert.True(t, response.Status.Bool())
	assert.Equal(t, "Authorization has been deactivated", response.Message)
	// Data should be nil or empty for this endpoint
}

func TestDeactivateAuthorizationRequestBuilder(t *testing.T) {
	t.Run("builds request with authorization code", func(t *testing.T) {
		builder := NewDeactivateAuthorizationRequestBuilder("AUTH_test123")
		req := builder.Build()

		assert.Equal(t, "AUTH_test123", req.AuthorizationCode)
	})

	t.Run("builds request with empty authorization code", func(t *testing.T) {
		builder := NewDeactivateAuthorizationRequestBuilder("")
		req := builder.Build()

		assert.Equal(t, "", req.AuthorizationCode)
	})
}

func TestDeactivateAuthorizationRequest_JSONSerialization(t *testing.T) {
	t.Run("serializes request correctly", func(t *testing.T) {
		builder := NewDeactivateAuthorizationRequestBuilder("AUTH_test123")
		req := builder.Build()

		jsonData, err := json.Marshal(req)
		require.NoError(t, err, "should marshal without error")

		var unmarshaled map[string]any
		err = json.Unmarshal(jsonData, &unmarshaled)
		require.NoError(t, err, "should unmarshal JSON without error")

		assert.Equal(t, "AUTH_test123", unmarshaled["authorization_code"], "authorization_code should match")
	})
}

func TestDeactivateAuthorizationResponse_FieldByFieldValidation(t *testing.T) {
	// Read the deactivate_authorization_200.json file
	responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "customers", "deactivate_authorization_200.json")
	responseData, err := os.ReadFile(responseFilePath)
	require.NoError(t, err, "failed to read deactivate_authorization_200.json")

	// Parse the raw JSON to get the original values
	var rawResponse map[string]any
	err = json.Unmarshal(responseData, &rawResponse)
	require.NoError(t, err, "failed to unmarshal raw JSON response")

	// Deserialize into the DeactivateAuthorizationResponse struct
	var response DeactivateAuthorizationResponse
	err = json.Unmarshal(responseData, &response)
	require.NoError(t, err, "failed to unmarshal into DeactivateAuthorizationResponse struct")

	// Validate top-level fields against the raw JSON
	assert.Equal(t, rawResponse["status"], response.Status.Bool(), "status field should match")
	assert.Equal(t, rawResponse["message"], response.Message, "message field should match")

	// Data field should not exist in this response
	_, hasData := rawResponse["data"]
	assert.False(t, hasData, "data field should not exist in deactivate authorization response")

	// Test round-trip serialization
	serialized, err := json.Marshal(response)
	require.NoError(t, err, "failed to marshal response back to JSON")

	var roundTripResponse DeactivateAuthorizationResponse
	err = json.Unmarshal(serialized, &roundTripResponse)
	require.NoError(t, err, "failed to unmarshal round-trip JSON")

	// Verify core fields survive round-trip
	assert.Equal(t, response.Status.Bool(), roundTripResponse.Status.Bool(), "status should survive round-trip")
	assert.Equal(t, response.Message, roundTripResponse.Message, "message should survive round-trip")
}
