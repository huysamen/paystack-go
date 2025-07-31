package applepay

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
)

// RegisterDomain registers a top-level domain or subdomain for Apple Pay integration
func (c *Client) RegisterDomain(ctx context.Context, req *RegisterDomainRequest) (*RegisterDomainResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

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
