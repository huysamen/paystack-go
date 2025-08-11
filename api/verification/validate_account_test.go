package verification

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestVerification_ValidateAccount_JSONDeserialization(t *testing.T) {
	p := filepath.Join("..", "..", "resources", "examples", "responses", "verification", "validate_account_200.json")
	b, err := os.ReadFile(p)
	require.NoError(t, err)

	var rsp ValidateAccountResponse
	err = json.Unmarshal(b, &rsp)
	require.NoError(t, err)
	assert.True(t, rsp.Status.Bool())
}

func TestVerification_ValidateAccount_Builder(t *testing.T) {
	r := NewValidateAccountRequestBuilder("n", "123", "personal", "058", "NG", "id").DocumentNumber("999").Build()
	assert.Equal(t, "n", r.AccountName)
	assert.NotNil(t, r.DocumentNumber)
}
