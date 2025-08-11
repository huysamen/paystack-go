package transactions

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTransactions_PartialDebit_JSONDeserialization(t *testing.T) {
	p := filepath.Join("..", "..", "resources", "examples", "responses", "transactions", "partial_debit_200.json")
	b, err := os.ReadFile(p)
	require.NoError(t, err)

	var rsp PartialDebitResponse
	err = json.Unmarshal(b, &rsp)
	require.NoError(t, err)
	assert.True(t, rsp.Status.Bool())
}

func TestTransactions_PartialDebit_Builder(t *testing.T) {
	r := NewPartialDebitRequestBuilder().AuthorizationCode("AUTH").Currency("NGN").Amount(1000).Email("e@example.com").Reference("ref").AtLeast("500").Build()
	assert.Equal(t, "AUTH", r.AuthorizationCode)
}
