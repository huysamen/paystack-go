package paymentpages

import (
	"context"

	"github.com/huysamen/paystack-go/net"
)

// Create creates a new payment page
func (c *Client) Create(ctx context.Context, req *CreatePaymentPageRequest) (*PaymentPage, error) {
	resp, err := net.Post[CreatePaymentPageRequest, PaymentPage](
		ctx, c.client, c.secret, paymentPagesBasePath, req, c.baseURL,
	)
	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}
