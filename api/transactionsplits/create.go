package transactionsplits

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// Create creates a new transaction split
func (c *Client) Create(ctx context.Context, builder *TransactionSplitCreateRequestBuilder) (*types.Response[TransactionSplit], error) {
	req := builder.Build()
	return net.Post[TransactionSplitCreateRequest, TransactionSplit](
		ctx, c.Client, c.Secret, basePath, req, c.BaseURL,
	)
}
