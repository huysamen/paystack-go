package subscriptions

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSubscriptions_List_JSONDeserialization(t *testing.T) {
	p := filepath.Join("..", "..", "resources", "examples", "responses", "subscriptions", "list_200.json")
	b, err := os.ReadFile(p)
	require.NoError(t, err)

	var rsp ListResponse
	err = json.Unmarshal(b, &rsp)
	require.NoError(t, err)

	assert.True(t, rsp.Status.Bool())
	assert.NotNil(t, rsp.Data)
}

func TestSubscriptions_List_BuilderQuery(t *testing.T) {
	list := NewListRequestBuilder().PerPage(20).Page(1).Customer(1).Plan(2).Build()
	q := list.toQuery()
	assert.Contains(t, q, "perPage=20")
	assert.Contains(t, q, "page=1")
}
