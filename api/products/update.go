package products

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type UpdateProductRequest struct {
	Name        *string         `json:"name,omitempty"`
	Description *string         `json:"description,omitempty"`
	Price       *int            `json:"price,omitempty"`
	Currency    *string         `json:"currency,omitempty"`
	Unlimited   *bool           `json:"unlimited,omitempty"`
	Quantity    *int            `json:"quantity,omitempty"`
	Metadata    *types.Metadata `json:"metadata,omitempty"`
}

type UpdateProductRequestBuilder struct {
	req *UpdateProductRequest
}

func NewUpdateProductRequest() *UpdateProductRequestBuilder {
	return &UpdateProductRequestBuilder{
		req: &UpdateProductRequest{},
	}
}

func (b *UpdateProductRequestBuilder) Name(name string) *UpdateProductRequestBuilder {
	b.req.Name = &name

	return b
}

func (b *UpdateProductRequestBuilder) Description(description string) *UpdateProductRequestBuilder {
	b.req.Description = &description

	return b
}

func (b *UpdateProductRequestBuilder) Price(price int) *UpdateProductRequestBuilder {
	b.req.Price = &price

	return b
}

func (b *UpdateProductRequestBuilder) Currency(currency string) *UpdateProductRequestBuilder {
	b.req.Currency = &currency

	return b
}

func (b *UpdateProductRequestBuilder) Unlimited(unlimited bool) *UpdateProductRequestBuilder {
	b.req.Unlimited = &unlimited

	return b
}

func (b *UpdateProductRequestBuilder) Quantity(quantity int) *UpdateProductRequestBuilder {
	b.req.Quantity = &quantity

	return b
}

func (b *UpdateProductRequestBuilder) Metadata(metadata *types.Metadata) *UpdateProductRequestBuilder {
	b.req.Metadata = metadata

	return b
}

func (b *UpdateProductRequestBuilder) Build() *UpdateProductRequest {
	return b.req
}

type UpdateProductResponse = types.Response[types.Product]

func (c *Client) Update(ctx context.Context, productID string, builder *UpdateProductRequestBuilder) (*UpdateProductResponse, error) {
	return net.Put[UpdateProductRequest, types.Product](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s", basePath, productID), builder.Build(), c.BaseURL)
}
