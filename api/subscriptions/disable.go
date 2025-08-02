package subscriptions

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// SubscriptionDisableRequest represents the request to disable a subscription
type SubscriptionDisableRequest struct {
	Code  string `json:"code"`  // Subscription code
	Token string `json:"token"` // Email token
}

// SubscriptionDisableRequestBuilder provides a fluent interface for building SubscriptionDisableRequest
type SubscriptionDisableRequestBuilder struct {
	req *SubscriptionDisableRequest
}

// NewSubscriptionDisableRequest creates a new builder for SubscriptionDisableRequest
func NewSubscriptionDisableRequest(code, token string) *SubscriptionDisableRequestBuilder {
	return &SubscriptionDisableRequestBuilder{
		req: &SubscriptionDisableRequest{
			Code:  code,
			Token: token,
		},
	}
}

// Build returns the constructed SubscriptionDisableRequest
func (b *SubscriptionDisableRequestBuilder) Build() *SubscriptionDisableRequest {
	return b.req
}

// SubscriptionDisableResponse represents the response from disabling a subscription.
type SubscriptionDisableResponse types.Response[struct {
	Message string `json:"message"`
}]

func (c *Client) Disable(ctx context.Context, builder *SubscriptionDisableRequestBuilder) (*SubscriptionDisableResponse, error) {
	if builder == nil {
		return nil, ErrBuilderRequired
	}

	req := builder.Build()
	path := subscriptionBasePath + "/disable"

	resp, err := net.Post[SubscriptionDisableRequest, struct {
		Message string `json:"message"`
	}](
		ctx,
		c.client,
		c.secret,
		path,
		req,
		c.baseURL,
	)
	if err != nil {
		return nil, err
	}

	response := SubscriptionDisableResponse(*resp)
	return &response, nil
}
