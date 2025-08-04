package terminal

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// TerminalUpdateRequest represents the request to update a terminal
type TerminalUpdateRequest struct {
	Name    *string `json:"name,omitempty"`    // Name of the terminal (optional)
	Address *string `json:"address,omitempty"` // Address of the terminal (optional)
}

// TerminalUpdateRequestBuilder provides a fluent interface for building TerminalUpdateRequest
type TerminalUpdateRequestBuilder struct {
	req *TerminalUpdateRequest
}

// NewTerminalUpdateRequest creates a new builder for TerminalUpdateRequest
func NewTerminalUpdateRequest() *TerminalUpdateRequestBuilder {
	return &TerminalUpdateRequestBuilder{
		req: &TerminalUpdateRequest{},
	}
}

// Name sets the terminal name
func (b *TerminalUpdateRequestBuilder) Name(name string) *TerminalUpdateRequestBuilder {
	b.req.Name = &name

	return b
}

// Address sets the terminal address
func (b *TerminalUpdateRequestBuilder) Address(address string) *TerminalUpdateRequestBuilder {
	b.req.Address = &address

	return b
}

// Build returns the constructed TerminalUpdateRequest
func (b *TerminalUpdateRequestBuilder) Build() *TerminalUpdateRequest {
	return b.req
}

// TerminalUpdateResponse represents the response from updating a terminal
type TerminalUpdateResponse = types.Response[types.Terminal]

// Update updates a terminal's details
func (c *Client) Update(ctx context.Context, terminalID string, builder *TerminalUpdateRequestBuilder) (*TerminalUpdateResponse, error) {
	req := builder.Build()
	endpoint := fmt.Sprintf("%s/%s", basePath, terminalID)
	return net.Put[TerminalUpdateRequest, types.Terminal](
		ctx, c.Client, c.Secret, endpoint, req, c.BaseURL,
	)
}
