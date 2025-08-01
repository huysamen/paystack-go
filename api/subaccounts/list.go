package subaccounts

import (
	"context"
	"net/url"
	"strconv"
	"time"

	"github.com/huysamen/paystack-go/net"
)

// List retrieves a list of subaccounts
func (c *Client) List(ctx context.Context, req *SubaccountListRequest) (*SubaccountListResponse, error) {
	params := url.Values{}

	if req != nil {
		if req.PerPage != nil {
			params.Set("perPage", strconv.Itoa(*req.PerPage))
		}
		if req.Page != nil {
			params.Set("page", strconv.Itoa(*req.Page))
		}
		if req.From != nil {
			params.Set("from", req.From.Format(time.RFC3339))
		}
		if req.To != nil {
			params.Set("to", req.To.Format(time.RFC3339))
		}
	}

	endpoint := subaccountBasePath
	if len(params) > 0 {
		endpoint += "?" + params.Encode()
	}

	resp, err := net.Get[SubaccountListResponse](ctx, c.client, c.secret, endpoint, c.baseURL)
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

// ListWithBuilder retrieves a list of subaccounts using the builder pattern
func (c *Client) ListWithBuilder(ctx context.Context, builder *SubaccountListRequestBuilder) (*SubaccountListResponse, error) {
	if builder == nil {
		return c.List(ctx, nil)
	}
	return c.List(ctx, builder.Build())
}
