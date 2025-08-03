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

// Request type
type CustomerRiskActionRequest struct {
	Customer   string      `json:"customer"`    // Customer code or email
	RiskAction *RiskAction `json:"risk_action"` // Optional, defaults to default
}

// Builder for CustomerRiskActionRequest
type CustomerRiskActionRequestBuilder struct {
	customer   string
	riskAction *RiskAction
}

// NewSetRiskActionRequest creates a new builder for setting risk action
func NewSetRiskActionRequest(customer string) *CustomerRiskActionRequestBuilder {
	return &CustomerRiskActionRequestBuilder{
		customer: customer,
	}
}

// RiskAction sets the risk action
func (b *CustomerRiskActionRequestBuilder) RiskAction(riskAction RiskAction) *CustomerRiskActionRequestBuilder {
	b.riskAction = &riskAction
	return b
}

// Build creates the CustomerRiskActionRequest
func (b *CustomerRiskActionRequestBuilder) Build() *CustomerRiskActionRequest {
	return &CustomerRiskActionRequest{
		Customer:   b.customer,
		RiskAction: b.riskAction,
	}
}

type SetRiskActionResponse = types.Response[types.Customer]

// SetRiskAction sets the risk action for a customer
func (c *Client) SetRiskAction(ctx context.Context, builder *CustomerRiskActionRequestBuilder) (*SetRiskActionResponse, error) {
	return net.Post[CustomerRiskActionRequest, types.Customer](ctx, c.Client, c.Secret, basePath+"/set_risk_action", builder.Build(), c.BaseURL)
}
