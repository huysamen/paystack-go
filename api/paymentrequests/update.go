package paymentrequests

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// UpdatePaymentRequestRequest represents the request to update a payment request
type UpdatePaymentRequestRequest struct {
	Customer         string         `json:"customer,omitempty"`
	Amount           *int           `json:"amount,omitempty"`
	DueDate          string         `json:"due_date,omitempty"`
	Description      string         `json:"description,omitempty"`
	LineItems        []LineItem     `json:"line_items,omitempty"`
	Tax              []Tax          `json:"tax,omitempty"`
	Currency         string         `json:"currency,omitempty"`
	SendNotification *bool          `json:"send_notification,omitempty"`
	Draft            *bool          `json:"draft,omitempty"`
	InvoiceNumber    *int           `json:"invoice_number,omitempty"`
	SplitCode        string         `json:"split_code,omitempty"`
	Metadata         types.Metadata `json:"metadata,omitempty"`
}

// UpdatePaymentRequestRequestBuilder provides a fluent interface for building UpdatePaymentRequestRequest
type UpdatePaymentRequestRequestBuilder struct {
	req *UpdatePaymentRequestRequest
}

// NewUpdatePaymentRequestRequest creates a new builder for UpdatePaymentRequestRequest
func NewUpdatePaymentRequestRequest() *UpdatePaymentRequestRequestBuilder {
	return &UpdatePaymentRequestRequestBuilder{
		req: &UpdatePaymentRequestRequest{},
	}
}

// Customer sets the customer
func (b *UpdatePaymentRequestRequestBuilder) Customer(customer string) *UpdatePaymentRequestRequestBuilder {
	b.req.Customer = customer
	return b
}

// Amount sets the payment amount
func (b *UpdatePaymentRequestRequestBuilder) Amount(amount int) *UpdatePaymentRequestRequestBuilder {
	b.req.Amount = &amount
	return b
}

// DueDate sets the due date
func (b *UpdatePaymentRequestRequestBuilder) DueDate(dueDate string) *UpdatePaymentRequestRequestBuilder {
	b.req.DueDate = dueDate
	return b
}

// Description sets the payment description
func (b *UpdatePaymentRequestRequestBuilder) Description(description string) *UpdatePaymentRequestRequestBuilder {
	b.req.Description = description
	return b
}

// LineItems sets the line items
func (b *UpdatePaymentRequestRequestBuilder) LineItems(lineItems []LineItem) *UpdatePaymentRequestRequestBuilder {
	b.req.LineItems = lineItems
	return b
}

// AddLineItem adds a single line item
func (b *UpdatePaymentRequestRequestBuilder) AddLineItem(lineItem LineItem) *UpdatePaymentRequestRequestBuilder {
	b.req.LineItems = append(b.req.LineItems, lineItem)
	return b
}

// Tax sets the tax information
func (b *UpdatePaymentRequestRequestBuilder) Tax(tax []Tax) *UpdatePaymentRequestRequestBuilder {
	b.req.Tax = tax
	return b
}

// AddTax adds a single tax entry
func (b *UpdatePaymentRequestRequestBuilder) AddTax(tax Tax) *UpdatePaymentRequestRequestBuilder {
	b.req.Tax = append(b.req.Tax, tax)
	return b
}

// Currency sets the currency
func (b *UpdatePaymentRequestRequestBuilder) Currency(currency string) *UpdatePaymentRequestRequestBuilder {
	b.req.Currency = currency
	return b
}

// SendNotification sets whether to send notification
func (b *UpdatePaymentRequestRequestBuilder) SendNotification(sendNotification bool) *UpdatePaymentRequestRequestBuilder {
	b.req.SendNotification = &sendNotification
	return b
}

// Draft sets whether this is a draft
func (b *UpdatePaymentRequestRequestBuilder) Draft(draft bool) *UpdatePaymentRequestRequestBuilder {
	b.req.Draft = &draft
	return b
}

// InvoiceNumber sets the invoice number
func (b *UpdatePaymentRequestRequestBuilder) InvoiceNumber(invoiceNumber int) *UpdatePaymentRequestRequestBuilder {
	b.req.InvoiceNumber = &invoiceNumber
	return b
}

// SplitCode sets the split code
func (b *UpdatePaymentRequestRequestBuilder) SplitCode(splitCode string) *UpdatePaymentRequestRequestBuilder {
	b.req.SplitCode = splitCode
	return b
}

// Metadata sets the metadata
func (b *UpdatePaymentRequestRequestBuilder) Metadata(metadata types.Metadata) *UpdatePaymentRequestRequestBuilder {
	b.req.Metadata = metadata
	return b
}

// Build returns the constructed UpdatePaymentRequestRequest
func (b *UpdatePaymentRequestRequestBuilder) Build() *UpdatePaymentRequestRequest {
	return b.req
}

// UpdatePaymentRequestResponse represents the response from updating a payment request
type UpdatePaymentRequestResponse = types.Response[PaymentRequest]

// Update updates a payment request details on your integration
func (c *Client) Update(ctx context.Context, idOrCode string, builder *UpdatePaymentRequestRequestBuilder) (*UpdatePaymentRequestResponse, error) {
	return net.Put[UpdatePaymentRequestRequest, PaymentRequest](ctx, c.Client, c.Secret, basePath+"/"+idOrCode, builder.Build(), c.BaseURL)
}
