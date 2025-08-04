package paymentrequests

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type FetchPaymentRequestResponse = types.Response[types.PaymentRequest]

func (c *Client) Fetch(ctx context.Context, idOrCode string) (*FetchPaymentRequestResponse, error) {
	return net.Get[types.PaymentRequest](ctx, c.Client, c.Secret, basePath+"/"+idOrCode, c.BaseURL)
}
