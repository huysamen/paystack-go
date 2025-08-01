package transfers

import (
	"context"
	"errors"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type TransferInitiateRequest struct {
	Source           string  `json:"source"`                      // Only "balance" supported for now
	Amount           int     `json:"amount"`                      // Amount in kobo (NGN) or pesewas (GHS)
	Recipient        string  `json:"recipient"`                   // Transfer recipient code
	Reason           *string `json:"reason,omitempty"`            // Reason for transfer
	Currency         *string `json:"currency,omitempty"`          // Defaults to NGN
	AccountReference *string `json:"account_reference,omitempty"` // Required for MPESA in Kenya
	Reference        *string `json:"reference,omitempty"`         // Unique identifier for transfer
}

func (c *Client) Initiate(ctx context.Context, req *TransferInitiateRequest) (*types.Response[Transfer], error) {
	if req == nil {
		return nil, errors.New("request cannot be nil")
	}

	return net.Post[TransferInitiateRequest, Transfer](
		ctx,
		c.client,
		c.secret,
		transferBasePath,
		req,
		c.baseURL,
	)
}
