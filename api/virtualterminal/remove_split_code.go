package virtualterminal

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type removeSplitCodeRequest struct {
	SplitCode string `json:"split_code"`
}

type RemoveSplitCodeRequestBuilder struct {
	splitCode string
}

func NewRemoveSplitCodeRequestBuilder(splitCode string) *RemoveSplitCodeRequestBuilder {
	return &RemoveSplitCodeRequestBuilder{
		splitCode: splitCode,
	}
}

func (b *RemoveSplitCodeRequestBuilder) SplitCode(splitCode string) *RemoveSplitCodeRequestBuilder {
	b.splitCode = splitCode

	return b
}

func (b *RemoveSplitCodeRequestBuilder) Build() *removeSplitCodeRequest {
	return &removeSplitCodeRequest{
		SplitCode: b.splitCode,
	}
}

type RemoveSplitCodeResponseData = any
type RemoveSplitCodeResponse = types.Response[RemoveSplitCodeResponseData]

func (c *Client) RemoveSplitCode(ctx context.Context, code string, builder RemoveSplitCodeRequestBuilder) (*RemoveSplitCodeResponse, error) {
	endpoint := fmt.Sprintf("%s/%s/split_code", basePath, code)

	return net.DeleteWithBody[removeSplitCodeRequest, RemoveSplitCodeResponseData](ctx, c.Client, c.Secret, endpoint, builder.Build(), c.BaseURL)
}
