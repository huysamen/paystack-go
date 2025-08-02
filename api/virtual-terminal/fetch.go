package virtualterminal

import (
	"context"
	"errors"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// Fetch retrieves a specific virtual terminal by code
func (c *Client) Fetch(ctx context.Context, code string) (*types.Response[VirtualTerminal], error) {
	if code == "" {
		return nil, errors.New("virtual terminal code is required")
	}

	endpoint := fmt.Sprintf("%s/%s", virtualTerminalBasePath, code)
	return net.Get[VirtualTerminal](
		ctx, c.client, c.secret, endpoint, c.baseURL,
	)
}
