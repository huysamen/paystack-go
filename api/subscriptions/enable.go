package subscriptions

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// SubscriptionEnableRequest represents the request to enable a subscription
type SubscriptionEnableRequest struct {
	Code  string `json:"code"`  // Subscription code
	Token string `json:"token"` // Email token
}

// SubscriptionEnableRequestBuilder provides a fluent interface for building SubscriptionEnableRequest
type SubscriptionEnableRequestBuilder struct {
	req *SubscriptionEnableRequest
}

// NewSubscriptionEnableRequest creates a new builder for SubscriptionEnableRequest
func NewSubscriptionEnableRequest(code, token string) *SubscriptionEnableRequestBuilder {
	return &SubscriptionEnableRequestBuilder{
		req: &SubscriptionEnableRequest{
			Code:  code,
			Token: token,
		},
	}
}

// Build returns the constructed SubscriptionEnableRequest
func (b *SubscriptionEnableRequestBuilder) Build() *SubscriptionEnableRequest {
	return b.req
}

// SubscriptionEnableResponse represents the response from enabling a subscription.
type SubscriptionEnableResponse types.Response[struct {
	Message string `json:"message"`
}]

func (c *Client) Enable(ctx context.Context, builder *SubscriptionEnableRequestBuilder) (*SubscriptionEnableResponse, error) {
	if builder == nil {
		return nil, ErrBuilderRequired
	}

	req := builder.Build()
	path := subscriptionBasePath + "/enable"

	resp, err := net.Post[SubscriptionEnableRequest, struct {
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

	response := SubscriptionEnableResponse(*resp)
	return &response, nil
}
