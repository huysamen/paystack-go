package products

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// FetchProductResponse represents the response from fetching a product
type FetchProductResponse = types.Response[Product]

// Fetch gets details of a product on your integration
func (c *Client) Fetch(ctx context.Context, productID string) (*FetchProductResponse, error) {
	return net.Get[Product](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s", basePath, productID), c.BaseURL)
}
