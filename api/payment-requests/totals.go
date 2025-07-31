package paymentrequests

import (
	"context"

	"github.com/huysamen/paystack-go/net"
)

// GetTotals gets payment requests metric
func (c *Client) GetTotals(ctx context.Context) (*PaymentRequestTotals, error) {
	resp, err := net.Get[PaymentRequestTotals](
		ctx, c.client, c.secret, paymentRequestsBasePath+"/totals", c.baseURL,
	)
	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}
