package disputes

import (
	"context"
	"errors"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type ListTransactionResponseData struct {
	History  []types.DisputeHistory `json:"history"`
	Messages []types.DisputeMessage `json:"messages"`
	Dispute  *types.Dispute         `json:"dispute,omitempty"`
}

type ListTransactionResponse = types.Response[ListTransactionResponseData]

func (c *Client) ListTransactionDisputes(ctx context.Context, transactionID string) (*ListTransactionResponse, error) {
	if transactionID == "" {
		return nil, errors.New("transaction ID is required")
	}

	return net.Get[ListTransactionResponseData](ctx, c.Client, c.Secret, "/transaction/"+transactionID+"/disputes", c.BaseURL)
}
