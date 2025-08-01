package applepay

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// ListDomainsRequest represents the request to list Apple Pay domains
type ListDomainsRequest struct {
	UseCursor *bool   `json:"use_cursor,omitempty"`
	Next      *string `json:"next,omitempty"`
	Previous  *string `json:"previous,omitempty"`
}

// ListDomainsRequestBuilder provides a fluent interface for building ListDomainsRequest
type ListDomainsRequestBuilder struct {
	req *ListDomainsRequest
}

// NewListDomainsRequest creates a new builder for ListDomainsRequest
func NewListDomainsRequest() *ListDomainsRequestBuilder {
	return &ListDomainsRequestBuilder{
		req: &ListDomainsRequest{},
	}
}

// UseCursor sets whether to use cursor-based pagination
func (b *ListDomainsRequestBuilder) UseCursor(useCursor bool) *ListDomainsRequestBuilder {
	b.req.UseCursor = &useCursor
	return b
}

// Next sets the cursor for next page
func (b *ListDomainsRequestBuilder) Next(next string) *ListDomainsRequestBuilder {
	b.req.Next = &next
	return b
}

// Previous sets the cursor for previous page
func (b *ListDomainsRequestBuilder) Previous(previous string) *ListDomainsRequestBuilder {
	b.req.Previous = &previous
	return b
}

// Build returns the constructed ListDomainsRequest
func (b *ListDomainsRequestBuilder) Build() *ListDomainsRequest {
	return b.req
}

// ListDomainsResponse represents the response from listing Apple Pay domains
type ListDomainsResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    struct {
		DomainNames []string `json:"domainNames"`
	} `json:"data"`
	Meta *types.Meta `json:"meta,omitempty"`
}

// ListDomains lists all registered domains on your integration
func (c *Client) ListDomains(ctx context.Context, builder *ListDomainsRequestBuilder) (*ListDomainsResponse, error) {
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
