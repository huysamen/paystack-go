package virtualterminal

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// CreateVirtualTerminalRequest represents the request to create a virtual terminal
type CreateVirtualTerminalRequest struct {
	Name         string                             `json:"name"`
	Destinations []types.VirtualTerminalDestination `json:"destinations,omitempty"`
	Metadata     *types.Metadata                    `json:"metadata,omitempty"`
	Currency     string                             `json:"currency,omitempty"`
	CustomFields []types.VirtualTerminalCustomField `json:"custom_fields,omitempty"`
}

// CreateVirtualTerminalRequestBuilder provides a fluent interface for building CreateVirtualTerminalRequest
type CreateVirtualTerminalRequestBuilder struct {
	req *CreateVirtualTerminalRequest
}

// NewCreateVirtualTerminalRequest creates a new builder for CreateVirtualTerminalRequest
func NewCreateVirtualTerminalRequest(name string) *CreateVirtualTerminalRequestBuilder {
	return &CreateVirtualTerminalRequestBuilder{
		req: &CreateVirtualTerminalRequest{
			Name: name,
		},
	}
}

// Build returns the constructed CreateVirtualTerminalRequest
func (b *CreateVirtualTerminalRequestBuilder) Build() *CreateVirtualTerminalRequest {
	return b.req
}

// CreateVirtualTerminalResponse represents the response from creating a virtual terminal
type CreateVirtualTerminalResponse = types.Response[types.VirtualTerminal]

// Create creates a new virtual terminal
func (c *Client) Create(ctx context.Context, builder *CreateVirtualTerminalRequestBuilder) (*CreateVirtualTerminalResponse, error) {
	return net.Post[CreateVirtualTerminalRequest, types.VirtualTerminal](ctx, c.Client, c.Secret, basePath, builder.Build(), c.BaseURL)
}
