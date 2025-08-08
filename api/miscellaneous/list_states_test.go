package miscellaneous

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListStatesResponse_JSONDeserialization(t *testing.T) {
	tests := []struct {
		name         string
		responseFile string
	}{
		{
			name:         "successful list states response",
			responseFile: "list_states_200.json",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Read the response file
			responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "miscellaneous", tt.responseFile)
			responseData, err := os.ReadFile(responseFilePath)
			require.NoError(t, err, "failed to read response file")

			// Deserialize the JSON response
			var response ListStatesResponse
			err = json.Unmarshal(responseData, &response)
			require.NoError(t, err, "failed to unmarshal JSON response")

			// Basic response validation
			assert.True(t, response.Status.Bool(), "status should be true")
			assert.Equal(t, "States retrieved", response.Message, "message should match")
			require.NotNil(t, response.Data, "data should not be nil")
			assert.Greater(t, len(response.Data), 0, "should have states in response")
		})
	}
}

func TestListStatesRequestBuilder(t *testing.T) {
	tests := []struct {
		name            string
		country         string
		expectedCountry string
	}{
		{
			name:            "builds request with canada",
			country:         "canada",
			expectedCountry: "canada",
		},
		{
			name:            "builds request with nigeria",
			country:         "nigeria",
			expectedCountry: "nigeria",
		},
		{
			name:            "builds request with usa",
			country:         "usa",
			expectedCountry: "usa",
		},
		{
			name:            "builds request with ghana",
			country:         "ghana",
			expectedCountry: "ghana",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			builder := NewListStatesRequestBuilder(tt.country)
			require.NotNil(t, builder, "builder should not be nil")

			req := builder.Build()
			require.NotNil(t, req, "built request should not be nil")
			assert.Equal(t, tt.expectedCountry, req.Country, "country should match expected value")
		})
	}
}

func TestListStatesRequest_QueryGeneration(t *testing.T) {
	tests := []struct {
		name          string
		country       string
		expectedQuery string
	}{
		{
			name:          "generates query for canada",
			country:       "canada",
			expectedQuery: "country=canada",
		},
		{
			name:          "generates query for nigeria",
			country:       "nigeria",
			expectedQuery: "country=nigeria",
		},
		{
			name:          "generates query for usa",
			country:       "usa",
			expectedQuery: "country=usa",
		},
		{
			name:          "generates query with special characters",
			country:       "south africa",
			expectedQuery: "country=south+africa",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			builder := NewListStatesRequestBuilder(tt.country)
			req := builder.Build()

			query := req.toQuery()
			assert.Equal(t, tt.expectedQuery, query, "query should match expected")
		})
	}
}

func TestListStatesResponse_FieldByFieldValidation(t *testing.T) {
	// Read the response file
	responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "miscellaneous", "list_states_200.json")
	responseData, err := os.ReadFile(responseFilePath)
	require.NoError(t, err, "failed to read response file")

	// Deserialize the JSON response
	var response ListStatesResponse
	err = json.Unmarshal(responseData, &response)
	require.NoError(t, err, "failed to unmarshal JSON response")

	// Validate response structure
	assert.True(t, response.Status.Bool())
	assert.Equal(t, "States retrieved", response.Message)
	require.NotNil(t, response.Data)
	assert.Len(t, response.Data, 2, "should have 2 states in test data")

	// Validate Alberta (first state)
	alberta := response.Data[0]
	assert.Equal(t, "Alberta", alberta.Name)
	assert.Equal(t, "alberta", alberta.Slug)
	assert.Equal(t, "AB", alberta.Abbreviation)

	// Validate British Columbia (second state)
	bc := response.Data[1]
	assert.Equal(t, "British Columbia", bc.Name)
	assert.Equal(t, "british-columbia", bc.Slug)
	assert.Equal(t, "BC", bc.Abbreviation)

	// Test JSON round-trip
	marshaled, err := json.Marshal(response)
	require.NoError(t, err)

	var roundTrip ListStatesResponse
	err = json.Unmarshal(marshaled, &roundTrip)
	require.NoError(t, err)

	assert.Equal(t, response.Status, roundTrip.Status)
	assert.Equal(t, response.Message, roundTrip.Message)
	assert.Len(t, roundTrip.Data, len(response.Data))
	if len(response.Data) > 0 {
		assert.Equal(t, response.Data[0].Name, roundTrip.Data[0].Name)
		assert.Equal(t, response.Data[0].Abbreviation, roundTrip.Data[0].Abbreviation)
	}
}
