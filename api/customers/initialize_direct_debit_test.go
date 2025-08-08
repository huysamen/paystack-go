package customers

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInitializeDirectDebitResponse_JSONDeserialization(t *testing.T) {
	// Read the initialize_direct_debit_200.json file
	responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "customers", "initialize_direct_debit_200.json")
	responseData, err := os.ReadFile(responseFilePath)
	require.NoError(t, err, "failed to read initialize_direct_debit_200.json")

	var response InitializeDirectDebitResponse
	err = json.Unmarshal(responseData, &response)
	require.NoError(t, err, "failed to unmarshal initialize direct debit response")

	// Basic validations
	assert.True(t, response.Status.Bool())
	assert.Equal(t, "Authorization initialized", response.Message)

	// Validate data structure
	require.NotNil(t, response.Data)
	assert.Equal(t, "https://link.paystack.com/ll6b0szngj1f27k", response.Data.RedirectURL)
	assert.Equal(t, "ll6b0szngj1f27k", response.Data.AccessCode)
	assert.Equal(t, "1er945lpy4txyki", response.Data.Reference)
}

func TestInitializeDirectDebitRequestBuilder(t *testing.T) {
	t.Run("builds request with all fields", func(t *testing.T) {
		builder := NewInitializeDirectDebitRequestBuilder(
			"0123456789",
			"011",
			"123 Main St",
			"Lagos",
			"Lagos State",
		)
		req := builder.Build()

		assert.Equal(t, "0123456789", req.Account.Number)
		assert.Equal(t, "011", req.Account.BankCode)
		assert.Equal(t, "123 Main St", req.Address.Street)
		assert.Equal(t, "Lagos", req.Address.City)
		assert.Equal(t, "Lagos State", req.Address.State)
	})

	t.Run("builds request with empty fields", func(t *testing.T) {
		builder := NewInitializeDirectDebitRequestBuilder("", "", "", "", "")
		req := builder.Build()

		assert.Equal(t, "", req.Account.Number)
		assert.Equal(t, "", req.Account.BankCode)
		assert.Equal(t, "", req.Address.Street)
		assert.Equal(t, "", req.Address.City)
		assert.Equal(t, "", req.Address.State)
	})

	t.Run("builds request with partial fields", func(t *testing.T) {
		builder := NewInitializeDirectDebitRequestBuilder(
			"9876543210",
			"044",
			"",
			"Abuja",
			"",
		)
		req := builder.Build()

		assert.Equal(t, "9876543210", req.Account.Number)
		assert.Equal(t, "044", req.Account.BankCode)
		assert.Equal(t, "", req.Address.Street)
		assert.Equal(t, "Abuja", req.Address.City)
		assert.Equal(t, "", req.Address.State)
	})
}

func TestInitializeDirectDebitRequest_JSONSerialization(t *testing.T) {
	t.Run("serializes complete request correctly", func(t *testing.T) {
		builder := NewInitializeDirectDebitRequestBuilder(
			"0123456789",
			"011",
			"123 Main St",
			"Lagos",
			"Lagos State",
		)
		req := builder.Build()

		jsonData, err := json.Marshal(req)
		require.NoError(t, err, "should marshal without error")

		var unmarshaled map[string]any
		err = json.Unmarshal(jsonData, &unmarshaled)
		require.NoError(t, err, "should unmarshal JSON without error")

		// Validate account fields
		account := unmarshaled["account"].(map[string]any)
		assert.Equal(t, "0123456789", account["number"], "account number should match")
		assert.Equal(t, "011", account["bank_code"], "bank code should match")

		// Validate address fields
		address := unmarshaled["address"].(map[string]any)
		assert.Equal(t, "123 Main St", address["street"], "street should match")
		assert.Equal(t, "Lagos", address["city"], "city should match")
		assert.Equal(t, "Lagos State", address["state"], "state should match")
	})

	t.Run("serializes empty request correctly", func(t *testing.T) {
		builder := NewInitializeDirectDebitRequestBuilder("", "", "", "", "")
		req := builder.Build()

		jsonData, err := json.Marshal(req)
		require.NoError(t, err, "should marshal without error")

		var unmarshaled map[string]any
		err = json.Unmarshal(jsonData, &unmarshaled)
		require.NoError(t, err, "should unmarshal JSON without error")

		// Validate account fields are empty
		account := unmarshaled["account"].(map[string]any)
		assert.Equal(t, "", account["number"], "account number should be empty")
		assert.Equal(t, "", account["bank_code"], "bank code should be empty")

		// Validate address fields are empty
		address := unmarshaled["address"].(map[string]any)
		assert.Equal(t, "", address["street"], "street should be empty")
		assert.Equal(t, "", address["city"], "city should be empty")
		assert.Equal(t, "", address["state"], "state should be empty")
	})
}

