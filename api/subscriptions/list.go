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
	// Optional
	PerPage  *int
	Page     *int
	Customer *int // Customer ID
	Plan     *int // Plan ID
}

// SubscriptionListRequestBuilder provides a fluent interface for building SubscriptionListRequest
type SubscriptionListRequestBuilder struct {
	req *SubscriptionListRequest
}

// NewSubscriptionListRequest creates a new builder for SubscriptionListRequest
func NewSubscriptionListRequest() *SubscriptionListRequestBuilder {
	return &SubscriptionListRequestBuilder{
		req: &SubscriptionListRequest{},
	}
}

// PerPage sets the number of records per page
func (b *SubscriptionListRequestBuilder) PerPage(perPage int) *SubscriptionListRequestBuilder {
	b.req.PerPage = optional.Int(perPage)

	return b
}

// Page sets the page number
func (b *SubscriptionListRequestBuilder) Page(page int) *SubscriptionListRequestBuilder {
	b.req.Page = optional.Int(page)

	return b
}

// Customer filters by customer ID
func (b *SubscriptionListRequestBuilder) Customer(customer int) *SubscriptionListRequestBuilder {
	b.req.Customer = optional.Int(customer)

	return b
}

// Plan filters by plan ID
func (b *SubscriptionListRequestBuilder) Plan(plan int) *SubscriptionListRequestBuilder {
	b.req.Plan = optional.Int(plan)

	return b
}

// Build returns the constructed SubscriptionListRequest
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

// SubscriptionListResponse represents the response from listing subscriptions
type SubscriptionListResponse = types.Response[[]types.Subscription]

// List lists subscriptions using a builder (fluent interface)
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
