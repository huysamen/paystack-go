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

func TestListResponse_JSONDeserialization(t *testing.T) {
	// Read the list_200.json file
	responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "customers", "list_200.json")
	responseData, err := os.ReadFile(responseFilePath)
	require.NoError(t, err, "failed to read list_200.json")

	var response ListResponse
	err = json.Unmarshal(responseData, &response)
	require.NoError(t, err, "failed to unmarshal list response")

	// Basic validations
	assert.True(t, response.Status.Bool())
	assert.Equal(t, "Customers retrieved", response.Message)
	assert.NotNil(t, response.Data)
	assert.Len(t, response.Data, 3, "should have 3 customers")

	// Validate first customer
	firstCustomer := response.Data[0]
	assert.Equal(t, "dom@gmail.com", firstCustomer.Email)
	assert.Equal(t, "CUS_c6wqvwmvwopw4ms", firstCustomer.CustomerCode)
	assert.Equal(t, uint64(90758908), firstCustomer.ID)

	// Validate second customer (has names and phone)
	secondCustomer := response.Data[1]
	assert.Equal(t, "okiki@sample.com", secondCustomer.Email)
	assert.Equal(t, "CUS_rki2ccocw7g8lsj", secondCustomer.CustomerCode)
	assert.Equal(t, uint64(90758301), secondCustomer.ID)
	assert.NotNil(t, secondCustomer.FirstName)
	assert.Equal(t, "Okiki", *secondCustomer.FirstName)
	assert.NotNil(t, secondCustomer.LastName)
	assert.Equal(t, "Sample", *secondCustomer.LastName)
	assert.NotNil(t, secondCustomer.Phone)
	assert.Equal(t, "09048829123", *secondCustomer.Phone)
}

func TestListRequestBuilder(t *testing.T) {
	t.Run("builds basic request with no parameters", func(t *testing.T) {
		builder := NewListRequestBuilder()
		req := builder.Build()

		assert.Nil(t, req.PerPage)
		assert.Nil(t, req.Page)
		assert.Nil(t, req.From)
		assert.Nil(t, req.To)
	})

	t.Run("builds request with pagination", func(t *testing.T) {
		builder := NewListRequestBuilder().
			PerPage(50).
			Page(2)

		req := builder.Build()

		assert.Equal(t, 50, *req.PerPage)
		assert.Equal(t, 2, *req.Page)
		assert.Nil(t, req.From)
		assert.Nil(t, req.To)
	})

	t.Run("builds request with date range", func(t *testing.T) {
		from := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
		to := time.Date(2023, 12, 31, 23, 59, 59, 0, time.UTC)

		builder := NewListRequestBuilder().
			DateRange(from, to)

		req := builder.Build()

		assert.Nil(t, req.PerPage)
		assert.Nil(t, req.Page)
		assert.True(t, from.Equal(*req.From))
		assert.True(t, to.Equal(*req.To))
	})

	t.Run("builds request with individual dates", func(t *testing.T) {
		from := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
		to := time.Date(2023, 12, 31, 23, 59, 59, 0, time.UTC)

		builder := NewListRequestBuilder().
			From(from).
			To(to)

		req := builder.Build()

		assert.True(t, from.Equal(*req.From))
		assert.True(t, to.Equal(*req.To))
	})
}

func TestListRequest_QueryGeneration(t *testing.T) {
	t.Run("generates query with all parameters", func(t *testing.T) {
		from := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
		to := time.Date(2023, 12, 31, 23, 59, 59, 0, time.UTC)

		builder := NewListRequestBuilder().
			PerPage(25).
			Page(1).
			From(from).
			To(to)

		req := builder.Build()
		query := req.toQuery()

		assert.Contains(t, query, "perPage=25")
		assert.Contains(t, query, "page=1")
		assert.Contains(t, query, "from=2023-01-01T00%3A00%3A00Z")
		assert.Contains(t, query, "to=2023-12-31T23%3A59%3A59Z")
	})

	t.Run("generates empty query for empty request", func(t *testing.T) {
		builder := NewListRequestBuilder()
		req := builder.Build()
		query := req.toQuery()

		assert.Empty(t, query)
	})
}

