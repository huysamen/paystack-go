package disputes

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetUploadURLResponse_JSONDeserialization(t *testing.T) {
	tests := []struct {
		name            string
		responseFile    string
		expectedStatus  bool
		expectedMessage string
	}{
		{
			name:            "successful get upload url response",
			responseFile:    "get_upload_url_200.json",
			expectedStatus:  true,
			expectedMessage: "Upload url generated",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Read the response file
			responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "disputes", tt.responseFile)
			responseData, err := os.ReadFile(responseFilePath)
			require.NoError(t, err, "failed to read response file: %s", responseFilePath)

			// Deserialize the JSON response
			var response GetUploadURLResponse
			err = json.Unmarshal(responseData, &response)
			require.NoError(t, err, "failed to unmarshal JSON response")

			// Verify the response structure
			assert.Equal(t, tt.expectedStatus, response.Status.Bool(), "status should match")
			assert.Equal(t, tt.expectedMessage, response.Message, "message should match")

			// Only verify data structure for successful responses
			if tt.expectedStatus {
				assert.NotNil(t, response.Data, "data should not be nil")

				// Verify upload URL data
				data := response.Data
				assert.NotEmpty(t, data.SignedURL, "signed URL should not be empty")
				assert.NotEmpty(t, data.FileName, "file name should not be empty")
			}
		})
	}
}

func TestGetUploadURLRequestBuilder(t *testing.T) {
	testCases := []struct {
		name     string
		setup    func() *GetUploadURLRequestBuilder
		expected *getUploadURLRequest
	}{
		{
			name: "builds request with upload filename",
			setup: func() *GetUploadURLRequestBuilder {
				return NewGetUploadURLRequestBuilder("evidence.pdf")
			},
			expected: &getUploadURLRequest{
				UploadFileName: "evidence.pdf",
			},
		},
		{
			name: "builds request with different filename",
			setup: func() *GetUploadURLRequestBuilder {
				return NewGetUploadURLRequestBuilder("dispute_document.jpg")
			},
			expected: &getUploadURLRequest{
				UploadFileName: "dispute_document.jpg",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.setup().Build()
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestGetUploadURLResponse_FieldByFieldValidation(t *testing.T) {
	// Read the response file
	responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "disputes", "get_upload_url_200.json")
	responseData, err := os.ReadFile(responseFilePath)
	require.NoError(t, err, "failed to read response file")

	// Deserialize the JSON response
	var response GetUploadURLResponse
	err = json.Unmarshal(responseData, &response)
	require.NoError(t, err, "failed to unmarshal JSON response")

	// Validate response structure
	assert.True(t, response.Status.Bool())
	assert.Equal(t, "Upload url generated", response.Message)
	require.NotNil(t, response.Data)

	// Validate upload URL data
	data := response.Data
	assert.Contains(t, data.SignedURL.String(), "https://s3.eu-west-1.amazonaws.com", "signed URL should be AWS S3 URL")
	assert.Equal(t, "qesp8a4df1xejihd9x5q", data.FileName.String())
	assert.Greater(t, len(data.SignedURL), 100, "signed URL should be reasonably long")

	// Test JSON round-trip
	marshaled, err := json.Marshal(response)
	require.NoError(t, err)

	var roundTrip GetUploadURLResponse
	err = json.Unmarshal(marshaled, &roundTrip)
	require.NoError(t, err)

	assert.Equal(t, response.Status, roundTrip.Status)
	assert.Equal(t, response.Message, roundTrip.Message)
	assert.Equal(t, response.Data.FileName, roundTrip.Data.FileName)
	assert.Equal(t, response.Data.SignedURL, roundTrip.Data.SignedURL)
}
