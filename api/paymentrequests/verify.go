package paymentrequests

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// VerifyPaymentRequestResponse represents the response from verifying a payment request
type VerifyPaymentRequestResponse = types.Response[PaymentRequest]

// Verify verifies details of a payment request on your integration
func (c *Client) Verify(ctx context.Context, code string) (*VerifyPaymentRequestResponse, error) {
	return net.Get[PaymentRequest](ctx, c.Client, c.Secret, basePath+"/verify/"+code, c.BaseURL)
}
