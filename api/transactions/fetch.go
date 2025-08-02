package transactions

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

func (c *Client) Fetch(ctx context.Context, id uint64) (*types.Response[types.Transaction], error) {
	return net.Get[types.Transaction](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%d", basePath, id), "", c.BaseURL)
}
