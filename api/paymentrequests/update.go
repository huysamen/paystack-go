package paymentrequests

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type UpdatePaymentRequestRequest struct {
	Customer         string           `json:"customer,omitempty"`
	Amount           *int             `json:"amount,omitempty"`
	DueDate          string           `json:"due_date,omitempty"`
	Description      string           `json:"description,omitempty"`
	LineItems        []types.LineItem `json:"line_items,omitempty"`
	Tax              []types.Tax      `json:"tax,omitempty"`
	Currency         string           `json:"currency,omitempty"`
	SendNotification *bool            `json:"send_notification,omitempty"`
	Draft            *bool            `json:"draft,omitempty"`
	InvoiceNumber    *int             `json:"invoice_number,omitempty"`
	SplitCode        string           `json:"split_code,omitempty"`
	Metadata         types.Metadata   `json:"metadata,omitempty"`
}

type UpdatePaymentRequestRequestBuilder struct {
	req *UpdatePaymentRequestRequest
}

func NewUpdatePaymentRequestRequest() *UpdatePaymentRequestRequestBuilder {
	return &UpdatePaymentRequestRequestBuilder{
		req: &UpdatePaymentRequestRequest{},
	}
}

func (b *UpdatePaymentRequestRequestBuilder) Customer(customer string) *UpdatePaymentRequestRequestBuilder {
	b.req.Customer = customer

	return b
}

func (b *UpdatePaymentRequestRequestBuilder) Amount(amount int) *UpdatePaymentRequestRequestBuilder {
	b.req.Amount = &amount

	return b
}

func (b *UpdatePaymentRequestRequestBuilder) DueDate(dueDate string) *UpdatePaymentRequestRequestBuilder {
	b.req.DueDate = dueDate

	return b
}

func (b *UpdatePaymentRequestRequestBuilder) Description(description string) *UpdatePaymentRequestRequestBuilder {
	b.req.Description = description

	return b
}

func (b *UpdatePaymentRequestRequestBuilder) LineItems(lineItems []types.LineItem) *UpdatePaymentRequestRequestBuilder {
	b.req.LineItems = lineItems

	return b
}

func (b *UpdatePaymentRequestRequestBuilder) AddLineItem(lineItem types.LineItem) *UpdatePaymentRequestRequestBuilder {
	b.req.LineItems = append(b.req.LineItems, lineItem)

	return b
}

func (b *UpdatePaymentRequestRequestBuilder) Tax(tax []types.Tax) *UpdatePaymentRequestRequestBuilder {
	b.req.Tax = tax

	return b
}

func (b *UpdatePaymentRequestRequestBuilder) AddTax(tax types.Tax) *UpdatePaymentRequestRequestBuilder {
	b.req.Tax = append(b.req.Tax, tax)

	return b
}

func (b *UpdatePaymentRequestRequestBuilder) Currency(currency string) *UpdatePaymentRequestRequestBuilder {
	b.req.Currency = currency

	return b
}

func (b *UpdatePaymentRequestRequestBuilder) SendNotification(sendNotification bool) *UpdatePaymentRequestRequestBuilder {
	b.req.SendNotification = &sendNotification

	return b
}

func (b *UpdatePaymentRequestRequestBuilder) Draft(draft bool) *UpdatePaymentRequestRequestBuilder {
	b.req.Draft = &draft

	return b
}

func (b *UpdatePaymentRequestRequestBuilder) InvoiceNumber(invoiceNumber int) *UpdatePaymentRequestRequestBuilder {
	b.req.InvoiceNumber = &invoiceNumber

	return b
}

func (b *UpdatePaymentRequestRequestBuilder) SplitCode(splitCode string) *UpdatePaymentRequestRequestBuilder {
	b.req.SplitCode = splitCode

	return b
}

func (b *UpdatePaymentRequestRequestBuilder) Metadata(metadata types.Metadata) *UpdatePaymentRequestRequestBuilder {
	b.req.Metadata = metadata

	return b
}

func (b *UpdatePaymentRequestRequestBuilder) Build() *UpdatePaymentRequestRequest {
	return b.req
}

type UpdatePaymentRequestResponse = types.Response[types.PaymentRequest]

func (c *Client) Update(ctx context.Context, idOrCode string, builder *UpdatePaymentRequestRequestBuilder) (*UpdatePaymentRequestResponse, error) {
	return net.Put[UpdatePaymentRequestRequest, types.PaymentRequest](ctx, c.Client, c.Secret, basePath+"/"+idOrCode, builder.Build(), c.BaseURL)
}
