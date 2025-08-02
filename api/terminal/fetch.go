package terminal

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
)

// Fetch retrieves a terminal by ID
func (c *Client) Fetch(ctx context.Context, terminalID string) (*TerminalFetchResponse, error) {
	return net.Get[Terminal](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s", basePath, terminalID), c.BaseURL)
}
