package virtualterminal

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestVirtualTerminal_RemoveSplitCode_JSONDeserialization(t *testing.T) {
	// Only a 400 example is present; we still validate envelope parsing
	p := filepath.Join("..", "..", "resources", "examples", "responses", "virtualterminal", "remove_split_code_400.json")
	b, err := os.ReadFile(p)
	require.NoError(t, err)

	var rsp RemoveSplitCodeResponse
	err = json.Unmarshal(b, &rsp)
	require.NoError(t, err)
}
