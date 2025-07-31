package paymentrequests

import (
	"context"

	"github.com/huysamen/paystack-go/net"
)

// Finalize finalizes a draft payment request
func (c *Client) Finalize(ctx context.Context, code string, req *FinalizePaymentRequestRequest) (*PaymentRequest, error) {
	if err := ValidateCode(code); err != nil {
		return nil, err
	}

	resp, err := net.Post[FinalizePaymentRequestRequest, PaymentRequest](
		ctx, c.client, c.secret, paymentRequestsBasePath+"/finalize/"+code, req, c.baseURL,
	)
	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}
