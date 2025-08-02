package subscriptions

import (
	"context"
	"time"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// SubscriptionCreateRequest represents the request to create a subscription
type SubscriptionCreateRequest struct {
	Customer      string     `json:"customer"`                // Customer email or customer code
	Plan          string     `json:"plan"`                    // Plan code
	Authorization *string    `json:"authorization,omitempty"` // Authorization code if customer has multiple
	StartDate     *time.Time `json:"start_date,omitempty"`    // Date for first debit (ISO 8601)
}

// SubscriptionCreateRequestBuilder provides a fluent interface for building SubscriptionCreateRequest
type SubscriptionCreateRequestBuilder struct {
	req *SubscriptionCreateRequest
}

// NewSubscriptionCreateRequest creates a new builder for SubscriptionCreateRequest
func NewSubscriptionCreateRequest(customer, plan string) *SubscriptionCreateRequestBuilder {
	return &SubscriptionCreateRequestBuilder{
		req: &SubscriptionCreateRequest{
			Customer: customer,
			Plan:     plan,
		},
	}
}

// Authorization sets the authorization code
func (b *SubscriptionCreateRequestBuilder) Authorization(authorization string) *SubscriptionCreateRequestBuilder {
	b.req.Authorization = &authorization
	return b
}

// StartDate sets the start date
func (b *SubscriptionCreateRequestBuilder) StartDate(startDate time.Time) *SubscriptionCreateRequestBuilder {
	b.req.StartDate = &startDate
	return b
}

// Build returns the constructed SubscriptionCreateRequest
func (b *SubscriptionCreateRequestBuilder) Build() *SubscriptionCreateRequest {
	return b.req
}

// SubscriptionCreateResponse represents the response from creating a subscription.
type SubscriptionCreateResponse = types.Response[Subscription]

func (c *Client) Create(ctx context.Context, builder *SubscriptionCreateRequestBuilder) (*SubscriptionCreateResponse, error) {
	return net.Post[SubscriptionCreateRequest, Subscription](ctx, c.Client, c.Secret, basePath, builder.Build(), c.BaseURL)
}
