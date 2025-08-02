package terminal

import (
	"context"
	"errors"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// Update updates a terminal's details
func (c *Client) Update(ctx context.Context, terminalID string, builder *TerminalUpdateRequestBuilder) (*types.Response[Terminal], error) {
	if terminalID == "" {
		return nil, errors.New("terminal ID is required")
	}
	if builder == nil {
		return nil, ErrBuilderRequired
	}

	req := builder.Build()
	endpoint := fmt.Sprintf("%s/%s", terminalBasePath, terminalID)
	return net.Put[TerminalUpdateRequest, Terminal](
		ctx, c.client, c.secret, endpoint, req, c.baseURL,
	)
}
