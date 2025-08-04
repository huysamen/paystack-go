package disputes

import (
	"context"
	"net/url"
	"strconv"
	"time"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type ListDisputesRequest struct {
	From        *time.Time           `json:"from,omitempty"`
	To          *time.Time           `json:"to,omitempty"`
	PerPage     *int                 `json:"per_page,omitempty"`
	Page        *int                 `json:"page,omitempty"`
	Transaction *string              `json:"transaction,omitempty"`
	Status      *types.DisputeStatus `json:"status,omitempty"`
}

type ListDisputesBuilder struct {
	req *ListDisputesRequest
}

func NewListDisputesBuilder() *ListDisputesBuilder {
	return &ListDisputesBuilder{
		req: &ListDisputesRequest{},
	}
}

func (b *ListDisputesBuilder) From(from time.Time) *ListDisputesBuilder {
	b.req.From = &from

	return b
}

func (b *ListDisputesBuilder) To(to time.Time) *ListDisputesBuilder {
	b.req.To = &to

	return b
}

func (b *ListDisputesBuilder) DateRange(from, to time.Time) *ListDisputesBuilder {
	b.req.From = &from
	b.req.To = &to

	return b
}

func (b *ListDisputesBuilder) PerPage(perPage int) *ListDisputesBuilder {
	b.req.PerPage = &perPage

	return b
}

func (b *ListDisputesBuilder) Page(page int) *ListDisputesBuilder {
	b.req.Page = &page

	return b
}

func (b *ListDisputesBuilder) Transaction(transaction string) *ListDisputesBuilder {
	b.req.Transaction = &transaction

	return b
}

func (b *ListDisputesBuilder) Status(status types.DisputeStatus) *ListDisputesBuilder {
	b.req.Status = &status

	return b
}

func (b *ListDisputesBuilder) Build() *ListDisputesRequest {
	return b.req
}

type ListDisputesResponse = types.Response[[]types.Dispute]

func (c *Client) List(ctx context.Context, builder *ListDisputesBuilder) (*ListDisputesResponse, error) {
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

	return net.Get[[]types.Dispute](ctx, c.Client, c.Secret, endpoint, c.BaseURL)
}
