package customers

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRiskActionResponse_JSONDeserialization(t *testing.T) {
	// Read the whitelist_blacklist_200.json file as it contains risk_action field
	responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "customers", "whitelist_blacklist_200.json")
	responseData, err := os.ReadFile(responseFilePath)
	require.NoError(t, err, "failed to read whitelist_blacklist_200.json")

	var response RiskActionResponse
	err = json.Unmarshal(responseData, &response)
	require.NoError(t, err, "failed to unmarshal risk action response")

	// Basic validations
	assert.True(t, response.Status.Bool())
	assert.Equal(t, "Customer updated", response.Message)

	// Validate customer data is present
	require.NotNil(t, response.Data)
}

func TestRiskActionRequestBuilder(t *testing.T) {
	t.Run("builds request with customer email only", func(t *testing.T) {
		builder := NewRiskActionRequestBuilder("customer@email.com")
		req := builder.Build()

		assert.Equal(t, "customer@email.com", req.Customer)
		assert.Nil(t, req.RiskAction, "risk_action should be nil when not set")
	})

	t.Run("builds request with customer code only", func(t *testing.T) {
		builder := NewRiskActionRequestBuilder("CUS_xr58yrr2ujlft9k")
		req := builder.Build()

		assert.Equal(t, "CUS_xr58yrr2ujlft9k", req.Customer)
		assert.Nil(t, req.RiskAction, "risk_action should be nil when not set")
	})

	t.Run("builds request with allow risk action", func(t *testing.T) {
		builder := NewRiskActionRequestBuilder("customer@email.com")
		builder.RiskAction(RiskActionAllow)
		req := builder.Build()

		assert.Equal(t, "customer@email.com", req.Customer)
		require.NotNil(t, req.RiskAction)
		assert.Equal(t, RiskActionAllow, *req.RiskAction)
	})

	t.Run("builds request with deny risk action", func(t *testing.T) {
		builder := NewRiskActionRequestBuilder("CUS_test123")
		builder.RiskAction(RiskActionDeny)
		req := builder.Build()

		assert.Equal(t, "CUS_test123", req.Customer)
		require.NotNil(t, req.RiskAction)
		assert.Equal(t, RiskActionDeny, *req.RiskAction)
	})

	t.Run("builds request with default risk action", func(t *testing.T) {
		builder := NewRiskActionRequestBuilder("customer@email.com")
		builder.RiskAction(RiskActionDefault)
		req := builder.Build()

		assert.Equal(t, "customer@email.com", req.Customer)
		require.NotNil(t, req.RiskAction)
		assert.Equal(t, RiskActionDefault, *req.RiskAction)
	})

	t.Run("builder allows method chaining", func(t *testing.T) {
		req := NewRiskActionRequestBuilder("customer@email.com").
			RiskAction(RiskActionAllow).
			Build()

		assert.Equal(t, "customer@email.com", req.Customer)
		require.NotNil(t, req.RiskAction)
		assert.Equal(t, RiskActionAllow, *req.RiskAction)
	})
}

func TestRiskActionRequest_JSONSerialization(t *testing.T) {
	t.Run("serializes request with customer only", func(t *testing.T) {
		builder := NewRiskActionRequestBuilder("customer@email.com")
		req := builder.Build()

		jsonData, err := json.Marshal(req)
		require.NoError(t, err, "should marshal without error")

		var unmarshaled map[string]any
		err = json.Unmarshal(jsonData, &unmarshaled)
		require.NoError(t, err, "should unmarshal JSON without error")

		assert.Equal(t, "customer@email.com", unmarshaled["customer"], "customer should match")
		// risk_action should be null when not set (Go marshals nil pointers as null)
		assert.Nil(t, unmarshaled["risk_action"], "risk_action should be null when nil")
	})

	t.Run("serializes request with risk action", func(t *testing.T) {
		builder := NewRiskActionRequestBuilder("CUS_test123")
		builder.RiskAction(RiskActionDeny)
		req := builder.Build()

		jsonData, err := json.Marshal(req)
		require.NoError(t, err, "should marshal without error")

		var unmarshaled map[string]any
		err = json.Unmarshal(jsonData, &unmarshaled)
		require.NoError(t, err, "should unmarshal JSON without error")

		assert.Equal(t, "CUS_test123", unmarshaled["customer"], "customer should match")
		assert.Equal(t, "deny", unmarshaled["risk_action"], "risk_action should match")
	})
}

