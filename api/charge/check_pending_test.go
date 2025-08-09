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
			if response.Status.Bool() {
				assert.NotEmpty(t, response.Data.Reference.String(), "reference should not be empty")
				assert.NotEmpty(t, response.Data.Status.String(), "status should not be empty")
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

		var unmarshaled map[string]any
		err = json.Unmarshal(jsonData, &unmarshaled)
		require.NoError(t, err, "should unmarshal JSON without error")

		assert.Equal(t, "ref_123456789", unmarshaled["reference"], "reference should match")
	})
}

func TestCheckPendingResponse_FieldByFieldValidation(t *testing.T) {
	// Read the check_pending_200.json file
	responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "charge", "check_pending_200.json")
	responseData, err := os.ReadFile(responseFilePath)
	require.NoError(t, err, "failed to read check_pending_200.json")

	// Parse the raw JSON to get the original values
	var rawResponse map[string]any
	err = json.Unmarshal(responseData, &rawResponse)
	require.NoError(t, err, "failed to unmarshal raw JSON response")

	// Deserialize into the CheckPendingResponse struct
	var response CheckPendingResponse
	err = json.Unmarshal(responseData, &response)
	require.NoError(t, err, "failed to unmarshal into CheckPendingResponse struct")

	// Validate top-level fields against the raw JSON
	assert.Equal(t, rawResponse["status"], response.Status.Bool(), "status field should match")
	assert.Equal(t, rawResponse["message"], response.Message, "message field should match")

	// Validate data object fields (reusing the same structure as create response)
	rawData := rawResponse["data"].(map[string]any)
	assert.Equal(t, int(rawData["amount"].(float64)), int(response.Data.Amount.Int64()), "amount should match")
	assert.Equal(t, rawData["currency"], string(response.Data.Currency), "currency should match")

	// For timestamp comparison, parse both and compare the actual time values
	expectedTime, err := time.Parse(time.RFC3339, rawData["transaction_date"].(string))
	require.NoError(t, err, "should parse expected transaction_date")
	actualTime, err := time.Parse(time.RFC3339, response.Data.TransactionDate.String())
	require.NoError(t, err, "should parse actual transaction_date")
	assert.True(t, expectedTime.Equal(actualTime), "transaction_date should represent the same moment")

	assert.Equal(t, rawData["status"], response.Data.Status.String(), "data status should match")
	assert.Equal(t, rawData["reference"], response.Data.Reference.String(), "reference should match")
	assert.Equal(t, rawData["domain"], response.Data.Domain.String(), "domain should match")
	assert.Equal(t, rawData["gateway_response"], response.Data.GatewayResponse.String(), "gateway_response should match")

	// Handle null message field
	if rawData["message"] == nil {
		assert.Empty(t, response.Data.Message, "data message should be empty for null")
	} else {
		assert.Equal(t, rawData["message"], response.Data.Message.String(), "data message should match")
	}

	assert.Equal(t, rawData["channel"], string(response.Data.Channel), "channel should match")
	assert.Equal(t, rawData["ip_address"], response.Data.IPAddress.String(), "ip_address should match")

	// Handle null log field
	if rawData["log"] == nil {
		assert.False(t, response.Data.Log.Valid, "log should be invalid for null")
	} else {
		assert.Equal(t, rawData["log"], response.Data.Log, "log should match")
	}

	assert.Equal(t, int(rawData["fees"].(float64)), int(response.Data.Fees.Int64()), "fees should match")

	// Handle null plan field
	if rawData["plan"] == nil {
		assert.Nil(t, response.Data.Plan, "plan should be nil")
	} else {
		assert.Equal(t, rawData["plan"], response.Data.Plan, "plan should match")
	}

	// Validate metadata object (comparing content, not pointer)
	rawMetadata := rawData["metadata"].(map[string]any)
	assert.True(t, response.Data.Metadata.Valid, "metadata should be valid")
	assert.Equal(t, rawMetadata, map[string]any(response.Data.Metadata.Metadata), "metadata content should match")

	// Validate authorization object
	rawAuthorization := rawData["authorization"].(map[string]any)
	assert.Equal(t, rawAuthorization["authorization_code"], response.Data.Authorization.AuthorizationCode.String(), "authorization_code should match")
	assert.Equal(t, rawAuthorization["bin"], response.Data.Authorization.Bin.String(), "bin should match")
	assert.Equal(t, rawAuthorization["last4"], response.Data.Authorization.Last4.String(), "last4 should match")
	assert.Equal(t, rawAuthorization["exp_month"], response.Data.Authorization.ExpMonth.String(), "exp_month should match")
	assert.Equal(t, rawAuthorization["exp_year"], response.Data.Authorization.ExpYear.String(), "exp_year should match")
	assert.Equal(t, rawAuthorization["channel"], string(response.Data.Authorization.Channel), "authorization channel should match")
	assert.Equal(t, rawAuthorization["card_type"], response.Data.Authorization.CardType.String(), "card_type should match")
	assert.Equal(t, rawAuthorization["bank"], response.Data.Authorization.Bank.String(), "bank should match")
	assert.Equal(t, rawAuthorization["country_code"], response.Data.Authorization.CountryCode.String(), "country_code should match")
	assert.Equal(t, rawAuthorization["brand"], response.Data.Authorization.Brand.String(), "brand should match")
	assert.Equal(t, rawAuthorization["reusable"], response.Data.Authorization.Reusable.Bool(), "reusable should match")
	assert.Equal(t, rawAuthorization["signature"], response.Data.Authorization.Signature.String(), "signature should match")

	// Handle account_name nullable field
	if rawAuthorization["account_name"] != nil {
		assert.Equal(t, rawAuthorization["account_name"], response.Data.Authorization.AccountName.String(), "account_name should match")
	} else {
		assert.False(t, response.Data.Authorization.AccountName.Valid, "account_name should be invalid")
	}

	// Validate customer object
	rawCustomer := rawData["customer"].(map[string]any)
	assert.Equal(t, uint64(rawCustomer["id"].(float64)), response.Data.Customer.ID.Uint64(), "customer id should match")

	// Handle nullable string fields
	if rawCustomer["first_name"] == nil {
		assert.False(t, response.Data.Customer.FirstName.Valid, "first_name should be invalid")
	} else {
		assert.Equal(t, rawCustomer["first_name"], response.Data.Customer.FirstName.String(), "first_name should match")
	}

	if rawCustomer["last_name"] == nil {
		assert.False(t, response.Data.Customer.LastName.Valid, "last_name should be invalid")
	} else {
		assert.Equal(t, rawCustomer["last_name"], response.Data.Customer.LastName.String(), "last_name should match")
	}

	assert.Equal(t, rawCustomer["email"], response.Data.Customer.Email.String(), "email should match")
	assert.Equal(t, rawCustomer["customer_code"], response.Data.Customer.CustomerCode.String(), "customer_code should match")

	if rawCustomer["phone"] == nil {
		assert.False(t, response.Data.Customer.Phone.Valid, "phone should be invalid")
	} else {
		assert.Equal(t, rawCustomer["phone"], response.Data.Customer.Phone.String(), "phone should match")
	}

	// Handle customer metadata - null in JSON becomes invalid metadata
	if rawCustomer["metadata"] == nil {
		assert.False(t, response.Data.Customer.Metadata.Valid, "customer metadata should be invalid for null")
	} else {
		assert.True(t, response.Data.Customer.Metadata.Valid, "customer metadata should be valid")
		assert.Equal(t, rawCustomer["metadata"], map[string]any(response.Data.Customer.Metadata.Metadata), "customer metadata should match")
	}

	assert.Equal(t, rawCustomer["risk_action"], response.Data.Customer.RiskAction.String(), "risk_action should match")

	// Test round-trip serialization
	serialized, err := json.Marshal(response)
	require.NoError(t, err, "failed to marshal response back to JSON")

	var roundTripResponse CheckPendingResponse
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
