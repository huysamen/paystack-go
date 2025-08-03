package applepay

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// UnregisterDomainRequest represents the request to unregister an Apple Pay domain
type UnregisterDomainRequest struct {
	DomainName string `json:"domainName"`
}

// UnregisterDomainRequestBuilder provides a fluent interface for building UnregisterDomainRequest
type UnregisterDomainRequestBuilder struct {
	req *UnregisterDomainRequest
}

// NewUnregisterDomainRequest creates a new builder for UnregisterDomainRequest
func NewUnregisterDomainRequest(domainName string) UnregisterDomainRequestBuilder {
	return UnregisterDomainRequestBuilder{
		req: &UnregisterDomainRequest{
			DomainName: domainName,
		},
	}
}

// Build returns the constructed UnregisterDomainRequest
func (b *UnregisterDomainRequestBuilder) Build() *UnregisterDomainRequest {
	return b.req
}

type UnregisterDomainResponse = types.Response[any]

// UnregisterDomain unregisters a top-level domain or subdomain previously used for Apple Pay integration
func (c *Client) UnregisterDomain(ctx context.Context, builder UnregisterDomainRequestBuilder) (*UnregisterDomainResponse, error) {
	return net.DeleteWithBody[UnregisterDomainRequest, any](ctx, c.Client, c.Secret, unregisterPath, builder.Build(), c.BaseURL)
}
