package applepay

import (
	"context"

	"github.com/huysamen/paystack-go/net"
)

// ListDomains lists all registered domains on your integration
func (c *Client) ListDomains(ctx context.Context, req *ListDomainsRequest) (*ListDomainsResponse, error) {
	if req == nil {
		req = &ListDomainsRequest{}
	}

	resp, err := net.Get[ListDomainsResponse](
		ctx,
		c.client,
		c.secret,
		applePayBasePath+"/domain",
		c.baseURL,
	)
	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}

// ListDomainsWithBuilder lists all registered domains using the builder pattern
func (c *Client) ListDomainsWithBuilder(ctx context.Context, builder *ListDomainsRequestBuilder) (*ListDomainsResponse, error) {
	if builder == nil {
		return c.ListDomains(ctx, nil)
	}
	return c.ListDomains(ctx, builder.Build())
}
