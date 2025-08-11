package transferscontrol

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTransfersControl_CheckBalance_JSONDeserialization(t *testing.T) {
	p := filepath.Join("..", "..", "resources", "examples", "responses", "transferscontrol", "check_balance_200.json")
	b, err := os.ReadFile(p)
	require.NoError(t, err)

	var rsp CheckBalanceResponse
	err = json.Unmarshal(b, &rsp)
	require.NoError(t, err)
	assert.True(t, rsp.Status.Bool())
}
