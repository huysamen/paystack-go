package plans

import (
	"encoding/json"
	"net/url"
	"os"
	"path/filepath"
	"testing"

	"github.com/huysamen/paystack-go/enums"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListResponse_JSONDeserialization(t *testing.T) {
	tests := []struct {
		name            string
		responseFile    string
		expectedStatus  bool
		expectedMessage string
	}{
		{
			name:            "successful list response",
			responseFile:    "list_200.json",
			expectedStatus:  true,
			expectedMessage: "Plans retrieved",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Read the JSON response file
			responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "plans", tt.responseFile)
			responseData, err := os.ReadFile(responseFilePath)
			require.NoError(t, err, "failed to read response file")

			// Deserialize the JSON response
			var response ListResponse
			err = json.Unmarshal(responseData, &response)
			require.NoError(t, err, "failed to unmarshal JSON response")

			// Verify the response structure
			assert.Equal(t, tt.expectedStatus, response.Status.Bool(), "status should match")
			assert.Equal(t, tt.expectedMessage, response.Message, "message should match")
			assert.NotNil(t, response.Data, "data should not be nil")
			assert.Len(t, response.Data, 2, "should have 2 plans")
		})
	}
}

func TestListResponse_FieldByFieldValidation(t *testing.T) {
	t.Run("list_200_json_comprehensive_field_validation", func(t *testing.T) {
		// Read the exact JSON response file
		responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "plans", "list_200.json")
		responseData, err := os.ReadFile(responseFilePath)
		require.NoError(t, err, "failed to read list_200.json")

		// Parse raw JSON to compare field by field
		var rawJSON map[string]any
		err = json.Unmarshal(responseData, &rawJSON)
		require.NoError(t, err, "failed to parse raw JSON")

		// Parse into our struct
		var response ListResponse
		err = json.Unmarshal(responseData, &response)
		require.NoError(t, err, "failed to unmarshal list_200.json")

		// Verify struct data is properly populated
		assert.NotNil(t, response.Data, "struct data field should not be nil")
		assert.Len(t, response.Data, 2, "should have 2 plans")

		// Get the data portion from raw JSON
		rawDataArray, ok := rawJSON["data"].([]any)
		require.True(t, ok, "data field should be an array")
		require.Len(t, rawDataArray, 2, "should have 2 plans in raw data")

		// Test first plan (Satin Flower - weekly)
		rawPlan := rawDataArray[0].(map[string]any)
		plan := response.Data[0]

		assert.Equal(t, "Satin Flower", rawPlan["name"], "first plan name in JSON should match")
		assert.Equal(t, "Satin Flower", plan.Name.String(), "first plan name in struct should match")

		assert.Equal(t, float64(100000), rawPlan["amount"], "first plan amount in JSON should match")
		assert.Equal(t, int64(100000), plan.Amount.Int64(), "first plan amount in struct should match")

		assert.Equal(t, "weekly", rawPlan["interval"], "first plan interval in JSON should match")
		assert.Equal(t, enums.IntervalWeekly, plan.Interval, "first plan interval in struct should match")

		assert.Equal(t, "PLN_lkozbpsoyd4je9t", rawPlan["plan_code"], "first plan plan_code in JSON should match")
		assert.Equal(t, "PLN_lkozbpsoyd4je9t", plan.PlanCode.String(), "first plan plan_code in struct should match")

		assert.Equal(t, float64(27), rawPlan["id"], "first plan id in JSON should match")
		assert.Equal(t, uint64(27), plan.ID.Uint64(), "first plan id in struct should match")

		// Test first plan subscriptions
		rawSubscriptions, ok := rawPlan["subscriptions"].([]any)
		require.True(t, ok, "first plan subscriptions should be an array")
		assert.Len(t, rawSubscriptions, 1, "first plan should have 1 subscription")
		assert.Len(t, plan.Subscriptions, 1, "first plan struct should have 1 subscription")

		// Test second plan (Monthly retainer - monthly)
		rawPlan2 := rawDataArray[1].(map[string]any)
		plan2 := response.Data[1]

		assert.Equal(t, "Monthly retainer", rawPlan2["name"], "second plan name in JSON should match")
		assert.Equal(t, "Monthly retainer", plan2.Name.String(), "second plan name in struct should match")

		assert.Equal(t, float64(50000), rawPlan2["amount"], "second plan amount in JSON should match")
		assert.Equal(t, int64(50000), plan2.Amount.Int64(), "second plan amount in struct should match")

		assert.Equal(t, "monthly", rawPlan2["interval"], "second plan interval in JSON should match")
		assert.Equal(t, enums.IntervalMonthly, plan2.Interval, "second plan interval in struct should match")

		assert.Equal(t, float64(28), rawPlan2["id"], "second plan id in JSON should match")
		assert.Equal(t, uint64(28), plan2.ID.Uint64(), "second plan id in struct should match")

		// Test second plan subscriptions (empty)
		rawSubscriptions2, ok := rawPlan2["subscriptions"].([]any)
		require.True(t, ok, "second plan subscriptions should be an array")
		assert.Empty(t, rawSubscriptions2, "second plan should have empty subscriptions")
		assert.Empty(t, plan2.Subscriptions, "second plan struct should have empty subscriptions")

		// Test meta information
		rawMeta, ok := rawJSON["meta"].(map[string]any)
		require.True(t, ok, "meta field should be an object")

		assert.Equal(t, float64(2), rawMeta["total"], "meta total should match")
		assert.Equal(t, float64(0), rawMeta["skipped"], "meta skipped should match")
		assert.Equal(t, float64(50), rawMeta["perPage"], "meta perPage should match")
		assert.Equal(t, float64(1), rawMeta["page"], "meta page should match")
		assert.Equal(t, float64(1), rawMeta["pageCount"], "meta pageCount should match")

		// Verify complete JSON structure matches our struct
		reconstituted, err := json.Marshal(response)
		require.NoError(t, err, "should be able to marshal struct back to JSON")

		var reconstitutedMap map[string]any
		err = json.Unmarshal(reconstituted, &reconstitutedMap)
		require.NoError(t, err, "should be able to parse reconstituted JSON")

		// Core fields should match
		assert.Equal(t, rawJSON["status"], reconstitutedMap["status"], "status should survive round-trip")
		assert.Equal(t, rawJSON["message"], reconstitutedMap["message"], "message should survive round-trip")
	})
}

