package bulkcharges

import (
	"context"
	"net/url"
	"strconv"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type ListRequest struct {
	PerPage *int    `json:"perPage,omitempty"`
	Page    *int    `json:"page,omitempty"`
	From    *string `json:"from,omitempty"`
	To      *string `json:"to,omitempty"`
}

type ListRequestBuilder struct {
	req *ListRequest
}

func NewListRequest() *ListRequestBuilder {
	return &ListRequestBuilder{
		req: &ListRequest{},
	}
}

func (b *ListRequestBuilder) PerPage(perPage int) *ListRequestBuilder {
	b.req.PerPage = &perPage

	return b
}

func (b *ListRequestBuilder) Page(page int) *ListRequestBuilder {
	b.req.Page = &page

	return b
}

func (b *ListRequestBuilder) From(from string) *ListRequestBuilder {
	b.req.From = &from

	return b
}

func (b *ListRequestBuilder) To(to string) *ListRequestBuilder {
	b.req.To = &to

	return b
}

func (b *ListRequestBuilder) DateRange(from, to string) *ListRequestBuilder {
	b.req.From = &from
	b.req.To = &to

	return b
}

func (b *ListRequestBuilder) Build() *ListRequest {
	return b.req
}

func (r *ListRequest) toQuery() string {
	params := url.Values{}

	if r.PerPage != nil {
		params.Set("perPage", strconv.Itoa(*r.PerPage))
	}

	if r.Page != nil {
		params.Set("page", strconv.Itoa(*r.Page))
	}

	if r.From != nil {
		params.Set("from", *r.From)
	}

	if r.To != nil {
		params.Set("to", *r.To)
	}

	return params.Encode()
}

type ListResponseData = []types.BulkChargeBatch
type ListResponse = types.Response[ListResponseData]

func (c *Client) List(ctx context.Context, builder ListRequestBuilder) (*ListResponse, error) {
	req := builder.Build()
	path := basePath

	if req != nil {
		if query := req.toQuery(); query != "" {
			path += "?" + query
		}
	}

	return net.Get[ListResponseData](ctx, c.Client, c.Secret, path, c.BaseURL)
}
