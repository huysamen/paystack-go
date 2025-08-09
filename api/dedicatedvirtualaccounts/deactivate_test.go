package dedicatedvirtualaccounts

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDeactivateResponse_JSONDeserialization(t *testing.T) {
	tests := []struct {
		name            string
		responseFile    string
		expectedStatus  bool
		expectedMessage string
	}{
		{
			name:            "successful deactivate dedicated virtual account response",
			responseFile:    "deactivate_200.json",
			expectedStatus:  true,
			expectedMessage: "Managed Account Successfully Unassigned",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Read the response file
			responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "dedicatedvirtualaccounts", tt.responseFile)
			responseData, err := os.ReadFile(responseFilePath)
			require.NoError(t, err, "failed to read response file: %s", responseFilePath)

			// Deserialize the JSON response
			var response DeactivateResponse
			err = json.Unmarshal(responseData, &response)
			require.NoError(t, err, "failed to unmarshal JSON response")

			// Verify the response structure
			assert.Equal(t, tt.expectedStatus, response.Status.Bool(), "status should match")
			assert.Equal(t, tt.expectedMessage, response.Message, "message should match")

			// Only verify data structure for successful responses
			if tt.expectedStatus {
				assert.NotEmpty(t, response.Data.AccountNumber, "account number should not be empty")
				assert.NotEmpty(t, response.Data.AccountName, "account name should not be empty")
				assert.NotEmpty(t, response.Data.Currency, "currency should not be empty")
				assert.Greater(t, response.Data.ID.Int64(), int64(0), "account ID should be greater than 0")
				assert.True(t, response.Data.Active.Bool(), "account should be active")
				// Note: assigned should be false for deactivated account
				assert.False(t, response.Data.Assigned.Bool(), "account should not be assigned after deactivation")

				// Verify bank data
				assert.NotEmpty(t, response.Data.Bank.Name, "bank name should not be empty")
				assert.NotEmpty(t, response.Data.Bank.Slug, "bank slug should not be empty")
				assert.Greater(t, response.Data.Bank.ID.Int64(), int64(0), "bank ID should be greater than 0")

				// Verify assignment data
				if response.Data.Assignment != nil {
					assert.NotEmpty(t, response.Data.Assignment.AccountType, "assignment account type should not be empty")
					assert.Greater(t, response.Data.Assignment.Integration.Int64(), int64(0), "assignment integration should be greater than 0")
					assert.Greater(t, response.Data.Assignment.AssigneeID.Int64(), int64(0), "assignment assignee ID should be greater than 0")
				}
			}
		})
	}
}

