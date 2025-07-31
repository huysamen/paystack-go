package virtualterminal

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
)

// Fetch retrieves a specific virtual terminal by code
func (c *Client) Fetch(ctx context.Context, code string) (*VirtualTerminal, error) {
	if err := validateCode(code); err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("%s/%s", virtualTerminalBasePath, code)
	resp, err := net.Get[VirtualTerminal](
		ctx, c.client, c.secret, endpoint, c.baseURL,
	)
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}
