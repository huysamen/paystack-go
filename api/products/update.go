package products

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type updateRequest struct {
	Name        *string         `json:"name,omitempty"`
	Description *string         `json:"description,omitempty"`
	Price       *int            `json:"price,omitempty"`
	Currency    *string         `json:"currency,omitempty"`
	Unlimited   *bool           `json:"unlimited,omitempty"`
	Quantity    *int            `json:"quantity,omitempty"`
	Metadata    *types.Metadata `json:"metadata,omitempty"`
}

type UpdateRequestBuilder struct {
	req *updateRequest
}

func NewUpdateRequestBuilder() *UpdateRequestBuilder {
	return &UpdateRequestBuilder{
		req: &updateRequest{},
	}
}

func (b *UpdateRequestBuilder) Name(name string) *UpdateRequestBuilder {
	b.req.Name = &name

	return b
}

func (b *UpdateRequestBuilder) Description(description string) *UpdateRequestBuilder {
	b.req.Description = &description

	return b
}

func (b *UpdateRequestBuilder) Price(price int) *UpdateRequestBuilder {
	b.req.Price = &price

	return b
}

func (b *UpdateRequestBuilder) Currency(currency string) *UpdateRequestBuilder {
	b.req.Currency = &currency

	return b
}

func (b *UpdateRequestBuilder) Unlimited(unlimited bool) *UpdateRequestBuilder {
	b.req.Unlimited = &unlimited

	return b
}

func (b *UpdateRequestBuilder) Quantity(quantity int) *UpdateRequestBuilder {
	b.req.Quantity = &quantity

	return b
}

func (b *UpdateRequestBuilder) Metadata(metadata *types.Metadata) *UpdateRequestBuilder {
	b.req.Metadata = metadata

	return b
}

func (b *UpdateRequestBuilder) Build() *updateRequest {
	return b.req
}

type UpdateResponseData = types.Product
type UpdateResponse = types.Response[UpdateResponseData]

func (c *Client) Update(ctx context.Context, productID string, builder UpdateRequestBuilder) (*UpdateResponse, error) {
	return net.Put[updateRequest, UpdateResponseData](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s", basePath, productID), builder.Build(), c.BaseURL)
}
