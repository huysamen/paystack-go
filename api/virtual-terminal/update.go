package virtualterminal

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
)

// Update updates a virtual terminal
func (c *Client) Update(ctx context.Context, code string, req *UpdateVirtualTerminalRequest) (*VirtualTerminal, error) {
	if err := validateCode(code); err != nil {
		return nil, err
	}
	if err := validateUpdateRequest(req); err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("%s/%s", virtualTerminalBasePath, code)
	resp, err := net.Put[UpdateVirtualTerminalRequest, VirtualTerminal](
		ctx, c.client, c.secret, endpoint, req, c.baseURL,
	)
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}
