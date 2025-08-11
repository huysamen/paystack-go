package terminal

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTerminal_List_JSONDeserialization(t *testing.T) {
	p := filepath.Join("..", "..", "resources", "examples", "responses", "terminal", "list_200.json")
	b, err := os.ReadFile(p)
	require.NoError(t, err)

	var rsp ListResponse
	err = json.Unmarshal(b, &rsp)
	require.NoError(t, err)

	assert.True(t, rsp.Status.Bool())
	assert.NotNil(t, rsp.Data)
}

func TestTerminal_List_BuilderQuery(t *testing.T) {
	list := NewListRequestBuilder().PerPage(10).Next("n").Previous("p").Build()
	q := list.toQuery()
	assert.Contains(t, q, "perPage=10")
	assert.Contains(t, q, "next=n")
	assert.Contains(t, q, "previous=p")
}
