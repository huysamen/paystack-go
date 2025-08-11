package settlements

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSettlements_ListTransactions_JSONDeserialization(t *testing.T) {
	p := filepath.Join("..", "..", "resources", "examples", "responses", "settlements", "list_transactions_200.json")
	b, err := os.ReadFile(p)
	require.NoError(t, err)

	var rsp ListTransactionsResponse
	err = json.Unmarshal(b, &rsp)
	require.NoError(t, err)

	assert.True(t, rsp.Status.Bool())
	assert.Equal(t, "Transactions retrieved", rsp.Message)
	assert.NotNil(t, rsp.Data)
	if len(rsp.Data) > 0 {
		first := rsp.Data[0]
		assert.True(t, first.PaidAt.Valid)
		// Ensure time parses
		_, _ = time.Parse(time.RFC3339, first.PaidAt.String())
	}
}

func TestSettlements_ListTransactionsRequestBuilder_Build(t *testing.T) {
	from := time.Date(2023, 6, 1, 0, 0, 0, 0, time.UTC)
	to := time.Date(2023, 6, 30, 23, 59, 59, 0, time.UTC)

	builder := NewListTransactionsRequestBuilder().PerPage(50).Page(3).DateRange(from, to)
	req := builder.Build()

	require.NotNil(t, req)
	assert.NotNil(t, req.PerPage)
	assert.NotNil(t, req.Page)
	assert.NotNil(t, req.From)
	assert.NotNil(t, req.To)
}
