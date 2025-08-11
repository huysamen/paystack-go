package disputes

import (
	"context"
	"net/url"
	"strconv"
	"time"

	"github.com/huysamen/paystack-go/enums"
	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type listRequest struct {
	From        *time.Time           `json:"from,omitempty"`
	To          *time.Time           `json:"to,omitempty"`
	PerPage     *int                 `json:"per_page,omitempty"`
	Page        *int                 `json:"page,omitempty"`
	Transaction *string              `json:"transaction,omitempty"`
	Status      *enums.DisputeStatus `json:"status,omitempty"`
}

type ListRequestBuilder struct {
	req *listRequest
}

func NewListRequestBuilder() *ListRequestBuilder {
	return &ListRequestBuilder{
		req: &listRequest{},
	}
}

func (b *ListRequestBuilder) From(from time.Time) *ListRequestBuilder {
	b.req.From = &from

	return b
}

func (b *ListRequestBuilder) To(to time.Time) *ListRequestBuilder {
	b.req.To = &to

	return b
}

func (b *ListRequestBuilder) DateRange(from, to time.Time) *ListRequestBuilder {
	b.req.From = &from
	b.req.To = &to

	return b
}

func (b *ListRequestBuilder) PerPage(perPage int) *ListRequestBuilder {
	b.req.PerPage = &perPage

	return b
}

func (b *ListRequestBuilder) Page(page int) *ListRequestBuilder {
	b.req.Page = &page

	return b
}

func (b *ListRequestBuilder) Transaction(transaction string) *ListRequestBuilder {
	b.req.Transaction = &transaction

	return b
}

func (b *ListRequestBuilder) Status(status enums.DisputeStatus) *ListRequestBuilder {
	b.req.Status = &status

	return b
}

func (b *ListRequestBuilder) Build() *listRequest {
	return b.req
}

func (r *listRequest) toQuery() string {
	params := url.Values{}
	if r.From != nil {
		params.Set("from", r.From.Format("2006-01-02"))
	}
	if r.To != nil {
		params.Set("to", r.To.Format("2006-01-02"))
	}
	if r.PerPage != nil {
		params.Set("per_page", strconv.Itoa(*r.PerPage))
	}
	if r.Page != nil {
		params.Set("page", strconv.Itoa(*r.Page))
	}
	if r.Transaction != nil {
		params.Set("transaction", *r.Transaction)
	}
	if r.Status != nil {
		params.Set("status", string(*r.Status))
	}

	return params.Encode()
}

type ListResponseData = []types.Dispute
type ListResponse = types.Response[ListResponseData]

func (c *Client) List(ctx context.Context, builder ListRequestBuilder) (*ListResponse, error) {
	path := basePath

	req := builder.Build()
	if req != nil {
		if query := req.toQuery(); query != "" {
			path += "?" + query
		}
	}

	return net.Get[ListResponseData](ctx, c.Client, c.Secret, path, c.BaseURL)
}
