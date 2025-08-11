package virtualterminal

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestVirtualTerminal_UnassignDestination_JSONDeserialization(t *testing.T) {
	p := filepath.Join("..", "..", "resources", "examples", "responses", "virtualterminal", "unassign_destination_200.json")
	b, err := os.ReadFile(p)
	require.NoError(t, err)

	var rsp UnassignDestinationResponse
	err = json.Unmarshal(b, &rsp)
	require.NoError(t, err)
}

func TestVirtualTerminal_UnassignDestination_Builder(t *testing.T) {
	unassign := NewUnassignDestinationRequestBuilder().
		AddTarget("bank_account:0123").
		Build()
	assert.NotNil(t, unassign.Targets)
}
