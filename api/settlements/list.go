package settlements

import (
	"context"
	"time"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// SettlementListRequest represents the request to list settlements
type SettlementListRequest struct {
	PerPage    *int                    `json:"perPage,omitempty"`    // Optional: records per page (default: 50)
	Page       *int                    `json:"page,omitempty"`       // Optional: page number (default: 1)
	Status     *types.SettlementStatus `json:"status,omitempty"`     // Optional: filter by status
	Subaccount *string                 `json:"subaccount,omitempty"` // Optional: filter by subaccount ID (use "none" for main account only)
	From       *time.Time              `json:"from,omitempty"`       // Optional: start date filter
	To         *time.Time              `json:"to,omitempty"`         // Optional: end date filter
}

// SettlementListRequestBuilder provides a fluent interface for building SettlementListRequest
type SettlementListRequestBuilder struct {
	req *SettlementListRequest
}

// NewSettlementListRequest creates a new builder for SettlementListRequest
func NewSettlementListRequest() *SettlementListRequestBuilder {
	return &SettlementListRequestBuilder{
		req: &SettlementListRequest{},
	}
}

// PerPage sets the number of records per page
func (b *SettlementListRequestBuilder) PerPage(perPage int) *SettlementListRequestBuilder {
	b.req.PerPage = &perPage

	return b
}

// Page sets the page number
func (b *SettlementListRequestBuilder) Page(page int) *SettlementListRequestBuilder {
	b.req.Page = &page

	return b
}

// Status filters by settlement status
func (b *SettlementListRequestBuilder) Status(status types.SettlementStatus) *SettlementListRequestBuilder {
	b.req.Status = &status

	return b
}

// Subaccount filters by subaccount ID (use "none" for main account only)
func (b *SettlementListRequestBuilder) Subaccount(subaccount string) *SettlementListRequestBuilder {
	b.req.Subaccount = &subaccount

	return b
}

// MainAccountOnly filters for main account settlements only
func (b *SettlementListRequestBuilder) MainAccountOnly() *SettlementListRequestBuilder {
	none := "none"
	b.req.Subaccount = &none

	return b
}

// DateRange sets both start and end date filters
func (b *SettlementListRequestBuilder) DateRange(from, to time.Time) *SettlementListRequestBuilder {
	b.req.From = &from
	b.req.To = &to

	return b
}

// From sets the start date filter
func (b *SettlementListRequestBuilder) From(from time.Time) *SettlementListRequestBuilder {
	b.req.From = &from

	return b
}

// To sets the end date filter
func (b *SettlementListRequestBuilder) To(to time.Time) *SettlementListRequestBuilder {
	b.req.To = &to

	return b
}

// Build returns the constructed SettlementListRequest
func (b *SettlementListRequestBuilder) Build() *SettlementListRequest {
	return b.req
}

// SettlementListResponse represents the response from listing settlements
type SettlementListResponse = types.Response[[]types.Settlement]

// List retrieves a list of settlements using a builder (fluent interface)
func (c *Client) List(ctx context.Context, builder *SettlementListRequestBuilder) (*SettlementListResponse, error) {
	return net.Get[[]types.Settlement](ctx, c.Client, c.Secret, basePath, c.BaseURL)
}
