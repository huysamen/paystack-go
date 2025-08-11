package transactionsplits

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/huysamen/paystack-go/enums"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTransactionSplits_Create_JSONDeserialization(t *testing.T) {
	p := filepath.Join("..", "..", "resources", "examples", "responses", "transactionsplits", "create_200.json")
	b, err := os.ReadFile(p)
	require.NoError(t, err)

	var rsp CreateResponse
	err = json.Unmarshal(b, &rsp)
	require.NoError(t, err)
	assert.True(t, rsp.Status.Bool())
}

func TestTransactionSplits_Create_Builder(t *testing.T) {
	b := NewCreateRequest("name", enums.TransactionSplitTypePercentage, enums.CurrencyNGN).
		AddSubaccount("SUB_1", 50).
		BearerType(enums.TransactionSplitBearerTypeAccount).
		BearerSubaccount("SUB_2").
		Build()
	assert.Equal(t, "name", b.Name)
}
