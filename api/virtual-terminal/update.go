package virtualterminal

import (
	"context"
	"errors"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// Update updates a virtual terminal
func (c *Client) Update(ctx context.Context, code string, builder *UpdateVirtualTerminalRequestBuilder) (*types.Response[VirtualTerminal], error) {
	if code == "" {
		return nil, errors.New("virtual terminal code is required")
	}
	if builder == nil {
		return nil, ErrBuilderRequired
	}

	req := builder.Build()
	endpoint := fmt.Sprintf("%s/%s", virtualTerminalBasePath, code)
	return net.Put[UpdateVirtualTerminalRequest, VirtualTerminal](
		ctx, c.client, c.secret, endpoint, req, c.baseURL,
	)
}
