package paymentpages

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// FetchPaymentPageResponse represents the response from fetching a payment page
type FetchPaymentPageResponse = types.Response[types.PaymentPage]

// Fetch gets details of a payment page on your integration
func (c *Client) Fetch(ctx context.Context, idOrSlug string) (*FetchPaymentPageResponse, error) {
	return net.Get[types.PaymentPage](ctx, c.Client, c.Secret, basePath+"/"+idOrSlug, c.BaseURL)
}
