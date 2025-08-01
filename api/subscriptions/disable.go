package subscriptions

import (
	"context"
	"errors"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type SubscriptionDisableRequest struct {
	Code  string `json:"code"`  // Subscription code
	Token string `json:"token"` // Email token
}


type SubscriptionDisableResponse struct {
	Message string `json:"message"`
}

func (c *Client) Disable(ctx context.Context, req *SubscriptionDisableRequest) (*types.Response[SubscriptionDisableResponse], error) {
	if req == nil {
		return nil, errors.New("request cannot be nil")
	}


	path := subscriptionBasePath + "/disable"

	return net.Post[SubscriptionDisableRequest, SubscriptionDisableResponse](
		ctx,
		c.client,
		c.secret,
		path,
		req,
		c.baseURL,
	)
}
