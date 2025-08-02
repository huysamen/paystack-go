package transaction_splits

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// Create creates a new transaction split
func (c *Client) Create(ctx context.Context, builder *TransactionSplitCreateRequestBuilder) (*types.Response[TransactionSplit], error) {
	if builder == nil {
		return nil, ErrBuilderRequired
	}

	req := builder.Build()
	return net.Post[TransactionSplitCreateRequest, TransactionSplit](
		ctx, c.client, c.secret, transactionSplitBasePath, req, c.baseURL,
	)
}
