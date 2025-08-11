package virtualterminal

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestVirtualTerminal_Create_JSONDeserialization(t *testing.T) {
	p := filepath.Join("..", "..", "resources", "examples", "responses", "virtualterminal", "create_200.json")
	b, err := os.ReadFile(p)
	require.NoError(t, err)

	var rsp CreateResponse
	err = json.Unmarshal(b, &rsp)
	require.NoError(t, err)
}
