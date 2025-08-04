package plans

import (
	"context"
	"net/url"
	"strconv"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type ListPlansRequest struct {
	PerPage  *int            `json:"perPage,omitempty"`
	Page     *int            `json:"page,omitempty"`
	Status   *string         `json:"status,omitempty"`
	Interval *types.Interval `json:"interval,omitempty"`
	Amount   *int            `json:"amount,omitempty"`
}

type ListPlansRequestBuilder struct {
	req *ListPlansRequest
}

func NewListPlansRequest() *ListPlansRequestBuilder {
	return &ListPlansRequestBuilder{
		req: &ListPlansRequest{},
	}
}

func (b *ListPlansRequestBuilder) PerPage(perPage int) *ListPlansRequestBuilder {
	b.req.PerPage = &perPage

	return b
}

func (b *ListPlansRequestBuilder) Page(page int) *ListPlansRequestBuilder {
	b.req.Page = &page

	return b
}

func (b *ListPlansRequestBuilder) Status(status string) *ListPlansRequestBuilder {
	b.req.Status = &status

	return b
}

func (b *ListPlansRequestBuilder) Interval(interval types.Interval) *ListPlansRequestBuilder {
	b.req.Interval = &interval

	return b
}

func (b *ListPlansRequestBuilder) Amount(amount int) *ListPlansRequestBuilder {
	b.req.Amount = &amount

	return b
}

func (b *ListPlansRequestBuilder) Build() *ListPlansRequest {
	return b.req
}

func (r *ListPlansRequest) toQuery() string {
	params := url.Values{}

	if r.PerPage != nil {
		params.Add("perPage", strconv.Itoa(*r.PerPage))
	}
	if r.Page != nil {
		params.Add("page", strconv.Itoa(*r.Page))
	}
	if r.Status != nil {
		params.Add("status", *r.Status)
	}
	if r.Interval != nil {
		params.Add("interval", r.Interval.String())
	}
	if r.Amount != nil {
		params.Add("amount", strconv.Itoa(*r.Amount))
	}

	return params.Encode()
}

type ListPlansResponse = types.Response[[]types.Plan]

func (c *Client) List(ctx context.Context, builder *ListPlansRequestBuilder) (*ListPlansResponse, error) {
	path := basePath

	if builder != nil {
		req := builder.Build()
		if query := req.toQuery(); query != "" {
			path += "?" + query
		}
	}

	return net.Get[[]types.Plan](ctx, c.Client, c.Secret, path, c.BaseURL)
}
