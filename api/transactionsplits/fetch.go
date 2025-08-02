package transactionsplits

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// Fetch retrieves a transaction split by ID
func (c *Client) Fetch(ctx context.Context, id string) (*types.Response[TransactionSplit], error) {
	return net.Get[TransactionSplit](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s", basePath, id), "", c.BaseURL)
}
