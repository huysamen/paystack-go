package customers

import (
	"context"
	"errors"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type RiskAction string

const (
	RiskActionDefault RiskAction = "default"
	RiskActionAllow   RiskAction = "allow" // whitelist
	RiskActionDeny    RiskAction = "deny"  // blacklist
)

type CustomerRiskActionRequest struct {
	Customer   string      `json:"customer"`    // Customer code or email
	RiskAction *RiskAction `json:"risk_action"` // Optional, defaults to default
}


func (c *Client) SetRiskAction(ctx context.Context, req *CustomerRiskActionRequest) (*types.Response[Customer], error) {
	if req == nil {
		return nil, errors.New("request cannot be nil")
	}


	path := customerBasePath + "/set_risk_action"

	return net.Post[CustomerRiskActionRequest, Customer](
		ctx,
		c.client,
		c.secret,
		path,
		req,
		c.baseURL,
	)
}
