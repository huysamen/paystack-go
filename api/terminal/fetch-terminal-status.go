package terminal

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type TerminalPresenceResponse = types.Response[types.TerminalPresenceStatus]

func (c *Client) FetchTerminalStatus(ctx context.Context, terminalID string) (*TerminalPresenceResponse, error) {
	endpoint := fmt.Sprintf("%s/%s/presence", basePath, terminalID)

	return net.Get[types.TerminalPresenceStatus](ctx, c.Client, c.Secret, endpoint, "", c.BaseURL)
}
