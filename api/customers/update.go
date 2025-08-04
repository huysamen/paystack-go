package customers

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type CustomerUpdateRequest struct {
	FirstName *string        `json:"first_name,omitempty"`
	LastName  *string        `json:"last_name,omitempty"`
	Phone     *string        `json:"phone,omitempty"`
	Metadata  map[string]any `json:"metadata,omitempty"`
}

type CustomerUpdateRequestBuilder struct {
	req *CustomerUpdateRequest
}

func NewUpdateCustomerRequest() *CustomerUpdateRequestBuilder {
	return &CustomerUpdateRequestBuilder{
		req: &CustomerUpdateRequest{},
	}
}

func (b *CustomerUpdateRequestBuilder) FirstName(firstName string) *CustomerUpdateRequestBuilder {
	b.req.FirstName = &firstName

	return b
}

func (b *CustomerUpdateRequestBuilder) LastName(lastName string) *CustomerUpdateRequestBuilder {
	b.req.LastName = &lastName

	return b
}

func (b *CustomerUpdateRequestBuilder) Phone(phone string) *CustomerUpdateRequestBuilder {
	b.req.Phone = &phone

	return b
}

func (b *CustomerUpdateRequestBuilder) Metadata(metadata map[string]any) *CustomerUpdateRequestBuilder {
	b.req.Metadata = metadata

	return b
}

func (b *CustomerUpdateRequestBuilder) Build() *CustomerUpdateRequest {
	return b.req
}

type UpdateCustomerResponse = types.Response[types.Customer]

func (c *Client) Update(ctx context.Context, customerCode string, builder *CustomerUpdateRequestBuilder) (*UpdateCustomerResponse, error) {
	path := fmt.Sprintf("%s/%s", basePath, customerCode)

	return net.Put[CustomerUpdateRequest, types.Customer](ctx, c.Client, c.Secret, path, builder.Build(), c.BaseURL)
}
