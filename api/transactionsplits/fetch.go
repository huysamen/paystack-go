package transactionsplits

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type TransactionSplitFetchResponse = types.Response[types.TransactionSplit]

func (c *Client) Fetch(ctx context.Context, id string) (*TransactionSplitFetchResponse, error) {
	return net.Get[types.TransactionSplit](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s", basePath, id), "", c.BaseURL)
}
