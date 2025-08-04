package products

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// UpdateProductRequest represents the request to update a product
type UpdateProductRequest struct {
	Name        *string         `json:"name,omitempty"`
	Description *string         `json:"description,omitempty"`
	Price       *int            `json:"price,omitempty"`
	Currency    *string         `json:"currency,omitempty"`
	Unlimited   *bool           `json:"unlimited,omitempty"`
	Quantity    *int            `json:"quantity,omitempty"`
	Metadata    *types.Metadata `json:"metadata,omitempty"`
}

// UpdateProductRequestBuilder provides a fluent interface for building UpdateProductRequest
type UpdateProductRequestBuilder struct {
	req *UpdateProductRequest
}

// NewUpdateProductRequest creates a new builder for UpdateProductRequest
func NewUpdateProductRequest() *UpdateProductRequestBuilder {
	return &UpdateProductRequestBuilder{
		req: &UpdateProductRequest{},
	}
}

// Name sets the product name
func (b *UpdateProductRequestBuilder) Name(name string) *UpdateProductRequestBuilder {
	b.req.Name = &name
	return b
}

// Description sets the product description
func (b *UpdateProductRequestBuilder) Description(description string) *UpdateProductRequestBuilder {
	b.req.Description = &description
	return b
}

// Price sets the product price
func (b *UpdateProductRequestBuilder) Price(price int) *UpdateProductRequestBuilder {
	b.req.Price = &price
	return b
}

// Currency sets the product currency
func (b *UpdateProductRequestBuilder) Currency(currency string) *UpdateProductRequestBuilder {
	b.req.Currency = &currency
	return b
}

// Unlimited sets whether the product has unlimited quantity
func (b *UpdateProductRequestBuilder) Unlimited(unlimited bool) *UpdateProductRequestBuilder {
	b.req.Unlimited = &unlimited
	return b
}

// Quantity sets the product quantity
func (b *UpdateProductRequestBuilder) Quantity(quantity int) *UpdateProductRequestBuilder {
	b.req.Quantity = &quantity
	return b
}

// Metadata sets the product metadata
func (b *UpdateProductRequestBuilder) Metadata(metadata *types.Metadata) *UpdateProductRequestBuilder {
	b.req.Metadata = metadata
	return b
}

// Build returns the constructed UpdateProductRequest
func (b *UpdateProductRequestBuilder) Build() *UpdateProductRequest {
	return b.req
}

// Update modifies a product details on your integration
func (c *Client) Update(ctx context.Context, productID string, builder *UpdateProductRequestBuilder) (*UpdateProductResponse, error) {
	return net.Put[UpdateProductRequest, types.Product](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s", basePath, productID), builder.Build(), c.BaseURL)
}
