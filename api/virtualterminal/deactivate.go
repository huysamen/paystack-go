package virtualterminal

import (
	"context"
	"errors"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// Deactivate deactivates a virtual terminal
func (c *Client) Deactivate(ctx context.Context, code string) (*types.Response[any], error) {
	if code == "" {
		return nil, errors.New("virtual terminal code is required")
	}

	endpoint := fmt.Sprintf("%s/%s/deactivate", virtualTerminalBasePath, code)
	return net.Put[any, any](
		ctx, c.client, c.secret, endpoint, nil, c.baseURL,
	)
}
