package transactions

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTransactions_ChargeAuthorization_JSONDeserialization(t *testing.T) {
	p := filepath.Join("..", "..", "resources", "examples", "responses", "transactions", "charge_authorization_200.json")
	b, err := os.ReadFile(p)
	require.NoError(t, err)

	var rsp ChargeAuthorizationResponse
	err = json.Unmarshal(b, &rsp)
	require.NoError(t, err)
	assert.NotNil(t, rsp.Data)
}

func TestTransactions_ChargeAuthorization_Builder(t *testing.T) {
	r := NewChargeAuthorizationRequestBuilder().Amount(1000).Email("e@example.com").AuthorizationCode("AUTH").Reference("ref").Queue(true).Build()
	assert.Equal(t, 1000, r.Amount)
}
