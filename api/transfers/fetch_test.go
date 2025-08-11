package transfers

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransfers_Fetch_JSONDeserialization(t *testing.T) {
	p := filepath.Join("..", "..", "resources", "examples", "responses", "transfers", "fetch_response.json")
	b, err := os.ReadFile(p)
	require.NoError(t, err)

	var rsp FetchResponse
	err = json.Unmarshal(b, &rsp)
	require.NoError(t, err)
}
