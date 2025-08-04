package products

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// ListProductsRequest represents the request to list products
type ListProductsRequest struct {
	PerPage *int    `json:"perPage,omitempty"`
	Page    *int    `json:"page,omitempty"`
	From    *string `json:"from,omitempty"`
	To      *string `json:"to,omitempty"`
}

// ListProductsRequestBuilder provides a fluent interface for building ListProductsRequest
type ListProductsRequestBuilder struct {
	req *ListProductsRequest
}

// NewListProductsRequest creates a new builder for ListProductsRequest
func NewListProductsRequest() *ListProductsRequestBuilder {
	return &ListProductsRequestBuilder{
		req: &ListProductsRequest{},
	}
}

// PerPage sets the number of records per page
func (b *ListProductsRequestBuilder) PerPage(perPage int) *ListProductsRequestBuilder {
	b.req.PerPage = &perPage

	return b
}

// Page sets the page number
func (b *ListProductsRequestBuilder) Page(page int) *ListProductsRequestBuilder {
	b.req.Page = &page

	return b
}

// From sets the start date filter
func (b *ListProductsRequestBuilder) From(from string) *ListProductsRequestBuilder {
	b.req.From = &from

	return b
}

// To sets the end date filter
func (b *ListProductsRequestBuilder) To(to string) *ListProductsRequestBuilder {
	b.req.To = &to

	return b
}

// DateRange sets both start and end date filters
func (b *ListProductsRequestBuilder) DateRange(from, to string) *ListProductsRequestBuilder {
	b.req.From = &from
	b.req.To = &to

	return b
}

// Build returns the constructed ListProductsRequest
func (b *ListProductsRequestBuilder) Build() *ListProductsRequest {
	return b.req
}

// ListProductsResponse represents the response from listing products
type ListProductsResponse = types.Response[[]types.Product]

// List retrieves products available on your integration using a builder
func (c *Client) List(ctx context.Context, builder *ListProductsRequestBuilder) (*ListProductsResponse, error) {
	// Build query parameters
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