func TestRiskActionResponse_FieldByFieldValidation(t *testing.T) {
	// Read the whitelist_blacklist_200.json file as it contains risk_action field
	responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "customers", "whitelist_blacklist_200.json")
	responseData, err := os.ReadFile(responseFilePath)
	require.NoError(t, err, "failed to read whitelist_blacklist_200.json")

	// Parse the raw JSON to get the original values
	var rawResponse map[string]any
	err = json.Unmarshal(responseData, &rawResponse)
	require.NoError(t, err, "failed to unmarshal raw JSON response")

	// Deserialize into the RiskActionResponse struct
	var response RiskActionResponse
	err = json.Unmarshal(responseData, &response)
	require.NoError(t, err, "failed to unmarshal into RiskActionResponse struct")

	// Validate top-level fields against the raw JSON
	assert.Equal(t, rawResponse["status"], response.Status.Bool(), "status field should match")
	assert.Equal(t, rawResponse["message"], response.Message, "message field should match")

	// Since response.Data is types.Customer struct, we validate its fields directly
	rawData := rawResponse["data"].(map[string]any)
	customer := response.Data

	// Validate key customer fields - handling NullString fields correctly
	if customer.FirstName.Valid {
		assert.Equal(t, rawData["first_name"], customer.FirstName.String(), "first_name should match")
	}
	if customer.LastName.Valid {
		assert.Equal(t, rawData["last_name"], customer.LastName.String(), "last_name should match")
	}
	assert.Equal(t, rawData["email"], customer.Email.String(), "email should match")
	// phone is null in JSON
	if rawData["phone"] == nil {
		assert.False(t, customer.Phone.Valid, "phone should be invalid when null in JSON")
	} else {
		assert.Equal(t, rawData["phone"], customer.Phone.String(), "phone should match")
	}
	assert.Equal(t, rawData["customer_code"], customer.CustomerCode.String(), "customer_code should match")
	assert.Equal(t, rawData["risk_action"], customer.RiskAction.String(), "risk_action should match")
	assert.Equal(t, rawData["id"], float64(customer.ID.Uint64()), "id should match")
	if customer.Integration.Valid {
		assert.Equal(t, rawData["integration"], float64(customer.Integration.Int), "integration should match")
	}
	assert.Equal(t, rawData["domain"], customer.Domain.String(), "domain should match")
	assert.Equal(t, rawData["identified"], customer.Identified.Bool(), "identified should match")
	// identifications is null in JSON
	if rawData["identifications"] == nil {
		assert.False(t, customer.Identifications.Valid, "identifications should be invalid when null in JSON")
	} else {
		assert.True(t, customer.Identifications.Valid, "identifications should be valid")
		assert.Equal(t, rawData["identifications"], map[string]any(customer.Identifications.Metadata), "identifications should match")
	}

	// Validate timestamp fields using MultiDateTime
	createdAtStr, ok := rawData["createdAt"].(string)
	require.True(t, ok, "createdAt should be a string")
	parsedCreatedAt, err := time.Parse("2006-01-02T15:04:05.000Z", createdAtStr)
	require.NoError(t, err, "should parse createdAt timestamp")
	assert.Equal(t, 2016, parsedCreatedAt.Year(), "createdAt year should be 2016")
	assert.Equal(t, 2016, customer.CreatedAt.Time().Year(), "customer CreatedAt year should match")

	updatedAtStr, ok := rawData["updatedAt"].(string)
	require.True(t, ok, "updatedAt should be a string")
	parsedUpdatedAt, err := time.Parse("2006-01-02T15:04:05.000Z", updatedAtStr)
	require.NoError(t, err, "should parse updatedAt timestamp")
	assert.Equal(t, 2016, parsedUpdatedAt.Year(), "updatedAt year should be 2016")
	assert.Equal(t, 2016, customer.UpdatedAt.Time().Year(), "customer UpdatedAt year should match")

	// Validate metadata field (should be empty object)
	metadata := rawData["metadata"].(map[string]any)
	assert.Empty(t, metadata, "metadata should be empty")

	// Test round-trip serialization
	serialized, err := json.Marshal(response)
	require.NoError(t, err, "failed to marshal response back to JSON")

	var roundTripResponse RiskActionResponse
	err = json.Unmarshal(serialized, &roundTripResponse)
	require.NoError(t, err, "failed to unmarshal round-trip JSON")

	// Verify core fields survive round-trip
	assert.Equal(t, response.Status.Bool(), roundTripResponse.Status.Bool(), "status should survive round-trip")
	assert.Equal(t, response.Message, roundTripResponse.Message, "message should survive round-trip")
}

func TestRiskActionConstants(t *testing.T) {
	t.Run("risk action constants have correct values", func(t *testing.T) {
		assert.Equal(t, "default", string(RiskActionDefault))
		assert.Equal(t, "allow", string(RiskActionAllow))
		assert.Equal(t, "deny", string(RiskActionDeny))
	})

	t.Run("risk action constants are different", func(t *testing.T) {
		assert.NotEqual(t, RiskActionDefault, RiskActionAllow)
		assert.NotEqual(t, RiskActionDefault, RiskActionDeny)
		assert.NotEqual(t, RiskActionAllow, RiskActionDeny)
	})
}
