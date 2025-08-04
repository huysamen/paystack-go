package terminal

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type TerminalUpdateRequest struct {
	Name    *string `json:"name,omitempty"`    // Name of the terminal (optional)
	Address *string `json:"address,omitempty"` // Address of the terminal (optional)
}

type TerminalUpdateRequestBuilder struct {
	req *TerminalUpdateRequest
}

func NewTerminalUpdateRequest() *TerminalUpdateRequestBuilder {
	return &TerminalUpdateRequestBuilder{
		req: &TerminalUpdateRequest{},
	}
}

func (b *TerminalUpdateRequestBuilder) Name(name string) *TerminalUpdateRequestBuilder {
	b.req.Name = &name

	return b
}

func (b *TerminalUpdateRequestBuilder) Address(address string) *TerminalUpdateRequestBuilder {
	b.req.Address = &address

	return b
}

func (b *TerminalUpdateRequestBuilder) Build() *TerminalUpdateRequest {
	return b.req
}

type TerminalUpdateResponse = types.Response[types.Terminal]

func (c *Client) Update(ctx context.Context, terminalID string, builder *TerminalUpdateRequestBuilder) (*TerminalUpdateResponse, error) {
	req := builder.Build()
	endpoint := fmt.Sprintf("%s/%s", basePath, terminalID)
	return net.Put[TerminalUpdateRequest, types.Terminal](
		ctx, c.Client, c.Secret, endpoint, req, c.BaseURL,
	)
}
