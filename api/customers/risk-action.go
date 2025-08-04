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

type CustomerRiskActionRequest struct {
	Customer   string      `json:"customer"`    // Customer code or email
	RiskAction *RiskAction `json:"risk_action"` // Optional, defaults to default
}

type CustomerRiskActionRequestBuilder struct {
	req *CustomerRiskActionRequest
}

func NewSetRiskActionRequest(customer string) *CustomerRiskActionRequestBuilder {
	return &CustomerRiskActionRequestBuilder{
		req: &CustomerRiskActionRequest{
			Customer: customer,
		},
	}
}

func (b *CustomerRiskActionRequestBuilder) RiskAction(riskAction RiskAction) *CustomerRiskActionRequestBuilder {
	b.req.RiskAction = &riskAction
	return b
}

func (b *CustomerRiskActionRequestBuilder) Build() *CustomerRiskActionRequest {
	return b.req
}

type SetRiskActionResponse = types.Response[types.Customer]

func (c *Client) SetRiskAction(ctx context.Context, builder *CustomerRiskActionRequestBuilder) (*SetRiskActionResponse, error) {
	return net.Post[CustomerRiskActionRequest, types.Customer](ctx, c.Client, c.Secret, basePath+"/set_risk_action", builder.Build(), c.BaseURL)
}
