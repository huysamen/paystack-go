package transfercontrol

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type FetchBalanceLedgerResponseData = []types.BalanceLedger
type FetchBalanceLedgerResponse = types.Response[FetchBalanceLedgerResponseData]

func (c *Client) FetchBalanceLedger(ctx context.Context) (*FetchBalanceLedgerResponse, error) {
	return net.Get[FetchBalanceLedgerResponseData](ctx, c.Client, c.Secret, "/balance/ledger", "", c.BaseURL)
}
