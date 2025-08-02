package virtualterminal

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// Fetch retrieves a specific virtual terminal by code
func (c *Client) Fetch(ctx context.Context, code string) (*types.Response[VirtualTerminal], error) {
	return net.Get[VirtualTerminal](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s", basePath, code), c.BaseURL)
}
