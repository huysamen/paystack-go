package virtualterminal

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// Create creates a new virtual terminal
func (c *Client) Create(ctx context.Context, builder *CreateVirtualTerminalRequestBuilder) (*types.Response[VirtualTerminal], error) {
	return net.Post[CreateVirtualTerminalRequest, VirtualTerminal](ctx, c.Client, c.Secret, basePath, builder.Build(), c.BaseURL)
}
