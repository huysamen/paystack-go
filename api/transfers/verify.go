package transfers

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

func (c *Client) Verify(ctx context.Context, reference string) (*types.Response[Transfer], error) {
	return net.Get[Transfer](ctx, c.Client, c.Secret, fmt.Sprintf("%s/verify/%s", basePath, reference), "", c.BaseURL)
}
