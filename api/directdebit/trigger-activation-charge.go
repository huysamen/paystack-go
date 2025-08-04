package directdebit

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// TriggerActivationChargeRequest represents the request to trigger activation charge
type TriggerActivationChargeRequest struct {
	CustomerIDs []uint64 `json:"customer_ids"`
}

// TriggerActivationChargeBuilder builds requests for triggering activation charges
type TriggerActivationChargeBuilder struct {
	req *TriggerActivationChargeRequest
}

// NewTriggerActivationChargeBuilder creates a new builder for triggering activation charges
func NewTriggerActivationChargeBuilder() *TriggerActivationChargeBuilder {
	return &TriggerActivationChargeBuilder{
		req: &TriggerActivationChargeRequest{},
	}
}

// CustomerIDs sets the customer IDs for the request
func (b *TriggerActivationChargeBuilder) CustomerIDs(customerIDs []uint64) *TriggerActivationChargeBuilder {
	b.req.CustomerIDs = customerIDs

	return b
}

// CustomerID adds a customer ID to the request
func (b *TriggerActivationChargeBuilder) CustomerID(customerID uint64) *TriggerActivationChargeBuilder {
	b.req.CustomerIDs = append(b.req.CustomerIDs, customerID)

	return b
}

// Build returns the built request
func (b *TriggerActivationChargeBuilder) Build() *TriggerActivationChargeRequest {
	return b.req
}

// TriggerActivationChargeResponse represents the response type for triggering activation charge
type TriggerActivationChargeResponse = types.Response[any]

// TriggerActivationCharge triggers an activation charge on pending mandates
func (c *Client) TriggerActivationCharge(ctx context.Context, builder *TriggerActivationChargeBuilder) (*TriggerActivationChargeResponse, error) {
	return net.Put[TriggerActivationChargeRequest, any](ctx, c.Client, c.Secret, basePath+"/activation-charge", builder.Build(), c.BaseURL)
}
