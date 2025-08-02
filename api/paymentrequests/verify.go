package paymentrequests

import (
	"context"

	"github.com/huysamen/paystack-go/net"
)

// VerifyPaymentRequestResponse represents the response from verifying a payment request
type VerifyPaymentRequestResponse struct {
	Status  bool           `json:"status"`
	Message string         `json:"message"`
	Data    PaymentRequest `json:"data"`
}

// Verify verifies details of a payment request on your integration
func (c *Client) Verify(ctx context.Context, code string) (*PaymentRequest, error) {
	resp, err := net.Get[PaymentRequest](
		ctx, c.client, c.secret, paymentRequestsBasePath+"/verify/"+code, c.baseURL,
	)
	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}
