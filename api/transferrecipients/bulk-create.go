package transferrecipients

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// BulkCreate creates multiple transfer recipients in a single request
func (c *Client) BulkCreate(ctx context.Context, builder *BulkCreateTransferRecipientRequestBuilder) (*types.Response[BulkCreateResult], error) {
	req := builder.Build()
	return net.Post[BulkCreateTransferRecipientRequest, BulkCreateResult](ctx, c.Client, c.Secret, basePath+"/bulk", req, c.BaseURL)
}
