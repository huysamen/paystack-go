package bulkcharges

import (
	"context"
	"net/url"
	"strconv"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// FetchChargesInBatchRequest represents the request to fetch charges in a batch
type FetchChargesInBatchRequest struct {
	Status  *string `json:"status,omitempty"`
	PerPage *int    `json:"perPage,omitempty"`
	Page    *int    `json:"page,omitempty"`
	From    *string `json:"from,omitempty"`
	To      *string `json:"to,omitempty"`
}

// FetchChargesInBatchRequestBuilder provides a fluent interface for building FetchChargesInBatchRequest
type FetchChargesInBatchRequestBuilder struct {
	req *FetchChargesInBatchRequest
}

// NewFetchChargesInBatchRequest creates a new builder for FetchChargesInBatchRequest
func NewFetchChargesInBatchRequest() *FetchChargesInBatchRequestBuilder {
	return &FetchChargesInBatchRequestBuilder{
		req: &FetchChargesInBatchRequest{},
	}
}

// Status filters by charge status
func (b *FetchChargesInBatchRequestBuilder) Status(status string) *FetchChargesInBatchRequestBuilder {
	b.req.Status = &status
	return b
}

// PerPage sets the number of records per page
func (b *FetchChargesInBatchRequestBuilder) PerPage(perPage int) *FetchChargesInBatchRequestBuilder {
	b.req.PerPage = &perPage
	return b
}

// Page sets the page number
func (b *FetchChargesInBatchRequestBuilder) Page(page int) *FetchChargesInBatchRequestBuilder {
	b.req.Page = &page
	return b
}

// From sets the start date filter
func (b *FetchChargesInBatchRequestBuilder) From(from string) *FetchChargesInBatchRequestBuilder {
	b.req.From = &from
	return b
}

// To sets the end date filter
func (b *FetchChargesInBatchRequestBuilder) To(to string) *FetchChargesInBatchRequestBuilder {
	b.req.To = &to
	return b
}

// DateRange sets both start and end date filters
func (b *FetchChargesInBatchRequestBuilder) DateRange(from, to string) *FetchChargesInBatchRequestBuilder {
	b.req.From = &from
	b.req.To = &to
	return b
}

// Build returns the constructed FetchChargesInBatchRequest
func (b *FetchChargesInBatchRequestBuilder) Build() *FetchChargesInBatchRequest {
	return b.req
}

// FetchChargesInBatch retrieves the charges associated with a specified batch code using a builder
func (c *Client) FetchChargesInBatch(ctx context.Context, idOrCode string, builder *FetchChargesInBatchRequestBuilder) (*types.Response[[]BulkChargeCharge], error) {
	req := builder.Build()

	params := url.Values{}
	if req.Status != nil {
		params.Set("status", *req.Status)
	}
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

	path := basePath + "/" + idOrCode + fetchChargesPath
	if len(params) > 0 {
		path += "?" + params.Encode()
	}

	return net.Get[[]BulkChargeCharge](ctx, c.Client, c.Secret, path, c.BaseURL)
}
