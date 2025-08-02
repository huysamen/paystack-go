package virtualterminal

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// Create creates a new virtual terminal
func (c *Client) Create(ctx context.Context, builder *CreateVirtualTerminalRequestBuilder) (*types.Response[VirtualTerminal], error) {
	if builder == nil {
		return nil, ErrBuilderRequired
	}

	req := builder.Build()
	return net.Post[CreateVirtualTerminalRequest, VirtualTerminal](
		ctx, c.client, c.secret, virtualTerminalBasePath, req, c.baseURL,
	)
}
