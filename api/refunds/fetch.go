package refunds

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type FetchResponseData = types.Refund
type FetchResponse = types.Response[FetchResponseData]

func (c *Client) Fetch(ctx context.Context, refundID string) (*FetchResponse, error) {
	return net.Get[FetchResponseData](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s", basePath, refundID), c.BaseURL)
}
