package subscriptions

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type SubscriptionDisableRequest struct {
	Code  string `json:"code"`  // Subscription code
	Token string `json:"token"` // Email token
}

type SubscriptionDisableRequestBuilder struct {
	req *SubscriptionDisableRequest
}

func NewSubscriptionDisableRequest(code, token string) *SubscriptionDisableRequestBuilder {
	return &SubscriptionDisableRequestBuilder{
		req: &SubscriptionDisableRequest{
			Code:  code,
			Token: token,
		},
	}
}

func (b *SubscriptionDisableRequestBuilder) Build() *SubscriptionDisableRequest {
	return b.req
}

type SubscriptionDisableResponse = types.Response[any]

func (c *Client) Disable(ctx context.Context, builder *SubscriptionDisableRequestBuilder) (*SubscriptionDisableResponse, error) {
	return net.Post[SubscriptionDisableRequest, any](ctx, c.Client, c.Secret, basePath+"/disable", builder.Build(), c.BaseURL)
}
