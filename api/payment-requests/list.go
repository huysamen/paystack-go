package paymentrequests

import (
	"context"
	"net/url"
	"strconv"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// List retrieves payment requests available on your integration
func (c *Client) List(ctx context.Context, req *ListPaymentRequestsRequest) (*types.Response[[]PaymentRequest], error) {
	params := url.Values{}

	if req != nil {
		if req.PerPage > 0 {
			params.Set("perPage", strconv.Itoa(req.PerPage))
		}
		if req.Page > 0 {
			params.Set("page", strconv.Itoa(req.Page))
		}
		if req.Customer != "" {
			params.Set("customer", req.Customer)
		}
		if req.Status != "" {
			params.Set("status", req.Status)
		}
		if req.Currency != "" {
			params.Set("currency", req.Currency)
		}
		if req.IncludeArchive != "" {
			params.Set("include_archive", req.IncludeArchive)
		}
		if req.From != "" {
			params.Set("from", req.From)
		}
		if req.To != "" {
			params.Set("to", req.To)
		}
	}

	endpoint := paymentRequestsBasePath
	if len(params) > 0 {
		endpoint += "?" + params.Encode()
	}

	return net.Get[[]PaymentRequest](
		ctx, c.client, c.secret, endpoint, c.baseURL,
	)
}
