package products

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// CreateProductRequest represents the request to create a product
type CreateProductRequest struct {
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Price       int             `json:"price"`
	Currency    string          `json:"currency"`
	Unlimited   *bool           `json:"unlimited,omitempty"`
	Quantity    *int            `json:"quantity,omitempty"`
	Metadata    *types.Metadata `json:"metadata,omitempty"`
}

// CreateProductRequestBuilder provides a fluent interface for building CreateProductRequest
type CreateProductRequestBuilder struct {
	req *CreateProductRequest
}

// NewCreateProductRequest creates a new builder for CreateProductRequest
func NewCreateProductRequest(name, description string, price int, currency string) *CreateProductRequestBuilder {
	return &CreateProductRequestBuilder{
		req: &CreateProductRequest{
			Name:        name,
			Description: description,
			Price:       price,
			Currency:    currency,
		},
	}
}

// Unlimited sets whether the product has unlimited quantity
func (b *CreateProductRequestBuilder) Unlimited(unlimited bool) *CreateProductRequestBuilder {
	b.req.Unlimited = &unlimited
	return b
}

// Quantity sets the product quantity (ignored if unlimited is true)
func (b *CreateProductRequestBuilder) Quantity(quantity int) *CreateProductRequestBuilder {
	b.req.Quantity = &quantity
	return b
}

// Metadata sets the product metadata
func (b *CreateProductRequestBuilder) Metadata(metadata *types.Metadata) *CreateProductRequestBuilder {
	b.req.Metadata = metadata
	return b
}

// Build returns the constructed CreateProductRequest
func (b *CreateProductRequestBuilder) Build() *CreateProductRequest {
	return b.req
}

// Create creates a product on your integration
func (c *Client) Create(ctx context.Context, builder *CreateProductRequestBuilder) (*Product, error) {
	if builder == nil {
		return nil, ErrBuilderRequired
	}

	resp, err := net.Post[CreateProductRequest, Product](
		ctx, c.client, c.secret, productsBasePath, builder.Build(), c.baseURL,
	)
	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}
