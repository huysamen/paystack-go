package virtualterminal

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// RemoveSplitCodeRequest represents the request to remove a split code from a virtual terminal
type RemoveSplitCodeRequest struct {
	SplitCode string `json:"split_code"`
}

// RemoveSplitCodeRequestBuilder provides a fluent interface for building RemoveSplitCodeRequest
type RemoveSplitCodeRequestBuilder struct {
	splitCode string
}

// NewRemoveSplitCodeRequest creates a new builder for removing a split code
func NewRemoveSplitCodeRequest(splitCode string) *RemoveSplitCodeRequestBuilder {
	return &RemoveSplitCodeRequestBuilder{
		splitCode: splitCode,
	}
}

// SplitCode sets the split code to remove
func (b *RemoveSplitCodeRequestBuilder) SplitCode(splitCode string) *RemoveSplitCodeRequestBuilder {
	b.splitCode = splitCode

	return b
}

// Build creates the RemoveSplitCodeRequest
func (b *RemoveSplitCodeRequestBuilder) Build() *RemoveSplitCodeRequest {
	return &RemoveSplitCodeRequest{
		SplitCode: b.splitCode,
	}
}

// RemoveSplitCodeResponse represents the response from removing a split code
type RemoveSplitCodeResponse = types.Response[any]

// RemoveSplitCode removes a split code from a virtual terminal
func (c *Client) RemoveSplitCode(ctx context.Context, code string, builder *RemoveSplitCodeRequestBuilder) (*RemoveSplitCodeResponse, error) {
	endpoint := fmt.Sprintf("%s/%s/split_code", basePath, code)

	return net.DeleteWithBody[RemoveSplitCodeRequest, any](ctx, c.Client, c.Secret, endpoint, builder.Build(), c.BaseURL)
}
