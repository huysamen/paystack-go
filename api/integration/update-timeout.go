package integration

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// UpdateTimeoutRequest represents the request to update payment session timeout
type UpdateTimeoutRequest struct {
	Timeout int `json:"timeout"`
}

// UpdateTimeoutRequestBuilder provides a fluent interface for building UpdateTimeoutRequest
type UpdateTimeoutRequestBuilder struct {
	req *UpdateTimeoutRequest
}

// NewUpdateTimeoutRequest creates a new builder for UpdateTimeoutRequest
func NewUpdateTimeoutRequest(timeout int) *UpdateTimeoutRequestBuilder {
	return &UpdateTimeoutRequestBuilder{
		req: &UpdateTimeoutRequest{
			Timeout: timeout,
		},
	}
}

// Build returns the constructed UpdateTimeoutRequest
func (b *UpdateTimeoutRequestBuilder) Build() *UpdateTimeoutRequest {
	return b.req
}

// UpdateTimeoutResponse represents the response from updating payment session timeout
type UpdateTimeoutResponse = types.Response[UpdateTimeoutData]

// UpdateTimeoutData contains the updated payment session timeout information
type UpdateTimeoutData struct {
	PaymentSessionTimeout int `json:"payment_session_timeout"`
}

// UpdateTimeout updates the payment session timeout on your integration
func (c *Client) UpdateTimeout(ctx context.Context, builder *UpdateTimeoutRequestBuilder) (*UpdateTimeoutResponse, error) {
	return net.Put[UpdateTimeoutRequest, UpdateTimeoutData](ctx, c.Client, c.Secret, basePath+"/payment_session_timeout", builder.Build(), c.BaseURL)
}
