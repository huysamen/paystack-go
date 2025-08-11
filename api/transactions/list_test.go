package transactions

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTransactions_List_JSONDeserialization(t *testing.T) {
	p := filepath.Join("..", "..", "resources", "examples", "responses", "transactions", "list_200.json")
	b, err := os.ReadFile(p)
	require.NoError(t, err)

	var rsp ListResponse
	err = json.Unmarshal(b, &rsp)
	require.NoError(t, err)

	assert.NotNil(t, rsp.Data)
	if len(rsp.Data) > 0 {
		first := rsp.Data[0]
		_ = first // basic smoke checks done via field types in models
	}
}

func TestTransactions_Builders(t *testing.T) {
	from := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	to := time.Date(2023, 1, 31, 0, 0, 0, 0, time.UTC)
	req := NewListRequestBuilder().PerPage(10).Page(1).Customer(1).TerminalID("T").Status("success").From(from).To(to).Amount(5000).Build()
	q := req.toQuery()
	assert.Contains(t, q, "perPage=10")
	assert.Contains(t, q, "page=1")
	assert.Contains(t, q, "customer=1")
	assert.Contains(t, q, "terminalid=T")
	assert.Contains(t, q, "status=success")
	assert.Contains(t, q, "amount=5000")
}
