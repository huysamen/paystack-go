package plans

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type PlanFetchResponse struct {
	types.Plan

	Subscriptions []types.Subscription `json:"subscriptions"`
	Pages         []types.Page         `json:"pages"`
	Subscribers   []types.Subscriber   `json:"subscribers"`
}

func (c *Client) FetchByID(ctx context.Context, id uint64) (*types.Response[PlanFetchResponse], error) {
	return net.Get[PlanFetchResponse](
		ctx,
		c.client,
		c.secret,
		fmt.Sprintf("%s/%d", planBasePath, id),
		c.baseURL,
	)
}

func (c *Client) FetchByCode(ctx context.Context, code string) (*types.Response[PlanFetchResponse], error) {
	return net.Get[PlanFetchResponse](
		ctx,
		c.client,
		c.secret,
		fmt.Sprintf("%s/%s", planBasePath, code),
		c.baseURL,
	)
}
