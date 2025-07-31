package terminal

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
)

// Fetch retrieves a terminal by ID
func (c *Client) Fetch(ctx context.Context, terminalID string) (*TerminalFetchResponse, error) {
	if err := validateTerminalID(terminalID); err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("%s/%s", terminalBasePath, terminalID)
	resp, err := net.Get[TerminalFetchResponse](ctx, c.client, c.secret, endpoint, c.baseURL)
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}
