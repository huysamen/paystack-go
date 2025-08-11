package directdebit

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type triggerActivationChargeRequest struct {
	CustomerIDs []uint64 `json:"customer_ids"`
}

type TriggerActivationChargeRequestBuilder struct {
	req *triggerActivationChargeRequest
}

func NewTriggerActivationChargeRequestBuilder() *TriggerActivationChargeRequestBuilder {
	return &TriggerActivationChargeRequestBuilder{
		req: &triggerActivationChargeRequest{},
	}
}

func (b *TriggerActivationChargeRequestBuilder) CustomerIDs(customerIDs []uint64) *TriggerActivationChargeRequestBuilder {
	b.req.CustomerIDs = customerIDs

	return b
}

func (b *TriggerActivationChargeRequestBuilder) CustomerID(customerID uint64) *TriggerActivationChargeRequestBuilder {
	b.req.CustomerIDs = append(b.req.CustomerIDs, customerID)

	return b
}

func (b *TriggerActivationChargeRequestBuilder) Build() *triggerActivationChargeRequest {
	return b.req
}

type TriggerActivationChargeResponseData = any
type TriggerActivationChargeResponse = types.Response[TriggerActivationChargeResponseData]

func (c *Client) TriggerActivationCharge(ctx context.Context, builder TriggerActivationChargeRequestBuilder) (*TriggerActivationChargeResponse, error) {
	return net.Put[triggerActivationChargeRequest, TriggerActivationChargeResponseData](ctx, c.Client, c.Secret, basePath+"/activation-charge", builder.Build(), c.BaseURL)
}
