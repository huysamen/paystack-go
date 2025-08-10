package bulkcharges

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInitiateResponse_JSONDeserialization(t *testing.T) {
	tests := []struct {
		name              string
		responseFile      string
		expectedStatus    bool
		expectedMessage   string
		expectedBatchCode string
	}{
		{
			name:              "successful initiate response",
			responseFile:      "initiate_200.json",
			expectedStatus:    true,
			expectedMessage:   "Charges have been queued",
			expectedBatchCode: "BCH_rrsbgwb4ivgzst1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Read the response file
			responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "bulkcharges", tt.responseFile)
			responseData, err := os.ReadFile(responseFilePath)
			require.NoError(t, err, "failed to read response file: %s", responseFilePath)

			// Deserialize the JSON response
			var response InitiateResponse
			err = json.Unmarshal(responseData, &response)
			require.NoError(t, err, "failed to unmarshal JSON response")

			// Verify the response structure
			assert.Equal(t, tt.expectedStatus, response.Status.Bool(), "status should match")
			assert.Equal(t, tt.expectedMessage, response.Message, "message should match")

			// Verify the data structure
			assert.NotNil(t, response.Data, "data should not be nil")
			assert.Equal(t, tt.expectedBatchCode, response.Data.BatchCode.String(), "batch code should match")
			assert.NotEmpty(t, response.Data.Reference.String(), "reference should not be empty")
			assert.Greater(t, response.Data.ID.Int64(), int64(0), "ID should be greater than 0")
		})
	}
}

