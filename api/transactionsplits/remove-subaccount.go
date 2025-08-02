package transactionsplits

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// RemoveSubaccount removes a subaccount from a transaction split
func (c *Client) RemoveSubaccount(ctx context.Context, id string, builder *TransactionSplitSubaccountRemoveRequestBuilder) (*types.Response[any], error) {
	req := builder.Build()
	return net.Post[TransactionSplitSubaccountRemoveRequest, any](
		ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s/subaccount/remove", basePath, id), req, c.BaseURL,
	)
}
