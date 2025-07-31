package terminal

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
)

// FetchTerminalStatus checks the availability of a terminal
func (c *Client) FetchTerminalStatus(ctx context.Context, terminalID string) (*TerminalPresenceResponse, error) {
	if err := validateTerminalID(terminalID); err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("%s/%s/presence", terminalBasePath, terminalID)
	resp, err := net.Get[TerminalPresenceResponse](ctx, c.client, c.secret, endpoint, c.baseURL)
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}
