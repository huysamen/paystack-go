package customers

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInitializeAuthorizationResponse_JSONDeserialization(t *testing.T) {
	// Read the initialize_authorization_200.json file
	responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "customers", "initialize_authorization_200.json")
	responseData, err := os.ReadFile(responseFilePath)
	require.NoError(t, err, "failed to read initialize_authorization_200.json")

	var response InitializeAuthorizationResponse
	err = json.Unmarshal(responseData, &response)
	require.NoError(t, err, "failed to unmarshal initialize authorization response")

	// Basic validations
	assert.True(t, response.Status.Bool())
	assert.Equal(t, "Authorization initialized", response.Message)
	assert.NotNil(t, response.Data)

	// Validate response data
	assert.Equal(t, "https://checkout.paystack.com/82t4mp5b5mfn51h", response.Data.RedirectURL)
	assert.Equal(t, "82t4mp5b5mfn51h", response.Data.AccessCode)
	assert.Equal(t, "dfbzfotsrbv4n5s82t4mp5b5mfn51h", response.Data.Reference)
}

func TestInitializeAuthorizationRequestBuilder(t *testing.T) {
	t.Run("builds basic request with required fields", func(t *testing.T) {
		builder := NewInitializeAuthorizationRequestBuilder("test@example.com", "direct_debit")
		req := builder.Build()

		assert.Equal(t, "test@example.com", req.Email)
		assert.Equal(t, "direct_debit", req.Channel)
		assert.Nil(t, req.CallbackURL)
		assert.Nil(t, req.Account)
		assert.Nil(t, req.Address)
	})

	t.Run("builds request with callback URL", func(t *testing.T) {
		builder := NewInitializeAuthorizationRequestBuilder("test@example.com", "direct_debit").
			CallbackURL("https://example.com/callback")

		req := builder.Build()

		assert.Equal(t, "test@example.com", req.Email)
		assert.Equal(t, "direct_debit", req.Channel)
		assert.Equal(t, "https://example.com/callback", *req.CallbackURL)
		assert.Nil(t, req.Account)
		assert.Nil(t, req.Address)
	})

	t.Run("builds request with account details", func(t *testing.T) {
		builder := NewInitializeAuthorizationRequestBuilder("test@example.com", "direct_debit").
			Account("0123456789", "044")

		req := builder.Build()

		assert.Equal(t, "test@example.com", req.Email)
		assert.Equal(t, "direct_debit", req.Channel)
		assert.NotNil(t, req.Account)
		assert.Equal(t, "0123456789", req.Account.Number)
		assert.Equal(t, "044", req.Account.BankCode)
		assert.Nil(t, req.Address)
	})

	t.Run("builds request with address", func(t *testing.T) {
		builder := NewInitializeAuthorizationRequestBuilder("test@example.com", "direct_debit").
			Address("123 Main St", "Lagos", "Lagos")

		req := builder.Build()

		assert.Equal(t, "test@example.com", req.Email)
		assert.Equal(t, "direct_debit", req.Channel)
		assert.NotNil(t, req.Address)
		assert.Equal(t, "123 Main St", req.Address.Street)
		assert.Equal(t, "Lagos", req.Address.City)
		assert.Equal(t, "Lagos", req.Address.State)
	})

	t.Run("builds request with all optional fields", func(t *testing.T) {
		builder := NewInitializeAuthorizationRequestBuilder("test@example.com", "direct_debit").
			CallbackURL("https://example.com/callback").
			Account("0123456789", "044").
			Address("123 Main St", "Lagos", "Lagos")

		req := builder.Build()

		assert.Equal(t, "test@example.com", req.Email)
		assert.Equal(t, "direct_debit", req.Channel)
		assert.Equal(t, "https://example.com/callback", *req.CallbackURL)
		assert.NotNil(t, req.Account)
		assert.Equal(t, "0123456789", req.Account.Number)
		assert.Equal(t, "044", req.Account.BankCode)
		assert.NotNil(t, req.Address)
		assert.Equal(t, "123 Main St", req.Address.Street)
		assert.Equal(t, "Lagos", req.Address.City)
		assert.Equal(t, "Lagos", req.Address.State)
	})
}

