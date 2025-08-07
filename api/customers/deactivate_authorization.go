package customers

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type deactivateAuthorizationRequest struct {
	AuthorizationCode string `json:"authorization_code"`
}

type DeactivateAuthorizationRequestBuilder struct {
	req *deactivateAuthorizationRequest
}

func NewDeactivateAuthorizationRequestBuilder(authorizationCode string) *DeactivateAuthorizationRequestBuilder {
	return &DeactivateAuthorizationRequestBuilder{
		req: &deactivateAuthorizationRequest{
			AuthorizationCode: authorizationCode,
		},
	}
}

func (b *DeactivateAuthorizationRequestBuilder) Build() *deactivateAuthorizationRequest {
	return b.req
}

type DeactivateAuthorizationResponseData = any
type DeactivateAuthorizationResponse = types.Response[DeactivateAuthorizationResponseData]

func (c *Client) DeactivateAuthorization(ctx context.Context, builder DeactivateAuthorizationRequestBuilder) (*DeactivateAuthorizationResponse, error) {
	return net.Post[deactivateAuthorizationRequest, DeactivateAuthorizationResponseData](ctx, c.Client, c.Secret, basePath+"/authorization/deactivate", builder.Build(), c.BaseURL)
}
