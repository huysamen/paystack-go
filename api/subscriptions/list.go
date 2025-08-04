package subscriptions

import (
	"context"
	"fmt"
	"net/url"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/optional"
	"github.com/huysamen/paystack-go/types"
)

type SubscriptionListRequest struct {
	PerPage  *int
	Page     *int
	Customer *int // Customer ID
	Plan     *int // Plan ID
}

type SubscriptionListRequestBuilder struct {
	req *SubscriptionListRequest
}

func NewSubscriptionListRequest() *SubscriptionListRequestBuilder {
	return &SubscriptionListRequestBuilder{
		req: &SubscriptionListRequest{},
	}
}

func (b *SubscriptionListRequestBuilder) PerPage(perPage int) *SubscriptionListRequestBuilder {
	b.req.PerPage = optional.Int(perPage)

	return b
}

func (b *SubscriptionListRequestBuilder) Page(page int) *SubscriptionListRequestBuilder {
	b.req.Page = optional.Int(page)

	return b
}

func (b *SubscriptionListRequestBuilder) Customer(customer int) *SubscriptionListRequestBuilder {
	b.req.Customer = optional.Int(customer)

	return b
}

func (b *SubscriptionListRequestBuilder) Plan(plan int) *SubscriptionListRequestBuilder {
	b.req.Plan = optional.Int(plan)

	return b
}

func (b *SubscriptionListRequestBuilder) Build() *SubscriptionListRequest {
	return b.req
}

func (r *SubscriptionListRequest) toQuery() string {
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

type SubscriptionListResponse = types.Response[[]types.Subscription]

func (c *Client) List(ctx context.Context, builder *SubscriptionListRequestBuilder) (*SubscriptionListResponse, error) {
	path := basePath

	if builder != nil {
		req := builder.Build()
		if query := req.toQuery(); query != "" {
			path += "?" + query
		}
	}

	return net.Get[[]types.Subscription](ctx, c.Client, c.Secret, path, c.BaseURL)
}
