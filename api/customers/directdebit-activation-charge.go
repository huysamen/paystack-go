package customers

import (
	"context"
	"errors"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type DirectDebitActivationChargeRequest struct {
	AuthorizationID int `json:"authorization_id"`
}

func (r *DirectDebitActivationChargeRequest) Validate() error {
	if r.AuthorizationID <= 0 {
		return errors.New("authorization_id must be a positive integer")
	}
	return nil
}

type DirectDebitActivationChargeResponse struct {
	Message string `json:"message"`
}

func (c *Client) DirectDebitActivationCharge(ctx context.Context, customerID string, req *DirectDebitActivationChargeRequest) (*types.Response[DirectDebitActivationChargeResponse], error) {
	if customerID == "" {
		return nil, errors.New("customer ID is required")
	}

	if req == nil {
		return nil, errors.New("request cannot be nil")
	}

	if err := req.Validate(); err != nil {
		return nil, err
	}

	path := fmt.Sprintf("%s/%s/directdebit-activation-charge", customerBasePath, customerID)

	return net.Put[DirectDebitActivationChargeRequest, DirectDebitActivationChargeResponse](
		ctx,
		c.client,
		c.secret,
		path,
		req,
		c.baseURL,
	)
}
