package transfer_recipients

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// Create creates a new transfer recipient
func (c *Client) Create(ctx context.Context, builder *TransferRecipientCreateRequestBuilder) (*types.Response[TransferRecipient], error) {
	if builder == nil {
		return nil, ErrBuilderRequired
	}

	req := builder.Build()
	return net.Post[TransferRecipientCreateRequest, TransferRecipient](
		ctx, c.client, c.secret, transferRecipientBasePath, req, c.baseURL,
	)
}
