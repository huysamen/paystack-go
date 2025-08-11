package transferrecipients

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTransferRecipients_List_JSONDeserialization(t *testing.T) {
	p := filepath.Join("..", "..", "resources", "examples", "responses", "transferrecipients", "list_response.json")
	b, err := os.ReadFile(p)
	require.NoError(t, err)

	var rsp ListResponse
	err = json.Unmarshal(b, &rsp)
	require.NoError(t, err)
	assert.NotNil(t, rsp.Data)
}

func TestTransferRecipients_List_BuilderQuery(t *testing.T) {
	from := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	to := time.Date(2023, 1, 31, 0, 0, 0, 0, time.UTC)
	req := NewListRequestBuilder().PerPage(10).Page(1).From(from).To(to).Build()
	q := req.toQuery()
	assert.Contains(t, q, "perPage=10")
	assert.Contains(t, q, "page=1")
}
