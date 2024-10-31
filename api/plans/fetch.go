package plans

import (
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

func (c *Client) Fetch(idOrCode string) (*types.Response[PlanFetchResponse], error) {
	return net.Get[PlanFetchResponse](
		c.client,
		c.secret,
		fmt.Sprintf("%s/%s", planBasePath, idOrCode),
	)
}
