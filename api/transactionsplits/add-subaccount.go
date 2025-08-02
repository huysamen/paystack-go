package transactionsplits

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// AddSubaccount adds or updates a subaccount in a transaction split
func (c *Client) AddSubaccount(ctx context.Context, id string, builder *TransactionSplitSubaccountAddRequestBuilder) (*types.Response[TransactionSplit], error) {
	req := builder.Build()
	return net.Post[TransactionSplitSubaccountAddRequest, TransactionSplit](
		ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s/subaccount/add", basePath, id), req, c.BaseURL,
	)
}
