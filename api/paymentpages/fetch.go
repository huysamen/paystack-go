package paymentpages

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type FetchResponseData = types.PaymentPage
type FetchResponse = types.Response[FetchResponseData]

func (c *Client) Fetch(ctx context.Context, idOrSlug string) (*FetchResponse, error) {
	return net.Get[FetchResponseData](ctx, c.Client, c.Secret, basePath+"/"+idOrSlug, c.BaseURL)
}
