package transfercontrol

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// FetchBalanceLedger fetches all pay-ins and pay-outs that occurred on your integration
func (c *Client) FetchBalanceLedger(ctx context.Context) (*types.Response[[]BalanceLedger], error) {
	return net.Get[[]BalanceLedger](ctx, c.client, c.secret, "/balance/ledger", "", c.baseURL)
}
