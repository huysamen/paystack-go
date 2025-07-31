package plans

import (
	"context"
	"errors"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type PlanUpdateRequest struct {
	// Required fields
	Name     string         `json:"name" validate:"required"`
	Amount   int            `json:"amount" validate:"required,min=1"`
	Interval types.Interval `json:"interval" validate:"required"`

	// Optional fields
	Description                 string         `json:"description,omitempty"`
	SendInvoices                *bool          `json:"send_invoices,omitempty"`
	SendSMS                     *bool          `json:"send_sms,omitempty"`
	Currency                    types.Currency `json:"currency,omitempty"`
	InvoiceLimit                *int           `json:"invoice_limit,omitempty"`
	UpdateExistingSubscriptions *bool          `json:"update_existing_subscriptions,omitempty"`
}

// Validate checks if the request has all required fields
func (r *PlanUpdateRequest) Validate() error {
	if r.Name == "" {
		return errors.New("name is required")
	}
	if r.Amount <= 0 {
		return errors.New("amount is required and must be greater than 0")
	}
	if r.Interval == types.IntervalUnknown {
		return errors.New("interval is required")
	}
	return nil
}

type PlanUpdateResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

func (c *Client) Update(ctx context.Context, idOrCode string, req *PlanUpdateRequest) (*types.Response[PlanUpdateResponse], error) {
	if idOrCode == "" {
		return nil, errors.New("plan ID or code is required")
	}

	if req == nil {
		return nil, errors.New("request cannot be nil")
	}

	if err := req.Validate(); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	return net.Put[PlanUpdateRequest, PlanUpdateResponse](
		ctx,
		c.client,
		c.secret,
		fmt.Sprintf("%s/%s", planBasePath, idOrCode),
		req,
		c.baseURL,
	)
}
