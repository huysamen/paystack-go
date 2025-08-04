package paymentrequests

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type VerifyResponseData = types.PaymentRequest
type VerifyResponse = types.Response[VerifyResponseData]

func (c *Client) Verify(ctx context.Context, code string) (*VerifyResponse, error) {
	return net.Get[VerifyResponseData](ctx, c.Client, c.Secret, basePath+"/verify/"+code, c.BaseURL)
}
