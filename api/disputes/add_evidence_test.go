package disputes

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddEvidenceResponse_JSONDeserialization(t *testing.T) {
	tests := []struct {
		name            string
		responseFile    string
		expectedStatus  bool
		expectedMessage string
	}{
		{
			name:            "successful add evidence response",
			responseFile:    "add_evidence_200.json",
			expectedStatus:  true,
			expectedMessage: "Evidence created",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Read the response file
			responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "disputes", tt.responseFile)
			responseData, err := os.ReadFile(responseFilePath)
			require.NoError(t, err, "failed to read response file: %s", responseFilePath)

			// Deserialize the JSON response
			var response AddEvidenceResponse
			err = json.Unmarshal(responseData, &response)
			require.NoError(t, err, "failed to unmarshal JSON response")

			// Verify the response structure
			assert.Equal(t, tt.expectedStatus, response.Status.Bool(), "status should match")
			assert.Equal(t, tt.expectedMessage, response.Message, "message should match")

			// Only verify data structure for successful responses
			if tt.expectedStatus {
				assert.NotNil(t, response.Data, "data should not be nil")

				// Verify evidence structure
				evidence := response.Data
				assert.Greater(t, evidence.ID, 0, "evidence ID should be positive")
				assert.NotEmpty(t, evidence.CustomerEmail, "customer email should not be empty")
				assert.NotEmpty(t, evidence.CustomerName, "customer name should not be empty")
				assert.NotEmpty(t, evidence.CustomerPhone, "customer phone should not be empty")
				assert.NotEmpty(t, evidence.ServiceDetails, "service details should not be empty")
				assert.Greater(t, evidence.Dispute, 0, "dispute ID should be positive")
			}
		})
	}
}

