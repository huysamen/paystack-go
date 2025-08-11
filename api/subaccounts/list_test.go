package subaccounts

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSubaccounts_List_JSONDeserialization(t *testing.T) {
	p := filepath.Join("..", "..", "resources", "examples", "responses", "subaccounts", "list_200.json")
	b, err := os.ReadFile(p)
	require.NoError(t, err)

	var rsp ListResponse
	err = json.Unmarshal(b, &rsp)
	require.NoError(t, err)

	assert.True(t, rsp.Status.Bool())
	assert.NotNil(t, rsp.Data)
	assert.GreaterOrEqual(t, len(rsp.Data), 1)
}

func TestSubaccounts_List_BuilderQuery(t *testing.T) {
	from := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	to := time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC)
	req := NewListRequestBuilder().PerPage(10).Page(1).From(from).To(to).Build()
	q := req.toQuery()
	assert.Contains(t, q, "perPage=10")
	assert.Contains(t, q, "page=1")
}

// Create/update builder tests live in corresponding *_test.go files
