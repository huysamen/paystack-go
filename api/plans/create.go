package plans

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type createRequest struct {
	Name     string         `json:"name"`
	Amount   int            `json:"amount"`
	Interval types.Interval `json:"interval"`

	Description  string         `json:"description,omitempty"`
	SendInvoices *bool          `json:"send_invoices,omitempty"`
	SendSMS      *bool          `json:"send_sms,omitempty"`
	Currency     types.Currency `json:"currency,omitempty"`
	InvoiceLimit *int           `json:"invoice_limit,omitempty"`
}

type CreateRequestBuilder struct {
	req *createRequest
}

func NewCreateRequestBuilder(name string, amount int, interval types.Interval) *CreateRequestBuilder {
	return &CreateRequestBuilder{
		req: &createRequest{
			Name:     name,
			Amount:   amount,
			Interval: interval,
		},
	}
}

func (b *CreateRequestBuilder) Description(description string) *CreateRequestBuilder {
	b.req.Description = description

	return b
}

func (b *CreateRequestBuilder) SendInvoices(sendInvoices bool) *CreateRequestBuilder {
	b.req.SendInvoices = &sendInvoices

	return b
}

func (b *CreateRequestBuilder) SendSMS(sendSMS bool) *CreateRequestBuilder {
	b.req.SendSMS = &sendSMS

	return b
}

func (b *CreateRequestBuilder) Currency(currency types.Currency) *CreateRequestBuilder {
	b.req.Currency = currency

	return b
}

func (b *CreateRequestBuilder) InvoiceLimit(limit int) *CreateRequestBuilder {
	b.req.InvoiceLimit = &limit

	return b
}

func (b *CreateRequestBuilder) Build() *createRequest {
	return b.req
}

type CreateResponseData = types.Plan
type CreateResponse = types.Response[CreateResponseData]

func (c *Client) Create(ctx context.Context, builder CreateRequestBuilder) (*CreateResponse, error) {
	return net.Post[createRequest, CreateResponseData](ctx, c.Client, c.Secret, basePath, builder.Build(), c.BaseURL)
}
