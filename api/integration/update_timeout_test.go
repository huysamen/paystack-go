package integration

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUpdateTimeoutResponse_JSONDeserialization(t *testing.T) {
	tests := []struct {
		name         string
		responseFile string
	}{
		{
			name:         "successful update timeout response",
			responseFile: "update_200.json",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Read the response file
			responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "integration", tt.responseFile)
			responseData, err := os.ReadFile(responseFilePath)
			require.NoError(t, err, "failed to read response file")

			// Deserialize the JSON response
			var response UpdateTimeoutResponse
			err = json.Unmarshal(responseData, &response)
			require.NoError(t, err, "failed to unmarshal JSON response")

			// Basic response validation
			assert.True(t, response.Status.Bool(), "status should be true")
			assert.Equal(t, "Payment session timeout updated", response.Message, "message should match")
			require.NotNil(t, response.Data, "data should not be nil")
		})
	}
}

func TestUpdateTimeoutRequestBuilder(t *testing.T) {
	tests := []struct {
		name            string
		timeout         int
		expectedTimeout int
	}{
		{
			name:            "builds request with 30 second timeout",
			timeout:         30,
			expectedTimeout: 30,
		},
		{
			name:            "builds request with 60 second timeout",
			timeout:         60,
			expectedTimeout: 60,
		},
		{
			name:            "builds request with 120 second timeout",
			timeout:         120,
			expectedTimeout: 120,
		},
		{
			name:            "builds request with minimum timeout",
			timeout:         1,
			expectedTimeout: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			builder := NewUpdateTimeoutRequestBuilder(tt.timeout)
			require.NotNil(t, builder, "builder should not be nil")

			req := builder.Build()
			require.NotNil(t, req, "built request should not be nil")
			assert.Equal(t, tt.expectedTimeout, req.Timeout, "timeout should match expected value")
		})
	}
}

func TestUpdateTimeoutRequest_JSONSerialization(t *testing.T) {
	tests := []struct {
		name         string
		timeout      int
		expectedJSON string
	}{
		{
			name:         "serializes 30 second timeout correctly",
			timeout:      30,
			expectedJSON: `{"timeout":30}`,
		},
		{
			name:         "serializes 60 second timeout correctly",
			timeout:      60,
			expectedJSON: `{"timeout":60}`,
		},
		{
			name:         "serializes 120 second timeout correctly",
			timeout:      120,
			expectedJSON: `{"timeout":120}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			builder := NewUpdateTimeoutRequestBuilder(tt.timeout)
			req := builder.Build()

			// Test JSON serialization
			jsonData, err := json.Marshal(req)
			require.NoError(t, err, "failed to marshal request to JSON")
			assert.JSONEq(t, tt.expectedJSON, string(jsonData), "JSON should match expected")

			// Test JSON deserialization round-trip
			var roundTrip updateTimeoutRequest
			err = json.Unmarshal(jsonData, &roundTrip)
			require.NoError(t, err, "failed to unmarshal JSON back to struct")
			assert.Equal(t, req.Timeout, roundTrip.Timeout, "timeout should match after round-trip")
		})
	}
}

func TestUpdateTimeoutResponse_FieldByFieldValidation(t *testing.T) {
	// Read the response file
	responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "integration", "update_200.json")
	responseData, err := os.ReadFile(responseFilePath)
	require.NoError(t, err, "failed to read response file")

	// Deserialize the JSON response
	var response UpdateTimeoutResponse
	err = json.Unmarshal(responseData, &response)
	require.NoError(t, err, "failed to unmarshal JSON response")

	// Validate response structure
	assert.True(t, response.Status.Bool())
	assert.Equal(t, "Payment session timeout updated", response.Message)
	require.NotNil(t, response.Data)

	// Validate timeout data
	data := response.Data
	assert.Equal(t, 30, data.PaymentSessionTimeout, "payment session timeout should match")

	// Test JSON round-trip
	marshaled, err := json.Marshal(response)
	require.NoError(t, err)

	var roundTrip UpdateTimeoutResponse
	err = json.Unmarshal(marshaled, &roundTrip)
	require.NoError(t, err)

	assert.Equal(t, response.Status, roundTrip.Status)
	assert.Equal(t, response.Message, roundTrip.Message)
	assert.Equal(t, response.Data.PaymentSessionTimeout, roundTrip.Data.PaymentSessionTimeout)
}
