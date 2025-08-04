package subscriptions

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type SubscriptionEnableRequest struct {
	Code  string `json:"code"`  // Subscription code
	Token string `json:"token"` // Email token
}

type SubscriptionEnableRequestBuilder struct {
	req *SubscriptionEnableRequest
}

func NewSubscriptionEnableRequest(code, token string) *SubscriptionEnableRequestBuilder {
	return &SubscriptionEnableRequestBuilder{
		req: &SubscriptionEnableRequest{
			Code:  code,
			Token: token,
		},
	}
}

func (b *SubscriptionEnableRequestBuilder) Build() *SubscriptionEnableRequest {
	return b.req
}

type SubscriptionEnableResponse = types.Response[any]

func (c *Client) Enable(ctx context.Context, builder *SubscriptionEnableRequestBuilder) (*SubscriptionEnableResponse, error) {
	return net.Post[SubscriptionEnableRequest, any](ctx, c.Client, c.Secret, basePath+"/enable", builder.Build(), c.BaseURL)
}
