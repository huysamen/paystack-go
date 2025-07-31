package transaction_splits

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
)

// Fetch retrieves a transaction split by ID
func (c *Client) Fetch(ctx context.Context, id string) (*TransactionSplitFetchResponse, error) {
	if err := validateTransactionSplitID(id); err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("%s/%s", transactionSplitBasePath, id)
	resp, err := net.Get[TransactionSplitFetchResponse](ctx, c.client, c.secret, endpoint, c.baseURL)
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}