func TestInitializeAuthorizationRequest_JSONSerialization(t *testing.T) {
	t.Run("serializes basic request correctly", func(t *testing.T) {
		builder := NewInitializeAuthorizationRequestBuilder("test@example.com", "direct_debit")
		req := builder.Build()

		jsonData, err := json.Marshal(req)
		require.NoError(t, err, "should marshal without error")

		var unmarshaled map[string]any
		err = json.Unmarshal(jsonData, &unmarshaled)
		require.NoError(t, err, "should unmarshal JSON without error")

		assert.Equal(t, "test@example.com", unmarshaled["email"], "email should match")
		assert.Equal(t, "direct_debit", unmarshaled["channel"], "channel should match")

		// Optional fields should not be present when nil
		_, hasCallbackURL := unmarshaled["callback_url"]
		_, hasAccount := unmarshaled["account"]
		_, hasAddress := unmarshaled["address"]
		assert.False(t, hasCallbackURL, "callback_url should be omitted when nil")
		assert.False(t, hasAccount, "account should be omitted when nil")
		assert.False(t, hasAddress, "address should be omitted when nil")
	})

	t.Run("includes all fields when provided", func(t *testing.T) {
		builder := NewInitializeAuthorizationRequestBuilder("test@example.com", "direct_debit").
			CallbackURL("https://example.com/callback").
			Account("0123456789", "044").
			Address("123 Main St", "Lagos", "Lagos")

		req := builder.Build()

		jsonData, err := json.Marshal(req)
		require.NoError(t, err, "should marshal without error")

		var unmarshaled map[string]any
		err = json.Unmarshal(jsonData, &unmarshaled)
		require.NoError(t, err, "should unmarshal JSON without error")

		assert.Equal(t, "test@example.com", unmarshaled["email"])
		assert.Equal(t, "direct_debit", unmarshaled["channel"])
		assert.Equal(t, "https://example.com/callback", unmarshaled["callback_url"])

		account := unmarshaled["account"].(map[string]any)
		assert.Equal(t, "0123456789", account["number"])
		assert.Equal(t, "044", account["bank_code"])

		address := unmarshaled["address"].(map[string]any)
		assert.Equal(t, "123 Main St", address["street"])
		assert.Equal(t, "Lagos", address["city"])
		assert.Equal(t, "Lagos", address["state"])
	})
}

func TestInitializeAuthorizationResponse_FieldByFieldValidation(t *testing.T) {
	// Read the initialize_authorization_200.json file
	responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "customers", "initialize_authorization_200.json")
	responseData, err := os.ReadFile(responseFilePath)
	require.NoError(t, err, "failed to read initialize_authorization_200.json")

	// Parse the raw JSON to get the original values
	var rawResponse map[string]any
	err = json.Unmarshal(responseData, &rawResponse)
	require.NoError(t, err, "failed to unmarshal raw JSON response")

	// Deserialize into the InitializeAuthorizationResponse struct
	var response InitializeAuthorizationResponse
	err = json.Unmarshal(responseData, &response)
	require.NoError(t, err, "failed to unmarshal into InitializeAuthorizationResponse struct")

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

	var roundTripResponse InitializeAuthorizationResponse
	err = json.Unmarshal(serialized, &roundTripResponse)
	require.NoError(t, err, "failed to unmarshal round-trip JSON")

	// Verify core fields survive round-trip
	assert.Equal(t, response.Status.Bool(), roundTripResponse.Status.Bool(), "status should survive round-trip")
	assert.Equal(t, response.Message, roundTripResponse.Message, "message should survive round-trip")
	assert.Equal(t, response.Data.RedirectURL, roundTripResponse.Data.RedirectURL, "redirect_url should survive round-trip")
	assert.Equal(t, response.Data.AccessCode, roundTripResponse.Data.AccessCode, "access_code should survive round-trip")
	assert.Equal(t, response.Data.Reference, roundTripResponse.Data.Reference, "reference should survive round-trip")
}
