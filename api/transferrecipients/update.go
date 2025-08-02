package transferrecipients

import (
	"context"
	"errors"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// Update updates an existing transfer recipient
func (c *Client) Update(ctx context.Context, idOrCode string, builder *TransferRecipientUpdateRequestBuilder) (*types.Response[TransferRecipient], error) {
	if idOrCode == "" {
		return nil, errors.New("id or code is required")
	}
	if builder == nil {
		return nil, ErrBuilderRequired
	}

	req := builder.Build()
	endpoint := fmt.Sprintf("%s/%s", transferRecipientBasePath, idOrCode)

	return net.Put[TransferRecipientUpdateRequest, TransferRecipient](
		ctx, c.client, c.secret, endpoint, req, c.baseURL,
	)
}
