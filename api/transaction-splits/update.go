package transaction_splits

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
)

// Update updates a transaction split
func (c *Client) Update(ctx context.Context, id string, req *TransactionSplitUpdateRequest) (*TransactionSplitUpdateResponse, error) {
	if err := validateTransactionSplitID(id); err != nil {
		return nil, err
	}
	if err := validateUpdateRequest(req); err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("%s/%s", transactionSplitBasePath, id)
	resp, err := net.Put[TransactionSplitUpdateRequest, TransactionSplitUpdateResponse](
		ctx, c.client, c.secret, endpoint, req, c.baseURL,
	)
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}
