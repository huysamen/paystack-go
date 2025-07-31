package products

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
)

// Fetch gets details of a product on your integration
func (c *Client) Fetch(ctx context.Context, productID string) (*Product, error) {
	if productID == "" {
		return nil, fmt.Errorf("productID is required")
	}

	path := fmt.Sprintf("%s/%s", productsBasePath, productID)

	resp, err := net.Get[Product](
		ctx,
		c.client,
		c.secret,
		path,
		c.baseURL,
	)
	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}
