package paymentrequests

import (
	"context"

	"github.com/huysamen/paystack-go/net"
)

// Update updates a payment request details on your integration
func (c *Client) Update(ctx context.Context, idOrCode string, req *UpdatePaymentRequestRequest) (*PaymentRequest, error) {
	if err := ValidateIDOrCode(idOrCode); err != nil {
		return nil, err
	}

	if err := ValidateUpdatePaymentRequestRequest(req); err != nil {
		return nil, err
	}

	resp, err := net.Put[UpdatePaymentRequestRequest, PaymentRequest](
		ctx, c.client, c.secret, paymentRequestsBasePath+"/"+idOrCode, req, c.baseURL,
	)
	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}
