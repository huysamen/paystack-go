package paymentrequests

import (
	"context"

	"github.com/huysamen/paystack-go/net"
)

// PaymentRequestTotalsResponse represents the response from getting payment request totals
type PaymentRequestTotalsResponse struct {
	Status  bool                 `json:"status"`
	Message string               `json:"message"`
	Data    PaymentRequestTotals `json:"data"`
}

// GetTotals gets payment requests metric
func (c *Client) GetTotals(ctx context.Context) (*PaymentRequestTotals, error) {
	resp, err := net.Get[PaymentRequestTotals](
		ctx, c.client, c.secret, paymentRequestsBasePath+"/totals", c.baseURL,
	)
	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}
