package customers

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/optional"
	"github.com/huysamen/paystack-go/types"
)

type CustomerListRequest struct {
	PerPage *int
	Page    *int
	From    *time.Time
	To      *time.Time
}

type CustomerListRequestBuilder struct {
	req *CustomerListRequest
}

func NewCustomerListRequest() *CustomerListRequestBuilder {
	return &CustomerListRequestBuilder{
		req: &CustomerListRequest{},
	}
}

func (b *CustomerListRequestBuilder) PerPage(perPage int) *CustomerListRequestBuilder {
	b.req.PerPage = optional.Int(perPage)

	return b
}

func (b *CustomerListRequestBuilder) Page(page int) *CustomerListRequestBuilder {
	b.req.Page = optional.Int(page)

	return b
}

func (b *CustomerListRequestBuilder) DateRange(from, to time.Time) *CustomerListRequestBuilder {
	b.req.From = optional.Time(from)
	b.req.To = optional.Time(to)

	return b
}

func (b *CustomerListRequestBuilder) From(from time.Time) *CustomerListRequestBuilder {
	b.req.From = optional.Time(from)

	return b
}

func (b *CustomerListRequestBuilder) To(to time.Time) *CustomerListRequestBuilder {
	b.req.To = optional.Time(to)

	return b
}

func (b *CustomerListRequestBuilder) Build() *CustomerListRequest {
	return b.req
}

func (r *CustomerListRequest) toQuery() string {
	params := url.Values{}

	if r.PerPage != nil {
		params.Add("perPage", fmt.Sprintf("%d", *r.PerPage))
	}

	if r.Page != nil {
		params.Add("page", fmt.Sprintf("%d", *r.Page))
	}

	if r.From != nil {
		params.Add("from", r.From.Format("2006-01-02T15:04:05.999Z"))
	}

	if r.To != nil {
		params.Add("to", r.To.Format("2006-01-02T15:04:05.999Z"))
	}

	return params.Encode()
}

type CustomerListResponseData struct {
	Data []types.Customer `json:"data"`
	Meta types.Meta       `json:"meta"`
}

type CustomerListResponse = types.Response[CustomerListResponseData]

func (c *Client) List(ctx context.Context, builder *CustomerListRequestBuilder) (*CustomerListResponse, error) {
	req := builder.Build()
	path := basePath

	if req != nil {
		if query := req.toQuery(); query != "" {
			path += "?" + query
		}
	}

	return net.Get[CustomerListResponseData](ctx, c.Client, c.Secret, path, c.BaseURL)
}
