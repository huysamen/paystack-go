package terminal

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
)

// SendEvent sends an event to a terminal
func (c *Client) SendEvent(ctx context.Context, terminalID string, req *TerminalSendEventRequest) (*TerminalSendEventResponse, error) {
	endpoint := fmt.Sprintf("%s/%s/event", basePath, terminalID)
	return net.Post[TerminalSendEventRequest, TerminalEventResult](
		ctx, c.Client, c.Secret, endpoint, req, c.BaseURL,
	)
}
