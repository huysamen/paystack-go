package dedicatedvirtualaccount

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type DeactivateResponse = types.Response[types.DedicatedVirtualAccount]

func (c *Client) Deactivate(ctx context.Context, dedicatedAccountID string) (*DeactivateResponse, error) {
	endpoint := fmt.Sprintf("%s/%s", basePath, dedicatedAccountID)

	return net.Delete[types.DedicatedVirtualAccount](ctx, c.Client, c.Secret, endpoint, c.BaseURL)
}
