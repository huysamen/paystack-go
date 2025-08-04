package transfercontrol

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// CheckBalanceResponse represents the response from checking balance
type CheckBalanceResponse = types.Response[[]types.Balance]

// CheckBalance fetches the available balance on your integration
func (c *Client) CheckBalance(ctx context.Context) (*CheckBalanceResponse, error) {
	return net.Get[[]types.Balance](ctx, c.Client, c.Secret, basePath, "", c.BaseURL)
}
