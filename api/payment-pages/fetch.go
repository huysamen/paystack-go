package paymentpages

import (
	"context"

	"github.com/huysamen/paystack-go/net"
)

// Fetch gets details of a payment page on your integration
func (c *Client) Fetch(ctx context.Context, idOrSlug string) (*PaymentPage, error) {
	if err := ValidateIDOrSlug(idOrSlug); err != nil {
		return nil, err
	}

	resp, err := net.Get[PaymentPage](
		ctx, c.client, c.secret, paymentPagesBasePath+"/"+idOrSlug, c.baseURL,
	)
	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}
