package customers

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type RiskAction string

const (
	RiskActionDefault RiskAction = "default"
	RiskActionAllow   RiskAction = "allow"
	RiskActionDeny    RiskAction = "deny"
)

type RiskActionRequest struct {
	Customer   string      `json:"customer"`    // Customer code or email
	RiskAction *RiskAction `json:"risk_action"` // Optional, defaults to default
}

type RiskActionRequestBuilder struct {
	req *RiskActionRequest
}

func NewRiskActionRequest(customer string) *RiskActionRequestBuilder {
	return &RiskActionRequestBuilder{
		req: &RiskActionRequest{
			Customer: customer,
		},
	}
}

func (b *RiskActionRequestBuilder) RiskAction(riskAction RiskAction) *RiskActionRequestBuilder {
	b.req.RiskAction = &riskAction
	return b
}

func (b *RiskActionRequestBuilder) Build() *RiskActionRequest {
	return b.req
}

type RiskActionResponseData = types.Customer
type RiskActionResponse = types.Response[RiskActionResponseData]

func (c *Client) SetRiskAction(ctx context.Context, builder RiskActionRequestBuilder) (*RiskActionResponse, error) {
	return net.Post[RiskActionRequest, RiskActionResponseData](ctx, c.Client, c.Secret, basePath+"/set_risk_action", builder.Build(), c.BaseURL)
}
