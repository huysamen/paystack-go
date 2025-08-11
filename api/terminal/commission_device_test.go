package terminal

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTerminal_CommissionDevice_JSONDeserialization(t *testing.T) {
	p := filepath.Join("..", "..", "resources", "examples", "responses", "terminal", "commission_200.json")
	b, err := os.ReadFile(p)
	require.NoError(t, err)

	var rsp CommissionDeviceResponse
	err = json.Unmarshal(b, &rsp)
	require.NoError(t, err)
}

func TestTerminal_CommissionDevice_Builder(t *testing.T) {
	commission := NewCommissionDeviceRequestBuilder("SN").Build()
	assert.Equal(t, "SN", commission.SerialNumber)
}
