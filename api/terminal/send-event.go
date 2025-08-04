package terminal

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// TerminalSendEventResponse represents the response from sending an event
type TerminalSendEventResponse = types.Response[types.TerminalEventResult]

// SendEvent sends an event to a terminal
func (c *Client) SendEvent(ctx context.Context, terminalID string, req *types.TerminalSendEventRequest) (*TerminalSendEventResponse, error) {
	endpoint := fmt.Sprintf("%s/%s/event", basePath, terminalID)
	return net.Post[types.TerminalSendEventRequest, types.TerminalEventResult](
		ctx, c.Client, c.Secret, endpoint, req, c.BaseURL,
	)
}
