package settlements

import (
	"context"
	"time"

	"github.com/huysamen/paystack-go/enums"
	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type listRequest struct {
	PerPage    *int                    `json:"perPage,omitempty"`    // Optional: records per page (default: 50)
	Page       *int                    `json:"page,omitempty"`       // Optional: page number (default: 1)
	Status     *enums.SettlementStatus `json:"status,omitempty"`     // Optional: filter by status
	Subaccount *string                 `json:"subaccount,omitempty"` // Optional: filter by subaccount ID (use "none" for main account only)
	From       *time.Time              `json:"from,omitempty"`       // Optional: start date filter
	To         *time.Time              `json:"to,omitempty"`         // Optional: end date filter
}

type ListRequestBuilder struct {
	req *listRequest
}

func NewListRequestBuilder() *ListRequestBuilder {
	return &ListRequestBuilder{
		req: &listRequest{},
	}
}

func (b *ListRequestBuilder) PerPage(perPage int) *ListRequestBuilder {
	b.req.PerPage = &perPage

	return b
}

func (b *ListRequestBuilder) Page(page int) *ListRequestBuilder {
	b.req.Page = &page

	return b
}

func (b *ListRequestBuilder) Status(status enums.SettlementStatus) *ListRequestBuilder {
	b.req.Status = &status

	return b
}

func (b *ListRequestBuilder) Subaccount(subaccount string) *ListRequestBuilder {
	b.req.Subaccount = &subaccount

	return b
}

func (b *ListRequestBuilder) MainAccountOnly() *ListRequestBuilder {
	none := "none"
	b.req.Subaccount = &none

	return b
}

func (b *ListRequestBuilder) DateRange(from, to time.Time) *ListRequestBuilder {
	b.req.From = &from
	b.req.To = &to

	return b
}

func (b *ListRequestBuilder) From(from time.Time) *ListRequestBuilder {
	b.req.From = &from

	return b
}

func (b *ListRequestBuilder) To(to time.Time) *ListRequestBuilder {
	b.req.To = &to

	return b
}

func (b *ListRequestBuilder) Build() *listRequest {
	return b.req
}

type ListResponseData = []types.Settlement
type ListResponse = types.Response[ListResponseData]

func (c *Client) List(ctx context.Context, builder ListRequestBuilder) (*ListResponse, error) {
	return net.Get[ListResponseData](ctx, c.Client, c.Secret, basePath, c.BaseURL)
}
