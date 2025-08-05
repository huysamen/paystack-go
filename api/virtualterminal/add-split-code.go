package virtualterminal

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type addSplitCodeRequest struct {
	SplitCode string `json:"split_code"`
}

type AddSplitCodeRequestBuilder struct {
	splitCode string
}

func NewAddSplitCodeRequestBuilder(splitCode string) *AddSplitCodeRequestBuilder {
	return &AddSplitCodeRequestBuilder{
		splitCode: splitCode,
	}
}

func (b *AddSplitCodeRequestBuilder) SplitCode(splitCode string) *AddSplitCodeRequestBuilder {
	b.splitCode = splitCode

	return b
}

func (b *AddSplitCodeRequestBuilder) Build() *addSplitCodeRequest {
	return &addSplitCodeRequest{
		SplitCode: b.splitCode,
	}
}

type AddSplitCodeResponseData = any
type AddSplitCodeResponse = types.Response[AddSplitCodeResponseData]

func (c *Client) AddSplitCode(ctx context.Context, code string, builder AddSplitCodeRequestBuilder) (*AddSplitCodeResponse, error) {
	return net.Put[addSplitCodeRequest, AddSplitCodeResponseData](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s/split_code", basePath, code), builder.Build(), c.BaseURL)
}
