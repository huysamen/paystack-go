package products

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type createRequest struct {
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Price       int             `json:"price"`
	Currency    string          `json:"currency"`
	Unlimited   *bool           `json:"unlimited,omitempty"`
	Quantity    *int            `json:"quantity,omitempty"`
	Metadata    *types.Metadata `json:"metadata,omitempty"`
}

type CreateRequestBuilder struct {
	req *createRequest
}

func NewCreateRequestBuilder(name, description string, price int, currency string) *CreateRequestBuilder {
	return &CreateRequestBuilder{
		req: &createRequest{
			Name:        name,
			Description: description,
			Price:       price,
			Currency:    currency,
		},
	}
}

func (b *CreateRequestBuilder) Unlimited(unlimited bool) *CreateRequestBuilder {
	b.req.Unlimited = &unlimited

	return b
}

func (b *CreateRequestBuilder) Quantity(quantity int) *CreateRequestBuilder {
	b.req.Quantity = &quantity

	return b
}

func (b *CreateRequestBuilder) Metadata(metadata *types.Metadata) *CreateRequestBuilder {
	b.req.Metadata = metadata

	return b
}

func (b *CreateRequestBuilder) Build() *createRequest {
	return b.req
}

type CreateResponseData = types.Product
type CreateResponse = types.Response[CreateResponseData]

func (c *Client) Create(ctx context.Context, builder CreateRequestBuilder) (*CreateResponse, error) {
	return net.Post[createRequest, CreateResponseData](ctx, c.Client, c.Secret, basePath, builder.Build(), c.BaseURL)
}
