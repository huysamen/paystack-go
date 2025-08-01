package directdebit

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// TriggerActivationCharge triggers an activation charge on pending mandates
func (c *Client) TriggerActivationCharge(ctx context.Context, req *TriggerActivationChargeRequest) (*types.Response[any], error) {
	if err := validateTriggerActivationChargeRequest(req); err != nil {
		return nil, err
	}

	endpoint := directDebitBasePath + "/activation-charge"
	resp, err := net.Put[TriggerActivationChargeRequest, any](
		ctx, c.client, c.secret, endpoint, req, c.baseURL,
	)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
