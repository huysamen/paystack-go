package terminal

import (
	"context"
	"errors"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// SendEvent sends an event to a terminal
func (c *Client) SendEvent(ctx context.Context, terminalID string, req *TerminalSendEventRequest) (*types.Response[any], error) {
	if terminalID == "" {
		return nil, errors.New("terminal ID is required")
	}
	if req == nil {
		return nil, errors.New("request cannot be nil")
	}

	endpoint := fmt.Sprintf("%s/%s/event", terminalBasePath, terminalID)
	return net.Post[TerminalSendEventRequest, any](
		ctx, c.client, c.secret, endpoint, req, c.baseURL,
	)
}
