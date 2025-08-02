package transactionsplits

import (
	"context"
	"errors"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// Update updates a transaction split
func (c *Client) Update(ctx context.Context, id string, builder *TransactionSplitUpdateRequestBuilder) (*types.Response[TransactionSplit], error) {
	if id == "" {
		return nil, errors.New("transaction split ID is required")
	}
	if builder == nil {
		return nil, ErrBuilderRequired
	}

	req := builder.Build()
	endpoint := fmt.Sprintf("%s/%s", transactionSplitBasePath, id)
	return net.Put[TransactionSplitUpdateRequest, TransactionSplit](
		ctx, c.client, c.secret, endpoint, req, c.baseURL,
	)
}
