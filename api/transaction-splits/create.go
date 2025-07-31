package transaction_splits

import (
	"context"

	"github.com/huysamen/paystack-go/net"
)

// Create creates a new transaction split
func (c *Client) Create(ctx context.Context, req *TransactionSplitCreateRequest) (*TransactionSplitCreateResponse, error) {
	if err := validateCreateRequest(req); err != nil {
		return nil, err
	}

	resp, err := net.Post[TransactionSplitCreateRequest, TransactionSplitCreateResponse](
		ctx, c.client, c.secret, transactionSplitBasePath, req, c.baseURL,
	)
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}
