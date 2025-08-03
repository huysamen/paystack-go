package transactions

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// Response type alias
type TransactionFetchResponse = types.Response[types.Transaction]

func (c *Client) Fetch(ctx context.Context, id uint64) (*TransactionFetchResponse, error) {
	return net.Get[types.Transaction](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%d", basePath, id), "", c.BaseURL)
}
