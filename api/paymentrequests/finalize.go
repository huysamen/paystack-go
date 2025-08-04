package paymentrequests

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type finalizeRequest struct {
	SendNotification *bool `json:"send_notification,omitempty"`
}

type FinalizeRequestBuilder struct {
	req *finalizeRequest
}

func NewFinalizeRequestBuilder() *FinalizeRequestBuilder {
	return &FinalizeRequestBuilder{
		req: &finalizeRequest{},
	}
}

func (b *FinalizeRequestBuilder) SendNotification(sendNotification bool) *FinalizeRequestBuilder {
	b.req.SendNotification = &sendNotification

	return b
}

func (b *FinalizeRequestBuilder) Build() *finalizeRequest {
	return b.req
}

type FinalizeResponseData = types.PaymentRequest
type FinalizeResponse = types.Response[FinalizeResponseData]

func (c *Client) Finalize(ctx context.Context, code string, builder FinalizeRequestBuilder) (*FinalizeResponse, error) {
	return net.Post[finalizeRequest, FinalizeResponseData](ctx, c.Client, c.Secret, basePath+"/finalize/"+code, builder.Build(), c.BaseURL)
}
