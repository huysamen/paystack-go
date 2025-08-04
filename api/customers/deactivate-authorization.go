package customers

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type DeactivateAuthorizationRequest struct {
	AuthorizationCode string `json:"authorization_code"`
}

type DeactivateAuthorizationRequestBuilder struct {
	req *DeactivateAuthorizationRequest
}

func NewDeactivateAuthorizationRequest(authorizationCode string) *DeactivateAuthorizationRequestBuilder {
	return &DeactivateAuthorizationRequestBuilder{
		req: &DeactivateAuthorizationRequest{
			AuthorizationCode: authorizationCode,
		},
	}
}

func (b *DeactivateAuthorizationRequestBuilder) Build() *DeactivateAuthorizationRequest {
	return b.req
}

type DeactivateAuthorizationResponse = types.Response[any]

func (c *Client) DeactivateAuthorization(ctx context.Context, builder *DeactivateAuthorizationRequestBuilder) (*DeactivateAuthorizationResponse, error) {
	return net.Post[DeactivateAuthorizationRequest, any](ctx, c.Client, c.Secret, basePath+"/authorization/deactivate", builder.Build(), c.BaseURL)
}
