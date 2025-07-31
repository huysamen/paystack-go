package customers

import (
	"context"
	"errors"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type DeactivateAuthorizationRequest struct {
	AuthorizationCode string `json:"authorization_code"`
}

func (r *DeactivateAuthorizationRequest) Validate() error {
	if r.AuthorizationCode == "" {
		return errors.New("authorization_code is required")
	}
	return nil
}

type DeactivateAuthorizationResponse struct {
	Message string `json:"message"`
}

func (c *Client) DeactivateAuthorization(ctx context.Context, req *DeactivateAuthorizationRequest) (*types.Response[DeactivateAuthorizationResponse], error) {
	if req == nil {
		return nil, errors.New("request cannot be nil")
	}

	if err := req.Validate(); err != nil {
		return nil, err
	}

	path := customerBasePath + "/authorization/deactivate"

	return net.Post[DeactivateAuthorizationRequest, DeactivateAuthorizationResponse](
		ctx,
		c.client,
		c.secret,
		path,
		req,
		c.baseURL,
	)
}
