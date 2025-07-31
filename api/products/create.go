package products

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
)

// Create creates a product on your integration
func (c *Client) Create(ctx context.Context, req *CreateProductRequest) (*Product, error) {
	if err := validateCreateRequest(req); err != nil {
		return nil, err
	}

	resp, err := net.Post[CreateProductRequest, Product](
		ctx, c.client, c.secret, productsBasePath, req, c.baseURL,
	)
	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}

func validateCreateRequest(req *CreateProductRequest) error {
	if req == nil {
		return fmt.Errorf("request cannot be nil")
	}

	if req.Name == "" {
		return fmt.Errorf("name is required")
	}

	if req.Description == "" {
		return fmt.Errorf("description is required")
	}

	if req.Price <= 0 {
		return fmt.Errorf("price must be greater than 0")
	}

	if req.Currency == "" {
		return fmt.Errorf("currency is required")
	}

	// If unlimited is false, quantity should be provided
	if req.Unlimited != nil && !*req.Unlimited && (req.Quantity == nil || *req.Quantity <= 0) {
		return fmt.Errorf("quantity must be greater than 0 when unlimited is false")
	}

	return nil
}
