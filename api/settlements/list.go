package settlements

import (
	"context"
	"time"

	"github.com/huysamen/paystack-go/net"
)

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
func (b *SettlementListRequestBuilder) Status(status SettlementStatus) *SettlementListRequestBuilder {
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

// List retrieves a list of settlements using a builder (fluent interface)
func (c *Client) List(ctx context.Context, builder *SettlementListRequestBuilder) (*SettlementListResponse, error) {
	if builder == nil {
		builder = NewSettlementListRequest()
	}

	url := c.baseURL + settlementBasePath
	response, err := net.Get[[]Settlement](ctx, c.client, c.secret, url)
	if err != nil {
		return nil, err
	}
	return response, nil
}
