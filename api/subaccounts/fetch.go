package subaccounts

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
)

// Fetch retrieves a specific subaccount by ID or code
func (c *Client) Fetch(ctx context.Context, idOrCode string) (*SubaccountFetchResponse, error) {
	if idOrCode == "" {
		return nil, fmt.Errorf("id_or_code is required")
	}

	endpoint := fmt.Sprintf("%s/%s", subaccountBasePath, idOrCode)

	return net.Get[Subaccount](ctx, c.client, c.secret, endpoint, c.baseURL)
}
