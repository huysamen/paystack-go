package charge

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSubmitPINResponse_JSONDeserialization(t *testing.T) {
	tests := []struct {
		name            string
		responseFile    string
		expectedStatus  bool
		expectedMessage string
	}{
		{
			name:            "successful submit PIN response",
			responseFile:    "submit_pin_200.json",
			expectedStatus:  true,
			expectedMessage: "Charge attempted",
		},
		{
			name:            "failed submit PIN response",
			responseFile:    "submit_pin_200_failed.json",
			expectedStatus:  true,
			expectedMessage: "Charge attempted",
		},
		{
			name:            "submit PIN with pending status",
			responseFile:    "submit_pin_200_pending.json",
			expectedStatus:  true,
			expectedMessage: "Charge attempted",
		},
		{
			name:            "submit PIN validation error",
			responseFile:    "submit_pin_400.json",
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
			var response SubmitPINResponse
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

func TestSubmitPINRequestBuilder(t *testing.T) {
	t.Run("builds request with PIN and reference", func(t *testing.T) {
		builder := NewSubmitPINRequestBuilder("1234", "ref_123456789")
		request := builder.Build()

		assert.Equal(t, "1234", request.PIN, "PIN should match")
		assert.Equal(t, "ref_123456789", request.Reference, "reference should match")
	})

	t.Run("builds request with empty values", func(t *testing.T) {
		builder := NewSubmitPINRequestBuilder("", "")
		request := builder.Build()

		assert.Equal(t, "", request.PIN, "PIN should be empty")
		assert.Equal(t, "", request.Reference, "reference should be empty")
	})
}

func TestSubmitPINRequest_JSONSerialization(t *testing.T) {
	t.Run("serializes request correctly", func(t *testing.T) {
		builder := NewSubmitPINRequestBuilder("1234", "ref_123456789")
		request := builder.Build()

		jsonData, err := json.Marshal(request)
		require.NoError(t, err, "should marshal to JSON without error")

		var unmarshaled map[string]any
		err = json.Unmarshal(jsonData, &unmarshaled)
		require.NoError(t, err, "should unmarshal JSON without error")

		assert.Equal(t, "1234", unmarshaled["pin"], "PIN should match")
		assert.Equal(t, "ref_123456789", unmarshaled["reference"], "reference should match")
	})
}

func TestSubmitPINResponse_FieldByFieldValidation(t *testing.T) {
	// Read the submit_pin_200.json file
	responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "charge", "submit_pin_200.json")
	responseData, err := os.ReadFile(responseFilePath)
	require.NoError(t, err, "failed to read submit_pin_200.json")

	// Parse the raw JSON to get the original values
	var rawResponse map[string]any
	err = json.Unmarshal(responseData, &rawResponse)
	require.NoError(t, err, "failed to unmarshal raw JSON response")

	// Deserialize into the SubmitPINResponse struct
	var response SubmitPINResponse
	err = json.Unmarshal(responseData, &response)
	require.NoError(t, err, "failed to unmarshal into SubmitPINResponse struct")

	// Validate top-level fields against the raw JSON
	assert.Equal(t, rawResponse["status"], response.Status.Bool(), "status field should match")
	assert.Equal(t, rawResponse["message"], response.Message, "message field should match")

	// Validate data object fields
	rawData := rawResponse["data"].(map[string]any)
	assert.Equal(t, int(rawData["id"].(float64)), response.Data.ID, "id should match")
	assert.Equal(t, rawData["domain"], response.Data.Domain, "domain should match")
	assert.Equal(t, rawData["status"], response.Data.Status, "data status should match")
	assert.Equal(t, rawData["reference"], response.Data.Reference, "reference should match")
	assert.Equal(t, int(rawData["amount"].(float64)), response.Data.Amount, "amount should match")
	assert.Equal(t, rawData["message"], response.Data.Message, "data message should match")
	assert.Equal(t, rawData["gateway_response"], response.Data.GatewayResponse, "gateway_response should match")

	// For timestamp comparisons, parse both and compare the actual time values
	expectedPaidAt, err := time.Parse(time.RFC3339, rawData["paid_at"].(string))
	require.NoError(t, err, "should parse expected paid_at")
	actualPaidAt, err := time.Parse(time.RFC3339, response.Data.PaidAt.String())
	require.NoError(t, err, "should parse actual paid_at")
	assert.True(t, expectedPaidAt.Equal(actualPaidAt), "paid_at should represent the same moment")

	expectedCreatedAt, err := time.Parse(time.RFC3339, rawData["created_at"].(string))
	require.NoError(t, err, "should parse expected created_at")
	actualCreatedAt, err := time.Parse(time.RFC3339, response.Data.CreatedAt.String())
	require.NoError(t, err, "should parse actual created_at")
	assert.True(t, expectedCreatedAt.Equal(actualCreatedAt), "created_at should represent the same moment")

	expectedTransactionDate, err := time.Parse(time.RFC3339, rawData["transaction_date"].(string))
	require.NoError(t, err, "should parse expected transaction_date")
	actualTransactionDate, err := time.Parse(time.RFC3339, response.Data.TransactionDate.String())
	require.NoError(t, err, "should parse actual transaction_date")
	assert.True(t, expectedTransactionDate.Equal(actualTransactionDate), "transaction_date should represent the same moment")

	// Validate other fields
	assert.Equal(t, rawData["channel"], string(response.Data.Channel), "channel should match")
	assert.Equal(t, rawData["currency"], string(response.Data.Currency), "currency should match")
	assert.Equal(t, rawData["ip_address"], response.Data.IPAddress, "ip_address should match")

	// Handle metadata - empty string in JSON becomes empty Metadata struct
	if rawData["metadata"] == "" || rawData["metadata"] == nil {
		if response.Data.Metadata == nil {
			assert.Nil(t, response.Data.Metadata, "metadata should be nil")
		} else {
			assert.Equal(t, map[string]any{}, map[string]any(*response.Data.Metadata), "metadata should be empty for empty string or null")
		}
	} else {
		assert.Equal(t, rawData["metadata"], map[string]any(*response.Data.Metadata), "metadata should match")
	}

	// Handle null log field
	if rawData["log"] == nil {
		assert.Nil(t, response.Data.Log, "log should be nil")
	} else {
		assert.Equal(t, rawData["log"], response.Data.Log, "log should match")
	}

	assert.Equal(t, int(rawData["fees"].(float64)), response.Data.Fees, "fees should match")
	assert.Equal(t, int(rawData["requested_amount"].(float64)), response.Data.RequestedAmount, "requested_amount should match")

	// Handle Plan object - can be null or have data
	if rawData["plan"] == nil {
		assert.Nil(t, response.Data.Plan, "plan should be nil")
	} else {
		assert.NotNil(t, response.Data.Plan, "plan should not be nil when present in JSON")
	}

	// Validate authorization object
	rawAuthorization := rawData["authorization"].(map[string]any)
	assert.Equal(t, rawAuthorization["authorization_code"], response.Data.Authorization.AuthorizationCode, "authorization_code should match")
	assert.Equal(t, rawAuthorization["bin"], response.Data.Authorization.Bin, "bin should match")
	assert.Equal(t, rawAuthorization["last4"], response.Data.Authorization.Last4, "last4 should match")
	assert.Equal(t, rawAuthorization["exp_month"], response.Data.Authorization.ExpMonth.String(), "exp_month should match")
	assert.Equal(t, rawAuthorization["exp_year"], response.Data.Authorization.ExpYear.String(), "exp_year should match")
	assert.Equal(t, rawAuthorization["channel"], string(response.Data.Authorization.Channel), "authorization channel should match")
	assert.Equal(t, rawAuthorization["card_type"], response.Data.Authorization.CardType, "card_type should match")
	assert.Equal(t, rawAuthorization["bank"], response.Data.Authorization.Bank, "bank should match")
	assert.Equal(t, rawAuthorization["country_code"], response.Data.Authorization.CountryCode, "country_code should match")
	assert.Equal(t, rawAuthorization["brand"], response.Data.Authorization.Brand, "brand should match")
	assert.Equal(t, rawAuthorization["reusable"], response.Data.Authorization.Reusable, "reusable should match")

	// Handle signature - null in JSON becomes empty string in struct
	if rawAuthorization["signature"] == nil {
		assert.Equal(t, "", response.Data.Authorization.Signature, "signature should be empty string when null in JSON")
	} else {
		assert.Equal(t, rawAuthorization["signature"], response.Data.Authorization.Signature, "signature should match")
	}

	// Handle nullable account_name field
	if rawAuthorization["account_name"] == nil {
		assert.Nil(t, response.Data.Authorization.AccountName, "account_name should be nil")
	} else {
		assert.Equal(t, rawAuthorization["account_name"], *response.Data.Authorization.AccountName, "account_name should match")
	}

	// Validate customer object
	rawCustomer := rawData["customer"].(map[string]any)
	assert.Equal(t, uint64(rawCustomer["id"].(float64)), response.Data.Customer.ID, "customer id should match")

	// Handle nullable string fields
	if rawCustomer["first_name"] == nil {
		assert.Nil(t, response.Data.Customer.FirstName, "first_name should be nil")
	} else {
		assert.Equal(t, rawCustomer["first_name"], *response.Data.Customer.FirstName, "first_name should match")
	}

	if rawCustomer["last_name"] == nil {
		assert.Nil(t, response.Data.Customer.LastName, "last_name should be nil")
	} else {
		assert.Equal(t, rawCustomer["last_name"], *response.Data.Customer.LastName, "last_name should match")
	}

	assert.Equal(t, rawCustomer["email"], response.Data.Customer.Email, "email should match")
	assert.Equal(t, rawCustomer["customer_code"], response.Data.Customer.CustomerCode, "customer_code should match")

	if rawCustomer["phone"] == nil {
		assert.Nil(t, response.Data.Customer.Phone, "phone should be nil")
	} else {
		assert.Equal(t, rawCustomer["phone"], *response.Data.Customer.Phone, "phone should match")
	}

	// Handle customer metadata - null in JSON becomes empty struct
	if rawCustomer["metadata"] == nil {
		assert.Equal(t, map[string]any{}, map[string]any(response.Data.Customer.Metadata), "customer metadata should be empty for null")
	} else {
		assert.Equal(t, rawCustomer["metadata"], map[string]any(response.Data.Customer.Metadata), "customer metadata should match")
	}

	assert.Equal(t, rawCustomer["risk_action"], response.Data.Customer.RiskAction, "risk_action should match")

	// Test round-trip serialization
	serialized, err := json.Marshal(response)
	require.NoError(t, err, "failed to marshal response back to JSON")

	var roundTripResponse SubmitPINResponse
	err = json.Unmarshal(serialized, &roundTripResponse)
	require.NoError(t, err, "failed to unmarshal round-trip JSON")

	// Verify core fields survive round-trip
	assert.Equal(t, response.Status.Bool(), roundTripResponse.Status.Bool(), "status should survive round-trip")
	assert.Equal(t, response.Message, roundTripResponse.Message, "message should survive round-trip")
	assert.Equal(t, response.Data.Amount, roundTripResponse.Data.Amount, "amount should survive round-trip")
	assert.Equal(t, response.Data.Reference, roundTripResponse.Data.Reference, "reference should survive round-trip")
	assert.Equal(t, response.Data.Customer.Email, roundTripResponse.Data.Customer.Email, "customer email should survive round-trip")
	assert.Equal(t, response.Data.Authorization.AuthorizationCode, roundTripResponse.Data.Authorization.AuthorizationCode, "authorization_code should survive round-trip")
}
