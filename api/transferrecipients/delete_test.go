package transferrecipients

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferRecipients_Delete_JSONDeserialization(t *testing.T) {
	// Reuse fetch_response shape which matches status/message/data envelope
	p := filepath.Join("..", "..", "resources", "examples", "responses", "transferrecipients", "fetch_response.json")
	b, err := os.ReadFile(p)
	require.NoError(t, err)

	var rsp DeleteResponse
	err = json.Unmarshal(b, &rsp)
	require.NoError(t, err)
}
