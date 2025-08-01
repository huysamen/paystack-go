package paymentrequests

import (
	"context"

	"github.com/huysamen/paystack-go/net"
)

// Create creates a payment request for a transaction on your integration
func (c *Client) Create(ctx context.Context, req *CreatePaymentRequestRequest) (*PaymentRequest, error) {

	resp, err := net.Post[CreatePaymentRequestRequest, PaymentRequest](
		ctx, c.client, c.secret, paymentRequestsBasePath, req, c.baseURL,
	)
	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}
