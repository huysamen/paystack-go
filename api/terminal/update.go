package terminal

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
)

// Update updates a terminal's details
func (c *Client) Update(ctx context.Context, terminalID string, builder *TerminalUpdateRequestBuilder) (*TerminalUpdateResponse, error) {
	req := builder.Build()
	endpoint := fmt.Sprintf("%s/%s", basePath, terminalID)
	return net.Put[TerminalUpdateRequest, Terminal](
		ctx, c.Client, c.Secret, endpoint, req, c.BaseURL,
	)
}
