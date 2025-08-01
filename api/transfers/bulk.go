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

func (c *Client) Bulk(ctx context.Context, req *BulkTransferRequest) (*types.Response[[]BulkTransferResponse], error) {
	if req == nil {
		return nil, errors.New("request cannot be nil")
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
