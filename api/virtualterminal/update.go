package virtualterminal

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// UpdateVirtualTerminalRequest represents the request to update a virtual terminal
type UpdateVirtualTerminalRequest struct {
	Name string `json:"name"`
}

// UpdateVirtualTerminalRequestBuilder provides a fluent interface for building UpdateVirtualTerminalRequest
type UpdateVirtualTerminalRequestBuilder struct {
	name string
}

// NewUpdateVirtualTerminalRequest creates a new builder for updating a virtual terminal
func NewUpdateVirtualTerminalRequest(name string) *UpdateVirtualTerminalRequestBuilder {
	return &UpdateVirtualTerminalRequestBuilder{
		name: name,
	}
}

// Name sets the name of the virtual terminal
func (b *UpdateVirtualTerminalRequestBuilder) Name(name string) *UpdateVirtualTerminalRequestBuilder {
	b.name = name
	return b
}

// Build creates the UpdateVirtualTerminalRequest
func (b *UpdateVirtualTerminalRequestBuilder) Build() *UpdateVirtualTerminalRequest {
	return &UpdateVirtualTerminalRequest{
		Name: b.name,
	}
}

// UpdateVirtualTerminalResponse represents the response from updating a virtual terminal
type UpdateVirtualTerminalResponse = types.Response[types.VirtualTerminal]

// Update updates a virtual terminal
func (c *Client) Update(ctx context.Context, code string, builder *UpdateVirtualTerminalRequestBuilder) (*types.Response[types.VirtualTerminal], error) {
	return net.Put[UpdateVirtualTerminalRequest, types.VirtualTerminal](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s", basePath, code), builder.Build(), c.BaseURL)
}
