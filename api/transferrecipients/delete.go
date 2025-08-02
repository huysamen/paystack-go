package transferrecipients

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// Delete deletes a transfer recipient (sets it to inactive)
func (c *Client) Delete(ctx context.Context, idOrCode string) (*types.Response[any], error) {
	return net.Delete[any](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s", basePath, idOrCode), c.BaseURL)
}
