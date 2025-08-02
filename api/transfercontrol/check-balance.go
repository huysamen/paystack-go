package transfercontrol

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// CheckBalance fetches the available balance on your integration
func (c *Client) CheckBalance(ctx context.Context) (*types.Response[[]Balance], error) {
	return net.Get[[]Balance](ctx, c.Client, c.Secret, basePath, "", c.BaseURL)
}
