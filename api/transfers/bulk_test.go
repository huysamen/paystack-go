package transfers

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTransfers_Bulk_JSONDeserialization(t *testing.T) {
	p := filepath.Join("..", "..", "resources", "examples", "responses", "transfers", "bulk_initiate_response.json")
	b, err := os.ReadFile(p)
	require.NoError(t, err)

	var rsp BulkResponse
	err = json.Unmarshal(b, &rsp)
	require.NoError(t, err)
}

func TestTransfers_Bulk_Builder(t *testing.T) {
	r := NewBulkRequestBuilder("balance").AddTransfer(BulkTransferItem{Amount: 10, Reference: "r", Reason: "x", Recipient: "R"}).Build()
	assert.Equal(t, "balance", r.Source)
	if len(r.Transfers) > 0 {
		assert.Equal(t, 10, r.Transfers[0].Amount)
	}
}
