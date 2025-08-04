package refunds

import (
	"context"
	"time"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type RefundListRequest struct {
	Transaction *string    `json:"transaction,omitempty"`
	Currency    *string    `json:"currency,omitempty"`
	From        *time.Time `json:"from,omitempty"`
	To          *time.Time `json:"to,omitempty"`
	PerPage     *int       `json:"perPage,omitempty"`
	Page        *int       `json:"page,omitempty"`
}

type RefundListRequestBuilder struct {
	req *RefundListRequest
}

func NewRefundListRequest() *RefundListRequestBuilder {
	return &RefundListRequestBuilder{
		req: &RefundListRequest{},
	}
}

func (b *RefundListRequestBuilder) Transaction(transaction string) *RefundListRequestBuilder {
	b.req.Transaction = &transaction

	return b
}

func (b *RefundListRequestBuilder) Currency(currency string) *RefundListRequestBuilder {
	b.req.Currency = &currency

	return b
}

func (b *RefundListRequestBuilder) DateRange(from, to time.Time) *RefundListRequestBuilder {
	b.req.From = &from
	b.req.To = &to

	return b
}

func (b *RefundListRequestBuilder) From(from time.Time) *RefundListRequestBuilder {
	b.req.From = &from

	return b
}

func (b *RefundListRequestBuilder) To(to time.Time) *RefundListRequestBuilder {
	b.req.To = &to

	return b
}

func (b *RefundListRequestBuilder) PerPage(perPage int) *RefundListRequestBuilder {
	b.req.PerPage = &perPage

	return b
}

func (b *RefundListRequestBuilder) Page(page int) *RefundListRequestBuilder {
	b.req.Page = &page

	return b
}

func (b *RefundListRequestBuilder) Build() *RefundListRequest {
	return b.req
}

type RefundListResponse = types.Response[[]types.Refund]

func (c *Client) List(ctx context.Context, builder *RefundListRequestBuilder) (*RefundListResponse, error) {
	return net.Get[[]types.Refund](ctx, c.Client, c.Secret, basePath, c.BaseURL)
}
