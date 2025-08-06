package plans

import (
	"context"
	"net/url"
	"strconv"

	"github.com/huysamen/paystack-go/enums"
	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type listRequest struct {
	PerPage  *int            `json:"perPage,omitempty"`
	Page     *int            `json:"page,omitempty"`
	Status   *string         `json:"status,omitempty"`
	Interval *enums.Interval `json:"interval,omitempty"`
	Amount   *int            `json:"amount,omitempty"`
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

func (b *ListRequestBuilder) Status(status string) *ListRequestBuilder {
	b.req.Status = &status

	return b
}

func (b *ListRequestBuilder) Interval(interval enums.Interval) *ListRequestBuilder {
	b.req.Interval = &interval

	return b
}

func (b *ListRequestBuilder) Amount(amount int) *ListRequestBuilder {
	b.req.Amount = &amount

	return b
}

func (b *ListRequestBuilder) Build() *listRequest {
	return b.req
}

func (r *listRequest) toQuery() string {
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

type ListResponseData = []types.Plan
type ListResponse = types.Response[ListResponseData]

func (c *Client) List(ctx context.Context, builder ListRequestBuilder) (*ListResponse, error) {
	path := basePath

	req := builder.Build()
	if query := req.toQuery(); query != "" {
		path += "?" + query
	}

	return net.Get[ListResponseData](ctx, c.Client, c.Secret, path, c.BaseURL)
}
