package transactionsplits

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// Update updates a transaction split
func (c *Client) Update(ctx context.Context, id string, builder *TransactionSplitUpdateRequestBuilder) (*types.Response[TransactionSplit], error) {
	req := builder.Build()
	return net.Put[TransactionSplitUpdateRequest, TransactionSplit](
		ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s", basePath, id), req, c.BaseURL,
	)
}
