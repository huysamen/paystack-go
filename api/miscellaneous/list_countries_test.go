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
	assert.Len(t, response.Data, 7, "should have 7 countries in test data")

	// Validate Nigeria (first country)
	nigeria := response.Data[0]
	assert.Equal(t, int64(1), nigeria.ID.Int64())
	assert.True(t, nigeria.ActiveForDashboardOnboarding.Bool())
	assert.Equal(t, "Nigeria", nigeria.Name.String())
	assert.Equal(t, "NG", nigeria.ISOCode.String())
	assert.Equal(t, "NGN", nigeria.DefaultCurrencyCode.String())
	assert.Equal(t, "+234", nigeria.CallingCode.String())
	assert.False(t, nigeria.PilotMode.Bool())
	require.NotNil(t, nigeria.Relationships)

	// Validate Nigeria currency relationships
	require.NotNil(t, nigeria.Relationships.Currency)
	assert.Equal(t, "currency", nigeria.Relationships.Currency.Type.String())

	// Convert data.String slice to regular string slice for assertion
	nigeriacurrencies := make([]string, len(nigeria.Relationships.Currency.Data))
	for i, curr := range nigeria.Relationships.Currency.Data {
		nigeriacurrencies[i] = curr.String()
	}
	assert.Contains(t, nigeriacurrencies, "NGN")

	// Validate supported currencies
	require.NotNil(t, nigeria.Relationships.Currency.SupportedCurrencies)
	assert.Contains(t, nigeria.Relationships.Currency.SupportedCurrencies, "NGN")
	assert.Contains(t, nigeria.Relationships.Currency.SupportedCurrencies, "USD")

	// Validate Nigeria integration type relationships
	require.NotNil(t, nigeria.Relationships.IntegrationType)
	assert.Equal(t, "integration_type", nigeria.Relationships.IntegrationType.Type.String())

	// Convert data.String slice to regular string slice for assertion
	nigeriaIntegrationTypes := make([]string, len(nigeria.Relationships.IntegrationType.Data))
	for i, intType := range nigeria.Relationships.IntegrationType.Data {
		nigeriaIntegrationTypes[i] = intType.String()
	}
	assert.Contains(t, nigeriaIntegrationTypes, "ITYPE_001")
	assert.Contains(t, nigeriaIntegrationTypes, "ITYPE_002")
	assert.Contains(t, nigeriaIntegrationTypes, "ITYPE_003")

	// Validate payment method relationships
	require.NotNil(t, nigeria.Relationships.PaymentMethod)
	assert.Equal(t, "payment_method", nigeria.Relationships.PaymentMethod.Type.String())

	// Validate Ghana (second country)
	ghana := response.Data[1]
	assert.Equal(t, int64(2), ghana.ID.Int64())
	assert.True(t, ghana.ActiveForDashboardOnboarding.Bool())
	assert.Equal(t, "Ghana", ghana.Name.String())
	assert.Equal(t, "GH", ghana.ISOCode.String())
	assert.Equal(t, "GHS", ghana.DefaultCurrencyCode.String())
	assert.Equal(t, "+233", ghana.CallingCode.String())
	assert.False(t, ghana.PilotMode.Bool())
	require.NotNil(t, ghana.Relationships)

	// Validate Ghana currency relationships
	require.NotNil(t, ghana.Relationships.Currency)
	assert.Equal(t, "currency", ghana.Relationships.Currency.Type.String())

	// Convert data.String slice to regular string slice for assertion
	ghanaCurrencies := make([]string, len(ghana.Relationships.Currency.Data))
	for i, curr := range ghana.Relationships.Currency.Data {
		ghanaCurrencies[i] = curr.String()
	}
	assert.Contains(t, ghanaCurrencies, "GHS")

	// Validate supported currencies for Ghana
	require.NotNil(t, ghana.Relationships.Currency.SupportedCurrencies)
	assert.Contains(t, ghana.Relationships.Currency.SupportedCurrencies, "GHS")
	assert.Contains(t, ghana.Relationships.Currency.SupportedCurrencies, "USD")

	// Validate GHS bank configuration for Ghana
	ghsConfig, exists := ghana.Relationships.Currency.SupportedCurrencies["GHS"]
	require.True(t, exists, "GHS configuration should exist")
	require.NotNil(t, ghsConfig.Bank, "GHS bank configuration should exist")
	assert.Equal(t, "ghipss", ghsConfig.Bank.BankType.String())
	assert.True(t, ghsConfig.Bank.AccountVerificationRequired.Bool())

	// Validate mobile money configuration for Ghana
	require.NotNil(t, ghsConfig.MobileMoney, "GHS mobile money configuration should exist")
	assert.Equal(t, "mobile_money", ghsConfig.MobileMoney.BankType.String())
	assert.Equal(t, "phoneNumber", ghsConfig.MobileMoney.PhoneNumberLabel.String())

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
