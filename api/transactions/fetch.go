package transactions

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type FetchResponseData = types.Transaction
type FetchResponse = types.Response[FetchResponseData]

func (c *Client) Fetch(ctx context.Context, id uint64) (*FetchResponse, error) {
	return net.Get[FetchResponseData](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%d", basePath, id), "", c.BaseURL)
}
