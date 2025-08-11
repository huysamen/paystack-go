package terminal

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTerminal_DecommissionDevice_JSONDeserialization(t *testing.T) {
	p := filepath.Join("..", "..", "resources", "examples", "responses", "terminal", "decommission_200.json")
	b, err := os.ReadFile(p)
	require.NoError(t, err)

	var rsp DecommissionDeviceResponse
	err = json.Unmarshal(b, &rsp)
	require.NoError(t, err)
}

func TestTerminal_DecommissionDevice_Builder(t *testing.T) {
	r := NewDecommissionDeviceRequest("SN").Build()
	assert.Equal(t, "SN", r.SerialNumber)
}
