package applepay

import (
	"context"
	"fmt"

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
func NewUnregisterDomainRequest(domainName string) *UnregisterDomainRequestBuilder {
	return &UnregisterDomainRequestBuilder{
		req: &UnregisterDomainRequest{
			DomainName: domainName,
		},
	}
}

// Build returns the constructed UnregisterDomainRequest
func (b *UnregisterDomainRequestBuilder) Build() *UnregisterDomainRequest {
	return b.req
}

// UnregisterDomainResponse represents the response from unregistering an Apple Pay domain
type UnregisterDomainResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

// UnregisterDomain unregisters a top-level domain or subdomain previously used for Apple Pay integration
func (c *Client) UnregisterDomain(ctx context.Context, builder *UnregisterDomainRequestBuilder) (*types.Response[UnregisterDomainResponse], error) {
	if builder == nil {
		return nil, fmt.Errorf("builder cannot be nil")
	}

	req := builder.Build()
	if req.DomainName == "" {
		return nil, fmt.Errorf("domainName is required")
	}

	resp, err := net.DeleteWithBody[UnregisterDomainRequest, UnregisterDomainResponse](
		ctx,
		c.client,
		c.secret,
		applePayBasePath+"/domain",
		req,
		c.baseURL,
	)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
