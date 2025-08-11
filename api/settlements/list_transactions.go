package settlements

import (
	"context"
	"fmt"
	"time"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type listTransactionsRequest struct {
	PerPage *int       `json:"perPage,omitempty"` // Optional: records per page (default: 50)
	Page    *int       `json:"page,omitempty"`    // Optional: page number (default: 1)
	From    *time.Time `json:"from,omitempty"`    // Optional: start date filter
	To      *time.Time `json:"to,omitempty"`      // Optional: end date filter
}

type ListTransactionsRequestBuilder struct {
	req *listTransactionsRequest
}

func NewListTransactionsRequestBuilder() *ListTransactionsRequestBuilder {
	return &ListTransactionsRequestBuilder{
		req: &listTransactionsRequest{},
	}
}

func (b *ListTransactionsRequestBuilder) PerPage(perPage int) *ListTransactionsRequestBuilder {
	b.req.PerPage = &perPage

	return b
}

func (b *ListTransactionsRequestBuilder) Page(page int) *ListTransactionsRequestBuilder {
	b.req.Page = &page

	return b
}

func (b *ListTransactionsRequestBuilder) DateRange(from, to time.Time) *ListTransactionsRequestBuilder {
	b.req.From = &from
	b.req.To = &to

	return b
}

func (b *ListTransactionsRequestBuilder) From(from time.Time) *ListTransactionsRequestBuilder {
	b.req.From = &from

	return b
}

func (b *ListTransactionsRequestBuilder) To(to time.Time) *ListTransactionsRequestBuilder {
	b.req.To = &to

	return b
}

func (b *ListTransactionsRequestBuilder) Build() *listTransactionsRequest {
	return b.req
}

type ListTransactionsResponseData = []types.SettlementTransaction
type ListTransactionsResponse = types.Response[ListTransactionsResponseData]

func (c *Client) ListTransactions(ctx context.Context, settlementID string, builder ListTransactionsRequestBuilder) (*ListTransactionsResponse, error) {
	return net.Get[ListTransactionsResponseData](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s/transactions", basePath, settlementID), c.BaseURL)
}
