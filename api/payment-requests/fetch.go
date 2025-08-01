package paymentrequests

import (
	"context"

	"github.com/huysamen/paystack-go/net"
)

// Fetch gets details of a payment request on your integration
func (c *Client) Fetch(ctx context.Context, idOrCode string) (*PaymentRequest, error) {

	resp, err := net.Get[PaymentRequest](
		ctx, c.client, c.secret, paymentRequestsBasePath+"/"+idOrCode, c.baseURL,
	)
	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}
