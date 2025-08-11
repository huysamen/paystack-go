package transfers

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTransfers_Initiate_JSONDeserialization(t *testing.T) {
	p := filepath.Join("..", "..", "resources", "examples", "responses", "transfers", "initiate_response.json")
	b, err := os.ReadFile(p)
	require.NoError(t, err)

	var rsp InitiateResponse
	err = json.Unmarshal(b, &rsp)
	require.NoError(t, err)
}

func TestTransfers_Initiate_Builder(t *testing.T) {
	r := NewInitiateRequestBuilder("balance", 1000, "RCP_123").Reason("test").Currency("NGN").AccountReference("AR").Reference("REF").Build()
	assert.Equal(t, "balance", r.Source)
	assert.Equal(t, 1000, r.Amount)
	assert.Equal(t, "RCP_123", r.Recipient)
}
