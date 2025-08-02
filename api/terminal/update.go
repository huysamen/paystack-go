package terminal

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// Update updates a terminal's details
func (c *Client) Update(ctx context.Context, terminalID string, builder *TerminalUpdateRequestBuilder) (*types.Response[Terminal], error) {
	req := builder.Build()
	endpoint := fmt.Sprintf("%s/%s", basePath, terminalID)
	return net.Put[TerminalUpdateRequest, Terminal](
		ctx, c.Client, c.Secret, endpoint, req, c.BaseURL,
	)
}
