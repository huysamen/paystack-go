package plans

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type CreatePlanRequest struct {
	Name     string         `json:"name"`
	Amount   int            `json:"amount"`
	Interval types.Interval `json:"interval"`

	Description  string         `json:"description,omitempty"`
	SendInvoices *bool          `json:"send_invoices,omitempty"`
	SendSMS      *bool          `json:"send_sms,omitempty"`
	Currency     types.Currency `json:"currency,omitempty"`
	InvoiceLimit *int           `json:"invoice_limit,omitempty"`
}

type CreatePlanRequestBuilder struct {
	req *CreatePlanRequest
}

func NewCreatePlanRequest(name string, amount int, interval types.Interval) *CreatePlanRequestBuilder {
	return &CreatePlanRequestBuilder{
		req: &CreatePlanRequest{
			Name:     name,
			Amount:   amount,
			Interval: interval,
		},
	}
}

func (b *CreatePlanRequestBuilder) Description(description string) *CreatePlanRequestBuilder {
	b.req.Description = description

	return b
}

func (b *CreatePlanRequestBuilder) SendInvoices(sendInvoices bool) *CreatePlanRequestBuilder {
	b.req.SendInvoices = &sendInvoices

	return b
}

func (b *CreatePlanRequestBuilder) SendSMS(sendSMS bool) *CreatePlanRequestBuilder {
	b.req.SendSMS = &sendSMS

	return b
}

func (b *CreatePlanRequestBuilder) Currency(currency types.Currency) *CreatePlanRequestBuilder {
	b.req.Currency = currency

	return b
}

func (b *CreatePlanRequestBuilder) InvoiceLimit(limit int) *CreatePlanRequestBuilder {
	b.req.InvoiceLimit = &limit

	return b
}

func (b *CreatePlanRequestBuilder) Build() *CreatePlanRequest {
	return b.req
}

type CreatePlanResponse = types.Response[types.Plan]

func (c *Client) Create(ctx context.Context, builder *CreatePlanRequestBuilder) (*CreatePlanResponse, error) {
	return net.Post[CreatePlanRequest, types.Plan](ctx, c.Client, c.Secret, basePath, builder.Build(), c.BaseURL)
}
