package directdebit

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type TriggerActivationChargeRequest struct {
	CustomerIDs []uint64 `json:"customer_ids"`
}

type TriggerActivationChargeBuilder struct {
	req *TriggerActivationChargeRequest
}

func NewTriggerActivationChargeBuilder() *TriggerActivationChargeBuilder {
	return &TriggerActivationChargeBuilder{
		req: &TriggerActivationChargeRequest{},
	}
}

func (b *TriggerActivationChargeBuilder) CustomerIDs(customerIDs []uint64) *TriggerActivationChargeBuilder {
	b.req.CustomerIDs = customerIDs

	return b
}

func (b *TriggerActivationChargeBuilder) CustomerID(customerID uint64) *TriggerActivationChargeBuilder {
	b.req.CustomerIDs = append(b.req.CustomerIDs, customerID)

	return b
}

func (b *TriggerActivationChargeBuilder) Build() *TriggerActivationChargeRequest {
	return b.req
}

type TriggerActivationChargeResponse = types.Response[any]

func (c *Client) TriggerActivationCharge(ctx context.Context, builder *TriggerActivationChargeBuilder) (*TriggerActivationChargeResponse, error) {
	return net.Put[TriggerActivationChargeRequest, any](ctx, c.Client, c.Secret, basePath+"/activation-charge", builder.Build(), c.BaseURL)
}
