package subscriptions

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type enableRequest struct {
	Code  string `json:"code"`  // Subscription code
	Token string `json:"token"` // Email token
}

type EnableRequestBuilder struct {
	req *enableRequest
}

func NewEnableRequestBuilder(code, token string) *EnableRequestBuilder {
	return &EnableRequestBuilder{
		req: &enableRequest{
			Code:  code,
			Token: token,
		},
	}
}

func (b *EnableRequestBuilder) Build() *enableRequest {
	return b.req
}

type EnableResponseData = any
type EnableResponse = types.Response[EnableResponseData]

func (c *Client) Enable(ctx context.Context, builder EnableRequestBuilder) (*EnableResponse, error) {
	return net.Post[enableRequest, EnableResponseData](ctx, c.Client, c.Secret, basePath+"/enable", builder.Build(), c.BaseURL)
}
