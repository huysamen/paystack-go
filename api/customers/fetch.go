package customers

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// CustomerWithRelations represents a customer with related data
type CustomerWithRelations struct {
	types.Customer
	Subscriptions  []types.Subscription  `json:"subscriptions"`
	Authorizations []types.Authorization `json:"authorizations"`
	Transactions   []types.Transaction   `json:"transactions"`
}

// Fetch retrieves a customer by email or customer code
func (c *Client) Fetch(ctx context.Context, emailOrCode string) (*types.Response[CustomerWithRelations], error) {
	path := fmt.Sprintf("%s/%s", basePath, emailOrCode)

	return net.Get[CustomerWithRelations](ctx, c.Client, c.Secret, path, c.BaseURL)
}
