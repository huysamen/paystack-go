package virtualterminal

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// Update updates a virtual terminal
func (c *Client) Update(ctx context.Context, code string, builder *UpdateVirtualTerminalRequestBuilder) (*types.Response[VirtualTerminal], error) {
	return net.Put[UpdateVirtualTerminalRequest, VirtualTerminal](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s", basePath, code), builder.Build(), c.BaseURL)
}
