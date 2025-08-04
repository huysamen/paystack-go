package transfercontrol

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type CheckBalanceResponse = types.Response[[]types.Balance]

func (c *Client) CheckBalance(ctx context.Context) (*CheckBalanceResponse, error) {
	return net.Get[[]types.Balance](ctx, c.Client, c.Secret, basePath, "", c.BaseURL)
}
