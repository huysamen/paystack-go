package transactionsplits

import (
	"context"
	"errors"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// AddSubaccount adds or updates a subaccount in a transaction split
func (c *Client) AddSubaccount(ctx context.Context, id string, builder *TransactionSplitSubaccountAddRequestBuilder) (*types.Response[TransactionSplit], error) {
	if id == "" {
		return nil, errors.New("transaction split ID is required")
	}
	if builder == nil {
		return nil, ErrBuilderRequired
	}

	req := builder.Build()
	endpoint := fmt.Sprintf("%s/%s/subaccount/add", transactionSplitBasePath, id)
	return net.Post[TransactionSplitSubaccountAddRequest, TransactionSplit](
		ctx, c.client, c.secret, endpoint, req, c.baseURL,
	)
}
