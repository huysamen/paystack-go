package transferrecipients

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// Create creates a new transfer recipient
func (c *Client) Create(ctx context.Context, builder *TransferRecipientCreateRequestBuilder) (*types.Response[TransferRecipient], error) {
	req := builder.Build()
	return net.Post[TransferRecipientCreateRequest, TransferRecipient](ctx, c.Client, c.Secret, basePath, req, c.BaseURL)
}
