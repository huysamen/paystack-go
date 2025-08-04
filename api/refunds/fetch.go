package refunds

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type RefundFetchResponse = types.Response[types.Refund]

func (c *Client) Fetch(ctx context.Context, refundID string) (*RefundFetchResponse, error) {
	return net.Get[types.Refund](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s", basePath, refundID), c.BaseURL)
}