func TestAddEvidenceRequestBuilder(t *testing.T) {
	testCases := []struct {
		name     string
		setup    func() *AddEvidenceRequestBuilder
		expected *addEvidenceRequest
	}{
		{
			name: "builds basic request with required fields only",
			setup: func() *AddEvidenceRequestBuilder {
				return NewAddEvidenceRequestBuilder(
					"test@example.com",
					"John Doe",
					"+1234567890",
					"Product delivered successfully",
				)
			},
			expected: &addEvidenceRequest{
				CustomerEmail:  "test@example.com",
				CustomerName:   "John Doe",
				CustomerPhone:  "+1234567890",
				ServiceDetails: "Product delivered successfully",
			},
		},
		{
			name: "builds request with delivery address",
			setup: func() *AddEvidenceRequestBuilder {
				return NewAddEvidenceRequestBuilder(
					"test@example.com",
					"John Doe",
					"+1234567890",
					"Product delivered successfully",
				).DeliveryAddress("123 Main St, City")
			},
			expected: &addEvidenceRequest{
				CustomerEmail:   "test@example.com",
				CustomerName:    "John Doe",
				CustomerPhone:   "+1234567890",
				ServiceDetails:  "Product delivered successfully",
				DeliveryAddress: stringPtr("123 Main St, City"),
			},
		},
		{
			name: "builds request with delivery date",
			setup: func() *AddEvidenceRequestBuilder {
				deliveryDate := time.Date(2023, 1, 15, 10, 0, 0, 0, time.UTC)
				return NewAddEvidenceRequestBuilder(
					"test@example.com",
					"John Doe",
					"+1234567890",
					"Product delivered successfully",
				).DeliveryDate(deliveryDate)
			},
			expected: &addEvidenceRequest{
				CustomerEmail:  "test@example.com",
				CustomerName:   "John Doe",
				CustomerPhone:  "+1234567890",
				ServiceDetails: "Product delivered successfully",
				DeliveryDate:   timePtr(time.Date(2023, 1, 15, 10, 0, 0, 0, time.UTC)),
			},
		},
		{
			name: "builds complete request with all fields",
			setup: func() *AddEvidenceRequestBuilder {
				deliveryDate := time.Date(2023, 1, 15, 10, 0, 0, 0, time.UTC)
				return NewAddEvidenceRequestBuilder(
					"customer@company.com",
					"Jane Smith",
					"+1555123456",
					"Digital service provided via email",
				).
					DeliveryAddress("456 Business Ave, Tech City").
					DeliveryDate(deliveryDate)
			},
			expected: &addEvidenceRequest{
				CustomerEmail:   "customer@company.com",
				CustomerName:    "Jane Smith",
				CustomerPhone:   "+1555123456",
				ServiceDetails:  "Digital service provided via email",
				DeliveryAddress: stringPtr("456 Business Ave, Tech City"),
				DeliveryDate:    timePtr(time.Date(2023, 1, 15, 10, 0, 0, 0, time.UTC)),
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

func TestAddEvidenceRequest_JSONSerialization(t *testing.T) {
	testCases := []struct {
		name     string
		request  *addEvidenceRequest
		expected string
	}{
		{
			name: "serializes basic request correctly",
			request: &addEvidenceRequest{
				CustomerEmail:  "test@example.com",
				CustomerName:   "John Doe",
				CustomerPhone:  "+1234567890",
				ServiceDetails: "Product delivered",
			},
			expected: `{"customer_email":"test@example.com","customer_name":"John Doe","customer_phone":"+1234567890","service_details":"Product delivered"}`,
		},
		{
			name: "includes optional fields when provided",
			request: &addEvidenceRequest{
				CustomerEmail:   "test@example.com",
				CustomerName:    "John Doe",
				CustomerPhone:   "+1234567890",
				ServiceDetails:  "Product delivered",
				DeliveryAddress: stringPtr("123 Main St"),
				DeliveryDate:    timePtr(time.Date(2023, 1, 15, 10, 0, 0, 0, time.UTC)),
			},
			expected: `{"customer_email":"test@example.com","customer_name":"John Doe","customer_phone":"+1234567890","service_details":"Product delivered","delivery_address":"123 Main St","delivery_date":"2023-01-15T10:00:00Z"}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			marshaled, err := json.Marshal(tc.request)
			require.NoError(t, err)
			assert.JSONEq(t, tc.expected, string(marshaled))
		})
	}
}

func TestAddEvidenceResponse_FieldByFieldValidation(t *testing.T) {
	// Read the response file
	responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "disputes", "add_evidence_200.json")
	responseData, err := os.ReadFile(responseFilePath)
	require.NoError(t, err, "failed to read response file")

	// Deserialize the JSON response
	var response AddEvidenceResponse
	err = json.Unmarshal(responseData, &response)
	require.NoError(t, err, "failed to unmarshal JSON response")

	// Validate response structure
	assert.True(t, response.Status.Bool())
	assert.Equal(t, "Evidence created", response.Message)
	require.NotNil(t, response.Data)

	// Validate evidence details
	evidence := response.Data
	assert.Equal(t, 21, evidence.ID)
	assert.Equal(t, "cus@gmail.com", evidence.CustomerEmail)
	assert.Equal(t, "Mensah King", evidence.CustomerName)
	assert.Equal(t, "0802345167", evidence.CustomerPhone)
	assert.Equal(t, "claim for buying product", evidence.ServiceDetails)
	require.NotNil(t, evidence.DeliveryAddress)
	assert.Equal(t, "3a ladoke street ogbomoso", *evidence.DeliveryAddress)
	assert.Equal(t, 624, evidence.Dispute)

	// Test JSON round-trip
	marshaled, err := json.Marshal(response)
	require.NoError(t, err)

	var roundTrip AddEvidenceResponse
	err = json.Unmarshal(marshaled, &roundTrip)
	require.NoError(t, err)

	assert.Equal(t, response.Status, roundTrip.Status)
	assert.Equal(t, response.Message, roundTrip.Message)
	assert.Equal(t, response.Data.ID, roundTrip.Data.ID)
}
