package bulkcharges

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

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
			assert.Equal(t, tt.expectedStatus, response.Status, "status should match")
			assert.Equal(t, tt.expectedMessage, response.Message, "message should match")

			// Verify the data structure
			assert.NotNil(t, response.Data, "data should not be nil")
			assert.Equal(t, tt.expectedBatchCode, response.Data.BatchCode, "batch code should match")
			assert.NotEmpty(t, response.Data.Reference, "reference should not be empty")
			assert.Greater(t, response.Data.ID, 0, "ID should be greater than 0")
		})
	}
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
