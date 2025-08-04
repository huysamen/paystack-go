package integration

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type UpdateTimeoutRequest struct {
	Timeout int `json:"timeout"`
}

type UpdateTimeoutRequestBuilder struct {
	req *UpdateTimeoutRequest
}

func NewUpdateTimeoutRequest(timeout int) *UpdateTimeoutRequestBuilder {
	return &UpdateTimeoutRequestBuilder{
		req: &UpdateTimeoutRequest{
			Timeout: timeout,
		},
	}
}

func (b *UpdateTimeoutRequestBuilder) Build() *UpdateTimeoutRequest {
	return b.req
}

type UpdateTimeoutResponse = types.Response[UpdateTimeoutData]

type UpdateTimeoutData struct {
	PaymentSessionTimeout int `json:"payment_session_timeout"`
}

func (c *Client) UpdateTimeout(ctx context.Context, builder *UpdateTimeoutRequestBuilder) (*UpdateTimeoutResponse, error) {
	return net.Put[UpdateTimeoutRequest, UpdateTimeoutData](ctx, c.Client, c.Secret, basePath+"/payment_session_timeout", builder.Build(), c.BaseURL)
}
