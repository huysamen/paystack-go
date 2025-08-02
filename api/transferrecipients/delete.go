package transferrecipients

import (
	"context"
	"errors"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// Delete deletes a transfer recipient (sets it to inactive)
func (c *Client) Delete(ctx context.Context, idOrCode string) (*types.Response[any], error) {
	if idOrCode == "" {
		return nil, errors.New("id or code is required")
	}

	endpoint := fmt.Sprintf("%s/%s", transferRecipientBasePath, idOrCode)
	return net.Delete[any](ctx, c.client, c.secret, endpoint, c.baseURL)
}
