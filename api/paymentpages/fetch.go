package paymentpages

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type FetchPaymentPageResponse = types.Response[types.PaymentPage]

func (c *Client) Fetch(ctx context.Context, idOrSlug string) (*FetchPaymentPageResponse, error) {
	return net.Get[types.PaymentPage](ctx, c.Client, c.Secret, basePath+"/"+idOrSlug, c.BaseURL)
}
