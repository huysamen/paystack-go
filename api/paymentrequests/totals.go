package paymentrequests

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type PaymentRequestTotalsResponse = types.Response[types.PaymentRequestTotals]

func (c *Client) GetTotals(ctx context.Context) (*PaymentRequestTotalsResponse, error) {
	return net.Get[types.PaymentRequestTotals](ctx, c.Client, c.Secret, basePath+"/totals", c.BaseURL)
}
