package settlements

import (
	"context"
	"fmt"
	"time"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// SettlementTransactionListRequest represents the request to list settlement transactions
type SettlementTransactionListRequest struct {
	PerPage *int       `json:"perPage,omitempty"` // Optional: records per page (default: 50)
	Page    *int       `json:"page,omitempty"`    // Optional: page number (default: 1)
	From    *time.Time `json:"from,omitempty"`    // Optional: start date filter
	To      *time.Time `json:"to,omitempty"`      // Optional: end date filter
}

// ListSettlementTransactionsRequestBuilder provides a fluent interface for building SettlementTransactionListRequest
type ListSettlementTransactionsRequestBuilder struct {
	req *SettlementTransactionListRequest
}

// NewSettlementTransactionListRequest creates a new builder for SettlementTransactionListRequest
func NewSettlementTransactionListRequest() *ListSettlementTransactionsRequestBuilder {
	return &ListSettlementTransactionsRequestBuilder{
		req: &SettlementTransactionListRequest{},
	}
}

// PerPage sets the number of records per page
func (b *ListSettlementTransactionsRequestBuilder) PerPage(perPage int) *ListSettlementTransactionsRequestBuilder {
	b.req.PerPage = &perPage

	return b
}

// Page sets the page number
func (b *ListSettlementTransactionsRequestBuilder) Page(page int) *ListSettlementTransactionsRequestBuilder {
	b.req.Page = &page

	return b
}

// DateRange sets both start and end date filters
func (b *ListSettlementTransactionsRequestBuilder) DateRange(from, to time.Time) *ListSettlementTransactionsRequestBuilder {
	b.req.From = &from
	b.req.To = &to

	return b
}

// From sets the start date filter
func (b *ListSettlementTransactionsRequestBuilder) From(from time.Time) *ListSettlementTransactionsRequestBuilder {
	b.req.From = &from

	return b
}

// To sets the end date filter
func (b *ListSettlementTransactionsRequestBuilder) To(to time.Time) *ListSettlementTransactionsRequestBuilder {
	b.req.To = &to

	return b
}

// Build returns the constructed SettlementTransactionListRequest
func (b *ListSettlementTransactionsRequestBuilder) Build() *SettlementTransactionListRequest {
	return b.req
}

// ListSettlementTransactionsResponse represents the response from listing settlement transactions
type ListSettlementTransactionsResponse = types.Response[[]types.SettlementTransaction]

// ListTransactions retrieves transactions for a specific settlement using a builder (fluent interface)
func (c *Client) ListTransactions(ctx context.Context, settlementID string, builder *ListSettlementTransactionsRequestBuilder) (*ListSettlementTransactionsResponse, error) {
	return net.Get[[]types.SettlementTransaction](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s/transactions", basePath, settlementID), c.BaseURL)
}
