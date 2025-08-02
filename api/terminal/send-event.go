package terminal

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// SendEvent sends an event to a terminal
func (c *Client) SendEvent(ctx context.Context, terminalID string, req *TerminalSendEventRequest) (*types.Response[any], error) {
	endpoint := fmt.Sprintf("%s/%s/event", basePath, terminalID)
	return net.Post[TerminalSendEventRequest, any](
		ctx, c.Client, c.Secret, endpoint, req, c.BaseURL,
	)
}
