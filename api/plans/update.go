package plans

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type UpdatePlanRequest struct {
	Name     string         `json:"name"`
	Amount   int            `json:"amount"`
	Interval types.Interval `json:"interval"`

	Description                 string         `json:"description,omitempty"`
	SendInvoices                *bool          `json:"send_invoices,omitempty"`
	SendSMS                     *bool          `json:"send_sms,omitempty"`
	Currency                    types.Currency `json:"currency,omitempty"`
	InvoiceLimit                *int           `json:"invoice_limit,omitempty"`
	UpdateExistingSubscriptions *bool          `json:"update_existing_subscriptions,omitempty"`
}

type UpdatePlanRequestBuilder struct {
	req *UpdatePlanRequest
}

func NewUpdatePlanRequest(name string, amount int, interval types.Interval) *UpdatePlanRequestBuilder {
	return &UpdatePlanRequestBuilder{
		req: &UpdatePlanRequest{
			Name:     name,
			Amount:   amount,
			Interval: interval,
		},
	}
}

func (b *UpdatePlanRequestBuilder) Description(description string) *UpdatePlanRequestBuilder {
	b.req.Description = description

	return b
}

func (b *UpdatePlanRequestBuilder) SendInvoices(sendInvoices bool) *UpdatePlanRequestBuilder {
	b.req.SendInvoices = &sendInvoices

	return b
}

func (b *UpdatePlanRequestBuilder) SendSMS(sendSMS bool) *UpdatePlanRequestBuilder {
	b.req.SendSMS = &sendSMS

	return b
}

func (b *UpdatePlanRequestBuilder) Currency(currency types.Currency) *UpdatePlanRequestBuilder {
	b.req.Currency = currency

	return b
}

func (b *UpdatePlanRequestBuilder) InvoiceLimit(limit int) *UpdatePlanRequestBuilder {
	b.req.InvoiceLimit = &limit

	return b
}

func (b *UpdatePlanRequestBuilder) UpdateExistingSubscriptions(update bool) *UpdatePlanRequestBuilder {
	b.req.UpdateExistingSubscriptions = &update

	return b
}

func (b *UpdatePlanRequestBuilder) Build() *UpdatePlanRequest {
	return b.req
}

type UpdatePlanResponse = types.Response[any]

func (c *Client) Update(ctx context.Context, idOrCode string, builder *UpdatePlanRequestBuilder) (*UpdatePlanResponse, error) {
	return net.Put[UpdatePlanRequest, any](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s", basePath, idOrCode), builder.Build(), c.BaseURL)
}
