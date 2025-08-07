package integration

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type updateTimeoutRequest struct {
	Timeout int `json:"timeout"`
}

type UpdateTimeoutRequestBuilder struct {
	req *updateTimeoutRequest
}

func NewUpdateTimeoutRequestBuilder(timeout int) *UpdateTimeoutRequestBuilder {
	return &UpdateTimeoutRequestBuilder{
		req: &updateTimeoutRequest{
			Timeout: timeout,
		},
	}
}

func (b *UpdateTimeoutRequestBuilder) Build() *updateTimeoutRequest {
	return b.req
}

type UpdateTimeoutResponseData struct {
	PaymentSessionTimeout int `json:"payment_session_timeout"`
}

type UpdateTimeoutResponse = types.Response[UpdateTimeoutResponseData]

func (c *Client) UpdateTimeout(ctx context.Context, builder UpdateTimeoutRequestBuilder) (*UpdateTimeoutResponse, error) {
	return net.Put[updateTimeoutRequest, UpdateTimeoutResponseData](ctx, c.Client, c.Secret, basePath+"/payment_session_timeout", builder.Build(), c.BaseURL)
}
