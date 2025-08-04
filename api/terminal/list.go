package terminal

import (
	"context"
	"net/url"
	"strconv"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type listRequest struct {
	PerPage  *int    `json:"perPage,omitempty"`  // Number of terminals per page (default: 50)
	Next     *string `json:"next,omitempty"`     // Cursor for next page (optional)
	Previous *string `json:"previous,omitempty"` // Cursor for previous page (optional)
}

type ListRequestBuilder struct {
	req *listRequest
}

func NewListRequestBuilder() *ListRequestBuilder {
	return &ListRequestBuilder{
		req: &listRequest{},
	}
}

func (b *ListRequestBuilder) PerPage(perPage int) *ListRequestBuilder {
	b.req.PerPage = &perPage

	return b
}

func (b *ListRequestBuilder) Next(next string) *ListRequestBuilder {
	b.req.Next = &next

	return b
}

func (b *ListRequestBuilder) Previous(previous string) *ListRequestBuilder {
	b.req.Previous = &previous

	return b
}

func (b *ListRequestBuilder) Build() *listRequest {
	return b.req
}

func (r *listRequest) toQuery() string {
	params := url.Values{}
	if r.PerPage != nil {
		params.Set("perPage", strconv.Itoa(*r.PerPage))
	}
	if r.Next != nil {
		params.Set("next", *r.Next)
	}
	if r.Previous != nil {
		params.Set("previous", *r.Previous)
	}

	return params.Encode()
}

type ListResponseData = []types.Terminal
type ListResponse = types.Response[ListResponseData]

func (c *Client) List(ctx context.Context, builder ListRequestBuilder) (*ListResponse, error) {
	req := builder.Build()
	path := basePath

	if req != nil {
		if query := req.toQuery(); query != "" {
			path += "?" + query
		}
	}

	return net.Get[[]types.Terminal](ctx, c.Client, c.Secret, path, c.BaseURL)
}
