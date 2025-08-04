package paymentrequests

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// FetchPaymentRequestResponse represents the response from fetching a payment request
type FetchPaymentRequestResponse = types.Response[types.PaymentRequest]

// Fetch gets details of a payment request on your integration
func (c *Client) Fetch(ctx context.Context, idOrCode string) (*FetchPaymentRequestResponse, error) {
	return net.Get[types.PaymentRequest](ctx, c.Client, c.Secret, basePath+"/"+idOrCode, c.BaseURL)
}
