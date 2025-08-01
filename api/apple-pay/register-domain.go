package applepay

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
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
func NewRegisterDomainRequest(domainName string) *RegisterDomainRequestBuilder {
	return &RegisterDomainRequestBuilder{
		req: &RegisterDomainRequest{
			DomainName: domainName,
		},
	}
}

// Build returns the constructed RegisterDomainRequest
func (b *RegisterDomainRequestBuilder) Build() *RegisterDomainRequest {
	return b.req
}

// RegisterDomainResponse represents the response from registering an Apple Pay domain
type RegisterDomainResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

// RegisterDomain registers a top-level domain or subdomain for Apple Pay integration
func (c *Client) RegisterDomain(ctx context.Context, builder *RegisterDomainRequestBuilder) (*RegisterDomainResponse, error) {
	if builder == nil {
		return nil, fmt.Errorf("builder cannot be nil")
	}

	req := builder.Build()
	if req.DomainName == "" {
		return nil, fmt.Errorf("domainName is required")
	}

	resp, err := net.Post[RegisterDomainRequest, RegisterDomainResponse](
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

	return &resp.Data, nil
}
