package virtualterminal

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// CreateVirtualTerminalResponse represents the response from creating a virtual terminal.
type CreateVirtualTerminalResponse types.Response[VirtualTerminal]

// Create creates a new virtual terminal
func (c *Client) Create(ctx context.Context, builder *CreateVirtualTerminalRequestBuilder) (*CreateVirtualTerminalResponse, error) {
	if builder == nil {
		return nil, ErrBuilderRequired
	}

	req := builder.Build()
	resp, err := net.Post[CreateVirtualTerminalRequest, VirtualTerminal](
		ctx, c.client, c.secret, virtualTerminalBasePath, req, c.baseURL,
	)
	if err != nil {
		return nil, err
	}

	response := CreateVirtualTerminalResponse(*resp)
	return &response, nil
}
