package paymentrequests

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// CreatePaymentRequestRequest represents the request to create a payment request
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

// CreatePaymentRequestRequestBuilder provides a fluent interface for building CreatePaymentRequestRequest
type CreatePaymentRequestRequestBuilder struct {
	req *CreatePaymentRequestRequest
}

// NewCreatePaymentRequestRequest creates a new builder for CreatePaymentRequestRequest
func NewCreatePaymentRequestRequest() *CreatePaymentRequestRequestBuilder {
	return &CreatePaymentRequestRequestBuilder{
		req: &CreatePaymentRequestRequest{},
	}
}

// Customer sets the customer (required)
func (b *CreatePaymentRequestRequestBuilder) Customer(customer string) *CreatePaymentRequestRequestBuilder {
	b.req.Customer = customer

	return b
}

// Amount sets the payment amount (required)
func (b *CreatePaymentRequestRequestBuilder) Amount(amount int) *CreatePaymentRequestRequestBuilder {
	b.req.Amount = amount

	return b
}

// DueDate sets the due date
func (b *CreatePaymentRequestRequestBuilder) DueDate(dueDate string) *CreatePaymentRequestRequestBuilder {
	b.req.DueDate = dueDate

	return b
}

// Description sets the payment description
func (b *CreatePaymentRequestRequestBuilder) Description(description string) *CreatePaymentRequestRequestBuilder {
	b.req.Description = description

	return b
}

// LineItems sets the line items
func (b *CreatePaymentRequestRequestBuilder) LineItems(lineItems []types.LineItem) *CreatePaymentRequestRequestBuilder {
	b.req.LineItems = lineItems

	return b
}

// AddLineItem adds a single line item
func (b *CreatePaymentRequestRequestBuilder) AddLineItem(lineItem types.LineItem) *CreatePaymentRequestRequestBuilder {
	b.req.LineItems = append(b.req.LineItems, lineItem)

	return b
}

// Tax sets the tax information
func (b *CreatePaymentRequestRequestBuilder) Tax(tax []types.Tax) *CreatePaymentRequestRequestBuilder {
	b.req.Tax = tax

	return b
}

// AddTax adds a single tax entry
func (b *CreatePaymentRequestRequestBuilder) AddTax(tax types.Tax) *CreatePaymentRequestRequestBuilder {
	b.req.Tax = append(b.req.Tax, tax)

	return b
}

// Currency sets the currency
func (b *CreatePaymentRequestRequestBuilder) Currency(currency string) *CreatePaymentRequestRequestBuilder {
	b.req.Currency = currency

	return b
}

// SendNotification sets whether to send notification
func (b *CreatePaymentRequestRequestBuilder) SendNotification(sendNotification bool) *CreatePaymentRequestRequestBuilder {
	b.req.SendNotification = &sendNotification

	return b
}

// Draft sets whether this is a draft
func (b *CreatePaymentRequestRequestBuilder) Draft(draft bool) *CreatePaymentRequestRequestBuilder {
	b.req.Draft = &draft

	return b
}

// HasInvoice sets whether this has an invoice
func (b *CreatePaymentRequestRequestBuilder) HasInvoice(hasInvoice bool) *CreatePaymentRequestRequestBuilder {
	b.req.HasInvoice = &hasInvoice

	return b
}

// InvoiceNumber sets the invoice number
func (b *CreatePaymentRequestRequestBuilder) InvoiceNumber(invoiceNumber int) *CreatePaymentRequestRequestBuilder {
	b.req.InvoiceNumber = &invoiceNumber

	return b
}

// SplitCode sets the split code
func (b *CreatePaymentRequestRequestBuilder) SplitCode(splitCode string) *CreatePaymentRequestRequestBuilder {
	b.req.SplitCode = splitCode

	return b
}

// Metadata sets the metadata
func (b *CreatePaymentRequestRequestBuilder) Metadata(metadata types.Metadata) *CreatePaymentRequestRequestBuilder {
	b.req.Metadata = metadata

	return b
}

// Build returns the constructed CreatePaymentRequestRequest
func (b *CreatePaymentRequestRequestBuilder) Build() *CreatePaymentRequestRequest {
	return b.req
}

// CreatePaymentRequestResponse represents the response from creating a payment request
type CreatePaymentRequestResponse = types.Response[types.PaymentRequest]

// Create creates a payment request for a transaction on your integration
func (c *Client) Create(ctx context.Context, builder *CreatePaymentRequestRequestBuilder) (*CreatePaymentRequestResponse, error) {
	return net.Post[CreatePaymentRequestRequest, types.PaymentRequest](ctx, c.Client, c.Secret, basePath, builder.Build(), c.BaseURL)
}
