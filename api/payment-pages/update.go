package paymentpages

import (
	"context"

	"github.com/huysamen/paystack-go/net"
)

// Update updates a payment page details on your integration
func (c *Client) Update(ctx context.Context, idOrSlug string, req *UpdatePaymentPageRequest) (*PaymentPage, error) {
	if err := ValidateIDOrSlug(idOrSlug); err != nil {
		return nil, err
	}

	if err := ValidateUpdatePaymentPageRequest(req); err != nil {
		return nil, err
	}

	resp, err := net.Put[UpdatePaymentPageRequest, PaymentPage](
		ctx, c.client, c.secret, paymentPagesBasePath+"/"+idOrSlug, req, c.baseURL,
	)
	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}
