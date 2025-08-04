package subscriptions

import (
	"context"
	"time"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type SubscriptionCreateRequest struct {
	Customer      string     `json:"customer"`                // Customer email or customer code
	Plan          string     `json:"plan"`                    // Plan code
	Authorization *string    `json:"authorization,omitempty"` // Authorization code if customer has multiple
	StartDate     *time.Time `json:"start_date,omitempty"`    // Date for first debit (ISO 8601)
}

type SubscriptionCreateRequestBuilder struct {
	req *SubscriptionCreateRequest
}

func NewSubscriptionCreateRequest(customer, plan string) *SubscriptionCreateRequestBuilder {
	return &SubscriptionCreateRequestBuilder{
		req: &SubscriptionCreateRequest{
			Customer: customer,
			Plan:     plan,
		},
	}
}

func (b *SubscriptionCreateRequestBuilder) Authorization(authorization string) *SubscriptionCreateRequestBuilder {
	b.req.Authorization = &authorization

	return b
}

func (b *SubscriptionCreateRequestBuilder) StartDate(startDate time.Time) *SubscriptionCreateRequestBuilder {
	b.req.StartDate = &startDate

	return b
}

func (b *SubscriptionCreateRequestBuilder) Build() *SubscriptionCreateRequest {
	return b.req
}

type SubscriptionCreateResponse = types.Response[types.Subscription]

func (c *Client) Create(ctx context.Context, builder *SubscriptionCreateRequestBuilder) (*SubscriptionCreateResponse, error) {
	return net.Post[SubscriptionCreateRequest, types.Subscription](ctx, c.Client, c.Secret, basePath, builder.Build(), c.BaseURL)
}
