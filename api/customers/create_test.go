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

func TestCreateResponse_JSONDeserialization(t *testing.T) {
	// Read the create_200.json file
	responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "customers", "create_200.json")
	responseData, err := os.ReadFile(responseFilePath)
	require.NoError(t, err, "failed to read create_200.json")

	var response CreateResponse
	err = json.Unmarshal(responseData, &response)
	require.NoError(t, err, "failed to unmarshal create response")

	// Basic validations
	assert.True(t, response.Status.Bool())
	assert.Equal(t, "Customer created", response.Message)
	assert.NotNil(t, response.Data)

	// Validate customer data
	assert.Equal(t, "customer@email.com", response.Data.Email)
	assert.Equal(t, "CUS_xnxdt6s1zg1f4nx", response.Data.CustomerCode)
	assert.Equal(t, uint64(1173), response.Data.ID)
	assert.False(t, response.Data.Identified)
	assert.Equal(t, "test", response.Data.Domain)
}

func TestCreateRequestBuilder(t *testing.T) {
	t.Run("builds basic request with email only", func(t *testing.T) {
		builder := NewCreateRequestBuilder("test@example.com")
		req := builder.Build()

		assert.Equal(t, "test@example.com", req.Email)
		assert.Nil(t, req.FirstName)
		assert.Nil(t, req.LastName)
		assert.Nil(t, req.Phone)
		assert.Nil(t, req.Metadata)
	})

	t.Run("builds request with all optional fields", func(t *testing.T) {
		metadata := map[string]any{"key": "value"}
		builder := NewCreateRequestBuilder("test@example.com").
			FirstName("John").
			LastName("Doe").
			Phone("+1234567890").
			Metadata(metadata)

		req := builder.Build()

		assert.Equal(t, "test@example.com", req.Email)
		assert.Equal(t, "John", *req.FirstName)
		assert.Equal(t, "Doe", *req.LastName)
		assert.Equal(t, "+1234567890", *req.Phone)
		assert.Equal(t, metadata, req.Metadata)
	})
}

func TestCreateRequest_JSONSerialization(t *testing.T) {
	t.Run("serializes basic request correctly", func(t *testing.T) {
		builder := NewCreateRequestBuilder("test@example.com")
		req := builder.Build()

		jsonData, err := json.Marshal(req)
		require.NoError(t, err, "should marshal without error")

		var unmarshaled map[string]any
		err = json.Unmarshal(jsonData, &unmarshaled)
		require.NoError(t, err, "should unmarshal JSON without error")

		assert.Equal(t, "test@example.com", unmarshaled["email"], "email should match")
		// Optional fields should not be present when nil
		_, hasFirstName := unmarshaled["first_name"]
		_, hasLastName := unmarshaled["last_name"]
		_, hasPhone := unmarshaled["phone"]
		_, hasMetadata := unmarshaled["metadata"]
		assert.False(t, hasFirstName, "first_name should be omitted when nil")
		assert.False(t, hasLastName, "last_name should be omitted when nil")
		assert.False(t, hasPhone, "phone should be omitted when nil")
		assert.False(t, hasMetadata, "metadata should be omitted when nil")
	})

	t.Run("includes all fields when provided", func(t *testing.T) {
		metadata := map[string]any{"custom_field": "custom_value"}
		builder := NewCreateRequestBuilder("test@example.com").
			FirstName("Jane").
			LastName("Smith").
			Phone("+1987654321").
			Metadata(metadata)

		req := builder.Build()

		jsonData, err := json.Marshal(req)
		require.NoError(t, err, "should marshal without error")

		var unmarshaled map[string]any
		err = json.Unmarshal(jsonData, &unmarshaled)
		require.NoError(t, err, "should unmarshal JSON without error")

		assert.Equal(t, "test@example.com", unmarshaled["email"])
		assert.Equal(t, "Jane", unmarshaled["first_name"])
		assert.Equal(t, "Smith", unmarshaled["last_name"])
		assert.Equal(t, "+1987654321", unmarshaled["phone"])
		assert.Equal(t, metadata, unmarshaled["metadata"])
	})
}

func TestCreateResponse_FieldByFieldValidation(t *testing.T) {
	// Read the create_200.json file
	responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "customers", "create_200.json")
	responseData, err := os.ReadFile(responseFilePath)
	require.NoError(t, err, "failed to read create_200.json")

	// Parse the raw JSON to get the original values
	var rawResponse map[string]any
	err = json.Unmarshal(responseData, &rawResponse)
	require.NoError(t, err, "failed to unmarshal raw JSON response")

	// Deserialize into the CreateResponse struct
	var response CreateResponse
	err = json.Unmarshal(responseData, &response)
	require.NoError(t, err, "failed to unmarshal into CreateResponse struct")

	// Validate top-level fields against the raw JSON
	assert.Equal(t, rawResponse["status"], response.Status.Bool(), "status field should match")
	assert.Equal(t, rawResponse["message"], response.Message, "message field should match")

	// Validate data object fields
	rawData := rawResponse["data"].(map[string]any)
	assert.Equal(t, rawData["email"], response.Data.Email, "email should match")
	assert.Equal(t, int(rawData["integration"].(float64)), *response.Data.Integration, "integration should match")
	assert.Equal(t, rawData["domain"], response.Data.Domain, "domain should match")
	assert.Equal(t, rawData["customer_code"], response.Data.CustomerCode, "customer_code should match")
	assert.Equal(t, uint64(rawData["id"].(float64)), response.Data.ID, "id should match")
	assert.Equal(t, rawResponse["data"].(map[string]any)["identified"], response.Data.Identified, "identified should match")

	// Handle null identifications field
	if rawData["identifications"] == nil {
		assert.Nil(t, response.Data.Identifications, "identifications should be nil")
	} else {
		assert.NotNil(t, response.Data.Identifications, "identifications should not be nil when present in JSON")
	}

	// For timestamp comparisons, parse both and compare the actual time values
	expectedCreatedAt, err := time.Parse(time.RFC3339, rawData["createdAt"].(string))
	require.NoError(t, err, "should parse expected createdAt")
	assert.True(t, expectedCreatedAt.Equal(response.Data.CreatedAt.Time), "createdAt should represent the same moment")

	expectedUpdatedAt, err := time.Parse(time.RFC3339, rawData["updatedAt"].(string))
	require.NoError(t, err, "should parse expected updatedAt")
	assert.True(t, expectedUpdatedAt.Equal(response.Data.UpdatedAt.Time), "updatedAt should represent the same moment")

	// Test round-trip serialization
	serialized, err := json.Marshal(response)
	require.NoError(t, err, "failed to marshal response back to JSON")

	var roundTripResponse CreateResponse
	err = json.Unmarshal(serialized, &roundTripResponse)
	require.NoError(t, err, "failed to unmarshal round-trip JSON")

	// Verify core fields survive round-trip
	assert.Equal(t, response.Status.Bool(), roundTripResponse.Status.Bool(), "status should survive round-trip")
	assert.Equal(t, response.Message, roundTripResponse.Message, "message should survive round-trip")
	assert.Equal(t, response.Data.Email, roundTripResponse.Data.Email, "email should survive round-trip")
	assert.Equal(t, response.Data.CustomerCode, roundTripResponse.Data.CustomerCode, "customer_code should survive round-trip")
	assert.Equal(t, response.Data.ID, roundTripResponse.Data.ID, "id should survive round-trip")
}
