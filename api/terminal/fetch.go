package terminal

import (
	"context"
	"errors"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// Fetch retrieves a terminal by ID
func (c *Client) Fetch(ctx context.Context, terminalID string) (*types.Response[Terminal], error) {
	if terminalID == "" {
		return nil, errors.New("terminal ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", terminalBasePath, terminalID)
	return net.Get[Terminal](ctx, c.client, c.secret, endpoint, c.baseURL)
}
