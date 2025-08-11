package transferrecipients

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/huysamen/paystack-go/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTransferRecipients_BulkCreate_JSONDeserialization(t *testing.T) {
	p := filepath.Join("..", "..", "resources", "examples", "responses", "transferrecipients", "bulk_create_200.json")
	b, err := os.ReadFile(p)
	require.NoError(t, err)

	var rsp BulkCreateResponse
	err = json.Unmarshal(b, &rsp)
	require.NoError(t, err)
}

func TestTransferRecipients_BulkCreate_Builder(t *testing.T) {
	item := types.BulkRecipientItem{}
	r := NewBulkCreateRequestBuilder().AddRecipient(item).Build()
	assert.Len(t, r.Batch, 1)
}
