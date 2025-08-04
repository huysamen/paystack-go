package virtualterminal

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type DeactivateVirtualTerminalResponse = types.Response[any]

func (c *Client) Deactivate(ctx context.Context, code string) (*DeactivateVirtualTerminalResponse, error) {
	return net.Put[any, any](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s/deactivate", basePath, code), nil, c.BaseURL)
}
