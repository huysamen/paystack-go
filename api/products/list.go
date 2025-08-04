package products

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type ListProductsRequest struct {
	PerPage *int    `json:"perPage,omitempty"`
	Page    *int    `json:"page,omitempty"`
	From    *string `json:"from,omitempty"`
	To      *string `json:"to,omitempty"`
}

type ListProductsRequestBuilder struct {
	req *ListProductsRequest
}

func NewListProductsRequest() *ListProductsRequestBuilder {
	return &ListProductsRequestBuilder{
		req: &ListProductsRequest{},
	}
}

func (b *ListProductsRequestBuilder) PerPage(perPage int) *ListProductsRequestBuilder {
	b.req.PerPage = &perPage

	return b
}

func (b *ListProductsRequestBuilder) Page(page int) *ListProductsRequestBuilder {
	b.req.Page = &page

	return b
}

func (b *ListProductsRequestBuilder) From(from string) *ListProductsRequestBuilder {
	b.req.From = &from

	return b
}

func (b *ListProductsRequestBuilder) To(to string) *ListProductsRequestBuilder {
	b.req.To = &to

	return b
}

func (b *ListProductsRequestBuilder) DateRange(from, to string) *ListProductsRequestBuilder {
	b.req.From = &from
	b.req.To = &to

	return b
}

func (b *ListProductsRequestBuilder) Build() *ListProductsRequest {
	return b.req
}

type ListProductsResponse = types.Response[[]types.Product]

func (c *Client) List(ctx context.Context, builder *ListProductsRequestBuilder) (*ListProductsResponse, error) {
	params := url.Values{}
	if builder != nil {
		req := builder.Build()
		if req.PerPage != nil {
			params.Set("perPage", strconv.Itoa(*req.PerPage))
		}
		if req.Page != nil {
			params.Set("page", strconv.Itoa(*req.Page))
		}
		if req.From != nil {
			params.Set("from", *req.From)
		}
		if req.To != nil {
			params.Set("to", *req.To)
		}
	}

	path := basePath
	if len(params) > 0 {
		path = fmt.Sprintf("%s?%s", basePath, params.Encode())
	}

	return net.Get[[]types.Product](ctx, c.Client, c.Secret, path, c.BaseURL)
}
