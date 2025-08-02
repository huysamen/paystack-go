package paymentrequests

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// FinalizePaymentRequestRequest represents the request to finalize a payment request
type FinalizePaymentRequestRequest struct {
	SendNotification *bool `json:"send_notification,omitempty"`
}

// FinalizePaymentRequestRequestBuilder provides a fluent interface for building FinalizePaymentRequestRequest
type FinalizePaymentRequestRequestBuilder struct {
	req *FinalizePaymentRequestRequest
}

// NewFinalizePaymentRequestRequest creates a new builder for FinalizePaymentRequestRequest
func NewFinalizePaymentRequestRequest() *FinalizePaymentRequestRequestBuilder {
	return &FinalizePaymentRequestRequestBuilder{
		req: &FinalizePaymentRequestRequest{},
	}
}

// SendNotification sets whether to send notification
func (b *FinalizePaymentRequestRequestBuilder) SendNotification(sendNotification bool) *FinalizePaymentRequestRequestBuilder {
	b.req.SendNotification = &sendNotification
	return b
}

// Build returns the constructed FinalizePaymentRequestRequest
func (b *FinalizePaymentRequestRequestBuilder) Build() *FinalizePaymentRequestRequest {
	return b.req
}

// FinalizePaymentRequestResponse represents the response from finalizing a payment request
type FinalizePaymentRequestResponse = types.Response[PaymentRequest]

// Finalize finalizes a draft payment request
func (c *Client) Finalize(ctx context.Context, code string, builder *FinalizePaymentRequestRequestBuilder) (*FinalizePaymentRequestResponse, error) {
	var req *FinalizePaymentRequestRequest
	if builder != nil {
		req = builder.Build()
	} else {
		req = &FinalizePaymentRequestRequest{}
	}

	return net.Post[FinalizePaymentRequestRequest, PaymentRequest](ctx, c.Client, c.Secret, basePath+"/finalize/"+code, req, c.BaseURL)
}
