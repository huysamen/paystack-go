package refunds

import (
	"context"
	"time"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// RefundListRequest represents the request payload for listing refunds
type RefundListRequest struct {
	Transaction *string    `json:"transaction,omitempty"`
	Currency    *string    `json:"currency,omitempty"`
	From        *time.Time `json:"from,omitempty"`
	To          *time.Time `json:"to,omitempty"`
	PerPage     *int       `json:"perPage,omitempty"`
	Page        *int       `json:"page,omitempty"`
}

// RefundListRequestBuilder provides a fluent interface for building RefundListRequest
type RefundListRequestBuilder struct {
	req *RefundListRequest
}

// NewRefundListRequest creates a new builder for RefundListRequest
func NewRefundListRequest() *RefundListRequestBuilder {
	return &RefundListRequestBuilder{
		req: &RefundListRequest{},
	}
}

// Transaction filters by transaction reference
func (b *RefundListRequestBuilder) Transaction(transaction string) *RefundListRequestBuilder {
	b.req.Transaction = &transaction

	return b
}

// Currency filters by currency
func (b *RefundListRequestBuilder) Currency(currency string) *RefundListRequestBuilder {
	b.req.Currency = &currency

	return b
}

// DateRange sets both start and end date filters
func (b *RefundListRequestBuilder) DateRange(from, to time.Time) *RefundListRequestBuilder {
	b.req.From = &from
	b.req.To = &to

	return b
}

// From sets the start date filter
func (b *RefundListRequestBuilder) From(from time.Time) *RefundListRequestBuilder {
	b.req.From = &from

	return b
}

// To sets the end date filter
func (b *RefundListRequestBuilder) To(to time.Time) *RefundListRequestBuilder {
	b.req.To = &to

	return b
}

// PerPage sets the number of records per page
func (b *RefundListRequestBuilder) PerPage(perPage int) *RefundListRequestBuilder {
	b.req.PerPage = &perPage

	return b
}

// Page sets the page number
func (b *RefundListRequestBuilder) Page(page int) *RefundListRequestBuilder {
	b.req.Page = &page

	return b
}

// Build returns the constructed RefundListRequest
func (b *RefundListRequestBuilder) Build() *RefundListRequest {
	return b.req
}

type RefundListResponse = types.Response[[]types.Refund]

// List retrieves all refunds available on your integration using a builder
func (c *Client) List(ctx context.Context, builder *RefundListRequestBuilder) (*RefundListResponse, error) {
	return net.Get[[]types.Refund](ctx, c.Client, c.Secret, basePath, c.BaseURL)
}
