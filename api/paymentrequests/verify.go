package paymentrequests

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type VerifyPaymentRequestResponse = types.Response[types.PaymentRequest]

func (c *Client) Verify(ctx context.Context, code string) (*VerifyPaymentRequestResponse, error) {
	return net.Get[types.PaymentRequest](ctx, c.Client, c.Secret, basePath+"/verify/"+code, c.BaseURL)
}
