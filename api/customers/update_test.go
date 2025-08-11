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

func TestUpdateResponse_JSONDeserialization(t *testing.T) {
	// Read the update_200.json file
	responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "customers", "update_200.json")
	responseData, err := os.ReadFile(responseFilePath)
	require.NoError(t, err, "failed to read update_200.json")

	var response UpdateResponse
	err = json.Unmarshal(responseData, &response)
	require.NoError(t, err, "failed to unmarshal update response")

	// Basic validations
	assert.True(t, response.Status.Bool())
	assert.Equal(t, "Customer updated", response.Message)
	assert.NotNil(t, response.Data)

	// Validate updated customer data
	assert.Equal(t, "bojack@horsinaround.com", response.Data.Email.String())
	assert.Equal(t, "CUS_xnxdt6s1zg1f4nx", response.Data.CustomerCode.String())
	assert.Equal(t, uint64(1173), response.Data.ID.Uint64())
	assert.Equal(t, "BoJack", response.Data.FirstName.String())
	assert.Equal(t, "Horseman", response.Data.LastName.String())
	assert.False(t, response.Data.Phone.Valid)

	// Validate metadata structure
	assert.True(t, response.Data.Metadata.Valid)
	metadata := map[string]any(response.Data.Metadata.Metadata)
	assert.Contains(t, metadata, "photos")
	photos := metadata["photos"].([]any)
	assert.Len(t, photos, 1)
}

func TestUpdateRequestBuilder(t *testing.T) {
	t.Run("builds empty request", func(t *testing.T) {
		builder := NewUpdateRequestBuilder()
		req := builder.Build()

		assert.Nil(t, req.FirstName)
		assert.Nil(t, req.LastName)
		assert.Nil(t, req.Phone)
		assert.Nil(t, req.Metadata)
	})

	t.Run("builds request with all fields", func(t *testing.T) {
		metadata := map[string]any{"updated": true}
		builder := NewUpdateRequestBuilder().
			FirstName("John").
			LastName("Doe").
			Phone("+1234567890").
			Metadata(metadata)

		req := builder.Build()

		assert.Equal(t, "John", *req.FirstName)
		assert.Equal(t, "Doe", *req.LastName)
		assert.Equal(t, "+1234567890", *req.Phone)
		assert.Equal(t, metadata, req.Metadata)
	})

	t.Run("builds request with partial updates", func(t *testing.T) {
		builder := NewUpdateRequestBuilder().
			FirstName("Jane").
			Phone("+1987654321")

		req := builder.Build()

		assert.Equal(t, "Jane", *req.FirstName)
		assert.Nil(t, req.LastName)
		assert.Equal(t, "+1987654321", *req.Phone)
		assert.Nil(t, req.Metadata)
	})
}

func TestUpdateRequest_JSONSerialization(t *testing.T) {
	t.Run("serializes empty request correctly", func(t *testing.T) {
		builder := NewUpdateRequestBuilder()
		req := builder.Build()

		jsonData, err := json.Marshal(req)
		require.NoError(t, err, "should marshal without error")

		var unmarshaled map[string]any
		err = json.Unmarshal(jsonData, &unmarshaled)
		require.NoError(t, err, "should unmarshal JSON without error")

		// All fields should be omitted when nil
		_, hasFirstName := unmarshaled["first_name"]
		_, hasLastName := unmarshaled["last_name"]
		_, hasPhone := unmarshaled["phone"]
		_, hasMetadata := unmarshaled["metadata"]
		assert.False(t, hasFirstName, "first_name should be omitted when nil")
		assert.False(t, hasLastName, "last_name should be omitted when nil")
		assert.False(t, hasPhone, "phone should be omitted when nil")
		assert.False(t, hasMetadata, "metadata should be omitted when nil")
	})

	t.Run("includes provided fields only", func(t *testing.T) {
		builder := NewUpdateRequestBuilder().
			FirstName("Alice").
			Metadata(map[string]any{"role": "admin"})

		req := builder.Build()

		jsonData, err := json.Marshal(req)
		require.NoError(t, err, "should marshal without error")

		var unmarshaled map[string]any
		err = json.Unmarshal(jsonData, &unmarshaled)
		require.NoError(t, err, "should unmarshal JSON without error")

		assert.Equal(t, "Alice", unmarshaled["first_name"])
		assert.Equal(t, map[string]any{"role": "admin"}, unmarshaled["metadata"])

		_, hasLastName := unmarshaled["last_name"]
		_, hasPhone := unmarshaled["phone"]
		assert.False(t, hasLastName, "last_name should be omitted when nil")
		assert.False(t, hasPhone, "phone should be omitted when nil")
	})
}

