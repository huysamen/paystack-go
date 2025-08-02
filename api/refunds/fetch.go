package refunds

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
)

// Fetch retrieves details of a specific refund
func (c *Client) Fetch(ctx context.Context, refundID string) (*RefundFetchResponse, error) {
	return net.Get[Refund](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s", basePath, refundID), c.BaseURL)
}
