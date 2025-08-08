package miscellaneous

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListBanksResponse_JSONDeserialization(t *testing.T) {
	tests := []struct {
		name         string
		responseFile string
	}{
		{
			name:         "successful list banks response",
			responseFile: "list_banks_200.json",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Read the response file
			responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "miscellaneous", tt.responseFile)
			responseData, err := os.ReadFile(responseFilePath)
			require.NoError(t, err, "failed to read response file")

			// Deserialize the JSON response
			var response ListBanksResponse
			err = json.Unmarshal(responseData, &response)
			require.NoError(t, err, "failed to unmarshal JSON response")

			// Basic response validation
			assert.True(t, response.Status.Bool(), "status should be true")
			assert.Equal(t, "Banks retrieved", response.Message, "message should match")
			require.NotNil(t, response.Data, "data should not be nil")
			assert.Greater(t, len(response.Data), 0, "should have banks in response")
		})
	}
}

func TestListBanksRequestBuilder(t *testing.T) {
	tests := []struct {
		name           string
		setupBuilder   func() *ListBanksRequestBuilder
		expectedFields map[string]interface{}
	}{
		{
			name: "builds empty request",
			setupBuilder: func() *ListBanksRequestBuilder {
				return NewListBanksRequestBuilder()
			},
			expectedFields: map[string]interface{}{},
		},
		{
			name: "builds request with country filter",
			setupBuilder: func() *ListBanksRequestBuilder {
				return NewListBanksRequestBuilder().Country("nigeria")
			},
			expectedFields: map[string]interface{}{
				"country": "nigeria",
			},
		},
		{
			name: "builds request with cursor pagination",
			setupBuilder: func() *ListBanksRequestBuilder {
				return NewListBanksRequestBuilder().UseCursor(true).PerPage(25)
			},
			expectedFields: map[string]interface{}{
				"use_cursor": true,
				"perPage":    25,
			},
		},
		{
			name: "builds request with payment filters",
			setupBuilder: func() *ListBanksRequestBuilder {
				return NewListBanksRequestBuilder().
					PayWithBank(true).
					PayWithBankTransfer(false).
					EnabledForVerification(true)
			},
			expectedFields: map[string]interface{}{
				"pay_with_bank":            true,
				"pay_with_bank_transfer":   false,
				"enabled_for_verification": true,
			},
		},
		{
			name: "builds request with pagination cursors",
			setupBuilder: func() *ListBanksRequestBuilder {
				return NewListBanksRequestBuilder().
					Next("next_cursor").
					Previous("prev_cursor")
			},
			expectedFields: map[string]interface{}{
				"next":     "next_cursor",
				"previous": "prev_cursor",
			},
		},
		{
			name: "builds request with type and gateway filters",
			setupBuilder: func() *ListBanksRequestBuilder {
				return NewListBanksRequestBuilder().
					Gateway("emandate").
					Type("nuban").
					Currency("NGN")
			},
			expectedFields: map[string]interface{}{
				"gateway":  "emandate",
				"type":     "nuban",
				"currency": "NGN",
			},
		},
		{
			name: "builds request with NIP sort code inclusion",
			setupBuilder: func() *ListBanksRequestBuilder {
				return NewListBanksRequestBuilder().IncludeNIPSortCode(true)
			},
			expectedFields: map[string]interface{}{
				"include_nip_sort_code": true,
			},
		},
		{
			name: "builds complete request with all fields",
			setupBuilder: func() *ListBanksRequestBuilder {
				return NewListBanksRequestBuilder().
					Country("nigeria").
					UseCursor(true).
					PerPage(50).
					PayWithBank(true).
					PayWithBankTransfer(true).
					EnabledForVerification(true).
					Next("next_token").
					Previous("prev_token").
					Gateway("emandate").
					Type("nuban").
					Currency("NGN").
					IncludeNIPSortCode(true)
			},
			expectedFields: map[string]interface{}{
				"country":                  "nigeria",
				"use_cursor":               true,
				"perPage":                  50,
				"pay_with_bank":            true,
				"pay_with_bank_transfer":   true,
				"enabled_for_verification": true,
				"next":                     "next_token",
				"previous":                 "prev_token",
				"gateway":                  "emandate",
				"type":                     "nuban",
				"currency":                 "NGN",
				"include_nip_sort_code":    true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			builder := tt.setupBuilder()
			require.NotNil(t, builder, "builder should not be nil")

			req := builder.Build()
			require.NotNil(t, req, "built request should not be nil")

			// Verify expected fields
			for fieldName, expectedValue := range tt.expectedFields {
				switch fieldName {
				case "country":
					if expectedValue != nil {
						assert.Equal(t, expectedValue, *req.Country)
					} else {
						assert.Nil(t, req.Country)
					}
				case "use_cursor":
					if expectedValue != nil {
						assert.Equal(t, expectedValue, *req.UseCursor)
					} else {
						assert.Nil(t, req.UseCursor)
					}
				case "perPage":
					if expectedValue != nil {
						assert.Equal(t, expectedValue, *req.PerPage)
					} else {
						assert.Nil(t, req.PerPage)
					}
				case "pay_with_bank":
					if expectedValue != nil {
						assert.Equal(t, expectedValue, *req.PayWithBank)
					} else {
						assert.Nil(t, req.PayWithBank)
					}
				case "pay_with_bank_transfer":
					if expectedValue != nil {
						assert.Equal(t, expectedValue, *req.PayWithBankTransfer)
					} else {
						assert.Nil(t, req.PayWithBankTransfer)
					}
				case "enabled_for_verification":
					if expectedValue != nil {
						assert.Equal(t, expectedValue, *req.EnabledForVerification)
					} else {
						assert.Nil(t, req.EnabledForVerification)
					}
				case "next":
					if expectedValue != nil {
						assert.Equal(t, expectedValue, *req.Next)
					} else {
						assert.Nil(t, req.Next)
					}
				case "previous":
					if expectedValue != nil {
						assert.Equal(t, expectedValue, *req.Previous)
					} else {
						assert.Nil(t, req.Previous)
					}
				case "gateway":
					if expectedValue != nil {
						assert.Equal(t, expectedValue, *req.Gateway)
					} else {
						assert.Nil(t, req.Gateway)
					}
				case "type":
					if expectedValue != nil {
						assert.Equal(t, expectedValue, *req.Type)
					} else {
						assert.Nil(t, req.Type)
					}
				case "currency":
					if expectedValue != nil {
						assert.Equal(t, expectedValue, *req.Currency)
					} else {
						assert.Nil(t, req.Currency)
					}
				case "include_nip_sort_code":
					if expectedValue != nil {
						assert.Equal(t, expectedValue, *req.IncludeNIPSortCode)
					} else {
						assert.Nil(t, req.IncludeNIPSortCode)
					}
				}
			}
		})
	}
}

