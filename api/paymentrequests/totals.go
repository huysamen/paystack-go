package paymentrequests

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// PaymentRequestTotalsResponse represents the response from getting payment request totals
type PaymentRequestTotalsResponse = types.Response[PaymentRequestTotals]

// GetTotals gets payment requests metric
func (c *Client) GetTotals(ctx context.Context) (*PaymentRequestTotalsResponse, error) {
	return net.Get[PaymentRequestTotals](ctx, c.Client, c.Secret, basePath+"/totals", c.BaseURL)
}
