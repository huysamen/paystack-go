package customers

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type UpdateRequest struct {
	FirstName *string        `json:"first_name,omitempty"`
	LastName  *string        `json:"last_name,omitempty"`
	Phone     *string        `json:"phone,omitempty"`
	Metadata  map[string]any `json:"metadata,omitempty"`
}

type UpdateRequestBuilder struct {
	req *UpdateRequest
}

func NewUpdateRequest() *UpdateRequestBuilder {
	return &UpdateRequestBuilder{
		req: &UpdateRequest{},
	}
}

func (b *UpdateRequestBuilder) FirstName(firstName string) *UpdateRequestBuilder {
	b.req.FirstName = &firstName

	return b
}

func (b *UpdateRequestBuilder) LastName(lastName string) *UpdateRequestBuilder {
	b.req.LastName = &lastName

	return b
}

func (b *UpdateRequestBuilder) Phone(phone string) *UpdateRequestBuilder {
	b.req.Phone = &phone

	return b
}

func (b *UpdateRequestBuilder) Metadata(metadata map[string]any) *UpdateRequestBuilder {
	b.req.Metadata = metadata

	return b
}

func (b *UpdateRequestBuilder) Build() *UpdateRequest {
	return b.req
}

type UpdateResponseData = types.Customer
type UpdateResponse = types.Response[UpdateResponseData]

func (c *Client) Update(ctx context.Context, customerCode string, builder UpdateRequestBuilder) (*UpdateResponse, error) {
	path := fmt.Sprintf("%s/%s", basePath, customerCode)

	return net.Put[UpdateRequest, UpdateResponseData](ctx, c.Client, c.Secret, path, builder.Build(), c.BaseURL)
}
