package transfercontrol

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// FetchBalanceLedgerResponse represents the response from fetching balance ledger
type FetchBalanceLedgerResponse = types.Response[[]types.BalanceLedger]

// FetchBalanceLedger fetches all pay-ins and pay-outs that occurred on your integration
func (c *Client) FetchBalanceLedger(ctx context.Context) (*FetchBalanceLedgerResponse, error) {
	return net.Get[[]types.BalanceLedger](ctx, c.Client, c.Secret, "/balance/ledger", "", c.BaseURL)
}
