package plans

import (
	"context"
	"errors"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type PlanCreateRequest struct {
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

type PlanCreateResponse struct {
	types.Plan
}

func (c *Client) Create(ctx context.Context, req *PlanCreateRequest) (*types.Response[PlanCreateResponse], error) {
	if req == nil {
		return nil, errors.New("request cannot be nil")
	}

	return net.Post[PlanCreateRequest, PlanCreateResponse](
		ctx,
		c.client,
		c.secret,
		planBasePath,
		req,
		c.baseURL,
	)
}
