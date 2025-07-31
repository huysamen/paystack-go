package transfercontrol

import (
	"context"

	"github.com/huysamen/paystack-go/net"
)

// FetchBalanceLedger fetches all pay-ins and pay-outs that occurred on your integration
func (c *Client) FetchBalanceLedger(ctx context.Context) (*FetchBalanceLedgerResponse, error) {
	resp, err := net.Get[FetchBalanceLedgerResponse](ctx, c.client, c.secret, "/balance/ledger", c.baseURL)
	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}
