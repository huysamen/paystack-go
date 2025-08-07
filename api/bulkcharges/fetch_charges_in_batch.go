package bulkcharges

import (
	"context"
	"net/url"
	"strconv"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type fetchInBatchRequest struct {
	Status  *string `json:"status,omitempty"`
	PerPage *int    `json:"perPage,omitempty"`
	Page    *int    `json:"page,omitempty"`
	From    *string `json:"from,omitempty"`
	To      *string `json:"to,omitempty"`
}

type FetchInBatchRequestBuilder struct {
	req *fetchInBatchRequest
}

func NewFetchInBatchRequestBuilder() *FetchInBatchRequestBuilder {
	return &FetchInBatchRequestBuilder{
		req: &fetchInBatchRequest{},
	}
}

func (b *FetchInBatchRequestBuilder) Status(status string) *FetchInBatchRequestBuilder {
	b.req.Status = &status

	return b
}

func (b *FetchInBatchRequestBuilder) PerPage(perPage int) *FetchInBatchRequestBuilder {
	b.req.PerPage = &perPage

	return b
}

func (b *FetchInBatchRequestBuilder) Page(page int) *FetchInBatchRequestBuilder {
	b.req.Page = &page

	return b
}

func (b *FetchInBatchRequestBuilder) From(from string) *FetchInBatchRequestBuilder {
	b.req.From = &from

	return b
}

func (b *FetchInBatchRequestBuilder) To(to string) *FetchInBatchRequestBuilder {
	b.req.To = &to

	return b
}

func (b *FetchInBatchRequestBuilder) DateRange(from, to string) *FetchInBatchRequestBuilder {
	b.req.From = &from
	b.req.To = &to

	return b
}

func (b *FetchInBatchRequestBuilder) Build() *fetchInBatchRequest {
	return b.req
}

func (r *fetchInBatchRequest) toQuery() string {
	params := url.Values{}

	if r.Status != nil {
		params.Set("status", *r.Status)
	}

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

type FetchInBatchResponseData = []types.BulkCharge
type FetchInBatchResponse = types.Response[FetchInBatchResponseData]

func (c *Client) FetchChargesInBatch(ctx context.Context, idOrCode string, builder FetchInBatchRequestBuilder) (*FetchInBatchResponse, error) {
	req := builder.Build()
	path := basePath + "/" + idOrCode + fetchChargesPath

	if req != nil {
		if query := req.toQuery(); query != "" {
			path += "?" + query
		}
	}

	return net.Get[FetchInBatchResponseData](ctx, c.Client, c.Secret, path, c.BaseURL)
}
