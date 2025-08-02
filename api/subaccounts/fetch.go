package subaccounts

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
)

// Fetch retrieves a specific subaccount by ID or code
func (c *Client) Fetch(ctx context.Context, idOrCode string) (*SubaccountFetchResponse, error) {
	return net.Get[Subaccount](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s", basePath, idOrCode), c.BaseURL)
}
