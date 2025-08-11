package terminal

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTerminal_Update_JSONDeserialization(t *testing.T) {
	p := filepath.Join("..", "..", "resources", "examples", "responses", "terminal", "update_200.json")
	b, err := os.ReadFile(p)
	require.NoError(t, err)

	var rsp UpdateResponse
	err = json.Unmarshal(b, &rsp)
	require.NoError(t, err)
}

func TestTerminal_Update_Builder(t *testing.T) {
	upd := NewUpdateRequestBuilder().Name("n").Address("a").Build()
	assert.NotNil(t, upd.Name)
	assert.NotNil(t, upd.Address)
}
