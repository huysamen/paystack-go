package plans

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// UpdatePlanRequest represents the request to update a plan
type UpdatePlanRequest struct {
	// Required fields
	Name     string         `json:"name"`
	Amount   int            `json:"amount"`
	Interval types.Interval `json:"interval"`

	// Optional fields
	Description                 string         `json:"description,omitempty"`
	SendInvoices                *bool          `json:"send_invoices,omitempty"`
	SendSMS                     *bool          `json:"send_sms,omitempty"`
	Currency                    types.Currency `json:"currency,omitempty"`
	InvoiceLimit                *int           `json:"invoice_limit,omitempty"`
	UpdateExistingSubscriptions *bool          `json:"update_existing_subscriptions,omitempty"`
}

// UpdatePlanRequestBuilder provides a fluent interface for building UpdatePlanRequest
type UpdatePlanRequestBuilder struct {
	req *UpdatePlanRequest
}

// NewUpdatePlanRequest creates a new builder for UpdatePlanRequest
func NewUpdatePlanRequest(name string, amount int, interval types.Interval) *UpdatePlanRequestBuilder {
	return &UpdatePlanRequestBuilder{
		req: &UpdatePlanRequest{
			Name:     name,
			Amount:   amount,
			Interval: interval,
		},
	}
}

// Description sets the plan description
func (b *UpdatePlanRequestBuilder) Description(description string) *UpdatePlanRequestBuilder {
	b.req.Description = description

	return b
}

// SendInvoices sets whether to send invoices to subscribers
func (b *UpdatePlanRequestBuilder) SendInvoices(sendInvoices bool) *UpdatePlanRequestBuilder {
	b.req.SendInvoices = &sendInvoices

	return b
}

// SendSMS sets whether to send SMS to subscribers
func (b *UpdatePlanRequestBuilder) SendSMS(sendSMS bool) *UpdatePlanRequestBuilder {
	b.req.SendSMS = &sendSMS

	return b
}

// Currency sets the plan currency
func (b *UpdatePlanRequestBuilder) Currency(currency types.Currency) *UpdatePlanRequestBuilder {
	b.req.Currency = currency

	return b
}

// InvoiceLimit sets the maximum number of invoices for the plan
func (b *UpdatePlanRequestBuilder) InvoiceLimit(limit int) *UpdatePlanRequestBuilder {
	b.req.InvoiceLimit = &limit

	return b
}

// UpdateExistingSubscriptions sets whether to update existing subscriptions
func (b *UpdatePlanRequestBuilder) UpdateExistingSubscriptions(update bool) *UpdatePlanRequestBuilder {
	b.req.UpdateExistingSubscriptions = &update

	return b
}

// Build returns the constructed UpdatePlanRequest
func (b *UpdatePlanRequestBuilder) Build() *UpdatePlanRequest {
	return b.req
}

// UpdatePlanResponse represents the response from updating a plan
type UpdatePlanResponse = types.Response[any]

// Update updates an existing subscription plan
func (c *Client) Update(ctx context.Context, idOrCode string, builder *UpdatePlanRequestBuilder) (*UpdatePlanResponse, error) {
	return net.Put[UpdatePlanRequest, any](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s", basePath, idOrCode), builder.Build(), c.BaseURL)
}
