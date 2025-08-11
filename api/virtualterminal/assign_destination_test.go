package virtualterminal

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/huysamen/paystack-go/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestVirtualTerminal_AssignDestination_JSONDeserialization(t *testing.T) {
	p := filepath.Join("..", "..", "resources", "examples", "responses", "virtualterminal", "assign_destination_200.json")
	b, err := os.ReadFile(p)
	require.NoError(t, err)

	var rsp AssignDestinationResponse
	err = json.Unmarshal(b, &rsp)
	require.NoError(t, err)
}

func TestVirtualTerminal_AssignDestination_Builder(t *testing.T) {
	assign := NewAssignDestinationRequestBuilder().
		AddDestination(types.VirtualTerminalDestination{Type: "bank_account", Target: "0123"}).
		Build()
	assert.NotNil(t, assign.Destinations)
}
