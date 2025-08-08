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

func TestCreateResponse_JSONDeserialization(t *testing.T) {
	tests := []struct {
		name            string
		responseFile    string
		expectedStatus  bool
		expectedMessage string
	}{
		{
			name:            "successful create charge response",
			responseFile:    "create_200.json",
			expectedStatus:  true,
			expectedMessage: "Charge attempted",
		},
		{
			name:            "create charge with address verification",
			responseFile:    "create_200_address.json",
			expectedStatus:  true,
			expectedMessage: "Charge attempted",
		},
		{
			name:            "create charge requiring OTP",
			responseFile:    "create_200_otp.json",
			expectedStatus:  true,
			expectedMessage: "Charge attempted",
		},
		{
			name:            "create charge requiring PIN",
			responseFile:    "create_200_pin.json",
			expectedStatus:  true,
			expectedMessage: "Charge attempted",
		},
		{
			name:            "create charge requiring bank authorization",
			responseFile:    "create_200_bank_auth.json",
			expectedStatus:  true,
			expectedMessage: "Charge attempted",
		},
		{
			name:            "create charge with birthday verification",
			responseFile:    "create_200_birthday.json",
			expectedStatus:  true,
			expectedMessage: "Charge attempted",
		},
		{
			name:            "create charge with mobile money",
			responseFile:    "create_200_momo.json",
			expectedStatus:  true,
			expectedMessage: "Charge attempted",
		},
		{
			name:            "create charge with pending status",
			responseFile:    "create_200_pending.json",
			expectedStatus:  true,
			expectedMessage: "Reference check successful",
		},
		{
			name:            "create charge requiring phone verification",
			responseFile:    "create_200_phone.json",
			expectedStatus:  true,
			expectedMessage: "Charge attempted",
		},
		{
			name:            "create charge with USSD",
			responseFile:    "create_200_ussd.json",
			expectedStatus:  true,
			expectedMessage: "Charge attempted",
		},
		{
			name:            "create charge validation error",
			responseFile:    "create_400.json",
			expectedStatus:  false,
			expectedMessage: "Email address is required for association with card",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Read the response file
			responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "charge", tt.responseFile)
			responseData, err := os.ReadFile(responseFilePath)
			require.NoError(t, err, "failed to read response file: %s", responseFilePath)

			// Deserialize the JSON response
			var response CreateChargeResponse
			err = json.Unmarshal(responseData, &response)
			require.NoError(t, err, "failed to unmarshal JSON response")

			// Verify the response structure
			assert.Equal(t, tt.expectedStatus, response.Status.Bool(), "status should match")
			assert.Equal(t, tt.expectedMessage, response.Message, "message should match")

			// Only verify data structure for successful responses
			if tt.expectedStatus {
				assert.NotNil(t, response.Data, "data should not be nil")

				// Some response types may not have amount/currency (like address verification flows)
				if response.Data.Amount > 0 {
					assert.Greater(t, response.Data.Amount, 0, "amount should be greater than 0")
				}
				if response.Data.Currency != "" {
					assert.NotEmpty(t, string(response.Data.Currency), "currency should not be empty")
				}

				assert.NotEmpty(t, response.Data.Reference, "reference should not be empty")
				assert.NotEmpty(t, response.Data.Status, "status should not be empty")

				// Verify customer data if present
				if response.Data.Customer != nil {
					assert.NotEmpty(t, response.Data.Customer.Email, "customer email should not be empty")
					assert.NotEmpty(t, response.Data.Customer.CustomerCode, "customer code should not be empty")
				}
			}
		})
	}
}

