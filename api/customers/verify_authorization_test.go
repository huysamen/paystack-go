package customers

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/huysamen/paystack-go/enums"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestVerifyAuthorizationResponse_JSONDeserialization(t *testing.T) {
	t.Run("deserializes success response (200)", func(t *testing.T) {
		responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "customers", "verify_authorization_200.json")
		responseData, err := os.ReadFile(responseFilePath)
		require.NoError(t, err, "failed to read verify_authorization_200.json")

		var response VerifyAuthorizationResponse
		err = json.Unmarshal(responseData, &response)
		require.NoError(t, err, "failed to unmarshal verify authorization response")

		// Basic validations
		assert.True(t, response.Status.Bool())
		assert.Equal(t, "Authorization retrieved successfully", response.Message)

		// Validate data structure
		require.NotNil(t, response.Data)
		assert.Equal(t, "AUTH_JV4T9Wawdj", response.Data.AuthorizationCode)
		assert.Equal(t, "direct_debit", response.Data.Channel)
		assert.Equal(t, "Guaranty Trust Bank", response.Data.Bank)
		assert.True(t, response.Data.Active)
		assert.Equal(t, "CUS_24lze1c8i2zl76y", response.Data.Customer.Code)
		assert.Equal(t, "ravi@demo.com", response.Data.Customer.Email)
	})

	t.Run("deserializes error response (404)", func(t *testing.T) {
		responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "customers", "verify_authorization_404.json")
		responseData, err := os.ReadFile(responseFilePath)
		require.NoError(t, err, "failed to read verify_authorization_404.json")

		var response VerifyAuthorizationResponse
		err = json.Unmarshal(responseData, &response)
		require.NoError(t, err, "failed to unmarshal verify authorization error response")

		// Basic validations
		assert.False(t, response.Status.Bool())
		assert.Equal(t, "Authorization does not exist or does not belong to integration", response.Message)
	})
}

func TestVerifyAuthorizationResponse_FieldByFieldValidation(t *testing.T) {
	t.Run("validates success response fields", func(t *testing.T) {
		responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "customers", "verify_authorization_200.json")
		responseData, err := os.ReadFile(responseFilePath)
		require.NoError(t, err, "failed to read verify_authorization_200.json")

		// Parse the raw JSON to get the original values
		var rawResponse map[string]any
		err = json.Unmarshal(responseData, &rawResponse)
		require.NoError(t, err, "failed to unmarshal raw JSON response")

		// Deserialize into the VerifyAuthorizationResponse struct
		var response VerifyAuthorizationResponse
		err = json.Unmarshal(responseData, &response)
		require.NoError(t, err, "failed to unmarshal into VerifyAuthorizationResponse struct")

		// Validate top-level fields against the raw JSON
		assert.Equal(t, rawResponse["status"], response.Status.Bool(), "status field should match")
		assert.Equal(t, rawResponse["message"], response.Message, "message field should match")

		// Validate data object fields
		rawData := rawResponse["data"].(map[string]any)
		assert.Equal(t, rawData["authorization_code"], response.Data.AuthorizationCode, "authorization_code should match")
		assert.Equal(t, rawData["channel"], response.Data.Channel, "channel should match")
		assert.Equal(t, rawData["bank"], response.Data.Bank, "bank should match")
		assert.Equal(t, rawData["active"], response.Data.Active, "active should match")

		// Validate nested customer object
		rawCustomer := rawData["customer"].(map[string]any)
		assert.Equal(t, rawCustomer["code"], response.Data.Customer.Code, "customer.code should match")
		assert.Equal(t, rawCustomer["email"], response.Data.Customer.Email, "customer.email should match")

		// Test round-trip serialization
		serialized, err := json.Marshal(response)
		require.NoError(t, err, "failed to marshal response back to JSON")

		var roundTripResponse VerifyAuthorizationResponse
		err = json.Unmarshal(serialized, &roundTripResponse)
		require.NoError(t, err, "failed to unmarshal round-trip JSON")

		// Verify core fields survive round-trip
		assert.Equal(t, response.Status.Bool(), roundTripResponse.Status.Bool(), "status should survive round-trip")
		assert.Equal(t, response.Message, roundTripResponse.Message, "message should survive round-trip")
		assert.Equal(t, response.Data.AuthorizationCode, roundTripResponse.Data.AuthorizationCode, "authorization_code should survive round-trip")
		assert.Equal(t, response.Data.Channel, roundTripResponse.Data.Channel, "channel should survive round-trip")
		assert.Equal(t, response.Data.Bank, roundTripResponse.Data.Bank, "bank should survive round-trip")
		assert.Equal(t, response.Data.Active, roundTripResponse.Data.Active, "active should survive round-trip")
		assert.Equal(t, response.Data.Customer.Code, roundTripResponse.Data.Customer.Code, "customer.code should survive round-trip")
		assert.Equal(t, response.Data.Customer.Email, roundTripResponse.Data.Customer.Email, "customer.email should survive round-trip")
	})

	t.Run("validates error response fields", func(t *testing.T) {
		responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "customers", "verify_authorization_404.json")
		responseData, err := os.ReadFile(responseFilePath)
		require.NoError(t, err, "failed to read verify_authorization_404.json")

		// Parse the raw JSON to get the original values
		var rawResponse map[string]any
		err = json.Unmarshal(responseData, &rawResponse)
		require.NoError(t, err, "failed to unmarshal raw JSON response")

		// Deserialize into the VerifyAuthorizationResponse struct
		var response VerifyAuthorizationResponse
		err = json.Unmarshal(responseData, &response)
		require.NoError(t, err, "failed to unmarshal into VerifyAuthorizationResponse struct")

		// Validate top-level fields against the raw JSON
		assert.Equal(t, rawResponse["status"], response.Status.Bool(), "status field should match")
		assert.Equal(t, rawResponse["message"], response.Message, "message field should match")

		// Data field should not exist in error response
		_, hasData := rawResponse["data"]
		assert.False(t, hasData, "data field should not exist in error response")
	})
}

