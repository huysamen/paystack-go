package virtualterminal

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type UpdateVirtualTerminalRequest struct {
	Name string `json:"name"`
}

type UpdateVirtualTerminalRequestBuilder struct {
	name string
}

func NewUpdateVirtualTerminalRequest(name string) *UpdateVirtualTerminalRequestBuilder {
	return &UpdateVirtualTerminalRequestBuilder{
		name: name,
	}
}

func (b *UpdateVirtualTerminalRequestBuilder) Name(name string) *UpdateVirtualTerminalRequestBuilder {
	b.name = name
	return b
}

func (b *UpdateVirtualTerminalRequestBuilder) Build() *UpdateVirtualTerminalRequest {
	return &UpdateVirtualTerminalRequest{
		Name: b.name,
	}
}

type UpdateVirtualTerminalResponse = types.Response[types.VirtualTerminal]

func (c *Client) Update(ctx context.Context, code string, builder *UpdateVirtualTerminalRequestBuilder) (*UpdateVirtualTerminalResponse, error) {
	return net.Put[UpdateVirtualTerminalRequest, types.VirtualTerminal](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s", basePath, code), builder.Build(), c.BaseURL)
}