func TestListRequestBuilder(t *testing.T) {
	tests := []struct {
		name             string
		setupBuilder     func() *ListRequestBuilder
		expectedPerPage  *int
		expectedPage     *int
		expectedStatus   *string
		expectedInterval *enums.Interval
		expectedAmount   *int
	}{
		{
			name: "builds request with no filters",
			setupBuilder: func() *ListRequestBuilder {
				return NewListRequestBuilder()
			},
		},
		{
			name: "builds request with pagination",
			setupBuilder: func() *ListRequestBuilder {
				return NewListRequestBuilder().
					PerPage(25).
					Page(2)
			},
			expectedPerPage: intPtr(25),
			expectedPage:    intPtr(2),
		},
		{
			name: "builds request with filters",
			setupBuilder: func() *ListRequestBuilder {
				return NewListRequestBuilder().
					Status("active").
					Interval(enums.IntervalMonthly).
					Amount(50000)
			},
			expectedStatus:   stringPtr("active"),
			expectedInterval: intervalPtr(enums.IntervalMonthly),
			expectedAmount:   intPtr(50000),
		},
		{
			name: "builds request with complete filters",
			setupBuilder: func() *ListRequestBuilder {
				return NewListRequestBuilder().
					PerPage(10).
					Page(1).
					Status("complete").
					Interval(enums.IntervalWeekly).
					Amount(100000)
			},
			expectedPerPage:  intPtr(10),
			expectedPage:     intPtr(1),
			expectedStatus:   stringPtr("complete"),
			expectedInterval: intervalPtr(enums.IntervalWeekly),
			expectedAmount:   intPtr(100000),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			builder := tt.setupBuilder()
			req := builder.Build()

			if tt.expectedPerPage != nil {
				require.NotNil(t, req.PerPage)
				assert.Equal(t, *tt.expectedPerPage, *req.PerPage)
			} else {
				assert.Nil(t, req.PerPage)
			}

			if tt.expectedPage != nil {
				require.NotNil(t, req.Page)
				assert.Equal(t, *tt.expectedPage, *req.Page)
			} else {
				assert.Nil(t, req.Page)
			}

			if tt.expectedStatus != nil {
				require.NotNil(t, req.Status)
				assert.Equal(t, *tt.expectedStatus, *req.Status)
			} else {
				assert.Nil(t, req.Status)
			}

			if tt.expectedInterval != nil {
				require.NotNil(t, req.Interval)
				assert.Equal(t, *tt.expectedInterval, *req.Interval)
			} else {
				assert.Nil(t, req.Interval)
			}

			if tt.expectedAmount != nil {
				require.NotNil(t, req.Amount)
				assert.Equal(t, *tt.expectedAmount, *req.Amount)
			} else {
				assert.Nil(t, req.Amount)
			}
		})
	}
}

func TestListRequest_URLQuery(t *testing.T) {
	tests := []struct {
		name             string
		builder          func() *ListRequestBuilder
		expectedContains []string
		expectedMissing  []string
	}{
		{
			name: "converts to URL query correctly",
			builder: func() *ListRequestBuilder {
				return NewListRequestBuilder().
					PerPage(20).
					Page(3).
					Status("active").
					Interval(enums.IntervalMonthly).
					Amount(25000)
			},
			expectedContains: []string{"perPage=20", "page=3", "status=active", "interval=monthly", "amount=25000"},
		},
		{
			name: "omits empty values",
			builder: func() *ListRequestBuilder {
				return NewListRequestBuilder().
					PerPage(50)
			},
			expectedContains: []string{"perPage=50"},
			expectedMissing:  []string{"page=", "status=", "interval=", "amount="},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			builder := tt.builder()
			req := builder.Build()
			query := req.toQuery()

			// Parse the query to verify it's valid
			parsedValues, err := url.ParseQuery(query)
			require.NoError(t, err, "query should be valid URL query string")

			// Check expected values are present
			for _, expected := range tt.expectedContains {
				assert.Contains(t, query, expected, "query should contain %s", expected)
			}

			// Check missing values are not present
			for _, missing := range tt.expectedMissing {
				assert.NotContains(t, query, missing, "query should not contain %s", missing)
			}

			// Verify we can reconstruct the values
			if len(tt.expectedContains) > 0 {
				assert.NotEmpty(t, parsedValues, "parsed values should not be empty")
			}
		})
	}
}

// Helper functions
func intervalPtr(i enums.Interval) *enums.Interval {
	return &i
}
