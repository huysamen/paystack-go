package dedicatedvirtualaccount

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type FetchDedicatedVirtualAccountResponse = types.Response[DedicatedVirtualAccount]

// Fetch gets details of a dedicated virtual account on your integration
func (c *Client) Fetch(ctx context.Context, dedicatedAccountID string) (*FetchDedicatedVirtualAccountResponse, error) {
	endpoint := fmt.Sprintf("%s/%s", basePath, dedicatedAccountID)

	return net.Get[DedicatedVirtualAccount](ctx, c.Client, c.Secret, endpoint, c.BaseURL)
}
