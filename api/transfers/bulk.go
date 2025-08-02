package transfers

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type BulkTransferRequest struct {
	Source    string             `json:"source"`    // Only "balance" supported for now
	Currency  *string            `json:"currency"`  // Optional, defaults to NGN
	Transfers []BulkTransferItem `json:"transfers"` // List of transfers
}

func (c *Client) Bulk(ctx context.Context, req *BulkTransferRequest) (*types.Response[[]BulkTransferResponse], error) {
	return net.Post[BulkTransferRequest, []BulkTransferResponse](ctx, c.Client, c.Secret, basePath+"/bulk", req, c.BaseURL)
}
