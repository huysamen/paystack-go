package virtualterminal

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestVirtualTerminal_Deactivate_JSONDeserialization(t *testing.T) {
	p := filepath.Join("..", "..", "resources", "examples", "responses", "virtualterminal", "deactivate_200.json")
	b, err := os.ReadFile(p)
	require.NoError(t, err)

	var rsp DeactivateResponse
	err = json.Unmarshal(b, &rsp)
	require.NoError(t, err)
}
