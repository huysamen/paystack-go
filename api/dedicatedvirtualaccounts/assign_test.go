package dedicatedvirtualaccounts

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAssignResponse_JSONDeserialization(t *testing.T) {
	tests := []struct {
		name            string
		responseFile    string
		expectedStatus  bool
		expectedMessage string
	}{
		{
			name:            "successful assign dedicated virtual account response",
			responseFile:    "assign_200.json",
			expectedStatus:  true,
			expectedMessage: "Assign dedicated account in progress",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Read the response file
			responseFilePath := filepath.Join("..", "..", "resources", "examples", "responses", "dedicatedvirtualaccounts", tt.responseFile)
			responseData, err := os.ReadFile(responseFilePath)
			require.NoError(t, err, "failed to read response file: %s", responseFilePath)

			// Deserialize the JSON response
			var response AssignDedicatedVirtualAccountResponse
			err = json.Unmarshal(responseData, &response)
			require.NoError(t, err, "failed to unmarshal JSON response")

			// Verify the response structure
			assert.Equal(t, tt.expectedStatus, response.Status.Bool(), "status should match")
			assert.Equal(t, tt.expectedMessage, response.Message, "message should match")
		})
	}
}

func TestAssignRequestBuilder(t *testing.T) {
	t.Run("builds request with required fields", func(t *testing.T) {
		builder := NewAssignRequestBuilder().
			Email("test@example.com").
			FirstName("John").
			LastName("Doe").
			Phone("+2348012345678").
			PreferredBank("wema-bank").
			Country("NG")

		request := builder.Build()

		assert.Equal(t, "test@example.com", request.Email, "email should match")
		assert.Equal(t, "John", request.FirstName, "first name should match")
		assert.Equal(t, "Doe", request.LastName, "last name should match")
		assert.Equal(t, "+2348012345678", request.Phone, "phone should match")
		assert.Equal(t, "wema-bank", request.PreferredBank, "preferred bank should match")
		assert.Equal(t, "NG", request.Country, "country should match")
	})

	t.Run("builds request with all fields", func(t *testing.T) {
		builder := NewAssignRequestBuilder().
			Email("test@example.com").
			FirstName("John").
			LastName("Doe").
			Phone("+2348012345678").
			PreferredBank("wema-bank").
			Country("NG").
			AccountNumber("0123456789").
			BVN("12345678901").
			BankCode("058").
			Subaccount("ACCT_subaccount_code").
			SplitCode("SPL_split_code").
			MiddleName("Michael")

		request := builder.Build()

		assert.Equal(t, "test@example.com", request.Email, "email should match")
		assert.Equal(t, "John", request.FirstName, "first name should match")
		assert.Equal(t, "Doe", request.LastName, "last name should match")
		assert.Equal(t, "+2348012345678", request.Phone, "phone should match")
		assert.Equal(t, "wema-bank", request.PreferredBank, "preferred bank should match")
		assert.Equal(t, "NG", request.Country, "country should match")
		assert.Equal(t, "0123456789", request.AccountNumber, "account number should match")
		assert.Equal(t, "12345678901", request.BVN, "BVN should match")
		assert.Equal(t, "058", request.BankCode, "bank code should match")
		assert.Equal(t, "ACCT_subaccount_code", request.Subaccount, "subaccount should match")
		assert.Equal(t, "SPL_split_code", request.SplitCode, "split code should match")
		assert.Equal(t, "Michael", request.MiddleName, "middle name should match")
	})

	t.Run("builds request with empty values", func(t *testing.T) {
		builder := NewAssignRequestBuilder()
		request := builder.Build()

		assert.Empty(t, request.Email, "email should be empty")
		assert.Empty(t, request.FirstName, "first name should be empty")
		assert.Empty(t, request.LastName, "last name should be empty")
		assert.Empty(t, request.Phone, "phone should be empty")
		assert.Empty(t, request.PreferredBank, "preferred bank should be empty")
		assert.Empty(t, request.Country, "country should be empty")
		assert.Empty(t, request.AccountNumber, "account number should be empty")
		assert.Empty(t, request.BVN, "BVN should be empty")
		assert.Empty(t, request.BankCode, "bank code should be empty")
		assert.Empty(t, request.Subaccount, "subaccount should be empty")
		assert.Empty(t, request.SplitCode, "split code should be empty")
		assert.Empty(t, request.MiddleName, "middle name should be empty")
	})
}

func TestAssignRequest_JSONSerialization(t *testing.T) {
	t.Run("serializes request correctly", func(t *testing.T) {
		builder := NewAssignRequestBuilder().
			Email("test@example.com").
			FirstName("John").
			LastName("Doe").
			Phone("+2348012345678").
			PreferredBank("wema-bank").
			Country("NG")

		request := builder.Build()

		jsonData, err := json.Marshal(request)
		require.NoError(t, err, "should marshal to JSON without error")

		var unmarshaled map[string]interface{}
		err = json.Unmarshal(jsonData, &unmarshaled)
		require.NoError(t, err, "should unmarshal JSON without error")

		assert.Equal(t, "test@example.com", unmarshaled["email"], "email should match")
		assert.Equal(t, "John", unmarshaled["first_name"], "first name should match")
		assert.Equal(t, "Doe", unmarshaled["last_name"], "last name should match")
		assert.Equal(t, "+2348012345678", unmarshaled["phone"], "phone should match")
		assert.Equal(t, "wema-bank", unmarshaled["preferred_bank"], "preferred bank should match")
		assert.Equal(t, "NG", unmarshaled["country"], "country should match")
	})

	t.Run("omits empty optional fields in JSON", func(t *testing.T) {
		builder := NewAssignRequestBuilder().
			Email("test@example.com").
			FirstName("John").
			LastName("Doe").
			Phone("+2348012345678").
			PreferredBank("wema-bank").
			Country("NG")

		request := builder.Build()

		jsonData, err := json.Marshal(request)
		require.NoError(t, err, "should marshal to JSON without error")

		var unmarshaled map[string]interface{}
		err = json.Unmarshal(jsonData, &unmarshaled)
		require.NoError(t, err, "should unmarshal JSON without error")

		// Required fields should be present
		assert.Equal(t, "test@example.com", unmarshaled["email"], "email should match")
		assert.Equal(t, "John", unmarshaled["first_name"], "first name should match")

		// Optional fields should not be present or should be empty
		_, hasAccountNumber := unmarshaled["account_number"]
		if hasAccountNumber {
			assert.Empty(t, unmarshaled["account_number"], "account number should be empty if present")
		}
	})
}
