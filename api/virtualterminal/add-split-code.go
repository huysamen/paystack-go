package virtualterminal

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// AddSplitCodeRequest represents the request to add a split code to a virtual terminal
type AddSplitCodeRequest struct {
	SplitCode string `json:"split_code"`
}

// AddSplitCodeRequestBuilder provides a fluent interface for building AddSplitCodeRequest
type AddSplitCodeRequestBuilder struct {
	splitCode string
}

// NewAddSplitCodeRequest creates a new builder for adding a split code
func NewAddSplitCodeRequest(splitCode string) *AddSplitCodeRequestBuilder {
	return &AddSplitCodeRequestBuilder{
		splitCode: splitCode,
	}
}

// SplitCode sets the split code
func (b *AddSplitCodeRequestBuilder) SplitCode(splitCode string) *AddSplitCodeRequestBuilder {
	b.splitCode = splitCode
	return b
}

// Build creates the AddSplitCodeRequest
func (b *AddSplitCodeRequestBuilder) Build() *AddSplitCodeRequest {
	return &AddSplitCodeRequest{
		SplitCode: b.splitCode,
	}
}

// AddSplitCodeResponse represents the response from adding a split code
type AddSplitCodeResponse = types.Response[any]

// AddSplitCode adds a split code to a virtual terminal
func (c *Client) AddSplitCode(ctx context.Context, code string, builder *AddSplitCodeRequestBuilder) (*types.Response[any], error) {
	return net.Put[AddSplitCodeRequest, any](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s/split_code", basePath, code), builder.Build(), c.BaseURL)
}
