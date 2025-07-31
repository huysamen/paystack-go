package transfercontrol

import (
	"context"

	"github.com/huysamen/paystack-go/net"
)

// CheckBalance fetches the available balance on your integration
func (c *Client) CheckBalance(ctx context.Context) (*CheckBalanceResponse, error) {
	resp, err := net.Get[CheckBalanceResponse](ctx, c.client, c.secret, "/balance", c.baseURL)
	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}
