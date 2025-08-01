package refunds

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
)

// Fetch retrieves details of a specific refund
func (c *Client) Fetch(ctx context.Context, refundID string) (*RefundFetchResponse, error) {
	if refundID == "" {
		return nil, fmt.Errorf("refund ID is required")
	}

	url := c.baseURL + refundsBasePath + "/" + refundID
	return net.Get[Refund](ctx, c.client, c.secret, url)
}
