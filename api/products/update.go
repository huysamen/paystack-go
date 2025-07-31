package products

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
)

// Update modifies a product details on your integration
func (c *Client) Update(ctx context.Context, productID string, req *UpdateProductRequest) (*Product, error) {
	if productID == "" {
		return nil, fmt.Errorf("productID is required")
	}

	if err := validateUpdateRequest(req); err != nil {
		return nil, err
	}

	path := fmt.Sprintf("%s/%s", productsBasePath, productID)

	resp, err := net.Put[UpdateProductRequest, Product](
		ctx, c.client, c.secret, path, req, c.baseURL,
	)
	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}

func validateUpdateRequest(req *UpdateProductRequest) error {
	if req == nil {
		return fmt.Errorf("request cannot be nil")
	}

	// At least one field should be provided for update
	if req.Name == nil && req.Description == nil && req.Price == nil &&
		req.Currency == nil && req.Unlimited == nil && req.Quantity == nil &&
		req.Metadata == nil {
		return fmt.Errorf("at least one field must be provided for update")
	}

	// Validate price if provided
	if req.Price != nil && *req.Price <= 0 {
		return fmt.Errorf("price must be greater than 0")
	}

	// If unlimited is false, quantity should be provided or existing
	if req.Unlimited != nil && !*req.Unlimited && req.Quantity != nil && *req.Quantity <= 0 {
		return fmt.Errorf("quantity must be greater than 0 when unlimited is false")
	}

	return nil
}
