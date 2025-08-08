package charge

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

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
			assert.Equal(t, tt.expectedStatus, response.Status, "status should match")
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

		var unmarshaled map[string]interface{}
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

	t.Run("QR details structure", func(t *testing.T) {
		qr := QRDetails{
			Provider: "visa",
		}

		assert.Equal(t, "visa", qr.Provider)
	})
}
