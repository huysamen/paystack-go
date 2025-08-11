package transfers

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTransfers_Finalize_JSONDeserialization(t *testing.T) {
	p := filepath.Join("..", "..", "resources", "examples", "responses", "transfers", "finalize_response.json")
	b, err := os.ReadFile(p)
	require.NoError(t, err)

	var rsp FinalizeResponse
	err = json.Unmarshal(b, &rsp)
	require.NoError(t, err)
}

func TestTransfers_Finalize_Builder(t *testing.T) {
	r := NewFinalizeRequestBuilder("TRF_123", "123456").Build()
	assert.Equal(t, "TRF_123", r.TransferCode)
	assert.Equal(t, "123456", r.OTP)
}
