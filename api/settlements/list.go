package settlements

import (
	"context"
	"time"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type SettlementListRequest struct {
	PerPage    *int                    `json:"perPage,omitempty"`    // Optional: records per page (default: 50)
	Page       *int                    `json:"page,omitempty"`       // Optional: page number (default: 1)
	Status     *types.SettlementStatus `json:"status,omitempty"`     // Optional: filter by status
	Subaccount *string                 `json:"subaccount,omitempty"` // Optional: filter by subaccount ID (use "none" for main account only)
	From       *time.Time              `json:"from,omitempty"`       // Optional: start date filter
	To         *time.Time              `json:"to,omitempty"`         // Optional: end date filter
}

type SettlementListRequestBuilder struct {
	req *SettlementListRequest
}

func NewSettlementListRequest() *SettlementListRequestBuilder {
	return &SettlementListRequestBuilder{
		req: &SettlementListRequest{},
	}
}

func (b *SettlementListRequestBuilder) PerPage(perPage int) *SettlementListRequestBuilder {
	b.req.PerPage = &perPage

	return b
}

func (b *SettlementListRequestBuilder) Page(page int) *SettlementListRequestBuilder {
	b.req.Page = &page

	return b
}

func (b *SettlementListRequestBuilder) Status(status types.SettlementStatus) *SettlementListRequestBuilder {
	b.req.Status = &status

	return b
}

func (b *SettlementListRequestBuilder) Subaccount(subaccount string) *SettlementListRequestBuilder {
	b.req.Subaccount = &subaccount

	return b
}

func (b *SettlementListRequestBuilder) MainAccountOnly() *SettlementListRequestBuilder {
	none := "none"
	b.req.Subaccount = &none

	return b
}

func (b *SettlementListRequestBuilder) DateRange(from, to time.Time) *SettlementListRequestBuilder {
	b.req.From = &from
	b.req.To = &to

	return b
}

func (b *SettlementListRequestBuilder) From(from time.Time) *SettlementListRequestBuilder {
	b.req.From = &from

	return b
}

func (b *SettlementListRequestBuilder) To(to time.Time) *SettlementListRequestBuilder {
	b.req.To = &to

	return b
}

func (b *SettlementListRequestBuilder) Build() *SettlementListRequest {
	return b.req
}

type SettlementListResponse = types.Response[[]types.Settlement]

func (c *Client) List(ctx context.Context, builder *SettlementListRequestBuilder) (*SettlementListResponse, error) {
	return net.Get[[]types.Settlement](ctx, c.Client, c.Secret, basePath, c.BaseURL)
}
