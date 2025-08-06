package plans

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/enums"
	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type updateRequest struct {
	Name     string         `json:"name"`
	Amount   int            `json:"amount"`
	Interval enums.Interval `json:"interval"`

	Description                 string         `json:"description,omitempty"`
	SendInvoices                *bool          `json:"send_invoices,omitempty"`
	SendSMS                     *bool          `json:"send_sms,omitempty"`
	Currency                    enums.Currency `json:"currency,omitempty"`
	InvoiceLimit                *int           `json:"invoice_limit,omitempty"`
	UpdateExistingSubscriptions *bool          `json:"update_existing_subscriptions,omitempty"`
}

type UpdateRequestBuilder struct {
	req *updateRequest
}

func NewUpdateRequestBuilder(name string, amount int, interval enums.Interval) *UpdateRequestBuilder {
	return &UpdateRequestBuilder{
		req: &updateRequest{
			Name:     name,
			Amount:   amount,
			Interval: interval,
		},
	}
}

func (b *UpdateRequestBuilder) Description(description string) *UpdateRequestBuilder {
	b.req.Description = description

	return b
}

func (b *UpdateRequestBuilder) SendInvoices(sendInvoices bool) *UpdateRequestBuilder {
	b.req.SendInvoices = &sendInvoices

	return b
}

func (b *UpdateRequestBuilder) SendSMS(sendSMS bool) *UpdateRequestBuilder {
	b.req.SendSMS = &sendSMS

	return b
}

func (b *UpdateRequestBuilder) Currency(currency enums.Currency) *UpdateRequestBuilder {
	b.req.Currency = currency

	return b
}

func (b *UpdateRequestBuilder) InvoiceLimit(limit int) *UpdateRequestBuilder {
	b.req.InvoiceLimit = &limit

	return b
}

func (b *UpdateRequestBuilder) UpdateExistingSubscriptions(update bool) *UpdateRequestBuilder {
	b.req.UpdateExistingSubscriptions = &update

	return b
}

func (b *UpdateRequestBuilder) Build() *updateRequest {
	return b.req
}

type UpdateResponseData = any
type UpdateResponse = types.Response[any]

func (c *Client) Update(ctx context.Context, idOrCode string, builder UpdateRequestBuilder) (*UpdateResponse, error) {
	return net.Put[updateRequest, UpdateResponseData](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s", basePath, idOrCode), builder.Build(), c.BaseURL)
}
