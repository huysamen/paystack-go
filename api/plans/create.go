package plans

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// CreatePlanRequest represents the request to create a plan
type CreatePlanRequest struct {
	// Required fields
	Name     string         `json:"name"`
	Amount   int            `json:"amount"`
	Interval types.Interval `json:"interval"`

	// Optional fields
	Description  string         `json:"description,omitempty"`
	SendInvoices *bool          `json:"send_invoices,omitempty"`
	SendSMS      *bool          `json:"send_sms,omitempty"`
	Currency     types.Currency `json:"currency,omitempty"`
	InvoiceLimit *int           `json:"invoice_limit,omitempty"`
}

// CreatePlanRequestBuilder provides a fluent interface for building CreatePlanRequest
type CreatePlanRequestBuilder struct {
	req *CreatePlanRequest
}

// NewCreatePlanRequest creates a new builder for CreatePlanRequest
func NewCreatePlanRequest(name string, amount int, interval types.Interval) *CreatePlanRequestBuilder {
	return &CreatePlanRequestBuilder{
		req: &CreatePlanRequest{
			Name:     name,
			Amount:   amount,
			Interval: interval,
		},
	}
}

// Description sets the plan description
func (b *CreatePlanRequestBuilder) Description(description string) *CreatePlanRequestBuilder {
	b.req.Description = description

	return b
}

// SendInvoices sets whether to send invoices to subscribers
func (b *CreatePlanRequestBuilder) SendInvoices(sendInvoices bool) *CreatePlanRequestBuilder {
	b.req.SendInvoices = &sendInvoices

	return b
}

// SendSMS sets whether to send SMS to subscribers
func (b *CreatePlanRequestBuilder) SendSMS(sendSMS bool) *CreatePlanRequestBuilder {
	b.req.SendSMS = &sendSMS

	return b
}

// Currency sets the plan currency
func (b *CreatePlanRequestBuilder) Currency(currency types.Currency) *CreatePlanRequestBuilder {
	b.req.Currency = currency

	return b
}

// InvoiceLimit sets the maximum number of invoices for the plan
func (b *CreatePlanRequestBuilder) InvoiceLimit(limit int) *CreatePlanRequestBuilder {
	b.req.InvoiceLimit = &limit

	return b
}

// Build returns the constructed CreatePlanRequest
func (b *CreatePlanRequestBuilder) Build() *CreatePlanRequest {
	return b.req
}

// CreatePlanResponse represents the response from creating a plan
type CreatePlanResponse = types.Response[types.Plan]

// Create creates a new subscription plan
func (c *Client) Create(ctx context.Context, builder *CreatePlanRequestBuilder) (*CreatePlanResponse, error) {
	return net.Post[CreatePlanRequest, types.Plan](ctx, c.Client, c.Secret, basePath, builder.Build(), c.BaseURL)
}