func TestCreateRequestBuilder(t *testing.T) {
	t.Run("builds basic request with required fields", func(t *testing.T) {
		builder := NewCreateRequestBuilder("test@example.com", "50000")
		request := builder.Build()

		assert.Equal(t, "test@example.com", request.Email, "email should match")
		assert.Equal(t, "50000", request.Amount, "amount should match")
	})

	t.Run("builds request with all optional fields", func(t *testing.T) {
		metadata := map[string]any{"key": "value"}
		bankDetails := &BankDetails{
			Code:          "057",
			AccountNumber: "0123456789",
		}

		builder := NewCreateRequestBuilder("test@example.com", "50000")
		request := builder.
			SplitCode("SPL_123").
			Subaccount("ACCT_456").
			TransactionCharge(100).
			Bearer("account").
			Bank(bankDetails).
			AuthorizationCode("AUTH_123").
			PIN("1234").
			Metadata(metadata).
			Reference("ref_123").
			DeviceID("device_123").
			Birthday("1990-01-01").
			Build()

		assert.Equal(t, "test@example.com", request.Email)
		assert.Equal(t, "50000", request.Amount)
		assert.Equal(t, "SPL_123", *request.SplitCode)
		assert.Equal(t, "ACCT_456", *request.Subaccount)
		assert.Equal(t, 100, *request.TransactionCharge)
		assert.Equal(t, "account", *request.Bearer)
		assert.Equal(t, bankDetails, request.Bank)
		assert.Equal(t, "AUTH_123", *request.AuthorizationCode)
		assert.Equal(t, "1234", *request.PIN)
		assert.Equal(t, metadata, request.Metadata)
		assert.Equal(t, "ref_123", *request.Reference)
		assert.Equal(t, "device_123", *request.DeviceID)
		assert.Equal(t, "1990-01-01", *request.Birthday)
	})

	t.Run("builds request with mobile money details", func(t *testing.T) {
		mobileMoneyDetails := &MobileMoneyDetails{
			Phone:    "+2348123456789",
			Provider: "mtn",
		}

		builder := NewCreateRequestBuilder("test@example.com", "50000")
		request := builder.
			MobileMoney(mobileMoneyDetails).
			Build()

		assert.Equal(t, mobileMoneyDetails, request.MobileMoney)
	})

	t.Run("builds request with USSD details", func(t *testing.T) {
		ussdDetails := &USSDDetails{
			Type: "737",
		}

		builder := NewCreateRequestBuilder("test@example.com", "50000")
		request := builder.
			USSD(ussdDetails).
			Build()

		assert.Equal(t, ussdDetails, request.USSD)
	})
}

func TestCreateRequest_JSONSerialization(t *testing.T) {
	t.Run("serializes basic request correctly", func(t *testing.T) {
		builder := NewCreateRequestBuilder("test@example.com", "50000")
		request := builder.Build()

		jsonData, err := json.Marshal(request)
		require.NoError(t, err, "should marshal to JSON without error")

		var unmarshaled map[string]any
		err = json.Unmarshal(jsonData, &unmarshaled)
		require.NoError(t, err, "should unmarshal JSON without error")

		assert.Equal(t, "test@example.com", unmarshaled["email"])
		assert.Equal(t, "50000", unmarshaled["amount"])
	})

	t.Run("omits nil fields in JSON", func(t *testing.T) {
		builder := NewCreateRequestBuilder("test@example.com", "50000")
		request := builder.Build()

		jsonData, err := json.Marshal(request)
		require.NoError(t, err, "should marshal to JSON without error")

		jsonString := string(jsonData)
		assert.NotContains(t, jsonString, "split_code", "should not contain nil split_code field")
		assert.NotContains(t, jsonString, "pin", "should not contain nil pin field")
		assert.NotContains(t, jsonString, "reference", "should not contain nil reference field")
	})
}

func TestPaymentMethodDetails(t *testing.T) {
	t.Run("bank details structure", func(t *testing.T) {
		bank := BankDetails{
			Code:          "057",
			AccountNumber: "0123456789",
		}

		assert.Equal(t, "057", bank.Code)
		assert.Equal(t, "0123456789", bank.AccountNumber)
	})

	t.Run("mobile money details structure", func(t *testing.T) {
		momo := MobileMoneyDetails{
			Phone:    "+2348123456789",
			Provider: "mtn",
		}

		assert.Equal(t, "+2348123456789", momo.Phone)
		assert.Equal(t, "mtn", momo.Provider)
	})

	t.Run("USSD details structure", func(t *testing.T) {
		ussd := USSDDetails{
			Type: "737",
		}

		assert.Equal(t, "737", ussd.Type)
	})
}

