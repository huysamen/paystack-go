package subscriptions

import (
	"context"
	"errors"
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


func (c *Client) Create(ctx context.Context, req *SubscriptionCreateRequest) (*types.Response[Subscription], error) {
	if req == nil {
		return nil, errors.New("request cannot be nil")
	}


	return net.Post[SubscriptionCreateRequest, Subscription](
		ctx,
		c.client,
		c.secret,
		subscriptionBasePath,
		req,
		c.baseURL,
	)
}
