package applepay

import (
	"context"
	"net/url"
	"strconv"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
	"github.com/huysamen/paystack-go/types/data"
)

type listDomainsRequest struct {
	UseCursor *bool   `json:"use_cursor,omitempty"`
	Next      *string `json:"next,omitempty"`
	Previous  *string `json:"previous,omitempty"`
}

type ListDomainsRequestBuilder struct {
	req *listDomainsRequest
}

func NewListDomainsRequest() *ListDomainsRequestBuilder {
	return &ListDomainsRequestBuilder{
		req: &listDomainsRequest{},
	}
}

func (b *ListDomainsRequestBuilder) UseCursor(useCursor bool) *ListDomainsRequestBuilder {
	b.req.UseCursor = &useCursor

	return b
}

func (b *ListDomainsRequestBuilder) Next(next string) *ListDomainsRequestBuilder {
	b.req.Next = &next

	return b
}

func (b *ListDomainsRequestBuilder) Previous(previous string) *ListDomainsRequestBuilder {
	b.req.Previous = &previous

	return b
}

func (b *ListDomainsRequestBuilder) Build() *listDomainsRequest {
	return b.req
}

func (r *listDomainsRequest) toQuery() string {
	params := url.Values{}

	if r.UseCursor != nil {
		params.Set("use_cursor", strconv.FormatBool(*r.UseCursor))
	}

	if r.Next != nil {
		params.Set("next", *r.Next)
	}

	if r.Previous != nil {
		params.Set("previous", *r.Previous)
	}

	return params.Encode()
}

type ListDomainsResponseData struct {
	DomainNames []data.String `json:"domainNames"`
}

type ListDomainsResponse = types.Response[ListDomainsResponseData]

func (c *Client) ListDomains(ctx context.Context, builder ListDomainsRequestBuilder) (*ListDomainsResponse, error) {
	req := builder.Build()
	path := listPath

	if req != nil {
		if query := req.toQuery(); query != "" {
			path += "?" + query
		}
	}

	return net.Get[ListDomainsResponseData](ctx, c.Client, c.Secret, path, c.BaseURL)
}
