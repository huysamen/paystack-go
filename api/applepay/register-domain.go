package applepay

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// RegisterDomainRequest represents the request to register an Apple Pay domain
type RegisterDomainRequest struct {
	DomainName string `json:"domainName"`
}

// RegisterDomainRequestBuilder provides a fluent interface for building RegisterDomainRequest
type RegisterDomainRequestBuilder struct {
	req *RegisterDomainRequest
}

// NewRegisterDomainRequest creates a new builder for RegisterDomainRequest
func NewRegisterDomainRequest(domainName string) RegisterDomainRequestBuilder {
	return RegisterDomainRequestBuilder{
		req: &RegisterDomainRequest{
			DomainName: domainName,
		},
	}
}

// Build returns the constructed RegisterDomainRequest
func (b *RegisterDomainRequestBuilder) Build() *RegisterDomainRequest {
	return b.req
}

// RegisterDomainResponse is the response type for registering an Apple Pay domain
type RegisterDomainResponse = types.Response[any]

// RegisterDomain registers a top-level domain or subdomain for Apple Pay integration
func (c *Client) RegisterDomain(ctx context.Context, builder RegisterDomainRequestBuilder) (*RegisterDomainResponse, error) {
	return net.Post[RegisterDomainRequest, any](ctx, c.Client, c.Secret, registerPath, builder.Build(), c.BaseURL)
}
