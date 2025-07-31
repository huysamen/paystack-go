package virtualterminal

import (
	"context"

	"github.com/huysamen/paystack-go/net"
)

// Create creates a new virtual terminal
func (c *Client) Create(ctx context.Context, req *CreateVirtualTerminalRequest) (*VirtualTerminal, error) {
	if err := validateCreateRequest(req); err != nil {
		return nil, err
	}

	resp, err := net.Post[CreateVirtualTerminalRequest, VirtualTerminal](
		ctx, c.client, c.secret, virtualTerminalBasePath, req, c.baseURL,
	)
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}
