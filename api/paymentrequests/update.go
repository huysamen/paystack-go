package paymentrequests

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type updateRequest struct {
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

type UpdateRequestBuilder struct {
	req *updateRequest
}

func NewUpdateRequestBuilder() *UpdateRequestBuilder {
	return &UpdateRequestBuilder{
		req: &updateRequest{},
	}
}

func (b *UpdateRequestBuilder) Customer(customer string) *UpdateRequestBuilder {
	b.req.Customer = customer

	return b
}

func (b *UpdateRequestBuilder) Amount(amount int) *UpdateRequestBuilder {
	b.req.Amount = &amount

	return b
}

func (b *UpdateRequestBuilder) DueDate(dueDate string) *UpdateRequestBuilder {
	b.req.DueDate = dueDate

	return b
}

func (b *UpdateRequestBuilder) Description(description string) *UpdateRequestBuilder {
	b.req.Description = description

	return b
}

func (b *UpdateRequestBuilder) LineItems(lineItems []types.LineItem) *UpdateRequestBuilder {
	b.req.LineItems = lineItems

	return b
}

func (b *UpdateRequestBuilder) AddLineItem(lineItem types.LineItem) *UpdateRequestBuilder {
	b.req.LineItems = append(b.req.LineItems, lineItem)

	return b
}

func (b *UpdateRequestBuilder) Tax(tax []types.Tax) *UpdateRequestBuilder {
	b.req.Tax = tax

	return b
}

func (b *UpdateRequestBuilder) AddTax(tax types.Tax) *UpdateRequestBuilder {
	b.req.Tax = append(b.req.Tax, tax)

	return b
}

func (b *UpdateRequestBuilder) Currency(currency string) *UpdateRequestBuilder {
	b.req.Currency = currency

	return b
}

func (b *UpdateRequestBuilder) SendNotification(sendNotification bool) *UpdateRequestBuilder {
	b.req.SendNotification = &sendNotification

	return b
}

func (b *UpdateRequestBuilder) Draft(draft bool) *UpdateRequestBuilder {
	b.req.Draft = &draft

	return b
}

func (b *UpdateRequestBuilder) InvoiceNumber(invoiceNumber int) *UpdateRequestBuilder {
	b.req.InvoiceNumber = &invoiceNumber

	return b
}

func (b *UpdateRequestBuilder) SplitCode(splitCode string) *UpdateRequestBuilder {
	b.req.SplitCode = splitCode

	return b
}

func (b *UpdateRequestBuilder) Metadata(metadata types.Metadata) *UpdateRequestBuilder {
	b.req.Metadata = metadata

	return b
}

func (b *UpdateRequestBuilder) Build() *updateRequest {
	return b.req
}

type UpdateResponseData = types.PaymentRequest
type UpdateResponse = types.Response[UpdateResponseData]

func (c *Client) Update(ctx context.Context, idOrCode string, builder UpdateRequestBuilder) (*UpdateResponse, error) {
	return net.Put[updateRequest, UpdateResponseData](ctx, c.Client, c.Secret, basePath+"/"+idOrCode, builder.Build(), c.BaseURL)
}
