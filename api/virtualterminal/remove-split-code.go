package virtualterminal

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type RemoveSplitCodeRequest struct {
	SplitCode string `json:"split_code"`
}

type RemoveSplitCodeRequestBuilder struct {
	splitCode string
}

func NewRemoveSplitCodeRequest(splitCode string) *RemoveSplitCodeRequestBuilder {
	return &RemoveSplitCodeRequestBuilder{
		splitCode: splitCode,
	}
}

func (b *RemoveSplitCodeRequestBuilder) SplitCode(splitCode string) *RemoveSplitCodeRequestBuilder {
	b.splitCode = splitCode

	return b
}

func (b *RemoveSplitCodeRequestBuilder) Build() *RemoveSplitCodeRequest {
	return &RemoveSplitCodeRequest{
		SplitCode: b.splitCode,
	}
}

type RemoveSplitCodeResponse = types.Response[any]

func (c *Client) RemoveSplitCode(ctx context.Context, code string, builder *RemoveSplitCodeRequestBuilder) (*RemoveSplitCodeResponse, error) {
	endpoint := fmt.Sprintf("%s/%s/split_code", basePath, code)

	return net.DeleteWithBody[RemoveSplitCodeRequest, any](ctx, c.Client, c.Secret, endpoint, builder.Build(), c.BaseURL)
}
