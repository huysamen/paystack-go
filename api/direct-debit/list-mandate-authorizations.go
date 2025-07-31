package directdebit

import (
	"context"
	"fmt"
	"net/url"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// ListMandateAuthorizations retrieves a list of direct debit mandate authorizations
func (c *Client) ListMandateAuthorizations(ctx context.Context, req *ListMandateAuthorizationsRequest) (*types.Response[[]MandateAuthorization], error) {
	endpoint := directDebitBasePath + "/mandate-authorizations"

	if req != nil {
		params := url.Values{}
		if req.Cursor != "" {
			params.Set("cursor", req.Cursor)
		}
		if req.Status != "" {
			params.Set("status", string(req.Status))
		}
		if req.PerPage > 0 {
			params.Set("per_page", fmt.Sprintf("%d", req.PerPage))
		}

		if len(params) > 0 {
			endpoint += "?" + params.Encode()
		}
	}

	resp, err := net.Get[[]MandateAuthorization](
		ctx, c.client, c.secret, endpoint, c.baseURL,
	)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
