package subscriptions

import (
	"context"
	"fmt"
	"net/url"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type SubscriptionListRequest struct {
	// Optional
	PerPage  *int
	Page     *int
	Customer *int // Customer ID
	Plan     *int // Plan ID
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

type SubscriptionListResponse struct {
	Data []Subscription `json:"data"`
	Meta types.Meta     `json:"meta"`
}

func (c *Client) List(ctx context.Context, req *SubscriptionListRequest) (*types.Response[SubscriptionListResponse], error) {
	path := subscriptionBasePath

	if req != nil {
		if query := req.toQuery(); query != "" {
			path += "?" + query
		}
	}

	return net.Get[SubscriptionListResponse](
		ctx,
		c.client,
		c.secret,
		path,
		c.baseURL,
	)
}
