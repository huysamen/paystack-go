package paymentpages

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type updateRequest struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Amount      *int   `json:"amount,omitempty"`
	Active      *bool  `json:"active,omitempty"`
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
	b.req.Name = name

	return b
}

func (b *UpdateRequestBuilder) Description(description string) *UpdateRequestBuilder {
	b.req.Description = description

	return b
}

func (b *UpdateRequestBuilder) Amount(amount int) *UpdateRequestBuilder {
	b.req.Amount = &amount

	return b
}

func (b *UpdateRequestBuilder) Active(active bool) *UpdateRequestBuilder {
	b.req.Active = &active

	return b
}

func (b *UpdateRequestBuilder) Build() *updateRequest {
	return b.req
}

type UpdateResponseData = types.PaymentPage
type UpdateResponse = types.Response[UpdateResponseData]

func (c *Client) Update(ctx context.Context, idOrSlug string, builder UpdateRequestBuilder) (*UpdateResponse, error) {
	return net.Put[updateRequest, UpdateResponseData](ctx, c.Client, c.Secret, basePath+"/"+idOrSlug, builder.Build(), c.BaseURL)
}
