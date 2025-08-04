package customers

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type CustomerCreateRequest struct {
	Email     string         `json:"email"`
	FirstName *string        `json:"first_name,omitempty"`
	LastName  *string        `json:"last_name,omitempty"`
	Phone     *string        `json:"phone,omitempty"`
	Metadata  map[string]any `json:"metadata,omitempty"`
}

type CustomerCreateRequestBuilder struct {
	req *CustomerCreateRequest
}

func NewCreateCustomerRequest(email string) *CustomerCreateRequestBuilder {
	return &CustomerCreateRequestBuilder{
		req: &CustomerCreateRequest{
			Email: email,
		},
	}
}

func (b *CustomerCreateRequestBuilder) FirstName(firstName string) *CustomerCreateRequestBuilder {
	b.req.FirstName = &firstName

	return b
}

func (b *CustomerCreateRequestBuilder) LastName(lastName string) *CustomerCreateRequestBuilder {
	b.req.LastName = &lastName

	return b
}

func (b *CustomerCreateRequestBuilder) Phone(phone string) *CustomerCreateRequestBuilder {
	b.req.Phone = &phone

	return b
}

func (b *CustomerCreateRequestBuilder) Metadata(metadata map[string]any) *CustomerCreateRequestBuilder {
	b.req.Metadata = metadata

	return b
}

func (b *CustomerCreateRequestBuilder) Build() *CustomerCreateRequest {
	return b.req
}

type CustomerCreateResponse = types.Response[types.Customer]

func (c *Client) Create(ctx context.Context, builder *CustomerCreateRequestBuilder) (*CustomerCreateResponse, error) {
	return net.Post[CustomerCreateRequest, types.Customer](ctx, c.Client, c.Secret, basePath, builder.Build(), c.BaseURL)
}