func TestUpdateResponse_FieldByFieldValidation(t *testing.T) {
	// Read the update_200.json file
	responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "customers", "update_200.json")
	responseData, err := os.ReadFile(responseFilePath)
	require.NoError(t, err, "failed to read update_200.json")

	// Parse the raw JSON to get the original values
	var rawResponse map[string]any
	err = json.Unmarshal(responseData, &rawResponse)
	require.NoError(t, err, "failed to unmarshal raw JSON response")

	// Deserialize into the UpdateResponse struct
	var response UpdateResponse
	err = json.Unmarshal(responseData, &response)
	require.NoError(t, err, "failed to unmarshal into UpdateResponse struct")

	// Validate top-level fields against the raw JSON
	assert.Equal(t, rawResponse["status"], response.Status.Bool(), "status field should match")
	assert.Equal(t, rawResponse["message"], response.Message, "message field should match")

	// Validate data object fields
	rawData := rawResponse["data"].(map[string]any)

	// Basic customer fields
	assert.Equal(t, rawData["email"], response.Data.Email.String(), "email should match")
	assert.Equal(t, int(rawData["integration"].(float64)), int(response.Data.Integration.Int), "integration should match")
	assert.Equal(t, rawData["domain"], response.Data.Domain.String(), "domain should match")
	assert.Equal(t, rawData["customer_code"], response.Data.CustomerCode.String(), "customer_code should match")
	assert.Equal(t, uint64(rawData["id"].(float64)), response.Data.ID.Uint64(), "id should match")
	assert.Equal(t, rawData["identified"], response.Data.Identified.Bool(), "identified should match")

	// Handle nullable fields
	if rawData["first_name"] == nil {
		assert.False(t, response.Data.FirstName.Valid, "first_name should be invalid")
	} else {
		assert.Equal(t, rawData["first_name"], response.Data.FirstName.String(), "first_name should match")
	}

	if rawData["last_name"] == nil {
		assert.False(t, response.Data.LastName.Valid, "last_name should be invalid")
	} else {
		assert.Equal(t, rawData["last_name"], response.Data.LastName.String(), "last_name should match")
	}

	if rawData["phone"] == nil {
		assert.False(t, response.Data.Phone.Valid, "phone should be invalid")
	} else {
		assert.Equal(t, rawData["phone"], response.Data.Phone.String(), "phone should match")
	}

	// Handle metadata - should match complex structure
	if rawData["metadata"] == nil {
		assert.False(t, response.Data.Metadata.Valid, "metadata should be invalid for null")
	} else {
		assert.True(t, response.Data.Metadata.Valid, "metadata should be valid")
		assert.Equal(t, rawData["metadata"], map[string]any(response.Data.Metadata.Metadata), "metadata should match")
	}

	if rawData["identifications"] == nil {
		assert.False(t, response.Data.Identifications.Valid, "identifications should be invalid")
	} else {
		assert.True(t, response.Data.Identifications.Valid, "identifications should be valid when present in JSON")
	}

	// For timestamp comparisons, parse both and compare the actual time values
	expectedCreatedAt, err := time.Parse(time.RFC3339, rawData["createdAt"].(string))
	require.NoError(t, err, "should parse expected createdAt")
	assert.True(t, expectedCreatedAt.Equal(response.Data.CreatedAt.Time()), "createdAt should represent the same moment")

	expectedUpdatedAt, err := time.Parse(time.RFC3339, rawData["updatedAt"].(string))
	require.NoError(t, err, "should parse expected updatedAt")
	assert.True(t, expectedUpdatedAt.Equal(response.Data.UpdatedAt.Time()), "updatedAt should represent the same moment")

	// Note: The update response JSON includes transactions, subscriptions, and authorizations arrays,
	// but the UpdateResponseData (types.Customer) doesn't include these fields.
	// This is expected as the update endpoint returns a basic customer structure,
	// while the fetch endpoint returns CustomerWithRelations.

	// Test round-trip serialization
	serialized, err := json.Marshal(response)
	require.NoError(t, err, "failed to marshal response back to JSON")

	var roundTripResponse UpdateResponse
	err = json.Unmarshal(serialized, &roundTripResponse)
	require.NoError(t, err, "failed to unmarshal round-trip JSON")

	// Verify core fields survive round-trip
	assert.Equal(t, response.Status.Bool(), roundTripResponse.Status.Bool(), "status should survive round-trip")
	assert.Equal(t, response.Message, roundTripResponse.Message, "message should survive round-trip")
	assert.Equal(t, response.Data.Email.String(), roundTripResponse.Data.Email.String(), "email should survive round-trip")
	assert.Equal(t, response.Data.CustomerCode.String(), roundTripResponse.Data.CustomerCode.String(), "customer_code should survive round-trip")
	assert.Equal(t, response.Data.ID.Uint64(), roundTripResponse.Data.ID.Uint64(), "id should survive round-trip")

	// Verify metadata complex structure survives round-trip
	assert.Equal(t, response.Data.Metadata.Valid, roundTripResponse.Data.Metadata.Valid, "metadata validity should survive round-trip")
	if response.Data.Metadata.Valid && roundTripResponse.Data.Metadata.Valid {
		originalMetadata := map[string]any(response.Data.Metadata.Metadata)
		roundTripMetadata := map[string]any(roundTripResponse.Data.Metadata.Metadata)
		assert.Equal(t, originalMetadata, roundTripMetadata, "complex metadata should survive round-trip")
	}
}
