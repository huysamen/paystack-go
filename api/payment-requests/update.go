package paymentrequests

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
)

// Update updates a payment request details on your integration
func (c *Client) Update(ctx context.Context, idOrCode string, req *UpdatePaymentRequestRequest) (*PaymentRequest, error) {

	resp, err := net.Put[UpdatePaymentRequestRequest, PaymentRequest](
		ctx, c.client, c.secret, paymentRequestsBasePath+"/"+idOrCode, req, c.baseURL,
	)
	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}

// UpdateWithBuilder updates a payment request using the builder pattern
func (c *Client) UpdateWithBuilder(ctx context.Context, idOrCode string, builder *UpdatePaymentRequestRequestBuilder) (*PaymentRequest, error) {
	if builder == nil {
		return nil, fmt.Errorf("builder cannot be nil")
	}
	return c.Update(ctx, idOrCode, builder.Build())
}
