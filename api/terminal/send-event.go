package terminal

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
)

// SendEvent sends an event to a terminal
func (c *Client) SendEvent(ctx context.Context, terminalID string, req *TerminalSendEventRequest) (*TerminalSendEventResponse, error) {
	if err := validateTerminalID(terminalID); err != nil {
		return nil, err
	}
	if err := validateSendEventRequest(req); err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("%s/%s/event", terminalBasePath, terminalID)
	resp, err := net.Post[TerminalSendEventRequest, TerminalSendEventResponse](
		ctx, c.client, c.secret, endpoint, req, c.baseURL,
	)
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}
