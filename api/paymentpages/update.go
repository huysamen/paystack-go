package paymentpages

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// UpdatePaymentPageRequest represents the request to update a payment page
type UpdatePaymentPageRequest struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Amount      *int   `json:"amount,omitempty"`
	Active      *bool  `json:"active,omitempty"`
}

// UpdatePaymentPageRequestBuilder provides a fluent interface for building UpdatePaymentPageRequest
type UpdatePaymentPageRequestBuilder struct {
	req *UpdatePaymentPageRequest
}

// NewUpdatePaymentPageRequest creates a new builder for UpdatePaymentPageRequest
func NewUpdatePaymentPageRequest() *UpdatePaymentPageRequestBuilder {
	return &UpdatePaymentPageRequestBuilder{
		req: &UpdatePaymentPageRequest{},
	}
}

// Name sets the name for the payment page
func (b *UpdatePaymentPageRequestBuilder) Name(name string) *UpdatePaymentPageRequestBuilder {
	b.req.Name = name
	return b
}

// Description sets the description for the payment page
func (b *UpdatePaymentPageRequestBuilder) Description(description string) *UpdatePaymentPageRequestBuilder {
	b.req.Description = description
	return b
}

// Amount sets the amount for the payment page (in kobo/cents)
func (b *UpdatePaymentPageRequestBuilder) Amount(amount int) *UpdatePaymentPageRequestBuilder {
	b.req.Amount = &amount
	return b
}

// Active sets whether the payment page is active
func (b *UpdatePaymentPageRequestBuilder) Active(active bool) *UpdatePaymentPageRequestBuilder {
	b.req.Active = &active
	return b
}

// Build returns the constructed UpdatePaymentPageRequest
func (b *UpdatePaymentPageRequestBuilder) Build() *UpdatePaymentPageRequest {
	return b.req
}

// UpdatePaymentPageResponse represents the response from updating a payment page
type UpdatePaymentPageResponse = types.Response[PaymentPage]

// Update updates a payment page details on your integration using the builder pattern
func (c *Client) Update(ctx context.Context, idOrSlug string, builder *UpdatePaymentPageRequestBuilder) (*UpdatePaymentPageResponse, error) {
	return net.Put[UpdatePaymentPageRequest, PaymentPage](ctx, c.Client, c.Secret, basePath+"/"+idOrSlug, builder.Build(), c.BaseURL)
}
