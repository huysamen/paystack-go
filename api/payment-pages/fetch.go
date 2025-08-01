package paymentpages

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// FetchPaymentPageResponse represents the response from fetching a payment page
type FetchPaymentPageResponse = types.Response[PaymentPage]

// Fetch gets details of a payment page on your integration
func (c *Client) Fetch(ctx context.Context, idOrSlug string) (*FetchPaymentPageResponse, error) {
	resp, err := net.Get[PaymentPage](
		ctx,
		c.client,
		c.secret,
		paymentPagesBasePath+"/"+idOrSlug,
		c.baseURL,
	)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
