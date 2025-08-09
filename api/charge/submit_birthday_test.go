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

func TestSubmitBirthdayResponse_JSONDeserialization(t *testing.T) {
	tests := []struct {
		name            string
		responseFile    string
		expectedStatus  bool
		expectedMessage string
	}{
		{
			name:            "successful submit birthday response",
			responseFile:    "submit_birthday_200.json",
			expectedStatus:  true,
			expectedMessage: "Charge attempted",
		},
		{
			name:            "submit birthday requiring bank authorization",
			responseFile:    "submit_birthday_200_bank_auth.json",
			expectedStatus:  true,
			expectedMessage: "Charge attempted",
		},
		{
			name:            "failed submit birthday response",
			responseFile:    "submit_birthday_200_failed.json",
			expectedStatus:  true,
			expectedMessage: "Charge attempted",
		},
		{
			name:            "submit birthday requiring OTP",
			responseFile:    "submit_birthday_200_otp.json",
			expectedStatus:  true,
			expectedMessage: "Charge attempted",
		},
		{
			name:            "submit birthday with pending status",
			responseFile:    "submit_birthday_200_pending.json",
			expectedStatus:  true,
			expectedMessage: "Charge attempted",
		},
		{
			name:            "submit birthday validation error",
			responseFile:    "submit_birthday_400.json",
			expectedStatus:  false,
			expectedMessage: "Transaction reference is invalid",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "charge", tt.responseFile)
			responseData, err := os.ReadFile(responseFilePath)
			require.NoError(t, err, "failed to read response file: %s", responseFilePath)

			var response SubmitBirthdayResponse
			err = json.Unmarshal(responseData, &response)
			require.NoError(t, err, "failed to unmarshal JSON response")

			assert.Equal(t, tt.expectedStatus, response.Status.Bool(), "status should match")
			assert.Equal(t, tt.expectedMessage, response.Message, "message should match")

			if tt.expectedStatus {
				// Some responses may not have reference/status fields (like OTP flow)
				if response.Data.Reference.String() != "" {
					assert.NotEmpty(t, response.Data.Reference.String(), "reference should not be empty")
				}
				if response.Data.Status.String() != "" {
					assert.NotEmpty(t, response.Data.Status.String(), "status should not be empty")
				}
			}
		})
	}
}

func TestSubmitBirthdayResponse_FieldByFieldValidation(t *testing.T) {
	// Read the submit_birthday_200.json file
	responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "charge", "submit_birthday_200.json")
	responseData, err := os.ReadFile(responseFilePath)
	require.NoError(t, err, "failed to read submit_birthday_200.json")

	// Parse the raw JSON to get the original values
	var rawResponse map[string]any
	err = json.Unmarshal(responseData, &rawResponse)
	require.NoError(t, err, "failed to unmarshal raw JSON response")

	// Deserialize into the SubmitBirthdayResponse struct
	var response SubmitBirthdayResponse
	err = json.Unmarshal(responseData, &response)
	require.NoError(t, err, "failed to unmarshal into SubmitBirthdayResponse struct")

	// Validate top-level fields against the raw JSON
	assert.Equal(t, rawResponse["status"], response.Status.Bool(), "status field should match")
	assert.Equal(t, rawResponse["message"], response.Message, "message field should match")

	// Validate data object fields
	rawData := rawResponse["data"].(map[string]any)
	assert.Equal(t, int(rawData["id"].(float64)), int(response.Data.ID.Int64()), "id should match")
	assert.Equal(t, rawData["domain"], response.Data.Domain.String(), "domain should match")
	assert.Equal(t, rawData["status"], response.Data.Status.String(), "data status should match")
	assert.Equal(t, rawData["reference"], response.Data.Reference.String(), "reference should match")
	assert.Equal(t, int(rawData["amount"].(float64)), int(response.Data.Amount.Int64()), "amount should match")
	assert.Equal(t, rawData["message"], response.Data.Message.String(), "data message should match")
	assert.Equal(t, rawData["gateway_response"], response.Data.GatewayResponse.String(), "gateway_response should match")

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
	assert.Equal(t, rawData["ip_address"], response.Data.IPAddress.String(), "ip_address should match")

	// Handle metadata - empty string in JSON becomes invalid metadata
	if rawData["metadata"] == "" || rawData["metadata"] == nil {
		assert.False(t, response.Data.Metadata.Valid, "metadata should be invalid for empty string or null")
	} else {
		assert.True(t, response.Data.Metadata.Valid, "metadata should be valid")
		assert.Equal(t, rawData["metadata"], map[string]any(response.Data.Metadata.Metadata), "metadata should match")
	}

	// Handle null log field
	if rawData["log"] == nil {
		assert.False(t, response.Data.Log.Valid, "log should be invalid for null")
	} else {
		assert.Equal(t, rawData["log"], response.Data.Log, "log should match")
	}

	assert.Equal(t, int(rawData["fees"].(float64)), int(response.Data.Fees.Int64()), "fees should match")
	assert.Equal(t, int(rawData["requested_amount"].(float64)), int(response.Data.RequestedAmount.Int64()), "requested_amount should match")

	// Handle Plan object - can be null or have data
	if rawData["plan"] == nil {
		assert.Nil(t, response.Data.Plan, "plan should be nil")
	} else {
		assert.NotNil(t, response.Data.Plan, "plan should not be nil when present in JSON")
	}

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

	// Handle signature - null in JSON becomes empty string in struct
	if rawAuthorization["signature"] == nil {
		assert.Equal(t, "", response.Data.Authorization.Signature.String(), "signature should be empty string when null in JSON")
	} else {
		assert.Equal(t, rawAuthorization["signature"], response.Data.Authorization.Signature, "signature should match")
	}

	// Handle nullable account_name field
	if rawAuthorization["account_name"] == nil {
		assert.False(t, response.Data.Authorization.AccountName.Valid, "account_name should be invalid for null")
	} else {
		assert.Equal(t, rawAuthorization["account_name"], response.Data.Authorization.AccountName.String(), "account_name should match")
	}

	// Validate customer object
	rawCustomer := rawData["customer"].(map[string]any)
	assert.Equal(t, uint64(rawCustomer["id"].(float64)), response.Data.Customer.ID.Uint64(), "customer id should match")

	// Handle nullable string fields
	if rawCustomer["first_name"] == nil {
		assert.False(t, response.Data.Customer.FirstName.Valid, "first_name should be invalid for null")
	} else {
		assert.Equal(t, rawCustomer["first_name"], response.Data.Customer.FirstName.String(), "first_name should match")
	}

	if rawCustomer["last_name"] == nil {
		assert.False(t, response.Data.Customer.LastName.Valid, "last_name should be invalid for null")
	} else {
		assert.Equal(t, rawCustomer["last_name"], response.Data.Customer.LastName.String(), "last_name should match")
	}

	assert.Equal(t, rawCustomer["email"], response.Data.Customer.Email.String(), "email should match")
	assert.Equal(t, rawCustomer["customer_code"], response.Data.Customer.CustomerCode.String(), "customer_code should match")

	if rawCustomer["phone"] == nil {
		assert.False(t, response.Data.Customer.Phone.Valid, "phone should be invalid for null")
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

	var roundTripResponse SubmitBirthdayResponse
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
