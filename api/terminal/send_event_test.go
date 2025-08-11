package terminal

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/huysamen/paystack-go/enums"
	"github.com/huysamen/paystack-go/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTerminal_SendEvent_JSONDeserialization(t *testing.T) {
	p := filepath.Join("..", "..", "resources", "examples", "responses", "terminal", "send_event_200.json")
	b, err := os.ReadFile(p)
	require.NoError(t, err)

	var rsp SendEventResponse
	err = json.Unmarshal(b, &rsp)
	require.NoError(t, err)
}

func TestTerminal_SendEvent_Builder(t *testing.T) {
	payload := types.TerminalEventData{"type": "display_text", "message": "Hello"}
	r := NewSendEventRequestBuilder(enums.TerminalEventTypePayment, enums.TerminalEventActionProcess, payload).Build()
	assert.Equal(t, "display_text", r.Data["type"])
}
