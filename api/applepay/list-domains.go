package applepay

import (
	"context"
	"net/url"
	"strconv"

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
type ListDomainsResponseData struct {
	DomainNames []string `json:"domainNames"`
}

// ListDomains lists all registered domains on your integration
func (c *Client) ListDomains(ctx context.Context, builder *ListDomainsRequestBuilder) (*types.Response[ListDomainsResponseData], error) {
	req := builder.Build()

	params := url.Values{}
	if req.UseCursor != nil {
		params.Set("use_cursor", strconv.FormatBool(*req.UseCursor))
	}
	if req.Next != nil {
		params.Set("next", *req.Next)
	}
	if req.Previous != nil {
		params.Set("previous", *req.Previous)
	}

	path := listPath
	if len(params) > 0 {
		path += "?" + params.Encode()
	}

	return net.Get[ListDomainsResponseData](ctx, c.Client, c.Secret, path, c.BaseURL)
}
