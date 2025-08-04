package settlements

import (
	"context"
	"fmt"
	"time"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type SettlementTransactionListRequest struct {
	PerPage *int       `json:"perPage,omitempty"` // Optional: records per page (default: 50)
	Page    *int       `json:"page,omitempty"`    // Optional: page number (default: 1)
	From    *time.Time `json:"from,omitempty"`    // Optional: start date filter
	To      *time.Time `json:"to,omitempty"`      // Optional: end date filter
}

type ListSettlementTransactionsRequestBuilder struct {
	req *SettlementTransactionListRequest
}

func NewSettlementTransactionListRequest() *ListSettlementTransactionsRequestBuilder {
	return &ListSettlementTransactionsRequestBuilder{
		req: &SettlementTransactionListRequest{},
	}
}

func (b *ListSettlementTransactionsRequestBuilder) PerPage(perPage int) *ListSettlementTransactionsRequestBuilder {
	b.req.PerPage = &perPage

	return b
}

func (b *ListSettlementTransactionsRequestBuilder) Page(page int) *ListSettlementTransactionsRequestBuilder {
	b.req.Page = &page

	return b
}

func (b *ListSettlementTransactionsRequestBuilder) DateRange(from, to time.Time) *ListSettlementTransactionsRequestBuilder {
	b.req.From = &from
	b.req.To = &to

	return b
}

func (b *ListSettlementTransactionsRequestBuilder) From(from time.Time) *ListSettlementTransactionsRequestBuilder {
	b.req.From = &from

	return b
}

func (b *ListSettlementTransactionsRequestBuilder) To(to time.Time) *ListSettlementTransactionsRequestBuilder {
	b.req.To = &to

	return b
}

func (b *ListSettlementTransactionsRequestBuilder) Build() *SettlementTransactionListRequest {
	return b.req
}

type ListSettlementTransactionsResponse = types.Response[[]types.SettlementTransaction]

func (c *Client) ListTransactions(ctx context.Context, settlementID string, builder *ListSettlementTransactionsRequestBuilder) (*ListSettlementTransactionsResponse, error) {
	return net.Get[[]types.SettlementTransaction](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s/transactions", basePath, settlementID), c.BaseURL)
}
