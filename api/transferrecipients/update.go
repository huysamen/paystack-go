package transferrecipients

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
)

// Update updates a transfer recipient
func (c *Client) Update(ctx context.Context, idOrCode string, builder *TransferRecipientUpdateRequestBuilder) (*TransferRecipientUpdateResponse, error) {
	req := builder.Build()
	return net.Put[TransferRecipientUpdateRequest, TransferRecipient](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s", basePath, idOrCode), req, c.BaseURL)
}
