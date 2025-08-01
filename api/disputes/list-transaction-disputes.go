package disputes

import (
"context"
"errors"

"github.com/huysamen/paystack-go/net"
"github.com/huysamen/paystack-go/types"
)

// ListTransactionDisputesResponse represents the response from listing transaction disputes
type ListTransactionDisputesResponse = types.Response[TransactionDisputeData]

// ListTransactionDisputes retrieves disputes for a transaction
func (c *Client) ListTransactionDisputes(ctx context.Context, transactionID string) (*types.Response[TransactionDisputeData], error) {
	if transactionID == "" {
		return nil, errors.New("transaction ID is required")
	}

	endpoint := c.baseURL + "/transaction/" + transactionID + "/disputes"
	resp, err := net.Get[TransactionDisputeData](ctx, c.client, c.secret, endpoint, c.baseURL)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
