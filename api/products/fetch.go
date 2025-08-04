package products

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// Fetch gets details of a product on your integration
func (c *Client) Fetch(ctx context.Context, productID string) (*FetchProductResponse, error) {
	return net.Get[types.Product](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s", basePath, productID), c.BaseURL)
}
