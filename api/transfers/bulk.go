package transfers

import (
	"context"
	"errors"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type BulkTransferRequest struct {
	Source    string             `json:"source"`    // Only "balance" supported for now
	Currency  *string            `json:"currency"`  // Optional, defaults to NGN
	Transfers []BulkTransferItem `json:"transfers"` // List of transfers
}

func (r *BulkTransferRequest) Validate() error {
	if r.Source == "" {
		return errors.New("source is required")
	}
	if r.Source != "balance" {
		return errors.New("only 'balance' source is supported")
	}
	if len(r.Transfers) == 0 {
		return errors.New("transfers list cannot be empty")
	}

	for i, transfer := range r.Transfers {
		if transfer.Amount <= 0 {
			return errors.New("all transfers must have amount greater than 0")
		}
		if transfer.Recipient == "" {
			return errors.New("all transfers must have a recipient")
		}
		if transfer.Reference == "" {
			return errors.New("all transfers must have a reference")
		}

		// Check for duplicate references
		for j, other := range r.Transfers {
			if i != j && transfer.Reference == other.Reference {
				return errors.New("transfer references must be unique")
			}
		}
	}

	return nil
}

func (c *Client) Bulk(ctx context.Context, req *BulkTransferRequest) (*types.Response[[]BulkTransferResponse], error) {
	if req == nil {
		return nil, errors.New("request cannot be nil")
	}

	if err := req.Validate(); err != nil {
		return nil, err
	}

	path := transferBasePath + "/bulk"

	return net.Post[BulkTransferRequest, []BulkTransferResponse](
		ctx,
		c.client,
		c.secret,
		path,
		req,
		c.baseURL,
	)
}
