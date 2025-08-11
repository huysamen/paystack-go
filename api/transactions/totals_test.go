package transactions

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTransactions_Totals_JSONDeserialization(t *testing.T) {
	p := filepath.Join("..", "..", "resources", "examples", "responses", "transactions", "totals_200.json")
	b, err := os.ReadFile(p)
	require.NoError(t, err)

	var rsp TotalsResponse
	err = json.Unmarshal(b, &rsp)
	require.NoError(t, err)

	assert.True(t, rsp.Status.Bool())
	assert.NotNil(t, rsp.Data)
}
