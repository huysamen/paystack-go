package paymentrequests

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type CreatePaymentRequestRequest struct {
	Customer         string           `json:"customer"`
	Amount           int              `json:"amount"`
	DueDate          string           `json:"due_date,omitempty"`
	Description      string           `json:"description,omitempty"`
	LineItems        []types.LineItem `json:"line_items,omitempty"`
	Tax              []types.Tax      `json:"tax,omitempty"`
	Currency         string           `json:"currency,omitempty"`
	SendNotification *bool            `json:"send_notification,omitempty"`
	Draft            *bool            `json:"draft,omitempty"`
	HasInvoice       *bool            `json:"has_invoice,omitempty"`
	InvoiceNumber    *int             `json:"invoice_number,omitempty"`
	SplitCode        string           `json:"split_code,omitempty"`
	Metadata         types.Metadata   `json:"metadata,omitempty"`
}

type CreatePaymentRequestRequestBuilder struct {
	req *CreatePaymentRequestRequest
}

func NewCreatePaymentRequestRequest() *CreatePaymentRequestRequestBuilder {
	return &CreatePaymentRequestRequestBuilder{
		req: &CreatePaymentRequestRequest{},
	}
}

func (b *CreatePaymentRequestRequestBuilder) Customer(customer string) *CreatePaymentRequestRequestBuilder {
	b.req.Customer = customer

	return b
}

func (b *CreatePaymentRequestRequestBuilder) Amount(amount int) *CreatePaymentRequestRequestBuilder {
	b.req.Amount = amount

	return b
}

func (b *CreatePaymentRequestRequestBuilder) DueDate(dueDate string) *CreatePaymentRequestRequestBuilder {
	b.req.DueDate = dueDate

	return b
}

func (b *CreatePaymentRequestRequestBuilder) Description(description string) *CreatePaymentRequestRequestBuilder {
	b.req.Description = description

	return b
}

func (b *CreatePaymentRequestRequestBuilder) LineItems(lineItems []types.LineItem) *CreatePaymentRequestRequestBuilder {
	b.req.LineItems = lineItems

	return b
}

func (b *CreatePaymentRequestRequestBuilder) AddLineItem(lineItem types.LineItem) *CreatePaymentRequestRequestBuilder {
	b.req.LineItems = append(b.req.LineItems, lineItem)

	return b
}

func (b *CreatePaymentRequestRequestBuilder) Tax(tax []types.Tax) *CreatePaymentRequestRequestBuilder {
	b.req.Tax = tax

	return b
}

func (b *CreatePaymentRequestRequestBuilder) AddTax(tax types.Tax) *CreatePaymentRequestRequestBuilder {
	b.req.Tax = append(b.req.Tax, tax)

	return b
}

func (b *CreatePaymentRequestRequestBuilder) Currency(currency string) *CreatePaymentRequestRequestBuilder {
	b.req.Currency = currency

	return b
}

func (b *CreatePaymentRequestRequestBuilder) SendNotification(sendNotification bool) *CreatePaymentRequestRequestBuilder {
	b.req.SendNotification = &sendNotification

	return b
}

func (b *CreatePaymentRequestRequestBuilder) Draft(draft bool) *CreatePaymentRequestRequestBuilder {
	b.req.Draft = &draft

	return b
}

func (b *CreatePaymentRequestRequestBuilder) HasInvoice(hasInvoice bool) *CreatePaymentRequestRequestBuilder {
	b.req.HasInvoice = &hasInvoice

	return b
}

func (b *CreatePaymentRequestRequestBuilder) InvoiceNumber(invoiceNumber int) *CreatePaymentRequestRequestBuilder {
	b.req.InvoiceNumber = &invoiceNumber

	return b
}

func (b *CreatePaymentRequestRequestBuilder) SplitCode(splitCode string) *CreatePaymentRequestRequestBuilder {
	b.req.SplitCode = splitCode

	return b
}

func (b *CreatePaymentRequestRequestBuilder) Metadata(metadata types.Metadata) *CreatePaymentRequestRequestBuilder {
	b.req.Metadata = metadata

	return b
}

func (b *CreatePaymentRequestRequestBuilder) Build() *CreatePaymentRequestRequest {
	return b.req
}

type CreatePaymentRequestResponse = types.Response[types.PaymentRequest]

func (c *Client) Create(ctx context.Context, builder *CreatePaymentRequestRequestBuilder) (*CreatePaymentRequestResponse, error) {
	return net.Post[CreatePaymentRequestRequest, types.PaymentRequest](ctx, c.Client, c.Secret, basePath, builder.Build(), c.BaseURL)
}
