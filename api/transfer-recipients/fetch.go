package transfer_recipients

import (
	"context"
	"errors"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// Fetch retrieves a specific transfer recipient by ID or code
func (c *Client) Fetch(ctx context.Context, idOrCode string) (*types.Response[TransferRecipient], error) {
	if idOrCode == "" {
		return nil, errors.New("id or code is required")
	}

	endpoint := fmt.Sprintf("%s/%s", transferRecipientBasePath, idOrCode)
	return net.Get[TransferRecipient](ctx, c.client, c.secret, endpoint, "", c.baseURL)
}
