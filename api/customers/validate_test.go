package customers

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidateResponse_JSONDeserialization(t *testing.T) {
	t.Run("deserializes success response (202)", func(t *testing.T) {
		responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "customers", "validate_202.json")
		responseData, err := os.ReadFile(responseFilePath)
		require.NoError(t, err, "failed to read validate_202.json")

		var response CustomerValidateResponse
		err = json.Unmarshal(responseData, &response)
		require.NoError(t, err, "failed to unmarshal validate response")

		// Basic validations
		assert.True(t, response.Status.Bool())
		assert.Equal(t, "Customer Identification in progress", response.Message)
	})

	t.Run("deserializes error response (400)", func(t *testing.T) {
		responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "customers", "validate_400.json")
		responseData, err := os.ReadFile(responseFilePath)
		require.NoError(t, err, "failed to read validate_400.json")

		var response CustomerValidateResponse
		err = json.Unmarshal(responseData, &response)
		require.NoError(t, err, "failed to unmarshal validate error response")

		// Basic validations
		assert.False(t, response.Status.Bool())
		assert.Equal(t, "Invalid BVN", response.Message)
	})
}

func TestValidateRequestBuilder(t *testing.T) {
	t.Run("builds basic request with required fields", func(t *testing.T) {
		builder := NewValidateRequestBuilder(
			"John",
			"Doe",
			"bank_account",
			"BVN",
			"NG",
			"12345678901",
		)
		req := builder.Build()

		assert.Equal(t, "John", req.FirstName)
		assert.Equal(t, "Doe", req.LastName)
		assert.Equal(t, "bank_account", req.Type)
		assert.Equal(t, "BVN", req.Value)
		assert.Equal(t, "NG", req.Country)
		assert.Equal(t, "12345678901", req.BVN)
		assert.Empty(t, req.BankCode, "BankCode should be empty when not set")
		assert.Empty(t, req.AccountNumber, "AccountNumber should be empty when not set")
		assert.Nil(t, req.MiddleName, "MiddleName should be nil when not set")
	})

	t.Run("builds request with bank account details", func(t *testing.T) {
		builder := NewValidateRequestBuilder(
			"Jane",
			"Smith",
			"bank_account",
			"BVN",
			"NG",
			"98765432109",
		)
		builder.BankCode("044")
		builder.AccountNumber("0123456789")
		req := builder.Build()

		assert.Equal(t, "Jane", req.FirstName)
		assert.Equal(t, "Smith", req.LastName)
		assert.Equal(t, "bank_account", req.Type)
		assert.Equal(t, "BVN", req.Value)
		assert.Equal(t, "NG", req.Country)
		assert.Equal(t, "98765432109", req.BVN)
		assert.Equal(t, "044", req.BankCode)
		assert.Equal(t, "0123456789", req.AccountNumber)
		assert.Nil(t, req.MiddleName, "MiddleName should be nil when not set")
	})

	t.Run("builds request with middle name", func(t *testing.T) {
		builder := NewValidateRequestBuilder(
			"John",
			"Doe",
			"bank_account",
			"BVN",
			"NG",
			"12345678901",
		)
		builder.MiddleName("Michael")
		req := builder.Build()

		assert.Equal(t, "John", req.FirstName)
		assert.Equal(t, "Doe", req.LastName)
		require.NotNil(t, req.MiddleName)
		assert.Equal(t, "Michael", *req.MiddleName)
	})

	t.Run("builds complete request with all fields", func(t *testing.T) {
		req := NewValidateRequestBuilder(
			"Alice",
			"Johnson",
			"bank_account",
			"BVN",
			"NG",
			"11122233344",
		).BankCode("011").
			AccountNumber("9876543210").
			MiddleName("Marie").
			Build()

		assert.Equal(t, "Alice", req.FirstName)
		assert.Equal(t, "Johnson", req.LastName)
		assert.Equal(t, "bank_account", req.Type)
		assert.Equal(t, "BVN", req.Value)
		assert.Equal(t, "NG", req.Country)
		assert.Equal(t, "11122233344", req.BVN)
		assert.Equal(t, "011", req.BankCode)
		assert.Equal(t, "9876543210", req.AccountNumber)
		require.NotNil(t, req.MiddleName)
		assert.Equal(t, "Marie", *req.MiddleName)
	})

	t.Run("builds request with empty fields", func(t *testing.T) {
		builder := NewValidateRequestBuilder("", "", "", "", "", "")
		req := builder.Build()

		assert.Equal(t, "", req.FirstName)
		assert.Equal(t, "", req.LastName)
		assert.Equal(t, "", req.Type)
		assert.Equal(t, "", req.Value)
		assert.Equal(t, "", req.Country)
		assert.Equal(t, "", req.BVN)
	})
}

