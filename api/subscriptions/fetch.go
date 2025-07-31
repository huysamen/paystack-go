package subscriptions

import (
	"context"
	"errors"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

func (c *Client) Fetch(ctx context.Context, idOrCode string) (*types.Response[SubscriptionWithInvoices], error) {
	if idOrCode == "" {
		return nil, errors.New("subscription ID or code is required")
	}

	path := fmt.Sprintf("%s/%s", subscriptionBasePath, idOrCode)

	return net.Get[SubscriptionWithInvoices](
		ctx,
		c.client,
		c.secret,
		path,
		c.baseURL,
	)
}
