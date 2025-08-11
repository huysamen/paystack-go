package paymentpages

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
		expectedName    string
		expectedSlug    string
	}{
		{
			name:            "successful create response",
			responseFile:    "create_200.json",
			expectedStatus:  true,
			expectedMessage: "Page created",
			expectedName:    "Onipan and Hassan",
			expectedSlug:    "1got2y8unp",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Read the response file
			responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "paymentpages", tt.responseFile)
			responseData, err := os.ReadFile(responseFilePath)
			require.NoError(t, err, "failed to read response file: %s", responseFilePath)

			// Deserialize the JSON response
			var response CreateResponse
			err = json.Unmarshal(responseData, &response)
			require.NoError(t, err, "failed to unmarshal JSON response")

			// Verify the response structure
			assert.Equal(t, tt.expectedStatus, response.Status.Bool(), "status should match")
			assert.Equal(t, tt.expectedMessage, response.Message, "message should match")

			// Verify the data structure
			assert.NotNil(t, response.Data, "data should not be nil")
			assert.Equal(t, tt.expectedName, response.Data.Name.String(), "name should match")
			assert.Equal(t, tt.expectedSlug, response.Data.Slug.String(), "slug should match")
			assert.Greater(t, response.Data.ID.Int64(), int64(0), "ID should be greater than 0")
		})
	}
}

