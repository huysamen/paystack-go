package subscriptions

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type disableRequest struct {
	Code  string `json:"code"`  // Subscription code
	Token string `json:"token"` // Email token
}

type DisableRequestBuilder struct {
	req *disableRequest
}

func NewDisableRequestBuilder(code, token string) *DisableRequestBuilder {
	return &DisableRequestBuilder{
		req: &disableRequest{
			Code:  code,
			Token: token,
		},
	}
}

func (b *DisableRequestBuilder) Build() *disableRequest {
	return b.req
}

type DisableResponseData = any
type DisableResponse = types.Response[DisableResponseData]

func (c *Client) Disable(ctx context.Context, builder DisableRequestBuilder) (*DisableResponse, error) {
	return net.Post[disableRequest, DisableResponseData](ctx, c.Client, c.Secret, basePath+"/disable", builder.Build(), c.BaseURL)
}
