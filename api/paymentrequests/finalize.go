package paymentrequests

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type FinalizePaymentRequestRequest struct {
	SendNotification *bool `json:"send_notification,omitempty"`
}

type FinalizePaymentRequestRequestBuilder struct {
	req *FinalizePaymentRequestRequest
}

func NewFinalizePaymentRequestRequest() *FinalizePaymentRequestRequestBuilder {
	return &FinalizePaymentRequestRequestBuilder{
		req: &FinalizePaymentRequestRequest{},
	}
}

func (b *FinalizePaymentRequestRequestBuilder) SendNotification(sendNotification bool) *FinalizePaymentRequestRequestBuilder {
	b.req.SendNotification = &sendNotification

	return b
}

func (b *FinalizePaymentRequestRequestBuilder) Build() *FinalizePaymentRequestRequest {
	return b.req
}

type FinalizePaymentRequestResponse = types.Response[types.PaymentRequest]

func (c *Client) Finalize(ctx context.Context, code string, builder *FinalizePaymentRequestRequestBuilder) (*FinalizePaymentRequestResponse, error) {
	var req *FinalizePaymentRequestRequest
	if builder != nil {
		req = builder.Build()
	} else {
		req = &FinalizePaymentRequestRequest{}
	}

	return net.Post[FinalizePaymentRequestRequest, types.PaymentRequest](ctx, c.Client, c.Secret, basePath+"/finalize/"+code, req, c.BaseURL)
}
