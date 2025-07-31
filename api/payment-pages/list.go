package paymentpages

import (
	"context"
	"net/url"
	"strconv"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// List retrieves payment pages available on your integration
func (c *Client) List(ctx context.Context, req *ListPaymentPagesRequest) (*types.Response[[]PaymentPage], error) {
	params := url.Values{}

	if req != nil {
		if req.PerPage > 0 {
			params.Set("perPage", strconv.Itoa(req.PerPage))
		}
		if req.Page > 0 {
			params.Set("page", strconv.Itoa(req.Page))
		}
		if req.From != "" {
			params.Set("from", req.From)
		}
		if req.To != "" {
			params.Set("to", req.To)
		}
	}

	endpoint := paymentPagesBasePath
	if len(params) > 0 {
		endpoint += "?" + params.Encode()
	}

	return net.Get[[]PaymentPage](
		ctx, c.client, c.secret, endpoint, c.baseURL,
	)
}
