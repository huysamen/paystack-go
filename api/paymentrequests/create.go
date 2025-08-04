package paymentrequests

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type createRequest struct {
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

type CreateRequestBuilder struct {
	req *createRequest
}

func NewCreateRequestBuilder() *CreateRequestBuilder {
	return &CreateRequestBuilder{
		req: &createRequest{},
	}
}

func (b *CreateRequestBuilder) Customer(customer string) *CreateRequestBuilder {
	b.req.Customer = customer

	return b
}

func (b *CreateRequestBuilder) Amount(amount int) *CreateRequestBuilder {
	b.req.Amount = amount

	return b
}

func (b *CreateRequestBuilder) DueDate(dueDate string) *CreateRequestBuilder {
	b.req.DueDate = dueDate

	return b
}

func (b *CreateRequestBuilder) Description(description string) *CreateRequestBuilder {
	b.req.Description = description

	return b
}

func (b *CreateRequestBuilder) LineItems(lineItems []types.LineItem) *CreateRequestBuilder {
	b.req.LineItems = lineItems

	return b
}

func (b *CreateRequestBuilder) AddLineItem(lineItem types.LineItem) *CreateRequestBuilder {
	b.req.LineItems = append(b.req.LineItems, lineItem)

	return b
}

func (b *CreateRequestBuilder) Tax(tax []types.Tax) *CreateRequestBuilder {
	b.req.Tax = tax

	return b
}

func (b *CreateRequestBuilder) AddTax(tax types.Tax) *CreateRequestBuilder {
	b.req.Tax = append(b.req.Tax, tax)

	return b
}

func (b *CreateRequestBuilder) Currency(currency string) *CreateRequestBuilder {
	b.req.Currency = currency

	return b
}

func (b *CreateRequestBuilder) SendNotification(sendNotification bool) *CreateRequestBuilder {
	b.req.SendNotification = &sendNotification

	return b
}

func (b *CreateRequestBuilder) Draft(draft bool) *CreateRequestBuilder {
	b.req.Draft = &draft

	return b
}

func (b *CreateRequestBuilder) HasInvoice(hasInvoice bool) *CreateRequestBuilder {
	b.req.HasInvoice = &hasInvoice

	return b
}

func (b *CreateRequestBuilder) InvoiceNumber(invoiceNumber int) *CreateRequestBuilder {
	b.req.InvoiceNumber = &invoiceNumber

	return b
}

func (b *CreateRequestBuilder) SplitCode(splitCode string) *CreateRequestBuilder {
	b.req.SplitCode = splitCode

	return b
}

func (b *CreateRequestBuilder) Metadata(metadata types.Metadata) *CreateRequestBuilder {
	b.req.Metadata = metadata

	return b
}

func (b *CreateRequestBuilder) Build() *createRequest {
	return b.req
}

type CreateResponseData = types.PaymentRequest
type CreateResponse = types.Response[CreateResponseData]

func (c *Client) Create(ctx context.Context, builder CreateRequestBuilder) (*CreateResponse, error) {
	return net.Post[createRequest, CreateResponseData](ctx, c.Client, c.Secret, basePath, builder.Build(), c.BaseURL)
}
