package dedicatedvirtualaccount

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type FetchResponseData = types.DedicatedVirtualAccount
type FetchResponse = types.Response[FetchResponseData]

func (c *Client) Fetch(ctx context.Context, dedicatedAccountID string) (*FetchResponse, error) {
	endpoint := fmt.Sprintf("%s/%s", basePath, dedicatedAccountID)

	return net.Get[FetchResponseData](ctx, c.Client, c.Secret, endpoint, c.BaseURL)
}
