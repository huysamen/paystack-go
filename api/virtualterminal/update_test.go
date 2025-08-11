package virtualterminal

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestVirtualTerminal_Update_JSONDeserialization(t *testing.T) {
	p := filepath.Join("..", "..", "resources", "examples", "responses", "virtualterminal", "update_400.json")
	b, err := os.ReadFile(p)
	require.NoError(t, err)

	var rsp UpdateResponse
	err = json.Unmarshal(b, &rsp)
	require.NoError(t, err)
}

func TestVirtualTerminal_Update_Builder(t *testing.T) {
	upd := NewUpdateRequestBuilder("n").Build()
	assert.Equal(t, "n", upd.Name)
}
