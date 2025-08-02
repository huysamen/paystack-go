package transaction_splits

import (
	"context"
	"errors"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// Fetch retrieves a transaction split by ID
func (c *Client) Fetch(ctx context.Context, id string) (*types.Response[TransactionSplit], error) {
	if id == "" {
		return nil, errors.New("transaction split ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", transactionSplitBasePath, id)
	return net.Get[TransactionSplit](ctx, c.client, c.secret, endpoint, c.baseURL)
}
