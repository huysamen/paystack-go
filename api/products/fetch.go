package products

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type FetchProductResponse = types.Response[types.Product]

func (c *Client) Fetch(ctx context.Context, productID string) (*FetchProductResponse, error) {
	return net.Get[types.Product](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s", basePath, productID), c.BaseURL)
}
