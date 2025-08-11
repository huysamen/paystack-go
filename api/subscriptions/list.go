package subscriptions

import (
	"context"
	"fmt"
	"net/url"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/optional"
	"github.com/huysamen/paystack-go/types"
)

type listRequest struct {
	PerPage  *int
	Page     *int
	Customer *int // Customer ID
	Plan     *int // Plan ID
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
	b.req.PerPage = optional.Int(perPage)

	return b
}

func (b *ListRequestBuilder) Page(page int) *ListRequestBuilder {
	b.req.Page = optional.Int(page)

	return b
}

func (b *ListRequestBuilder) Customer(customer int) *ListRequestBuilder {
	b.req.Customer = optional.Int(customer)

	return b
}

func (b *ListRequestBuilder) Plan(plan int) *ListRequestBuilder {
	b.req.Plan = optional.Int(plan)

	return b
}

func (b *ListRequestBuilder) Build() *listRequest {
	return b.req
}

func (r *listRequest) toQuery() string {
	params := url.Values{}

	if r.PerPage != nil {
		params.Add("perPage", fmt.Sprintf("%d", *r.PerPage))
	}
	if r.Page != nil {
		params.Add("page", fmt.Sprintf("%d", *r.Page))
	}
	if r.Customer != nil {
		params.Add("customer", fmt.Sprintf("%d", *r.Customer))
	}
	if r.Plan != nil {
		params.Add("plan", fmt.Sprintf("%d", *r.Plan))
	}

	return params.Encode()
}

type ListResponseData = []types.Subscription
type ListResponse = types.Response[ListResponseData]

func (c *Client) List(ctx context.Context, builder ListRequestBuilder) (*ListResponse, error) {
	path := basePath

	req := builder.Build()
	if query := req.toQuery(); query != "" {
		path += "?" + query
	}

	return net.Get[ListResponseData](ctx, c.Client, c.Secret, path, c.BaseURL)
}
