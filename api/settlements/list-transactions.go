package settlements

import (
	"context"
	"fmt"
	"time"

	"github.com/huysamen/paystack-go/net"
)

// SettlementTransactionListRequestBuilder provides a fluent interface for building SettlementTransactionListRequest
type SettlementTransactionListRequestBuilder struct {
	req *SettlementTransactionListRequest
}

// NewSettlementTransactionListRequest creates a new builder for SettlementTransactionListRequest
func NewSettlementTransactionListRequest() *SettlementTransactionListRequestBuilder {
	return &SettlementTransactionListRequestBuilder{
		req: &SettlementTransactionListRequest{},
	}
}

// PerPage sets the number of records per page
func (b *SettlementTransactionListRequestBuilder) PerPage(perPage int) *SettlementTransactionListRequestBuilder {
	b.req.PerPage = &perPage
	return b
}

// Page sets the page number
func (b *SettlementTransactionListRequestBuilder) Page(page int) *SettlementTransactionListRequestBuilder {
	b.req.Page = &page
	return b
}

// DateRange sets both start and end date filters
func (b *SettlementTransactionListRequestBuilder) DateRange(from, to time.Time) *SettlementTransactionListRequestBuilder {
	b.req.From = &from
	b.req.To = &to
	return b
}

// From sets the start date filter
func (b *SettlementTransactionListRequestBuilder) From(from time.Time) *SettlementTransactionListRequestBuilder {
	b.req.From = &from
	return b
}

// To sets the end date filter
func (b *SettlementTransactionListRequestBuilder) To(to time.Time) *SettlementTransactionListRequestBuilder {
	b.req.To = &to
	return b
}

// Build returns the constructed SettlementTransactionListRequest
func (b *SettlementTransactionListRequestBuilder) Build() *SettlementTransactionListRequest {
	return b.req
}

// ListTransactions retrieves transactions for a specific settlement using a builder (fluent interface)
func (c *Client) ListTransactions(ctx context.Context, settlementID string, builder *SettlementTransactionListRequestBuilder) (*SettlementTransactionListResponse, error) {
	if builder == nil {
		return nil, ErrBuilderRequired
	}

	url := fmt.Sprintf("%s%s/%s/transactions", c.baseURL, settlementBasePath, settlementID)
	response, err := net.Get[[]SettlementTransaction](ctx, c.client, c.secret, url)
	if err != nil {
		return nil, err
	}
	return response, nil
}
