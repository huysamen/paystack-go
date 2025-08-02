package products

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
)

// Fetch gets details of a product on your integration
func (c *Client) Fetch(ctx context.Context, productID string) (*FetchProductResponse, error) {
	return net.Get[Product](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s", basePath, productID), c.BaseURL)
}
