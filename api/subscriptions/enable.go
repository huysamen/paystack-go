package subscriptions

import (
	"context"
	"errors"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type SubscriptionEnableRequest struct {
	Code  string `json:"code"`  // Subscription code
	Token string `json:"token"` // Email token
}

type SubscriptionEnableResponse struct {
	Message string `json:"message"`
}

func (c *Client) Enable(ctx context.Context, req *SubscriptionEnableRequest) (*types.Response[SubscriptionEnableResponse], error) {
	if req == nil {
		return nil, errors.New("request cannot be nil")
	}

	path := subscriptionBasePath + "/enable"

	return net.Post[SubscriptionEnableRequest, SubscriptionEnableResponse](
		ctx,
		c.client,
		c.secret,
		path,
		req,
		c.baseURL,
	)
}
