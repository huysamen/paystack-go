package bulkcharges

import (
	"context"
	"net/url"
	"strconv"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// ListBulkChargeBatchesRequest represents the request to list bulk charge batches
type ListBulkChargeBatchesRequest struct {
	PerPage *int    `json:"perPage,omitempty"`
	Page    *int    `json:"page,omitempty"`
	From    *string `json:"from,omitempty"`
	To      *string `json:"to,omitempty"`
}

// ListBulkChargeBatchesRequestBuilder provides a fluent interface for building ListBulkChargeBatchesRequest
type ListBulkChargeBatchesRequestBuilder struct {
	req *ListBulkChargeBatchesRequest
}

// NewListBulkChargeBatchesRequest creates a new builder for ListBulkChargeBatchesRequest
func NewListBulkChargeBatchesRequest() *ListBulkChargeBatchesRequestBuilder {
	return &ListBulkChargeBatchesRequestBuilder{
		req: &ListBulkChargeBatchesRequest{},
	}
}

// PerPage sets the number of records per page
func (b *ListBulkChargeBatchesRequestBuilder) PerPage(perPage int) *ListBulkChargeBatchesRequestBuilder {
	b.req.PerPage = &perPage
	return b
}

// Page sets the page number
func (b *ListBulkChargeBatchesRequestBuilder) Page(page int) *ListBulkChargeBatchesRequestBuilder {
	b.req.Page = &page
	return b
}

// From sets the start date filter
func (b *ListBulkChargeBatchesRequestBuilder) From(from string) *ListBulkChargeBatchesRequestBuilder {
	b.req.From = &from
	return b
}

// To sets the end date filter
func (b *ListBulkChargeBatchesRequestBuilder) To(to string) *ListBulkChargeBatchesRequestBuilder {
	b.req.To = &to
	return b
}

// DateRange sets both start and end date filters
func (b *ListBulkChargeBatchesRequestBuilder) DateRange(from, to string) *ListBulkChargeBatchesRequestBuilder {
	b.req.From = &from
	b.req.To = &to
	return b
}

// Build returns the constructed ListBulkChargeBatchesRequest
func (b *ListBulkChargeBatchesRequestBuilder) Build() *ListBulkChargeBatchesRequest {
	return b.req
}

// List retrieves all bulk charge batches created by the integration using a builder
func (c *Client) List(ctx context.Context, builder *ListBulkChargeBatchesRequestBuilder) (*types.Response[[]BulkChargeBatch], error) {
	req := builder.Build()

	params := url.Values{}
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

	path := basePath
	if len(params) > 0 {
		path += "?" + params.Encode()
	}

	return net.Get[[]BulkChargeBatch](ctx, c.Client, c.Secret, path, c.BaseURL)
}
