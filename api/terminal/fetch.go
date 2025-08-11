package terminal

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type FetchResponseData = types.Terminal
type FetchResponse = types.Response[FetchResponseData]

func (c *Client) Fetch(ctx context.Context, terminalID string) (*FetchResponse, error) {
	return net.Get[FetchResponseData](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s", basePath, terminalID), c.BaseURL)
}
