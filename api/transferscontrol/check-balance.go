package transfercontrol

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type CheckBalanceResponseData = []types.Balance
type CheckBalanceResponse = types.Response[CheckBalanceResponseData]

func (c *Client) CheckBalance(ctx context.Context) (*CheckBalanceResponse, error) {
	return net.Get[CheckBalanceResponseData](ctx, c.Client, c.Secret, basePath, "", c.BaseURL)
}
