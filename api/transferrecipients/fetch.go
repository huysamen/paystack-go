package transferrecipients

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// Fetch retrieves a specific transfer recipient by ID or code
func (c *Client) Fetch(ctx context.Context, idOrCode string) (*types.Response[TransferRecipient], error) {
	return net.Get[TransferRecipient](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s", basePath, idOrCode), "", c.BaseURL)
}