func TestInitiateResponse_FieldByFieldValidation(t *testing.T) {
	t.Run("initiate_200_json_comprehensive_field_validation", func(t *testing.T) {
		// Read the exact JSON response file
		responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "bulkcharges", "initiate_200.json")
		responseData, err := os.ReadFile(responseFilePath)
		require.NoError(t, err, "failed to read initiate_200.json")

		// Parse into our struct
		var response InitiateResponse
		err = json.Unmarshal(responseData, &response)
		require.NoError(t, err, "failed to unmarshal initiate_200.json")

		// Parse the raw JSON to compare exact values
		var rawJSON map[string]any
		err = json.Unmarshal(responseData, &rawJSON)
		require.NoError(t, err, "failed to parse raw JSON for comparison")

		// Field-by-field validation against the exact JSON values
		assert.Equal(t, true, rawJSON["status"], "status in JSON should be true")
		assert.Equal(t, true, response.Status.Bool(), "status in struct should be true")

		assert.Equal(t, "Charges have been queued", rawJSON["message"], "message in JSON should match")
		assert.Equal(t, "Charges have been queued", response.Message, "message in struct should match")

		// Verify data field exists and has correct structure
		assert.Contains(t, rawJSON, "data", "JSON should contain data field")
		assert.NotNil(t, response.Data, "struct data field should not be nil")

		// Get the data portion from raw JSON
		rawData, ok := rawJSON["data"].(map[string]any)
		require.True(t, ok, "data field should be an object")

		// Validate each data field
		assert.Equal(t, "BCH_rrsbgwb4ivgzst1", rawData["batch_code"], "batch_code in JSON should match")
		assert.Equal(t, "BCH_rrsbgwb4ivgzst1", response.Data.BatchCode.String(), "batch_code in struct should match")

		assert.Equal(t, "bulkcharge-1663150565684-p18nyoa68a", rawData["reference"], "reference in JSON should match")
		assert.Equal(t, "bulkcharge-1663150565684-p18nyoa68a", response.Data.Reference.String(), "reference in struct should match")

		assert.Equal(t, float64(66608171), rawData["id"], "id in JSON should match")
		assert.Equal(t, int64(66608171), response.Data.ID.Int64(), "id in struct should match")

		assert.Equal(t, float64(463433), rawData["integration"], "integration in JSON should match")
		assert.Equal(t, int64(463433), response.Data.Integration.Int64(), "integration in struct should match")

		assert.Equal(t, "test", rawData["domain"], "domain in JSON should match")
		assert.Equal(t, "test", response.Data.Domain.String(), "domain in struct should match")

		assert.Equal(t, "active", rawData["status"], "status in JSON should match")
		assert.Equal(t, "active", response.Data.Status.String(), "status in struct should match")

		assert.Equal(t, float64(2), rawData["total_charges"], "total_charges in JSON should match")
		assert.Equal(t, int64(2), response.Data.TotalCharges.Int64(), "total_charges in struct should match")

		assert.Equal(t, float64(2), rawData["pending_charges"], "pending_charges in JSON should match")
		assert.Equal(t, int64(2), response.Data.PendingCharges.Int64(), "pending_charges in struct should match")

		assert.Equal(t, "2022-09-14T10:16:05.000Z", rawData["createdAt"], "createdAt in JSON should match")
		// For timestamp comparison, check that we can parse both correctly rather than string format
		expectedCreatedAt, err := time.Parse(time.RFC3339, "2022-09-14T10:16:05.000Z")
		require.NoError(t, err, "should parse expected createdAt")
		actualCreatedAt, err := time.Parse(time.RFC3339, response.Data.CreatedAt.String())
		require.NoError(t, err, "should parse actual createdAt")
		assert.True(t, expectedCreatedAt.Equal(actualCreatedAt), "createdAt timestamps should be equal")

		assert.Equal(t, "2022-09-14T10:16:05.000Z", rawData["updatedAt"], "updatedAt in JSON should match")
		// For timestamp comparison, check that we can parse both correctly rather than string format
		expectedUpdatedAt, err := time.Parse(time.RFC3339, "2022-09-14T10:16:05.000Z")
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

		assert.Equal(t, rawData["batch_code"], reconstitutedData["batch_code"], "batch_code should survive round-trip")
		assert.Equal(t, rawData["reference"], reconstitutedData["reference"], "reference should survive round-trip")
		assert.Equal(t, rawData["id"], reconstitutedData["id"], "id should survive round-trip")
		assert.Equal(t, rawData["integration"], reconstitutedData["integration"], "integration should survive round-trip")
		assert.Equal(t, rawData["domain"], reconstitutedData["domain"], "domain should survive round-trip")
		assert.Equal(t, rawData["status"], reconstitutedData["status"], "status should survive round-trip")
		assert.Equal(t, rawData["total_charges"], reconstitutedData["total_charges"], "total_charges should survive round-trip")
		assert.Equal(t, rawData["pending_charges"], reconstitutedData["pending_charges"], "pending_charges should survive round-trip")

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

func TestInitiateRequestBuilder(t *testing.T) {
	t.Run("builds request with single item", func(t *testing.T) {
		builder := NewInitiateRequestBuilder()
		request := builder.
			AddItem("AUTH_test123", 50000, "ref_123").
			Build()

		assert.Len(t, *request, 1, "should have one item")

		item := (*request)[0]
		assert.Equal(t, "AUTH_test123", item.Authorization, "authorization should match")
		assert.Equal(t, int64(50000), item.Amount, "amount should match")
		assert.Equal(t, "ref_123", item.Reference, "reference should match")
	})

	t.Run("builds request with multiple items", func(t *testing.T) {
		builder := NewInitiateRequestBuilder()
		request := builder.
			AddItem("AUTH_test123", 50000, "ref_123").
			AddItem("AUTH_test456", 75000, "ref_456").
			Build()

		assert.Len(t, *request, 2, "should have two items")

		assert.Equal(t, "AUTH_test123", (*request)[0].Authorization)
		assert.Equal(t, int64(50000), (*request)[0].Amount)
		assert.Equal(t, "ref_123", (*request)[0].Reference)

		assert.Equal(t, "AUTH_test456", (*request)[1].Authorization)
		assert.Equal(t, int64(75000), (*request)[1].Amount)
		assert.Equal(t, "ref_456", (*request)[1].Reference)
	})

	t.Run("builds request with items slice", func(t *testing.T) {
		items := []BulkChargeItem{
			{Authorization: "AUTH_test789", Amount: 25000, Reference: "ref_789"},
			{Authorization: "AUTH_test012", Amount: 35000, Reference: "ref_012"},
		}

		builder := NewInitiateRequestBuilder()
		request := builder.
			AddItems(items).
			Build()

		assert.Len(t, *request, 2, "should have two items")
		assert.Equal(t, items[0], (*request)[0], "first item should match")
		assert.Equal(t, items[1], (*request)[1], "second item should match")
	})

	t.Run("builds request combining single items and slice", func(t *testing.T) {
		items := []BulkChargeItem{
			{Authorization: "AUTH_slice1", Amount: 10000, Reference: "slice_ref1"},
		}

		builder := NewInitiateRequestBuilder()
		request := builder.
			AddItem("AUTH_single", 20000, "single_ref").
			AddItems(items).
			AddItem("AUTH_another", 30000, "another_ref").
			Build()

		assert.Len(t, *request, 3, "should have three items")
		assert.Equal(t, "AUTH_single", (*request)[0].Authorization)
		assert.Equal(t, "AUTH_slice1", (*request)[1].Authorization)
		assert.Equal(t, "AUTH_another", (*request)[2].Authorization)
	})

	t.Run("builds empty request", func(t *testing.T) {
		builder := NewInitiateRequestBuilder()
		request := builder.Build()

		assert.Empty(t, *request, "should be empty")
	})
}

func TestInitiateRequest_JSONSerialization(t *testing.T) {
	t.Run("serializes request correctly", func(t *testing.T) {
		builder := NewInitiateRequestBuilder()
		request := builder.
			AddItem("AUTH_test123", 50000, "ref_123").
			AddItem("AUTH_test456", 75000, "ref_456").
			Build()

		jsonData, err := json.Marshal(request)
		require.NoError(t, err, "should marshal to JSON without error")

		var unmarshaled []BulkChargeItem
		err = json.Unmarshal(jsonData, &unmarshaled)
		require.NoError(t, err, "should unmarshal JSON without error")

		assert.Len(t, unmarshaled, 2, "should have two items")
		assert.Equal(t, "AUTH_test123", unmarshaled[0].Authorization)
		assert.Equal(t, int64(50000), unmarshaled[0].Amount)
		assert.Equal(t, "ref_123", unmarshaled[0].Reference)
	})
}

func TestBulkChargeItem_Structure(t *testing.T) {
	t.Run("item has correct field types", func(t *testing.T) {
		item := BulkChargeItem{
			Authorization: "AUTH_test",
			Amount:        12345,
			Reference:     "ref_test",
		}

		assert.IsType(t, "", item.Authorization)
		assert.IsType(t, int64(0), item.Amount)
		assert.IsType(t, "", item.Reference)
	})
}
