package virtualterminal

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type FetchResponseData = types.VirtualTerminal
type FetchResponse = types.Response[FetchResponseData]

func (c *Client) Fetch(ctx context.Context, code string) (*FetchResponse, error) {
	return net.Get[FetchResponseData](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s", basePath, code), c.BaseURL)
}
