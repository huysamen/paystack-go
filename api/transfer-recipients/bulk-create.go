package transfer_recipients

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// BulkCreate creates multiple transfer recipients in a single request
func (c *Client) BulkCreate(ctx context.Context, builder *BulkCreateTransferRecipientRequestBuilder) (*types.Response[BulkCreateResult], error) {
	if builder == nil {
		return nil, ErrBuilderRequired
	}

	req := builder.Build()
	return net.Post[BulkCreateTransferRecipientRequest, BulkCreateResult](
		ctx, c.client, c.secret, transferRecipientBasePath+"/bulk", req, c.baseURL,
	)
}