func TestDeactivateResponse_FieldByFieldValidation(t *testing.T) {
	// Read the deactivate_200.json file
	responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "dedicatedvirtualaccounts", "deactivate_200.json")
	responseData, err := os.ReadFile(responseFilePath)
	require.NoError(t, err, "failed to read deactivate_200.json")

	// Parse the raw JSON to get the original values
	var rawResponse map[string]any
	err = json.Unmarshal(responseData, &rawResponse)
	require.NoError(t, err, "failed to unmarshal raw JSON response")

	// Deserialize into the DeactivateResponse struct
	var response DeactivateResponse
	err = json.Unmarshal(responseData, &response)
	require.NoError(t, err, "failed to unmarshal into DeactivateResponse struct")

	// Validate top-level fields against the raw JSON
	assert.Equal(t, rawResponse["status"], response.Status.Bool(), "status field should match")
	assert.Equal(t, rawResponse["message"], response.Message, "message field should match")

	// Validate data object fields
	rawData := rawResponse["data"].(map[string]any)
	data := response.Data

	// Basic account fields
	assert.Equal(t, rawData["account_name"], data.AccountName.String(), "account_name should match")
	assert.Equal(t, rawData["account_number"], data.AccountNumber.String(), "account_number should match")
	assert.Equal(t, rawData["assigned"], data.Assigned.Bool(), "assigned should match")
	assert.Equal(t, rawData["currency"], data.Currency.String(), "currency should match")
	assert.Equal(t, rawData["active"], data.Active.Bool(), "active should match")
	assert.Equal(t, rawData["id"], float64(data.ID.Int64()), "id should match")

	// Metadata field (null in JSON)
	if rawData["metadata"] == nil {
		assert.False(t, data.Metadata.Valid, "metadata should be invalid when null in JSON")
	} else {
		assert.Equal(t, rawData["metadata"], data.Metadata.Metadata, "metadata should match")
	}

	// Bank object validation
	rawBank := rawData["bank"].(map[string]any)
	bank := data.Bank
	assert.Equal(t, rawBank["name"], bank.Name.String(), "bank.name should match")
	assert.Equal(t, rawBank["id"], float64(bank.ID.Int64()), "bank.id should match")
	assert.Equal(t, rawBank["slug"], bank.Slug.String(), "bank.slug should match")

	// Assignment object validation
	rawAssignment := rawData["assignment"].(map[string]any)
	assignment := data.Assignment
	require.NotNil(t, assignment, "assignment should not be nil")
	assert.Equal(t, rawAssignment["assignee_id"], float64(assignment.AssigneeID.Int64()), "assignment.assignee_id should match")
	assert.Equal(t, rawAssignment["assignee_type"], assignment.AssigneeType.String(), "assignment.assignee_type should match")
	assert.Equal(t, rawAssignment["integration"], float64(assignment.Integration.Int64()), "assignment.integration should match")
	assert.Equal(t, rawAssignment["account_type"], assignment.AccountType.String(), "assignment.account_type should match")

	// Timestamp validation using MultiDateTime
	createdAtStr, ok := rawData["created_at"].(string)
	require.True(t, ok, "created_at should be a string")
	parsedCreatedAt, err := time.Parse("2006-01-02T15:04:05.000Z", createdAtStr)
	require.NoError(t, err, "should parse created_at timestamp")
	assert.Equal(t, 2019, parsedCreatedAt.Year(), "created_at year should be 2019")
	assert.Equal(t, 2019, data.CreatedAt.Time().Year(), "data CreatedAt year should match")

	updatedAtStr, ok := rawData["updated_at"].(string)
	require.True(t, ok, "updated_at should be a string")
	parsedUpdatedAt, err := time.Parse("2006-01-02T15:04:05.000Z", updatedAtStr)
	require.NoError(t, err, "should parse updated_at timestamp")
	assert.Equal(t, 2020, parsedUpdatedAt.Year(), "updated_at year should be 2020")
	assert.Equal(t, 2020, data.UpdatedAt.Time().Year(), "data UpdatedAt year should match")

	// Assignment timestamp validation
	assignedAtStr, ok := rawAssignment["assigned_at"].(string)
	require.True(t, ok, "assigned_at should be a string")
	parsedAssignedAt, err := time.Parse("2006-01-02T15:04:05.000Z", assignedAtStr)
	require.NoError(t, err, "should parse assigned_at timestamp")
	assert.Equal(t, 2019, parsedAssignedAt.Year(), "assigned_at year should be 2019")
	assert.Equal(t, 2019, assignment.AssignedAt.Time().Year(), "assignment AssignedAt year should match")

	// Test round-trip serialization
	serialized, err := json.Marshal(response)
	require.NoError(t, err, "failed to marshal response back to JSON")

	var roundTripResponse DeactivateResponse
	err = json.Unmarshal(serialized, &roundTripResponse)
	require.NoError(t, err, "failed to unmarshal round-trip JSON")

	// Verify core fields survive round-trip
	assert.Equal(t, response.Status.Bool(), roundTripResponse.Status.Bool(), "status should survive round-trip")
	assert.Equal(t, response.Message, roundTripResponse.Message, "message should survive round-trip")
	assert.Equal(t, response.Data.AccountNumber, roundTripResponse.Data.AccountNumber, "account_number should survive round-trip")
	assert.Equal(t, response.Data.Assigned, roundTripResponse.Data.Assigned, "assigned should survive round-trip")
}
