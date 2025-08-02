package transactionsplits

import (
	"context"
	"errors"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// RemoveSubaccount removes a subaccount from a transaction split
func (c *Client) RemoveSubaccount(ctx context.Context, id string, builder *TransactionSplitSubaccountRemoveRequestBuilder) (*types.Response[any], error) {
	if id == "" {
		return nil, errors.New("transaction split ID is required")
	}
	if builder == nil {
		return nil, ErrBuilderRequired
	}

	req := builder.Build()
	endpoint := fmt.Sprintf("%s/%s/subaccount/remove", transactionSplitBasePath, id)
	return net.Post[TransactionSplitSubaccountRemoveRequest, any](
		ctx, c.client, c.secret, endpoint, req, c.baseURL,
	)
}
