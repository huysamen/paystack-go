package subaccounts

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// SubaccountFetchResponse represents the response from fetching a subaccount
type SubaccountFetchResponse = types.Response[types.Subaccount]

// Fetch retrieves a specific subaccount by ID or code
func (c *Client) Fetch(ctx context.Context, idOrCode string) (*SubaccountFetchResponse, error) {
	return net.Get[types.Subaccount](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s", basePath, idOrCode), c.BaseURL)
}
