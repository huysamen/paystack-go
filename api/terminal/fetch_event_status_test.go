package terminal

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTerminal_FetchEventStatus_JSONDeserialization(t *testing.T) {
	p := filepath.Join("..", "..", "resources", "examples", "responses", "terminal", "fetch_event_status_200.json")
	b, err := os.ReadFile(p)
	require.NoError(t, err)

	var rsp FetchEventStatusResponse
	err = json.Unmarshal(b, &rsp)
	require.NoError(t, err)
}
