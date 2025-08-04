package products

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type CreateProductRequest struct {
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Price       int             `json:"price"`
	Currency    string          `json:"currency"`
	Unlimited   *bool           `json:"unlimited,omitempty"`
	Quantity    *int            `json:"quantity,omitempty"`
	Metadata    *types.Metadata `json:"metadata,omitempty"`
}

type CreateProductRequestBuilder struct {
	req *CreateProductRequest
}

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

func (b *CreateProductRequestBuilder) Unlimited(unlimited bool) *CreateProductRequestBuilder {
	b.req.Unlimited = &unlimited

	return b
}

func (b *CreateProductRequestBuilder) Quantity(quantity int) *CreateProductRequestBuilder {
	b.req.Quantity = &quantity

	return b
}

func (b *CreateProductRequestBuilder) Metadata(metadata *types.Metadata) *CreateProductRequestBuilder {
	b.req.Metadata = metadata

	return b
}

func (b *CreateProductRequestBuilder) Build() *CreateProductRequest {
	return b.req
}

type CreateProductResponse = types.Response[types.Product]

func (c *Client) Create(ctx context.Context, builder *CreateProductRequestBuilder) (*CreateProductResponse, error) {
	return net.Post[CreateProductRequest, types.Product](ctx, c.Client, c.Secret, basePath, builder.Build(), c.BaseURL)
}