func TestValidateRequest_JSONSerialization(t *testing.T) {
	t.Run("serializes basic request correctly", func(t *testing.T) {
		builder := NewValidateRequestBuilder(
			"John",
			"Doe",
			"bank_account",
			"BVN",
			"NG",
			"12345678901",
		)
		req := builder.Build()

		jsonData, err := json.Marshal(req)
		require.NoError(t, err, "should marshal without error")

		var unmarshaled map[string]any
		err = json.Unmarshal(jsonData, &unmarshaled)
		require.NoError(t, err, "should unmarshal JSON without error")

		assert.Equal(t, "John", unmarshaled["first_name"], "first_name should match")
		assert.Equal(t, "Doe", unmarshaled["last_name"], "last_name should match")
		assert.Equal(t, "bank_account", unmarshaled["type"], "type should match")
		assert.Equal(t, "BVN", unmarshaled["value"], "value should match")
		assert.Equal(t, "NG", unmarshaled["country"], "country should match")
		assert.Equal(t, "12345678901", unmarshaled["bvn"], "bvn should match")
		assert.Equal(t, "", unmarshaled["bank_code"], "bank_code should be empty")
		assert.Equal(t, "", unmarshaled["account_number"], "account_number should be empty")
		// middle_name should not be present when nil (omitempty tag)
		_, hasMiddleName := unmarshaled["middle_name"]
		assert.False(t, hasMiddleName, "middle_name should not be present when nil")
	})

	t.Run("serializes complete request correctly", func(t *testing.T) {
		req := NewValidateRequestBuilder(
			"Alice",
			"Johnson",
			"bank_account",
			"BVN",
			"NG",
			"11122233344",
		).BankCode("011").
			AccountNumber("9876543210").
			MiddleName("Marie").
			Build()

		jsonData, err := json.Marshal(req)
		require.NoError(t, err, "should marshal without error")

		var unmarshaled map[string]any
		err = json.Unmarshal(jsonData, &unmarshaled)
		require.NoError(t, err, "should unmarshal JSON without error")

		assert.Equal(t, "Alice", unmarshaled["first_name"], "first_name should match")
		assert.Equal(t, "Johnson", unmarshaled["last_name"], "last_name should match")
		assert.Equal(t, "bank_account", unmarshaled["type"], "type should match")
		assert.Equal(t, "BVN", unmarshaled["value"], "value should match")
		assert.Equal(t, "NG", unmarshaled["country"], "country should match")
		assert.Equal(t, "11122233344", unmarshaled["bvn"], "bvn should match")
		assert.Equal(t, "011", unmarshaled["bank_code"], "bank_code should match")
		assert.Equal(t, "9876543210", unmarshaled["account_number"], "account_number should match")
		assert.Equal(t, "Marie", unmarshaled["middle_name"], "middle_name should match")
	})
}

func TestValidateResponse_FieldByFieldValidation(t *testing.T) {
	t.Run("validates success response fields", func(t *testing.T) {
		responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "customers", "validate_202.json")
		responseData, err := os.ReadFile(responseFilePath)
		require.NoError(t, err, "failed to read validate_202.json")

		// Parse the raw JSON to get the original values
		var rawResponse map[string]any
		err = json.Unmarshal(responseData, &rawResponse)
		require.NoError(t, err, "failed to unmarshal raw JSON response")

		// Deserialize into the CustomerValidateResponse struct
		var response CustomerValidateResponse
		err = json.Unmarshal(responseData, &response)
		require.NoError(t, err, "failed to unmarshal into CustomerValidateResponse struct")

		// Validate top-level fields against the raw JSON
		assert.Equal(t, rawResponse["status"], response.Status.Bool(), "status field should match")
		assert.Equal(t, rawResponse["message"], response.Message, "message field should match")

		// Data field should not exist in this response
		_, hasData := rawResponse["data"]
		assert.False(t, hasData, "data field should not exist in validate response")

		// Test round-trip serialization
		serialized, err := json.Marshal(response)
		require.NoError(t, err, "failed to marshal response back to JSON")

		var roundTripResponse CustomerValidateResponse
		err = json.Unmarshal(serialized, &roundTripResponse)
		require.NoError(t, err, "failed to unmarshal round-trip JSON")

		// Verify core fields survive round-trip
		assert.Equal(t, response.Status.Bool(), roundTripResponse.Status.Bool(), "status should survive round-trip")
		assert.Equal(t, response.Message, roundTripResponse.Message, "message should survive round-trip")
	})

	t.Run("validates error response fields", func(t *testing.T) {
		responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "customers", "validate_400.json")
		responseData, err := os.ReadFile(responseFilePath)
		require.NoError(t, err, "failed to read validate_400.json")

		// Parse the raw JSON to get the original values
		var rawResponse map[string]any
		err = json.Unmarshal(responseData, &rawResponse)
		require.NoError(t, err, "failed to unmarshal raw JSON response")

		// Deserialize into the CustomerValidateResponse struct
		var response CustomerValidateResponse
		err = json.Unmarshal(responseData, &response)
		require.NoError(t, err, "failed to unmarshal into CustomerValidateResponse struct")

		// Validate top-level fields against the raw JSON
		assert.Equal(t, rawResponse["status"], response.Status.Bool(), "status field should match")
		assert.Equal(t, rawResponse["message"], response.Message, "message field should match")

		// Data field should not exist in error response
		_, hasData := rawResponse["data"]
		assert.False(t, hasData, "data field should not exist in error response")
	})
}
