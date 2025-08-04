package dedicatedvirtualaccount

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type FetchDedicatedVirtualAccountResponse = types.Response[types.DedicatedVirtualAccount]

func (c *Client) Fetch(ctx context.Context, dedicatedAccountID string) (*FetchDedicatedVirtualAccountResponse, error) {
	endpoint := fmt.Sprintf("%s/%s", basePath, dedicatedAccountID)

	return net.Get[types.DedicatedVirtualAccount](ctx, c.Client, c.Secret, endpoint, c.BaseURL)
}
