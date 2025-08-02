package paymentrequests

import (
	"context"

	"github.com/huysamen/paystack-go/net"
)

// FetchPaymentRequestResponse represents the response from fetching a payment request
type FetchPaymentRequestResponse struct {
	Status  bool           `json:"status"`
	Message string         `json:"message"`
	Data    PaymentRequest `json:"data"`
}

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
