package dedicatedvirtualaccount

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type DeactivateResponse = types.Response[DedicatedVirtualAccount]

// Deactivate deactivates a dedicated virtual account on your integration
func (c *Client) Deactivate(ctx context.Context, dedicatedAccountID string) (*DeactivateResponse, error) {
	endpoint := fmt.Sprintf("%s/%s", basePath, dedicatedAccountID)

	return net.Delete[DedicatedVirtualAccount](ctx, c.Client, c.Secret, endpoint, c.BaseURL)
}
