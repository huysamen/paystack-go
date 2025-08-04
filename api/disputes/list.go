package disputes

import (
	"context"
	"net/url"
	"strconv"
	"time"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type listRequest struct {
	From        *time.Time           `json:"from,omitempty"`
	To          *time.Time           `json:"to,omitempty"`
	PerPage     *int                 `json:"per_page,omitempty"`
	Page        *int                 `json:"page,omitempty"`
	Transaction *string              `json:"transaction,omitempty"`
	Status      *types.DisputeStatus `json:"status,omitempty"`
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

func (b *ListRequestBuilder) Status(status types.DisputeStatus) *ListRequestBuilder {
	b.req.Status = &status

	return b
}

func (b *ListRequestBuilder) Build() *listRequest {
	return b.req
}

type ListResponseData = []types.Dispute
type ListDisputesResponse = types.Response[ListResponseData]

func (c *Client) List(ctx context.Context, builder *ListRequestBuilder) (*ListDisputesResponse, error) {
	endpoint := basePath
	req := builder.Build()

	params := url.Values{}
	if req.From != nil {
		params.Set("from", req.From.Format("2006-01-02"))
	}
	if req.To != nil {
		params.Set("to", req.To.Format("2006-01-02"))
	}
	if req.PerPage != nil {
		params.Set("per_page", strconv.Itoa(*req.PerPage))
	}
	if req.Page != nil {
		params.Set("page", strconv.Itoa(*req.Page))
	}
	if req.Transaction != nil {
		params.Set("transaction", *req.Transaction)
	}
	if req.Status != nil {
		params.Set("status", string(*req.Status))
	}

	if len(params) > 0 {
		endpoint += "?" + params.Encode()
	}

	return net.Get[ListResponseData](ctx, c.Client, c.Secret, endpoint, c.BaseURL)
}
