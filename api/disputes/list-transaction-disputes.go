package disputes

import (
	"context"
	"errors"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type ListTransactionDisputesResponse = types.Response[TransactionDisputeData]

func (c *Client) ListTransactionDisputes(ctx context.Context, transactionID string) (*ListTransactionDisputesResponse, error) {
	if transactionID == "" {
		return nil, errors.New("transaction ID is required")
	}

	return net.Get[TransactionDisputeData](ctx, c.Client, c.Secret, "/transaction/"+transactionID+"/disputes", c.BaseURL)
}
