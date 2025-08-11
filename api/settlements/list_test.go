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

func TestSettlements_List_JSONDeserialization(t *testing.T) {
	p := filepath.Join("..", "..", "resources", "examples", "responses", "settlements", "list_200.json")
	b, err := os.ReadFile(p)
	require.NoError(t, err)

	var rsp ListResponse
	err = json.Unmarshal(b, &rsp)
	require.NoError(t, err)

	assert.True(t, rsp.Status.Bool())
	assert.Equal(t, "Settlements retrieved", rsp.Message)
	assert.NotNil(t, rsp.Data)
	assert.GreaterOrEqual(t, len(rsp.Data), 1)
}

func TestSettlements_ListRequestBuilder_Build(t *testing.T) {
	from := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	to := time.Date(2023, 12, 31, 23, 59, 59, 0, time.UTC)

	builder := NewListRequestBuilder().PerPage(25).Page(2).Status("success").Subaccount("SUB_123").DateRange(from, to)
	req := builder.Build()

	require.NotNil(t, req)
	assert.NotNil(t, req.PerPage)
	assert.NotNil(t, req.Page)
	assert.NotNil(t, req.Status)
	assert.NotNil(t, req.Subaccount)
	assert.NotNil(t, req.From)
	assert.NotNil(t, req.To)
}