func TestCreateResponse_FieldByFieldValidation(t *testing.T) {
	t.Run("create_200_json_comprehensive_field_validation", func(t *testing.T) {
		// Read the exact JSON response file
		responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "paymentpages", "create_200.json")
		responseData, err := os.ReadFile(responseFilePath)
		require.NoError(t, err, "failed to read create_200.json")

		// Parse into our struct
		var response CreateResponse
		err = json.Unmarshal(responseData, &response)
		require.NoError(t, err, "failed to unmarshal create_200.json")

		// Parse the raw JSON to compare exact values
		var rawJSON map[string]any
		err = json.Unmarshal(responseData, &rawJSON)
		require.NoError(t, err, "failed to parse raw JSON for comparison")

		// Field-by-field validation against the exact JSON values
		assert.Equal(t, true, rawJSON["status"], "status in JSON should be true")
		assert.Equal(t, true, response.Status.Bool(), "status in struct should be true")

		assert.Equal(t, "Page created", rawJSON["message"], "message in JSON should match")
		assert.Equal(t, "Page created", response.Message, "message in struct should match")

		// Verify data field exists and has correct structure
		assert.Contains(t, rawJSON, "data", "JSON should contain data field")
		assert.NotNil(t, response.Data, "struct data field should not be nil")

		// Get the data portion from raw JSON
		rawData, ok := rawJSON["data"].(map[string]any)
		require.True(t, ok, "data field should be an object")

		// Validate each data field
		assert.Equal(t, "Onipan and Hassan", rawData["name"], "name in JSON should match")
		assert.Equal(t, "Onipan and Hassan", response.Data.Name.String(), "name in struct should match")

		assert.Equal(t, "", rawData["description"], "description in JSON should match")
		assert.Equal(t, "null", response.Data.Description.String(), "description in struct should match")

		assert.Equal(t, float64(10000), rawData["amount"], "amount in JSON should match")
		assert.Equal(t, int64(10000), response.Data.Amount.Int, "amount in struct should match")

		assert.Equal(t, "SPL_yqSS1FvrBz", rawData["split_code"], "split_code in JSON should match")
		assert.Equal(t, "SPL_yqSS1FvrBz", response.Data.SplitCode.String(), "split_code in struct should match")

		assert.Equal(t, float64(463433), rawData["integration"], "integration in JSON should match")
		assert.Equal(t, int64(463433), response.Data.Integration.Int64(), "integration in struct should match")

		assert.Equal(t, "test", rawData["domain"], "domain in JSON should match")
		assert.Equal(t, "test", response.Data.Domain.String(), "domain in struct should match")

		assert.Equal(t, "1got2y8unp", rawData["slug"], "slug in JSON should match")
		assert.Equal(t, "1got2y8unp", response.Data.Slug.String(), "slug in struct should match")

		assert.Equal(t, "NGN", rawData["currency"], "currency in JSON should match")
		assert.Equal(t, "NGN", response.Data.Currency.String(), "currency in struct should match")

		assert.Equal(t, "payment", rawData["type"], "type in JSON should match")
		assert.Equal(t, "payment", response.Data.Type.String(), "type in struct should match")

		assert.Equal(t, false, rawData["collect_phone"], "collect_phone in JSON should match")
		assert.Equal(t, false, response.Data.CollectPhone.Bool, "collect_phone in struct should match")

		assert.Equal(t, true, rawData["active"], "active in JSON should match")
		assert.Equal(t, true, response.Data.Active.Bool(), "active in struct should match")

		assert.Equal(t, true, rawData["published"], "published in JSON should match")
		assert.Equal(t, true, response.Data.Published.Bool, "published in struct should match")

		assert.Equal(t, false, rawData["migrate"], "migrate in JSON should match")
		assert.Equal(t, false, response.Data.Migrate.Bool, "migrate in struct should match")

		assert.Equal(t, float64(1308510), rawData["id"], "id in JSON should match")
		assert.Equal(t, int64(1308510), response.Data.ID.Int64(), "id in struct should match")

		assert.Equal(t, "2023-03-21T11:44:16.412Z", rawData["createdAt"], "createdAt in JSON should match")
		// For timestamp comparison, check that we can parse both correctly rather than string format
		expectedCreatedAt, err := time.Parse(time.RFC3339, "2023-03-21T11:44:16.412Z")
		require.NoError(t, err, "should parse expected createdAt")
		actualCreatedAt, err := time.Parse(time.RFC3339, response.Data.CreatedAt.String())
		require.NoError(t, err, "should parse actual createdAt")
		assert.True(t, expectedCreatedAt.Equal(actualCreatedAt), "createdAt timestamps should be equal")

		assert.Equal(t, "2023-03-21T11:44:16.412Z", rawData["updatedAt"], "updatedAt in JSON should match")
		// For timestamp comparison, check that we can parse both correctly rather than string format
		expectedUpdatedAt, err := time.Parse(time.RFC3339, "2023-03-21T11:44:16.412Z")
		require.NoError(t, err, "should parse expected updatedAt")
		actualUpdatedAt, err := time.Parse(time.RFC3339, response.Data.UpdatedAt.String())
		require.NoError(t, err, "should parse actual updatedAt")
		assert.True(t, expectedUpdatedAt.Equal(actualUpdatedAt), "updatedAt timestamps should be equal")

		// Verify complete JSON structure matches our struct
		reconstituted, err := json.Marshal(response)
		require.NoError(t, err, "should be able to marshal struct back to JSON")

		var reconstitutedMap map[string]any
		err = json.Unmarshal(reconstituted, &reconstitutedMap)
		require.NoError(t, err, "should be able to parse reconstituted JSON")

		// Core fields should match
		assert.Equal(t, rawJSON["status"], reconstitutedMap["status"], "status should survive round-trip")
		assert.Equal(t, rawJSON["message"], reconstitutedMap["message"], "message should survive round-trip")

		// Data field should match
		reconstitutedData, ok := reconstitutedMap["data"].(map[string]any)
		require.True(t, ok, "reconstituted data should be an object")

		assert.Equal(t, rawData["name"], reconstitutedData["name"], "name should survive round-trip")
		// Note: empty string gets converted to null by NullString, so we expect null in reconstituted data
		assert.Equal(t, nil, reconstitutedData["description"], "description should be null after round-trip (empty string becomes null)")
		assert.Equal(t, rawData["amount"], reconstitutedData["amount"], "amount should survive round-trip")
		assert.Equal(t, rawData["split_code"], reconstitutedData["split_code"], "split_code should survive round-trip")
		assert.Equal(t, rawData["integration"], reconstitutedData["integration"], "integration should survive round-trip")
		assert.Equal(t, rawData["domain"], reconstitutedData["domain"], "domain should survive round-trip")
		assert.Equal(t, rawData["slug"], reconstitutedData["slug"], "slug should survive round-trip")
		assert.Equal(t, rawData["currency"], reconstitutedData["currency"], "currency should survive round-trip")
		assert.Equal(t, rawData["type"], reconstitutedData["type"], "type should survive round-trip")
		assert.Equal(t, rawData["collect_phone"], reconstitutedData["collect_phone"], "collect_phone should survive round-trip")
		assert.Equal(t, rawData["active"], reconstitutedData["active"], "active should survive round-trip")
		assert.Equal(t, rawData["published"], reconstitutedData["published"], "published should survive round-trip")
		assert.Equal(t, rawData["migrate"], reconstitutedData["migrate"], "migrate should survive round-trip")
		assert.Equal(t, rawData["id"], reconstitutedData["id"], "id should survive round-trip")

		// For timestamps, verify they represent the same moment in time
		originalCreatedAt, err := time.Parse(time.RFC3339, rawData["createdAt"].(string))
		require.NoError(t, err, "should parse original createdAt")
		roundTripCreatedAt, err := time.Parse(time.RFC3339, reconstitutedData["createdAt"].(string))
		require.NoError(t, err, "should parse round-trip createdAt")
		assert.True(t, originalCreatedAt.Equal(roundTripCreatedAt), "createdAt should survive round-trip")

		originalUpdatedAt, err := time.Parse(time.RFC3339, rawData["updatedAt"].(string))
		require.NoError(t, err, "should parse original updatedAt")
		roundTripUpdatedAt, err := time.Parse(time.RFC3339, reconstitutedData["updatedAt"].(string))
		require.NoError(t, err, "should parse round-trip updatedAt")
		assert.True(t, originalUpdatedAt.Equal(roundTripUpdatedAt), "updatedAt should survive round-trip")
	})
}

