package transfercontrol

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type FetchBalanceLedgerResponse = types.Response[[]types.BalanceLedger]

func (c *Client) FetchBalanceLedger(ctx context.Context) (*FetchBalanceLedgerResponse, error) {
	return net.Get[[]types.BalanceLedger](ctx, c.Client, c.Secret, "/balance/ledger", "", c.BaseURL)
}
