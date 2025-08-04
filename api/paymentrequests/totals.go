package paymentrequests

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type TotalsResponseData = types.PaymentRequestTotals
type TotalsResponse = types.Response[TotalsResponseData]

func (c *Client) GetTotals(ctx context.Context) (*TotalsResponse, error) {
	return net.Get[TotalsResponseData](ctx, c.Client, c.Secret, basePath+"/totals", c.BaseURL)
}
