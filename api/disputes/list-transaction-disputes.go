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

	return net.Get[TransactionDisputeData](ctx, c.Client, c.Secret, "/transaction/"+transactionID+"/disputes", c.BaseURL)
}