func TestCustomerReference_JSONSerialization(t *testing.T) {
	t.Run("serializes customer reference correctly", func(t *testing.T) {
		customerRef := CustomerReference{
			Code:  "CUS_24lze1c8i2zl76y",
			Email: "ravi@demo.com",
		}

		jsonData, err := json.Marshal(customerRef)
		require.NoError(t, err, "should marshal CustomerReference without error")

		var unmarshaled map[string]any
		err = json.Unmarshal(jsonData, &unmarshaled)
		require.NoError(t, err, "should unmarshal CustomerReference JSON without error")

		assert.Equal(t, "CUS_24lze1c8i2zl76y", unmarshaled["code"], "code should match")
		assert.Equal(t, "ravi@demo.com", unmarshaled["email"], "email should match")
	})

	t.Run("deserializes customer reference correctly", func(t *testing.T) {
		jsonData := `{"code":"CUS_test123","email":"test@example.com"}`

		var customerRef CustomerReference
		err := json.Unmarshal([]byte(jsonData), &customerRef)
		require.NoError(t, err, "should unmarshal CustomerReference without error")

		assert.Equal(t, "CUS_test123", customerRef.Code)
		assert.Equal(t, "test@example.com", customerRef.Email)
	})
}

func TestVerifyAuthorizationResponseData_JSONSerialization(t *testing.T) {
	t.Run("serializes verify authorization response data correctly", func(t *testing.T) {
		responseData := verifyAuthorizationResponseData{
			AuthorizationCode: "AUTH_JV4T9Wawdj",
			Channel:           "direct_debit",
			Bank:              "Guaranty Trust Bank",
			Active:            true,
			Customer: CustomerReference{
				Code:  "CUS_24lze1c8i2zl76y",
				Email: "ravi@demo.com",
			},
		}

		jsonData, err := json.Marshal(responseData)
		require.NoError(t, err, "should marshal verifyAuthorizationResponseData without error")

		var unmarshaled map[string]any
		err = json.Unmarshal(jsonData, &unmarshaled)
		require.NoError(t, err, "should unmarshal response data JSON without error")

		assert.Equal(t, "AUTH_JV4T9Wawdj", unmarshaled["authorization_code"], "authorization_code should match")
		assert.Equal(t, "direct_debit", unmarshaled["channel"], "channel should match")
		assert.Equal(t, "Guaranty Trust Bank", unmarshaled["bank"], "bank should match")
		assert.Equal(t, true, unmarshaled["active"], "active should match")

		// Validate nested customer object
		customer := unmarshaled["customer"].(map[string]any)
		assert.Equal(t, "CUS_24lze1c8i2zl76y", customer["code"], "customer.code should match")
		assert.Equal(t, "ravi@demo.com", customer["email"], "customer.email should match")
	})
}

func TestChannelHandling_Integration(t *testing.T) {
	t.Run("handles direct_debit channel as string", func(t *testing.T) {
		// Test that direct_debit channel works as a plain string
		channel := "direct_debit"
		assert.Equal(t, "direct_debit", channel)

		// Test JSON marshaling with string channel
		jsonData, err := json.Marshal(channel)
		require.NoError(t, err, "should marshal channel string without error")
		assert.Equal(t, `"direct_debit"`, string(jsonData), "channel should marshal to correct JSON string")

		// Test JSON unmarshaling with string channel
		var unmarshaledChannel string
		err = json.Unmarshal([]byte(`"direct_debit"`), &unmarshaledChannel)
		require.NoError(t, err, "should unmarshal channel string without error")
		assert.Equal(t, "direct_debit", unmarshaledChannel, "channel should unmarshal correctly")
	})

	t.Run("demonstrates enum limitation", func(t *testing.T) {
		// This test shows that the current Channel enum doesn't support direct_debit
		// and would need to be updated to handle all channel types returned by the API
		var channel enums.Channel
		err := json.Unmarshal([]byte(`"direct_debit"`), &channel)
		assert.Error(t, err, "should fail to unmarshal unknown channel value with current enum")

		// Test with a known valid channel
		err = json.Unmarshal([]byte(`"card"`), &channel)
		require.NoError(t, err, "should unmarshal valid channel enum without error")
		assert.Equal(t, enums.ChannelCard, channel, "known channel should unmarshal correctly")
	})
}
