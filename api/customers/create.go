package customers

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type CreateRequest struct {
	Email     string         `json:"email"`
	FirstName *string        `json:"first_name,omitempty"`
	LastName  *string        `json:"last_name,omitempty"`
	Phone     *string        `json:"phone,omitempty"`
	Metadata  map[string]any `json:"metadata,omitempty"`
}

type CreateRequestBuilder struct {
	req *CreateRequest
}

func NewCreateRequest(email string) *CreateRequestBuilder {
	return &CreateRequestBuilder{
		req: &CreateRequest{
			Email: email,
		},
	}
}

func (b *CreateRequestBuilder) FirstName(firstName string) *CreateRequestBuilder {
	b.req.FirstName = &firstName

	return b
}

func (b *CreateRequestBuilder) LastName(lastName string) *CreateRequestBuilder {
	b.req.LastName = &lastName

	return b
}

func (b *CreateRequestBuilder) Phone(phone string) *CreateRequestBuilder {
	b.req.Phone = &phone

	return b
}

func (b *CreateRequestBuilder) Metadata(metadata map[string]any) *CreateRequestBuilder {
	b.req.Metadata = metadata

	return b
}

func (b *CreateRequestBuilder) Build() *CreateRequest {
	return b.req
}

type CreateResponseData = types.Customer
type CreateResponse = types.Response[CreateResponseData]

func (c *Client) Create(ctx context.Context, builder CreateRequestBuilder) (*CreateResponse, error) {
	return net.Post[CreateRequest, CreateResponseData](ctx, c.Client, c.Secret, basePath, builder.Build(), c.BaseURL)
}
