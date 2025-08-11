package transactionsplits

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTransactionSplits_List_JSONDeserialization(t *testing.T) {
	p := filepath.Join("..", "..", "resources", "examples", "responses", "transactionsplits", "list_200.json")
	b, err := os.ReadFile(p)
	require.NoError(t, err)

	var rsp ListResponse
	err = json.Unmarshal(b, &rsp)
	require.NoError(t, err)
	assert.True(t, rsp.Status.Bool())
}

func TestTransactionSplits_List_Builder(t *testing.T) {
	lb := NewListRequestBuilder().Name("n").Active(true).SortBy("createdAt").PerPage(10).Page(1).Build()
	q := lb.toQuery()
	assert.Contains(t, q, "perPage=10")
}
