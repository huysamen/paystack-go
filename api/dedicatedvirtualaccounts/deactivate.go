package dedicatedvirtualaccount

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type DeactivateResponseData = types.DedicatedVirtualAccount
type DeactivateResponse = types.Response[DeactivateResponseData]

func (c *Client) Deactivate(ctx context.Context, dedicatedAccountID string) (*DeactivateResponse, error) {
	endpoint := fmt.Sprintf("%s/%s", basePath, dedicatedAccountID)

	return net.Delete[DeactivateResponseData](ctx, c.Client, c.Secret, endpoint, c.BaseURL)
}
