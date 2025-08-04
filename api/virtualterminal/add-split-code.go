package virtualterminal

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type AddSplitCodeRequest struct {
	SplitCode string `json:"split_code"`
}

type AddSplitCodeRequestBuilder struct {
	splitCode string
}

func NewAddSplitCodeRequest(splitCode string) *AddSplitCodeRequestBuilder {
	return &AddSplitCodeRequestBuilder{
		splitCode: splitCode,
	}
}

func (b *AddSplitCodeRequestBuilder) SplitCode(splitCode string) *AddSplitCodeRequestBuilder {
	b.splitCode = splitCode

	return b
}

func (b *AddSplitCodeRequestBuilder) Build() *AddSplitCodeRequest {
	return &AddSplitCodeRequest{
		SplitCode: b.splitCode,
	}
}

type AddSplitCodeResponse = types.Response[any]

func (c *Client) AddSplitCode(ctx context.Context, code string, builder *AddSplitCodeRequestBuilder) (*AddSplitCodeResponse, error) {
	return net.Put[AddSplitCodeRequest, any](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s/split_code", basePath, code), builder.Build(), c.BaseURL)
}