func TestListBanksRequest_QueryGeneration(t *testing.T) {
	tests := []struct {
		name          string
		setupBuilder  func() *ListBanksRequestBuilder
		expectedQuery string
	}{
		{
			name: "generates empty query for empty request",
			setupBuilder: func() *ListBanksRequestBuilder {
				return NewListBanksRequestBuilder()
			},
			expectedQuery: "",
		},
		{
			name: "generates query with country only",
			setupBuilder: func() *ListBanksRequestBuilder {
				return NewListBanksRequestBuilder().Country("nigeria")
			},
			expectedQuery: "country=nigeria",
		},
		{
			name: "generates query with pagination parameters",
			setupBuilder: func() *ListBanksRequestBuilder {
				return NewListBanksRequestBuilder().
					UseCursor(true).
					PerPage(25)
			},
			expectedQuery: "perPage=25&use_cursor=true",
		},
		{
			name: "generates query with payment filters",
			setupBuilder: func() *ListBanksRequestBuilder {
				return NewListBanksRequestBuilder().
					PayWithBank(true).
					PayWithBankTransfer(false).
					EnabledForVerification(true)
			},
			expectedQuery: "enabled_for_verification=true&pay_with_bank=true&pay_with_bank_transfer=false",
		},
		{
			name: "generates query with all parameters",
			setupBuilder: func() *ListBanksRequestBuilder {
				return NewListBanksRequestBuilder().
					Country("nigeria").
					UseCursor(true).
					PerPage(50).
					PayWithBank(true).
					EnabledForVerification(true).
					Gateway("emandate").
					Type("nuban").
					Currency("NGN").
					IncludeNIPSortCode(true)
			},
			expectedQuery: "country=nigeria&currency=NGN&enabled_for_verification=true&gateway=emandate&include_nip_sort_code=true&pay_with_bank=true&perPage=50&type=nuban&use_cursor=true",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			builder := tt.setupBuilder()
			req := builder.Build()

			query := req.toQuery()
			assert.Equal(t, tt.expectedQuery, query, "query should match expected")
		})
	}
}

func TestListBanksResponse_FieldByFieldValidation(t *testing.T) {
	// Read the response file
	responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "miscellaneous", "list_banks_200.json")
	responseData, err := os.ReadFile(responseFilePath)
	require.NoError(t, err, "failed to read response file")

	// Deserialize the JSON response
	var response ListBanksResponse
	err = json.Unmarshal(responseData, &response)
	require.NoError(t, err, "failed to unmarshal JSON response")

	// Validate response structure
	assert.True(t, response.Status.Bool())
	assert.Equal(t, "Banks retrieved", response.Message)
	require.NotNil(t, response.Data)
	assert.Greater(t, len(response.Data), 0, "should have banks in response")

	// Validate first bank structure
	bank := response.Data[0]
	assert.Equal(t, "Abbey Mortgage Bank", bank.Name)
	assert.Equal(t, "abbey-mortgage-bank", bank.Slug)
	assert.Equal(t, "801", bank.Code)
	assert.Equal(t, "", bank.LongCode)
	assert.Nil(t, bank.Gateway)
	assert.False(t, bank.PayWithBank)
	assert.True(t, bank.Active)
	assert.Equal(t, "Nigeria", bank.Country)
	assert.Equal(t, "NGN", string(bank.Currency))
	assert.Equal(t, "nuban", bank.Type)
	assert.Equal(t, 174, bank.ID)

	// Validate metadata if present
	if response.Meta != nil {
		meta := response.Meta
		assert.NotEmpty(t, meta.Next, "next cursor should be present")
		assert.Nil(t, meta.Previous, "previous cursor should be nil for first page")
		assert.Equal(t, 5, meta.PerPage, "per page should match")
	}

	// Test JSON round-trip
	marshaled, err := json.Marshal(response)
	require.NoError(t, err)

	var roundTrip ListBanksResponse
	err = json.Unmarshal(marshaled, &roundTrip)
	require.NoError(t, err)

	assert.Equal(t, response.Status, roundTrip.Status)
	assert.Equal(t, response.Message, roundTrip.Message)
	assert.Len(t, roundTrip.Data, len(response.Data))
	if len(response.Data) > 0 {
		assert.Equal(t, response.Data[0].Name, roundTrip.Data[0].Name)
	}
}
