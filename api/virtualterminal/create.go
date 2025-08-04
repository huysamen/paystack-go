package virtualterminal

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type CreateVirtualTerminalRequest struct {
	Name         string                             `json:"name"`
	Destinations []types.VirtualTerminalDestination `json:"destinations,omitempty"`
	Metadata     *types.Metadata                    `json:"metadata,omitempty"`
	Currency     string                             `json:"currency,omitempty"`
	CustomFields []types.VirtualTerminalCustomField `json:"custom_fields,omitempty"`
}

type CreateVirtualTerminalRequestBuilder struct {
	req *CreateVirtualTerminalRequest
}

func NewCreateVirtualTerminalRequest(name string) *CreateVirtualTerminalRequestBuilder {
	return &CreateVirtualTerminalRequestBuilder{
		req: &CreateVirtualTerminalRequest{
			Name: name,
		},
	}
}

func (b *CreateVirtualTerminalRequestBuilder) Build() *CreateVirtualTerminalRequest {
	return b.req
}

type CreateVirtualTerminalResponse = types.Response[types.VirtualTerminal]

func (c *Client) Create(ctx context.Context, builder *CreateVirtualTerminalRequestBuilder) (*CreateVirtualTerminalResponse, error) {
	return net.Post[CreateVirtualTerminalRequest, types.VirtualTerminal](ctx, c.Client, c.Secret, basePath, builder.Build(), c.BaseURL)
}