func TestCreateRequestBuilder(t *testing.T) {
	t.Run("builds request with required name only", func(t *testing.T) {
		builder := NewCreateRequestBuilder("Test Page")
		request := builder.Build()

		assert.Equal(t, "Test Page", request.Name, "name should match")
		assert.Equal(t, "", request.Description, "description should be empty by default")
		assert.Nil(t, request.Amount, "amount should be nil by default")
		assert.Equal(t, "", request.Currency, "currency should be empty by default")
		assert.Equal(t, "", request.Slug, "slug should be empty by default")
	})

	t.Run("builds request with all fields", func(t *testing.T) {
		builder := NewCreateRequestBuilder("Test Page")
		request := builder.
			Description("Test Description").
			Amount(10000).
			Currency("NGN").
			Slug("test-page").
			Type("payment").
			Plan("plan123").
			FixedAmount(true).
			SplitCode("SPL_123").
			RedirectURL("https://redirect.url").
			SuccessMessage("Success!").
			NotificationEmail("notify@email.com").
			CollectPhone(true).
			Build()

		assert.Equal(t, "Test Page", request.Name)
		assert.Equal(t, "Test Description", request.Description)
		assert.Equal(t, 10000, *request.Amount)
		assert.Equal(t, "NGN", request.Currency)
		assert.Equal(t, "test-page", request.Slug)
		assert.Equal(t, "payment", request.Type)
		assert.Equal(t, "plan123", request.Plan)
		assert.Equal(t, true, *request.FixedAmount)
		assert.Equal(t, "SPL_123", request.SplitCode)
		assert.Equal(t, "https://redirect.url", request.RedirectURL)
		assert.Equal(t, "Success!", request.SuccessMessage)
		assert.Equal(t, "notify@email.com", request.NotificationEmail)
		assert.Equal(t, true, *request.CollectPhone)
	})

	t.Run("builds request with custom fields", func(t *testing.T) {
		builder := NewCreateRequestBuilder("Test Page")
		request := builder.
			AddCustomField("Full Name", "full_name", true).
			AddCustomField("Age", "age", false).
			Build()

		assert.Len(t, request.CustomFields, 2)
		assert.Equal(t, "Full Name", request.CustomFields[0].DisplayName)
		assert.Equal(t, "full_name", request.CustomFields[0].VariableName)
		assert.Equal(t, true, request.CustomFields[0].Required)
		assert.Equal(t, "Age", request.CustomFields[1].DisplayName)
		assert.Equal(t, "age", request.CustomFields[1].VariableName)
		assert.Equal(t, false, request.CustomFields[1].Required)
	})
}

func TestCreateRequest_JSONSerialization(t *testing.T) {
	t.Run("serializes request correctly", func(t *testing.T) {
		builder := NewCreateRequestBuilder("Test Page").
			Description("Test Description").
			Amount(10000).
			Currency("NGN")

		request := builder.Build()

		jsonData, err := json.Marshal(request)
		require.NoError(t, err, "should marshal to JSON without error")

		var unmarshaled map[string]any
		err = json.Unmarshal(jsonData, &unmarshaled)
		require.NoError(t, err, "should unmarshal JSON without error")

		assert.Equal(t, "Test Page", unmarshaled["name"])
		assert.Equal(t, "Test Description", unmarshaled["description"])
		assert.Equal(t, float64(10000), unmarshaled["amount"])
		assert.Equal(t, "NGN", unmarshaled["currency"])
	})
}
