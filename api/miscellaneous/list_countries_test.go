package miscellaneous

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListCountriesResponse_JSONDeserialization(t *testing.T) {
	tests := []struct {
		name         string
		responseFile string
	}{
		{
			name:         "successful list countries response",
			responseFile: "list_countries_200.json",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Read the response file
			responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "miscellaneous", tt.responseFile)
			responseData, err := os.ReadFile(responseFilePath)
			require.NoError(t, err, "failed to read response file")

			// Deserialize the JSON response
			var response ListCountriesResponse
			err = json.Unmarshal(responseData, &response)
			require.NoError(t, err, "failed to unmarshal JSON response")

			// Basic response validation
			assert.True(t, response.Status.Bool(), "status should be true")
			assert.Equal(t, "Countries retrieved", response.Message, "message should match")
			require.NotNil(t, response.Data, "data should not be nil")
			assert.Greater(t, len(response.Data), 0, "should have countries in response")
		})
	}
}

func TestListCountriesResponse_FieldByFieldValidation(t *testing.T) {
	// Read the response file
	responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "miscellaneous", "list_countries_200.json")
	responseData, err := os.ReadFile(responseFilePath)
	require.NoError(t, err, "failed to read response file")

	// Deserialize the JSON response
	var response ListCountriesResponse
	err = json.Unmarshal(responseData, &response)
	require.NoError(t, err, "failed to unmarshal JSON response")

	// Validate response structure
	assert.True(t, response.Status.Bool())
	assert.Equal(t, "Countries retrieved", response.Message)
	require.NotNil(t, response.Data)
	assert.Len(t, response.Data, 2, "should have 2 countries in test data")

	// Validate Nigeria (first country)
	nigeria := response.Data[0]
	assert.Equal(t, 1, nigeria.ID)
	assert.Equal(t, "Nigeria", nigeria.Name)
	assert.Equal(t, "NG", nigeria.ISOCode)
	assert.Equal(t, "NGN", nigeria.DefaultCurrencyCode)
	require.NotNil(t, nigeria.Relationships)

	// Validate Nigeria currency relationships
	require.NotNil(t, nigeria.Relationships.Currency)
	assert.Equal(t, "currency", nigeria.Relationships.Currency.Type)
	assert.Contains(t, nigeria.Relationships.Currency.Data, "NGN")
	assert.Contains(t, nigeria.Relationships.Currency.Data, "USD")

	// Validate Nigeria integration type relationships
	require.NotNil(t, nigeria.Relationships.IntegrationType)
	assert.Equal(t, "integration_type", nigeria.Relationships.IntegrationType.Type)
	assert.Contains(t, nigeria.Relationships.IntegrationType.Data, "ITYPE_001")
	assert.Contains(t, nigeria.Relationships.IntegrationType.Data, "ITYPE_002")
	assert.Contains(t, nigeria.Relationships.IntegrationType.Data, "ITYPE_003")

	// Validate Ghana (second country)
	ghana := response.Data[1]
	assert.Equal(t, 2, ghana.ID)
	assert.Equal(t, "Ghana", ghana.Name)
	assert.Equal(t, "GH", ghana.ISOCode)
	assert.Equal(t, "GHS", ghana.DefaultCurrencyCode)
	require.NotNil(t, ghana.Relationships)

	// Validate Ghana currency relationships
	require.NotNil(t, ghana.Relationships.Currency)
	assert.Equal(t, "currency", ghana.Relationships.Currency.Type)
	assert.Contains(t, ghana.Relationships.Currency.Data, "GHS")
	assert.Contains(t, ghana.Relationships.Currency.Data, "USD")

	// Test JSON round-trip
	marshaled, err := json.Marshal(response)
	require.NoError(t, err)

	var roundTrip ListCountriesResponse
	err = json.Unmarshal(marshaled, &roundTrip)
	require.NoError(t, err)

	assert.Equal(t, response.Status, roundTrip.Status)
	assert.Equal(t, response.Message, roundTrip.Message)
	assert.Len(t, roundTrip.Data, len(response.Data))
	if len(response.Data) > 0 {
		assert.Equal(t, response.Data[0].Name, roundTrip.Data[0].Name)
		assert.Equal(t, response.Data[0].ISOCode, roundTrip.Data[0].ISOCode)
	}
}
