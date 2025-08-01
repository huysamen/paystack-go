package paymentrequests

import (
	"context"

	"github.com/huysamen/paystack-go/net"
)

// Finalize finalizes a draft payment request
func (c *Client) Finalize(ctx context.Context, code string, req *FinalizePaymentRequestRequest) (*PaymentRequest, error) {

	resp, err := net.Post[FinalizePaymentRequestRequest, PaymentRequest](
		ctx, c.client, c.secret, paymentRequestsBasePath+"/finalize/"+code, req, c.baseURL,
	)
	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}

// FinalizeWithBuilder finalizes a draft payment request using the builder pattern
func (c *Client) FinalizeWithBuilder(ctx context.Context, code string, builder *FinalizePaymentRequestRequestBuilder) (*PaymentRequest, error) {
	if builder == nil {
		return c.Finalize(ctx, code, &FinalizePaymentRequestRequest{})
	}
	return c.Finalize(ctx, code, builder.Build())
}
