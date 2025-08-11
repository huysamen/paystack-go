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

func TestFetchMandateAuthorizationsResponse_JSONDeserialization(t *testing.T) {
	// Read the fetch_mandate_authorizations_200.json file
	responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "customers", "fetch_mandate_authorizations_200.json")
	responseData, err := os.ReadFile(responseFilePath)
	require.NoError(t, err, "failed to read fetch_mandate_authorizations_200.json")

	var response FetchMandateAuthorizationsResponse
	err = json.Unmarshal(responseData, &response)
	require.NoError(t, err, "failed to unmarshal fetch mandate authorizations response")

	// Basic validations
	assert.True(t, response.Status.Bool())
	assert.Equal(t, "Mandate authorizations retrieved successfully", response.Message)

	// Data should be an array
	require.NotNil(t, response.Data)
	require.Len(t, response.Data, 1, "should have 1 mandate authorization")

	// Validate metadata - now properly handles both per_page and perPage field names
	// The JSON uses "per_page" which is now correctly parsed into PerPage
	assert.Equal(t, int64(50), response.Meta.PerPage.Int64(), "PerPage should be parsed from per_page field")
	assert.False(t, response.Meta.Next.Valid, "Next should be null")
	require.True(t, response.Meta.Total.Valid, "Total should not be null")
	assert.Equal(t, int64(1), response.Meta.Total.Int)
}

func TestFetchMandateAuthorizationsResponse_FieldByFieldValidation(t *testing.T) {
	// Read the fetch_mandate_authorizations_200.json file
	responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "customers", "fetch_mandate_authorizations_200.json")
	responseData, err := os.ReadFile(responseFilePath)
	require.NoError(t, err, "failed to read fetch_mandate_authorizations_200.json")

	// Parse the raw JSON to get the original values
	var rawResponse map[string]any
	err = json.Unmarshal(responseData, &rawResponse)
	require.NoError(t, err, "failed to unmarshal raw JSON response")

	// Deserialize into the FetchMandateAuthorizationsResponse struct
	var response FetchMandateAuthorizationsResponse
	err = json.Unmarshal(responseData, &response)
	require.NoError(t, err, "failed to unmarshal into FetchMandateAuthorizationsResponse struct")

	// Validate top-level fields against the raw JSON
	assert.Equal(t, rawResponse["status"], response.Status.Bool(), "status field should match")
	assert.Equal(t, rawResponse["message"], response.Message, "message field should match")

	// Validate data array exists and has correct length
	rawData, hasData := rawResponse["data"]
	require.True(t, hasData, "data field should exist")
	rawDataArray, ok := rawData.([]any)
	require.True(t, ok, "data should be an array")
	require.Len(t, rawDataArray, 1, "should have 1 item")
	require.Len(t, response.Data, 1, "response data should have 1 item")

	// Validate first mandate authorization
	rawAuth := rawDataArray[0].(map[string]any)
	auth := response.Data[0]

	// Validate the actual MandateAuthorization struct fields
	assert.Equal(t, rawAuth["id"], float64(auth.ID), "id should match")
	assert.Equal(t, rawAuth["status"], string(auth.Status), "status should match")
	assert.Equal(t, rawAuth["mandate_id"], float64(auth.MandateID), "mandate_id should match")
	assert.Equal(t, rawAuth["authorization_id"], float64(auth.AuthorizationID), "authorization_id should match")
	assert.Equal(t, rawAuth["authorization_code"], auth.AuthorizationCode.String(), "authorization_code should match")
	assert.Equal(t, rawAuth["integration_id"], float64(auth.IntegrationID), "integration_id should match")
	assert.Equal(t, rawAuth["account_number"], auth.AccountNumber.String(), "account_number should match")
	assert.Equal(t, rawAuth["bank_code"], auth.BankCode.String(), "bank_code should match")
	// bank_name is null in JSON but becomes empty string in struct
	if rawAuth["bank_name"] == nil {
		assert.Empty(t, auth.BankName.String(), "bank_name should be empty when null in JSON")
	} else {
		assert.Equal(t, rawAuth["bank_name"], auth.BankName.String(), "bank_name should match")
	}

	// Additional timestamp parsing validation
	authorizedAtStr, ok := rawAuth["authorized_at"].(string)
	require.True(t, ok, "authorized_at should be a string")
	parsedTime, err := time.Parse("2006-01-02T15:04:05.999Z", authorizedAtStr)
	require.NoError(t, err, "should parse authorized_at timestamp")
	assert.Equal(t, 2024, parsedTime.Year(), "year should be 2024")
	assert.Equal(t, time.September, parsedTime.Month(), "month should be September")

	// Validate meta fields - now properly handles both per_page and perPage field names
	rawMeta := rawResponse["meta"].(map[string]any)
	// per_page field should now be properly parsed into PerPage
	expectedPerPage := int64(rawMeta["per_page"].(float64))
	assert.Equal(t, expectedPerPage, response.Meta.PerPage.Int64(), "PerPage should be parsed from per_page field")
	// next field comparisons - null vs nil pointer handling
	if rawMeta["next"] == nil {
		assert.False(t, response.Meta.Next.Valid, "Next should be invalid when null in JSON")
	} else {
		assert.Equal(t, rawMeta["next"], response.Meta.Next.String, "Next should match")
	}
	// Note: The response JSON has "count" but our Meta struct only has "total"
	assert.Equal(t, rawMeta["total"], float64(response.Meta.Total.Int), "meta.total should match")

	// Test round-trip serialization
	serialized, err := json.Marshal(response)
	require.NoError(t, err, "failed to marshal response back to JSON")

	var roundTripResponse FetchMandateAuthorizationsResponse
	err = json.Unmarshal(serialized, &roundTripResponse)
	require.NoError(t, err, "failed to unmarshal round-trip JSON")

	// Verify core fields survive round-trip
	assert.Equal(t, response.Status.Bool(), roundTripResponse.Status.Bool(), "status should survive round-trip")
	assert.Equal(t, response.Message, roundTripResponse.Message, "message should survive round-trip")
	assert.Equal(t, len(response.Data), len(roundTripResponse.Data), "data array length should survive round-trip")
}
