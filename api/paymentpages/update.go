package paymentpages

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type UpdatePaymentPageRequest struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Amount      *int   `json:"amount,omitempty"`
	Active      *bool  `json:"active,omitempty"`
}

type UpdatePaymentPageRequestBuilder struct {
	req *UpdatePaymentPageRequest
}

func NewUpdatePaymentPageRequest() *UpdatePaymentPageRequestBuilder {
	return &UpdatePaymentPageRequestBuilder{
		req: &UpdatePaymentPageRequest{},
	}
}

func (b *UpdatePaymentPageRequestBuilder) Name(name string) *UpdatePaymentPageRequestBuilder {
	b.req.Name = name

	return b
}

func (b *UpdatePaymentPageRequestBuilder) Description(description string) *UpdatePaymentPageRequestBuilder {
	b.req.Description = description

	return b
}

func (b *UpdatePaymentPageRequestBuilder) Amount(amount int) *UpdatePaymentPageRequestBuilder {
	b.req.Amount = &amount

	return b
}

func (b *UpdatePaymentPageRequestBuilder) Active(active bool) *UpdatePaymentPageRequestBuilder {
	b.req.Active = &active

	return b
}

func (b *UpdatePaymentPageRequestBuilder) Build() *UpdatePaymentPageRequest {
	return b.req
}

type UpdatePaymentPageResponse = types.Response[types.PaymentPage]

func (c *Client) Update(ctx context.Context, idOrSlug string, builder *UpdatePaymentPageRequestBuilder) (*UpdatePaymentPageResponse, error) {
	return net.Put[UpdatePaymentPageRequest, types.PaymentPage](ctx, c.Client, c.Secret, basePath+"/"+idOrSlug, builder.Build(), c.BaseURL)
}
