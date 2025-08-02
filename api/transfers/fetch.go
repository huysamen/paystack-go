package transfers

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

func (c *Client) Fetch(ctx context.Context, idOrCode string) (*types.Response[Transfer], error) {
	return net.Get[Transfer](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s", basePath, idOrCode), "", c.BaseURL)
}