func TestInitializeDirectDebitResponse_FieldByFieldValidation(t *testing.T) {
	// Read the initialize_direct_debit_200.json file
	responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "customers", "initialize_direct_debit_200.json")
	responseData, err := os.ReadFile(responseFilePath)
	require.NoError(t, err, "failed to read initialize_direct_debit_200.json")

	// Parse the raw JSON to get the original values
	var rawResponse map[string]any
	err = json.Unmarshal(responseData, &rawResponse)
	require.NoError(t, err, "failed to unmarshal raw JSON response")

	// Deserialize into the InitializeDirectDebitResponse struct
	var response InitializeDirectDebitResponse
	err = json.Unmarshal(responseData, &response)
	require.NoError(t, err, "failed to unmarshal into InitializeDirectDebitResponse struct")

	// Validate top-level fields against the raw JSON
	assert.Equal(t, rawResponse["status"], response.Status.Bool(), "status field should match")
	assert.Equal(t, rawResponse["message"], response.Message, "message field should match")

	// Validate data object fields
	rawData := rawResponse["data"].(map[string]any)
	assert.Equal(t, rawData["redirect_url"], response.Data.RedirectURL, "redirect_url should match")
	assert.Equal(t, rawData["access_code"], response.Data.AccessCode, "access_code should match")
	assert.Equal(t, rawData["reference"], response.Data.Reference, "reference should match")

	// Test round-trip serialization
	serialized, err := json.Marshal(response)
	require.NoError(t, err, "failed to marshal response back to JSON")

	var roundTripResponse InitializeDirectDebitResponse
	err = json.Unmarshal(serialized, &roundTripResponse)
	require.NoError(t, err, "failed to unmarshal round-trip JSON")

	// Verify core fields survive round-trip
	assert.Equal(t, response.Status.Bool(), roundTripResponse.Status.Bool(), "status should survive round-trip")
	assert.Equal(t, response.Message, roundTripResponse.Message, "message should survive round-trip")
	assert.Equal(t, response.Data.RedirectURL, roundTripResponse.Data.RedirectURL, "redirect_url should survive round-trip")
	assert.Equal(t, response.Data.AccessCode, roundTripResponse.Data.AccessCode, "access_code should survive round-trip")
	assert.Equal(t, response.Data.Reference, roundTripResponse.Data.Reference, "reference should survive round-trip")
}

func TestAccount_JSONSerialization(t *testing.T) {
	account := Account{
		Number:   "0123456789",
		BankCode: "011",
	}

	jsonData, err := json.Marshal(account)
	require.NoError(t, err, "should marshal Account without error")

	var unmarshaled map[string]any
	err = json.Unmarshal(jsonData, &unmarshaled)
	require.NoError(t, err, "should unmarshal Account JSON without error")

	assert.Equal(t, "0123456789", unmarshaled["number"], "number should match")
	assert.Equal(t, "011", unmarshaled["bank_code"], "bank_code should match")
}

func TestAddress_JSONSerialization(t *testing.T) {
	address := Address{
		Street: "123 Main St",
		City:   "Lagos",
		State:  "Lagos State",
	}

	jsonData, err := json.Marshal(address)
	require.NoError(t, err, "should marshal Address without error")

	var unmarshaled map[string]any
	err = json.Unmarshal(jsonData, &unmarshaled)
	require.NoError(t, err, "should unmarshal Address JSON without error")

	assert.Equal(t, "123 Main St", unmarshaled["street"], "street should match")
	assert.Equal(t, "Lagos", unmarshaled["city"], "city should match")
	assert.Equal(t, "Lagos State", unmarshaled["state"], "state should match")
}
