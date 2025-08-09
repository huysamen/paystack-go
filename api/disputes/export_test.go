package disputes

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/huysamen/paystack-go/enums"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExportResponse_JSONDeserialization(t *testing.T) {
	tests := []struct {
		name            string
		responseFile    string
		expectedStatus  bool
		expectedMessage string
	}{
		{
			name:            "successful export response",
			responseFile:    "export_200.json",
			expectedStatus:  true,
			expectedMessage: "Export successful",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Read the response file
			responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "disputes", tt.responseFile)
			responseData, err := os.ReadFile(responseFilePath)
			require.NoError(t, err, "failed to read response file: %s", responseFilePath)

			// Deserialize the JSON response
			var response ExportResponse
			err = json.Unmarshal(responseData, &response)
			require.NoError(t, err, "failed to unmarshal JSON response")

			// Verify the response structure
			assert.Equal(t, tt.expectedStatus, response.Status.Bool(), "status should match")
			assert.Equal(t, tt.expectedMessage, response.Message, "message should match")

			// Only verify data structure for successful responses
			if tt.expectedStatus {
				assert.NotNil(t, response.Data, "data should not be nil")

				// Verify export data
				data := response.Data
				assert.NotEmpty(t, data.Path, "export path should not be empty")
			}
		})
	}
}

func TestExportRequestBuilder(t *testing.T) {
	testCases := []struct {
		name     string
		setup    func() *ExportRequestBuilder
		expected *exportRequest
	}{
		{
			name: "builds empty request",
			setup: func() *ExportRequestBuilder {
				return NewExportRequestBuilder()
			},
			expected: &exportRequest{},
		},
		{
			name: "builds request with pagination",
			setup: func() *ExportRequestBuilder {
				return NewExportRequestBuilder().
					PerPage(20).
					Page(2)
			},
			expected: &exportRequest{
				PerPage: intPtr(20),
				Page:    intPtr(2),
			},
		},
		{
			name: "builds request with date range",
			setup: func() *ExportRequestBuilder {
				from := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
				to := time.Date(2023, 1, 31, 0, 0, 0, 0, time.UTC)
				return NewExportRequestBuilder().DateRange(from, to)
			},
			expected: &exportRequest{
				From: timePtr(time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)),
				To:   timePtr(time.Date(2023, 1, 31, 0, 0, 0, 0, time.UTC)),
			},
		},
		{
			name: "builds request with transaction and status",
			setup: func() *ExportRequestBuilder {
				return NewExportRequestBuilder().
					Transaction("TXN_123").
					Status(enums.DisputeStatusResolved)
			},
			expected: &exportRequest{
				Transaction: stringPtr("TXN_123"),
				Status:      statusPtr(enums.DisputeStatusResolved),
			},
		},
		{
			name: "builds complete request with all fields",
			setup: func() *ExportRequestBuilder {
				from := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
				to := time.Date(2023, 1, 31, 0, 0, 0, 0, time.UTC)
				return NewExportRequestBuilder().
					From(from).
					To(to).
					PerPage(25).
					Page(3).
					Transaction("TXN_456").
					Status(enums.DisputeStatusPending)
			},
			expected: &exportRequest{
				From:        timePtr(time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)),
				To:          timePtr(time.Date(2023, 1, 31, 0, 0, 0, 0, time.UTC)),
				PerPage:     intPtr(25),
				Page:        intPtr(3),
				Transaction: stringPtr("TXN_456"),
				Status:      statusPtr(enums.DisputeStatusPending),
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

func TestExportRequest_QueryGeneration(t *testing.T) {
	testCases := []struct {
		name     string
		request  *exportRequest
		expected string
	}{
		{
			name:     "generates empty query for empty request",
			request:  &exportRequest{},
			expected: "",
		},
		{
			name: "generates query with all parameters",
			request: &exportRequest{
				From:        timePtr(time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)),
				To:          timePtr(time.Date(2023, 1, 31, 0, 0, 0, 0, time.UTC)),
				PerPage:     intPtr(25),
				Page:        intPtr(2),
				Transaction: stringPtr("TXN_123"),
				Status:      statusPtr(enums.DisputeStatusArchived),
			},
			expected: "from=2023-01-01&page=2&per_page=25&status=archived&to=2023-01-31&transaction=TXN_123",
		},
		{
			name: "generates query with date range only",
			request: &exportRequest{
				From: timePtr(time.Date(2023, 6, 15, 0, 0, 0, 0, time.UTC)),
				To:   timePtr(time.Date(2023, 6, 30, 0, 0, 0, 0, time.UTC)),
			},
			expected: "from=2023-06-15&to=2023-06-30",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.request.toQuery()
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestExportResponse_FieldByFieldValidation(t *testing.T) {
	// Read the response file
	responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "disputes", "export_200.json")
	responseData, err := os.ReadFile(responseFilePath)
	require.NoError(t, err, "failed to read response file")

	// Deserialize the JSON response
	var response ExportResponse
	err = json.Unmarshal(responseData, &response)
	require.NoError(t, err, "failed to unmarshal JSON response")

	// Validate response structure
	assert.True(t, response.Status.Bool())
	assert.Equal(t, "Export successful", response.Message)
	require.NotNil(t, response.Data)

	// Validate export data
	data := response.Data
	assert.Contains(t, data.Path, "https://s3.eu-west-1.amazonaws.com", "export path should be AWS S3 URL")
	assert.Contains(t, data.Path, "disputes", "export path should contain 'disputes'")
	assert.Contains(t, data.Path, ".csv", "export path should be CSV file")

	// Validate expiration date if present
	if data.ExpiresAt.Valid {
		// Should be a valid datetime
		assert.True(t, data.ExpiresAt.Valid, "expires_at should be present")
	}

	// Test JSON round-trip
	marshaled, err := json.Marshal(response)
	require.NoError(t, err)

	var roundTrip ExportResponse
	err = json.Unmarshal(marshaled, &roundTrip)
	require.NoError(t, err)

	assert.Equal(t, response.Status, roundTrip.Status)
	assert.Equal(t, response.Message, roundTrip.Message)
	assert.Equal(t, response.Data.Path, roundTrip.Data.Path)
}
