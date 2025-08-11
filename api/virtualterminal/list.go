package virtualterminal

import (
	"context"
	"net/url"
	"strconv"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type listRequest struct {
	Status   string `json:"status,omitempty"`
	PerPage  int    `json:"perPage,omitempty"`
	Search   string `json:"search,omitempty"`
	Next     string `json:"next,omitempty"`
	Previous string `json:"previous,omitempty"`
}

type ListRequestBuilder struct {
	req *listRequest
}

func NewListRequestBuilder() *ListRequestBuilder {
	return &ListRequestBuilder{
		req: &listRequest{},
	}
}

func (b *ListRequestBuilder) Status(status string) *ListRequestBuilder {
	b.req.Status = status

	return b
}

func (b *ListRequestBuilder) PerPage(perPage int) *ListRequestBuilder {
	b.req.PerPage = perPage

	return b
}

func (b *ListRequestBuilder) Search(search string) *ListRequestBuilder {
	b.req.Search = search

	return b
}

func (b *ListRequestBuilder) Next(next string) *ListRequestBuilder {
	b.req.Next = next

	return b
}

func (b *ListRequestBuilder) Previous(previous string) *ListRequestBuilder {
	b.req.Previous = previous

	return b
}

func (b *ListRequestBuilder) Build() *listRequest {
	return b.req
}

func (r *listRequest) toQuery() string {
	params := url.Values{}
	if r.Status != "" {
		params.Set("status", r.Status)
	}
	if r.PerPage > 0 {
		params.Set("perPage", strconv.Itoa(r.PerPage))
	}
	if r.Search != "" {
		params.Set("search", r.Search)
	}
	if r.Next != "" {
		params.Set("next", r.Next)
	}
	if r.Previous != "" {
		params.Set("previous", r.Previous)
	}

	return params.Encode()
}

type ListResponseData = []types.VirtualTerminal
type ListResponse = types.Response[ListResponseData]

func (c *Client) List(ctx context.Context, builder ListRequestBuilder) (*ListResponse, error) {
	req := builder.Build()
	path := basePath

	if req != nil {
		if query := req.toQuery(); query != "" {
			path += "?" + query
		}
	}

	return net.Get[[]types.VirtualTerminal](ctx, c.Client, c.Secret, path, c.BaseURL)
}
