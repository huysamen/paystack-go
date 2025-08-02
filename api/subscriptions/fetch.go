package subscriptions

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// SubscriptionFetchResponse represents the response from fetching a subscription
type SubscriptionFetchResponse = types.Response[SubscriptionWithInvoices]

func (c *Client) Fetch(ctx context.Context, idOrCode string) (*SubscriptionFetchResponse, error) {
	return net.Get[SubscriptionWithInvoices](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s", basePath, idOrCode), c.BaseURL)
}
