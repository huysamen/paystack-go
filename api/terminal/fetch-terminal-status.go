package terminal

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// TerminalPresenceResponse represents the response from checking terminal presence
type TerminalPresenceResponse = types.Response[types.TerminalPresenceStatus]

// FetchTerminalStatus fetches the status of a terminal
func (c *Client) FetchTerminalStatus(ctx context.Context, terminalID string) (*TerminalPresenceResponse, error) {
	endpoint := fmt.Sprintf("%s/%s/presence", basePath, terminalID)

	return net.Get[types.TerminalPresenceStatus](ctx, c.Client, c.Secret, endpoint, "", c.BaseURL)
}