func TestCreateResponse_FieldByFieldValidation(t *testing.T) {
	// Read the create_200.json file
	responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "charge", "create_200.json")
	responseData, err := os.ReadFile(responseFilePath)
	require.NoError(t, err, "failed to read create_200.json")

	// Parse the raw JSON to get the original values
	var rawResponse map[string]any
	err = json.Unmarshal(responseData, &rawResponse)
	require.NoError(t, err, "failed to unmarshal raw JSON response")

	// Deserialize into the CreateChargeResponse struct
	var response CreateChargeResponse
	err = json.Unmarshal(responseData, &response)
	require.NoError(t, err, "failed to unmarshal into CreateChargeResponse struct")

	// Validate top-level fields against the raw JSON
	assert.Equal(t, rawResponse["status"], response.Status.Bool(), "status field should match")
	assert.Equal(t, rawResponse["message"], response.Message, "message field should match")

	// Validate data object fields
	rawData := rawResponse["data"].(map[string]any)
	assert.Equal(t, int(rawData["amount"].(float64)), response.Data.Amount, "amount should match")
	assert.Equal(t, rawData["currency"], string(response.Data.Currency), "currency should match")

	// For timestamp comparison, parse both and compare the actual time values
	expectedTime, err := time.Parse(time.RFC3339, rawData["transaction_date"].(string))
	require.NoError(t, err, "should parse expected transaction_date")
	actualTime, err := time.Parse(time.RFC3339, response.Data.TransactionDate.String())
	require.NoError(t, err, "should parse actual transaction_date")
	assert.True(t, expectedTime.Equal(actualTime), "transaction_date should represent the same moment")

	assert.Equal(t, rawData["status"], response.Data.Status, "data status should match")
	assert.Equal(t, rawData["reference"], response.Data.Reference, "reference should match")
	assert.Equal(t, rawData["domain"], response.Data.Domain, "domain should match")
	assert.Equal(t, rawData["gateway_response"], response.Data.GatewayResponse, "gateway_response should match")

	// Handle null message field
	if rawData["message"] == nil {
		assert.Empty(t, response.Data.Message, "data message should be empty for null")
	} else {
		assert.Equal(t, rawData["message"], response.Data.Message, "data message should match")
	}

	assert.Equal(t, rawData["channel"], string(response.Data.Channel), "channel should match")
	assert.Equal(t, rawData["ip_address"], response.Data.IPAddress, "ip_address should match")

	// Handle null log field
	if rawData["log"] == nil {
		assert.Nil(t, response.Data.Log, "log should be nil")
	} else {
		assert.Equal(t, rawData["log"], response.Data.Log, "log should match")
	}

	assert.Equal(t, int(rawData["fees"].(float64)), response.Data.Fees, "fees should match")

	// Handle null plan field
	if rawData["plan"] == nil {
		assert.Nil(t, response.Data.Plan, "plan should be nil")
	} else {
		assert.Equal(t, rawData["plan"], response.Data.Plan, "plan should match")
	}

	// Validate metadata object (comparing content, not pointer)
	rawMetadata := rawData["metadata"].(map[string]any)
	assert.Equal(t, rawMetadata, map[string]any(*response.Data.Metadata), "metadata content should match")

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
	assert.Equal(t, rawAuthorization["signature"], response.Data.Authorization.Signature, "signature should match")

	// Handle account_name pointer field
	if rawAuthorization["account_name"] != nil {
		assert.Equal(t, rawAuthorization["account_name"], *response.Data.Authorization.AccountName, "account_name should match")
	} else {
		assert.Nil(t, response.Data.Authorization.AccountName, "account_name should be nil")
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

	var roundTripResponse CreateChargeResponse
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
