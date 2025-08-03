package customers

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// Request and Response types
type DeactivateAuthorizationRequest struct {
	AuthorizationCode string `json:"authorization_code"`
}

// Builder for DeactivateAuthorizationRequest
type DeactivateAuthorizationRequestBuilder struct {
	authorizationCode string
}

// NewDeactivateAuthorizationRequest creates a new builder for deactivating authorization
func NewDeactivateAuthorizationRequest(authorizationCode string) *DeactivateAuthorizationRequestBuilder {
	return &DeactivateAuthorizationRequestBuilder{
		authorizationCode: authorizationCode,
	}
}

// Build creates the DeactivateAuthorizationRequest
func (b *DeactivateAuthorizationRequestBuilder) Build() *DeactivateAuthorizationRequest {
	return &DeactivateAuthorizationRequest{
		AuthorizationCode: b.authorizationCode,
	}
}

type DeactivateAuthorizationResponse = types.Response[any]

// DeactivateAuthorization deactivates an authorization with the provided builder
func (c *Client) DeactivateAuthorization(ctx context.Context, builder *DeactivateAuthorizationRequestBuilder) (*DeactivateAuthorizationResponse, error) {
	return net.Post[DeactivateAuthorizationRequest, any](ctx, c.Client, c.Secret, basePath+"/authorization/deactivate", builder.Build(), c.BaseURL)
}
