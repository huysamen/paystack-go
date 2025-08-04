package virtualterminal

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// FetchVirtualTerminalResponse represents the response from fetching a virtual terminal
type FetchVirtualTerminalResponse = types.Response[types.VirtualTerminal]

// Fetch retrieves a specific virtual terminal by code
func (c *Client) Fetch(ctx context.Context, code string) (*FetchVirtualTerminalResponse, error) {
	return net.Get[types.VirtualTerminal](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s", basePath, code), c.BaseURL)
}