func TestListResponse_FieldByFieldValidation(t *testing.T) {
	// Read the list_200.json file
	responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "customers", "list_200.json")
	responseData, err := os.ReadFile(responseFilePath)
	require.NoError(t, err, "failed to read list_200.json")

	// Parse the raw JSON to get the original values
	var rawResponse map[string]any
	err = json.Unmarshal(responseData, &rawResponse)
	require.NoError(t, err, "failed to unmarshal raw JSON response")

	// Deserialize into the ListResponse struct
	var response ListResponse
	err = json.Unmarshal(responseData, &response)
	require.NoError(t, err, "failed to unmarshal into ListResponse struct")

	// Validate top-level fields against the raw JSON
	assert.Equal(t, rawResponse["status"], response.Status.Bool(), "status field should match")
	assert.Equal(t, rawResponse["message"], response.Message, "message field should match")

	// Validate data array
	rawData := rawResponse["data"].([]any)
	assert.Len(t, response.Data, len(rawData), "data array length should match")

	// Validate each customer in the list
	for i, rawCustomer := range rawData {
		customerData := rawCustomer.(map[string]any)
		customer := response.Data[i]

		// Basic fields
		assert.Equal(t, customerData["email"], customer.Email, "email should match for customer %d", i)
		assert.Equal(t, int(customerData["integration"].(float64)), *customer.Integration, "integration should match for customer %d", i)
		assert.Equal(t, customerData["domain"], customer.Domain, "domain should match for customer %d", i)
		assert.Equal(t, customerData["customer_code"], customer.CustomerCode, "customer_code should match for customer %d", i)
		assert.Equal(t, uint64(customerData["id"].(float64)), customer.ID, "id should match for customer %d", i)
		assert.Equal(t, customerData["risk_action"], customer.RiskAction, "risk_action should match for customer %d", i)

		// Handle nullable fields
		if customerData["first_name"] == nil {
			assert.Nil(t, customer.FirstName, "first_name should be nil for customer %d", i)
		} else {
			assert.Equal(t, customerData["first_name"], *customer.FirstName, "first_name should match for customer %d", i)
		}

		if customerData["last_name"] == nil {
			assert.Nil(t, customer.LastName, "last_name should be nil for customer %d", i)
		} else {
			assert.Equal(t, customerData["last_name"], *customer.LastName, "last_name should match for customer %d", i)
		}

		if customerData["phone"] == nil {
			assert.Nil(t, customer.Phone, "phone should be nil for customer %d", i)
		} else {
			assert.Equal(t, customerData["phone"], *customer.Phone, "phone should match for customer %d", i)
		}

		// Handle metadata - null or empty object in JSON becomes empty struct
		if customerData["metadata"] == nil {
			assert.Equal(t, map[string]any{}, map[string]any(customer.Metadata), "metadata should be empty for null for customer %d", i)
		} else {
			assert.Equal(t, customerData["metadata"], map[string]any(customer.Metadata), "metadata should match for customer %d", i)
		}

		// For timestamp comparisons, parse both and compare the actual time values
		expectedCreatedAt, err := time.Parse(time.RFC3339, customerData["createdAt"].(string))
		require.NoError(t, err, "should parse expected createdAt for customer %d", i)
		assert.True(t, expectedCreatedAt.Equal(customer.CreatedAt.Time), "createdAt should represent the same moment for customer %d", i)

		expectedUpdatedAt, err := time.Parse(time.RFC3339, customerData["updatedAt"].(string))
		require.NoError(t, err, "should parse expected updatedAt for customer %d", i)
		assert.True(t, expectedUpdatedAt.Equal(customer.UpdatedAt.Time), "updatedAt should represent the same moment for customer %d", i)
	}

	// Validate meta object if present
	if rawMeta, exists := rawResponse["meta"]; exists {
		metaData := rawMeta.(map[string]any)

		// We don't have a meta field in the struct, but we can validate it exists in the raw response
		assert.NotNil(t, metaData["next"], "meta.next should be present")
		assert.Nil(t, metaData["previous"], "meta.previous should be null")
		assert.Equal(t, float64(3), metaData["perPage"], "meta.perPage should match")
	}

	// Test round-trip serialization
	serialized, err := json.Marshal(response)
	require.NoError(t, err, "failed to marshal response back to JSON")

	var roundTripResponse ListResponse
	err = json.Unmarshal(serialized, &roundTripResponse)
	require.NoError(t, err, "failed to unmarshal round-trip JSON")

	// Verify core fields survive round-trip
	assert.Equal(t, response.Status.Bool(), roundTripResponse.Status.Bool(), "status should survive round-trip")
	assert.Equal(t, response.Message, roundTripResponse.Message, "message should survive round-trip")
	assert.Len(t, roundTripResponse.Data, len(response.Data), "data array length should survive round-trip")

	if len(response.Data) > 0 {
		assert.Equal(t, response.Data[0].Email, roundTripResponse.Data[0].Email, "first customer email should survive round-trip")
		assert.Equal(t, response.Data[0].CustomerCode, roundTripResponse.Data[0].CustomerCode, "first customer code should survive round-trip")
	}
}
