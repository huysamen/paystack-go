package terminal

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type TerminalFetchResponse = types.Response[types.Terminal]

func (c *Client) Fetch(ctx context.Context, terminalID string) (*TerminalFetchResponse, error) {
	return net.Get[types.Terminal](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s", basePath, terminalID), c.BaseURL)
}
